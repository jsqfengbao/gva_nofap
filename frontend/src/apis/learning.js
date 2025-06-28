/**
 * 学习相关 API
 */
import { request } from '@/utils/auth'
import { buildApiUrl } from '@/config/env.js'

export const learningApi = {
  // 获取学习统计
  getStats: () => {
    return request({
      url: buildApiUrl('/learning/stats'),      
      method: 'GET'
    })
  },

  // 获取学习内容列表
  getContents: (params = {}) => {
    return request({
      url: buildApiUrl('/learning/contents'),
      method: 'GET',
      data: params
    })
  },

  // 获取学习内容列表 (兼容旧接口)
  getList: (params = {}) => {
    return learningApi.getContents(params)
  },

  // 获取学习内容详情
  getDetail: (id) => {
    return request({
      url: buildApiUrl(`/learning/${id}`),
      method: 'GET'
    })
  },

  // 开始学习记录
  startLearning: (contentId) => {
    return request({
      url: buildApiUrl('/learning/start'),
      method: 'POST',
      data: { contentId }
    })
  },

  // 记录学习进度
  recordProgress: (id, data) => {
    return request({
      url: buildApiUrl(`/learning/${id}/progress`),
      method: 'POST',
      data
    })
  },

  // 完成学习
  completeLearning: (contentId, data = {}) => {
    return request({
      url: buildApiUrl('/learning/complete'),
      method: 'POST',
      data: { contentId, ...data }
    })
  },

  // 点赞内容
  likeContent: (contentId) => {
    return request({
      url: buildApiUrl(`/learning/${contentId}/like`),
      method: 'POST'
    })
  },

  // 收藏内容
  collectContent: (contentId) => {
    return request({
      url: buildApiUrl(`/learning/${contentId}/collect`),
      method: 'POST'
    })
  },

  // 搜索内容
  searchContents: (params = {}) => {
    return request({
      url: buildApiUrl('/learning/search'),
      method: 'GET',
      data: params
    })
  },

  // 获取推荐内容
  getRecommendations: (params = {}) => {
    return request({
      url: buildApiUrl('/learning/recommendations'),
      method: 'GET',
      data: params
    })
  },

  // 获取分类统计
  getCategoryStats: () => {
    return request({
      url: buildApiUrl('/learning/category-stats'),
      method: 'GET'
    })
  }
} 