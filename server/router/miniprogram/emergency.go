package miniprogram

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmergencyRouter struct{}

// InitEmergencyRouter 初始化紧急求助路由
func (e *EmergencyRouter) InitEmergencyRouter(Router *gin.RouterGroup) {
	emergencyRouter := Router.Group("emergency").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	emergencyRouterWithoutRecord := Router.Group("emergency").Use(middleware.JWTAuth())
	emergencyApi := v1.ApiGroupApp.MiniprogramApiGroup.EmergencyApi
	{
		// 紧急求助管理
		emergencyRouter.POST("help", emergencyApi.CreateEmergencyHelp)       // 创建紧急求助
		emergencyRouter.GET("help/list", emergencyApi.GetEmergencyHelps)     // 获取紧急求助列表
		emergencyRouter.GET("help/:id", emergencyApi.GetEmergencyHelpDetail) // 获取紧急求助详情
		emergencyRouter.PUT("help", emergencyApi.UpdateEmergencyHelp)        // 更新紧急求助

		// 志愿者响应
		emergencyRouter.POST("response", emergencyApi.CreateEmergencyResponse) // 创建紧急求助响应
		emergencyRouter.POST("connect", emergencyApi.ConnectVolunteer)         // 连接志愿者

		// 志愿者管理
		emergencyRouter.POST("volunteer/register", emergencyApi.RegisterVolunteer)  // 注册志愿者
		emergencyRouter.PUT("volunteer/status", emergencyApi.UpdateVolunteerStatus) // 更新志愿者状态
		emergencyRouter.GET("volunteers/online", emergencyApi.GetOnlineVolunteers)  // 获取在线志愿者

		// 紧急求助资源
		emergencyRouter.GET("resources", emergencyApi.GetEmergencyResources) // 获取紧急求助资源
		emergencyRouter.POST("resources/:id/use", emergencyApi.UseResource)  // 使用资源
		emergencyRouter.POST("resources/rate", emergencyApi.RateResource)    // 资源评分

		// 统计数据
		emergencyRouter.GET("stats", emergencyApi.GetEmergencyStats) // 获取紧急求助统计
	}
	{
		// 无记录路由 - 使用不同的路径避免冲突
		emergencyRouterWithoutRecord.GET("public/resources", emergencyApi.GetEmergencyResources) // 获取紧急求助资源(公开访问)
		emergencyRouterWithoutRecord.GET("public/volunteers", emergencyApi.GetOnlineVolunteers)  // 获取在线志愿者数量(公开访问)
	}
}
