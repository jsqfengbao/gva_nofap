package miniprogram

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram"
	miniprogramReq "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/request"
	miniprogramRes "github.com/flipped-aurora/gin-vue-admin/server/model/miniprogram/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct{}

// FindOrCreateWxUser 查找或创建微信用户
func (s *UserService) FindOrCreateWxUser(openid, unionid string, userInfo *miniprogramReq.WxUserInfo) (*miniprogram.WxUser, error) {
	var user miniprogram.WxUser

	// 先查找用户是否存在
	err := global.GVA_DB.Where("openid = ?", openid).First(&user).Error
	if err == nil {
		// 用户存在，更新最后登录时间和用户信息
		updates := map[string]interface{}{
			"last_login_at": time.Now(),
		}

		// 如果有新的用户信息，更新用户资料
		if userInfo != nil {
			if userInfo.Nickname != "" {
				updates["nickname"] = userInfo.Nickname
			}
			if userInfo.AvatarUrl != "" {
				updates["avatar_url"] = userInfo.AvatarUrl
			}
			updates["gender"] = userInfo.Gender
			updates["city"] = userInfo.City
			updates["province"] = userInfo.Province
			updates["country"] = userInfo.Country
		}

		err = global.GVA_DB.Model(&user).Updates(updates).Error
		if err != nil {
			return nil, err
		}

		return &user, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 用户不存在，创建新用户
	user = miniprogram.WxUser{
		Openid:       openid,
		Unionid:      unionid,
		PrivacyLevel: 1, // 默认低隐私级别
		Status:       1, // 默认正常状态
		LastLoginAt:  time.Now(),
	}

	// 如果有用户信息，填充用户资料
	if userInfo != nil {
		user.Nickname = userInfo.Nickname
		user.AvatarUrl = userInfo.AvatarUrl
		user.Gender = userInfo.Gender
		user.City = userInfo.City
		user.Province = userInfo.Province
		user.Country = userInfo.Country
	}

	err = global.GVA_DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// 为新用户创建戒色记录和设置
	err = s.CreateAbstinenceRecord(user.ID)
	if err != nil {
		global.GVA_LOG.Error("创建戒色记录失败", zap.Error(err))
		// 不影响登录流程，只记录错误
	}

	err = s.CreateUserSettings(user.ID)
	if err != nil {
		global.GVA_LOG.Error("创建用户设置失败", zap.Error(err))
		// 不影响登录流程，只记录错误
	}

	return &user, nil
}

// CreateAbstinenceRecord 为新用户创建戒色记录
func (s *UserService) CreateAbstinenceRecord(userID uint) error {
	record := miniprogram.AbstinenceRecord{
		UserID:        userID,
		StartDate:     time.Now(),
		CurrentStreak: 0,
		LongestStreak: 0,
		TotalDays:     0,
		SuccessRate:   0.00,
		Level:         1,
		Experience:    0,
		Status:        1, // 进行中
	}

	return global.GVA_DB.Create(&record).Error
}

// CreateUserSettings 为新用户创建默认设置
func (s *UserService) CreateUserSettings(userID uint) error {
	settings := miniprogram.UserSettings{
		UserID:             userID,
		CheckinReminder:    true,
		CommunityReply:     true,
		AchievementUnlock:  true,
		WeeklyReport:       true,
		EmergencyAlert:     true,
		LearningReminder:   true,
		ShowProfile:        true,
		ShowStats:          true,
		ShowAchievements:   true,
		AllowFriendRequest: true,
		ShowOnlineStatus:   true,
		ExportCount:        0,
	}

	return global.GVA_DB.Create(&settings).Error
}

// GetUserByID 根据ID获取用户信息
func (s *UserService) GetUserByID(userID uint) (*miniprogram.WxUser, error) {
	var user miniprogram.WxUser
	err := global.GVA_DB.Where("id = ? AND status = 1", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(userID uint, updates map[string]interface{}) error {
	return global.GVA_DB.Model(&miniprogram.WxUser{}).Where("id = ?", userID).Updates(updates).Error
}

// UpdatePrivacyLevel 更新用户隐私级别
func (s *UserService) UpdatePrivacyLevel(userID uint, level int) error {
	if level < 1 || level > 3 {
		return errors.New("隐私级别必须在1-3之间")
	}
	return global.GVA_DB.Model(&miniprogram.WxUser{}).Where("id = ?", userID).Update("privacy_level", level).Error
}

// GetUserProfile 获取用户详细资料（包括戒色记录）
func (s *UserService) GetUserProfile(userID uint) (*miniprogram.WxUser, *miniprogram.AbstinenceRecord, error) {
	// 获取用户信息
	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, nil, err
	}

	// 获取戒色记录
	var record miniprogram.AbstinenceRecord
	err = global.GVA_DB.Where("user_id = ? AND status = 1", userID).First(&record).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有戒色记录，创建一个
		err = s.CreateAbstinenceRecord(userID)
		if err != nil {
			return user, nil, err
		}
		err = global.GVA_DB.Where("user_id = ?", userID).First(&record).Error
		if err != nil {
			return user, nil, err
		}
	}

	return user, &record, nil
}

// GetProfileStats 获取个人中心统计数据
func (s *UserService) GetProfileStats(userID uint) (*miniprogramRes.ProfileStatsResponse, error) {
	// 获取用户基本信息和戒色记录
	user, record, err := s.GetUserProfile(userID)
	if err != nil {
		return nil, err
	}

	// 获取用户设置
	settings, err := s.GetUserSettings(userID)
	if err != nil {
		return nil, err
	}

	// 计算加入天数
	joinDays := int(time.Since(user.CreatedAt).Hours() / 24)

	// 获取成就数量
	var achievementCount int64
	global.GVA_DB.Model(&miniprogram.UserAchievement{}).
		Where("user_id = ?", userID).Count(&achievementCount)

	// 获取帮助他人次数（社区回复数量）
	var helpCount int64
	global.GVA_DB.Model(&miniprogram.CommunityComment{}).
		Where("user_id = ?", userID).Count(&helpCount)

	// 获取等级称号
	levelTitle := s.GetLevelTitle(record.Level)

	// 构建基础统计
	basicStats := miniprogramRes.BasicStats{
		CurrentStreak:    record.CurrentStreak,
		LongestStreak:    record.LongestStreak,
		Experience:       record.Experience,
		Level:            record.Level,
		LevelTitle:       levelTitle,
		JoinDays:         joinDays,
		AchievementCount: int(achievementCount),
		HelpCount:        int(helpCount),
		SuccessRate:      record.SuccessRate,
	}

	// 获取最近成就
	recentAchievements, err := s.GetRecentAchievements(userID, 3)
	if err != nil {
		return nil, err
	}

	// 构建用户设置
	userSettings := miniprogramRes.UserSettings{
		NotificationSettings: miniprogramRes.NotificationSettings{
			CheckinReminder:   settings.CheckinReminder,
			CommunityReply:    settings.CommunityReply,
			AchievementUnlock: settings.AchievementUnlock,
			WeeklyReport:      settings.WeeklyReport,
			EmergencyAlert:    settings.EmergencyAlert,
			LearningReminder:  settings.LearningReminder,
		},
		PrivacySettings: miniprogramRes.PrivacySettings{
			ShowProfile:        settings.ShowProfile,
			ShowStats:          settings.ShowStats,
			ShowAchievements:   settings.ShowAchievements,
			AllowFriendRequest: settings.AllowFriendRequest,
			ShowOnlineStatus:   settings.ShowOnlineStatus,
		},
	}

	return &miniprogramRes.ProfileStatsResponse{
		BasicStats:         basicStats,
		RecentAchievements: recentAchievements,
		UserSettings:       userSettings,
	}, nil
}

// GetLevelTitle 根据等级获取称号
func (s *UserService) GetLevelTitle(level int) string {
	titleMap := map[int]string{
		1:  "新手上路",
		2:  "初有成效",
		3:  "小有成就",
		4:  "渐入佳境",
		5:  "稳步前进",
		6:  "坚定不移",
		7:  "意志坚强",
		8:  "稳定期",
		9:  "自律达人",
		10: "戒色大师",
	}

	if title, exists := titleMap[level]; exists {
		return title
	}
	return "戒色专家"
}

// GetRecentAchievements 获取最近成就
func (s *UserService) GetRecentAchievements(userID uint, limit int) ([]miniprogramRes.RecentAchievement, error) {
	var userAchievements []miniprogram.UserAchievement
	err := global.GVA_DB.Where("user_id = ?", userID).
		Order("unlocked_at DESC").
		Limit(limit).
		Find(&userAchievements).Error
	if err != nil {
		return nil, err
	}

	var recentAchievements []miniprogramRes.RecentAchievement
	for _, ua := range userAchievements {
		// 获取成就详情
		var achievement miniprogram.Achievement
		err := global.GVA_DB.Where("id = ?", ua.AchievementID).First(&achievement).Error
		if err != nil {
			continue
		}

		// 计算几天前获得
		daysAgo := int(time.Since(ua.UnlockedAt).Hours() / 24)

		// 转换稀有度为字符串
		rarityMap := map[int]string{
			1: "普通",
			2: "稀有",
			3: "史诗",
			4: "传说",
		}
		rarityStr := rarityMap[achievement.Rarity]
		if rarityStr == "" {
			rarityStr = "普通"
		}

		recentAchievement := miniprogramRes.RecentAchievement{
			ID:          achievement.ID,
			Title:       achievement.Name,
			Description: achievement.Description,
			Icon:        achievement.IconUrl,
			Rarity:      rarityStr,
			UnlockedAt:  ua.UnlockedAt,
			DaysAgo:     daysAgo,
		}

		recentAchievements = append(recentAchievements, recentAchievement)
	}

	return recentAchievements, nil
}

// GetUserSettings 获取用户设置
func (s *UserService) GetUserSettings(userID uint) (*miniprogram.UserSettings, error) {
	var settings miniprogram.UserSettings
	err := global.GVA_DB.Where("user_id = ?", userID).First(&settings).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有设置记录，创建默认设置
			err = s.CreateUserSettings(userID)
			if err != nil {
				return nil, err
			}
			err = global.GVA_DB.Where("user_id = ?", userID).First(&settings).Error
		}
		if err != nil {
			return nil, err
		}
	}
	return &settings, nil
}

// UpdateNotificationSettings 更新通知设置
func (s *UserService) UpdateNotificationSettings(userID uint, req *miniprogramReq.NotificationSettingsRequest) error {
	updates := map[string]interface{}{
		"checkin_reminder":   req.CheckinReminder,
		"community_reply":    req.CommunityReply,
		"achievement_unlock": req.AchievementUnlock,
		"weekly_report":      req.WeeklyReport,
		"emergency_alert":    req.EmergencyAlert,
		"learning_reminder":  req.LearningReminder,
	}

	return global.GVA_DB.Model(&miniprogram.UserSettings{}).
		Where("user_id = ?", userID).
		Updates(updates).Error
}

// UpdatePrivacySettings 更新隐私设置
func (s *UserService) UpdatePrivacySettings(userID uint, req *miniprogramReq.PrivacySettingsRequest) error {
	updates := map[string]interface{}{
		"show_profile":         req.ShowProfile,
		"show_stats":           req.ShowStats,
		"show_achievements":    req.ShowAchievements,
		"allow_friend_request": req.AllowFriendRequest,
		"show_online_status":   req.ShowOnlineStatus,
	}

	return global.GVA_DB.Model(&miniprogram.UserSettings{}).
		Where("user_id = ?", userID).
		Updates(updates).Error
}

// CreateDataExport 创建数据导出任务
func (s *UserService) CreateDataExport(userID uint, req *miniprogramReq.DataExportRequest) (*miniprogramRes.DataExportResponse, error) {
	// 检查导出频率限制（每天最多5次）
	var todayCount int64
	today := time.Now().Format("2006-01-02")
	global.GVA_DB.Model(&miniprogram.DataExport{}).
		Where("user_id = ? AND DATE(created_at) = ?", userID, today).
		Count(&todayCount)

	if todayCount >= 5 {
		return nil, errors.New("今日导出次数已达上限（5次）")
	}

	// 序列化数据类型
	dataTypesJSON, err := json.Marshal(req.DataTypes)
	if err != nil {
		return nil, err
	}

	// 生成文件名
	timestamp := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("nofap_data_%s_%s.%s", strconv.Itoa(int(userID)), timestamp, req.ExportType)

	// 设置过期时间（7天后）
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	// 创建导出记录
	export := miniprogram.DataExport{
		UserID:     userID,
		ExportType: req.ExportType,
		DataTypes:  string(dataTypesJSON),
		DateRange:  req.DateRange,
		FileName:   fileName,
		Status:     1, // 处理中
		ExpiresAt:  &expiresAt,
	}

	err = global.GVA_DB.Create(&export).Error
	if err != nil {
		return nil, err
	}

	// 异步处理数据导出
	go s.ProcessDataExport(&export, req)

	// 更新用户设置中的导出计数
	global.GVA_DB.Model(&miniprogram.UserSettings{}).
		Where("user_id = ?", userID).
		UpdateColumn("export_count", gorm.Expr("export_count + ?", 1))

	return &miniprogramRes.DataExportResponse{
		DownloadURL: "", // 处理完成后会更新
		FileName:    fileName,
		FileSize:    0,
		ExpiresAt:   expiresAt,
	}, nil
}

// ProcessDataExport 处理数据导出（异步）
func (s *UserService) ProcessDataExport(export *miniprogram.DataExport, req *miniprogramReq.DataExportRequest) {
	// 这里实现具体的数据导出逻辑
	// 根据 req.DataTypes 和 req.DateRange 导出相应的数据
	// 生成文件并上传到文件服务器
	// 更新 export 记录的状态和下载链接

	// 模拟处理时间
	time.Sleep(5 * time.Second)

	// 生成文件路径
	filePath := filepath.Join("exports", export.FileName)
	downloadURL := fmt.Sprintf("/api/v1/miniprogram/profile/download/%d", export.ID)

	// 更新导出记录
	updates := map[string]interface{}{
		"file_path":    filePath,
		"file_size":    1024, // 模拟文件大小
		"status":       2,    // 完成
		"download_url": downloadURL,
	}

	global.GVA_DB.Model(&miniprogram.DataExport{}).
		Where("id = ?", export.ID).
		Updates(updates)
}

// ValidateUserLogin 验证用户登录
func (s *UserService) ValidateUserLogin(phone, password string) (*miniprogram.WxUser, error) {
	var user miniprogram.WxUser

	// 根据手机号查找用户
	err := global.GVA_DB.Where("phone = ? AND status = 1", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// 这里应该验证密码哈希，暂时简单比较
	// TODO: 实现密码哈希验证
	if user.Password != password {
		return nil, errors.New("密码错误")
	}

	// 更新最后登录时间
	global.GVA_DB.Model(&user).Update("last_login_at", time.Now())

	return &user, nil
}

// CreateUser 创建新用户
func (s *UserService) CreateUser(req *miniprogramReq.RegisterRequest) (*miniprogram.WxUser, error) {
	// 检查手机号是否已存在
	var existingUser miniprogram.WxUser
	err := global.GVA_DB.Where("phone = ?", req.Phone).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("手机号已被注册")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 创建新用户
	user := miniprogram.WxUser{
		Phone:        req.Phone,
		Password:     req.Password, // TODO: 应该进行密码哈希
		Nickname:     req.Nickname,
		PrivacyLevel: 1, // 默认低隐私级别
		Status:       1, // 默认正常状态
		LastLoginAt:  time.Now(),
	}

	err = global.GVA_DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// 为新用户创建戒色记录和设置
	err = s.CreateAbstinenceRecord(user.ID)
	if err != nil {
		global.GVA_LOG.Error("创建戒色记录失败", zap.Error(err))
	}

	err = s.CreateUserSettings(user.ID)
	if err != nil {
		global.GVA_LOG.Error("创建用户设置失败", zap.Error(err))
	}

	return &user, nil
}
