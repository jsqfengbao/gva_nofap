package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CommunityComment 社区评论表
type CommunityComment struct {
	global.GVA_MODEL
	PostID      uint       `json:"postId" gorm:"not null;index;comment:动态ID"`
	UserID      uint       `json:"userId" gorm:"not null;index;comment:用户ID"`
	Content     string     `json:"content" gorm:"type:text;not null;comment:评论内容"`
	ParentID    uint       `json:"parentId" gorm:"index;comment:父评论ID,0为顶级评论"`
	IsAnonymous bool       `json:"isAnonymous" gorm:"default:false;comment:是否匿名"`
	LikeCount   int        `json:"likeCount" gorm:"default:0;comment:点赞数"`
	Status      int        `json:"status" gorm:"default:1;comment:状态:0删除,1正常,2待审核"`
	AuditAt     *time.Time `json:"auditAt" gorm:"comment:审核时间"`

	// 关联表
	User     WxUser             `json:"user" gorm:"foreignKey:UserID"`
	Post     CommunityPost      `json:"post" gorm:"foreignKey:PostID"`
	Parent   *CommunityComment  `json:"parent" gorm:"foreignKey:ParentID"`
	Children []CommunityComment `json:"children" gorm:"foreignKey:ParentID"`
}

// TableName 设置表名
func (CommunityComment) TableName() string {
	return "nofap_community_comments"
}
