package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	// 小程序路由 - 认证相关接口
	miniprogramRouter := router.RouterGroupApp.Miniprogram
	miniprogramGroup := publicGroup.Group("api/v1/miniprogram")
	{
		miniprogramRouter.AuthRouter.InitAuthRouter(miniprogramGroup)               // 小程序认证路由
		miniprogramRouter.UserRouter.InitUserRouter(miniprogramGroup)               // 小程序用户路由
		miniprogramRouter.ProfileRouter.InitProfileRouter(miniprogramGroup)         // 小程序个人中心路由
		miniprogramRouter.AssessmentRouter.InitAssessmentRouter(miniprogramGroup)   // 小程序评估路由
		miniprogramRouter.CheckinRouter.InitCheckinRouter(miniprogramGroup)         // 小程序打卡路由
		miniprogramRouter.AchievementRouter.InitAchievementRouter(miniprogramGroup) // 小程序成就路由
		miniprogramRouter.CommunityRouter.InitCommunityRouter(miniprogramGroup)     // 小程序社区路由
		miniprogramRouter.EmergencyRouter.InitEmergencyRouter(miniprogramGroup)     // 小程序紧急求助路由
		miniprogramRouter.LearningRouter.InitLearningRouter(miniprogramGroup)       // 小程序学习内容路由
	}

	holder(publicGroup, privateGroup)
}
