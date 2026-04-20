package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserAchievement 用户成就表
type UserAchievement struct {
	global.GVA_MODEL
	UserID        uint      `json:"userId" gorm:"not null;index;comment:用户ID"`
	AchievementID uint      `json:"achievementId" gorm:"not null;index;comment:成就ID"`
	UnlockedAt    time.Time `json:"unlockedAt" gorm:"not null;comment:解锁时间"`
	Progress      int       `json:"progress" gorm:"default:0;comment:完成进度"`
	IsNotified    bool      `json:"isNotified" gorm:"default:false;comment:是否已通知"`

	// 关联表
	User        WxUser      `json:"user" gorm:"foreignKey:UserID"`
	Achievement Achievement `json:"achievement" gorm:"foreignKey:AchievementID"`
}

// TableName 设置表名
func (UserAchievement) TableName() string {
	return "nofap_user_achievements"
}
