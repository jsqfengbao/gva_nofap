package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model/request"
	miniResponse "github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmergencyApi struct{}

// CreateEmergencyHelp 创建紧急求助
// @Tags Emergency
// @Summary 创建紧急求助
// @Description 用户创建紧急求助请求
// @Accept json
// @Produce json
// @Param data body request.CreateEmergencyHelpRequest true "求助信息"
// @Success 200 {object} response.Response{data=miniprogram.EmergencyHelp} "创建成功"
// @Router /emergency/help [post]
// @Security ApiKeyAuth
func (e *EmergencyApi) CreateEmergencyHelp(c *gin.Context) {
	var req request.CreateEmergencyHelpRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取用户ID
	userID := utils.GetUserID(c)

	help, err := service.ServiceGroupApp.EmergencyService.CreateEmergencyHelp(userID, req)
	if err != nil {
		global.GVA_LOG.Error("创建紧急求助失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}

	response.OkWithData(help, c)
}

// GetEmergencyHelps 获取紧急求助列表
// @Tags Emergency
// @Summary 获取紧急求助列表
// @Description 获取用户的紧急求助记录列表
// @Accept json
// @Produce json
// @Param data query request.GetEmergencyHelpsRequest true "查询参数"
// @Success 200 {object} response.Response{data=miniResponse.EmergencyHelpListResponse} "获取成功"
// @Router /emergency/help/list [get]
// @Security ApiKeyAuth
func (e *EmergencyApi) GetEmergencyHelps(c *gin.Context) {
	var req request.GetEmergencyHelpsRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 设置默认分页
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	userID := utils.GetUserID(c)

	list, err := service.ServiceGroupApp.EmergencyService.GetEmergencyHelps(userID, req)
	if err != nil {
		global.GVA_LOG.Error("获取紧急求助列表失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}

// GetEmergencyHelpDetail 获取紧急求助详情
// @Tags Emergency
// @Summary 获取紧急求助详情
// @Description 获取特定紧急求助的详细信息
// @Accept json
// @Produce json
// @Param id path int true "求助ID"
// @Success 200 {object} response.Response{data=miniResponse.EmergencyHelpDetailResponse} "获取成功"
// @Router /emergency/help/{id} [get]
// @Security ApiKeyAuth
func (e *EmergencyApi) GetEmergencyHelpDetail(c *gin.Context) {
	idStr := c.Param("id")
	helpID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的求助ID", c)
		return
	}

	userID := utils.GetUserID(c)

	detail, err := service.ServiceGroupApp.EmergencyService.GetEmergencyHelpDetail(userID, uint(helpID))
	if err != nil {
		global.GVA_LOG.Error("获取紧急求助详情失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithData(detail, c)
}

// UpdateEmergencyHelp 更新紧急求助
// @Tags Emergency
// @Summary 更新紧急求助
// @Description 更新紧急求助状态或信息
// @Accept json
// @Produce json
// @Param data body request.UpdateEmergencyHelpRequest true "更新信息"
// @Success 200 {object} response.Response "更新成功"
// @Router /emergency/help [put]
// @Security ApiKeyAuth
func (e *EmergencyApi) UpdateEmergencyHelp(c *gin.Context) {
	var req request.UpdateEmergencyHelpRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)

	err = service.ServiceGroupApp.EmergencyService.UpdateEmergencyHelp(userID, req)
	if err != nil {
		global.GVA_LOG.Error("更新紧急求助失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// CreateEmergencyResponse 创建紧急求助响应
// @Tags Emergency
// @Summary 创建紧急求助响应
// @Description 志愿者响应紧急求助
// @Accept json
// @Produce json
// @Param data body request.CreateEmergencyResponseRequest true "响应信息"
// @Success 200 {object} response.Response "响应成功"
// @Router /emergency/response [post]
// @Security ApiKeyAuth
func (e *EmergencyApi) CreateEmergencyResponse(c *gin.Context) {
	var req request.CreateEmergencyResponseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	volunteerID := utils.GetUserID(c)

	err = service.ServiceGroupApp.EmergencyService.CreateEmergencyResponse(volunteerID, req)
	if err != nil {
		global.GVA_LOG.Error("创建紧急求助响应失败!", zap.Error(err))
		response.FailWithMessage("响应失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("响应成功", c)
}

// ConnectVolunteer 连接志愿者
// @Tags Emergency
// @Summary 连接志愿者
// @Description 用户主动连接志愿者寻求帮助
// @Accept json
// @Produce json
// @Param data body request.ConnectVolunteerRequest true "连接信息"
// @Success 200 {object} response.Response{data=miniResponse.ConnectVolunteerResponse} "连接成功"
// @Router /emergency/connect [post]
// @Security ApiKeyAuth
func (e *EmergencyApi) ConnectVolunteer(c *gin.Context) {
	var req request.ConnectVolunteerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)

	result, err := service.ServiceGroupApp.EmergencyService.ConnectVolunteer(userID, req)
	if err != nil {
		global.GVA_LOG.Error("连接志愿者失败!", zap.Error(err))
		response.FailWithMessage("连接失败: "+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// RegisterVolunteer 注册志愿者
// @Tags Emergency
// @Summary 注册志愿者
// @Description 用户申请成为紧急求助志愿者
// @Accept json
// @Produce json
// @Param data body request.EmergencyVolunteerRequest true "志愿者信息"
// @Success 200 {object} response.Response "注册成功"
// @Router /emergency/volunteer/register [post]
// @Security ApiKeyAuth
func (e *EmergencyApi) RegisterVolunteer(c *gin.Context) {
	var req request.EmergencyVolunteerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)

	err = service.ServiceGroupApp.EmergencyService.RegisterVolunteer(userID, req)
	if err != nil {
		global.GVA_LOG.Error("注册志愿者失败!", zap.Error(err))
		response.FailWithMessage("注册失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("注册成功，等待审核", c)
}

// UpdateVolunteerStatus 更新志愿者状态
// @Tags Emergency
// @Summary 更新志愿者状态
// @Description 志愿者更新在线状态
// @Accept json
// @Produce json
// @Param data body request.UpdateVolunteerStatusRequest true "状态信息"
// @Success 200 {object} response.Response "更新成功"
// @Router /emergency/volunteer/status [put]
// @Security ApiKeyAuth
func (e *EmergencyApi) UpdateVolunteerStatus(c *gin.Context) {
	var req request.UpdateVolunteerStatusRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)

	err = service.ServiceGroupApp.EmergencyService.UpdateVolunteerStatus(userID, req)
	if err != nil {
		global.GVA_LOG.Error("更新志愿者状态失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// GetEmergencyResources 获取紧急求助资源
// @Tags Emergency
// @Summary 获取紧急求助资源
// @Description 获取紧急情况下的辅助资源(呼吸练习、冥想、音乐等)
// @Accept json
// @Produce json
// @Param data query request.GetEmergencyResourcesRequest true "查询参数"
// @Success 200 {object} response.Response{data=miniResponse.ResourceListResponse} "获取成功"
// @Router /emergency/resources [get]
// @Security ApiKeyAuth
func (e *EmergencyApi) GetEmergencyResources(c *gin.Context) {
	var req request.GetEmergencyResourcesRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 设置默认分页
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, err := service.ServiceGroupApp.EmergencyService.GetEmergencyResources(req)
	if err != nil {
		global.GVA_LOG.Error("获取紧急求助资源失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}

// UseResource 使用资源
// @Tags Emergency
// @Summary 使用资源
// @Description 记录用户使用某个紧急求助资源
// @Accept json
// @Produce json
// @Param id path int true "资源ID"
// @Success 200 {object} response.Response "使用成功"
// @Router /emergency/resources/{id}/use [post]
// @Security ApiKeyAuth
func (e *EmergencyApi) UseResource(c *gin.Context) {
	idStr := c.Param("id")
	resourceID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的资源ID", c)
		return
	}

	err = service.ServiceGroupApp.EmergencyService.UseResource(uint(resourceID))
	if err != nil {
		global.GVA_LOG.Error("使用资源失败!", zap.Error(err))
		response.FailWithMessage("使用失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("使用成功", c)
}

// RateResource 资源评分
// @Tags Emergency
// @Summary 资源评分
// @Description 用户对紧急求助资源进行评分
// @Accept json
// @Produce json
// @Param data body request.RateResourceRequest true "评分信息"
// @Success 200 {object} response.Response "评分成功"
// @Router /emergency/resources/rate [post]
// @Security ApiKeyAuth
func (e *EmergencyApi) RateResource(c *gin.Context) {
	var req request.RateResourceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)

	err = service.ServiceGroupApp.EmergencyService.RateResource(userID, req)
	if err != nil {
		global.GVA_LOG.Error("资源评分失败!", zap.Error(err))
		response.FailWithMessage("评分失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("评分成功", c)
}

// GetEmergencyStats 获取紧急求助统计
// @Tags Emergency
// @Summary 获取紧急求助统计
// @Description 获取紧急求助相关的统计数据
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=miniResponse.EmergencyStatsResponse} "获取成功"
// @Router /emergency/stats [get]
// @Security ApiKeyAuth
func (e *EmergencyApi) GetEmergencyStats(c *gin.Context) {
	stats, err := service.ServiceGroupApp.EmergencyService.GetEmergencyStats()
	if err != nil {
		global.GVA_LOG.Error("获取紧急求助统计失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithData(stats, c)
}

// GetOnlineVolunteers 获取在线志愿者
// @Tags Emergency
// @Summary 获取在线志愿者
// @Description 获取当前在线的志愿者列表
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=miniResponse.OnlineVolunteersResponse} "获取成功"
// @Router /emergency/volunteers/online [get]
// @Security ApiKeyAuth
func (e *EmergencyApi) GetOnlineVolunteers(c *gin.Context) {
	// 简化实现，获取在线志愿者数量
	var count int64
	global.GVA_DB.Model(&miniprogram.EmergencyVolunteer{}).
		Where("is_online = ? AND status = ?", true, 2).Count(&count)

	result := miniResponse.OnlineVolunteersResponse{
		Count:      int(count),
		Volunteers: []miniResponse.EmergencyVolunteerItem{}, // 简化处理，不返回具体志愿者信息
	}

	response.OkWithData(result, c)
}
