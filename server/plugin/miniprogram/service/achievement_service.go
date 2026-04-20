package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AchievementService struct{}

// AchievementCondition 成就条件结构
type AchievementCondition struct {
	CheckinCount          int     `json:"checkin_count,omitempty"`
	StreakDays            int     `json:"streak_days,omitempty"`
	Level                 int     `json:"level,omitempty"`
	PostCount             int     `json:"post_count,omitempty"`
	PostLikes             int     `json:"post_likes,omitempty"`
	TotalLikes            int     `json:"total_likes,omitempty"`
	LearningCompleted     int     `json:"learning_completed,omitempty"`
	LearningHours         int     `json:"learning_hours,omitempty"`
	AllTypesCompleted     bool    `json:"all_types_completed,omitempty"`
	HelpResponses         int     `json:"help_responses,omitempty"`
	FirstAssessmentNormal bool    `json:"first_assessment_normal,omitempty"`
	RiskLevelImprovement  string  `json:"risk_level_improvement,omitempty"`
	MonthlyMoodAvg        float64 `json:"monthly_mood_avg,omitempty"`
}

// CheckAchievements 检查并解锁成就
func (s *AchievementService) CheckAchievements(userID uint, tx *gorm.DB) ([]string, error) {
	if tx == nil {
		tx = global.GVA_DB
	}

	var unlockedAchievements []string

	// 获取所有启用的成就
	var achievements []model.Achievement
	err := tx.Where("is_active = ?", true).Find(&achievements).Error
	if err != nil {
		return nil, err
	}

	// 获取用户已解锁的成就
	var userAchievements []model.UserAchievement
	err = tx.Where("user_id = ?", userID).Find(&userAchievements).Error
	if err != nil {
		return nil, err
	}

	// 创建已解锁成就的映射
	unlockedMap := make(map[uint]bool)
	for _, ua := range userAchievements {
		unlockedMap[ua.AchievementID] = true
	}

	// 检查每个成就
	for _, achievement := range achievements {
		// 跳过已解锁的成就
		if unlockedMap[achievement.ID] {
			continue
		}

		// 检查成就条件
		isUnlocked, err := s.checkAchievementCondition(userID, achievement, tx)
		if err != nil {
			global.GVA_LOG.Error("检查成就条件失败", zap.Error(err))
			continue
		}

		if isUnlocked {
			// 解锁成就
			err = s.unlockAchievement(userID, achievement.ID, tx)
			if err != nil {
				global.GVA_LOG.Error("解锁成就失败", zap.Error(err))
				continue
			}

			unlockedAchievements = append(unlockedAchievements, achievement.Name)

			// 奖励经验值
			if achievement.Rewards > 0 {
				err = s.addExperienceReward(userID, achievement.Rewards, tx)
				if err != nil {
					global.GVA_LOG.Error("添加成就奖励经验值失败", zap.Error(err))
				}
			}
		}
	}

	return unlockedAchievements, nil
}

// checkAchievementCondition 检查单个成就条件
func (s *AchievementService) checkAchievementCondition(userID uint, achievement model.Achievement, tx *gorm.DB) (bool, error) {
	var condition AchievementCondition
	err := json.Unmarshal([]byte(achievement.Condition), &condition)
	if err != nil {
		return false, err
	}

	switch achievement.Category {
	case 1: // 打卡类成就
		return s.checkCheckinAchievement(userID, condition, tx)
	case 2: // 等级类成就
		return s.checkLevelAchievement(userID, condition, tx)
	case 3: // 社区类成就
		return s.checkCommunityAchievement(userID, condition, tx)
	case 4: // 学习类成就
		return s.checkLearningAchievement(userID, condition, tx)
	case 5: // 特殊成就
		return s.checkSpecialAchievement(userID, condition, tx)
	default:
		return false, errors.New("未知的成就分类")
	}
}

// checkCheckinAchievement 检查打卡类成就
func (s *AchievementService) checkCheckinAchievement(userID uint, condition AchievementCondition, tx *gorm.DB) (bool, error) {
	// 获取戒色记录
	var record model.AbstinenceRecord
	err := tx.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	// 检查打卡次数
	if condition.CheckinCount > 0 {
		return record.TotalDays >= condition.CheckinCount, nil
	}

	// 检查连续天数
	if condition.StreakDays > 0 {
		return record.CurrentStreak >= condition.StreakDays, nil
	}

	return false, nil
}

// checkLevelAchievement 检查等级类成就
func (s *AchievementService) checkLevelAchievement(userID uint, condition AchievementCondition, tx *gorm.DB) (bool, error) {
	var record model.AbstinenceRecord
	err := tx.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	if condition.Level > 0 {
		return record.Level >= condition.Level, nil
	}

	return false, nil
}

// checkCommunityAchievement 检查社区类成就
func (s *AchievementService) checkCommunityAchievement(userID uint, condition AchievementCondition, tx *gorm.DB) (bool, error) {
	// 检查发帖数量
	if condition.PostCount > 0 {
		var count int64
		err := tx.Model(&model.CommunityPost{}).Where("user_id = ?", userID).Count(&count).Error
		if err != nil {
			return false, err
		}
		return int(count) >= condition.PostCount, nil
	}

	// 检查单篇动态点赞数
	if condition.PostLikes > 0 {
		var maxLikes int64
		err := tx.Model(&model.CommunityPost{}).
			Where("user_id = ?", userID).
			Select("MAX(likes_count)").
			Scan(&maxLikes).Error
		if err != nil {
			return false, err
		}
		return int(maxLikes) >= condition.PostLikes, nil
	}

	// 检查累计点赞数
	if condition.TotalLikes > 0 {
		var totalLikes int64
		err := tx.Model(&model.CommunityPost{}).
			Where("user_id = ?", userID).
			Select("SUM(likes_count)").
			Scan(&totalLikes).Error
		if err != nil {
			return false, err
		}
		return int(totalLikes) >= condition.TotalLikes, nil
	}

	return false, nil
}

// checkLearningAchievement 检查学习类成就
func (s *AchievementService) checkLearningAchievement(userID uint, condition AchievementCondition, tx *gorm.DB) (bool, error) {
	// 检查完成的学习内容数量
	if condition.LearningCompleted > 0 {
		var count int64
		err := tx.Model(&model.UserLearningRecord{}).
			Where("user_id = ? AND is_completed = ?", userID, true).
			Count(&count).Error
		if err != nil {
			return false, err
		}
		return int(count) >= condition.LearningCompleted, nil
	}

	// 检查学习时长（小时）
	if condition.LearningHours > 0 {
		var totalDuration int64
		err := tx.Model(&model.UserLearningRecord{}).
			Where("user_id = ?", userID).
			Select("SUM(duration)").
			Scan(&totalDuration).Error
		if err != nil {
			return false, err
		}
		totalHours := int(totalDuration / 3600) // 转换为小时
		return totalHours >= condition.LearningHours, nil
	}

	// 检查是否完成所有类型的学习内容
	if condition.AllTypesCompleted {
		// 获取所有学习内容类型
		var allTypes []int
		err := tx.Model(&model.LearningContent{}).
			Distinct("content_type").
			Pluck("content_type", &allTypes).Error
		if err != nil {
			return false, err
		}

		// 检查用户是否完成了每种类型的内容
		for _, contentType := range allTypes {
			var hasCompleted int64
			err := tx.Table("nofap_user_learning_records ulr").
				Joins("JOIN nofap_learning_contents lc ON ulr.content_id = lc.id").
				Where("ulr.user_id = ? AND lc.content_type = ? AND ulr.is_completed = ?",
					userID, contentType, true).
				Count(&hasCompleted).Error
			if err != nil {
				return false, err
			}
			if hasCompleted == 0 {
				return false, nil
			}
		}
		return true, nil
	}

	return false, nil
}

// checkSpecialAchievement 检查特殊成就
func (s *AchievementService) checkSpecialAchievement(userID uint, condition AchievementCondition, tx *gorm.DB) (bool, error) {
	// 检查帮助响应次数
	if condition.HelpResponses > 0 {
		var count int64
		err := tx.Model(&model.HelpResponse{}).
			Where("user_id = ?", userID).
			Count(&count).Error
		if err != nil {
			return false, err
		}
		return int(count) >= condition.HelpResponses, nil
	}

	// 检查首次评估是否正常
	if condition.FirstAssessmentNormal {
		var firstAssessment model.AssessmentResult
		err := tx.Where("user_id = ?", userID).
			Order("test_date ASC").
			First(&firstAssessment).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return false, nil
			}
			return false, err
		}
		return firstAssessment.RiskLevel == 1, nil // 1表示正常
	}

	// 检查风险等级改善
	if condition.RiskLevelImprovement != "" {
		// 获取第一次和最后一次评估
		var firstAssessment, lastAssessment model.AssessmentResult

		err := tx.Where("user_id = ?", userID).
			Order("test_date ASC").
			First(&firstAssessment).Error
		if err != nil {
			return false, err
		}

		err = tx.Where("user_id = ?", userID).
			Order("test_date DESC").
			First(&lastAssessment).Error
		if err != nil {
			return false, err
		}

		if condition.RiskLevelImprovement == "severe_to_normal" {
			return firstAssessment.RiskLevel >= 4 && lastAssessment.RiskLevel == 1, nil
		}
	}

	// 检查月度平均心情
	if condition.MonthlyMoodAvg > 0 {
		// 获取本月的平均心情
		now := time.Now()
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endOfMonth := startOfMonth.AddDate(0, 1, -1)

		var avgMood sql.NullFloat64
		err := tx.Model(&model.DailyCheckin{}).
			Where("user_id = ? AND checkin_date BETWEEN ? AND ?", userID, startOfMonth, endOfMonth).
			Select("AVG(mood_level)").
			Scan(&avgMood).Error
		if err != nil {
			return false, err
		}

		// 处理NULL值
		avgMoodValue := 0.0
		if avgMood.Valid {
			avgMoodValue = avgMood.Float64
		}

		return avgMoodValue >= condition.MonthlyMoodAvg, nil
	}

	return false, nil
}

// unlockAchievement 解锁成就
func (s *AchievementService) unlockAchievement(userID uint, achievementID uint, tx *gorm.DB) error {
	userAchievement := model.UserAchievement{
		UserID:        userID,
		AchievementID: achievementID,
		UnlockedAt:    time.Now(),
		Progress:      100,
		IsNotified:    false,
	}

	return tx.Create(&userAchievement).Error
}

// addExperienceReward 添加经验值奖励
func (s *AchievementService) addExperienceReward(userID uint, rewards int, tx *gorm.DB) error {
	var record model.AbstinenceRecord
	err := tx.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil {
		return err
	}

	record.Experience += rewards
	return tx.Save(&record).Error
}

// GetUserAchievements 获取用户成就列表
func (s *AchievementService) GetUserAchievements(userID uint) (*response.UserAchievementsResponse, error) {
	// 获取所有启用的成就
	var allAchievements []model.Achievement
	err := global.GVA_DB.Where("is_active = ?", true).
		Order("category ASC, display_order ASC").
		Find(&allAchievements).Error
	if err != nil {
		return nil, err
	}

	// 获取用户已解锁的成就
	var userAchievements []model.UserAchievement
	err = global.GVA_DB.Preload("Achievement").
		Where("user_id = ?", userID).
		Find(&userAchievements).Error
	if err != nil {
		return nil, err
	}

	// 创建已解锁成就映射
	unlockedMap := make(map[uint]model.UserAchievement)
	for _, ua := range userAchievements {
		unlockedMap[ua.AchievementID] = ua
	}

	// 构建响应数据
	var achievementList []response.AchievementItem
	unlockedCount := 0

	for _, achievement := range allAchievements {
		item := response.AchievementItem{
			ID:          achievement.ID,
			Name:        achievement.Name,
			Description: achievement.Description,
			IconUrl:     achievement.IconUrl,
			Category:    achievement.Category,
			Rarity:      achievement.Rarity,
			Rewards:     achievement.Rewards,
			IsUnlocked:  false,
		}

		if userAchievement, exists := unlockedMap[achievement.ID]; exists {
			item.IsUnlocked = true
			item.UnlockedAt = &userAchievement.UnlockedAt
			item.Progress = userAchievement.Progress
			unlockedCount++
		}

		achievementList = append(achievementList, item)
	}

	// 按分类分组
	categories := map[int][]response.AchievementItem{
		1: {}, // 打卡类
		2: {}, // 等级类
		3: {}, // 社区类
		4: {}, // 学习类
		5: {}, // 特殊类
	}

	for _, item := range achievementList {
		categories[item.Category] = append(categories[item.Category], item)
	}

	return &response.UserAchievementsResponse{
		TotalAchievements:    len(allAchievements),
		UnlockedAchievements: unlockedCount,
		UnlockRate:           float64(unlockedCount) / float64(len(allAchievements)) * 100,
		Categories:           categories,
		RecentUnlocked:       s.getRecentUnlocked(userAchievements, 5),
	}, nil
}

// getRecentUnlocked 获取最近解锁的成就
func (s *AchievementService) getRecentUnlocked(userAchievements []model.UserAchievement, limit int) []response.AchievementItem {
	if len(userAchievements) == 0 {
		return []response.AchievementItem{}
	}

	// 按解锁时间倒序排序
	achievements := make([]model.UserAchievement, len(userAchievements))
	copy(achievements, userAchievements)

	// 简单排序：按解锁时间倒序
	for i := 0; i < len(achievements)-1; i++ {
		for j := i + 1; j < len(achievements); j++ {
			if achievements[i].UnlockedAt.Before(achievements[j].UnlockedAt) {
				achievements[i], achievements[j] = achievements[j], achievements[i]
			}
		}
	}

	// 取前limit个
	if len(achievements) > limit {
		achievements = achievements[:limit]
	}

	var recent []response.AchievementItem
	for _, ua := range achievements {
		item := response.AchievementItem{
			ID:          ua.Achievement.ID,
			Name:        ua.Achievement.Name,
			Description: ua.Achievement.Description,
			IconUrl:     ua.Achievement.IconUrl,
			Category:    ua.Achievement.Category,
			Rarity:      ua.Achievement.Rarity,
			Rewards:     ua.Achievement.Rewards,
			IsUnlocked:  true,
			UnlockedAt:  &ua.UnlockedAt,
			Progress:    ua.Progress,
		}
		recent = append(recent, item)
	}

	return recent
}

// GetAchievementStats 获取成就统计
func (s *AchievementService) GetAchievementStats(userID uint) (*response.AchievementStatsResponse, error) {
	// 获取总成就数
	var totalAchievements int64
	err := global.GVA_DB.Model(&model.Achievement{}).
		Where("is_active = ?", true).
		Count(&totalAchievements).Error
	if err != nil {
		return nil, err
	}

	// 获取用户解锁的成就数
	var unlockedAchievements int64
	err = global.GVA_DB.Model(&model.UserAchievement{}).
		Where("user_id = ?", userID).
		Count(&unlockedAchievements).Error
	if err != nil {
		return nil, err
	}

	// 获取各稀有度成就统计
	rarityStats := make(map[int]response.RarityStats)
	for rarity := 1; rarity <= 4; rarity++ {
		var total, unlocked int64

		// 获取该稀有度总数
		err = global.GVA_DB.Model(&model.Achievement{}).
			Where("is_active = ? AND rarity = ?", true, rarity).
			Count(&total).Error
		if err != nil {
			return nil, err
		}

		// 获取用户解锁数
		err = global.GVA_DB.Table("nofap_user_achievements ua").
			Joins("JOIN nofap_achievements a ON ua.achievement_id = a.id").
			Where("ua.user_id = ? AND a.rarity = ? AND a.is_active = ?", userID, rarity, true).
			Count(&unlocked).Error
		if err != nil {
			return nil, err
		}

		rarityStats[rarity] = response.RarityStats{
			Total:    int(total),
			Unlocked: int(unlocked),
		}
	}

	// 计算完成率
	unlockRate := float64(0)
	if totalAchievements > 0 {
		unlockRate = float64(unlockedAchievements) / float64(totalAchievements) * 100
	}

	return &response.AchievementStatsResponse{
		TotalAchievements:    int(totalAchievements),
		UnlockedAchievements: int(unlockedAchievements),
		UnlockRate:           unlockRate,
		RarityStats:          rarityStats,
	}, nil
}

// GetAchievementProgress 获取成就进度信息
func (s *AchievementService) GetAchievementProgress(userID uint) ([]response.AchievementProgressResponse, error) {
	// 获取所有未完成的成就
	var achievements []model.Achievement
	err := global.GVA_DB.Where("is_active = ?", true).Find(&achievements).Error
	if err != nil {
		return nil, err
	}

	// 获取用户已解锁的成就
	var userAchievements []model.UserAchievement
	err = global.GVA_DB.Where("user_id = ?", userID).Find(&userAchievements).Error
	if err != nil {
		return nil, err
	}

	// 创建已解锁成就映射
	unlockedMap := make(map[uint]bool)
	for _, ua := range userAchievements {
		unlockedMap[ua.AchievementID] = true
	}

	var progressList []response.AchievementProgressResponse

	// 只返回正在进行中的成就（未解锁且有进度的）
	for _, achievement := range achievements {
		// 跳过已解锁的成就
		if unlockedMap[achievement.ID] {
			continue
		}

		// 计算当前进度
		progress, targetValue, err := s.calculateAchievementProgress(userID, achievement)
		if err != nil {
			continue
		}

		// 只返回有进度的成就
		if progress > 0 {
			categoryName := s.getCategoryName(achievement.Category)
			rarityName := s.getRarityName(achievement.Rarity)

			progressList = append(progressList, response.AchievementProgressResponse{
				ID:          achievement.ID,
				Name:        achievement.Name,
				Description: achievement.Description,
				Category:    categoryName,
				Rarity:      rarityName,
				Progress:    progress,
				TargetValue: targetValue,
				IconURL:     achievement.IconUrl,
				IsCompleted: false,
			})
		}
	}

	return progressList, nil
}

// calculateAchievementProgress 计算成就进度
func (s *AchievementService) calculateAchievementProgress(userID uint, achievement model.Achievement) (int, int, error) {
	var condition AchievementCondition
	err := json.Unmarshal([]byte(achievement.Condition), &condition)
	if err != nil {
		return 0, 0, err
	}

	switch achievement.Category {
	case 1: // 打卡类成就
		return s.getCheckinProgress(userID, condition)
	case 2: // 等级类成就
		return s.getLevelProgress(userID, condition)
	case 3: // 社区类成就
		return s.getCommunityProgress(userID, condition)
	case 4: // 学习类成就
		return s.getLearningProgress(userID, condition)
	case 5: // 特殊成就
		return s.getSpecialProgress(userID, condition)
	default:
		return 0, 0, errors.New("未知的成就分类")
	}
}

// getCheckinProgress 获取打卡类成就进度
func (s *AchievementService) getCheckinProgress(userID uint, condition AchievementCondition) (int, int, error) {
	// 获取戒色记录
	var record model.AbstinenceRecord
	err := global.GVA_DB.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, 0, nil
		}
		return 0, 0, err
	}

	// 检查打卡次数进度
	if condition.CheckinCount > 0 {
		return record.TotalDays, condition.CheckinCount, nil
	}

	// 检查连续天数进度
	if condition.StreakDays > 0 {
		return record.CurrentStreak, condition.StreakDays, nil
	}

	return 0, 0, nil
}

// getLevelProgress 获取等级类成就进度
func (s *AchievementService) getLevelProgress(userID uint, condition AchievementCondition) (int, int, error) {
	var record model.AbstinenceRecord
	err := global.GVA_DB.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, 0, nil
		}
		return 0, 0, err
	}

	if condition.Level > 0 {
		return record.Level, condition.Level, nil
	}

	return 0, 0, nil
}

// getCommunityProgress 获取社区类成就进度
func (s *AchievementService) getCommunityProgress(userID uint, condition AchievementCondition) (int, int, error) {
	// 检查发帖数量进度
	if condition.PostCount > 0 {
		var count int64
		err := global.GVA_DB.Model(&model.CommunityPost{}).Where("user_id = ?", userID).Count(&count).Error
		if err != nil {
			return 0, 0, err
		}
		return int(count), condition.PostCount, nil
	}

	// 检查累计点赞数进度
	if condition.TotalLikes > 0 {
		var totalLikes int64
		err := global.GVA_DB.Model(&model.CommunityPost{}).
			Where("user_id = ?", userID).
			Select("SUM(likes_count)").
			Scan(&totalLikes).Error
		if err != nil {
			return 0, 0, err
		}
		return int(totalLikes), condition.TotalLikes, nil
	}

	return 0, 0, nil
}

// getLearningProgress 获取学习类成就进度
func (s *AchievementService) getLearningProgress(userID uint, condition AchievementCondition) (int, int, error) {
	// 检查完成的学习内容数量进度
	if condition.LearningCompleted > 0 {
		var count int64
		err := global.GVA_DB.Model(&model.UserLearningRecord{}).
			Where("user_id = ? AND is_completed = ?", userID, true).
			Count(&count).Error
		if err != nil {
			return 0, 0, err
		}
		return int(count), condition.LearningCompleted, nil
	}

	// 检查学习时长进度
	if condition.LearningHours > 0 {
		var totalDuration int64
		err := global.GVA_DB.Model(&model.UserLearningRecord{}).
			Where("user_id = ?", userID).
			Select("SUM(duration)").
			Scan(&totalDuration).Error
		if err != nil {
			return 0, 0, err
		}
		totalHours := int(totalDuration / 3600) // 转换为小时
		return totalHours, condition.LearningHours, nil
	}

	return 0, 0, nil
}

// getSpecialProgress 获取特殊成就进度
func (s *AchievementService) getSpecialProgress(userID uint, condition AchievementCondition) (int, int, error) {
	// 检查帮助响应次数进度
	if condition.HelpResponses > 0 {
		var count int64
		err := global.GVA_DB.Model(&model.HelpResponse{}).
			Where("user_id = ?", userID).
			Count(&count).Error
		if err != nil {
			return 0, 0, err
		}
		return int(count), condition.HelpResponses, nil
	}

	return 0, 0, nil
}

// getCategoryName 获取分类名称
func (s *AchievementService) getCategoryName(category int) string {
	switch category {
	case 1:
		return "checkin"
	case 2:
		return "level"
	case 3:
		return "community"
	case 4:
		return "learning"
	case 5:
		return "special"
	default:
		return "unknown"
	}
}

// getRarityName 获取稀有度名称
func (s *AchievementService) getRarityName(rarity int) string {
	switch rarity {
	case 1:
		return "common"
	case 2:
		return "rare"
	case 3:
		return "epic"
	case 4:
		return "legendary"
	default:
		return "common"
	}
}

// GetUserAchievementsForProfile 获取用户成就（用于个人中心）
func (s *AchievementService) GetUserAchievementsForProfile(userID uint, limit int, recent bool) (*response.UserProfileAchievementsResponse, error) {
	var achievements []response.AchievementItem

	if recent {
		// 获取最近解锁的成就
		var userAchievements []model.UserAchievement
		err := global.GVA_DB.Preload("Achievement").
			Where("user_id = ?", userID).
			Order("unlocked_at DESC").
			Limit(limit).
			Find(&userAchievements).Error
		if err != nil {
			return nil, err
		}

		for _, ua := range userAchievements {
			item := response.AchievementItem{
				ID:          ua.Achievement.ID,
				Name:        ua.Achievement.Name,
				Description: ua.Achievement.Description,
				IconUrl:     ua.Achievement.IconUrl,
				Category:    ua.Achievement.Category,
				Rarity:      ua.Achievement.Rarity,
				Rewards:     ua.Achievement.Rewards,
				IsUnlocked:  true,
				UnlockedAt:  &ua.UnlockedAt,
				Progress:    ua.Progress,
			}
			achievements = append(achievements, item)
		}
	} else {
		// 获取所有成就，包括未解锁的
		var allAchievements []model.Achievement
		err := global.GVA_DB.Where("is_active = ?", true).
			Order("category ASC, display_order ASC").
			Limit(limit).
			Find(&allAchievements).Error
		if err != nil {
			return nil, err
		}

		// 获取用户已解锁的成就
		var userAchievements []model.UserAchievement
		err = global.GVA_DB.Where("user_id = ?", userID).Find(&userAchievements).Error
		if err != nil {
			return nil, err
		}

		// 创建已解锁成就映射
		unlockedMap := make(map[uint]model.UserAchievement)
		for _, ua := range userAchievements {
			unlockedMap[ua.AchievementID] = ua
		}

		for _, achievement := range allAchievements {
			item := response.AchievementItem{
				ID:          achievement.ID,
				Name:        achievement.Name,
				Description: achievement.Description,
				IconUrl:     achievement.IconUrl,
				Category:    achievement.Category,
				Rarity:      achievement.Rarity,
				Rewards:     achievement.Rewards,
				IsUnlocked:  false,
			}

			if userAchievement, exists := unlockedMap[achievement.ID]; exists {
				item.IsUnlocked = true
				item.UnlockedAt = &userAchievement.UnlockedAt
				item.Progress = userAchievement.Progress
			}

			achievements = append(achievements, item)
		}
	}

	return &response.UserProfileAchievementsResponse{
		List: achievements,
	}, nil
}
