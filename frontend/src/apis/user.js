/**
 * 用户相关API
 */
import { authGet, authPost, authPut, uploadFile } from '@/utils/request.js'

/**
 * 获取用户资料
 */
export function getProfile() {
  return authGet('/user/profile')
}

/**
 * 更新用户信息
 * @param {Object} data 用户信息
 * @param {string} data.nickname 昵称
 * @param {string} data.avatarUrl 头像URL
 */
export function updateUserInfo(data) {
  return authPut('/user/info', data)
}

/**
 * 保存微信头像
 * @param {Object} data 头像数据
 * @param {string} data.tempUrl 临时头像URL
 */
export function saveWxAvatar(data) {
  return authPost('/user/save-wx-avatar', data)
}

/**
 * 获取用户统计数据
 */
export function getStats() {
  return authGet('/user/stats')
}

/**
 * 更新隐私设置
 * @param {Object} data 隐私设置
 */
export function updatePrivacySettings(data) {
  return authPut('/user/privacy-settings', data)
}

/**
 * 更新通知设置
 * @param {Object} data 通知设置
 */
export function updateNotificationSettings(data) {
  return authPut('/user/notification-settings', data)
}

/**
 * 创建数据导出
 * @param {Object} data 导出配置
 */
export function createDataExport(data) {
  return authPost('/user/export', data)
}

/**
 * 获取用户设置
 */
export function getUserSettings() {
  return authGet('/user/settings')
}

/**
 * 上传头像文件
 * @param {string} filePath 文件路径
 */
export function uploadAvatar(filePath) {
  return uploadFile('/user/upload-avatar', filePath)
}

// 导出用户API对象
const userApi = {
  getProfile,
  updateUserInfo,
  saveWxAvatar,
  getStats,
  updatePrivacySettings,
  updateNotificationSettings,
  createDataExport,
  getUserSettings,
  uploadAvatar
}

export default userApi 