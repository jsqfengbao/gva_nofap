package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/middleware"
	"github.com/gin-gonic/gin"
)

type LearningRouter struct{}

// InitLearningRouter 初始化学习内容路由
func (r *LearningRouter) InitLearningRouter(Router *gin.RouterGroup) {
	learningRouter := Router.Group("learning").Use(middleware.MiniprogramJWTAuth())
	learningApi := api.ApiGroupApp.LearningApi

	{
		// 需要JWT认证的路由 - 只保留实际存在的方法
		learningRouter.POST("content", learningApi.CreateLearningContent)       // 创建学习内容
		learningRouter.PUT("content", learningApi.UpdateLearningContent)        // 更新学习内容
		learningRouter.DELETE("content/:id", learningApi.DeleteLearningContent) // 删除学习内容

		// 已实现的方法
		learningRouter.GET("stats", learningApi.GetLearningStats) // 获取学习统计

		// 以下方法暂时移除，等待实现
		// learningRouter.GET("content/:id", learningApi.GetLearningContent)       // 获取学习内容详情
		// learningRouter.GET("contents", learningApi.GetLearningContents)         // 获取学习内容列表
		// learningRouter.POST("start", learningApi.StartLearning)                 // 开始学习
		// learningRouter.PUT("progress", learningApi.UpdateLearningProgress)      // 更新学习进度
		// learningRouter.POST("like", learningApi.ToggleLike)                     // 切换点赞状态
		// learningRouter.POST("collect", learningApi.ToggleCollect)               // 切换收藏状态
		// learningRouter.GET("recommendations", learningApi.GetRecommendations)   // 获取推荐内容
		// learningRouter.GET("records", learningApi.GetUserLearningRecords)       // 获取用户学习记录
		// learningRouter.POST("rate", learningApi.RateLearningContent)            // 评分学习内容
		// learningRouter.GET("category-stats", learningApi.GetCategoryStats)      // 获取分类统计
		// learningRouter.GET("homepage", learningApi.GetLearningHomepage)         // 获取学习首页数据
	}
}
