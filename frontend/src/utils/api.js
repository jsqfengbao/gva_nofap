/**
 * API接口调用模块
 */
import { request } from './auth'
import { buildApiUrl } from '../config/env.js'

// 用户相关API
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

// 成就相关API
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
  }
}

// 打卡相关API
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
  }
}

// 社区相关API
export const communityApi = {
  // 获取帖子列表
  getPosts: (params = {}) => {
    return request({
      url: buildApiUrl('/community/posts'),
      method: 'GET',
      data: params
    })
  },

  // 获取帖子详情
  getPostDetail: (id) => {
    return request({
      url: buildApiUrl(`/community/posts/${id}`),
      method: 'GET'
    })
  },

  // 发布帖子
  createPost: (data) => {
    return request({
      url: buildApiUrl('/community/posts'),
      method: 'POST',
      data
    })
  },

  // 回复帖子
  replyPost: (postId, data) => {
    return request({
      url: buildApiUrl(`/community/posts/${postId}/replies`),
      method: 'POST',
      data
    })
  }
}

// 学习相关API
export const learningApi = {
  // 获取学习内容列表
  getList: (params = {}) => {
    return request({
      url: buildApiUrl('/learning/list'),
      method: 'GET',
      data: params
    })
  },

  // 获取学习内容详情
  getDetail: (id) => {
    return request({
      url: buildApiUrl(`/learning/${id}`),
      method: 'GET'
    })
  },

  // 记录学习进度
  recordProgress: (id, data) => {
    return request({
      url: buildApiUrl(`/learning/${id}/progress`),
      method: 'POST',
      data
    })
  }
}

// 评估相关API
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

// 紧急求助相关API
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

// 统计相关API
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

// 导出所有API
export default {
  user: userApi,
  achievement: achievementApi,
  checkin: checkinApi,
  community: communityApi,
  learning: learningApi,
  assessment: assessmentApi,
  emergency: emergencyApi,
  stats: statsApi
} 