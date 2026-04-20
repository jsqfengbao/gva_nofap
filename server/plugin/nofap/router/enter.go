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
}

var RouterGroupApp = new(RouterGroup)
