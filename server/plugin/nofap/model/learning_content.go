package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LearningContent 学习内容表
type LearningContent struct {
	global.GVA_MODEL
	Title        string     `json:"title" gorm:"size:200;not null;comment:标题"`
	Summary      string     `json:"summary" gorm:"size:500;comment:摘要"`
	Content      string     `json:"content" gorm:"type:longtext;not null;comment:内容"`
	ContentType  int        `json:"contentType" gorm:"not null;comment:内容类型:1文章,2视频,3音频"`
	Category     int        `json:"category" gorm:"not null;comment:分类:1科普知识,2康复指导,3心理健康,4经验分享"`
	Difficulty   int        `json:"difficulty" gorm:"default:1;comment:难度等级:1入门,2初级,3中级,4高级"`
	Duration     int        `json:"duration" gorm:"comment:时长(分钟)"`
	ThumbnailUrl string     `json:"thumbnailUrl" gorm:"type:text;comment:缩略图URL"`
	MediaUrl     string     `json:"mediaUrl" gorm:"type:text;comment:媒体文件URL"`
	Author       string     `json:"author" gorm:"size:100;comment:作者"`
	ViewCount    int        `json:"viewCount" gorm:"default:0;comment:观看次数"`
	LikeCount    int        `json:"likeCount" gorm:"default:0;comment:点赞数"`
	CollectCount int        `json:"collectCount" gorm:"default:0;comment:收藏数"`
	CommentCount int        `json:"commentCount" gorm:"default:0;comment:评论数"`
	Status       int        `json:"status" gorm:"default:1;comment:状态:0删除,1正常,2草稿"`
	PublishAt    *time.Time `json:"publishAt" gorm:"comment:发布时间"`
	Tags         string     `json:"tags" gorm:"size:500;comment:标签,逗号分隔"`

	// 关联表
	LearningRecords []UserLearningRecord `json:"learningRecords" gorm:"foreignKey:ContentID"`
}

// TableName 设置表名
func (LearningContent) TableName() string {
	return "nofap_learning_contents"
}
