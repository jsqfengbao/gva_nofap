package admin

import (
	"fmt"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/service"
	"github.com/gin-gonic/gin"
)

type UserAdminApi struct{}

// GetUserList 获取小程序用户列表（分页）
// @Summary 获取小程序用户列表
// @Description 获取小程序用户列表分页
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/users [get]
func (a *UserAdminApi) GetUserList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	var nickname *string
	var status *int

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 可选查询参数
	if v, ok := c.GetQuery("nickname"); ok && v != "" {
		nickname = &v
	}
	if v, ok := c.GetQuery("status"); ok {
		var s int
		if _, err := fmt.Sscanf(v, "%d", &s); err == nil {
			status = &s
		}
	}

	list, total, err := service.ServiceGroupApp.UserService.GetUserAdminList(pageInfo, nickname, status)
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

// UpdateUserStatus 更新用户状态（启用/禁用）
// @Summary 更新用户状态
// @Description 更新用户启用/禁用状态
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/user/{id}/status [put]
func (a *UserAdminApi) UpdateUserStatus(c *gin.Context) {
	idStr := c.Param("id")
	var req struct {
		Status int `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.UserService.UpdateUserStatusByAdmin(idStr, req.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("操作成功", c)
}

// GetUserDetail 获取用户详情
// @Summary 获取用户详情
// @Description 获取用户详情
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/user/{id} [get]
func (a *UserAdminApi) GetUserDetail(c *gin.Context) {
	idStr := c.Param("id")
	user, err := service.ServiceGroupApp.UserService.GetUserDetailByAdmin(idStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/user/{id} [delete]
func (a *UserAdminApi) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	err := service.ServiceGroupApp.UserService.DeleteUserByAdmin(idStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
