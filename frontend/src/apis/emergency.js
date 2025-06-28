/**
 * 紧急求助相关 API
 */
import { request } from '@/utils/auth'
import { buildApiUrl } from '@/config/env.js'

export const emergencyApi = {
  // 获取紧急资源
  getResources: (params = {}) => {
    return request({
      url: buildApiUrl('/emergency/resources'),
      method: 'GET',
      data: params
    })
  },

  // 获取励志文章
  getArticles: (params = {}) => {
    return request({
      url: buildApiUrl('/emergency/articles'),
      method: 'GET',
      data: params
    })
  },

  // 记录求助记录
  recordHelp: (data) => {
    return request({
      url: buildApiUrl('/emergency/help'),
      method: 'POST',
      data
    })
  }
} 