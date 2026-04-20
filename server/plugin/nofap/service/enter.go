package service

type ServiceGroup struct {
	AuthService
	UserService
	AssessmentService
	CheckinService
	AchievementService
	GamificationService
	CommunityService
	EmergencyService
	LearningService
}

var ServiceGroupApp = new(ServiceGroup)
