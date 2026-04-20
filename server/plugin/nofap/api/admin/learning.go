package admin

import (
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/service"
	"github.com/gin-gonic/gin"
)

type LearningAdminApi struct{}

// GetLearningList 获取学习内容列表（分页）
// @Summary 获取学习内容列表
// @Description 获取学习内容列表分页
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/learning [get]
func (a *LearningAdminApi) GetLearningList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	var title *string
	var category *string

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if v, ok := c.GetQuery("title"); ok && v != "" {
		title = &v
	}
	if v, ok := c.GetQuery("category"); ok && v != "" {
		category = &v
	}

	list, total, err := service.ServiceGroupApp.LearningService.GetLearningAdminList(pageInfo, title, nil, category, nil)
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

// CreateLearning 创建学习内容
// @Summary 创建学习内容
// @Description 创建学习内容
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/learning [post]
func (a *LearningAdminApi) CreateLearning(c *gin.Context) {
	var learning model.LearningContent
	if err := c.ShouldBindJSON(&learning); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.LearningService.CreateLearningContentAdmin(&learning)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// UpdateLearning 更新学习内容
// @Summary 更新学习内容
// @Description 更新学习内容
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/learning [put]
func (a *LearningAdminApi) UpdateLearning(c *gin.Context) {
	var learning model.LearningContent
	if err := c.ShouldBindJSON(&learning); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.LearningService.UpdateLearningContentAdmin(&learning)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteLearning 删除学习内容
// @Summary 删除学习内容
// @Description 删除学习内容
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/learning/{id} [delete]
func (a *LearningAdminApi) DeleteLearning(c *gin.Context) {
	idStr := c.Param("id")
	err := service.ServiceGroupApp.LearningService.DeleteLearningContentAdmin(idStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
