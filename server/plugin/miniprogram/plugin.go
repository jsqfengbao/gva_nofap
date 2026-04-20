package miniprogram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/initialize"
	"github.com/gin-gonic/gin"
)

type Plugin struct{}

// GetPlugin 插件实例获取
func GetPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Register(publicGroup *gin.RouterGroup) {
	// 初始化数据库
	err := initialize.InitializeDB()
	if err != nil {
		panic(err)
	}
	// 初始化路由
	initialize.InitializeRouter(publicGroup)
}

func (p *Plugin) RouterPath() string {
	return "v1/miniprogram"
}
