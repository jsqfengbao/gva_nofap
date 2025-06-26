package miniprogram

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommunityService struct{}

// GetCategoryName 获取分类名称
func (s *CommunityService) GetCategoryName(category int) string {
	categories := map[int]string{
		1: "经验分享",
		2: "求助求鼓励",
		3: "日常打卡",
		4: "成功故事",
	}
	if name, exists := categories[category]; exists {
		return name
	}
	return "未知分类"
}

// CreatePost 创建帖子
func (s *CommunityService) CreatePost(userID uint, req request.CreatePostRequest) (*response.CreatePostResponse, error) {
	// 内容敏感词检查
	if err := s.checkSensitiveWords(req.Content); err != nil {
		return nil, err
	}

	// 创建帖子
	post := miniprogram.CommunityPost{
		UserID:      userID,
		Title:       strings.TrimSpace(req.Title),
		Content:     strings.TrimSpace(req.Content),
		Category:    req.Category,
		IsAnonymous: req.IsAnonymous,
		Status:      1, // 正常状态，如需审核可设为2
	}

	err := global.GVA_DB.Create(&post).Error
	if err != nil {
		return nil, fmt.Errorf("创建帖子失败: %w", err)
	}

	return &response.CreatePostResponse{
		ID:        post.ID,
		Title:     post.Title,
		CreatedAt: post.CreatedAt,
		Status:    post.Status,
		Message:   "帖子发布成功",
	}, nil
}

// UpdatePost 更新帖子
func (s *CommunityService) UpdatePost(userID uint, postID uint, req request.UpdatePostRequest) error {
	// 查找帖子
	var post miniprogram.CommunityPost
	err := global.GVA_DB.Where("id = ? AND user_id = ?", postID, userID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("帖子不存在或无权限修改")
		}
		return fmt.Errorf("查询帖子失败: %w", err)
	}

	// 内容敏感词检查
	if err := s.checkSensitiveWords(req.Content); err != nil {
		return err
	}

	// 更新帖子
	post.Title = strings.TrimSpace(req.Title)
	post.Content = strings.TrimSpace(req.Content)
	post.Category = req.Category
	post.IsAnonymous = req.IsAnonymous

	err = global.GVA_DB.Save(&post).Error
	if err != nil {
		return fmt.Errorf("更新帖子失败: %w", err)
	}

	return nil
}

// DeletePost 删除帖子
func (s *CommunityService) DeletePost(userID uint, postID uint) error {
	// 查找帖子
	var post miniprogram.CommunityPost
	err := global.GVA_DB.Where("id = ? AND user_id = ?", postID, userID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("帖子不存在或无权限删除")
		}
		return fmt.Errorf("查询帖子失败: %w", err)
	}

	// 开启事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 删除相关评论
	err = tx.Where("post_id = ?", postID).Delete(&miniprogram.CommunityComment{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("删除评论失败: %w", err)
	}

	// 删除相关点赞
	err = tx.Where("post_id = ?", postID).Delete(&miniprogram.CommunityLike{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("删除点赞失败: %w", err)
	}

	// 删除帖子
	err = tx.Delete(&post).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("删除帖子失败: %w", err)
	}

	tx.Commit()
	return nil
}

// GetPosts 获取帖子列表
func (s *CommunityService) GetPosts(userID uint, req request.GetPostsRequest) (*response.PostListResponse, error) {
	query := global.GVA_DB.Model(&miniprogram.CommunityPost{}).
		Preload("User").
		Where("status = 1") // 只显示正常状态的帖子

	// 分类筛选
	if req.Category > 0 {
		query = query.Where("category = ?", req.Category)
	}

	// 获取总数
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, fmt.Errorf("获取帖子总数失败: %w", err)
	}

	// 排序
	orderBy := fmt.Sprintf("%s %s", req.SortBy, req.Order)
	query = query.Order(orderBy)

	// 分页
	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	var posts []miniprogram.CommunityPost
	err = query.Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("获取帖子列表失败: %w", err)
	}

	// 获取用户点赞状态
	likedPosts := make(map[uint]bool)
	if userID > 0 {
		var likes []miniprogram.CommunityLike
		postIDs := make([]uint, len(posts))
		for i, post := range posts {
			postIDs[i] = post.ID
		}
		global.GVA_DB.Where("user_id = ? AND like_type = 1 AND target_id IN ?", userID, postIDs).Find(&likes)
		for _, like := range likes {
			likedPosts[like.TargetID] = true
		}
	}

	// 构建响应
	var list []response.PostItem
	for _, post := range posts {
		item := response.PostItem{
			ID:           post.ID,
			Title:        post.Title,
			Content:      s.truncateContent(post.Content, 200), // 列表中截断内容
			Category:     post.Category,
			CategoryName: s.GetCategoryName(post.Category),
			IsAnonymous:  post.IsAnonymous,
			ViewCount:    post.ViewCount,
			LikeCount:    post.LikeCount,
			CommentCount: post.CommentCount,
			Status:       post.Status,
			CreatedAt:    post.CreatedAt,
			UpdatedAt:    post.UpdatedAt,
			UserID:       post.UserID,
			IsLiked:      likedPosts[post.ID],
			CanEdit:      post.UserID == userID,
			CanDelete:    post.UserID == userID,
		}

		// 设置用户信息（如果不是匿名）
		if !post.IsAnonymous && post.User.ID > 0 {
			item.UserNickname = post.User.Nickname
			item.UserAvatar = post.User.AvatarUrl
		} else {
			item.UserNickname = "匿名用户"
			item.UserAvatar = ""
		}

		list = append(list, item)
	}

	return &response.PostListResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		HasMore:  int64(req.Page*req.PageSize) < total,
	}, nil
}

// GetPostDetail 获取帖子详情
func (s *CommunityService) GetPostDetail(userID uint, postID uint) (*response.PostDetailResponse, error) {
	// 查找帖子
	var post miniprogram.CommunityPost
	err := global.GVA_DB.Preload("User").Where("id = ? AND status = 1", postID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("帖子不存在")
		}
		return nil, fmt.Errorf("查询帖子失败: %w", err)
	}

	// 增加浏览次数
	global.GVA_DB.Model(&post).UpdateColumn("view_count", gorm.Expr("view_count + 1"))

	// 检查用户是否点赞
	var isLiked bool
	if userID > 0 {
		var count int64
		global.GVA_DB.Model(&miniprogram.CommunityLike{}).
			Where("user_id = ? AND like_type = 1 AND target_id = ?", userID, postID).
			Count(&count)
		isLiked = count > 0
	}

	// 获取评论列表
	comments, err := s.getPostComments(userID, postID)
	if err != nil {
		global.GVA_LOG.Error("获取评论列表失败", zap.Error(err))
		comments = []response.CommentItem{} // 评论获取失败不影响帖子展示
	}

	// 构建响应
	item := response.PostItem{
		ID:           post.ID,
		Title:        post.Title,
		Content:      post.Content, // 详情页显示完整内容
		Category:     post.Category,
		CategoryName: s.GetCategoryName(post.Category),
		IsAnonymous:  post.IsAnonymous,
		ViewCount:    post.ViewCount + 1, // 包含本次浏览
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
		Status:       post.Status,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
		UserID:       post.UserID,
		IsLiked:      isLiked,
		CanEdit:      post.UserID == userID,
		CanDelete:    post.UserID == userID,
	}

	// 设置用户信息
	if !post.IsAnonymous && post.User.ID > 0 {
		item.UserNickname = post.User.Nickname
		item.UserAvatar = post.User.AvatarUrl
	} else {
		item.UserNickname = "匿名用户"
		item.UserAvatar = ""
	}

	return &response.PostDetailResponse{
		PostItem: item,
		Comments: comments,
	}, nil
}

// CreateComment 创建评论
func (s *CommunityService) CreateComment(userID uint, req request.CreateCommentRequest) (*response.CreateCommentResponse, error) {
	// 检查帖子是否存在
	var post miniprogram.CommunityPost
	err := global.GVA_DB.Where("id = ? AND status = 1", req.PostID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("帖子不存在")
		}
		return nil, fmt.Errorf("查询帖子失败: %w", err)
	}

	// 如果是回复评论，检查父评论是否存在
	if req.ParentID > 0 {
		var parentComment miniprogram.CommunityComment
		err = global.GVA_DB.Where("id = ? AND post_id = ? AND status = 1", req.ParentID, req.PostID).First(&parentComment).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("父评论不存在")
			}
			return nil, fmt.Errorf("查询父评论失败: %w", err)
		}
	}

	// 内容敏感词检查
	if err := s.checkSensitiveWords(req.Content); err != nil {
		return nil, err
	}

	// 开启事务
	tx := global.GVA_DB.Begin()

	// 创建评论
	comment := miniprogram.CommunityComment{
		PostID:      req.PostID,
		UserID:      userID,
		Content:     strings.TrimSpace(req.Content),
		ParentID:    req.ParentID,
		IsAnonymous: req.IsAnonymous,
		Status:      1, // 正常状态
	}

	err = tx.Create(&comment).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建评论失败: %w", err)
	}

	// 更新帖子评论数
	err = tx.Model(&post).UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新评论数失败: %w", err)
	}

	tx.Commit()

	return &response.CreateCommentResponse{
		ID:        comment.ID,
		PostID:    comment.PostID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Status:    comment.Status,
		Message:   "评论发布成功",
	}, nil
}

// DeleteComment 删除评论
func (s *CommunityService) DeleteComment(userID uint, commentID uint) error {
	// 查找评论
	var comment miniprogram.CommunityComment
	err := global.GVA_DB.Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在或无权限删除")
		}
		return fmt.Errorf("查询评论失败: %w", err)
	}

	// 开启事务
	tx := global.GVA_DB.Begin()

	// 删除评论及其所有子评论
	err = tx.Where("id = ? OR parent_id = ?", commentID, commentID).Delete(&miniprogram.CommunityComment{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("删除评论失败: %w", err)
	}

	// 更新帖子评论数
	var remainingCount int64
	tx.Model(&miniprogram.CommunityComment{}).Where("post_id = ? AND status = 1", comment.PostID).Count(&remainingCount)
	tx.Model(&miniprogram.CommunityPost{}).Where("id = ?", comment.PostID).UpdateColumn("comment_count", remainingCount)

	tx.Commit()
	return nil
}

// ToggleLike 点赞/取消点赞
func (s *CommunityService) ToggleLike(userID uint, req request.LikeRequest) (*response.LikeResponse, error) {
	// 检查目标是否存在
	if err := s.checkLikeTarget(req.TargetID, req.LikeType); err != nil {
		return nil, err
	}

	// 查找现有点赞记录
	var existingLike miniprogram.CommunityLike
	err := global.GVA_DB.Where("user_id = ? AND like_type = ? AND target_id = ?", userID, req.LikeType, req.TargetID).First(&existingLike).Error

	tx := global.GVA_DB.Begin()
	var isLiked bool
	var newLikeCount int

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 创建点赞记录
		like := miniprogram.CommunityLike{
			UserID:   userID,
			LikeType: req.LikeType,
			TargetID: req.TargetID,
		}

		// 设置PostID
		if req.LikeType == 1 {
			like.PostID = req.TargetID
		} else {
			// 通过评论ID获取PostID
			var comment miniprogram.CommunityComment
			err = global.GVA_DB.Where("id = ?", req.TargetID).First(&comment).Error
			if err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("查询评论失败: %w", err)
			}
			like.PostID = comment.PostID
		}

		err = tx.Create(&like).Error
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("创建点赞记录失败: %w", err)
		}

		isLiked = true
		newLikeCount, _ = s.updateLikeCount(tx, req.TargetID, req.LikeType, 1)
	} else if err == nil {
		// 删除点赞记录
		err = tx.Delete(&existingLike).Error
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("删除点赞记录失败: %w", err)
		}

		isLiked = false
		newLikeCount, _ = s.updateLikeCount(tx, req.TargetID, req.LikeType, -1)
	} else {
		tx.Rollback()
		return nil, fmt.Errorf("查询点赞记录失败: %w", err)
	}

	tx.Commit()

	message := "取消点赞"
	if isLiked {
		message = "点赞成功"
	}

	return &response.LikeResponse{
		TargetID:  req.TargetID,
		LikeType:  req.LikeType,
		IsLiked:   isLiked,
		LikeCount: newLikeCount,
		Message:   message,
	}, nil
}

// 辅助方法

// checkSensitiveWords 敏感词检查
func (s *CommunityService) checkSensitiveWords(content string) error {
	// 简单的敏感词列表，实际项目中应该使用更完善的敏感词过滤系统
	sensitiveWords := []string{"垃圾", "spam", "广告"}

	content = strings.ToLower(content)
	for _, word := range sensitiveWords {
		if strings.Contains(content, word) {
			return errors.New("内容包含敏感词，请修改后再试")
		}
	}
	return nil
}

// truncateContent 截断内容
func (s *CommunityService) truncateContent(content string, maxLen int) string {
	runes := []rune(content)
	if len(runes) <= maxLen {
		return content
	}
	return string(runes[:maxLen]) + "..."
}

// getPostComments 获取帖子评论
func (s *CommunityService) getPostComments(userID uint, postID uint) ([]response.CommentItem, error) {
	var comments []miniprogram.CommunityComment
	err := global.GVA_DB.Preload("User").
		Where("post_id = ? AND status = 1", postID).
		Order("created_at ASC").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}

	// 获取用户点赞状态
	likedComments := make(map[uint]bool)
	if userID > 0 {
		var likes []miniprogram.CommunityLike
		commentIDs := make([]uint, len(comments))
		for i, comment := range comments {
			commentIDs[i] = comment.ID
		}
		global.GVA_DB.Where("user_id = ? AND like_type = 2 AND target_id IN ?", userID, commentIDs).Find(&likes)
		for _, like := range likes {
			likedComments[like.TargetID] = true
		}
	}

	// 构建树形结构
	var result []response.CommentItem
	commentMap := make(map[uint]*response.CommentItem)

	// 先创建所有评论项
	for _, comment := range comments {
		item := &response.CommentItem{
			ID:          comment.ID,
			PostID:      comment.PostID,
			Content:     comment.Content,
			ParentID:    comment.ParentID,
			IsAnonymous: comment.IsAnonymous,
			LikeCount:   comment.LikeCount,
			Status:      comment.Status,
			CreatedAt:   comment.CreatedAt,
			UserID:      comment.UserID,
			IsLiked:     likedComments[comment.ID],
			CanDelete:   comment.UserID == userID,
			Children:    []response.CommentItem{},
		}

		// 设置用户信息
		if !comment.IsAnonymous && comment.User.ID > 0 {
			item.UserNickname = comment.User.Nickname
			item.UserAvatar = comment.User.AvatarUrl
		} else {
			item.UserNickname = "匿名用户"
			item.UserAvatar = ""
		}

		commentMap[comment.ID] = item
	}

	// 构建父子关系
	for _, comment := range comments {
		item := commentMap[comment.ID]
		if comment.ParentID == 0 {
			// 顶级评论
			result = append(result, *item)
		} else {
			// 子评论
			if parent, exists := commentMap[comment.ParentID]; exists {
				parent.Children = append(parent.Children, *item)
				// 设置父评论信息
				if parentComment := commentMap[comment.ParentID]; parentComment != nil {
					item.ParentContent = s.truncateContent(parentComment.Content, 50)
					item.ParentNickname = parentComment.UserNickname
				}
			}
		}
	}

	return result, nil
}

// checkLikeTarget 检查点赞目标是否存在
func (s *CommunityService) checkLikeTarget(targetID uint, likeType int) error {
	if likeType == 1 {
		// 检查帖子
		var count int64
		err := global.GVA_DB.Model(&miniprogram.CommunityPost{}).Where("id = ? AND status = 1", targetID).Count(&count).Error
		if err != nil {
			return fmt.Errorf("查询帖子失败: %w", err)
		}
		if count == 0 {
			return errors.New("帖子不存在")
		}
	} else if likeType == 2 {
		// 检查评论
		var count int64
		err := global.GVA_DB.Model(&miniprogram.CommunityComment{}).Where("id = ? AND status = 1", targetID).Count(&count).Error
		if err != nil {
			return fmt.Errorf("查询评论失败: %w", err)
		}
		if count == 0 {
			return errors.New("评论不存在")
		}
	} else {
		return errors.New("无效的点赞类型")
	}
	return nil
}

// updateLikeCount 更新点赞数
func (s *CommunityService) updateLikeCount(tx *gorm.DB, targetID uint, likeType int, delta int) (int, error) {
	var newCount int64

	if likeType == 1 {
		// 更新帖子点赞数
		err := tx.Model(&miniprogram.CommunityPost{}).Where("id = ?", targetID).
			UpdateColumn("like_count", gorm.Expr("like_count + ?", delta)).Error
		if err != nil {
			return 0, err
		}

		// 获取新的点赞数
		tx.Model(&miniprogram.CommunityPost{}).Where("id = ?", targetID).Select("like_count").Scan(&newCount)
	} else {
		// 更新评论点赞数
		err := tx.Model(&miniprogram.CommunityComment{}).Where("id = ?", targetID).
			UpdateColumn("like_count", gorm.Expr("like_count + ?", delta)).Error
		if err != nil {
			return 0, err
		}

		// 获取新的点赞数
		tx.Model(&miniprogram.CommunityComment{}).Where("id = ?", targetID).Select("like_count").Scan(&newCount)
	}

	return int(newCount), nil
}

// GetCommunityStats 获取社区统计信息
func (s *CommunityService) GetCommunityStats() (*response.CommunityStatsResponse, error) {
	var stats response.CommunityStatsResponse

	// 获取总帖子数
	global.GVA_DB.Model(&miniprogram.CommunityPost{}).Where("status = 1").Count(&stats.TotalPosts)

	// 获取总评论数
	global.GVA_DB.Model(&miniprogram.CommunityComment{}).Where("status = 1").Count(&stats.TotalComments)

	// 获取总点赞数
	global.GVA_DB.Model(&miniprogram.CommunityLike{}).Count(&stats.TotalLikes)

	// 获取今日帖子数
	today := time.Now().Format("2006-01-02")
	global.GVA_DB.Model(&miniprogram.CommunityPost{}).
		Where("status = 1 AND DATE(created_at) = ?", today).
		Count(&stats.TodayPosts)

	// 获取分类统计
	var categoryStats []response.CategoryStats
	for i := 1; i <= 4; i++ {
		var count int64
		global.GVA_DB.Model(&miniprogram.CommunityPost{}).
			Where("status = 1 AND category = ?", i).
			Count(&count)

		categoryStats = append(categoryStats, response.CategoryStats{
			Category:     i,
			CategoryName: s.GetCategoryName(i),
			PostCount:    count,
		})
	}
	stats.CategoryStats = categoryStats

	return &stats, nil
}

// GetUserPosts 获取用户发帖记录
func (s *CommunityService) GetUserPosts(userID uint, req request.GetPostsRequest) (*response.UserPostsResponse, error) {
	query := global.GVA_DB.Model(&miniprogram.CommunityPost{}).
		Preload("User").
		Where("user_id = ? AND status = 1", userID)

	// 分类筛选
	if req.Category > 0 {
		query = query.Where("category = ?", req.Category)
	}

	// 获取总数
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, fmt.Errorf("获取用户帖子总数失败: %w", err)
	}

	// 排序
	orderBy := fmt.Sprintf("%s %s", req.SortBy, req.Order)
	query = query.Order(orderBy)

	// 分页
	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	var posts []miniprogram.CommunityPost
	err = query.Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("获取用户帖子列表失败: %w", err)
	}

	// 构建响应
	var list []response.PostItem
	for _, post := range posts {
		item := response.PostItem{
			ID:           post.ID,
			Title:        post.Title,
			Content:      s.truncateContent(post.Content, 200),
			Category:     post.Category,
			CategoryName: s.GetCategoryName(post.Category),
			IsAnonymous:  post.IsAnonymous,
			ViewCount:    post.ViewCount,
			LikeCount:    post.LikeCount,
			CommentCount: post.CommentCount,
			Status:       post.Status,
			CreatedAt:    post.CreatedAt,
			UpdatedAt:    post.UpdatedAt,
			UserID:       post.UserID,
			IsLiked:      false, // 自己的帖子不显示点赞状态
			CanEdit:      true,  // 都是自己的帖子，可以编辑
			CanDelete:    true,  // 都是自己的帖子，可以删除
		}

		// 设置用户信息
		if !post.IsAnonymous && post.User.ID > 0 {
			item.UserNickname = post.User.Nickname
			item.UserAvatar = post.User.AvatarUrl
		} else {
			item.UserNickname = "匿名用户"
			item.UserAvatar = ""
		}

		list = append(list, item)
	}

	return &response.UserPostsResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		HasMore:  int64(req.Page*req.PageSize) < total,
	}, nil
}
