package miniprogram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Plugin struct{}

func (p *Plugin) Register(logger *zap.Logger) error {
	// 初始化数据库
	err := initialize.InitializeDB()
	if err != nil {
		return err
	}
	logger.Info("miniprogram plugin database initialized")
	return nil
}

func (p *Plugin) RouterPath() string {
	return "v1/miniprogram"
}

func (p *Plugin) InitializeRouter(publicGroup *gin.RouterGroup) {
	initialize.InitializeRouter(publicGroup)
}

var _ system.Plugin = (*Plugin)(nil)
