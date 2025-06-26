package miniprogram

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AchievementRouter struct{}

// InitAchievementRouter 初始化成就系统路由
func (s *AchievementRouter) InitAchievementRouter(Router *gin.RouterGroup) {
	achievementRouter := Router.Group("achievement").Use(middleware.MiniprogramJWTAuth())
	achievementApi := v1.ApiGroupApp.MiniprogramApiGroup.AchievementApi

	{
		achievementRouter.GET("list", achievementApi.GetUserAchievements)           // 获取用户成就列表
		achievementRouter.GET("user", achievementApi.GetUserAchievementsForProfile) // 获取用户成就（个人中心）
		achievementRouter.GET("stats", achievementApi.GetAchievementStats)          // 获取成就统计
		achievementRouter.GET("game-stats", achievementApi.GetGameStats)            // 获取游戏化统计
		achievementRouter.GET("level-progress", achievementApi.GetLevelProgress)    // 获取等级进度
		achievementRouter.GET("progress", achievementApi.GetAchievementProgress)    // 获取成就进度
	}
}
