/**
 * API 统一入口文件
 * 导出所有 API 模块
 */

import { userApi } from './user'
import { achievementApi } from './achievement'
import { checkinApi } from './checkin'
import { communityApi } from './community'
import { learningApi } from './learning'
import learningService from './learning-service'
import { assessmentApi } from './assessment'
import { emergencyApi } from './emergency'
import { statsApi } from './stats'
import { homeApi } from './home'

// 统一导出所有 API
export {
  userApi,
  achievementApi,
  checkinApi,
  communityApi,
  learningApi,
  learningService,
  assessmentApi,
  emergencyApi,
  statsApi,
  homeApi
}

// 默认导出对象形式
export default {
  user: userApi,
  achievement: achievementApi,
  checkin: checkinApi,
  community: communityApi,
  learning: learningApi,
  learningService: learningService,
  assessment: assessmentApi,
  emergency: emergencyApi,
  stats: statsApi,
  home: homeApi
} 