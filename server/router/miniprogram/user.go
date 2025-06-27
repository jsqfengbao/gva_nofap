package miniprogram

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// InitUserRouter 初始化用户路由
func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.MiniprogramJWTAuth())
	userApi := v1.ApiGroupApp.MiniprogramApiGroup.UserApi

	{
		// 用户相关的路由
		userRouter.GET("profile", userApi.GetUserProfile)           // 获取用户详细资料
		userRouter.GET("info", userApi.GetUserProfile)              // 获取用户信息（别名，兼容性）
		userRouter.PUT("info", userApi.UpdateUserInfo)              // 更新用户信息
		userRouter.PUT("privacy", userApi.UpdatePrivacyLevel)       // 更新隐私级别
		userRouter.POST("upload-avatar", userApi.UploadAvatar)      // 上传用户头像
		userRouter.POST("save-wx-avatar", userApi.SaveWxTempAvatar) // 保存微信临时头像
	}
}
