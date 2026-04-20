package model

// AdminStatistics 管理端统计数据
type AdminStatistics struct {
	TotalUsers      int64 `json:"totalUsers"`      // 总用户数
	ActiveUsers     int64 `json:"activeUsers"`     // 活跃用户（7天内活跃）
	TotalPosts      int64 `json:"totalPosts"`      // 总帖子数
	TotalCheckins   int64 `json:"totalCheckins"`   // 总打卡天数
}
