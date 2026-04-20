package response

import (
	"time"

	miniprogramModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
)

// AssessmentResultResponse 评估结果响应
type AssessmentResultResponse struct {
	ID              uint               `json:"id" comment:"评估结果ID"`
	TotalScore      int                `json:"totalScore" comment:"总分"`
	RiskLevel       string             `json:"riskLevel" comment:"风险等级"`
	CategoryScores  map[string]float64 `json:"categoryScores" comment:"各维度分数"`
	CompletedAt     time.Time          `json:"completedAt" comment:"完成时间"`
	Recommendations []string           `json:"recommendations" comment:"建议方案"`
}

// AssessmentHistoryResponse 评估历史响应
type AssessmentHistoryResponse struct {
	List     []miniprogramModel.AssessmentResult `json:"list" comment:"评估列表"`
	Total    int64                               `json:"total" comment:"总数"`
	Page     int                                 `json:"page" comment:"当前页"`
	PageSize int                                 `json:"pageSize" comment:"每页数量"`
}

// AssessmentStatsResponse 评估统计响应
type AssessmentStatsResponse struct {
	TotalAssessments int                  `json:"totalAssessments" comment:"总评估次数"`
	LatestScore      int                  `json:"latestScore" comment:"最新得分"`
	LatestRiskLevel  string               `json:"latestRiskLevel" comment:"最新风险等级"`
	ScoreTrend       []int                `json:"scoreTrend" comment:"分数趋势"`
	CategoryTrends   map[string][]float64 `json:"categoryTrends" comment:"各维度趋势"`
	LastAssessmentAt *time.Time           `json:"lastAssessmentAt" comment:"最后评估时间"`
}
