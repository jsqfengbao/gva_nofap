package api

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	miniprogramRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StatsApi struct{}

// GetOverall 获取总体统计数据
// @Tags      MiniProgram
// @Summary   获取总体统计数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=miniprogramRes.OverallStatsResponse,msg=string}  "获取成功"
// @Router    /miniprogram/stats/overall [get]
func (s *StatsApi) GetOverall(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	// 获取戒色记录
	var abstinenceRecord struct {
		CurrentStreak int     `json:"current_streak"`
		LongestStreak int     `json:"longest_streak"`
		TotalDays     int     `json:"total_days"`
		SuccessRate   float64 `json:"success_rate"`
		StartDate     string  `json:"start_date"`
		Level         int     `json:"level"`
		Experience    int     `json:"experience"`
	}
	err := global.GVA_DB.Table("nofap_abstinence_records").
		Select("current_streak, longest_streak, total_days, success_rate, start_date, level, experience").
		Where("user_id = ? AND status = 1", userID).
		First(&abstinenceRecord).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		global.GVA_LOG.Error("获取戒色记录失败", zap.Error(err))
		response.FailWithMessage("获取统计数据失败", c)
		return
	}

	// 获取签到统计
	var checkinStats struct {
		TotalCheckins int `json:"total_checkins"`
		MonthCheckins int `json:"month_checkins"`
	}
	monthAgo := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	global.GVA_DB.Table("nofap_daily_checkins").
		Select("COUNT(*) as total_checkins, SUM(CASE WHEN checkin_date >= ? THEN 1 ELSE 0 END) as month_checkins", monthAgo).
		Where("user_id = ?", userID).
		Row().Scan(&checkinStats.TotalCheckins, &checkinStats.MonthCheckins)

	// 获取成就统计
	var achievementStats struct {
		TotalAchievements    int64 `json:"total_achievements"`
		UnlockedAchievements int64 `json:"unlocked_achievements"`
	}
	global.GVA_DB.Table("nofap_achievements").Count(&achievementStats.TotalAchievements)
	global.GVA_DB.Table("nofap_user_achievements").
		Where("user_id = ? AND unlocked_at IS NOT NULL", userID).
		Count(&achievementStats.UnlockedAchievements)

	// 获取学习统计
	var learningStats struct {
		TotalArticles int64 `json:"total_articles"`
		ReadArticles  int `json:"read_articles"`
		TotalReadTime int `json:"total_read_time"`
	}
	global.GVA_DB.Table("nofap_learning_content").Where("status = 1").Count(&learningStats.TotalArticles)
	global.GVA_DB.Table("nofap_user_learning_records").
		Select("COUNT(*) as read_articles, COALESCE(SUM(read_duration), 0) as total_read_time").
		Where("user_id = ?", userID).
		Row().Scan(&learningStats.ReadArticles, &learningStats.TotalReadTime)

	// 获取社区统计
	var communityStats struct {
		TotalPosts int `json:"total_posts"`
		TotalLikes int `json:"total_likes"`
	}
	global.GVA_DB.Table("nofap_community_posts").
		Select("COUNT(*) as total_posts, COALESCE(SUM(like_count), 0) as total_likes").
		Where("user_id = ? AND status = 1", userID).
		Row().Scan(&communityStats.TotalPosts, &communityStats.TotalLikes)

	result := miniprogramRes.OverallStatsResponse{
		Abstinence: miniprogramRes.AbstinenceStats{
			CurrentStreak: abstinenceRecord.CurrentStreak,
			LongestStreak: abstinenceRecord.LongestStreak,
			TotalDays:     abstinenceRecord.TotalDays,
			SuccessRate:   abstinenceRecord.SuccessRate,
			StartDate:     abstinenceRecord.StartDate,
			Level:         abstinenceRecord.Level,
			Experience:    abstinenceRecord.Experience,
		},
		Checkin: miniprogramRes.CheckinStats{
			TotalCheckins: checkinStats.TotalCheckins,
			MonthCheckins: checkinStats.MonthCheckins,
		},
		Achievement: miniprogramRes.AchievementStats{
			TotalAchievements:    int(achievementStats.TotalAchievements),
			UnlockedAchievements: int(achievementStats.UnlockedAchievements),
		},
		Learning: miniprogramRes.LearningStats{
			TotalArticles: int(learningStats.TotalArticles),
			ReadArticles:  learningStats.ReadArticles,
			TotalReadTime: learningStats.TotalReadTime,
		},
		Community: miniprogramRes.CommunityStats{
			TotalPosts: communityStats.TotalPosts,
			TotalLikes: communityStats.TotalLikes,
		},
	}

	response.OkWithData(result, c)
}

// GetTrends 获取趋势数据
// @Tags      MiniProgram
// @Summary   获取趋势数据（近30天）
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=miniprogramRes.TrendsResponse,msg=string}  "获取成功"
// @Router    /miniprogram/stats/trends [get]
func (s *StatsApi) GetTrends(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	// 获取近30天的打卡数据
	daysAgo := time.Now().AddDate(0, 0, -30)
	var checkinData []struct {
		CheckinDate string `json:"date"`
		IsSuccess   bool   `json:"is_success"`
		Streak      int    `json:"streak"`
	}

	global.GVA_DB.Table("nofap_daily_checkins").
		Select("checkin_date, is_success, current_streak as streak").
		Where("user_id = ? AND checkin_date >= ?", userID, daysAgo.Format("2006-01-02")).
		Order("checkin_date ASC").
		Find(&checkinData)

	// 计算每日经验增长（这里简化处理）
	var trendPoints []miniprogramRes.TrendPoint
	for _, d := range checkinData {
		point := miniprogramRes.TrendPoint{
			Date:    d.CheckinDate,
			Success: d.IsSuccess,
			Streak:  d.Streak,
		}
		trendPoints = append(trendPoints, point)
	}

	result := miniprogramRes.TrendsResponse{
		Trends: trendPoints,
	}

	response.OkWithData(result, c)
}
