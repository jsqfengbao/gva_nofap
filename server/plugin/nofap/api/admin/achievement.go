package admin

import (
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/service"
	"github.com/gin-gonic/gin"
)

type AchievementAdminApi struct{}

// GetAchievementList 获取成就列表（分页）
// @Summary 获取成就列表
// @Description 获取成就列表分页
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/achievements [get]
func (a *AchievementAdminApi) GetAchievementList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	var name *string
	var category *string

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if v, ok := c.GetQuery("name"); ok && v != "" {
		name = &v
	}
	if v, ok := c.GetQuery("category"); ok && v != "" {
		category = &v
	}

	list, total, err := service.ServiceGroupApp.AchievementService.GetAchievementAdminList(pageInfo, name, category)
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

// CreateAchievement 创建成就
// @Summary 创建成就
// @Description 创建成就
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/achievements [post]
func (a *AchievementAdminApi) CreateAchievement(c *gin.Context) {
	var achievement model.Achievement
	if err := c.ShouldBindJSON(&achievement); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.AchievementService.CreateAchievement(&achievement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// UpdateAchievement 更新成就
// @Summary 更新成就
// @Description 更新成就
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/achievements [put]
func (a *AchievementAdminApi) UpdateAchievement(c *gin.Context) {
	var achievement model.Achievement
	if err := c.ShouldBindJSON(&achievement); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.AchievementService.UpdateAchievement(&achievement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteAchievement 删除成就
// @Summary 删除成就
// @Description 删除成就
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/achievements/{id} [delete]
func (a *AchievementAdminApi) DeleteAchievement(c *gin.Context) {
	idStr := c.Param("id")
	err := service.ServiceGroupApp.AchievementService.DeleteAchievement(idStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
