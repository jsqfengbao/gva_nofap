package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// CreateEmergencyHelpRequest 创建紧急求助请求
type CreateEmergencyHelpRequest struct {
	Type        int    `json:"type" binding:"required,min=1,max=4" form:"type"`                   // 求助类型(1-紧急冲动 2-情绪低落 3-复发担忧 4-其他)
	Priority    int    `json:"priority" binding:"omitempty,min=1,max=4" form:"priority"`          // 优先级(1-低 2-中 3-高 4-紧急)
	Description string `json:"description" binding:"required,min=10,max=1000" form:"description"` // 描述说明
	Location    string `json:"location" binding:"omitempty,max=200" form:"location"`              // 地理位置(可选)
	IsAnonymous bool   `json:"isAnonymous" form:"isAnonymous"`                                    // 是否匿名求助
}

// UpdateEmergencyHelpRequest 更新紧急求助请求
type UpdateEmergencyHelpRequest struct {
	ID          uint   `json:"id" binding:"required" form:"id"`                                    // 求助ID
	Status      int    `json:"status" binding:"omitempty,min=1,max=4" form:"status"`               // 状态
	Description string `json:"description" binding:"omitempty,min=10,max=1000" form:"description"` // 描述说明
	Rating      *int   `json:"rating" binding:"omitempty,min=1,max=5" form:"rating"`               // 用户评分
	Feedback    string `json:"feedback" binding:"omitempty,max=500" form:"feedback"`               // 用户反馈
}

// GetEmergencyHelpsRequest 获取紧急求助列表请求
type GetEmergencyHelpsRequest struct {
	request.PageInfo
	UserID   uint `json:"userId" form:"userId"`     // 用户ID(管理员查看所有)
	Type     int  `json:"type" form:"type"`         // 求助类型筛选
	Status   int  `json:"status" form:"status"`     // 状态筛选
	Priority int  `json:"priority" form:"priority"` // 优先级筛选
}

// CreateEmergencyResponseRequest 创建紧急求助响应请求
type CreateEmergencyResponseRequest struct {
	HelpID  uint   `json:"helpId" binding:"required" form:"helpId"`                  // 求助记录ID
	Type    int    `json:"type" binding:"required,min=1,max=4" form:"type"`          // 响应类型(1-文字回复 2-语音通话 3-视频通话 4-陪伴聊天)
	Content string `json:"content" binding:"required,min=1,max=1000" form:"content"` // 响应内容
}

// UpdateEmergencyResponseRequest 更新紧急求助响应请求
type UpdateEmergencyResponseRequest struct {
	ID       uint       `json:"id" binding:"required" form:"id"`                     // 响应ID
	Content  string     `json:"content" binding:"omitempty,max=1000" form:"content"` // 响应内容
	Duration int        `json:"duration" form:"duration"`                            // 持续时间(分钟)
	IsActive bool       `json:"isActive" form:"isActive"`                            // 是否活跃
	EndedAt  *time.Time `json:"endedAt" form:"endedAt"`                              // 结束时间
}

// EmergencyVolunteerRequest 志愿者注册/更新请求
type EmergencyVolunteerRequest struct {
	Specialties string `json:"specialties" binding:"omitempty,max=500" form:"specialties"` // 专长领域
	MaxLoad     int    `json:"maxLoad" binding:"omitempty,min=1,max=10" form:"maxLoad"`    // 最大同时服务数
}

// UpdateVolunteerStatusRequest 更新志愿者状态请求
type UpdateVolunteerStatusRequest struct {
	IsOnline bool `json:"isOnline" form:"isOnline"` // 是否在线
}

// GetEmergencyResourcesRequest 获取紧急求助资源请求
type GetEmergencyResourcesRequest struct {
	request.PageInfo
	Type       int    `json:"type" form:"type"`             // 资源类型筛选
	Difficulty int    `json:"difficulty" form:"difficulty"` // 难度等级筛选
	Tags       string `json:"tags" form:"tags"`             // 标签筛选
}

// CreateEmergencyResourceRequest 创建紧急求助资源请求
type CreateEmergencyResourceRequest struct {
	Title       string `json:"title" binding:"required,min=2,max=200" form:"title"`          // 标题
	Type        int    `json:"type" binding:"required,min=1,max=5" form:"type"`              // 资源类型
	Content     string `json:"content" binding:"omitempty,max=2000" form:"content"`          // 内容
	MediaUrl    string `json:"mediaUrl" binding:"omitempty,max=500" form:"mediaUrl"`         // 媒体文件URL
	Duration    int    `json:"duration" form:"duration"`                                     // 时长(秒)
	Difficulty  int    `json:"difficulty" binding:"omitempty,min=1,max=5" form:"difficulty"` // 难度等级
	Tags        string `json:"tags" binding:"omitempty,max=500" form:"tags"`                 // 标签
	Description string `json:"description" binding:"omitempty,max=1000" form:"description"`  // 描述
}

// UpdateEmergencyResourceRequest 更新紧急求助资源请求
type UpdateEmergencyResourceRequest struct {
	ID          uint   `json:"id" binding:"required" form:"id"`                              // 资源ID
	Title       string `json:"title" binding:"omitempty,min=2,max=200" form:"title"`         // 标题
	Content     string `json:"content" binding:"omitempty,max=2000" form:"content"`          // 内容
	MediaUrl    string `json:"mediaUrl" binding:"omitempty,max=500" form:"mediaUrl"`         // 媒体文件URL
	Duration    int    `json:"duration" form:"duration"`                                     // 时长(秒)
	Difficulty  int    `json:"difficulty" binding:"omitempty,min=1,max=5" form:"difficulty"` // 难度等级
	IsActive    bool   `json:"isActive" form:"isActive"`                                     // 是否启用
	Tags        string `json:"tags" binding:"omitempty,max=500" form:"tags"`                 // 标签
	Description string `json:"description" binding:"omitempty,max=1000" form:"description"`  // 描述
}

// RateResourceRequest 资源评分请求
type RateResourceRequest struct {
	ResourceID uint `json:"resourceId" binding:"required" form:"resourceId"`     // 资源ID
	Rating     int  `json:"rating" binding:"required,min=1,max=5" form:"rating"` // 评分(1-5)
}

// ConnectVolunteerRequest 连接志愿者请求
type ConnectVolunteerRequest struct {
	HelpID uint `json:"helpId" binding:"required" form:"helpId"` // 求助记录ID
}

// SearchEmergencyRequest 搜索紧急求助记录请求
type SearchEmergencyRequest struct {
	request.PageInfo
	Keyword   string     `json:"keyword" form:"keyword"`     // 关键词搜索
	StartTime *time.Time `json:"startTime" form:"startTime"` // 开始时间
	EndTime   *time.Time `json:"endTime" form:"endTime"`     // 结束时间
	Type      int        `json:"type" form:"type"`           // 类型筛选
	Status    int        `json:"status" form:"status"`       // 状态筛选
	Priority  int        `json:"priority" form:"priority"`   // 优先级筛选
}
