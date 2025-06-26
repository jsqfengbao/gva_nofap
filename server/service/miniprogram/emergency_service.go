package miniprogram

import (
	"database/sql"
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/request"
	miniprogramRes "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/response"
	"gorm.io/gorm"
)

type EmergencyService struct{}

// CreateEmergencyHelp 创建紧急求助
func (e *EmergencyService) CreateEmergencyHelp(userID uint, req request.CreateEmergencyHelpRequest) (miniprogram.EmergencyHelp, error) {
	help := miniprogram.EmergencyHelp{
		UserID:      userID,
		Type:        req.Type,
		Priority:    req.Priority,
		Description: req.Description,
		Location:    req.Location,
		IsAnonymous: req.IsAnonymous,
		Status:      1, // 待响应
	}

	// 设置默认优先级
	if help.Priority == 0 {
		help.Priority = 2 // 默认中等优先级
	}

	if err := global.GVA_DB.Create(&help).Error; err != nil {
		return help, err
	}

	// 自动匹配志愿者(异步处理)
	go e.autoMatchVolunteer(help.ID)

	return help, nil
}

// GetEmergencyHelps 获取紧急求助列表
func (e *EmergencyService) GetEmergencyHelps(userID uint, req request.GetEmergencyHelpsRequest) (miniprogramRes.EmergencyHelpListResponse, error) {
	var helps []miniprogram.EmergencyHelp
	var total int64

	db := global.GVA_DB.Model(&miniprogram.EmergencyHelp{})

	// 如果不是管理员，只能看自己的记录
	if req.UserID == 0 {
		db = db.Where("user_id = ?", userID)
	}

	// 筛选条件
	if req.Type > 0 {
		db = db.Where("type = ?", req.Type)
	}
	if req.Status > 0 {
		db = db.Where("status = ?", req.Status)
	}
	if req.Priority > 0 {
		db = db.Where("priority = ?", req.Priority)
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return miniprogramRes.EmergencyHelpListResponse{}, err
	}

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	if err := db.Preload("User").Preload("Volunteer").Preload("Responses").
		Order("priority DESC, created_at DESC").
		Offset(offset).Limit(req.PageSize).Find(&helps).Error; err != nil {
		return miniprogramRes.EmergencyHelpListResponse{}, err
	}

	// 转换响应
	var list []miniprogramRes.EmergencyHelpItem
	for _, help := range helps {
		list = append(list, miniprogramRes.ConvertToEmergencyHelpItem(help))
	}

	return miniprogramRes.EmergencyHelpListResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		HasMore:  int64(offset+req.PageSize) < total,
	}, nil
}

// GetEmergencyHelpDetail 获取紧急求助详情
func (e *EmergencyService) GetEmergencyHelpDetail(userID uint, helpID uint) (miniprogramRes.EmergencyHelpDetailResponse, error) {
	var help miniprogram.EmergencyHelp

	if err := global.GVA_DB.Preload("User").Preload("Volunteer").
		Preload("Responses.Volunteer").First(&help, helpID).Error; err != nil {
		return miniprogramRes.EmergencyHelpDetailResponse{}, err
	}

	// 转换响应项目
	item := miniprogramRes.ConvertToEmergencyHelpItem(help)

	// 转换响应列表
	var responses []miniprogramRes.EmergencyResponseItem
	for _, resp := range help.Responses {
		typeName := "文字回复"
		switch resp.Type {
		case 1:
			typeName = "文字回复"
		case 2:
			typeName = "语音通话"
		case 3:
			typeName = "视频通话"
		case 4:
			typeName = "陪伴聊天"
		}

		responseItem := miniprogramRes.EmergencyResponseItem{
			ID:                resp.ID,
			Type:              resp.Type,
			TypeName:          typeName,
			Content:           resp.Content,
			Duration:          resp.Duration,
			IsActive:          resp.IsActive,
			VolunteerNickname: resp.Volunteer.Nickname,
			VolunteerAvatar:   resp.Volunteer.AvatarUrl,
			CreatedAt:         resp.CreatedAt,
			EndedAt:           resp.EndedAt,
		}
		responses = append(responses, responseItem)
	}

	// 检查用户权限
	canRespond := e.canUserRespond(userID, help)
	isOwner := help.UserID == userID

	return miniprogramRes.EmergencyHelpDetailResponse{
		EmergencyHelpItem: item,
		Responses:         responses,
		CanRespond:        canRespond,
		IsOwner:           isOwner,
	}, nil
}

// UpdateEmergencyHelp 更新紧急求助
func (e *EmergencyService) UpdateEmergencyHelp(userID uint, req request.UpdateEmergencyHelpRequest) error {
	var help miniprogram.EmergencyHelp

	if err := global.GVA_DB.First(&help, req.ID).Error; err != nil {
		return err
	}

	// 检查权限
	if help.UserID != userID {
		return errors.New("无权限操作")
	}

	// 更新字段
	updates := make(map[string]interface{})
	if req.Status > 0 {
		updates["status"] = req.Status
		if req.Status == 3 { // 已解决
			now := time.Now()
			updates["resolved_at"] = &now
		}
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Rating != nil {
		updates["rating"] = *req.Rating
	}
	if req.Feedback != "" {
		updates["feedback"] = req.Feedback
	}

	return global.GVA_DB.Model(&help).Updates(updates).Error
}

// CreateEmergencyResponse 创建紧急求助响应
func (e *EmergencyService) CreateEmergencyResponse(volunteerID uint, req request.CreateEmergencyResponseRequest) error {
	// 检查志愿者权限
	if !e.isActiveVolunteer(volunteerID) {
		return errors.New("非活跃志愿者")
	}

	// 检查求助记录
	var help miniprogram.EmergencyHelp
	if err := global.GVA_DB.First(&help, req.HelpID).Error; err != nil {
		return err
	}

	if help.Status == 3 || help.Status == 4 {
		return errors.New("求助已结束")
	}

	// 创建响应
	response := miniprogram.EmergencyResponse{
		HelpID:      req.HelpID,
		VolunteerID: volunteerID,
		Type:        req.Type,
		Content:     req.Content,
		IsActive:    true,
	}

	tx := global.GVA_DB.Begin()

	// 创建响应记录
	if err := tx.Create(&response).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新求助状态
	now := time.Now()
	updates := map[string]interface{}{
		"status":       2, // 进行中
		"volunteer_id": volunteerID,
		"responded_at": &now,
	}
	if err := tx.Model(&help).Updates(updates).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新志愿者当前服务数
	if err := tx.Model(&miniprogram.EmergencyVolunteer{}).
		Where("user_id = ?", volunteerID).
		Update("current_load", gorm.Expr("current_load + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// ConnectVolunteer 连接志愿者
func (e *EmergencyService) ConnectVolunteer(userID uint, req request.ConnectVolunteerRequest) (miniprogramRes.ConnectVolunteerResponse, error) {
	// 检查求助记录
	var help miniprogram.EmergencyHelp
	if err := global.GVA_DB.First(&help, req.HelpID).Error; err != nil {
		return miniprogramRes.ConnectVolunteerResponse{}, err
	}

	if help.UserID != userID {
		return miniprogramRes.ConnectVolunteerResponse{}, errors.New("无权限操作")
	}

	if help.Status != 1 {
		return miniprogramRes.ConnectVolunteerResponse{}, errors.New("求助状态不允许连接")
	}

	// 查找可用志愿者
	volunteer, err := e.findAvailableVolunteer(help.Type, help.Priority)
	if err != nil {
		return miniprogramRes.ConnectVolunteerResponse{
			Success: false,
			Message: "暂时没有可用志愿者，请稍后再试",
		}, nil
	}

	// 创建连接响应
	respReq := request.CreateEmergencyResponseRequest{
		HelpID:  req.HelpID,
		Type:    4, // 陪伴聊天
		Content: "志愿者已连接，开始提供帮助",
	}

	if err := e.CreateEmergencyResponse(volunteer.UserID, respReq); err != nil {
		return miniprogramRes.ConnectVolunteerResponse{
			Success: false,
			Message: "连接失败，请重试",
		}, nil
	}

	volunteerItem := miniprogramRes.ConvertToEmergencyVolunteerItem(*volunteer)

	return miniprogramRes.ConnectVolunteerResponse{
		Success:   true,
		Message:   "成功连接志愿者",
		Volunteer: &volunteerItem,
	}, nil
}

// RegisterVolunteer 注册志愿者
func (e *EmergencyService) RegisterVolunteer(userID uint, req request.EmergencyVolunteerRequest) error {
	// 检查是否已注册
	var existing miniprogram.EmergencyVolunteer
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&existing).Error; err == nil {
		return errors.New("已经是志愿者")
	}

	volunteer := miniprogram.EmergencyVolunteer{
		UserID:       userID,
		Status:       1, // 待审核
		IsOnline:     false,
		Specialties:  req.Specialties,
		MaxLoad:      req.MaxLoad,
		CurrentLoad:  0,
		TotalHelped:  0,
		AvgRating:    0,
		RatingCount:  0,
		LastActiveAt: time.Now(),
	}

	if volunteer.MaxLoad == 0 {
		volunteer.MaxLoad = 3 // 默认最大服务3人
	}

	return global.GVA_DB.Create(&volunteer).Error
}

// UpdateVolunteerStatus 更新志愿者状态
func (e *EmergencyService) UpdateVolunteerStatus(userID uint, req request.UpdateVolunteerStatusRequest) error {
	return global.GVA_DB.Model(&miniprogram.EmergencyVolunteer{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"is_online":      req.IsOnline,
			"last_active_at": time.Now(),
		}).Error
}

// GetEmergencyResources 获取紧急求助资源
func (e *EmergencyService) GetEmergencyResources(req request.GetEmergencyResourcesRequest) (miniprogramRes.ResourceListResponse, error) {
	var resources []miniprogram.EmergencyResource
	var total int64

	db := global.GVA_DB.Model(&miniprogram.EmergencyResource{}).Where("is_active = ?", true)

	// 筛选条件
	if req.Type > 0 {
		db = db.Where("type = ?", req.Type)
	}
	if req.Difficulty > 0 {
		db = db.Where("difficulty = ?", req.Difficulty)
	}
	if req.Tags != "" {
		db = db.Where("tags LIKE ?", "%"+req.Tags+"%")
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return miniprogramRes.ResourceListResponse{}, err
	}

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	if err := db.Order("usage_count DESC, created_at DESC").
		Offset(offset).Limit(req.PageSize).Find(&resources).Error; err != nil {
		return miniprogramRes.ResourceListResponse{}, err
	}

	// 转换响应
	var list []miniprogramRes.EmergencyResourceItem
	for _, resource := range resources {
		list = append(list, miniprogramRes.ConvertToEmergencyResourceItem(resource))
	}

	return miniprogramRes.ResourceListResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		HasMore:  int64(offset+req.PageSize) < total,
	}, nil
}

// UseResource 使用资源(增加使用次数)
func (e *EmergencyService) UseResource(resourceID uint) error {
	return global.GVA_DB.Model(&miniprogram.EmergencyResource{}).
		Where("id = ?", resourceID).
		Update("usage_count", gorm.Expr("usage_count + ?", 1)).Error
}

// RateResource 资源评分
func (e *EmergencyService) RateResource(userID uint, req request.RateResourceRequest) error {
	// 检查资源是否存在
	var resource miniprogram.EmergencyResource
	if err := global.GVA_DB.First(&resource, req.ResourceID).Error; err != nil {
		return err
	}

	// 更新评分(简化处理，实际应该记录每个用户的评分)
	newRatingCount := resource.RatingCount + 1
	newAvgRating := (resource.AvgRating*float64(resource.RatingCount) + float64(req.Rating)) / float64(newRatingCount)

	return global.GVA_DB.Model(&resource).Updates(map[string]interface{}{
		"avg_rating":   newAvgRating,
		"rating_count": newRatingCount,
	}).Error
}

// GetEmergencyStats 获取紧急求助统计
func (e *EmergencyService) GetEmergencyStats() (miniprogramRes.EmergencyStatsResponse, error) {
	var stats miniprogramRes.EmergencyStatsResponse

	// 基础统计
	global.GVA_DB.Model(&miniprogram.EmergencyHelp{}).Count(&stats.TotalHelps)
	global.GVA_DB.Model(&miniprogram.EmergencyHelp{}).Where("status = ?", 1).Count(&stats.PendingHelps)
	global.GVA_DB.Model(&miniprogram.EmergencyHelp{}).Where("status = ?", 2).Count(&stats.InProgressHelps)
	global.GVA_DB.Model(&miniprogram.EmergencyHelp{}).Where("status = ?", 3).Count(&stats.ResolvedHelps)
	global.GVA_DB.Model(&miniprogram.EmergencyVolunteer{}).Where("is_online = ? AND status = ?", true, 2).Count(&stats.OnlineVolunteers)

	// 平均响应时间(分钟)
	var avgResponse sql.NullFloat64
	global.GVA_DB.Model(&miniprogram.EmergencyHelp{}).
		Where("responded_at IS NOT NULL").
		Select("AVG(TIMESTAMPDIFF(MINUTE, created_at, responded_at))").
		Scan(&avgResponse)
	if avgResponse.Valid {
		stats.AvgResponseTime = avgResponse.Float64
	}

	// 平均解决时间(分钟)
	var avgResolution sql.NullFloat64
	global.GVA_DB.Model(&miniprogram.EmergencyHelp{}).
		Where("resolved_at IS NOT NULL").
		Select("AVG(TIMESTAMPDIFF(MINUTE, created_at, resolved_at))").
		Scan(&avgResolution)
	if avgResolution.Valid {
		stats.AvgResolutionTime = avgResolution.Float64
	}

	// 按类型统计
	var typeStats []struct {
		Type  int   `json:"type"`
		Count int64 `json:"count"`
	}
	global.GVA_DB.Model(&miniprogram.EmergencyHelp{}).
		Select("type, COUNT(*) as count").
		Group("type").Find(&typeStats)

	for _, ts := range typeStats {
		help := miniprogram.EmergencyHelp{Type: ts.Type}
		stats.HelpsByType = append(stats.HelpsByType, miniprogramRes.TypeStat{
			Type:     ts.Type,
			TypeName: help.GetTypeName(),
			Count:    ts.Count,
		})
	}

	// 按优先级统计
	var priorityStats []struct {
		Priority int   `json:"priority"`
		Count    int64 `json:"count"`
	}
	global.GVA_DB.Model(&miniprogram.EmergencyHelp{}).
		Select("priority, COUNT(*) as count").
		Group("priority").Find(&priorityStats)

	for _, ps := range priorityStats {
		help := miniprogram.EmergencyHelp{Priority: ps.Priority}
		stats.HelpsByPriority = append(stats.HelpsByPriority, miniprogramRes.PriorityStat{
			Priority:     ps.Priority,
			PriorityName: help.GetPriorityName(),
			Count:        ps.Count,
		})
	}

	return stats, nil
}

// 私有方法

// autoMatchVolunteer 自动匹配志愿者
func (e *EmergencyService) autoMatchVolunteer(helpID uint) {
	var help miniprogram.EmergencyHelp
	if err := global.GVA_DB.First(&help, helpID).Error; err != nil {
		return
	}

	volunteer, err := e.findAvailableVolunteer(help.Type, help.Priority)
	if err != nil {
		return
	}

	// 发送通知给志愿者(这里可以集成推送服务)
	// TODO: 发送推送通知
	_ = volunteer
}

// findAvailableVolunteer 查找可用志愿者
func (e *EmergencyService) findAvailableVolunteer(helpType, priority int) (*miniprogram.EmergencyVolunteer, error) {
	var volunteer miniprogram.EmergencyVolunteer

	// 优先选择专长匹配且当前负载较低的志愿者
	db := global.GVA_DB.Preload("User").
		Where("status = ? AND is_online = ? AND current_load < max_load", 2, true)

	// 根据优先级排序
	if priority >= 3 { // 高优先级
		db = db.Order("current_load ASC, avg_rating DESC")
	} else {
		db = db.Order("avg_rating DESC, current_load ASC")
	}

	if err := db.First(&volunteer).Error; err != nil {
		return nil, err
	}

	return &volunteer, nil
}

// canUserRespond 检查用户是否可以响应
func (e *EmergencyService) canUserRespond(userID uint, help miniprogram.EmergencyHelp) bool {
	// 不能响应自己的求助
	if help.UserID == userID {
		return false
	}

	// 检查是否是活跃志愿者
	return e.isActiveVolunteer(userID)
}

// isActiveVolunteer 检查是否是活跃志愿者
func (e *EmergencyService) isActiveVolunteer(userID uint) bool {
	var count int64
	global.GVA_DB.Model(&miniprogram.EmergencyVolunteer{}).
		Where("user_id = ? AND status = ? AND current_load < max_load", userID, 2).
		Count(&count)
	return count > 0
}
