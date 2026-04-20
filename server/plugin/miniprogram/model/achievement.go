package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Achievement 成就系统表
type Achievement struct {
	global.GVA_MODEL
	Name         string `json:"name" gorm:"size:100;not null;comment:成就名称"`
	Description  string `json:"description" gorm:"size:500;comment:成就描述"`
	IconUrl      string `json:"iconUrl" gorm:"type:text;comment:成就图标URL"`
	Category     int    `json:"category" gorm:"not null;comment:成就分类:1打卡,2等级,3社区,4学习,5特殊"`
	Type         int    `json:"type" gorm:"not null;comment:成就类型:1累计,2连续,3一次性"`
	Condition    string `json:"condition" gorm:"type:text;comment:解锁条件JSON"`
	Rewards      int    `json:"rewards" gorm:"default:0;comment:奖励经验值"`
	Rarity       int    `json:"rarity" gorm:"default:1;comment:稀有度:1普通,2稀有,3史诗,4传说"`
	DisplayOrder int    `json:"displayOrder" gorm:"default:0;comment:显示顺序"`
	IsActive     bool   `json:"isActive" gorm:"default:true;comment:是否启用"`

	// 关联表
	UserAchievements []UserAchievement `json:"userAchievements" gorm:"foreignKey:AchievementID"`
}

// TableName 设置表名
func (Achievement) TableName() string {
	return "nofap_achievements"
}
