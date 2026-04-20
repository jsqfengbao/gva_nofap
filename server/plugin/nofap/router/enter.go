package router

type RouterGroup struct {
	AuthRouter
	UserRouter
	ProfileRouter
	AssessmentRouter
	CheckinRouter
	AchievementRouter
	CommunityRouter
	EmergencyRouter
	LearningRouter
	StatsRouter
	AdminRouter
}

var RouterGroupApp = new(RouterGroup)
