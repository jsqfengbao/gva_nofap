package api

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	miniprogramReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model/request"
	miniprogramRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/nofap/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

// GetUserProfile 获取用户详细资料
// @Tags      MiniProgram
// @Summary   获取用户详细资料
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=miniprogramRes.UserProfileResponse,msg=string}  "获取成功"
// @Router    /miniprogram/user/profile [get]
func (u *UserApi) GetUserProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	user, record, err := userService.GetUserProfile(userID)
	if err != nil {
		global.GVA_LOG.Error("获取用户资料失败", zap.Error(err))
		response.FailWithMessage("获取用户资料失败", c)
		return
	}

	profileRes := miniprogramRes.UserProfileResponse{
		User:             user,
		AbstinenceRecord: record,
	}

	response.OkWithData(profileRes, c)
}

// UpdateUserInfo 更新用户信息
// @Tags      MiniProgram
// @Summary   更新用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      miniprogramReq.UpdateUserInfoRequest                                true  "更新用户信息请求"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /miniprogram/user/info [put]
func (u *UserApi) UpdateUserInfo(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	var updateReq miniprogramReq.UpdateUserInfoRequest
	err := c.ShouldBindJSON(&updateReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 构建更新数据
	updates := make(map[string]interface{})
	if updateReq.Nickname != "" {
		updates["nickname"] = updateReq.Nickname
	}
	if updateReq.AvatarUrl != "" {
		updates["avatar_url"] = updateReq.AvatarUrl
	}
	if updateReq.Gender > 0 {
		updates["gender"] = updateReq.Gender
	}

	// 更新用户信息
	err = userService.UpdateUser(userID, updates)
	if err != nil {
		global.GVA_LOG.Error("更新用户信息失败", zap.Error(err))
		response.FailWithMessage("更新用户信息失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// UpdatePrivacyLevel 更新隐私级别
func (u *UserApi) UpdatePrivacyLevel(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	var privacyReq miniprogramReq.UpdatePrivacyLevelRequest
	err := c.ShouldBindJSON(&privacyReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.UpdatePrivacyLevel(userID, privacyReq.PrivacyLevel)
	if err != nil {
		global.GVA_LOG.Error("更新隐私级别失败", zap.Error(err))
		response.FailWithMessage("更新隐私级别失败", c)
		return
	}

	response.OkWithMessage("隐私级别更新成功", c)
}

// GetProfileStats 获取个人中心统计数据
func (u *UserApi) GetProfileStats(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	stats, err := userService.GetProfileStats(userID)
	if err != nil {
		global.GVA_LOG.Error("获取个人中心统计数据失败", zap.Error(err))
		response.FailWithMessage("获取统计数据失败", c)
		return
	}

	response.OkWithData(stats, c)
}

// UpdateNotificationSettings 更新通知设置
func (u *UserApi) UpdateNotificationSettings(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	var req miniprogramReq.NotificationSettingsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.UpdateNotificationSettings(userID, &req)
	if err != nil {
		global.GVA_LOG.Error("更新通知设置失败", zap.Error(err))
		response.FailWithMessage("更新通知设置失败", c)
		return
	}

	response.OkWithMessage("通知设置更新成功", c)
}

// UpdatePrivacySettings 更新隐私设置
func (u *UserApi) UpdatePrivacySettings(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	var req miniprogramReq.PrivacySettingsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.UpdatePrivacySettings(userID, &req)
	if err != nil {
		global.GVA_LOG.Error("更新隐私设置失败", zap.Error(err))
		response.FailWithMessage("更新隐私设置失败", c)
		return
	}

	response.OkWithMessage("隐私设置更新成功", c)
}

// CreateDataExport 创建数据导出任务
func (u *UserApi) CreateDataExport(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	var req miniprogramReq.DataExportRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	exportRes, err := userService.CreateDataExport(userID, &req)
	if err != nil {
		global.GVA_LOG.Error("创建数据导出任务失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(exportRes, c)
}

// GetDataExportStatus 获取数据导出状态
func (u *UserApi) GetDataExportStatus(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	exportIDStr := c.Param("exportId")
	exportID, err := strconv.ParseUint(exportIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("导出ID格式错误", c)
		return
	}

	var export struct {
		ID          uint   `json:"id"`
		Status      int    `json:"status"`
		FileName    string `json:"fileName"`
		FileSize    int64  `json:"fileSize"`
		DownloadURL string `json:"downloadUrl"`
		ExpiresAt   string `json:"expiresAt"`
	}

	err = global.GVA_DB.Table("nofap_data_exports").
		Select("id, status, file_name, file_size, download_url, expires_at").
		Where("id = ? AND user_id = ?", uint(exportID), userID).
		First(&export).Error

	if err != nil {
		global.GVA_LOG.Error("获取导出状态失败", zap.Error(err))
		response.FailWithMessage("获取导出状态失败", c)
		return
	}

	response.OkWithData(export, c)
}

// DownloadDataExport 下载导出文件
func (u *UserApi) DownloadDataExport(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	exportIDStr := c.Param("exportId")
	exportID, err := strconv.ParseUint(exportIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("导出ID格式错误", c)
		return
	}

	var export struct {
		FilePath string `json:"filePath"`
		FileName string `json:"fileName"`
		Status   int    `json:"status"`
	}

	err = global.GVA_DB.Table("nofap_data_exports").
		Select("file_path, file_name, status").
		Where("id = ? AND user_id = ? AND status = 2", uint(exportID), userID).
		First(&export).Error

	if err != nil {
		global.GVA_LOG.Error("获取导出文件失败", zap.Error(err))
		response.FailWithMessage("文件不存在或未完成", c)
		return
	}

	// 这里应该实现实际的文件下载逻辑
	// 暂时返回文件信息
	response.OkWithData(gin.H{
		"message":  "文件下载功能暂未实现",
		"filePath": export.FilePath,
		"fileName": export.FileName,
	}, c)
}

// GetUserSettings 获取用户设置
func (u *UserApi) GetUserSettings(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	settings, err := userService.GetUserSettings(userID)
	if err != nil {
		global.GVA_LOG.Error("获取用户设置失败", zap.Error(err))
		response.FailWithMessage("获取用户设置失败", c)
		return
	}

	response.OkWithData(settings, c)
}

// UploadAvatar 上传用户头像
// @Tags      MiniProgram
// @Summary   上传用户头像
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                   true  "头像文件"
// @Success   200   {object}  response.Response{data=gin.H,msg=string}  "上传成功"
// @Router    /miniprogram/user/upload-avatar [post]
func (u *UserApi) UploadAvatar(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage("获取文件失败", c)
		return
	}
	defer file.Close()

	// 验证文件类型
	allowedTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/gif"}
	contentType := header.Header.Get("Content-Type")
	isAllowed := false
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		response.FailWithMessage("只支持 JPG、PNG、GIF 格式的图片", c)
		return
	}

	// 验证文件大小 (限制为 2MB)
	if header.Size > 2*1024*1024 {
		response.FailWithMessage("文件大小不能超过 2MB", c)
		return
	}

	// 这里应该调用文件上传服务
	// 暂时返回一个示例URL，实际应该上传到云存储或本地存储
	avatarURL := fmt.Sprintf("/uploads/avatars/user_%d_%d%s",
		userID,
		time.Now().Unix(),
		filepath.Ext(header.Filename))

	// 更新用户头像URL
	err = userService.UpdateUser(userID, map[string]interface{}{
		"avatar_url": avatarURL,
	})
	if err != nil {
		global.GVA_LOG.Error("更新用户头像失败", zap.Error(err))
		response.FailWithMessage("更新头像失败", c)
		return
	}

	response.OkWithData(gin.H{
		"url":     avatarURL,
		"message": "头像上传成功",
	}, c)
}

// SaveWxTempAvatar 保存微信临时头像
// @Tags      MiniProgram
// @Summary   保存微信临时头像到服务器
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      gin.H{tempUrl=string}                                true  "临时头像URL"
// @Success   200   {object}  response.Response{data=gin.H,msg=string}  "保存成功"
// @Router    /miniprogram/user/save-wx-avatar [post]
func (u *UserApi) SaveWxTempAvatar(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	var req struct {
		TempUrl string `json:"tempUrl" binding:"required"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该下载微信临时头像并保存到服务器
	// 暂时直接使用临时URL，实际应该下载并保存
	avatarURL := fmt.Sprintf("/uploads/avatars/wx_temp_%d_%d.jpg",
		userID,
		time.Now().Unix())

	// 更新用户头像URL
	err = userService.UpdateUser(userID, map[string]interface{}{
		"avatar_url": avatarURL,
	})
	if err != nil {
		global.GVA_LOG.Error("更新用户头像失败", zap.Error(err))
		response.FailWithMessage("更新头像失败", c)
		return
	}

	response.OkWithData(gin.H{
		"url":     avatarURL,
		"message": "头像保存成功",
	}, c)
}
