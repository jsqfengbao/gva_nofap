package api

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CheckinApi struct{}

var checkinService = service.ServiceGroupApp.CheckinService

// DailyCheckin 每日打卡
// @Tags CheckinApi
// @Summary 每日打卡
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CheckinRequest true "打卡数据"
// @Success 200 {object} response.Response{data=miniresponse.CheckinResponse,msg=string} "每日打卡成功"
// @Router /miniprogram/checkin/daily [post]
func (c *CheckinApi) DailyCheckin(ctx *gin.Context) {
	var req request.CheckinRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 获取用户ID
	userID := ctx.GetUint("userID")
	if userID == 0 {
		response.FailWithMessage("用户未登录", ctx)
		return
	}

	// 验证心情等级
	if req.MoodLevel < 1 || req.MoodLevel > 5 {
		response.FailWithMessage("心情等级必须在1-5之间", ctx)
		return
	}

	// 调用服务层
	checkinResp, err := checkinService.DailyCheckin(userID, req)
	if err != nil {
		global.GVA_LOG.Error("每日打卡失败!", zap.Error(err))
		response.FailWithMessage("打卡失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(checkinResp, ctx)
}

// GetCheckinHistory 获取打卡历史
// @Tags CheckinApi
// @Summary 获取打卡历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param month query string false "月份筛选 YYYY-MM"
// @Success 200 {object} response.Response{data=miniresponse.CheckinHistoryResponse,msg=string} "获取打卡历史成功"
// @Router /miniprogram/checkin/history [get]
func (c *CheckinApi) GetCheckinHistory(ctx *gin.Context) {
	// 获取用户ID
	userID := ctx.GetUint("userID")
	if userID == 0 {
		response.FailWithMessage("用户未登录", ctx)
		return
	}

	// 获取分页参数
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "30"))
	month := ctx.Query("month")

	// 调用服务层
	historyResp, err := checkinService.GetCheckinHistory(userID, page, pageSize, month)
	if err != nil {
		global.GVA_LOG.Error("获取打卡历史失败!", zap.Error(err))
		response.FailWithMessage("获取历史失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(historyResp, ctx)
}

// GetTodayCheckin 检查今日打卡状态
// @Tags CheckinApi
// @Summary 检查今日打卡状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=miniresponse.TodayCheckinResponse,msg=string} "获取今日打卡状态成功"
// @Router /miniprogram/checkin/today [get]
func (c *CheckinApi) GetTodayCheckin(ctx *gin.Context) {
	// 获取用户ID
	userID := ctx.GetUint("userID")
	if userID == 0 {
		response.FailWithMessage("用户未登录", ctx)
		return
	}

	// 调用服务层
	todayResp, err := checkinService.GetTodayCheckin(userID)
	if err != nil {
		global.GVA_LOG.Error("获取今日打卡状态失败!", zap.Error(err))
		response.FailWithMessage("获取状态失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(todayResp, ctx)
}

// GetCheckinStatistics 获取打卡统计
// @Tags CheckinApi
// @Summary 获取打卡统计
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=miniresponse.CheckinStatsResponse,msg=string} "获取打卡统计成功"
// @Router /miniprogram/checkin/statistics [get]
func (c *CheckinApi) GetCheckinStatistics(ctx *gin.Context) {
	// 获取用户ID
	userID := ctx.GetUint("userID")
	if userID == 0 {
		response.FailWithMessage("用户未登录", ctx)
		return
	}

	// 调用服务层
	statsResp, err := checkinService.GetCheckinStatistics(userID)
	if err != nil {
		global.GVA_LOG.Error("获取打卡统计失败!", zap.Error(err))
		response.FailWithMessage("获取统计失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(statsResp, ctx)
}

// GetWeeklyProgress 获取本周进度
// @Tags CheckinApi
// @Summary 获取本周进度
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=miniresponse.WeeklyProgressResponse,msg=string} "获取本周进度成功"
// @Router /miniprogram/checkin/weekly [get]
func (c *CheckinApi) GetWeeklyProgress(ctx *gin.Context) {
	// 获取用户ID
	userID := ctx.GetUint("userID")
	if userID == 0 {
		response.FailWithMessage("用户未登录", ctx)
		return
	}

	// 调用服务层
	weeklyResp, err := checkinService.GetWeeklyProgress(userID)
	if err != nil {
		global.GVA_LOG.Error("获取本周进度失败!", zap.Error(err))
		response.FailWithMessage("获取进度失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(weeklyResp, ctx)
}

// GetMonthlyCalendar 获取月度日历
// @Tags CheckinApi
// @Summary 获取月度日历
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param year query int false "年份"
// @Param month query int false "月份"
// @Success 200 {object} response.Response{data=miniresponse.MonthlyCalendarResponse,msg=string} "获取月度日历成功"
// @Router /miniprogram/checkin/calendar [get]
func (c *CheckinApi) GetMonthlyCalendar(ctx *gin.Context) {
	// 获取用户ID
	userID := ctx.GetUint("userID")
	if userID == 0 {
		response.FailWithMessage("用户未登录", ctx)
		return
	}

	// 获取年月参数
	now := time.Now()
	year, _ := strconv.Atoi(ctx.DefaultQuery("year", strconv.Itoa(now.Year())))
	month, _ := strconv.Atoi(ctx.DefaultQuery("month", strconv.Itoa(int(now.Month()))))

	// 验证参数
	if year < 2020 || year > 2030 {
		response.FailWithMessage("年份参数无效", ctx)
		return
	}
	if month < 1 || month > 12 {
		response.FailWithMessage("月份参数无效", ctx)
		return
	}

	// 调用服务层
	calendarResp, err := checkinService.GetMonthlyCalendar(userID, year, month)
	if err != nil {
		global.GVA_LOG.Error("获取月度日历失败!", zap.Error(err))
		response.FailWithMessage("获取日历失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(calendarResp, ctx)
}

// GetWeeklyProgressForChart 获取本周进度数据供图表使用
// @Tags CheckinApi
// @Summary 获取本周进度图表数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=miniresponse.WeeklyChartResponse,msg=string} "获取本周进度图表数据成功"
// @Router /miniprogram/checkin/weekly-progress [get]
func (c *CheckinApi) GetWeeklyProgressForChart(ctx *gin.Context) {
	// 获取用户ID
	userID := ctx.GetUint("userID")
	if userID == 0 {
		response.FailWithMessage("用户未登录", ctx)
		return
	}

	// 调用服务层
	chartResp, err := checkinService.GetWeeklyProgressForChart(userID)
	if err != nil {
		global.GVA_LOG.Error("获取本周进度图表数据失败!", zap.Error(err))
		response.FailWithMessage("获取进度数据失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(chartResp, ctx)
}
