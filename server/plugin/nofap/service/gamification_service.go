package service

import (
	"math"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model/response"
	"gorm.io/gorm"
)

type GamificationService struct{}

const (
	MaxLevel    = 50  // 最高等级
	ExpPerLevel = 100 // 每级所需经验值
)

// GetGameStats 获取游戏化统计数据
func (s *GamificationService) GetGameStats(userID uint) (*response.GameStatsResponse, error) {
	// 获取戒色记录
	var record model.AbstinenceRecord
	err := global.GVA_DB.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 创建默认记录
			record = model.AbstinenceRecord{
				Level:      1,
				Experience: 0,
			}
		} else {
			return nil, err
		}
	}

	// 计算等级进度
	levelProgress := s.calculateLevelProgress(record.Experience, record.Level)

	// 获取成就统计
	achievementService := AchievementService{}
	achievementStats, err := achievementService.GetAchievementStats(userID)
	if err != nil {
		return nil, err
	}

	// 获取最近解锁的成就
	userAchievements, err := achievementService.GetUserAchievements(userID)
	if err != nil {
		return nil, err
	}

	return &response.GameStatsResponse{
		// 等级信息
		CurrentLevel:  record.Level,
		CurrentExp:    record.Experience,
		NextLevelExp:  s.getExpForLevel(record.Level + 1),
		LevelProgress: levelProgress,

		// 成就信息
		TotalAchievements:    achievementStats.TotalAchievements,
		UnlockedAchievements: achievementStats.UnlockedAchievements,
		AchievementRate:      achievementStats.UnlockRate,

		// 打卡信息
		CurrentStreak: record.CurrentStreak,
		LongestStreak: record.LongestStreak,
		TotalDays:     record.TotalDays,
		SuccessRate:   record.SuccessRate,

		// 最近成就
		RecentAchievements: userAchievements.RecentUnlocked,
	}, nil
}

// GetLevelProgress 获取等级进度信息
func (s *GamificationService) GetLevelProgress(userID uint) (*response.LevelProgressResponse, error) {
	// 获取戒色记录
	var record model.AbstinenceRecord
	err := global.GVA_DB.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 创建默认记录
			record = model.AbstinenceRecord{
				Level:      1,
				Experience: 0,
			}
		} else {
			return nil, err
		}
	}

	currentLevelExp := s.getExpForLevel(record.Level)
	nextLevelExp := s.getExpForLevel(record.Level + 1)
	levelProgress := s.calculateLevelProgress(record.Experience, record.Level)

	return &response.LevelProgressResponse{
		CurrentLevel:   record.Level,
		CurrentExp:     record.Experience,
		NextLevelExp:   nextLevelExp,
		LevelProgress:  levelProgress,
		TotalExpNeeded: nextLevelExp - currentLevelExp,
		ExpToNextLevel: nextLevelExp - record.Experience,
		IsMaxLevel:     record.Level >= MaxLevel,
	}, nil
}

// AddExperience 添加经验值并检查升级
func (s *GamificationService) AddExperience(userID uint, exp int, tx *gorm.DB) (*response.RewardResponse, error) {
	if tx == nil {
		tx = global.GVA_DB
	}

	// 获取戒色记录
	var record model.AbstinenceRecord
	err := tx.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil {
		return nil, err
	}

	oldLevel := record.Level
	record.Experience += exp

	// 计算新等级
	newLevel := s.calculateLevel(record.Experience)
	levelUp := newLevel > oldLevel && newLevel <= MaxLevel

	if levelUp {
		record.Level = newLevel
	}

	// 保存记录
	err = tx.Save(&record).Error
	if err != nil {
		return nil, err
	}

	// 构建奖励响应
	reward := &response.RewardResponse{
		Type:        "exp",
		Amount:      exp,
		Description: "获得经验值",
	}

	if levelUp {
		reward.Type = "level_up"
		reward.Amount = newLevel
		reward.Description = "恭喜升级！"
	}

	return reward, nil
}

// CalculateCheckinRewards 计算打卡奖励
func (s *GamificationService) CalculateCheckinRewards(streak int, moodLevel int) int {
	baseRewards := 10
	streakBonus := s.calculateStreakBonus(streak)
	moodBonus := s.calculateMoodBonus(moodLevel)

	return baseRewards + streakBonus + moodBonus
}

// calculateLevel 根据经验值计算等级
func (s *GamificationService) calculateLevel(experience int) int {
	level := experience/ExpPerLevel + 1
	if level > MaxLevel {
		level = MaxLevel
	}
	return level
}

// getExpForLevel 获取指定等级所需的经验值
func (s *GamificationService) getExpForLevel(level int) int {
	if level <= 1 {
		return 0
	}
	if level > MaxLevel {
		level = MaxLevel
	}
	return (level - 1) * ExpPerLevel
}

// calculateLevelProgress 计算当前等级的进度百分比
func (s *GamificationService) calculateLevelProgress(currentExp int, currentLevel int) float64 {
	if currentLevel >= MaxLevel {
		return 100.0
	}

	currentLevelExp := s.getExpForLevel(currentLevel)
	nextLevelExp := s.getExpForLevel(currentLevel + 1)

	if nextLevelExp <= currentLevelExp {
		return 100.0
	}

	expInCurrentLevel := currentExp - currentLevelExp
	expNeededForNextLevel := nextLevelExp - currentLevelExp

	progress := float64(expInCurrentLevel) / float64(expNeededForNextLevel) * 100
	return math.Round(progress*10) / 10 // 保留一位小数
}

// calculateStreakBonus 计算连续天数奖励
func (s *GamificationService) calculateStreakBonus(streak int) int {
	switch {
	case streak >= 365:
		return 100
	case streak >= 180:
		return 50
	case streak >= 90:
		return 30
	case streak >= 30:
		return 20
	case streak >= 7:
		return 10
	case streak >= 3:
		return 5
	default:
		return 0
	}
}

// calculateMoodBonus 计算心情奖励
func (s *GamificationService) calculateMoodBonus(moodLevel int) int {
	switch moodLevel {
	case 5:
		return 5
	case 4:
		return 3
	case 3:
		return 1
	case 2:
		return 0
	case 1:
		return 0
	default:
		return 0
	}
}

// GetLevelTitle 根据等级获取称号
func (s *GamificationService) GetLevelTitle(level int) string {
	switch {
	case level >= 50:
		return "传奇大师"
	case level >= 40:
		return "宗师"
	case level >= 30:
		return "专家"
	case level >= 25:
		return "资深导师"
	case level >= 20:
		return "高级导师"
	case level >= 15:
		return "导师"
	case level >= 10:
		return "进阶行者"
	case level >= 5:
		return "初级学徒"
	case level >= 1:
		return "新手"
	default:
		return "未知"
	}
}

// GetRarityName 获取稀有度名称
func (s *GamificationService) GetRarityName(rarity int) string {
	switch rarity {
	case 1:
		return "普通"
	case 2:
		return "稀有"
	case 3:
		return "史诗"
	case 4:
		return "传说"
	default:
		return "未知"
	}
}

// GetCategoryName 获取成就分类名称
func (s *GamificationService) GetCategoryName(category int) string {
	switch category {
	case 1:
		return "打卡类"
	case 2:
		return "等级类"
	case 3:
		return "社区类"
	case 4:
		return "学习类"
	case 5:
		return "特殊类"
	default:
		return "未知"
	}
}
