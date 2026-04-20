package admin

import (
	"fmt"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/service"
	"github.com/gin-gonic/gin"
)

type EmergencyAdminApi struct{}

// GetEmergencyList 获取紧急资源列表（分页）
// @Summary 获取紧急资源列表
// @Description 获取紧急资源列表分页
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/resources [get]
func (a *EmergencyAdminApi) GetEmergencyList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	var title *string
	var resourceType *int
	var isActive *bool

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if v, ok := c.GetQuery("title"); ok && v != "" {
		title = &v
	}
	if v, ok := c.GetQuery("type"); ok {
		var t int
		if _, err := fmt.Sscanf(v, "%d", &t); err == nil {
			resourceType = &t
		}
	}
	if v, ok := c.GetQuery("isActive"); ok {
		var a bool
		if _, err := fmt.Sscanf(v, "%t", &a); err == nil {
			isActive = &a
		}
	}

	list, total, err := service.ServiceGroupApp.EmergencyService.GetResourceAdminList(pageInfo, title, resourceType, isActive)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}

// CreateEmergency 创建紧急资源
// @Summary 创建紧急资源
// @Description 创建紧急资源
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/resources [post]
func (a *EmergencyAdminApi) CreateEmergency(c *gin.Context) {
	var resource model.EmergencyResource
	if err := c.ShouldBindJSON(&resource); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.EmergencyService.CreateResourceAdmin(&resource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// UpdateEmergency 更新紧急资源
// @Summary 更新紧急资源
// @Description 更新紧急资源
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/resources [put]
func (a *EmergencyAdminApi) UpdateEmergency(c *gin.Context) {
	var resource model.EmergencyResource
	if err := c.ShouldBindJSON(&resource); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.EmergencyService.UpdateResourceAdmin(&resource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteEmergency 删除紧急资源
// @Summary 删除紧急资源
// @Description 删除紧急资源
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/resources/{id} [delete]
func (a *EmergencyAdminApi) DeleteEmergency(c *gin.Context) {
	var idStr = c.Param("id")
	err := service.ServiceGroupApp.EmergencyService.DeleteResourceAdmin(idStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetHelpAdminList 获取紧急求助列表（分页）
// @Summary 获取紧急求助列表
// @Description 获取紧急求助列表分页
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/emergency/helps [get]
func (a *EmergencyAdminApi) GetHelpAdminList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	var status *int
	var helpType *int

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if v, ok := c.GetQuery("status"); ok {
		var s int
		if _, err := fmt.Sscanf(v, "%d", &s); err == nil {
			status = &s
		}
	}
	if v, ok := c.GetQuery("type"); ok {
		var t int
		if _, err := fmt.Sscanf(v, "%d", &t); err == nil {
			helpType = &t
		}
	}

	list, total, err := service.ServiceGroupApp.EmergencyService.GetHelpAdminList(pageInfo, status, helpType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}

// UpdateHelpStatus 更新紧急求助状态
// @Summary 更新紧急求助状态
// @Description 更新紧急求助状态
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/emergency/helps/{id}/status [put]
func (a *EmergencyAdminApi) UpdateHelpStatus(c *gin.Context) {
	idStr := c.Param("id")
	var req struct {
		Status int `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.EmergencyService.UpdateHelpStatus(idStr, req.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteHelpAdmin 删除紧急求助
// @Summary 删除紧急求助
// @Description 删除紧急求助
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/emergency/helps/{id} [delete]
func (a *EmergencyAdminApi) DeleteHelpAdmin(c *gin.Context) {
	idStr := c.Param("id")
	err := service.ServiceGroupApp.EmergencyService.DeleteHelpAdmin(idStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
