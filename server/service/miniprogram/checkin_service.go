package miniprogram

import (
	"errors"
	"math"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CheckinService struct{}

// DailyCheckin 每日打卡
func (s *CheckinService) DailyCheckin(userID uint, req request.CheckinRequest) (*response.CheckinResponse, error) {
	today := time.Now().Format("2006-01-02")

	// 检查今日是否已经打卡
	var existingCheckin miniprogram.DailyCheckin
	err := global.GVA_DB.Where("user_id = ? AND DATE(checkin_date) = ?", userID, today).First(&existingCheckin).Error
	if err == nil {
		return nil, errors.New("今日已经打卡过了")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 获取或创建戒色记录
	var abstinenceRecord miniprogram.AbstinenceRecord
	err = global.GVA_DB.Where("user_id = ? AND status = 1", userID).First(&abstinenceRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 创建新的戒色记录
		abstinenceRecord = miniprogram.AbstinenceRecord{
			UserID:        userID,
			StartDate:     time.Now(),
			CurrentStreak: 0,
			LongestStreak: 0,
			TotalDays:     0,
			SuccessRate:   0.00,
			Level:         1,
			Experience:    0,
			Status:        1,
		}
		err = global.GVA_DB.Create(&abstinenceRecord).Error
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	// 计算连续天数
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	var yesterdayCheckin miniprogram.DailyCheckin
	err = global.GVA_DB.Where("user_id = ? AND DATE(checkin_date) = ?", userID, yesterday).First(&yesterdayCheckin).Error

	var newStreak int
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 昨天没有打卡，重新开始
		newStreak = 1
	} else if err == nil {
		// 昨天有打卡，连续天数+1
		newStreak = abstinenceRecord.CurrentStreak + 1
	} else {
		return nil, err
	}

	// 使用游戏化服务计算奖励经验值
	gamificationService := GamificationService{}
	totalRewards := gamificationService.CalculateCheckinRewards(newStreak, req.MoodLevel)

	// 创建打卡记录
	checkin := miniprogram.DailyCheckin{
		UserID:      userID,
		CheckinDate: time.Now(),
		MoodLevel:   req.MoodLevel,
		Notes:       req.Notes,
		Rewards:     totalRewards,
		IsSuccess:   true,
		Streak:      newStreak,
	}

	// 开始数据库事务
	tx := global.GVA_DB.Begin()

	// 保存打卡记录
	err = tx.Create(&checkin).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 更新戒色记录
	oldLevel := abstinenceRecord.Level
	abstinenceRecord.CurrentStreak = newStreak
	if newStreak > abstinenceRecord.LongestStreak {
		abstinenceRecord.LongestStreak = newStreak
	}
	abstinenceRecord.TotalDays++
	abstinenceRecord.Experience += totalRewards

	// 计算新等级
	newLevel := s.calculateLevel(abstinenceRecord.Experience)
	levelUp := newLevel > oldLevel
	abstinenceRecord.Level = newLevel

	// 计算成功率
	totalCheckins := abstinenceRecord.TotalDays
	daysSinceStart := int(time.Since(abstinenceRecord.StartDate).Hours()/24) + 1
	abstinenceRecord.SuccessRate = float64(totalCheckins) / float64(daysSinceStart) * 100

	err = tx.Save(&abstinenceRecord).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 检查成就解锁
	var achievement *string
	if achievementName := s.checkAchievements(userID, newStreak, newLevel, tx); achievementName != "" {
		achievement = &achievementName
	}

	// 提交事务
	tx.Commit()

	// 构建响应
	resp := &response.CheckinResponse{
		ID:          checkin.ID,
		CheckinDate: checkin.CheckinDate,
		MoodLevel:   checkin.MoodLevel,
		Notes:       checkin.Notes,
		Rewards:     checkin.Rewards,
		Streak:      checkin.Streak,
		IsSuccess:   checkin.IsSuccess,
		LevelUp:     levelUp,
		NewLevel:    abstinenceRecord.Level,
		Experience:  abstinenceRecord.Experience,
		Achievement: achievement,
	}

	return resp, nil
}

// GetTodayCheckin 获取今日打卡状态
func (s *CheckinService) GetTodayCheckin(userID uint) (*response.TodayCheckinResponse, error) {
	today := time.Now().Format("2006-01-02")

	var checkin miniprogram.DailyCheckin
	err := global.GVA_DB.Where("user_id = ? AND DATE(checkin_date) = ?", userID, today).First(&checkin).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 今日未打卡
		return &response.TodayCheckinResponse{
			HasChecked: false,
		}, nil
	}

	if err != nil {
		return nil, err
	}

	// 今日已打卡
	return &response.TodayCheckinResponse{
		HasChecked:  true,
		CheckinDate: &checkin.CheckinDate,
		MoodLevel:   checkin.MoodLevel,
		Notes:       checkin.Notes,
		Rewards:     checkin.Rewards,
		CheckinTime: &checkin.CreatedAt,
	}, nil
}

// GetCheckinHistory 获取打卡历史
func (s *CheckinService) GetCheckinHistory(userID uint, page, pageSize int, month string) (*response.CheckinHistoryResponse, error) {
	offset := (page - 1) * pageSize

	query := global.GVA_DB.Where("user_id = ?", userID)

	// 如果指定了月份，添加月份筛选
	if month != "" {
		query = query.Where("DATE_FORMAT(checkin_date, '%Y-%m') = ?", month)
	}

	var total int64
	err := query.Model(&miniprogram.DailyCheckin{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	var checkins []miniprogram.DailyCheckin
	err = query.Order("checkin_date DESC").Offset(offset).Limit(pageSize).Find(&checkins).Error
	if err != nil {
		return nil, err
	}

	return &response.CheckinHistoryResponse{
		List:     checkins,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// GetCheckinStatistics 获取打卡统计
func (s *CheckinService) GetCheckinStatistics(userID uint) (*response.CheckinStatsResponse, error) {
	// 获取戒色记录
	var abstinenceRecord miniprogram.AbstinenceRecord
	err := global.GVA_DB.Where("user_id = ? AND status = 1", userID).First(&abstinenceRecord).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有记录，返回默认值
			return &response.CheckinStatsResponse{
				TotalDays:     0,
				CurrentStreak: 0,
				LongestStreak: 0,
				SuccessRate:   0,
				AverageMood:   0,
				TotalRewards:  0,
				ThisMonth:     0,
				ThisWeek:      0,
				Level:         1,
				Experience:    0,
			}, nil
		}
		return nil, err
	}

	// 计算总打卡天数
	var totalDays int64
	err = global.GVA_DB.Model(&miniprogram.DailyCheckin{}).Where("user_id = ?", userID).Count(&totalDays).Error
	if err != nil {
		return nil, err
	}

	// 计算平均心情
	var avgMood float64
	err = global.GVA_DB.Model(&miniprogram.DailyCheckin{}).Where("user_id = ?", userID).Select("AVG(mood_level)").Scan(&avgMood).Error
	if err != nil {
		return nil, err
	}

	// 计算总奖励
	var totalRewards int64
	err = global.GVA_DB.Model(&miniprogram.DailyCheckin{}).Where("user_id = ?", userID).Select("SUM(rewards)").Scan(&totalRewards).Error
	if err != nil {
		return nil, err
	}

	// 计算本月打卡天数
	thisMonth := time.Now().Format("2006-01")
	var thisMonthDays int64
	err = global.GVA_DB.Model(&miniprogram.DailyCheckin{}).Where("user_id = ? AND DATE_FORMAT(checkin_date, '%Y-%m') = ?", userID, thisMonth).Count(&thisMonthDays).Error
	if err != nil {
		return nil, err
	}

	// 计算本周打卡天数
	startOfWeek := s.getStartOfWeek(time.Now())
	endOfWeek := startOfWeek.AddDate(0, 0, 6)
	var thisWeekDays int64
	err = global.GVA_DB.Model(&miniprogram.DailyCheckin{}).Where("user_id = ? AND checkin_date BETWEEN ? AND ?", userID, startOfWeek, endOfWeek).Count(&thisWeekDays).Error
	if err != nil {
		return nil, err
	}

	return &response.CheckinStatsResponse{
		TotalDays:     int(totalDays),
		CurrentStreak: abstinenceRecord.CurrentStreak,
		LongestStreak: abstinenceRecord.LongestStreak,
		SuccessRate:   abstinenceRecord.SuccessRate,
		AverageMood:   math.Round(avgMood*10) / 10,
		TotalRewards:  int(totalRewards),
		ThisMonth:     int(thisMonthDays),
		ThisWeek:      int(thisWeekDays),
		Level:         abstinenceRecord.Level,
		Experience:    abstinenceRecord.Experience,
	}, nil
}

// GetWeeklyProgress 获取本周进度
func (s *CheckinService) GetWeeklyProgress(userID uint) (*response.WeeklyProgressResponse, error) {
	startOfWeek := s.getStartOfWeek(time.Now())
	weekDays := make([]response.WeekDay, 7)
	weekdays := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}

	var checkedDays int
	var totalMood int
	today := time.Now().Format("2006-01-02")

	for i := 0; i < 7; i++ {
		date := startOfWeek.AddDate(0, 0, i)
		dateStr := date.Format("2006-01-02")

		// 查询当天的打卡记录
		var checkin miniprogram.DailyCheckin
		err := global.GVA_DB.Where("user_id = ? AND DATE(checkin_date) = ?", userID, dateStr).First(&checkin).Error

		hasChecked := !errors.Is(err, gorm.ErrRecordNotFound)
		moodLevel := 0

		if hasChecked {
			checkedDays++
			moodLevel = checkin.MoodLevel
			totalMood += moodLevel
		}

		weekDays[i] = response.WeekDay{
			Date:       dateStr,
			Weekday:    weekdays[i],
			HasChecked: hasChecked,
			MoodLevel:  moodLevel,
			IsToday:    dateStr == today,
		}
	}

	// 计算本周统计
	successRate := float64(checkedDays) / 7.0 * 100
	averageMood := 0.0
	if checkedDays > 0 {
		averageMood = float64(totalMood) / float64(checkedDays)
	}

	summary := response.WeekSummary{
		CheckedDays: checkedDays,
		TotalDays:   7,
		SuccessRate: math.Round(successRate*10) / 10,
		AverageMood: math.Round(averageMood*10) / 10,
	}

	return &response.WeeklyProgressResponse{
		WeekDays: weekDays,
		Summary:  summary,
	}, nil
}

// GetMonthlyCalendar 获取月度日历
func (s *CheckinService) GetMonthlyCalendar(userID uint, year, month int) (*response.MonthlyCalendarResponse, error) {
	// 获取月份的第一天和最后一天
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastDay := firstDay.AddDate(0, 1, -1)

	// 获取本月的打卡记录
	var checkins []miniprogram.DailyCheckin
	err := global.GVA_DB.Where("user_id = ? AND checkin_date BETWEEN ? AND ?", userID, firstDay, lastDay).Find(&checkins).Error
	if err != nil {
		return nil, err
	}

	// 创建日期映射
	checkinMap := make(map[string]miniprogram.DailyCheckin)
	for _, checkin := range checkins {
		date := checkin.CheckinDate.Format("2006-01-02")
		checkinMap[date] = checkin
	}

	// 生成日历数据
	days := make([]response.CalendarDay, 0)
	today := time.Now().Format("2006-01-02")

	var checkedDays int
	var totalMood int
	var currentStreak int
	var bestStreak int

	for day := 1; day <= lastDay.Day(); day++ {
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		dateStr := date.Format("2006-01-02")

		checkin, hasChecked := checkinMap[dateStr]
		moodLevel := 0

		if hasChecked {
			checkedDays++
			moodLevel = checkin.MoodLevel
			totalMood += moodLevel
			currentStreak++
			if currentStreak > bestStreak {
				bestStreak = currentStreak
			}
		} else {
			currentStreak = 0
		}

		days = append(days, response.CalendarDay{
			Day:         day,
			Date:        dateStr,
			HasChecked:  hasChecked,
			MoodLevel:   moodLevel,
			IsToday:     dateStr == today,
			IsThisMonth: true,
		})
	}

	// 计算月度统计
	successRate := float64(checkedDays) / float64(lastDay.Day()) * 100
	averageMood := 0.0
	if checkedDays > 0 {
		averageMood = float64(totalMood) / float64(checkedDays)
	}

	summary := response.MonthSummary{
		CheckedDays: checkedDays,
		TotalDays:   lastDay.Day(),
		SuccessRate: math.Round(successRate*10) / 10,
		AverageMood: math.Round(averageMood*10) / 10,
		BestStreak:  bestStreak,
	}

	return &response.MonthlyCalendarResponse{
		Year:    year,
		Month:   month,
		Days:    days,
		Summary: summary,
	}, nil
}

// 注意：奖励计算方法已迁移到 GamificationService 中统一管理

// calculateLevel 根据经验值计算等级
func (s *CheckinService) calculateLevel(experience int) int {
	// 简单的等级计算：每100经验值升1级
	level := experience/100 + 1
	if level > 50 {
		level = 50 // 最高50级
	}
	return level
}

// checkAchievements 检查成就解锁
func (s *CheckinService) checkAchievements(userID uint, streak, level int, tx *gorm.DB) string {
	// 调用成就服务检查解锁
	achievementService := AchievementService{}
	unlockedAchievements, err := achievementService.CheckAchievements(userID, tx)
	if err != nil {
		global.GVA_LOG.Error("检查成就失败", zap.Error(err))
		return ""
	}

	// 返回第一个解锁的成就名称
	if len(unlockedAchievements) > 0 {
		return unlockedAchievements[0]
	}

	return ""
}

// getStartOfWeek 获取一周的开始时间（周一）
func (s *CheckinService) getStartOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	return t.AddDate(0, 0, -int(weekday-1)).Truncate(24 * time.Hour)
}

// GetWeeklyProgressForChart 获取本周进度数据供图表使用
func (s *CheckinService) GetWeeklyProgressForChart(userID uint) (*response.WeeklyChartResponse, error) {
	startOfWeek := s.getStartOfWeek(time.Now())
	weekdays := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}

	weeklyCheckins := make([]response.WeeklyCheckinData, 7)
	checkedCount := 0

	for i := 0; i < 7; i++ {
		date := startOfWeek.AddDate(0, 0, i)
		dateStr := date.Format("2006-01-02")

		// 查询当天的打卡记录
		var checkin miniprogram.DailyCheckin
		err := global.GVA_DB.Where("user_id = ? AND DATE(checkin_date) = ?", userID, dateStr).First(&checkin).Error

		hasChecked := !errors.Is(err, gorm.ErrRecordNotFound)
		moodLevel := 0

		if hasChecked {
			checkedCount++
			moodLevel = checkin.MoodLevel
		}

		weeklyCheckins[i] = response.WeeklyCheckinData{
			Date:       dateStr,
			Weekday:    weekdays[i],
			HasChecked: hasChecked,
			MoodLevel:  moodLevel,
		}
	}

	// 计算完成率
	completionRate := float64(checkedCount) / 7.0 * 100

	return &response.WeeklyChartResponse{
		WeeklyCheckins: weeklyCheckins,
		CompletionRate: math.Round(completionRate*10) / 10,
	}, nil
}
