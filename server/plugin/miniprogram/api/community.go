package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CommunityApi struct{}

// CreatePost 创建帖子
// @Tags Community
// @Summary 创建社区帖子
// @Description 用户创建社区帖子
// @Accept json
// @Produce json
// @Param data body request.CreatePostRequest true "创建帖子信息"
// @Success 200 {object} response.Response{data=response.CreatePostResponse} "创建成功"
// @Router /miniprogram/community/post [post]
func (a *CommunityApi) CreatePost(c *gin.Context) {
	var req request.CreatePostRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	communityService := service.ServiceGroupApp.CommunityService
	result, err := communityService.CreatePost(userID, req)
	if err != nil {
		global.GVA_LOG.Error("创建帖子失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// GetPosts 获取帖子列表
// @Tags Community
// @Summary 获取社区帖子列表
// @Description 获取社区帖子列表，支持分类筛选和排序
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(20)
// @Param category query int false "分类" default(0)
// @Param sortBy query string false "排序字段" default(created_at)
// @Param order query string false "排序方向" default(desc)
// @Success 200 {object} response.Response{data=response.PostListResponse} "获取成功"
// @Router /miniprogram/community/posts [get]
func (a *CommunityApi) GetPosts(c *gin.Context) {
	var req request.GetPostsRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取用户ID（可能未登录）
	userID := utils.GetUserID(c)

	communityService := service.ServiceGroupApp.CommunityService
	result, err := communityService.GetPosts(userID, req)
	if err != nil {
		global.GVA_LOG.Error("获取帖子列表失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// GetPostDetail 获取帖子详情
// @Tags Community
// @Summary 获取帖子详情
// @Description 获取社区帖子详情，包含评论列表
// @Accept json
// @Produce json
// @Param id path uint true "帖子ID"
// @Success 200 {object} response.Response{data=response.PostDetailResponse} "获取成功"
// @Router /miniprogram/community/post/:id [get]
func (a *CommunityApi) GetPostDetail(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的帖子ID", c)
		return
	}

	// 获取用户ID（可能未登录）
	userID := utils.GetUserID(c)

	communityService := service.ServiceGroupApp.CommunityService
	result, err := communityService.GetPostDetail(userID, uint(postID))
	if err != nil {
		global.GVA_LOG.Error("获取帖子详情失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// UpdatePost 更新帖子
// @Tags Community
// @Summary 更新帖子
// @Description 用户更新自己的帖子
// @Accept json
// @Produce json
// @Param id path uint true "帖子ID"
// @Param data body request.UpdatePostRequest true "更新帖子信息"
// @Success 200 {object} response.Response "更新成功"
// @Router /miniprogram/community/post/:id [put]
func (a *CommunityApi) UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的帖子ID", c)
		return
	}

	var req request.UpdatePostRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	communityService := service.ServiceGroupApp.CommunityService
	err = communityService.UpdatePost(userID, uint(postID), req)
	if err != nil {
		global.GVA_LOG.Error("更新帖子失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeletePost 删除帖子
// @Tags Community
// @Summary 删除帖子
// @Description 用户删除自己的帖子
// @Accept json
// @Produce json
// @Param id path uint true "帖子ID"
// @Success 200 {object} response.Response "删除成功"
// @Router /miniprogram/community/post/:id [delete]
func (a *CommunityApi) DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的帖子ID", c)
		return
	}

	// 获取用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	communityService := service.ServiceGroupApp.CommunityService
	err = communityService.DeletePost(userID, uint(postID))
	if err != nil {
		global.GVA_LOG.Error("删除帖子失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// CreateComment 创建评论
// @Tags Community
// @Summary 创建评论
// @Description 用户创建评论或回复
// @Accept json
// @Produce json
// @Param data body request.CreateCommentRequest true "创建评论信息"
// @Success 200 {object} response.Response{data=response.CreateCommentResponse} "创建成功"
// @Router /miniprogram/community/comment [post]
func (a *CommunityApi) CreateComment(c *gin.Context) {
	var req request.CreateCommentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	communityService := service.ServiceGroupApp.CommunityService
	result, err := communityService.CreateComment(userID, req)
	if err != nil {
		global.GVA_LOG.Error("创建评论失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// DeleteComment 删除评论
// @Tags Community
// @Summary 删除评论
// @Description 用户删除自己的评论
// @Accept json
// @Produce json
// @Param id path uint true "评论ID"
// @Success 200 {object} response.Response "删除成功"
// @Router /miniprogram/community/comment/:id [delete]
func (a *CommunityApi) DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的评论ID", c)
		return
	}

	// 获取用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	communityService := service.ServiceGroupApp.CommunityService
	err = communityService.DeleteComment(userID, uint(commentID))
	if err != nil {
		global.GVA_LOG.Error("删除评论失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// ToggleLike 点赞/取消点赞
// @Tags Community
// @Summary 点赞或取消点赞
// @Description 用户对帖子或评论进行点赞或取消点赞
// @Accept json
// @Produce json
// @Param data body request.LikeRequest true "点赞信息"
// @Success 200 {object} response.Response{data=response.LikeResponse} "操作成功"
// @Router /miniprogram/community/like [post]
func (a *CommunityApi) ToggleLike(c *gin.Context) {
	var req request.LikeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	communityService := service.ServiceGroupApp.CommunityService
	result, err := communityService.ToggleLike(userID, req)
	if err != nil {
		global.GVA_LOG.Error("点赞操作失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// GetCommunityStats 获取社区统计信息
// @Tags Community
// @Summary 获取社区统计
// @Description 获取社区统计信息，包括帖子数、评论数等
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=response.CommunityStatsResponse} "获取成功"
// @Router /miniprogram/community/stats [get]
func (a *CommunityApi) GetCommunityStats(c *gin.Context) {
	communityService := service.ServiceGroupApp.CommunityService
	result, err := communityService.GetCommunityStats()
	if err != nil {
		global.GVA_LOG.Error("获取社区统计失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// GetUserPosts 获取用户发帖记录
// @Tags Community
// @Summary 获取用户发帖记录
// @Description 获取当前用户的发帖记录
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(20)
// @Success 200 {object} response.Response{data=response.UserPostsResponse} "获取成功"
// @Router /miniprogram/community/my-posts [get]
func (a *CommunityApi) GetUserPosts(c *gin.Context) {
	// 获取用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	var req request.GetPostsRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	communityService := service.ServiceGroupApp.CommunityService
	result, err := communityService.GetUserPosts(userID, req)
	if err != nil {
		global.GVA_LOG.Error("获取用户发帖记录失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}
