package response

// OverallStatsResponse 总体统计响应
type OverallStatsResponse struct {
	Abstinence AbstinenceStats `json:"abstinence"`
	Checkin     CheckinStats   `json:"checkin"`
	Achievement AchievementStats `json:"achievement"`
	Learning    LearningStats  `json:"learning"`
	Community   CommunityStats `json:"community"`
}

// AbstinenceStats 戒色统计
type AbstinenceStats struct {
	CurrentStreak int     `json:"current_streak"`  // 当前连续天数
	LongestStreak int     `json:"longest_streak"`  // 最长连续天数
	TotalDays     int     `json:"total_days"`      // 总成功天数
	SuccessRate   float64 `json:"success_rate"`    // 成功率
	StartDate     string  `json:"start_date"`      // 开始日期
	Level         int     `json:"level"`           // 当前等级
	Experience    int     `json:"experience"`      // 当前经验值
}

// CheckinStats 打卡统计
type CheckinStats struct {
	TotalCheckins  int `json:"total_checkins"`  // 总打卡次数
	MonthCheckins int `json:"month_checkins"`  // 近30天打卡次数
}

// AchievementStats 成就统计
type AchievementStats struct {
	TotalAchievements    int `json:"total_achievements"`    // 总成就数
	UnlockedAchievements int `json:"unlocked_achievements"` // 已解锁成就数
}

// LearningStats 学习统计
type LearningStats struct {
	TotalArticles int `json:"total_articles"` // 总文章数
	ReadArticles  int `json:"read_articles"`  // 已读文章数
	TotalReadTime int `json:"total_read_time"` // 总阅读时长（秒）
}

// CommunityStats 社区统计
type CommunityStats struct {
	TotalPosts int `json:"total_posts"` // 用户发帖数
	TotalLikes int `json:"total_likes"` // 获得点赞总数
}

// TrendsResponse 趋势数据响应
type TrendsResponse struct {
	Trends []TrendPoint `json:"trends"`
}

// TrendPoint 趋势数据点
type TrendPoint struct {
	Date    string `json:"date"`    // 日期
	Success bool   `json:"success"` // 是否成功打卡
	Streak  int    `json:"streak"`  // 当前连续天数
}
