package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AbstinenceRecord 戒色记录表
type AbstinenceRecord struct {
	global.GVA_MODEL
	UserID        uint       `json:"userId" gorm:"not null;index;comment:用户ID"`
	StartDate     time.Time  `json:"startDate" gorm:"not null;comment:开始戒色日期"`
	CurrentStreak int        `json:"currentStreak" gorm:"default:0;comment:当前连续天数"`
	LongestStreak int        `json:"longestStreak" gorm:"default:0;comment:最长连续天数"`
	TotalDays     int        `json:"totalDays" gorm:"default:0;comment:总戒色天数"`
	SuccessRate   float64    `json:"successRate" gorm:"type:decimal(5,2);default:0;comment:成功率百分比"`
	Level         int        `json:"level" gorm:"default:1;comment:等级(1-50)"`
	Experience    int        `json:"experience" gorm:"default:0;comment:经验值"`
	LastRelapseAt *time.Time `json:"lastRelapseAt" gorm:"comment:最后复发时间"`
	Status        int        `json:"status" gorm:"default:1;comment:状态:0结束,1进行中"`

	// 关联表
	User          WxUser         `json:"user" gorm:"foreignKey:UserID"`
	DailyCheckins []DailyCheckin `json:"dailyCheckins" gorm:"foreignKey:UserID"`
}

// TableName 设置表名
func (AbstinenceRecord) TableName() string {
	return "nofap_abstinence_records"
}
