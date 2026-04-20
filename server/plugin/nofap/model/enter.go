package model

// 小程序相关的所有数据模型

// 用户相关
type (
	// WxUser 微信用户
	WxUserModel = WxUser
	// UserSettings 用户设置
	UserSettingsModel = UserSettings
	// DataExport 数据导出
	DataExportModel = DataExport
)

// 戒色记录相关
type (
	// AbstinenceRecord 戒色记录
	AbstinenceRecordModel = AbstinenceRecord
	// DailyCheckin 每日打卡
	DailyCheckinModel = DailyCheckin
	// AssessmentResult 评估结果
	AssessmentResultModel = AssessmentResult
)

// 社区相关
type (
	// CommunityPost 社区动态
	CommunityPostModel = CommunityPost
	// CommunityComment 社区评论
	CommunityCommentModel = CommunityComment
	// CommunityLike 社区点赞
	CommunityLikeModel = CommunityLike
)

// 紧急求助相关
type (
	// EmergencyHelp 紧急求助
	EmergencyHelpModel = EmergencyHelp
	// HelpResponse 求助响应
	HelpResponseModel = HelpResponse
)

// 学习内容相关
type (
	// LearningContent 学习内容
	LearningContentModel = LearningContent
	// UserLearningRecord 用户学习记录
	UserLearningRecordModel = UserLearningRecord
)

// 成就系统相关
type (
	// Achievement 成就
	AchievementModel = Achievement
	// UserAchievement 用户成就
	UserAchievementModel = UserAchievement
)
