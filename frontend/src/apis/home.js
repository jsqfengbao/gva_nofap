/**
 * 首页相关 API
 */
import { checkinApi } from './checkin'
import { achievementApi } from './achievement'

export const homeApi = {
  // 获取首页所有数据
  getHomeData: async () => {
    try {
      const [statsRes, todayRes, gameRes] = await Promise.all([
        checkinApi.getStatistics(),
        checkinApi.getTodayStatus(),
        achievementApi.getGameStats()
      ])
      
      return {
        code: 0,
        data: {
          userStats: statsRes.data || {},
          todayStatus: todayRes.data || {},
          gameStats: gameRes.data || {}
        }
      }
    } catch (error) {
      console.error('获取首页数据失败:', error)
      return {
        code: -1,
        message: '获取首页数据失败',
        error
      }
    }
  },

  // 获取打卡统计数据
  getCheckinStats: () => {
    return checkinApi.getStatistics()
  },

  // 获取今日打卡状态
  getTodayCheckinStatus: () => {
    return checkinApi.getTodayStatus()
  },

  // 获取游戏化统计数据
  getGameStats: () => {
    return achievementApi.getGameStats()
  }
} 