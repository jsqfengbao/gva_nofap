package miniprogram

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

// InitAuthRouter 初始化认证路由
func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth")
	authApi := v1.ApiGroupApp.MiniprogramApiGroup.AuthApi
	userApi := v1.ApiGroupApp.MiniprogramApiGroup.UserApi

	{
		// 无需认证的路由
		authRouter.POST("wx-login", authApi.WxLogin)                      // 微信登录
		authRouter.POST("login", authApi.Login)                           // 普通登录
		authRouter.POST("register", authApi.Register)                     // 注册
		authRouter.POST("refresh", authApi.RefreshToken)                  // 刷新token
		authRouter.POST("clear-auth-failures", authApi.ClearAuthFailures) // 清除认证失败记录（开发测试用）
	}

	// 需要JWT认证的路由
	authRouterWithJWT := authRouter.Group("")
	authRouterWithJWT.Use(middleware.MiniprogramJWTAuth())
	{
		authRouterWithJWT.GET("profile", userApi.GetUserProfile)     // 获取用户资料
		authRouterWithJWT.PUT("profile", userApi.UpdateUserInfo)     // 更新用户信息
		authRouterWithJWT.PUT("privacy", userApi.UpdatePrivacyLevel) // 更新隐私级别
	}
}
