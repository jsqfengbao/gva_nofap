package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/router"
	"github.com/gin-gonic/gin"
)

// InitializeRouter 初始化小程序路由
func InitializeRouter(publicGroup *gin.RouterGroup) {
	rg := router.RouterGroupApp
	rg.AuthRouter.InitAuthRouter(publicGroup)
	rg.UserRouter.InitUserRouter(publicGroup)
	rg.ProfileRouter.InitProfileRouter(publicGroup)
	rg.AssessmentRouter.InitAssessmentRouter(publicGroup)
	rg.CheckinRouter.InitCheckinRouter(publicGroup)
	rg.AchievementRouter.InitAchievementRouter(publicGroup)
	rg.CommunityRouter.InitCommunityRouter(publicGroup)
	rg.EmergencyRouter.InitEmergencyRouter(publicGroup)
	rg.LearningRouter.InitLearningRouter(publicGroup)
	rg.StatsRouter.InitStatsRouter(publicGroup)
	rg.AdminRouter.InitAdminRouter(publicGroup)
}