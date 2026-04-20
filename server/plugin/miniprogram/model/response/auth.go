package response

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model"
)

// WxLoginResponse 微信登录响应
type WxLoginResponse struct {
	Token string              `json:"token"` // JWT token
	User  *model.WxUser `json:"user"`  // 用户信息
}

// RefreshTokenResponse 刷新token响应
type RefreshTokenResponse struct {
	Token string `json:"token"` // 新的JWT token
}

// LoginResponse 普通登录响应
type LoginResponse struct {
	Token string              `json:"token"` // JWT token
	User  *model.WxUser `json:"user"`  // 用户信息
}

// RegisterResponse 用户注册响应
type RegisterResponse struct {
	User *model.WxUser `json:"user"` // 用户信息
}

// UserProfileResponse 用户详细资料响应
type UserProfileResponse struct {
	User             *model.WxUser           `json:"user"`             // 用户基本信息
	AbstinenceRecord *model.AbstinenceRecord `json:"abstinenceRecord"` // 戒色记录
}

// ProfileStatsResponse 个人中心统计数据响应
type ProfileStatsResponse struct {
	BasicStats         BasicStats          `json:"basicStats"`         // 基础统计
	RecentAchievements []RecentAchievement `json:"recentAchievements"` // 最近成就
	UserSettings       UserSettings        `json:"userSettings"`       // 用户设置
}

// BasicStats 基础统计数据
type BasicStats struct {
	CurrentStreak    int     `json:"currentStreak"`    // 当前连击天数
	LongestStreak    int     `json:"longestStreak"`    // 最长连击天数
	Experience       int     `json:"experience"`       // 经验值
	Level            int     `json:"level"`            // 等级
	LevelTitle       string  `json:"levelTitle"`       // 等级称号
	JoinDays         int     `json:"joinDays"`         // 加入天数
	AchievementCount int     `json:"achievementCount"` // 成就数量
	HelpCount        int     `json:"helpCount"`        // 帮助他人次数
	SuccessRate      float64 `json:"successRate"`      // 成功率
}

// RecentAchievement 最近成就
type RecentAchievement struct {
	ID          uint      `json:"id"`          // 成就ID
	Title       string    `json:"title"`       // 成就标题
	Description string    `json:"description"` // 成就描述
	Icon        string    `json:"icon"`        // 成就图标
	Rarity      string    `json:"rarity"`      // 稀有度
	UnlockedAt  time.Time `json:"unlockedAt"`  // 解锁时间
	DaysAgo     int       `json:"daysAgo"`     // 几天前获得
}

// UserSettings 用户设置
type UserSettings struct {
	NotificationSettings NotificationSettings `json:"notificationSettings"` // 通知设置
	PrivacySettings      PrivacySettings      `json:"privacySettings"`      // 隐私设置
}

// NotificationSettings 通知设置
type NotificationSettings struct {
	CheckinReminder   bool `json:"checkinReminder"`   // 打卡提醒
	CommunityReply    bool `json:"communityReply"`    // 社区回复提醒
	AchievementUnlock bool `json:"achievementUnlock"` // 成就解锁提醒
	WeeklyReport      bool `json:"weeklyReport"`      // 周报提醒
	EmergencyAlert    bool `json:"emergencyAlert"`    // 紧急求助提醒
	LearningReminder  bool `json:"learningReminder"`  // 学习提醒
}

// PrivacySettings 隐私设置
type PrivacySettings struct {
	ShowProfile        bool `json:"showProfile"`        // 显示个人资料
	ShowStats          bool `json:"showStats"`          // 显示统计数据
	ShowAchievements   bool `json:"showAchievements"`   // 显示成就
	AllowFriendRequest bool `json:"allowFriendRequest"` // 允许好友申请
	ShowOnlineStatus   bool `json:"showOnlineStatus"`   // 显示在线状态
}

// DataExportResponse 数据导出响应
type DataExportResponse struct {
	DownloadURL string    `json:"downloadUrl"` // 下载链接
	FileName    string    `json:"fileName"`    // 文件名
	FileSize    int64     `json:"fileSize"`    // 文件大小(字节)
	ExpiresAt   time.Time `json:"expiresAt"`   // 过期时间
}
