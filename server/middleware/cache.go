package middleware

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CacheConfig 缓存配置
type CacheConfig struct {
	Duration    time.Duration // 缓存时间
	Prefix      string        // 缓存前缀
	SkipPaths   []string      // 跳过缓存的路径
	OnlyMethods []string      // 只缓存特定方法
}

// ResponseCache 响应缓存中间件
func ResponseCache(config CacheConfig) gin.HandlerFunc {
	if config.Prefix == "" {
		config.Prefix = "api_cache:"
	}
	if len(config.OnlyMethods) == 0 {
		config.OnlyMethods = []string{"GET"}
	}
	if config.Duration == 0 {
		config.Duration = 5 * time.Minute
	}

	return func(c *gin.Context) {
		// 检查Redis是否可用
		if global.GVA_REDIS == nil {
			c.Next()
			return
		}

		// 检查是否跳过缓存
		if shouldSkipCache(c, config) {
			c.Next()
			return
		}

		// 生成缓存键
		cacheKey := generateCacheKey(c, config.Prefix)

		// 尝试从缓存获取响应
		cached, err := global.GVA_REDIS.Get(c, cacheKey).Result()
		if err == nil && cached != "" {
			// 缓存命中
			var cachedResponse CachedResponse
			if err := json.Unmarshal([]byte(cached), &cachedResponse); err == nil {
				// 设置响应头
				for key, value := range cachedResponse.Headers {
					c.Header(key, value)
				}
				c.Header("X-Cache", "HIT")
				c.Data(cachedResponse.StatusCode, cachedResponse.ContentType, cachedResponse.Body)
				return
			}
		}

		// 缓存未命中，创建响应写入器
		writer := &CacheWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		// 处理请求
		c.Next()

		// 检查是否应该缓存响应
		if shouldCacheResponse(c, writer) {
			cachedResponse := CachedResponse{
				StatusCode:  writer.Status(),
				ContentType: writer.Header().Get("Content-Type"),
				Headers:     make(map[string]string),
				Body:        writer.body.Bytes(),
			}

			// 复制重要的响应头
			for _, header := range []string{"Content-Type", "Content-Encoding", "ETag"} {
				if value := writer.Header().Get(header); value != "" {
					cachedResponse.Headers[header] = value
				}
			}

			// 序列化并存储到缓存
			if data, err := json.Marshal(cachedResponse); err == nil {
				global.GVA_REDIS.Set(c, cacheKey, string(data), config.Duration)
				global.GVA_LOG.Debug("缓存响应",
					zap.String("key", cacheKey),
					zap.Duration("duration", config.Duration))
			}
		}

		// 添加缓存状态头
		c.Header("X-Cache", "MISS")
	}
}

// CachedResponse 缓存的响应结构
type CachedResponse struct {
	StatusCode  int               `json:"status_code"`
	ContentType string            `json:"content_type"`
	Headers     map[string]string `json:"headers"`
	Body        []byte            `json:"body"`
}

// CacheWriter 缓存写入器
type CacheWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CacheWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func (w *CacheWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// generateCacheKey 生成缓存键
func generateCacheKey(c *gin.Context, prefix string) string {
	// 包含路径、查询参数、用户ID等信息
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	userID := ""

	// 获取用户ID（如果存在）
	if uid, exists := c.Get("userID"); exists {
		userID = fmt.Sprintf("%v", uid)
	}

	// 创建唯一标识
	identifier := fmt.Sprintf("%s?%s&user=%s", path, query, userID)

	// 使用MD5哈希
	hash := md5.Sum([]byte(identifier))
	hashString := hex.EncodeToString(hash[:])

	return prefix + hashString
}

// shouldSkipCache 检查是否应该跳过缓存
func shouldSkipCache(c *gin.Context, config CacheConfig) bool {
	// 检查请求方法
	methodAllowed := false
	for _, method := range config.OnlyMethods {
		if c.Request.Method == method {
			methodAllowed = true
			break
		}
	}
	if !methodAllowed {
		return true
	}

	// 检查跳过路径
	for _, path := range config.SkipPaths {
		if strings.HasPrefix(c.Request.URL.Path, path) {
			return true
		}
	}

	// 检查特殊头部
	if c.GetHeader("Cache-Control") == "no-cache" {
		return true
	}

	return false
}

// shouldCacheResponse 检查是否应该缓存响应
func shouldCacheResponse(c *gin.Context, writer *CacheWriter) bool {
	// 只缓存成功响应
	if writer.Status() < 200 || writer.Status() >= 300 {
		return false
	}

	// 检查响应大小（不缓存过大的响应）
	if writer.body.Len() > 1024*1024 { // 1MB
		return false
	}

	// 检查Content-Type
	contentType := writer.Header().Get("Content-Type")
	if strings.Contains(contentType, "application/json") ||
		strings.Contains(contentType, "text/") {
		return true
	}

	return false
}

// ETagMiddleware ETag中间件
func ETagMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 只对GET请求处理ETag
		if c.Request.Method != "GET" {
			c.Next()
			return
		}

		// 创建ETag写入器
		writer := &ETagWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		// 处理请求
		c.Next()

		// 生成ETag
		if writer.body.Len() > 0 {
			hash := md5.Sum(writer.body.Bytes())
			etag := `"` + hex.EncodeToString(hash[:]) + `"`

			// 检查If-None-Match头
			if ifNoneMatch := c.GetHeader("If-None-Match"); ifNoneMatch == etag {
				c.Status(http.StatusNotModified)
				return
			}

			// 设置ETag头
			c.Header("ETag", etag)
		}
	}
}

// ETagWriter ETag写入器
type ETagWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *ETagWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func (w *ETagWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// CompressionMiddleware 压缩中间件
func CompressionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查客户端是否支持压缩
		if !strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		// 检查Content-Type是否适合压缩
		contentType := c.GetHeader("Content-Type")
		if !shouldCompress(contentType) {
			c.Next()
			return
		}

		// 这里可以集成gzip中间件
		// 由于gin本身支持gzip，这里只做检查
		c.Next()
	}
}

// shouldCompress 检查是否应该压缩
func shouldCompress(contentType string) bool {
	compressibleTypes := []string{
		"text/",
		"application/json",
		"application/javascript",
		"application/xml",
		"application/rss+xml",
		"application/atom+xml",
	}

	for _, t := range compressibleTypes {
		if strings.HasPrefix(contentType, t) {
			return true
		}
	}
	return false
}

// CacheInvalidation 缓存失效中间件
func CacheInvalidation(patterns []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 对于修改操作，清除相关缓存
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			if global.GVA_REDIS != nil {
				for _, pattern := range patterns {
					if strings.HasPrefix(c.Request.URL.Path, pattern) {
						// 清除匹配的缓存键
						keys, err := global.GVA_REDIS.Keys(c, "api_cache:*").Result()
						if err == nil {
							for _, key := range keys {
								if shouldInvalidateKey(key, pattern) {
									global.GVA_REDIS.Del(c, key)
								}
							}
						}
						break
					}
				}
			}
		}
		c.Next()
	}
}

// shouldInvalidateKey 检查是否应该失效缓存键
func shouldInvalidateKey(key, pattern string) bool {
	// 简单的模式匹配，实际可以更复杂
	return strings.Contains(key, pattern)
}

// MemoryCache 内存缓存中间件（用于小数据量的快速缓存）
type MemoryCacheItem struct {
	Data      []byte
	ExpiresAt time.Time
	Headers   map[string]string
}

var memoryCache = make(map[string]*MemoryCacheItem)

func MemoryCache(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 只缓存GET请求
		if c.Request.Method != "GET" {
			c.Next()
			return
		}

		// 生成缓存键
		cacheKey := "mem:" + c.Request.URL.Path + "?" + c.Request.URL.RawQuery

		// 检查内存缓存
		if item, exists := memoryCache[cacheKey]; exists {
			if time.Now().Before(item.ExpiresAt) {
				// 缓存命中
				for key, value := range item.Headers {
					c.Header(key, value)
				}
				c.Header("X-Memory-Cache", "HIT")
				c.Data(200, "application/json", item.Data)
				return
			} else {
				// 缓存过期，删除
				delete(memoryCache, cacheKey)
			}
		}

		// 创建写入器
		writer := &MemoryCacheWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		// 处理请求
		c.Next()

		// 存储到内存缓存
		if writer.Status() == 200 && writer.body.Len() < 64*1024 { // 只缓存小于64KB的响应
			headers := make(map[string]string)
			headers["Content-Type"] = writer.Header().Get("Content-Type")

			memoryCache[cacheKey] = &MemoryCacheItem{
				Data:      writer.body.Bytes(),
				ExpiresAt: time.Now().Add(duration),
				Headers:   headers,
			}
		}

		c.Header("X-Memory-Cache", "MISS")
	}
}

// MemoryCacheWriter 内存缓存写入器
type MemoryCacheWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *MemoryCacheWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func (w *MemoryCacheWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
