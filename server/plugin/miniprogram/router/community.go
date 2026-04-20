package router

import (
	sys_mw "github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/api"
	miniprogram_mw "github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/middleware"
	"github.com/gin-gonic/gin"
)

type CommunityRouter struct{}

// InitCommunityRouter 初始化社区路由
func (s *CommunityRouter) InitCommunityRouter(Router *gin.RouterGroup) {
	communityRouter := Router.Group("community").Use(miniprogram_mw.MiniprogramJWTAuth()).Use(sys_mw.CasbinHandler())
	communityRouterWithoutRecord := Router.Group("community").Use(miniprogram_mw.MiniprogramJWTAuth())
	var communityApi = api.ApiGroupApp.CommunityApi
	{
		// 需要记录操作日志的路由
		communityRouter.POST("post", communityApi.CreatePost)             // 创建帖子
		communityRouter.PUT("post/:id", communityApi.UpdatePost)          // 更新帖子
		communityRouter.DELETE("post/:id", communityApi.DeletePost)       // 删除帖子
		communityRouter.POST("comment", communityApi.CreateComment)       // 创建评论
		communityRouter.DELETE("comment/:id", communityApi.DeleteComment) // 删除评论
		communityRouter.POST("like", communityApi.ToggleLike)             // 点赞/取消点赞
	}
	{
		// 不需要记录操作日志的路由
		communityRouterWithoutRecord.GET("posts", communityApi.GetPosts)          // 获取帖子列表
		communityRouterWithoutRecord.GET("post/:id", communityApi.GetPostDetail)  // 获取帖子详情
		communityRouterWithoutRecord.GET("stats", communityApi.GetCommunityStats) // 获取社区统计
		communityRouterWithoutRecord.GET("my-posts", communityApi.GetUserPosts)   // 获取用户发帖记录
	}
}
