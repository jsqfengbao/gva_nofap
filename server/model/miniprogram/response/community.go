package response

import (
	"time"
)

// PostItem 帖子项目
type PostItem struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Category     int       `json:"category"`
	CategoryName string    `json:"categoryName"`
	IsAnonymous  bool      `json:"isAnonymous"`
	ViewCount    int       `json:"viewCount"`
	LikeCount    int       `json:"likeCount"`
	CommentCount int       `json:"commentCount"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`

	// 用户信息
	UserID       uint   `json:"userId"`
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`

	// 交互状态
	IsLiked   bool `json:"isLiked"`   // 当前用户是否已点赞
	CanEdit   bool `json:"canEdit"`   // 是否可以编辑
	CanDelete bool `json:"canDelete"` // 是否可以删除
}

// PostListResponse 帖子列表响应
type PostListResponse struct {
	List     []PostItem `json:"list"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
	HasMore  bool       `json:"hasMore"`
}

// PostDetailResponse 帖子详情响应
type PostDetailResponse struct {
	PostItem
	Comments []CommentItem `json:"comments"`
}

// CommentItem 评论项目
type CommentItem struct {
	ID          uint      `json:"id"`
	PostID      uint      `json:"postId"`
	Content     string    `json:"content"`
	ParentID    uint      `json:"parentId"`
	IsAnonymous bool      `json:"isAnonymous"`
	LikeCount   int       `json:"likeCount"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`

	// 用户信息
	UserID       uint   `json:"userId"`
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`

	// 父评论信息
	ParentContent  string `json:"parentContent,omitempty"`
	ParentNickname string `json:"parentNickname,omitempty"`

	// 子评论
	Children []CommentItem `json:"children,omitempty"`

	// 交互状态
	IsLiked   bool `json:"isLiked"`   // 当前用户是否已点赞
	CanDelete bool `json:"canDelete"` // 是否可以删除
}

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	List     []CommentItem `json:"list"`
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	HasMore  bool          `json:"hasMore"`
}

// CreatePostResponse 创建帖子响应
type CreatePostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Status    int       `json:"status"`
	Message   string    `json:"message"`
}

// CreateCommentResponse 创建评论响应
type CreateCommentResponse struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"postId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Status    int       `json:"status"`
	Message   string    `json:"message"`
}

// LikeResponse 点赞响应
type LikeResponse struct {
	TargetID  uint   `json:"targetId"`
	LikeType  int    `json:"likeType"`
	IsLiked   bool   `json:"isLiked"`   // 点赞后的状态
	LikeCount int    `json:"likeCount"` // 点赞后的总数
	Message   string `json:"message"`
}

// CategoryStats 分类统计
type CategoryStats struct {
	Category     int    `json:"category"`
	CategoryName string `json:"categoryName"`
	PostCount    int64  `json:"postCount"`
}

// CommunityStatsResponse 社区统计响应
type CommunityStatsResponse struct {
	TotalPosts    int64           `json:"totalPosts"`
	TotalComments int64           `json:"totalComments"`
	TotalLikes    int64           `json:"totalLikes"`
	TodayPosts    int64           `json:"todayPosts"`
	CategoryStats []CategoryStats `json:"categoryStats"`
}

// UserPostsResponse 用户帖子响应
type UserPostsResponse struct {
	List     []PostItem `json:"list"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
	HasMore  bool       `json:"hasMore"`
}
