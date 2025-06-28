/**
 * 首页相关API
 */
import checkinApi from './checkin.js'
import achievementApi from './achievement.js'

/**
 * 获取首页所有数据
 * 并发请求多个接口，提高加载效率
 */
export async function getHomeData() {
  try {
    const [statsRes, todayRes, gameRes] = await Promise.all([
      checkinApi.getStatistics(),
      checkinApi.getTodayStatus(),
      achievementApi.getGameStats()
    ])
    
    return {
      code: 0,
      data: {
        userStats: statsRes.data?.data || {},
        todayStatus: todayRes.data?.data || {},
        gameStats: gameRes.data?.data || {}
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
}

/**
 * 获取打卡统计数据
 */
export function getCheckinStats() {
  return checkinApi.getStatistics()
}

/**
 * 获取今日打卡状态
 */
export function getTodayCheckinStatus() {
  return checkinApi.getTodayStatus()
}

/**
 * 获取游戏化统计数据
 */
export function getGameStats() {
  return achievementApi.getGameStats()
}

// 导出首页API对象
const homeApi = {
  getHomeData,
  getCheckinStats,
  getTodayCheckinStatus,
  getGameStats
}

export default homeApi 