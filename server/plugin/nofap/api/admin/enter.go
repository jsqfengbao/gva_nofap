package admin

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	StatisticsApi
	UserAdminApi
	LearningAdminApi
	EmergencyAdminApi
	PostAdminApi
	CommentAdminApi
	AchievementAdminApi
}
