package miniprogram

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserSettings 用户设置表
type UserSettings struct {
	global.GVA_MODEL
	UserID uint `json:"userId" gorm:"uniqueIndex;not null;comment:用户ID"`

	// 通知设置
	CheckinReminder   bool `json:"checkinReminder" gorm:"default:true;comment:打卡提醒"`
	CommunityReply    bool `json:"communityReply" gorm:"default:true;comment:社区回复提醒"`
	AchievementUnlock bool `json:"achievementUnlock" gorm:"default:true;comment:成就解锁提醒"`
	WeeklyReport      bool `json:"weeklyReport" gorm:"default:true;comment:周报提醒"`
	EmergencyAlert    bool `json:"emergencyAlert" gorm:"default:true;comment:紧急求助提醒"`
	LearningReminder  bool `json:"learningReminder" gorm:"default:true;comment:学习提醒"`

	// 隐私设置
	ShowProfile        bool `json:"showProfile" gorm:"default:true;comment:显示个人资料"`
	ShowStats          bool `json:"showStats" gorm:"default:true;comment:显示统计数据"`
	ShowAchievements   bool `json:"showAchievements" gorm:"default:true;comment:显示成就"`
	AllowFriendRequest bool `json:"allowFriendRequest" gorm:"default:true;comment:允许好友申请"`
	ShowOnlineStatus   bool `json:"showOnlineStatus" gorm:"default:true;comment:显示在线状态"`

	// 数据导出设置
	LastExportAt *time.Time `json:"lastExportAt" gorm:"comment:最后导出时间"`
	ExportCount  int        `json:"exportCount" gorm:"default:0;comment:导出次数"`
}

// TableName 设置表名
func (UserSettings) TableName() string {
	return "nofap_user_settings"
}

// DataExport 数据导出记录表
type DataExport struct {
	global.GVA_MODEL
	UserID      uint       `json:"userId" gorm:"not null;comment:用户ID"`
	ExportType  string     `json:"exportType" gorm:"size:10;not null;comment:导出格式:json,csv"`
	DataTypes   string     `json:"dataTypes" gorm:"type:text;not null;comment:导出数据类型(JSON数组)"`
	DateRange   string     `json:"dateRange" gorm:"size:20;comment:日期范围"`
	FilePath    string     `json:"filePath" gorm:"type:text;comment:文件路径"`
	FileName    string     `json:"fileName" gorm:"size:255;comment:文件名"`
	FileSize    int64      `json:"fileSize" gorm:"comment:文件大小(字节)"`
	Status      int        `json:"status" gorm:"default:1;comment:状态:1处理中,2完成,3失败"`
	DownloadURL string     `json:"downloadUrl" gorm:"type:text;comment:下载链接"`
	ExpiresAt   *time.Time `json:"expiresAt" gorm:"comment:过期时间"`
}

// TableName 设置表名
func (DataExport) TableName() string {
	return "nofap_data_exports"
}
