package response

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model"
)

// CheckinResponse 打卡响应
type CheckinResponse struct {
	ID          uint      `json:"id"`
	CheckinDate time.Time `json:"checkinDate"`
	MoodLevel   int       `json:"moodLevel"`
	Notes       string    `json:"notes"`
	Rewards     int       `json:"rewards"`
	Streak      int       `json:"streak"`
	IsSuccess   bool      `json:"isSuccess"`
	LevelUp     bool      `json:"levelUp"`     // 是否升级
	NewLevel    int       `json:"newLevel"`    // 新等级
	Experience  int       `json:"experience"`  // 当前经验值
	Achievement *string   `json:"achievement"` // 解锁的成就
}

// TodayCheckinResponse 今日打卡状态响应
type TodayCheckinResponse struct {
	HasChecked  bool       `json:"hasChecked"`
	CheckinDate *time.Time `json:"checkinDate"`
	MoodLevel   int        `json:"moodLevel"`
	Notes       string     `json:"notes"`
	Rewards     int        `json:"rewards"`
	CheckinTime *time.Time `json:"checkinTime"`
}

// CheckinHistoryResponse 打卡历史响应
type CheckinHistoryResponse struct {
	List     []miniprogram.DailyCheckin `json:"list"`
	Total    int64                      `json:"total"`
	Page     int                        `json:"page"`
	PageSize int                        `json:"pageSize"`
}

// CheckinStatsResponse 打卡统计响应
type CheckinStatsResponse struct {
	TotalDays     int     `json:"totalDays"`     // 累计打卡天数
	CurrentStreak int     `json:"currentStreak"` // 当前连续天数
	LongestStreak int     `json:"longestStreak"` // 最长连续天数
	SuccessRate   float64 `json:"successRate"`   // 成功率
	AverageMood   float64 `json:"averageMood"`   // 平均心情
	TotalRewards  int     `json:"totalRewards"`  // 累计奖励
	ThisMonth     int     `json:"thisMonth"`     // 本月打卡天数
	ThisWeek      int     `json:"thisWeek"`      // 本周打卡天数
	Level         int     `json:"level"`         // 当前等级
	Experience    int     `json:"experience"`    // 当前经验值
}

// WeeklyProgressResponse 本周进度响应
type WeeklyProgressResponse struct {
	WeekDays []WeekDay   `json:"weekDays"`
	Summary  WeekSummary `json:"summary"`
}

type WeekDay struct {
	Date       string `json:"date"`       // 日期 YYYY-MM-DD
	Weekday    string `json:"weekday"`    // 星期几
	HasChecked bool   `json:"hasChecked"` // 是否已打卡
	MoodLevel  int    `json:"moodLevel"`  // 心情等级
	IsToday    bool   `json:"isToday"`    // 是否今天
}

type WeekSummary struct {
	CheckedDays int     `json:"checkedDays"` // 本周已打卡天数
	TotalDays   int     `json:"totalDays"`   // 本周总天数
	SuccessRate float64 `json:"successRate"` // 本周成功率
	AverageMood float64 `json:"averageMood"` // 本周平均心情
}

// MonthlyCalendarResponse 月度日历响应
type MonthlyCalendarResponse struct {
	Year    int           `json:"year"`
	Month   int           `json:"month"`
	Days    []CalendarDay `json:"days"`
	Summary MonthSummary  `json:"summary"`
}

type CalendarDay struct {
	Day         int    `json:"day"`         // 日期
	Date        string `json:"date"`        // 完整日期 YYYY-MM-DD
	HasChecked  bool   `json:"hasChecked"`  // 是否已打卡
	MoodLevel   int    `json:"moodLevel"`   // 心情等级
	IsToday     bool   `json:"isToday"`     // 是否今天
	IsThisMonth bool   `json:"isThisMonth"` // 是否本月
}

type MonthSummary struct {
	CheckedDays int     `json:"checkedDays"` // 本月已打卡天数
	TotalDays   int     `json:"totalDays"`   // 本月总天数
	SuccessRate float64 `json:"successRate"` // 本月成功率
	AverageMood float64 `json:"averageMood"` // 本月平均心情
	BestStreak  int     `json:"bestStreak"`  // 本月最长连击
}

// WeeklyChartResponse 本周进度图表响应
type WeeklyChartResponse struct {
	WeeklyCheckins []WeeklyCheckinData `json:"weeklyCheckins"`
	CompletionRate float64             `json:"completionRate"`
}

type WeeklyCheckinData struct {
	Date       string `json:"date"`
	Weekday    string `json:"weekday"`
	HasChecked bool   `json:"hasChecked"`
	MoodLevel  int    `json:"moodLevel"`
}
