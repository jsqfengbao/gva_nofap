package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// CreateLearningContentRequest 创建学习内容请求
type CreateLearningContentRequest struct {
	Title       string `json:"title" binding:"required,min=2,max=200" form:"title"`          // 标题
	Type        int    `json:"type" binding:"required,min=1,max=5" form:"type"`              // 内容类型
	Category    int    `json:"category" binding:"required,min=1,max=10" form:"category"`     // 分类
	Content     string `json:"content" binding:"omitempty,max=5000" form:"content"`          // 内容
	MediaUrl    string `json:"mediaUrl" binding:"omitempty,max=500" form:"mediaUrl"`         // 媒体文件URL
	Duration    int    `json:"duration" form:"duration"`                                     // 时长(秒)
	Difficulty  int    `json:"difficulty" binding:"omitempty,min=1,max=5" form:"difficulty"` // 难度等级
	Tags        string `json:"tags" binding:"omitempty,max=500" form:"tags"`                 // 标签
	Description string `json:"description" binding:"omitempty,max=1000" form:"description"`  // 描述
}

// UpdateLearningContentRequest 更新学习内容请求
type UpdateLearningContentRequest struct {
	ID          uint   `json:"id" binding:"required" form:"id"`                              // 内容ID
	Title       string `json:"title" binding:"omitempty,min=2,max=200" form:"title"`         // 标题
	Content     string `json:"content" binding:"omitempty,max=5000" form:"content"`          // 内容
	MediaUrl    string `json:"mediaUrl" binding:"omitempty,max=500" form:"mediaUrl"`         // 媒体文件URL
	Duration    int    `json:"duration" form:"duration"`                                     // 时长(秒)
	Difficulty  int    `json:"difficulty" binding:"omitempty,min=1,max=5" form:"difficulty"` // 难度等级
	IsActive    bool   `json:"isActive" form:"isActive"`                                     // 是否启用
	Tags        string `json:"tags" binding:"omitempty,max=500" form:"tags"`                 // 标签
	Description string `json:"description" binding:"omitempty,max=1000" form:"description"`  // 描述
}

// CreateLearningProgressRequest 创建学习进度请求
type CreateLearningProgressRequest struct {
	ContentID   uint `json:"contentId" binding:"required" form:"contentId"`             // 学习内容ID
	Progress    int  `json:"progress" binding:"required,min=0,max=100" form:"progress"` // 学习进度(0-100)
	Duration    int  `json:"duration" form:"duration"`                                  // 学习时长(秒)
	IsCompleted bool `json:"isCompleted" form:"isCompleted"`                            // 是否完成
}

// GetLearningContentsRequest 获取学习内容列表请求
type GetLearningContentsRequest struct {
	request.PageInfo
	Type       int    `json:"type" form:"type"`             // 内容类型筛选
	Category   int    `json:"category" form:"category"`     // 分类筛选
	Difficulty int    `json:"difficulty" form:"difficulty"` // 难度等级筛选
	Tags       string `json:"tags" form:"tags"`             // 标签筛选
	IsActive   *bool  `json:"isActive" form:"isActive"`     // 是否启用筛选
}

// RateLearningContentRequest 学习内容评分请求
type RateLearningContentRequest struct {
	ContentID uint   `json:"contentId" binding:"required" form:"contentId"`       // 内容ID
	Rating    int    `json:"rating" binding:"required,min=1,max=5" form:"rating"` // 评分(1-5)
	Comment   string `json:"comment" binding:"omitempty,max=500" form:"comment"`  // 评论
}

// GetLearningProgressRequest 获取学习进度请求
type GetLearningProgressRequest struct {
	UserID      uint  `json:"userId" form:"userId"`           // 用户ID(管理员查看)
	ContentID   uint  `json:"contentId" form:"contentId"`     // 内容ID筛选
	IsCompleted *bool `json:"isCompleted" form:"isCompleted"` // 是否完成筛选
}

// CreateLearningPlanRequest 创建学习计划请求
type CreateLearningPlanRequest struct {
	Title       string     `json:"title" binding:"required,min=2,max=200" form:"title"`         // 计划标题
	Description string     `json:"description" binding:"omitempty,max=1000" form:"description"` // 计划描述
	StartDate   *time.Time `json:"startDate" binding:"required" form:"startDate"`               // 开始日期
	EndDate     *time.Time `json:"endDate" binding:"required" form:"endDate"`                   // 结束日期
	ContentIDs  []uint     `json:"contentIds" binding:"required,min=1" form:"contentIds"`       // 学习内容ID列表
}

// UpdateLearningPlanRequest 更新学习计划请求
type UpdateLearningPlanRequest struct {
	ID          uint       `json:"id" binding:"required" form:"id"`                             // 计划ID
	Title       string     `json:"title" binding:"omitempty,min=2,max=200" form:"title"`        // 计划标题
	Description string     `json:"description" binding:"omitempty,max=1000" form:"description"` // 计划描述
	StartDate   *time.Time `json:"startDate" form:"startDate"`                                  // 开始日期
	EndDate     *time.Time `json:"endDate" form:"endDate"`                                      // 结束日期
	Status      int        `json:"status" binding:"omitempty,min=1,max=4" form:"status"`        // 状态
	ContentIDs  []uint     `json:"contentIds" form:"contentIds"`                                // 学习内容ID列表
}

// GetLearningPlansRequest 获取学习计划列表请求
type GetLearningPlansRequest struct {
	request.PageInfo
	UserID uint `json:"userId" form:"userId"` // 用户ID(管理员查看)
	Status int  `json:"status" form:"status"` // 状态筛选
}
