package miniprogram

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StatsRouter struct{}

// InitStatsRouter 初始化统计路由
func (s *StatsRouter) InitStatsRouter(Router *gin.RouterGroup) {
	statsRouter := Router.Group("stats").Use(middleware.MiniprogramJWTAuth())
	{
		var statsApi = v1.ApiGroupApp.MiniprogramApiGroup.StatsApi
		statsRouter.GET("overall", statsApi.GetOverall) // 获取总体统计
		statsRouter.GET("trends", statsApi.GetTrends)   // 获取趋势数据
	}
}
