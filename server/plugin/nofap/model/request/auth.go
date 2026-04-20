package request

// WxLoginRequest 微信登录请求
type WxLoginRequest struct {
	Code          string `json:"code" binding:"required"` // 微信登录凭证
	EncryptedData string `json:"encryptedData"`           // 加密的用户信息
	Iv            string `json:"iv"`                      // 加密算法的初始向量
}

// WxUserInfo 微信用户信息
type WxUserInfo struct {
	Openid    string `json:"openId"`
	Nickname  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
}

// RefreshTokenRequest 刷新token请求
type RefreshTokenRequest struct {
	Token string `json:"token" binding:"required"` // 需要刷新的token
}

// LoginRequest 普通登录请求
type LoginRequest struct {
	Phone    string `json:"phone" binding:"required"`    // 手机号
	Password string `json:"password" binding:"required"` // 密码
}

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Phone    string `json:"phone" binding:"required"`    // 手机号
	Password string `json:"password" binding:"required"` // 密码
	Nickname string `json:"nickname" binding:"required"` // 昵称
}

// UpdateUserInfoRequest 更新用户信息请求
type UpdateUserInfoRequest struct {
	Nickname  string `json:"nickname"`  // 用户昵称
	AvatarUrl string `json:"avatarUrl"` // 头像URL
	Gender    int    `json:"gender"`    // 性别
}

// UpdatePrivacyLevelRequest 更新隐私级别请求
type UpdatePrivacyLevelRequest struct {
	PrivacyLevel int `json:"privacyLevel" binding:"required,min=1,max=3"` // 隐私级别：1低，2中，3高
}

// NotificationSettingsRequest 通知设置请求
type NotificationSettingsRequest struct {
	CheckinReminder   bool `json:"checkinReminder"`   // 打卡提醒
	CommunityReply    bool `json:"communityReply"`    // 社区回复提醒
	AchievementUnlock bool `json:"achievementUnlock"` // 成就解锁提醒
	WeeklyReport      bool `json:"weeklyReport"`      // 周报提醒
	EmergencyAlert    bool `json:"emergencyAlert"`    // 紧急求助提醒
	LearningReminder  bool `json:"learningReminder"`  // 学习提醒
}

// DataExportRequest 数据导出请求
type DataExportRequest struct {
	ExportType   string   `json:"exportType" binding:"required,oneof=json csv"` // 导出格式：json、csv
	DataTypes    []string `json:"dataTypes" binding:"required"`                 // 导出数据类型
	DateRange    string   `json:"dateRange"`                                    // 日期范围：all、last_month、last_3months、last_year
	IncludeStats bool     `json:"includeStats"`                                 // 是否包含统计数据
}

// PrivacySettingsRequest 隐私设置请求
type PrivacySettingsRequest struct {
	ShowProfile        bool `json:"showProfile"`        // 显示个人资料
	ShowStats          bool `json:"showStats"`          // 显示统计数据
	ShowAchievements   bool `json:"showAchievements"`   // 显示成就
	AllowFriendRequest bool `json:"allowFriendRequest"` // 允许好友申请
	ShowOnlineStatus   bool `json:"showOnlineStatus"`   // 显示在线状态
}
