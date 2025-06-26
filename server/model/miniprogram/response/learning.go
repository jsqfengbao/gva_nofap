package response

import (
	"time"
)

// LearningContentResponse 学习内容响应
type LearningContentResponse struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	Title        string     `json:"title"`
	Summary      string     `json:"summary"`
	Content      string     `json:"content"`
	ContentType  int        `json:"contentType"`
	Category     int        `json:"category"`
	Difficulty   int        `json:"difficulty"`
	Duration     int        `json:"duration"`
	ThumbnailUrl string     `json:"thumbnailUrl"`
	MediaUrl     string     `json:"mediaUrl"`
	Author       string     `json:"author"`
	ViewCount    int        `json:"viewCount"`
	LikeCount    int        `json:"likeCount"`
	CollectCount int        `json:"collectCount"`
	CommentCount int        `json:"commentCount"`
	Status       int        `json:"status"`
	PublishAt    *time.Time `json:"publishAt"`
	Tags         string     `json:"tags"`
	TagList      []string   `json:"tagList"` // 处理后的标签数组

	// 用户相关状态
	IsLiked     bool `json:"isLiked"`     // 当前用户是否点赞
	IsCollected bool `json:"isCollected"` // 当前用户是否收藏
	UserRating  int  `json:"userRating"`  // 当前用户评分

	// 学习进度
	LearningProgress *UserLearningProgressResponse `json:"learningProgress,omitempty"`
}

// UserLearningProgressResponse 用户学习进度响应
type UserLearningProgressResponse struct {
	ID          uint       `json:"id"`
	StartTime   time.Time  `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	Duration    int        `json:"duration"`
	Progress    int        `json:"progress"`
	IsCompleted bool       `json:"isCompleted"`
	Rating      int        `json:"rating"`
	Notes       string     `json:"notes"`
}

// LearningContentListResponse 学习内容列表响应
type LearningContentListResponse struct {
	List     []LearningContentResponse `json:"list"`
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"pageSize"`
}

// UserLearningRecordResponse 用户学习记录响应
type UserLearningRecordResponse struct {
	ID          uint                    `json:"id"`
	StartTime   time.Time               `json:"startTime"`
	EndTime     *time.Time              `json:"endTime"`
	Duration    int                     `json:"duration"`
	Progress    int                     `json:"progress"`
	IsCompleted bool                    `json:"isCompleted"`
	IsLiked     bool                    `json:"isLiked"`
	IsCollected bool                    `json:"isCollected"`
	Rating      int                     `json:"rating"`
	Notes       string                  `json:"notes"`
	Content     LearningContentResponse `json:"content"`
}

// UserLearningRecordListResponse 用户学习记录列表响应
type UserLearningRecordListResponse struct {
	List     []UserLearningRecordResponse `json:"list"`
	Total    int64                        `json:"total"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"pageSize"`
}

// LearningStatsResponse 学习统计响应
type LearningStatsResponse struct {
	TotalContents      int64      `json:"totalContents"`      // 总内容数
	CompletedContents  int64      `json:"completedContents"`  // 已完成内容数
	LikedContents      int64      `json:"likedContents"`      // 点赞内容数
	CollectedContents  int64      `json:"collectedContents"`  // 收藏内容数
	TotalLearningTime  int        `json:"totalLearningTime"`  // 总学习时长(分钟)
	AvgLearningTime    int        `json:"avgLearningTime"`    // 平均学习时长(分钟)
	CompletionRate     int        `json:"completionRate"`     // 完成率(百分比)
	ContinuousLearning int        `json:"continuousLearning"` // 连续学习天数
	LastLearningTime   *time.Time `json:"lastLearningTime"`   // 最后学习时间
}

// CategoryStatsResponse 分类统计响应
type CategoryStatsResponse struct {
	Category          int    `json:"category"`          // 分类
	CategoryName      string `json:"categoryName"`      // 分类名称
	TotalContents     int64  `json:"totalContents"`     // 总内容数
	CompletedContents int64  `json:"completedContents"` // 已完成内容数
	CompletionRate    int    `json:"completionRate"`    // 完成率
}

// LearningRecommendationResponse 学习推荐响应
type LearningRecommendationResponse struct {
	RecommendType string                    `json:"recommendType"` // 推荐类型: continue, similar, trending, new
	Reason        string                    `json:"reason"`        // 推荐理由
	Contents      []LearningContentResponse `json:"contents"`      // 推荐内容列表
}

// LearningHomepageResponse 学习首页响应
type LearningHomepageResponse struct {
	Stats           LearningStatsResponse            `json:"stats"`           // 学习统计
	CategoryStats   []CategoryStatsResponse          `json:"categoryStats"`   // 分类统计
	Recommendations []LearningRecommendationResponse `json:"recommendations"` // 推荐内容
	RecentContents  []LearningContentResponse        `json:"recentContents"`  // 最近学习
	PopularContents []LearningContentResponse        `json:"popularContents"` // 热门内容
}
