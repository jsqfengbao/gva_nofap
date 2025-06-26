package miniprogram

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// EmergencyHelp 紧急求助记录
type EmergencyHelp struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:主键ID"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`

	UserID      uint       `json:"userId" gorm:"not null;index;comment:用户ID"`
	Type        int        `json:"type" gorm:"not null;comment:求助类型(1-紧急冲动 2-情绪低落 3-复发担忧 4-其他)"`
	Status      int        `json:"status" gorm:"default:1;comment:状态(1-待响应 2-进行中 3-已解决 4-已关闭)"`
	Priority    int        `json:"priority" gorm:"default:2;comment:优先级(1-低 2-中 3-高 4-紧急)"`
	Description string     `json:"description" gorm:"type:text;comment:描述说明"`
	Location    string     `json:"location" gorm:"size:200;comment:地理位置(可选)"`
	IsAnonymous bool       `json:"isAnonymous" gorm:"default:true;comment:是否匿名求助"`
	RespondedAt *time.Time `json:"respondedAt" gorm:"comment:响应时间"`
	ResolvedAt  *time.Time `json:"resolvedAt" gorm:"comment:解决时间"`
	VolunteerID *uint      `json:"volunteerId" gorm:"index;comment:响应志愿者ID"`
	Rating      *int       `json:"rating" gorm:"comment:用户评分(1-5)"`
	Feedback    string     `json:"feedback" gorm:"type:text;comment:用户反馈"`

	// 关联
	User      WxUser              `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Volunteer *WxUser             `json:"volunteer,omitempty" gorm:"foreignKey:VolunteerID"`
	Responses []EmergencyResponse `json:"responses,omitempty" gorm:"foreignKey:HelpID"`
}

// EmergencyResponse 紧急求助响应记录
type EmergencyResponse struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:主键ID"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`

	HelpID      uint       `json:"helpId" gorm:"not null;index;comment:求助记录ID"`
	VolunteerID uint       `json:"volunteerId" gorm:"not null;index;comment:志愿者ID"`
	Type        int        `json:"type" gorm:"not null;comment:响应类型(1-文字回复 2-语音通话 3-视频通话 4-陪伴聊天)"`
	Content     string     `json:"content" gorm:"type:text;comment:响应内容"`
	Duration    int        `json:"duration" gorm:"default:0;comment:持续时间(分钟)"`
	IsActive    bool       `json:"isActive" gorm:"default:true;comment:是否活跃"`
	EndedAt     *time.Time `json:"endedAt" gorm:"comment:结束时间"`

	// 关联
	Help      EmergencyHelp `json:"help,omitempty" gorm:"foreignKey:HelpID"`
	Volunteer WxUser        `json:"volunteer,omitempty" gorm:"foreignKey:VolunteerID"`
}

// EmergencyVolunteer 紧急求助志愿者
type EmergencyVolunteer struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:主键ID"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`

	UserID       uint      `json:"userId" gorm:"not null;uniqueIndex;comment:用户ID"`
	Status       int       `json:"status" gorm:"default:1;comment:状态(1-待审核 2-已激活 3-暂停 4-停用)"`
	IsOnline     bool      `json:"isOnline" gorm:"default:false;comment:是否在线"`
	Specialties  string    `json:"specialties" gorm:"size:500;comment:专长领域"`
	MaxLoad      int       `json:"maxLoad" gorm:"default:3;comment:最大同时服务数"`
	CurrentLoad  int       `json:"currentLoad" gorm:"default:0;comment:当前服务数"`
	TotalHelped  int       `json:"totalHelped" gorm:"default:0;comment:累计帮助人数"`
	AvgRating    float64   `json:"avgRating" gorm:"default:0;comment:平均评分"`
	RatingCount  int       `json:"ratingCount" gorm:"default:0;comment:评分次数"`
	LastActiveAt time.Time `json:"lastActiveAt" gorm:"comment:最后活跃时间"`

	// 关联
	User WxUser `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// EmergencyResource 紧急求助资源(冥想、音乐、文章等)
type EmergencyResource struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:主键ID"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`

	Title       string  `json:"title" gorm:"size:200;not null;comment:标题"`
	Type        int     `json:"type" gorm:"not null;comment:资源类型(1-呼吸练习 2-冥想指导 3-舒缓音乐 4-励志文章 5-运动指导)"`
	Content     string  `json:"content" gorm:"type:text;comment:内容"`
	MediaUrl    string  `json:"mediaUrl" gorm:"size:500;comment:媒体文件URL"`
	Duration    int     `json:"duration" gorm:"default:0;comment:时长(秒)"`
	Difficulty  int     `json:"difficulty" gorm:"default:1;comment:难度等级(1-5)"`
	IsActive    bool    `json:"isActive" gorm:"default:true;comment:是否启用"`
	UsageCount  int     `json:"usageCount" gorm:"default:0;comment:使用次数"`
	AvgRating   float64 `json:"avgRating" gorm:"default:0;comment:平均评分"`
	RatingCount int     `json:"ratingCount" gorm:"default:0;comment:评分次数"`
	Tags        string  `json:"tags" gorm:"size:500;comment:标签(逗号分隔)"`
	Description string  `json:"description" gorm:"type:text;comment:描述"`
}

// TableName 设置表名
func (EmergencyHelp) TableName() string {
	return "nofap_emergency_help"
}

func (EmergencyResponse) TableName() string {
	return "nofap_emergency_response"
}

func (EmergencyVolunteer) TableName() string {
	return "nofap_emergency_volunteer"
}

func (EmergencyResource) TableName() string {
	return "nofap_emergency_resource"
}

// 创建表结构和索引
func (e *EmergencyHelp) AutoMigrate() error {
	db := global.GVA_DB

	// 自动迁移表结构
	if err := db.AutoMigrate(&EmergencyHelp{}, &EmergencyResponse{}, &EmergencyVolunteer{}, &EmergencyResource{}); err != nil {
		return err
	}

	// 创建额外索引
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_emergency_help_type_status ON nofap_emergency_help(type, status)").Error; err != nil {
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_emergency_help_created_at ON nofap_emergency_help(created_at)").Error; err != nil {
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_emergency_volunteer_online_status ON nofap_emergency_volunteer(is_online, status)").Error; err != nil {
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_emergency_resource_type_active ON nofap_emergency_resource(type, is_active)").Error; err != nil {
		return err
	}

	return nil
}

// GetTypeName 获取求助类型名称
func (e *EmergencyHelp) GetTypeName() string {
	switch e.Type {
	case 1:
		return "紧急冲动"
	case 2:
		return "情绪低落"
	case 3:
		return "复发担忧"
	case 4:
		return "其他"
	default:
		return "未知"
	}
}

// GetStatusName 获取状态名称
func (e *EmergencyHelp) GetStatusName() string {
	switch e.Status {
	case 1:
		return "待响应"
	case 2:
		return "进行中"
	case 3:
		return "已解决"
	case 4:
		return "已关闭"
	default:
		return "未知"
	}
}

// GetPriorityName 获取优先级名称
func (e *EmergencyHelp) GetPriorityName() string {
	switch e.Priority {
	case 1:
		return "低"
	case 2:
		return "中"
	case 3:
		return "高"
	case 4:
		return "紧急"
	default:
		return "中"
	}
}

// GetResourceTypeName 获取资源类型名称
func (e *EmergencyResource) GetResourceTypeName() string {
	switch e.Type {
	case 1:
		return "呼吸练习"
	case 2:
		return "冥想指导"
	case 3:
		return "舒缓音乐"
	case 4:
		return "励志文章"
	case 5:
		return "运动指导"
	default:
		return "其他"
	}
}
