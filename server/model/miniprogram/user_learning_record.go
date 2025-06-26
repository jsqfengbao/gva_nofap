package miniprogram

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserLearningRecord 用户学习记录表
type UserLearningRecord struct {
	global.GVA_MODEL
	UserID      uint       `json:"userId" gorm:"not null;index;comment:用户ID"`
	ContentID   uint       `json:"contentId" gorm:"not null;index;comment:内容ID"`
	StartTime   time.Time  `json:"startTime" gorm:"not null;comment:开始学习时间"`
	EndTime     *time.Time `json:"endTime" gorm:"comment:结束学习时间"`
	Duration    int        `json:"duration" gorm:"default:0;comment:学习时长(秒)"`
	Progress    int        `json:"progress" gorm:"default:0;comment:学习进度(百分比)"`
	IsCompleted bool       `json:"isCompleted" gorm:"default:false;comment:是否完成"`
	IsLiked     bool       `json:"isLiked" gorm:"default:false;comment:是否点赞"`
	IsCollected bool       `json:"isCollected" gorm:"default:false;comment:是否收藏"`
	Rating      int        `json:"rating" gorm:"default:0;comment:评分(1-5)"`
	Notes       string     `json:"notes" gorm:"type:text;comment:学习笔记"`

	// 关联表
	User    WxUser          `json:"user" gorm:"foreignKey:UserID"`
	Content LearningContent `json:"content" gorm:"foreignKey:ContentID"`
}

// TableName 设置表名
func (UserLearningRecord) TableName() string {
	return "nofap_user_learning_records"
}
