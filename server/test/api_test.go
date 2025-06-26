package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	miniprogramReq "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 初始化测试环境
func setupTestEnv() *gin.Engine {
	// 初始化配置
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()

	// 设置gin为测试模式
	gin.SetMode(gin.TestMode)

	// 初始化路由
	router := initialize.Routers()

	return router
}

// TestWxLogin 测试微信登录接口
func TestWxLogin(t *testing.T) {
	router := setupTestEnv()

	// 准备测试数据
	loginReq := miniprogramReq.WxLoginRequest{
		Code: "test_code_123",
	}

	jsonData, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/miniprogram/auth/wx-login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 执行请求
	router.ServeHTTP(w, req)

	// 验证响应
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 验证响应结构
	assert.Contains(t, response, "code")
	assert.Contains(t, response, "msg")
}

// TestGetUserProfile 测试获取用户资料接口
func TestGetUserProfile(t *testing.T) {
	router := setupTestEnv()

	req, _ := http.NewRequest("GET", "/miniprogram/user/profile", nil)
	req.Header.Set("Content-Type", "application/json")
	// 这里应该添加认证token，但为了测试简化先跳过

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 验证响应状态码
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestCheckinDaily 测试每日打卡接口
func TestCheckinDaily(t *testing.T) {
	router := setupTestEnv()

	checkinReq := miniprogramReq.CheckinRequest{
		MoodLevel: 4,
		Notes:     "今天感觉不错",
	}

	jsonData, _ := json.Marshal(checkinReq)
	req, _ := http.NewRequest("POST", "/miniprogram/checkin/daily", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetAchievements 测试获取成就列表接口
func TestGetAchievements(t *testing.T) {
	router := setupTestEnv()

	req, _ := http.NewRequest("GET", "/miniprogram/achievement/list", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestCreateCommunityPost 测试创建社区帖子接口
func TestCreateCommunityPost(t *testing.T) {
	router := setupTestEnv()

	postReq := miniprogramReq.CreatePostRequest{
		Content:     "这是一个测试帖子",
		Category:    "experience",
		IsAnonymous: false,
	}

	jsonData, _ := json.Marshal(postReq)
	req, _ := http.NewRequest("POST", "/miniprogram/community/posts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetLearningContent 测试获取学习内容接口
func TestGetLearningContent(t *testing.T) {
	router := setupTestEnv()

	req, _ := http.NewRequest("GET", "/miniprogram/learning/content?category=knowledge&page=1&pageSize=10", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestCreateEmergencyHelp 测试创建紧急求助接口
func TestCreateEmergencyHelp(t *testing.T) {
	router := setupTestEnv()

	helpReq := miniprogramReq.CreateEmergencyHelpRequest{
		Type:        "urgent_impulse",
		Description: "需要紧急帮助",
		IsAnonymous: true,
	}

	jsonData, _ := json.Marshal(helpReq)
	req, _ := http.NewRequest("POST", "/miniprogram/emergency/help", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestSubmitAssessment 测试提交评估接口
func TestSubmitAssessment(t *testing.T) {
	router := setupTestEnv()

	assessmentReq := miniprogramReq.SubmitAssessmentRequest{
		AssessmentType: "addiction",
		Answers: []miniprogramReq.AssessmentAnswer{
			{QuestionID: 1, Answer: "经常", Score: 3},
			{QuestionID: 2, Answer: "有时", Score: 2},
		},
	}

	jsonData, _ := json.Marshal(assessmentReq)
	req, _ := http.NewRequest("POST", "/miniprogram/assessment/submit", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// BenchmarkWxLogin 微信登录接口性能测试
func BenchmarkWxLogin(b *testing.B) {
	router := setupTestEnv()

	loginReq := miniprogramReq.WxLoginRequest{
		Code: "test_code_123",
	}
	jsonData, _ := json.Marshal(loginReq)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/miniprogram/auth/wx-login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

// BenchmarkGetUserProfile 用户资料接口性能测试
func BenchmarkGetUserProfile(b *testing.B) {
	router := setupTestEnv()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/miniprogram/user/profile", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}
