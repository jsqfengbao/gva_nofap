/**
 * 成就相关 API
 */
import { request } from '@/utils/auth'
import { buildApiUrl } from '@/config/env.js'

export const achievementApi = {
  // 获取成就列表
  getList: (params = {}) => {
    return request({
      url: buildApiUrl('/achievement/list'),
      method: 'GET',
      data: params
    })
  },

  // 获取用户成就
  getUserAchievements: (params = {}) => {
    return request({
      url: buildApiUrl('/achievement/user'),
      method: 'GET',
      data: params
    })
  },

  // 获取游戏化统计数据（首页使用）
  getGameStats: () => {
    return request({
      url: buildApiUrl('/achievement/game-stats'),
      method: 'GET'
    })
  }
} 