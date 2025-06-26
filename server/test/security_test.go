package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 初始化测试环境
func setupSecurityTestEnv() *gin.Engine {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// 添加安全中间件
	router.Use(middleware.SecurityHeaders())
	router.Use(middleware.XSSProtection())
	router.Use(middleware.SQLInjectionProtection())
	router.Use(middleware.RequestSizeLimit(1024 * 1024)) // 1MB限制

	return router
}

// TestXSSProtection 测试XSS防护
func TestXSSProtection(t *testing.T) {
	router := setupSecurityTestEnv()

	// 添加测试路由
	router.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 测试XSS攻击载荷
	xssPayloads := []string{
		`<script>alert('xss')</script>`,
		`javascript:alert('xss')`,
		`<img src=x onerror=alert('xss')>`,
		`<iframe src="javascript:alert('xss')"></iframe>`,
		`<svg onload=alert('xss')>`,
	}

	for _, payload := range xssPayloads {
		reqBody := map[string]string{
			"content": payload,
		}

		jsonData, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// XSS攻击应该被拦截
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "请求包含不安全内容")
	}
}

// TestSQLInjectionProtection 测试SQL注入防护
func TestSQLInjectionProtection(t *testing.T) {
	router := setupSecurityTestEnv()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 测试SQL注入攻击载荷
	sqlPayloads := []string{
		"1' OR '1'='1",
		"'; DROP TABLE users; --",
		"UNION SELECT * FROM users",
		"1; DELETE FROM users",
		"admin'--",
	}

	for _, payload := range sqlPayloads {
		req, _ := http.NewRequest("GET", "/test?id="+payload, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// SQL注入攻击应该被拦截
		assert.NotEqual(t, http.StatusOK, w.Code)
	}
}

// TestSecurityHeaders 测试安全响应头
func TestSecurityHeaders(t *testing.T) {
	router := setupSecurityTestEnv()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 检查安全头
	assert.Equal(t, "DENY", w.Header().Get("X-Frame-Options"))
	assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"))
	assert.Equal(t, "1; mode=block", w.Header().Get("X-XSS-Protection"))
	assert.Contains(t, w.Header().Get("Content-Security-Policy"), "default-src 'self'")
	assert.Equal(t, "strict-origin-when-cross-origin", w.Header().Get("Referrer-Policy"))
}

// TestRequestSizeLimit 测试请求大小限制
func TestRequestSizeLimit(t *testing.T) {
	router := setupSecurityTestEnv()

	router.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 创建超过限制大小的请求
	largeData := strings.Repeat("a", 2*1024*1024) // 2MB数据
	req, _ := http.NewRequest("POST", "/test", strings.NewReader(largeData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 应该被拦截
	assert.NotEqual(t, http.StatusOK, w.Code)
}

// TestDataEncryption 测试数据加密
func TestDataEncryption(t *testing.T) {
	testData := "sensitive-information-123"

	// 测试加密
	encrypted, err := utils.AESEncrypt(testData)
	assert.NoError(t, err)
	assert.NotEmpty(t, encrypted)
	assert.NotEqual(t, testData, encrypted)

	// 测试解密
	decrypted, err := utils.AESDecrypt(encrypted)
	assert.NoError(t, err)
	assert.Equal(t, testData, decrypted)
}

// TestDataMasking 测试数据脱敏
func TestDataMasking(t *testing.T) {
	// 测试手机号脱敏
	phone := "13812345678"
	masked := utils.MaskSensitiveData(phone, "phone")
	assert.Equal(t, "138****5678", masked)

	// 测试邮箱脱敏
	email := "test@example.com"
	masked = utils.MaskSensitiveData(email, "email")
	assert.Equal(t, "te***t@example.com", masked)

	// 测试姓名脱敏
	name := "张三丰"
	masked = utils.MaskSensitiveData(name, "name")
	assert.Equal(t, "张*丰", masked)
}

// TestPasswordHashing 测试密码哈希
func TestPasswordHashing(t *testing.T) {
	password := "MySecurePassword123!"

	// 测试密码哈希
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)

	// 测试密码验证
	valid := utils.CheckPassword(password, hash)
	assert.True(t, valid)

	// 测试错误密码
	invalid := utils.CheckPassword("WrongPassword", hash)
	assert.False(t, invalid)
}

// TestInputSanitization 测试输入清理
func TestInputSanitization(t *testing.T) {
	// 测试HTML实体转义
	input := `<script>alert("xss")</script>`
	sanitized := utils.SanitizeInput(input)
	assert.Contains(t, sanitized, "&lt;script&gt;")
	assert.Contains(t, sanitized, "&quot;")

	// 测试控制字符移除
	input = "test\x00\x1fdata"
	sanitized = utils.SanitizeInput(input)
	assert.Equal(t, "testdata", sanitized)
}

// TestInputValidation 测试输入验证
func TestInputValidation(t *testing.T) {
	// 测试正常输入
	valid := utils.IsValidInput("normal text", 100)
	assert.True(t, valid)

	// 测试长度超限
	valid = utils.IsValidInput("very long text", 5)
	assert.False(t, valid)

	// 测试危险模式
	dangerous := []string{
		"<script>alert('xss')</script>",
		"javascript:alert('xss')",
		"union select * from users",
		"drop table users",
	}

	for _, input := range dangerous {
		valid = utils.IsValidInput(input, 1000)
		assert.False(t, valid, "Input should be invalid: %s", input)
	}
}

// TestRateLimiting 测试速率限制
func TestRateLimiting(t *testing.T) {
	router := setupSecurityTestEnv()

	// 添加速率限制中间件
	router.Use(middleware.DefaultLimit())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 发送多个请求测试速率限制
	for i := 0; i < 5; i++ {
		req, _ := http.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "127.0.0.1:8080" // 模拟同一IP

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// 前几个请求应该成功
		if i < 3 {
			assert.Equal(t, http.StatusOK, w.Code)
		}
	}
}

// BenchmarkEncryption 加密性能测试
func BenchmarkEncryption(b *testing.B) {
	testData := "benchmark test data for encryption performance"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encrypted, _ := utils.AESEncrypt(testData)
		utils.AESDecrypt(encrypted)
	}
}

// BenchmarkHashing 哈希性能测试
func BenchmarkHashing(b *testing.B) {
	password := "benchmark password"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash, _ := utils.HashPassword(password)
		utils.CheckPassword(password, hash)
	}
}

// BenchmarkInputValidation 输入验证性能测试
func BenchmarkInputValidation(b *testing.B) {
	input := "normal input text for validation benchmark"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		utils.IsValidInput(input, 1000)
		utils.SanitizeInput(input)
	}
}
