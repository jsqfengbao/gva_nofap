package miniprogram

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AssessmentResult 评估结果表
type AssessmentResult struct {
	global.GVA_MODEL
	UserID     uint      `json:"userId" gorm:"not null;index;comment:用户ID"`
	TotalScore int       `json:"totalScore" gorm:"not null;comment:总分"`
	RiskLevel  int       `json:"riskLevel" gorm:"not null;comment:风险等级:1正常,2轻度,3中度,4重度,5严重"`
	Answers    string    `json:"answers" gorm:"type:text;comment:答案JSON"`
	TestDate   time.Time `json:"testDate" gorm:"not null;comment:测试日期"`
	TestType   int       `json:"testType" gorm:"default:1;comment:测试类型:1初次,2复评"`
	Duration   int       `json:"duration" gorm:"comment:测试耗时(秒)"`

	// 关联表
	User WxUser `json:"user" gorm:"foreignKey:UserID"`
}

// TableName 设置表名
func (AssessmentResult) TableName() string {
	return "nofap_assessment_results"
}
