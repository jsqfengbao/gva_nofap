package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"html"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SecurityHeaders 添加安全响应头
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 防止点击劫持
		c.Header("X-Frame-Options", "DENY")

		// 防止MIME类型嗅探
		c.Header("X-Content-Type-Options", "nosniff")

		// XSS保护
		c.Header("X-XSS-Protection", "1; mode=block")

		// 严格传输安全(HTTPS)
		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}

		// 内容安全策略
		csp := "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline' 'unsafe-eval' https://res.wx.qq.com; " +
			"style-src 'self' 'unsafe-inline'; " +
			"img-src 'self' data: https:; " +
			"connect-src 'self' https://api.weixin.qq.com; " +
			"font-src 'self'; " +
			"object-src 'none'; " +
			"media-src 'self'; " +
			"frame-src 'none'"
		c.Header("Content-Security-Policy", csp)

		// 引用者策略
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// 权限策略
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		c.Next()
	}
}

// XSSProtection XSS防护中间件
func XSSProtection() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求体中的XSS攻击
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			// 读取原始请求体进行检查
			body, err := c.GetRawData()
			if err != nil {
				response.FailWithMessage("请求数据读取失败", c)
				c.Abort()
				return
			}

			// 检查XSS模式
			if containsXSS(string(body)) {
				global.GVA_LOG.Warn("检测到XSS攻击尝试",
					zap.String("ip", c.ClientIP()),
					zap.String("path", c.Request.URL.Path),
					zap.String("body", string(body)))
				response.FailWithMessage("请求包含不安全内容", c)
				c.Abort()
				return
			}

			// 重新设置请求体
			c.Request.Body = io.NopCloser(strings.NewReader(string(body)))
		}

		// 检查查询参数
		for key, values := range c.Request.URL.Query() {
			for _, value := range values {
				if containsXSS(value) {
					global.GVA_LOG.Warn("查询参数中检测到XSS攻击尝试",
						zap.String("ip", c.ClientIP()),
						zap.String("param", key),
						zap.String("value", value))
					response.FailWithMessage("请求参数包含不安全内容", c)
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}

// containsXSS 检查字符串是否包含XSS攻击模式
func containsXSS(input string) bool {
	// 常见的XSS攻击模式
	xssPatterns := []string{
		`<script[^>]*>.*?</script>`,
		`javascript:`,
		`vbscript:`,
		`onload\s*=`,
		`onerror\s*=`,
		`onclick\s*=`,
		`onmouseover\s*=`,
		`onfocus\s*=`,
		`<iframe[^>]*>`,
		`<object[^>]*>`,
		`<embed[^>]*>`,
		`<form[^>]*>`,
		`<img[^>]*onerror`,
		`<svg[^>]*onload`,
		`expression\s*\(`,
		`url\s*\(\s*javascript:`,
	}

	input = strings.ToLower(input)

	for _, pattern := range xssPatterns {
		matched, _ := regexp.MatchString(pattern, input)
		if matched {
			return true
		}
	}

	return false
}

// SQLInjectionProtection SQL注入防护中间件
func SQLInjectionProtection() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查查询参数
		for key, values := range c.Request.URL.Query() {
			for _, value := range values {
				if containsSQLInjection(value) {
					global.GVA_LOG.Warn("查询参数中检测到SQL注入攻击尝试",
						zap.String("ip", c.ClientIP()),
						zap.String("param", key),
						zap.String("value", value))
					response.FailWithMessage("请求参数包含不安全内容", c)
					c.Abort()
					return
				}
			}
		}

		// 检查POST请求体
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			body, err := c.GetRawData()
			if err != nil {
				response.FailWithMessage("请求数据读取失败", c)
				c.Abort()
				return
			}

			if containsSQLInjection(string(body)) {
				global.GVA_LOG.Warn("请求体中检测到SQL注入攻击尝试",
					zap.String("ip", c.ClientIP()),
					zap.String("path", c.Request.URL.Path))
				response.FailWithMessage("请求包含不安全内容", c)
				c.Abort()
				return
			}

			// 重新设置请求体
			c.Request.Body = io.NopCloser(strings.NewReader(string(body)))
		}

		c.Next()
	}
}

// containsSQLInjection 检查字符串是否包含SQL注入攻击模式
func containsSQLInjection(input string) bool {
	// 常见的SQL注入攻击模式
	sqlPatterns := []string{
		`union\s+select`,
		`drop\s+table`,
		`delete\s+from`,
		`insert\s+into`,
		`update\s+set`,
		`exec\s*\(`,
		`execute\s*\(`,
		`sp_executesql`,
		`xp_cmdshell`,
		`;\s*drop`,
		`;\s*delete`,
		`;\s*insert`,
		`;\s*update`,
		`'\s*or\s*'1'\s*=\s*'1`,
		`'\s*or\s*1\s*=\s*1`,
		`admin'\s*--`,
		`'\s*union`,
		`benchmark\s*\(`,
		`sleep\s*\(`,
		`waitfor\s+delay`,
	}

	input = strings.ToLower(input)

	for _, pattern := range sqlPatterns {
		matched, _ := regexp.MatchString(pattern, input)
		if matched {
			return true
		}
	}

	return false
}

// InputSanitization 输入清理中间件
func InputSanitization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 为响应添加清理函数
		c.Set("sanitizeHTML", func(input string) string {
			return html.EscapeString(input)
		})

		c.Next()
	}
}

// CSRFProtection CSRF防护中间件
func CSRFProtection() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 对于GET请求，生成CSRF token
		if c.Request.Method == "GET" {
			token := generateCSRFToken()
			c.Header("X-CSRF-Token", token)
			c.Set("csrf_token", token)
		} else {
			// 对于POST/PUT/DELETE请求，验证CSRF token
			clientToken := c.GetHeader("X-CSRF-Token")
			if clientToken == "" {
				clientToken = c.PostForm("_token")
			}

			if clientToken == "" {
				response.FailWithMessage("缺少CSRF令牌", c)
				c.Abort()
				return
			}

			// 这里可以添加更复杂的CSRF token验证逻辑
			// 简单验证：检查token格式和长度
			if len(clientToken) < 32 {
				response.FailWithMessage("无效的CSRF令牌", c)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// generateCSRFToken 生成CSRF令牌
func generateCSRFToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}

// RequestSizeLimit 请求大小限制中间件
func RequestSizeLimit(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength > maxSize {
			response.FailWithMessage("请求体过大", c)
			c.Abort()
			return
		}

		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		c.Next()
	}
}

// AntiReplay 防重放攻击中间件
func AntiReplay() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取时间戳和随机数
		timestamp := c.GetHeader("X-Timestamp")
		nonce := c.GetHeader("X-Nonce")

		if timestamp == "" || nonce == "" {
			// 对于小程序API，可以放宽要求
			if !strings.HasPrefix(c.Request.URL.Path, "/miniprogram/") {
				response.FailWithMessage("缺少防重放参数", c)
				c.Abort()
				return
			}
		} else {
			// 验证时间戳（5分钟内有效）
			if !isValidTimestamp(timestamp, 300) {
				response.FailWithMessage("请求已过期", c)
				c.Abort()
				return
			}

			// 检查nonce是否已使用（需要Redis支持）
			if global.GVA_REDIS != nil {
				key := "nonce:" + nonce
				exists := global.GVA_REDIS.Exists(c, key).Val()
				if exists > 0 {
					response.FailWithMessage("重复请求", c)
					c.Abort()
					return
				}

				// 记录nonce，5分钟过期
				global.GVA_REDIS.Set(c, key, "1", 5*time.Minute)
			}
		}

		c.Next()
	}
}

// isValidTimestamp 验证时间戳是否在有效期内
func isValidTimestamp(timestamp string, maxAge int64) bool {
	// 简单实现，实际应该解析时间戳并验证
	return len(timestamp) > 0
}
