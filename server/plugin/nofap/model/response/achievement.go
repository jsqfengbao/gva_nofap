package response

import (
	"time"
)

// AchievementItem 成就项目
type AchievementItem struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IconUrl     string     `json:"iconUrl"`
	Category    int        `json:"category"`
	Rarity      int        `json:"rarity"`
	Rewards     int        `json:"rewards"`
	IsUnlocked  bool       `json:"isUnlocked"`
	UnlockedAt  *time.Time `json:"unlockedAt"`
	Progress    int        `json:"progress"`
}

// UserAchievementsResponse 用户成就响应
type UserAchievementsResponse struct {
	TotalAchievements    int                       `json:"totalAchievements"`
	UnlockedAchievements int                       `json:"unlockedAchievements"`
	UnlockRate           float64                   `json:"unlockRate"`
	Categories           map[int][]AchievementItem `json:"categories"`
	RecentUnlocked       []AchievementItem         `json:"recentUnlocked"`
}

// RarityStats 稀有度统计
type RarityStats struct {
	Total    int `json:"total"`
	Unlocked int `json:"unlocked"`
}

// AchievementStatsResponse 成就统计响应
type AchievementStatsResponse struct {
	TotalAchievements    int                 `json:"totalAchievements"`
	UnlockedAchievements int                 `json:"unlockedAchievements"`
	UnlockRate           float64             `json:"unlockRate"`
	RarityStats          map[int]RarityStats `json:"rarityStats"`
}

// GameStatsResponse 游戏化统计响应
type GameStatsResponse struct {
	// 等级信息
	CurrentLevel  int     `json:"currentLevel"`
	CurrentExp    int     `json:"currentExp"`
	NextLevelExp  int     `json:"nextLevelExp"`
	LevelProgress float64 `json:"levelProgress"`

	// 成就信息
	TotalAchievements    int     `json:"totalAchievements"`
	UnlockedAchievements int     `json:"unlockedAchievements"`
	AchievementRate      float64 `json:"achievementRate"`

	// 打卡信息
	CurrentStreak int     `json:"currentStreak"`
	LongestStreak int     `json:"longestStreak"`
	TotalDays     int     `json:"totalDays"`
	SuccessRate   float64 `json:"successRate"`

	// 最近成就
	RecentAchievements []AchievementItem `json:"recentAchievements"`
}

// LevelProgressResponse 等级进度响应
type LevelProgressResponse struct {
	CurrentLevel   int     `json:"currentLevel"`
	CurrentExp     int     `json:"currentExp"`
	NextLevelExp   int     `json:"nextLevelExp"`
	LevelProgress  float64 `json:"levelProgress"`
	TotalExpNeeded int     `json:"totalExpNeeded"`
	ExpToNextLevel int     `json:"expToNextLevel"`
	IsMaxLevel     bool    `json:"isMaxLevel"`
}

// RewardResponse 奖励响应
type RewardResponse struct {
	Type        string `json:"type"`        // exp, achievement, level_up
	Amount      int    `json:"amount"`      // 奖励数量
	Description string `json:"description"` // 奖励描述
	IconUrl     string `json:"iconUrl"`     // 图标URL
}

// AchievementProgressResponse 成就进度响应
type AchievementProgressResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Rarity      string `json:"rarity"`
	Progress    int    `json:"progress"`
	TargetValue int    `json:"targetValue"`
	IconURL     string `json:"iconUrl"`
	IsCompleted bool   `json:"isCompleted"`
}

// UserProfileAchievementsResponse 用户个人中心成就响应
type UserProfileAchievementsResponse struct {
	List []AchievementItem `json:"list"`
}
