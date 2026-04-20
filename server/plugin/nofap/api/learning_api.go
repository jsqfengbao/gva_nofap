package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LearningApi struct{}

// CreateLearningContent
// @Tags Learning
// @Summary 创建学习内容
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateLearningContentRequest true "创建学习内容请求"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /learning/content [post]
func (l *LearningApi) CreateLearningContent(c *gin.Context) {
	var req request.CreateLearningContentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	content, err := service.ServiceGroupApp.LearningService.CreateLearningContent(req)
	if err != nil {
		global.GVA_LOG.Error("创建学习内容失败!", zap.Error(err))
		response.FailWithMessage("创建学习内容失败: "+err.Error(), c)
		return
	}

	response.OkWithData(content, c)
}

// UpdateLearningContent
// @Tags Learning
// @Summary 更新学习内容
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateLearningContentRequest true "更新学习内容请求"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /learning/content [put]
func (l *LearningApi) UpdateLearningContent(c *gin.Context) {
	var req request.UpdateLearningContentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.LearningService.UpdateLearningContent(req)
	if err != nil {
		global.GVA_LOG.Error("更新学习内容失败!", zap.Error(err))
		response.FailWithMessage("更新学习内容失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteLearningContent
// @Tags Learning
// @Summary 删除学习内容
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "内容ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /learning/content/{id} [delete]
func (l *LearningApi) DeleteLearningContent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的内容ID", c)
		return
	}

	err = service.ServiceGroupApp.LearningService.DeleteLearningContent(uint(id))
	if err != nil {
		global.GVA_LOG.Error("删除学习内容失败!", zap.Error(err))
		response.FailWithMessage("删除学习内容失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetLearningStats
// @Tags Learning
// @Summary 获取学习统计
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=miniprogramRes.LearningStatsResponse,msg=string} "获取成功"
// @Router /learning/stats [get]
func (l *LearningApi) GetLearningStats(c *gin.Context) {
	// 获取当前用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	stats, err := service.ServiceGroupApp.LearningService.GetLearningStats(userID)
	if err != nil {
		global.GVA_LOG.Error("获取学习统计失败!", zap.Error(err))
		response.FailWithMessage("获取学习统计失败: "+err.Error(), c)
		return
	}

	response.OkWithData(stats, c)
}

// 注意：其他学习相关的API方法暂时移除，等待相关请求类型和服务方法定义完成后再添加
// 这包括：GetLearningContent, GetLearningContents, StartLearning, UpdateLearningProgress,
// ToggleLike, ToggleCollect, GetRecommendations, GetUserLearningRecords, RateLearningContent,
// GetCategoryStats, GetLearningHomepage
