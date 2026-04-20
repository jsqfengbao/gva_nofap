package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CommunityLike 社区点赞表
type CommunityLike struct {
	global.GVA_MODEL
	PostID   uint `json:"postId" gorm:"not null;index;comment:动态ID"`
	UserID   uint `json:"userId" gorm:"not null;index;comment:用户ID"`
	LikeType int  `json:"likeType" gorm:"default:1;comment:点赞类型:1动态,2评论"`
	TargetID uint `json:"targetId" gorm:"not null;comment:目标ID(动态ID或评论ID)"`

	// 关联表
	User WxUser        `json:"user" gorm:"foreignKey:UserID"`
	Post CommunityPost `json:"post" gorm:"foreignKey:PostID"`
}

// TableName 设置表名
func (CommunityLike) TableName() string {
	return "nofap_community_likes"
}
