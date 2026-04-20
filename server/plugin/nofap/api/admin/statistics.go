package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/service"
	"github.com/gin-gonic/gin"
)

type StatisticsApi struct{}

// GetStatistics 获取管理端统计数据
// @Summary 获取管理端统计数据
// @Description 获取管理端统计数据
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/statistics [get]
func (a *StatisticsApi) GetStatistics(c *gin.Context) {
	var stats model.AdminStatistics
	var err error
	stats, err = service.ServiceGroupApp.UserService.GetAdminStatistics()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(stats, c)
}
