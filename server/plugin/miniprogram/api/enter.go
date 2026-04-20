package api

type ApiGroup struct {
	AuthApi        AuthApi
	UserApi        UserApi
	AssessmentApi  AssessmentApi
	CheckinApi     CheckinApi
	AchievementApi AchievementApi
	CommunityApi   CommunityApi
	EmergencyApi   EmergencyApi
	LearningApi    LearningApi
	StatsApi       StatsApi
}
