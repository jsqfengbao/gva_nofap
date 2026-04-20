package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CommunityPost 社区动态表
type CommunityPost struct {
	global.GVA_MODEL
	UserID       uint       `json:"userId" gorm:"not null;index;comment:用户ID"`
	Title        string     `json:"title" gorm:"size:200;not null;comment:标题"`
	Content      string     `json:"content" gorm:"type:text;not null;comment:内容"`
	Category     int        `json:"category" gorm:"not null;comment:分类:1经验分享,2求助求鼓励,3日常打卡,4成功故事"`
	IsAnonymous  bool       `json:"isAnonymous" gorm:"default:false;comment:是否匿名"`
	ViewCount    int        `json:"viewCount" gorm:"default:0;comment:查看次数"`
	LikeCount    int        `json:"likeCount" gorm:"default:0;comment:点赞数"`
	CommentCount int        `json:"commentCount" gorm:"default:0;comment:评论数"`
	Status       int        `json:"status" gorm:"default:1;comment:状态:0删除,1正常,2待审核,3已拒绝"`
	AuditAt      *time.Time `json:"auditAt" gorm:"comment:审核时间"`
	AuditBy      uint       `json:"auditBy" gorm:"comment:审核人ID"`
	AuditReason  string     `json:"auditReason" gorm:"size:500;comment:审核原因"`

	// 关联表
	User     WxUser             `json:"user" gorm:"foreignKey:UserID"`
	Comments []CommunityComment `json:"comments" gorm:"foreignKey:PostID"`
	Likes    []CommunityLike    `json:"likes" gorm:"foreignKey:PostID"`
}

// TableName 设置表名
func (CommunityPost) TableName() string {
	return "nofap_community_posts"
}
