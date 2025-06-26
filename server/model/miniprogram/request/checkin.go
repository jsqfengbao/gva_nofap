package request

// CheckinRequest 每日打卡请求
type CheckinRequest struct {
	MoodLevel int    `json:"moodLevel" binding:"required,min=1,max=5"` // 心情等级 1-5
	Notes     string `json:"notes"`                                    // 打卡备注
}
