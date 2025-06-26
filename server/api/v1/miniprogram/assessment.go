package miniprogram

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	miniprogramModel "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram"
	miniprogramReq "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/request"
	miniprogramRes "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/response"
	miniprogramService "github.com/flipped-aurora/gin-vue-admin/server/service/miniprogram"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AssessmentApi struct{}

var assessmentService = new(miniprogramService.AssessmentService)

// SubmitAssessment 提交评估结果
// @Tags 小程序-评估系统
// @Summary 提交评估结果
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body miniprogramReq.SubmitAssessmentRequest true "评估数据"
// @Success 200 {object} response.Response{data=miniprogramRes.AssessmentResultResponse,msg=string} "提交成功"
// @Router /miniprogram/assessment/submit [post]
func (a *AssessmentApi) SubmitAssessment(c *gin.Context) {
	var req miniprogramReq.SubmitAssessmentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从JWT中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未认证", c)
		return
	}

	// 转换RiskLevel为int
	riskLevelMap := map[string]int{
		"正常": 1, "轻度": 2, "中度": 3, "重度": 4, "严重": 5,
		"low": 1, "mild": 2, "moderate": 3, "high": 4, "severe": 5,
	}
	riskLevel := riskLevelMap[req.RiskLevel]
	if riskLevel == 0 {
		riskLevel = 1 // 默认正常
	}

	// 序列化答案为JSON字符串
	answersJSON, err := json.Marshal(req.Answers)
	if err != nil {
		response.FailWithMessage("答案数据格式错误", c)
		return
	}

	// 创建评估结果记录
	assessmentResult := miniprogramModel.AssessmentResult{
		UserID:     userID.(uint),
		TotalScore: req.TotalScore,
		RiskLevel:  riskLevel,
		Answers:    string(answersJSON),
		TestDate:   time.Now(),
		TestType:   1, // 默认初次评估
	}

	// 保存到数据库
	err = assessmentService.CreateAssessmentResult(&assessmentResult)
	if err != nil {
		global.GVA_LOG.Error("保存评估结果失败", zap.Error(err))
		response.FailWithMessage("保存评估结果失败", c)
		return
	}

	// 更新用户的最后评估时间
	err = assessmentService.UpdateUserLastAssessment(userID.(uint), time.Now())
	if err != nil {
		global.GVA_LOG.Error("更新用户评估时间失败", zap.Error(err))
		// 不影响主流程，仅记录日志
	}

	// 返回评估结果
	result := miniprogramRes.AssessmentResultResponse{
		ID:              assessmentResult.ID,
		TotalScore:      assessmentResult.TotalScore,
		RiskLevel:       req.RiskLevel, // 返回原始字符串
		CompletedAt:     assessmentResult.TestDate,
		Recommendations: generateRecommendations(req.RiskLevel),
	}

	response.OkWithDetailed(result, "评估提交成功", c)
}

// GetAssessmentHistory 获取评估历史
// @Tags 小程序-评估系统
// @Summary 获取用户评估历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} response.Response{data=miniprogramRes.AssessmentHistoryResponse,msg=string} "获取成功"
// @Router /miniprogram/assessment/history [get]
func (a *AssessmentApi) GetAssessmentHistory(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未认证", c)
		return
	}

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 获取评估历史
	history, total, err := assessmentService.GetAssessmentHistory(userID.(uint), page, pageSize)
	if err != nil {
		global.GVA_LOG.Error("获取评估历史失败", zap.Error(err))
		response.FailWithMessage("获取评估历史失败", c)
		return
	}

	// 构造响应数据
	result := miniprogramRes.AssessmentHistoryResponse{
		List:     history,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	response.OkWithDetailed(result, "获取评估历史成功", c)
}

// GetLatestAssessment 获取最新评估结果
// @Tags 小程序-评估系统
// @Summary 获取用户最新评估结果
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=miniprogramRes.AssessmentResultResponse,msg=string} "获取成功"
// @Router /miniprogram/assessment/latest [get]
func (a *AssessmentApi) GetLatestAssessment(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未认证", c)
		return
	}

	// 获取最新评估结果
	assessment, err := assessmentService.GetLatestAssessment(userID.(uint))
	if err != nil {
		global.GVA_LOG.Error("获取最新评估结果失败", zap.Error(err))
		response.FailWithMessage("获取最新评估结果失败", c)
		return
	}

	if assessment == nil {
		response.OkWithDetailed(nil, "暂无评估记录", c)
		return
	}

	// 转换风险等级为字符串
	riskLevelMap := map[int]string{
		1: "正常", 2: "轻度", 3: "中度", 4: "重度", 5: "严重",
	}
	riskLevelStr := riskLevelMap[assessment.RiskLevel]
	if riskLevelStr == "" {
		riskLevelStr = "正常"
	}

	// 构造响应数据
	result := miniprogramRes.AssessmentResultResponse{
		ID:              assessment.ID,
		TotalScore:      assessment.TotalScore,
		RiskLevel:       riskLevelStr,
		CompletedAt:     assessment.TestDate,
		Recommendations: generateRecommendations(riskLevelStr),
	}

	response.OkWithDetailed(result, "获取最新评估结果成功", c)
}

// GetAssessmentStats 获取评估统计数据
// @Tags 小程序-评估系统
// @Summary 获取用户评估统计数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=miniprogramRes.AssessmentStatsResponse,msg=string} "获取成功"
// @Router /miniprogram/assessment/stats [get]
func (a *AssessmentApi) GetAssessmentStats(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未认证", c)
		return
	}

	// 获取评估统计数据
	stats, err := assessmentService.GetAssessmentStats(userID.(uint))
	if err != nil {
		global.GVA_LOG.Error("获取评估统计数据失败", zap.Error(err))
		response.FailWithMessage("获取评估统计数据失败", c)
		return
	}

	response.OkWithDetailed(stats, "获取评估统计数据成功", c)
}

// generateRecommendations 根据风险等级生成建议
func generateRecommendations(riskLevel string) []string {
	switch riskLevel {
	case "low":
		return []string{
			"继续保持健康的生活方式",
			"定期进行自我评估",
			"培养多样化的兴趣爱好",
			"保持规律的作息时间",
		}
	case "mild":
		return []string{
			"开始记录和监控相关行为",
			"寻找健康的替代活动",
			"建立支持网络",
			"学习压力管理技巧",
			"考虑寻求专业指导",
		}
	case "moderate":
		return []string{
			"制定详细的康复计划",
			"寻求专业心理咨询师帮助",
			"参加支持小组",
			"建立问责制度",
			"开始行为认知疗法",
			"使用戒色助手的各项功能",
		}
	case "high":
		return []string{
			"立即寻求专业医疗帮助",
			"联系心理健康专家",
			"考虑住院或密集治疗",
			"通知信任的家人或朋友",
			"使用危机干预热线",
			"建立24小时支持系统",
		}
	default:
		return []string{
			"建议进行完整评估",
			"寻求专业指导",
		}
	}
}
