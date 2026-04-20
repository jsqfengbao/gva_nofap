package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/router"
	"github.com/gin-gonic/gin"
)

// InitializeRouter 初始化小程序路由
func InitializeRouter(publicGroup *gin.RouterGroup) {
	router.InitMiniprogramRouter(publicGroup)
}