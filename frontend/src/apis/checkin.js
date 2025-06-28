/**
 * 打卡相关 API
 */
import { request } from '@/utils/auth'
import { buildApiUrl } from '@/config/env.js'

export const checkinApi = {
  // 今日打卡
  checkin: (data) => {
    return request({
      url: buildApiUrl('/checkin'),
      method: 'POST',
      data
    })
  },

  // 获取打卡历史
  getHistory: (params = {}) => {
    return request({
      url: buildApiUrl('/checkin/history'),
      method: 'GET',
      data: params
    })
  },

  // 获取打卡统计
  getStats: (params = {}) => {
    return request({
      url: buildApiUrl('/checkin/stats'),
      method: 'GET',
      data: params
    })
  },

  // 获取打卡统计数据（首页使用）
  getStatistics: () => {
    return request({
      url: buildApiUrl('/checkin/statistics'),
      method: 'GET'
    })
  },

  // 获取今日打卡状态（首页使用）
  getTodayStatus: () => {
    return request({
      url: buildApiUrl('/checkin/today'),
      method: 'GET'
    })
  }
} 