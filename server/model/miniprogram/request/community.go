package request

// CreatePostRequest 创建帖子请求
type CreatePostRequest struct {
	Title       string `json:"title" binding:"required,max=200" message:"标题不能为空且长度不超过200字符"`
	Content     string `json:"content" binding:"required,max=5000" message:"内容不能为空且长度不超过5000字符"`
	Category    int    `json:"category" binding:"required,min=1,max=4" message:"分类必须在1-4之间"`
	IsAnonymous bool   `json:"isAnonymous"`
}

// UpdatePostRequest 更新帖子请求
type UpdatePostRequest struct {
	Title       string `json:"title" binding:"required,max=200" message:"标题不能为空且长度不超过200字符"`
	Content     string `json:"content" binding:"required,max=5000" message:"内容不能为空且长度不超过5000字符"`
	Category    int    `json:"category" binding:"required,min=1,max=4" message:"分类必须在1-4之间"`
	IsAnonymous bool   `json:"isAnonymous"`
}

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	PostID      uint   `json:"postId" binding:"required" message:"帖子ID不能为空"`
	Content     string `json:"content" binding:"required,max=1000" message:"评论内容不能为空且长度不超过1000字符"`
	ParentID    uint   `json:"parentId"` // 0表示顶级评论
	IsAnonymous bool   `json:"isAnonymous"`
}

// GetPostsRequest 获取帖子列表请求
type GetPostsRequest struct {
	Page     int    `form:"page,default=1" binding:"min=1"`
	PageSize int    `form:"pageSize,default=20" binding:"min=1,max=50"`
	Category int    `form:"category" binding:"min=0,max=4"` // 0表示全部分类
	SortBy   string `form:"sortBy,default=created_at"`      // created_at, like_count, comment_count, view_count
	Order    string `form:"order,default=desc"`             // asc, desc
}

// GetCommentsRequest 获取评论列表请求
type GetCommentsRequest struct {
	PostID   uint `form:"postId" binding:"required" message:"帖子ID不能为空"`
	Page     int  `form:"page,default=1" binding:"min=1"`
	PageSize int  `form:"pageSize,default=20" binding:"min=1,max=50"`
}

// LikeRequest 点赞/取消点赞请求
type LikeRequest struct {
	TargetID uint `json:"targetId" binding:"required" message:"目标ID不能为空"`
	LikeType int  `json:"likeType" binding:"required,min=1,max=2" message:"点赞类型必须是1(帖子)或2(评论)"`
}

// ReportRequest 举报请求
type ReportRequest struct {
	TargetID   uint   `json:"targetId" binding:"required" message:"目标ID不能为空"`
	TargetType int    `json:"targetType" binding:"required,min=1,max=2" message:"举报类型必须是1(帖子)或2(评论)"`
	Reason     string `json:"reason" binding:"required,max=500" message:"举报原因不能为空且长度不超过500字符"`
}
