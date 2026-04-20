package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/middleware"
	"github.com/gin-gonic/gin"
)

type ProfileRouter struct{}

// InitProfileRouter 初始化个人中心路由
func (p *ProfileRouter) InitProfileRouter(Router *gin.RouterGroup) {
	profileRouter := Router.Group("profile").Use(middleware.MiniprogramJWTAuth())
	{
		var userApi = api.ApiGroupApp.UserApi
		profileRouter.GET("stats", userApi.GetProfileStats)                   // 获取个人中心统计数据
		profileRouter.GET("settings", userApi.GetUserSettings)                // 获取用户设置
		profileRouter.PUT("notification", userApi.UpdateNotificationSettings) // 更新通知设置
		profileRouter.PUT("privacy", userApi.UpdatePrivacySettings)           // 更新隐私设置
		profileRouter.POST("export", userApi.CreateDataExport)                // 创建数据导出任务
		// profileRouter.GET("export/:exportId", userApi.GetExportStatus)        // 获取导出状态 - 暂未实现
		// profileRouter.GET("download/:exportId", userApi.DownloadExport)       // 下载导出文件 - 暂未实现
	}
	{
		// profileRouter.GET("help", userApi.GetHelpContent)                      // 获取帮助内容 - 暂未实现
	}
}
