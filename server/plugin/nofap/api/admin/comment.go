package admin

import (
	"fmt"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/service"
	"github.com/gin-gonic/gin"
)

type CommentAdminApi struct{}

// GetCommentList 获取社区评论列表（分页）
// @Summary 获取社区评论列表
// @Description 获取社区评论列表分页
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/comments [get]
func (a *CommentAdminApi) GetCommentList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	var postId *uint
	var status *int

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if v, ok := c.GetQuery("postId"); ok && v != "" {
		var pid uint
		if _, err := fmt.Sscanf(v, "%d", &pid); err == nil {
			postId = &pid
		}
	}
	if v, ok := c.GetQuery("status"); ok {
		var s int
		if _, err := fmt.Sscanf(v, "%d", &s); err == nil {
			status = &s
		}
	}

	list, total, err := service.ServiceGroupApp.CommunityService.GetCommentAdminList(pageInfo, postId, status)
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

// UpdateCommentStatus 更新评论状态（通过/屏蔽）
// @Summary 更新评论状态
// @Description 更新评论状态（通过/屏蔽）
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/comments/{id}/status [put]
func (a *CommentAdminApi) UpdateCommentStatus(c *gin.Context) {
	var idStr = c.Param("id")
	var req struct {
		Status int `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.ServiceGroupApp.CommunityService.UpdateCommentStatusByAdmin(idStr, req.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("操作成功", c)
}

// DeleteComment 删除评论
// @Summary 删除评论
// @Description 删除评论
// @Tags Nofap-Admin
// @Authorization Bearer
// @Router /v1/nofap/admin/comments/{id} [delete]
func (a *CommentAdminApi) DeleteComment(c *gin.Context) {
	var idStr = c.Param("id")
	err := service.ServiceGroupApp.CommunityService.DeleteCommentByAdmin(idStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
