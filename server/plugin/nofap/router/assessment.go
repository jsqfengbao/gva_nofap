package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/middleware"
	"github.com/gin-gonic/gin"
)

type AssessmentRouter struct{}

// InitAssessmentRouter 初始化评估路由信息
func (s *AssessmentRouter) InitAssessmentRouter(Router *gin.RouterGroup) {
	assessmentRouter := Router.Group("assessment").Use(middleware.MiniprogramJWTAuth())
	assessmentRouterWithoutRecord := Router.Group("assessment").Use(middleware.MiniprogramJWTAuth())
	miniprogramApi := api.ApiGroupApp

	{
		assessmentRouter.POST("submit", miniprogramApi.AssessmentApi.SubmitAssessment)     // 提交评估结果
		assessmentRouter.GET("history", miniprogramApi.AssessmentApi.GetAssessmentHistory) // 获取评估历史
		assessmentRouter.GET("latest", miniprogramApi.AssessmentApi.GetLatestAssessment)   // 获取最新评估结果
		assessmentRouter.GET("stats", miniprogramApi.AssessmentApi.GetAssessmentStats)     // 获取评估统计数据
	}

	{
		assessmentRouterWithoutRecord.GET("", miniprogramApi.AssessmentApi.GetLatestAssessment) // 获取最新评估结果（无记录版本）
	}
}
