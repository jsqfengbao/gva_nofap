/**
 * 成就相关API
 */
import { authGet, guestGet } from '@/utils/request.js'

/**
 * 获取成就列表
 * @param {Object} params 查询参数
 * @param {string} params.category 成就分类
 * @param {boolean} params.unlocked 是否只显示已解锁
 */
export function getList(params = {}) {
  return guestGet('/achievement/list', params)
}

/**
 * 获取用户成就
 * @param {Object} params 查询参数
 * @param {number} params.limit 返回数量限制
 * @param {boolean} params.recent 是否只返回最近获得的
 */
export function getUserAchievements(params = {}) {
  return authGet('/achievement/user', params)
}

/**
 * 获取成就详情
 * @param {string|number} id 成就ID
 */
export function getAchievementDetail(id) {
  return guestGet(`/achievement/${id}`)
}

/**
 * 获取成就进度
 * @param {string|number} id 成就ID
 */
export function getAchievementProgress(id) {
  return authGet(`/achievement/${id}/progress`)
}

/**
 * 获取游戏化统计数据（首页使用）
 */
export function getGameStats() {
  return authGet('/achievement/game-stats')
}

/**
 * 获取成就排行榜
 * @param {Object} params 查询参数
 * @param {string} params.type 排行类型
 * @param {number} params.limit 返回数量
 */
export function getLeaderboard(params = {}) {
  return guestGet('/achievement/leaderboard', params)
}

/**
 * 获取成就分类
 */
export function getCategories() {
  return guestGet('/achievement/categories')
}

/**
 * 获取用户等级信息
 */
export function getUserLevel() {
  return authGet('/achievement/user-level')
}

/**
 * 获取经验值历史
 * @param {Object} params 查询参数
 */
export function getExpHistory(params = {}) {
  return authGet('/achievement/exp-history', params)
}

// 导出成就API对象
const achievementApi = {
  getList,
  getUserAchievements,
  getAchievementDetail,
  getAchievementProgress,
  getGameStats,
  getLeaderboard,
  getCategories,
  getUserLevel,
  getExpHistory
}

export default achievementApi 