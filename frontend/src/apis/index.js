/**
 * 统一API入口
 * 整合各模块的API接口
 */

// 导入各模块API
export * from './auth.js'
export * from './user.js'
export * from './checkin.js'
export * from './community.js'
export * from './learning.js'
export * from './achievement.js'
export * from './assessment.js'
export * from './emergency.js'
export * from './home.js'

// 导入API模块对象
import authApi from './auth.js'
import userApi from './user.js'
import checkinApi from './checkin.js'
import communityApi from './community.js'
import learningApi from './learning.js'
import achievementApi from './achievement.js'
import assessmentApi from './assessment.js'
import emergencyApi from './emergency.js'
import homeApi from './home.js'

// 统一导出
export {
  authApi,
  userApi,
  checkinApi,
  communityApi,
  learningApi,
  achievementApi,
  assessmentApi,
  emergencyApi,
  homeApi
}

// 默认导出所有API
export default {
  auth: authApi,
  user: userApi,
  checkin: checkinApi,
  community: communityApi,
  learning: learningApi,
  achievement: achievementApi,
  assessment: assessmentApi,
  emergency: emergencyApi,
  home: homeApi
} 