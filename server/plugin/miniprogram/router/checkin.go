package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/middleware"
	"github.com/gin-gonic/gin"
)

type CheckinRouter struct{}

// InitCheckinRouter 初始化 打卡 路由信息
func (s *CheckinRouter) InitCheckinRouter(Router *gin.RouterGroup) {
	checkinRouter := Router.Group("checkin").Use(middleware.MiniprogramJWTAuth())
	checkinRouterWithoutRecord := Router.Group("checkin").Use(middleware.MiniprogramJWTAuth())
	var checkinApi = api.ApiGroupApp.CheckinApi
	{
		checkinRouter.POST("daily", checkinApi.DailyCheckin)                       // 每日打卡
		checkinRouter.GET("today", checkinApi.GetTodayCheckin)                     // 获取今日打卡状态
		checkinRouter.GET("history", checkinApi.GetCheckinHistory)                 // 获取打卡历史
		checkinRouter.GET("statistics", checkinApi.GetCheckinStatistics)           // 获取打卡统计
		checkinRouter.GET("weekly", checkinApi.GetWeeklyProgress)                  // 获取本周进度
		checkinRouter.GET("weekly-progress", checkinApi.GetWeeklyProgressForChart) // 获取本周进度图表数据
		checkinRouter.GET("calendar", checkinApi.GetMonthlyCalendar)               // 获取月度日历
	}
	{
		checkinRouterWithoutRecord.GET("test", checkinApi.GetTodayCheckin) // 测试接口
	}
}
