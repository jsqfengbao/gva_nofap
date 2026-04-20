package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/api/admin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/middleware"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

// InitAdminRouter 初始化 admin 路由
func (a *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouter := Router.Group("admin").Use(middleware.MiniprogramJWTAuth())
	{
		// 统计数据
		adminRouter.GET("statistics", admin.ApiGroupApp.StatisticsApi.GetStatistics)

		// 用户管理
		adminRouter.GET("users", admin.ApiGroupApp.UserAdminApi.GetUserList)
		adminRouter.GET("users/:id", admin.ApiGroupApp.UserAdminApi.GetUserDetail)
		adminRouter.PUT("users/:id/status", admin.ApiGroupApp.UserAdminApi.UpdateUserStatus)
		adminRouter.DELETE("users/:id", admin.ApiGroupApp.UserAdminApi.DeleteUser)

		// 学习内容管理
		adminRouter.GET("learning", admin.ApiGroupApp.LearningAdminApi.GetLearningList)
		adminRouter.POST("learning", admin.ApiGroupApp.LearningAdminApi.CreateLearning)
		adminRouter.PUT("learning", admin.ApiGroupApp.LearningAdminApi.UpdateLearning)
		adminRouter.DELETE("learning/:id", admin.ApiGroupApp.LearningAdminApi.DeleteLearning)

		// 紧急资源管理
		adminRouter.GET("emergency/resources", admin.ApiGroupApp.EmergencyAdminApi.GetEmergencyList)
		adminRouter.POST("emergency/resources", admin.ApiGroupApp.EmergencyAdminApi.CreateEmergency)
		adminRouter.PUT("emergency/resources", admin.ApiGroupApp.EmergencyAdminApi.UpdateEmergency)
		adminRouter.DELETE("emergency/resources/:id", admin.ApiGroupApp.EmergencyAdminApi.DeleteEmergency)
		adminRouter.GET("emergency/helps", admin.ApiGroupApp.EmergencyAdminApi.GetHelpAdminList)
		adminRouter.PUT("emergency/helps/:id/status", admin.ApiGroupApp.EmergencyAdminApi.UpdateHelpStatus)
		adminRouter.DELETE("emergency/helps/:id", admin.ApiGroupApp.EmergencyAdminApi.DeleteHelpAdmin)

		// 帖子管理
		adminRouter.GET("posts", admin.ApiGroupApp.PostAdminApi.GetPostList)
		adminRouter.PUT("posts/:id/status", admin.ApiGroupApp.PostAdminApi.UpdatePostStatus)
		adminRouter.DELETE("posts/:id", admin.ApiGroupApp.PostAdminApi.DeletePost)

		// 评论管理
		adminRouter.GET("comments", admin.ApiGroupApp.CommentAdminApi.GetCommentList)
		adminRouter.PUT("comments/:id/status", admin.ApiGroupApp.CommentAdminApi.UpdateCommentStatus)
		adminRouter.DELETE("comments/:id", admin.ApiGroupApp.CommentAdminApi.DeleteComment)

		// 成就管理
		adminRouter.GET("achievements", admin.ApiGroupApp.AchievementAdminApi.GetAchievementList)
		adminRouter.POST("achievements", admin.ApiGroupApp.AchievementAdminApi.CreateAchievement)
		adminRouter.PUT("achievements", admin.ApiGroupApp.AchievementAdminApi.UpdateAchievement)
		adminRouter.DELETE("achievements/:id", admin.ApiGroupApp.AchievementAdminApi.DeleteAchievement)
	}
}