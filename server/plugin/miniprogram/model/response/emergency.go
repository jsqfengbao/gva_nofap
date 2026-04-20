package response

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model"
)

// EmergencyHelpItem 紧急求助项目
type EmergencyHelpItem struct {
	ID                uint       `json:"id"`
	Type              int        `json:"type"`
	TypeName          string     `json:"typeName"`
	Status            int        `json:"status"`
	StatusName        string     `json:"statusName"`
	Priority          int        `json:"priority"`
	PriorityName      string     `json:"priorityName"`
	Description       string     `json:"description"`
	Location          string     `json:"location"`
	IsAnonymous       bool       `json:"isAnonymous"`
	UserNickname      string     `json:"userNickname"`
	VolunteerNickname string     `json:"volunteerNickname,omitempty"`
	ResponseCount     int        `json:"responseCount"`
	Rating            *int       `json:"rating"`
	Feedback          string     `json:"feedback"`
	CreatedAt         time.Time  `json:"createdAt"`
	RespondedAt       *time.Time `json:"respondedAt,omitempty"`
	ResolvedAt        *time.Time `json:"resolvedAt,omitempty"`
}

// EmergencyHelpListResponse 紧急求助列表响应
type EmergencyHelpListResponse struct {
	List     []EmergencyHelpItem `json:"list"`
	Total    int64               `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"pageSize"`
	HasMore  bool                `json:"hasMore"`
}

// EmergencyHelpDetailResponse 紧急求助详情响应
type EmergencyHelpDetailResponse struct {
	EmergencyHelpItem
	Responses  []EmergencyResponseItem `json:"responses"`
	CanRespond bool                    `json:"canRespond"`
	IsOwner    bool                    `json:"isOwner"`
}

// EmergencyResponseItem 紧急求助响应项目
type EmergencyResponseItem struct {
	ID                uint       `json:"id"`
	Type              int        `json:"type"`
	TypeName          string     `json:"typeName"`
	Content           string     `json:"content"`
	Duration          int        `json:"duration"`
	IsActive          bool       `json:"isActive"`
	VolunteerNickname string     `json:"volunteerNickname"`
	VolunteerAvatar   string     `json:"volunteerAvatar"`
	CreatedAt         time.Time  `json:"createdAt"`
	EndedAt           *time.Time `json:"endedAt,omitempty"`
}

// EmergencyVolunteerItem 紧急求助志愿者项目
type EmergencyVolunteerItem struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"userId"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Status       int       `json:"status"`
	StatusName   string    `json:"statusName"`
	IsOnline     bool      `json:"isOnline"`
	Specialties  string    `json:"specialties"`
	MaxLoad      int       `json:"maxLoad"`
	CurrentLoad  int       `json:"currentLoad"`
	TotalHelped  int       `json:"totalHelped"`
	AvgRating    float64   `json:"avgRating"`
	RatingCount  int       `json:"ratingCount"`
	LastActiveAt time.Time `json:"lastActiveAt"`
}

// VolunteerListResponse 志愿者列表响应
type VolunteerListResponse struct {
	List     []EmergencyVolunteerItem `json:"list"`
	Total    int64                    `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
	HasMore  bool                     `json:"hasMore"`
}

// OnlineVolunteersResponse 在线志愿者响应
type OnlineVolunteersResponse struct {
	Count      int                      `json:"count"`
	Volunteers []EmergencyVolunteerItem `json:"volunteers"`
}

// EmergencyResourceItem 紧急求助资源项目
type EmergencyResourceItem struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Type        int       `json:"type"`
	TypeName    string    `json:"typeName"`
	Content     string    `json:"content,omitempty"`
	MediaUrl    string    `json:"mediaUrl,omitempty"`
	Duration    int       `json:"duration"`
	Difficulty  int       `json:"difficulty"`
	IsActive    bool      `json:"isActive"`
	UsageCount  int       `json:"usageCount"`
	AvgRating   float64   `json:"avgRating"`
	RatingCount int       `json:"ratingCount"`
	Tags        string    `json:"tags"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

// ResourceListResponse 资源列表响应
type ResourceListResponse struct {
	List     []EmergencyResourceItem `json:"list"`
	Total    int64                   `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
	HasMore  bool                    `json:"hasMore"`
}

// EmergencyStatsResponse 紧急求助统计响应
type EmergencyStatsResponse struct {
	TotalHelps        int64          `json:"totalHelps"`        // 总求助数
	PendingHelps      int64          `json:"pendingHelps"`      // 待响应数
	InProgressHelps   int64          `json:"inProgressHelps"`   // 进行中数
	ResolvedHelps     int64          `json:"resolvedHelps"`     // 已解决数
	OnlineVolunteers  int64          `json:"onlineVolunteers"`  // 在线志愿者数
	AvgResponseTime   float64        `json:"avgResponseTime"`   // 平均响应时间(分钟)
	AvgResolutionTime float64        `json:"avgResolutionTime"` // 平均解决时间(分钟)
	HelpsByType       []TypeStat     `json:"helpsByType"`       // 按类型统计
	HelpsByPriority   []PriorityStat `json:"helpsByPriority"`   // 按优先级统计
}

// TypeStat 按类型统计
type TypeStat struct {
	Type     int    `json:"type"`
	TypeName string `json:"typeName"`
	Count    int64  `json:"count"`
}

// PriorityStat 按优先级统计
type PriorityStat struct {
	Priority     int    `json:"priority"`
	PriorityName string `json:"priorityName"`
	Count        int64  `json:"count"`
}

// ConnectVolunteerResponse 连接志愿者响应
type ConnectVolunteerResponse struct {
	Success   bool                    `json:"success"`
	Message   string                  `json:"message"`
	Volunteer *EmergencyVolunteerItem `json:"volunteer,omitempty"`
	Response  *EmergencyResponseItem  `json:"response,omitempty"`
}

// VolunteerMatchResponse 志愿者匹配响应
type VolunteerMatchResponse struct {
	HelpID        uint                   `json:"helpId"`
	Volunteer     EmergencyVolunteerItem `json:"volunteer"`
	EstimatedWait int                    `json:"estimatedWait"` // 预计等待时间(分钟)
}

// EmergencyDashboardResponse 紧急求助仪表板响应
type EmergencyDashboardResponse struct {
	Stats                EmergencyStatsResponse   `json:"stats"`
	RecentHelps          []EmergencyHelpItem      `json:"recentHelps"`
	RecommendedResources []EmergencyResourceItem  `json:"recommendedResources"`
	OnlineVolunteers     OnlineVolunteersResponse `json:"onlineVolunteers"`
}

// ConvertToEmergencyHelpItem 转换为求助项目响应
func ConvertToEmergencyHelpItem(help model.EmergencyHelp) EmergencyHelpItem {
	item := EmergencyHelpItem{
		ID:            help.ID,
		Type:          help.Type,
		TypeName:      help.GetTypeName(),
		Status:        help.Status,
		StatusName:    help.GetStatusName(),
		Priority:      help.Priority,
		PriorityName:  help.GetPriorityName(),
		Description:   help.Description,
		Location:      help.Location,
		IsAnonymous:   help.IsAnonymous,
		UserNickname:  "匿名用户",
		ResponseCount: len(help.Responses),
		Rating:        help.Rating,
		Feedback:      help.Feedback,
		CreatedAt:     help.CreatedAt,
		RespondedAt:   help.RespondedAt,
		ResolvedAt:    help.ResolvedAt,
	}

	// 如果不是匿名且有用户信息
	if !help.IsAnonymous && help.User.Nickname != "" {
		item.UserNickname = help.User.Nickname
	}

	// 如果有志愿者响应
	if help.Volunteer != nil && help.Volunteer.Nickname != "" {
		item.VolunteerNickname = help.Volunteer.Nickname
	}

	return item
}

// ConvertToEmergencyResourceItem 转换为资源项目响应
func ConvertToEmergencyResourceItem(resource model.EmergencyResource) EmergencyResourceItem {
	return EmergencyResourceItem{
		ID:          resource.ID,
		Title:       resource.Title,
		Type:        resource.Type,
		TypeName:    resource.GetResourceTypeName(),
		Content:     resource.Content,
		MediaUrl:    resource.MediaUrl,
		Duration:    resource.Duration,
		Difficulty:  resource.Difficulty,
		IsActive:    resource.IsActive,
		UsageCount:  resource.UsageCount,
		AvgRating:   resource.AvgRating,
		RatingCount: resource.RatingCount,
		Tags:        resource.Tags,
		Description: resource.Description,
		CreatedAt:   resource.CreatedAt,
	}
}

// ConvertToEmergencyVolunteerItem 转换为志愿者项目响应
func ConvertToEmergencyVolunteerItem(volunteer model.EmergencyVolunteer) EmergencyVolunteerItem {
	statusName := "待审核"
	switch volunteer.Status {
	case 1:
		statusName = "待审核"
	case 2:
		statusName = "已激活"
	case 3:
		statusName = "暂停"
	case 4:
		statusName = "停用"
	}

	return EmergencyVolunteerItem{
		ID:           volunteer.ID,
		UserID:       volunteer.UserID,
		Nickname:     volunteer.User.Nickname,
		Avatar:       volunteer.User.AvatarUrl,
		Status:       volunteer.Status,
		StatusName:   statusName,
		IsOnline:     volunteer.IsOnline,
		Specialties:  volunteer.Specialties,
		MaxLoad:      volunteer.MaxLoad,
		CurrentLoad:  volunteer.CurrentLoad,
		TotalHelped:  volunteer.TotalHelped,
		AvgRating:    volunteer.AvgRating,
		RatingCount:  volunteer.RatingCount,
		LastActiveAt: volunteer.LastActiveAt,
	}
}
