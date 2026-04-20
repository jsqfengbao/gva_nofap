package admin

import (
	"fmt"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/service"
	"github.com/gin-gonic/gin"
)

type PostAdminApi struct{}

// GetPostList 获取社区帖子列表（分页）
// @Summary 获取社区帖子列表
// @Description 获取社区帖子列表分页
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/posts [get]
func (a *PostAdminApi) GetPostList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	var title *string
	var category *int
	var status *int

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if v, ok := c.GetQuery("title"); ok && v != "" {
		title = &v
	}
	if v, ok := c.GetQuery("category"); ok {
		var cat int
		if _, err := fmt.Sscanf(v, "%d", &cat); err == nil {
			category = &cat
		}
	}
	if v, ok := c.GetQuery("status"); ok {
		var s int
		if _, err := fmt.Sscanf(v, "%d", &s); err == nil {
			status = &s
		}
	}

	list, total, err := service.ServiceGroupApp.CommunityService.GetPostAdminList(pageInfo, title, category, status)
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

// UpdatePostStatus 更新帖子状态（通过/屏蔽）
// @Summary 更新帖子状态
// @Description 更新帖子状态（通过/屏蔽）
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/posts/{id}/status [put]
func (a *PostAdminApi) UpdatePostStatus(c *gin.Context) {
	idStr := c.Param("id")
	var req struct {
		Status int `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.CommunityService.UpdatePostStatusByAdmin(idStr, req.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("操作成功", c)
}

// DeletePost 删除帖子
// @Summary 删除帖子
// @Description 删除帖子
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/posts/{id} [delete]
func (a *PostAdminApi) DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	err := service.ServiceGroupApp.CommunityService.DeletePostByAdmin(idStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
