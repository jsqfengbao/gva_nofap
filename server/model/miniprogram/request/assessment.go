package request

// SubmitAssessmentRequest 提交评估请求
type SubmitAssessmentRequest struct {
	TotalScore     int                    `json:"totalScore" binding:"required,min=0,max=200" comment:"总分"`
	RiskLevel      string                 `json:"riskLevel" binding:"required,oneof=low mild moderate high" comment:"风险等级"`
	CategoryScores map[string]float64     `json:"categoryScores" binding:"required" comment:"各维度分数"`
	Answers        map[string]interface{} `json:"answers" binding:"required" comment:"题目答案"`
}

// GetAssessmentHistoryRequest 获取评估历史请求
type GetAssessmentHistoryRequest struct {
	Page     int `form:"page" binding:"min=1" comment:"页码"`
	PageSize int `form:"pageSize" binding:"min=1,max=100" comment:"每页数量"`
}
