package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DailyCheckin 每日打卡表
type DailyCheckin struct {
	global.GVA_MODEL
	UserID      uint      `json:"userId" gorm:"not null;index;comment:用户ID"`
	CheckinDate time.Time `json:"checkinDate" gorm:"not null;index;comment:打卡日期"`
	MoodLevel   int       `json:"moodLevel" gorm:"not null;comment:情绪等级(1-5)"`
	Notes       string    `json:"notes" gorm:"type:text;comment:打卡备注"`
	Rewards     int       `json:"rewards" gorm:"default:0;comment:获得奖励经验值"`
	IsSuccess   bool      `json:"isSuccess" gorm:"default:true;comment:是否成功戒色"`
	Streak      int       `json:"streak" gorm:"default:0;comment:当前连续天数"`

	// 关联表
	User WxUser `json:"user" gorm:"foreignKey:UserID"`
}

// TableName 设置表名
func (DailyCheckin) TableName() string {
	return "nofap_daily_checkins"
}
