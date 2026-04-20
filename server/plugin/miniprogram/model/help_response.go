package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HelpResponse 求助响应表
type HelpResponse struct {
	global.GVA_MODEL
	HelpID       uint      `json:"helpId" gorm:"not null;index;comment:求助ID"`
	ResponderID  uint      `json:"responderId" gorm:"not null;index;comment:响应者ID"`
	Content      string    `json:"content" gorm:"type:text;not null;comment:响应内容"`
	ResponseType int       `json:"responseType" gorm:"default:1;comment:响应类型:1文字回复,2申请私聊,3专业建议"`
	IsVolunteer  bool      `json:"isVolunteer" gorm:"default:false;comment:是否志愿者"`
	Status       int       `json:"status" gorm:"default:1;comment:状态:1正常,2已删除"`
	HelpfulCount int       `json:"helpfulCount" gorm:"default:0;comment:有用评价数"`
	ResponseAt   time.Time `json:"responseAt" gorm:"not null;comment:响应时间"`

	// 关联表
	Help      EmergencyHelp `json:"help" gorm:"foreignKey:HelpID"`
	Responder WxUser        `json:"responder" gorm:"foreignKey:ResponderID"`
}

// TableName 设置表名
func (HelpResponse) TableName() string {
	return "nofap_help_responses"
}
