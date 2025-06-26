import service from '@/utils/request'

// 微信小程序认证相关
export const wxLogin = (data) => {
  return service({
    url: '/api/v1/miniprogram/auth/login',
    method: 'post',
    data
  })
}

export const refreshToken = (data) => {
  return service({
    url: '/api/v1/miniprogram/auth/refresh',
    method: 'post',
    data
  })
}

export const getUserProfile = () => {
  return service({
    url: '/api/v1/miniprogram/auth/profile',
    method: 'get'
  })
}

export const updateUserInfo = (data) => {
  return service({
    url: '/api/v1/miniprogram/auth/profile',
    method: 'put',
    data
  })
}

export const updatePrivacyLevel = (data) => {
  return service({
    url: '/api/v1/miniprogram/auth/privacy',
    method: 'put',
    data
  })
}

// 用户管理（管理端使用）
export const getUserList = (params) => {
  return service({
    url: '/api/v1/miniprogram/user/list',
    method: 'get',
    params
  })
}

export const getUserDetail = (id) => {
  return service({
    url: `/api/v1/miniprogram/user/${id}`,
    method: 'get'
  })
}

export const updateUserStatus = (data) => {
  return service({
    url: '/api/v1/miniprogram/user/status',
    method: 'put',
    data
  })
}

// 评估管理
export const getAssessmentList = (params) => {
  return service({
    url: '/api/v1/miniprogram/assessment/list',
    method: 'get',
    params
  })
}

export const createAssessment = (data) => {
  return service({
    url: '/api/v1/miniprogram/assessment',
    method: 'post',
    data
  })
}

export const updateAssessment = (data) => {
  return service({
    url: '/api/v1/miniprogram/assessment',
    method: 'put',
    data
  })
}

export const deleteAssessment = (id) => {
  return service({
    url: `/api/v1/miniprogram/assessment/${id}`,
    method: 'delete'
  })
}

// 社区管理
export const getCommunityPostList = (params) => {
  return service({
    url: '/api/v1/miniprogram/community/posts',
    method: 'get',
    params
  })
}

export const auditCommunityPost = (data) => {
  return service({
    url: '/api/v1/miniprogram/community/audit',
    method: 'put',
    data
  })
}

export const deleteCommunityPost = (id) => {
  return service({
    url: `/api/v1/miniprogram/community/posts/${id}`,
    method: 'delete'
  })
}

// 学习内容管理
export const getLearningContentList = (params) => {
  return service({
    url: '/api/v1/miniprogram/learning/content/list',
    method: 'get',
    params
  })
}

export const createLearningContent = (data) => {
  return service({
    url: '/api/v1/miniprogram/learning/content',
    method: 'post',
    data
  })
}

export const updateLearningContent = (data) => {
  return service({
    url: '/api/v1/miniprogram/learning/content',
    method: 'put',
    data
  })
}

export const deleteLearningContent = (id) => {
  return service({
    url: `/api/v1/miniprogram/learning/content/${id}`,
    method: 'delete'
  })
}

// 成就管理
export const getAchievementList = (params) => {
  return service({
    url: '/api/v1/miniprogram/achievement/list',
    method: 'get',
    params
  })
}

export const createAchievement = (data) => {
  return service({
    url: '/api/v1/miniprogram/achievement',
    method: 'post',
    data
  })
}

export const updateAchievement = (data) => {
  return service({
    url: '/api/v1/miniprogram/achievement',
    method: 'put',
    data
  })
}

export const deleteAchievement = (id) => {
  return service({
    url: `/api/v1/miniprogram/achievement/${id}`,
    method: 'delete'
  })
}

// 数据统计
export const getUserStatistics = () => {
  return service({
    url: '/api/v1/miniprogram/statistics/users',
    method: 'get'
  })
}

export const getActivityStatistics = () => {
  return service({
    url: '/api/v1/miniprogram/statistics/activity',
    method: 'get'
  })
}

export const getContentStatistics = () => {
  return service({
    url: '/api/v1/miniprogram/statistics/content',
    method: 'get'
  })
} 