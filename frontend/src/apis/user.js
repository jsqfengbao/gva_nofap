/**
 * 用户相关 API
 */
import { request } from '@/utils/auth'
import { buildApiUrl } from '@/config/env.js'

export const userApi = {
  // 获取用户资料
  getProfile: () => {
    return request({
      url: buildApiUrl('/user/profile'),
      method: 'GET'
    })
  },

  // 更新用户信息
  updateUserInfo: (data) => {
    return request({
      url: buildApiUrl('/user/info'),
      method: 'PUT',
      data
    })
  },

  // 保存微信头像
  saveWxAvatar: (data) => {
    return request({
      url: buildApiUrl('/user/save-wx-avatar'),
      method: 'POST',
      data
    })
  },

  // 获取用户统计数据
  getStats: () => {
    return request({
      url: buildApiUrl('/user/stats'),
      method: 'GET'
    })
  },

  // 更新隐私设置
  updatePrivacySettings: (data) => {
    return request({
      url: buildApiUrl('/user/privacy-settings'),
      method: 'PUT',
      data
    })
  },

  // 更新通知设置
  updateNotificationSettings: (data) => {
    return request({
      url: buildApiUrl('/user/notification-settings'),
      method: 'PUT',
      data
    })
  },

  // 创建数据导出
  createDataExport: (data) => {
    return request({
      url: buildApiUrl('/user/export'),
      method: 'POST',
      data
    })
  },

  // 获取用户设置
  getUserSettings: () => {
    return request({
      url: buildApiUrl('/user/settings'),
      method: 'GET'
    })
  }
} 