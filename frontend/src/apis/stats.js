/**
 * 统计相关 API
 */
import { request } from '@/utils/auth'
import { buildApiUrl } from '@/config/env.js'

export const statsApi = {
  // 获取总体统计
  getOverall: () => {
    return request({
      url: buildApiUrl('/stats/overall'),
      method: 'GET'
    })
  },

  // 获取趋势数据
  getTrends: (params = {}) => {
    return request({
      url: buildApiUrl('/stats/trends'),
      method: 'GET',
      data: params
    })
  }
} 