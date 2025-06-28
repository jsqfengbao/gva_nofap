/**
 * 打卡相关API
 */
import { authGet, authPost, guestGet } from '@/utils/request.js'

/**
 * 今日打卡
 * @param {Object} data 打卡数据
 * @param {number} data.moodLevel 心情等级 1-5
 * @param {string} data.notes 打卡备注
 */
export function dailyCheckin(data) {
  return authPost('/checkin/daily', data)
}

/**
 * 获取今日打卡状态
 */
export function getTodayStatus() {
  return authGet('/checkin/today')
}

/**
 * 获取打卡历史
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.month 月份过滤 (YYYY-MM)
 */
export function getHistory(params = {}) {
  return authGet('/checkin/history', params)
}

/**
 * 获取打卡统计数据
 * @param {Object} params 查询参数
 * @param {string} params.period 统计周期 (week|month|year)
 */
export function getStatistics(params = {}) {
  return authGet('/checkin/statistics', params)
}

/**
 * 获取本周进度
 */
export function getWeeklyProgress() {
  return authGet('/checkin/weekly-progress')
}

/**
 * 获取打卡日历数据
 * @param {Object} params 查询参数
 * @param {string} params.year 年份
 * @param {string} params.month 月份
 */
export function getCalendarData(params = {}) {
  return authGet('/checkin/calendar', params)
}

/**
 * 获取连续打卡记录
 */
export function getStreakRecord() {
  return authGet('/checkin/streak')
}

/**
 * 补签打卡（如果支持）
 * @param {Object} data 补签数据
 * @param {string} data.date 补签日期
 * @param {string} data.reason 补签原因
 */
export function makeupCheckin(data) {
  return authPost('/checkin/makeup', data)
}

/**
 * 获取打卡排行榜（游客可访问）
 * @param {Object} params 查询参数
 * @param {string} params.type 排行类型 (streak|monthly)
 * @param {number} params.limit 返回数量
 */
export function getLeaderboard(params = {}) {
  return guestGet('/checkin/leaderboard', params)
}

// 导出打卡API对象
const checkinApi = {
  dailyCheckin,
  getTodayStatus,
  getHistory,
  getStatistics,
  getWeeklyProgress,
  getCalendarData,
  getStreakRecord,
  makeupCheckin,
  getLeaderboard,
  
  // 兼容旧的方法名
  checkin: dailyCheckin,
  getStats: getStatistics
}

export default checkinApi 