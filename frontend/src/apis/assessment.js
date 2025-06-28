/**
 * 评估相关 API
 */
import { request } from '@/utils/auth'
import { buildApiUrl } from '@/config/env.js'

export const assessmentApi = {
  // 获取评估问题
  getQuestions: (type) => {
    return request({
      url: buildApiUrl(`/assessment/questions/${type}`),
      method: 'GET'
    })
  },

  // 提交评估结果
  submitResult: (data) => {
    return request({
      url: buildApiUrl('/assessment/submit'),
      method: 'POST',
      data
    })
  },

  // 获取评估历史
  getHistory: (params = {}) => {
    return request({
      url: buildApiUrl('/assessment/history'),
      method: 'GET',
      data: params
    })
  }
} 