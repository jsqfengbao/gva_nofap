package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AchievementApi struct{}

// GetUserAchievements 获取用户成就列表
// @Tags Achievement
// @Summary 获取用户成就列表
// @Description 获取用户的所有成就，包括已解锁和未解锁的
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=response.UserAchievementsResponse} "成功"
// @Router /api/v1/miniprogram/achievement/list [get]
func (achievementApi *AchievementApi) GetUserAchievements(c *gin.Context) {
	// 获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未登录", c)
		return
	}

	uid := userID.(uint)

	// 调用服务
	achievementService := service.ServiceGroupApp.AchievementService
	achievements, err := achievementService.GetUserAchievements(uid)
	if err != nil {
		global.GVA_LOG.Error("获取用户成就失败!", zap.Error(err))
		response.FailWithMessage("获取用户成就失败", c)
		return
	}

	response.OkWithData(achievements, c)
}

// GetAchievementStats 获取成就统计
// @Tags Achievement
// @Summary 获取成就统计信息
// @Description 获取用户的成就统计数据
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=response.AchievementStatsResponse} "成功"
// @Router /api/v1/miniprogram/achievement/stats [get]
func (achievementApi *AchievementApi) GetAchievementStats(c *gin.Context) {
	// 获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未登录", c)
		return
	}

	uid := userID.(uint)

	// 调用服务
	achievementService := service.ServiceGroupApp.AchievementService
	stats, err := achievementService.GetAchievementStats(uid)
	if err != nil {
		global.GVA_LOG.Error("获取成就统计失败!", zap.Error(err))
		response.FailWithMessage("获取成就统计失败", c)
		return
	}

	response.OkWithData(stats, c)
}

// GetGameStats 获取游戏化统计数据
// @Tags Achievement
// @Summary 获取游戏化统计数据
// @Description 获取用户的等级、经验值、成就等游戏化数据
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=response.GameStatsResponse} "成功"
// @Router /api/v1/miniprogram/achievement/game-stats [get]
func (achievementApi *AchievementApi) GetGameStats(c *gin.Context) {
	// 获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未登录", c)
		return
	}

	uid := userID.(uint)

	// 调用服务
	gamificationService := service.ServiceGroupApp.GamificationService
	stats, err := gamificationService.GetGameStats(uid)
	if err != nil {
		global.GVA_LOG.Error("获取游戏化统计失败!", zap.Error(err))
		response.FailWithMessage("获取游戏化统计失败", c)
		return
	}

	response.OkWithData(stats, c)
}

// GetLevelProgress 获取等级进度
// @Tags Achievement
// @Summary 获取等级进度信息
// @Description 获取用户的等级进度详细信息
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=response.LevelProgressResponse} "成功"
// @Router /api/v1/miniprogram/achievement/level-progress [get]
func (achievementApi *AchievementApi) GetLevelProgress(c *gin.Context) {
	// 获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未登录", c)
		return
	}

	uid := userID.(uint)

	// 调用服务
	gamificationService := service.ServiceGroupApp.GamificationService
	progress, err := gamificationService.GetLevelProgress(uid)
	if err != nil {
		global.GVA_LOG.Error("获取等级进度失败!", zap.Error(err))
		response.FailWithMessage("获取等级进度失败", c)
		return
	}

	response.OkWithData(progress, c)
}

// GetAchievementProgress 获取成就进度信息
// @Tags Achievement
// @Summary 获取成就进度信息
// @Description 获取用户正在进行中的成就的进度信息
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=[]response.AchievementProgressResponse} "成功"
// @Router /api/v1/miniprogram/achievement/progress [get]
func (achievementApi *AchievementApi) GetAchievementProgress(c *gin.Context) {
	// 获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未登录", c)
		return
	}

	uid := userID.(uint)

	// 调用服务
	achievementService := service.ServiceGroupApp.AchievementService
	progress, err := achievementService.GetAchievementProgress(uid)
	if err != nil {
		global.GVA_LOG.Error("获取成就进度失败!", zap.Error(err))
		response.FailWithMessage("获取成就进度失败", c)
		return
	}

	response.OkWithData(progress, c)
}

// GetUserAchievementsForProfile 获取用户成就（用于个人中心）
// @Tags Achievement
// @Summary 获取用户成就（用于个人中心）
// @Description 获取用户的成就数据，支持限制数量和仅显示最近解锁的成就
// @Security ApiKeyAuth
// @Produce application/json
// @Param limit query int false "限制返回数量"
// @Param recent query bool false "是否只返回最近解锁的成就"
// @Success 200 {object} response.Response{data=response.UserProfileAchievementsResponse} "成功"
// @Router /api/v1/miniprogram/achievement/user [get]
func (achievementApi *AchievementApi) GetUserAchievementsForProfile(c *gin.Context) {
	// 获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage("用户未登录", c)
		return
	}

	uid := userID.(uint)

	// 获取查询参数
	limitStr := c.Query("limit")
	recentStr := c.Query("recent")

	limit := 10 // 默认限制
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	recent := false
	if recentStr == "true" {
		recent = true
	}

	// 调用服务
	achievementService := service.ServiceGroupApp.AchievementService
	achievements, err := achievementService.GetUserAchievementsForProfile(uid, limit, recent)
	if err != nil {
		global.GVA_LOG.Error("获取用户成就失败!", zap.Error(err))
		response.FailWithMessage("获取用户成就失败", c)
		return
	}

	response.OkWithData(achievements, c)
}
