package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/model/response"
	"gorm.io/gorm"
)

type LearningService struct{}

// CreateLearningContent 创建学习内容
func (s *LearningService) CreateLearningContent(req request.CreateLearningContentRequest) (model.LearningContent, error) {
	content := model.LearningContent{
		Title:        req.Title,
		ContentType:  req.Type,
		Category:     req.Category,
		Content:      req.Content,
		MediaUrl:     req.MediaUrl,
		Duration:     req.Duration,
		Difficulty:   req.Difficulty,
		Tags:         req.Tags,
		Status:       1, // 默认正常状态
		ViewCount:    0,
		LikeCount:    0,
		CommentCount: 0,
	}

	if err := global.GVA_DB.Create(&content).Error; err != nil {
		return content, err
	}

	return content, nil
}

// UpdateLearningContent 更新学习内容
func (s *LearningService) UpdateLearningContent(req request.UpdateLearningContentRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.MediaUrl != "" {
		updates["media_url"] = req.MediaUrl
	}
	if req.Duration > 0 {
		updates["duration"] = req.Duration
	}
	if req.Difficulty > 0 {
		updates["difficulty"] = req.Difficulty
	}
	if req.Tags != "" {
		updates["tags"] = req.Tags
	}

	return global.GVA_DB.Model(&model.LearningContent{}).
		Where("id = ?", req.ID).
		Updates(updates).Error
}

// DeleteLearningContent 删除学习内容
func (s *LearningService) DeleteLearningContent(id uint) error {
	return global.GVA_DB.Model(&model.LearningContent{}).Where("id = ?", id).Update("status", 0).Error
}

// GetLearningContent 获取单个学习内容
func (s *LearningService) GetLearningContent(id uint, userID uint) (response.LearningContentResponse, error) {
	var content model.LearningContent
	if err := global.GVA_DB.Where("id = ? AND status != 0", id).First(&content).Error; err != nil {
		return response.LearningContentResponse{}, err
	}

	// 增加浏览量
	global.GVA_DB.Model(&content).Update("view_count", gorm.Expr("view_count + 1"))

	resp := s.convertToResponse(content)

	// 获取用户相关信息
	if userID > 0 {
		s.setUserRelatedInfo(&resp, userID)
	}

	return resp, nil
}

// GetLearningContents 获取学习内容列表
func (s *LearningService) GetLearningContents(req request.GetLearningContentsRequest, userID uint) (response.LearningContentListResponse, error) {
	var contents []model.LearningContent
	var total int64

	db := global.GVA_DB.Model(&model.LearningContent{}).Where("status != 0")

	// 构建查询条件
	if req.Category > 0 {
		db = db.Where("category = ?", req.Category)
	}
	if req.Type > 0 {
		db = db.Where("content_type = ?", req.Type)
	}
	if req.Difficulty > 0 {
		db = db.Where("difficulty = ?", req.Difficulty)
	}
	if req.IsActive != nil {
		db = db.Where("status = ?", map[bool]int{true: 1, false: 0}[*req.IsActive])
	}
	if req.Tags != "" {
		tags := strings.Split(req.Tags, ",")
		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			if tag != "" {
				db = db.Where("tags LIKE ?", "%"+tag+"%")
			}
		}
	}

	// 计算总数
	db.Count(&total)

	// 排序 - 使用默认排序
	sortBy := "created_at"
	order := "desc"

	// 分页
	offset := (req.Page - 1) * req.PageSize
	if err := db.Order(fmt.Sprintf("%s %s", sortBy, order)).Offset(offset).Limit(req.PageSize).Find(&contents).Error; err != nil {
		return response.LearningContentListResponse{}, err
	}

	// 转换响应
	var respList []response.LearningContentResponse
	for _, content := range contents {
		resp := s.convertToResponse(content)
		if userID > 0 {
			s.setUserRelatedInfo(&resp, userID)
		}
		respList = append(respList, resp)
	}

	return response.LearningContentListResponse{
		List:     respList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// StartLearning 开始学习
func (s *LearningService) StartLearning(userID uint, contentID uint) (model.UserLearningRecord, error) {
	// 检查内容是否存在
	var content model.LearningContent
	if err := global.GVA_DB.Where("id = ? AND status = 1", contentID).First(&content).Error; err != nil {
		return model.UserLearningRecord{}, errors.New("学习内容不存在")
	}

	// 检查是否已存在学习记录
	var existingRecord model.UserLearningRecord
	if err := global.GVA_DB.Where("user_id = ? AND content_id = ?", userID, contentID).First(&existingRecord).Error; err == nil {
		// 已存在记录，更新开始时间
		existingRecord.StartTime = time.Now()
		if existingRecord.EndTime != nil {
			existingRecord.EndTime = nil
		}
		return existingRecord, global.GVA_DB.Save(&existingRecord).Error
	}

	// 创建新的学习记录
	record := model.UserLearningRecord{
		UserID:    userID,
		ContentID: contentID,
		StartTime: time.Now(),
		Progress:  0,
	}

	return record, global.GVA_DB.Create(&record).Error
}

// UpdateLearningProgress 更新学习进度
func (s *LearningService) UpdateLearningProgress(userID uint, contentID uint, progress int, duration int, isCompleted bool) error {
	var record model.UserLearningRecord
	if err := global.GVA_DB.Where("user_id = ? AND content_id = ?", userID, contentID).First(&record).Error; err != nil {
		return errors.New("学习记录不存在")
	}

	// 更新进度
	updates := map[string]interface{}{
		"progress": progress,
	}

	if duration > 0 {
		updates["duration"] = record.Duration + duration
	}

	if isCompleted {
		updates["is_completed"] = true
		if record.EndTime == nil {
			now := time.Now()
			updates["end_time"] = &now
		}
	}

	return global.GVA_DB.Model(&record).Updates(updates).Error
}

// ToggleLike 切换点赞状态
func (s *LearningService) ToggleLike(userID uint, contentID uint) (bool, error) {
	var record model.UserLearningRecord
	if err := global.GVA_DB.Where("user_id = ? AND content_id = ?", userID, contentID).First(&record).Error; err != nil {
		// 创建学习记录
		record = model.UserLearningRecord{
			UserID:    userID,
			ContentID: contentID,
			StartTime: time.Now(),
			IsLiked:   true,
		}
		if err := global.GVA_DB.Create(&record).Error; err != nil {
			return false, err
		}
		// 增加内容点赞数
		global.GVA_DB.Model(&model.LearningContent{}).Where("id = ?", contentID).Update("like_count", gorm.Expr("like_count + 1"))
		return true, nil
	}

	newLikedStatus := !record.IsLiked
	if err := global.GVA_DB.Model(&record).Update("is_liked", newLikedStatus).Error; err != nil {
		return false, err
	}

	// 更新内容点赞数
	if newLikedStatus {
		global.GVA_DB.Model(&model.LearningContent{}).Where("id = ?", contentID).Update("like_count", gorm.Expr("like_count + 1"))
	} else {
		global.GVA_DB.Model(&model.LearningContent{}).Where("id = ?", contentID).Update("like_count", gorm.Expr("like_count - 1"))
	}

	return newLikedStatus, nil
}

// ToggleCollect 切换收藏状态
func (s *LearningService) ToggleCollect(userID uint, contentID uint) (bool, error) {
	var record model.UserLearningRecord
	if err := global.GVA_DB.Where("user_id = ? AND content_id = ?", userID, contentID).First(&record).Error; err != nil {
		// 创建学习记录
		record = model.UserLearningRecord{
			UserID:      userID,
			ContentID:   contentID,
			StartTime:   time.Now(),
			IsCollected: true,
		}
		if err := global.GVA_DB.Create(&record).Error; err != nil {
			return false, err
		}
		// 增加内容收藏数
		global.GVA_DB.Model(&model.LearningContent{}).Where("id = ?", contentID).Update("collect_count", gorm.Expr("collect_count + 1"))
		return true, nil
	}

	newCollectedStatus := !record.IsCollected
	if err := global.GVA_DB.Model(&record).Update("is_collected", newCollectedStatus).Error; err != nil {
		return false, err
	}

	// 更新内容收藏数
	if newCollectedStatus {
		global.GVA_DB.Model(&model.LearningContent{}).Where("id = ?", contentID).Update("collect_count", gorm.Expr("collect_count + 1"))
	} else {
		global.GVA_DB.Model(&model.LearningContent{}).Where("id = ?", contentID).Update("collect_count", gorm.Expr("collect_count - 1"))
	}

	return newCollectedStatus, nil
}

// GetRecommendations 获取推荐内容（个性化推荐算法）
// 此方法暂时移除，等待相关请求类型定义完成

// convertToResponse 转换为响应格式
func (s *LearningService) convertToResponse(content model.LearningContent) response.LearningContentResponse {
	resp := response.LearningContentResponse{
		ID:           content.ID,
		CreatedAt:    content.CreatedAt,
		UpdatedAt:    content.UpdatedAt,
		Title:        content.Title,
		Summary:      content.Summary,
		Content:      content.Content,
		ContentType:  content.ContentType,
		Category:     content.Category,
		Difficulty:   content.Difficulty,
		Duration:     content.Duration,
		ThumbnailUrl: content.ThumbnailUrl,
		MediaUrl:     content.MediaUrl,
		Author:       content.Author,
		ViewCount:    content.ViewCount,
		LikeCount:    content.LikeCount,
		CollectCount: content.CollectCount,
		CommentCount: content.CommentCount,
		Status:       content.Status,
		PublishAt:    content.PublishAt,
		Tags:         content.Tags,
	}

	// 处理标签
	if content.Tags != "" {
		resp.TagList = strings.Split(content.Tags, ",")
		for i, tag := range resp.TagList {
			resp.TagList[i] = strings.TrimSpace(tag)
		}
	}

	return resp
}

// setUserRelatedInfo 设置用户相关信息
func (s *LearningService) setUserRelatedInfo(resp *response.LearningContentResponse, userID uint) {
	var record model.UserLearningRecord
	if err := global.GVA_DB.Where("user_id = ? AND content_id = ?", userID, resp.ID).First(&record).Error; err == nil {
		resp.IsLiked = record.IsLiked
		resp.IsCollected = record.IsCollected
		resp.UserRating = record.Rating

		if record.Progress > 0 {
			resp.LearningProgress = &response.UserLearningProgressResponse{
				ID:          record.ID,
				StartTime:   record.StartTime,
				EndTime:     record.EndTime,
				Duration:    record.Duration,
				Progress:    record.Progress,
				IsCompleted: record.IsCompleted,
				Rating:      record.Rating,
				Notes:       record.Notes,
			}
		}
	}
}

// GetUserLearningRecords 获取用户学习记录
// 此方法暂时移除，等待相关请求类型定义完成

// RateLearningContent 评分学习内容
func (s *LearningService) RateLearningContent(userID uint, req request.RateLearningContentRequest) error {
	var record model.UserLearningRecord
	if err := global.GVA_DB.Where("user_id = ? AND content_id = ?", userID, req.ContentID).First(&record).Error; err != nil {
		// 创建学习记录
		record = model.UserLearningRecord{
			UserID:    userID,
			ContentID: req.ContentID,
			StartTime: time.Now(),
			Rating:    req.Rating,
		}
		return global.GVA_DB.Create(&record).Error
	}

	// 更新评分
	return global.GVA_DB.Model(&record).Update("rating", req.Rating).Error
}

// GetLearningStats 获取学习统计
func (s *LearningService) GetLearningStats(userID uint) (response.LearningStatsResponse, error) {
	var stats response.LearningStatsResponse

	// 总内容数（用户有学习记录的）
	global.GVA_DB.Model(&model.UserLearningRecord{}).Where("user_id = ?", userID).Count(&stats.TotalContents)

	// 已完成内容数
	global.GVA_DB.Model(&model.UserLearningRecord{}).Where("user_id = ? AND is_completed = 1", userID).Count(&stats.CompletedContents)

	// 点赞内容数
	global.GVA_DB.Model(&model.UserLearningRecord{}).Where("user_id = ? AND is_liked = 1", userID).Count(&stats.LikedContents)

	// 收藏内容数
	global.GVA_DB.Model(&model.UserLearningRecord{}).Where("user_id = ? AND is_collected = 1", userID).Count(&stats.CollectedContents)

	// 总学习时长
	var totalDuration int
	global.GVA_DB.Model(&model.UserLearningRecord{}).Where("user_id = ?", userID).Select("COALESCE(SUM(duration), 0)").Scan(&totalDuration)
	stats.TotalLearningTime = totalDuration / 60 // 转换为分钟

	// 平均学习时长
	if stats.TotalContents > 0 {
		stats.AvgLearningTime = stats.TotalLearningTime / int(stats.TotalContents)
	}

	// 完成率
	if stats.TotalContents > 0 {
		stats.CompletionRate = int((stats.CompletedContents * 100) / stats.TotalContents)
	}

	// 连续学习天数（简化计算，基于最近的学习记录）
	var recentRecords []model.UserLearningRecord
	global.GVA_DB.Where("user_id = ?", userID).Order("start_time desc").Limit(30).Find(&recentRecords)

	continuous := 0
	yesterday := time.Now().AddDate(0, 0, -1)
	for _, record := range recentRecords {
		if record.StartTime.After(yesterday.Add(-24*time.Hour)) && record.StartTime.Before(yesterday.Add(24*time.Hour)) {
			continuous++
			yesterday = yesterday.AddDate(0, 0, -1)
		} else {
			break
		}
	}
	stats.ContinuousLearning = continuous

	// 最后学习时间
	var lastRecord model.UserLearningRecord
	if err := global.GVA_DB.Where("user_id = ?", userID).Order("start_time desc").First(&lastRecord).Error; err == nil {
		stats.LastLearningTime = &lastRecord.StartTime
	}

	return stats, nil
}

// GetCategoryStats 获取分类统计
func (s *LearningService) GetCategoryStats(userID uint) ([]response.CategoryStatsResponse, error) {
	var stats []response.CategoryStatsResponse

	categoryNames := map[int]string{
		1: "科普知识",
		2: "康复指导",
		3: "心理健康",
		4: "经验分享",
	}

	for category, name := range categoryNames {
		var stat response.CategoryStatsResponse
		stat.Category = category
		stat.CategoryName = name

		// 总内容数
		global.GVA_DB.Model(&model.UserLearningRecord{}).
			Joins("JOIN nofap_learning_contents ON nofap_user_learning_records.content_id = nofap_learning_contents.id").
			Where("nofap_user_learning_records.user_id = ? AND nofap_learning_contents.category = ?", userID, category).
			Count(&stat.TotalContents)

		// 已完成内容数
		global.GVA_DB.Model(&model.UserLearningRecord{}).
			Joins("JOIN nofap_learning_contents ON nofap_user_learning_records.content_id = nofap_learning_contents.id").
			Where("nofap_user_learning_records.user_id = ? AND nofap_learning_contents.category = ? AND nofap_user_learning_records.is_completed = 1", userID, category).
			Count(&stat.CompletedContents)

		// 完成率
		if stat.TotalContents > 0 {
			stat.CompletionRate = int((stat.CompletedContents * 100) / stat.TotalContents)
		}

		stats = append(stats, stat)
	}

	return stats, nil
}

// GetLearningHomepage 获取学习首页数据
func (s *LearningService) GetLearningHomepage(userID uint) (response.LearningHomepageResponse, error) {
	var homepage response.LearningHomepageResponse

	// 获取学习统计
	stats, err := s.GetLearningStats(userID)
	if err != nil {
		return homepage, err
	}
	homepage.Stats = stats

	// 获取分类统计
	categoryStats, err := s.GetCategoryStats(userID)
	if err != nil {
		return homepage, err
	}
	homepage.CategoryStats = categoryStats

	// 获取最近学习的内容
	var recentRecords []model.UserLearningRecord
	global.GVA_DB.Where("user_id = ?", userID).Preload("Content").Order("start_time desc").Limit(5).Find(&recentRecords)

	var recentContents []response.LearningContentResponse
	for _, record := range recentRecords {
		if record.Content.Status == 1 {
			resp := s.convertToResponse(record.Content)
			s.setUserRelatedInfo(&resp, userID)
			recentContents = append(recentContents, resp)
		}
	}
	homepage.RecentContents = recentContents

	// 获取热门内容
	var popularContents []model.LearningContent
	global.GVA_DB.Where("status = 1").Order("(view_count * 0.3 + like_count * 0.5 + collect_count * 0.2) desc").Limit(5).Find(&popularContents)

	var popularContentResp []response.LearningContentResponse
	for _, content := range popularContents {
		resp := s.convertToResponse(content)
		s.setUserRelatedInfo(&resp, userID)
		popularContentResp = append(popularContentResp, resp)
	}
	homepage.PopularContents = popularContentResp

	return homepage, nil
}
