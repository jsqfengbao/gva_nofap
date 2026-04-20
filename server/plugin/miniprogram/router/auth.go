package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/api"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

// InitAuthRouter 初始化认证路由
func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth")
	authApi := api.ApiGroupApp.AuthApi

	{
		// 无需认证的路由
		authRouter.POST("wx-login", authApi.WxLogin)                      // 微信登录
		authRouter.POST("login", authApi.Login)                           // 普通登录
		authRouter.POST("register", authApi.Register)                     // 注册
		authRouter.POST("refresh", authApi.RefreshToken)                  // 刷新token
		authRouter.POST("clear-auth-failures", authApi.ClearAuthFailures) // 清除认证失败记录（开发测试用）
	}

}
