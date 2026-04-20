package service

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	miniprogramModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
	miniprogramRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model/response"
)

type AssessmentService struct{}

// CreateAssessmentResult 创建评估结果
func (assessmentService *AssessmentService) CreateAssessmentResult(assessment *miniprogramModel.AssessmentResult) error {
	return global.GVA_DB.Create(assessment).Error
}

// GetAssessmentHistory 获取用户评估历史
func (assessmentService *AssessmentService) GetAssessmentHistory(userID uint, page, pageSize int) ([]miniprogramModel.AssessmentResult, int64, error) {
	var assessments []miniprogramModel.AssessmentResult
	var total int64

	db := global.GVA_DB.Model(&miniprogramModel.AssessmentResult{}).Where("user_id = ?", userID)

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	err = db.Order("test_date DESC").Offset(offset).Limit(pageSize).Find(&assessments).Error
	if err != nil {
		return nil, 0, err
	}

	return assessments, total, nil
}

// GetLatestAssessment 获取最新评估结果
func (assessmentService *AssessmentService) GetLatestAssessment(userID uint) (*miniprogramModel.AssessmentResult, error) {
	var assessment miniprogramModel.AssessmentResult

	err := global.GVA_DB.Where("user_id = ?", userID).
		Order("test_date DESC").
		First(&assessment).Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}

	return &assessment, nil
}

// GetAssessmentStats 获取评估统计数据
func (assessmentService *AssessmentService) GetAssessmentStats(userID uint) (*miniprogramRes.AssessmentStatsResponse, error) {
	var stats miniprogramRes.AssessmentStatsResponse

	// 获取总评估次数
	var totalCount int64
	err := global.GVA_DB.Model(&miniprogramModel.AssessmentResult{}).
		Where("user_id = ?", userID).
		Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	stats.TotalAssessments = int(totalCount)

	// 获取最新评估结果
	latest, err := assessmentService.GetLatestAssessment(userID)
	if err != nil {
		return nil, err
	}

	if latest != nil {
		stats.LatestScore = latest.TotalScore
		// 转换风险等级为字符串
		riskLevelMap := map[int]string{
			1: "正常",
			2: "轻度",
			3: "中度",
			4: "重度",
			5: "严重",
		}
		stats.LatestRiskLevel = riskLevelMap[latest.RiskLevel]
		if stats.LatestRiskLevel == "" {
			stats.LatestRiskLevel = "正常"
		}
		stats.LastAssessmentAt = &latest.TestDate
	}

	// 获取最近10次评估的分数趋势
	var recentAssessments []miniprogramModel.AssessmentResult
	err = global.GVA_DB.Where("user_id = ?", userID).
		Order("test_date DESC").
		Limit(10).
		Find(&recentAssessments).Error
	if err != nil {
		return nil, err
	}

	// 构建分数趋势（按时间正序）
	for i := len(recentAssessments) - 1; i >= 0; i-- {
		stats.ScoreTrend = append(stats.ScoreTrend, recentAssessments[i].TotalScore)
	}

	// 注意：CategoryTrends功能暂时移除，因为AssessmentResult模型中没有CategoryScores字段
	stats.CategoryTrends = make(map[string][]float64)

	return &stats, nil
}

// UpdateUserLastAssessment 更新用户最后评估时间
func (assessmentService *AssessmentService) UpdateUserLastAssessment(userID uint, assessmentTime time.Time) error {
	return global.GVA_DB.Model(&miniprogramModel.WxUser{}).
		Where("id = ?", userID).
		Update("last_assessment_at", assessmentTime).Error
}
