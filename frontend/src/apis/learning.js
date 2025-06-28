/**
 * 学习相关API
 */
import { authGet, authPost, guestGet } from '@/utils/request.js'

/**
 * 获取学习内容列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.category 分类
 * @param {string} params.difficulty 难度
 */
export function getList(params = {}) {
  return guestGet('/learning/list', params)
}

/**
 * 获取学习内容详情
 * @param {string|number} id 内容ID
 */
export function getDetail(id) {
  return guestGet(`/learning/${id}`)
}

/**
 * 记录学习进度
 * @param {string|number} id 内容ID
 * @param {Object} data 进度数据
 * @param {number} data.progress 进度百分比
 * @param {number} data.duration 学习时长（秒）
 */
export function recordProgress(id, data) {
  return authPost(`/learning/${id}/progress`, data)
}

/**
 * 获取学习记录
 * @param {Object} params 查询参数
 */
export function getLearningRecord(params = {}) {
  return authGet('/learning/record', params)
}

/**
 * 收藏学习内容
 * @param {string|number} id 内容ID
 */
export function collectContent(id) {
  return authPost(`/learning/${id}/collect`)
}

/**
 * 取消收藏学习内容
 * @param {string|number} id 内容ID
 */
export function uncollectContent(id) {
  return authPost(`/learning/${id}/uncollect`)
}

/**
 * 获取我的收藏
 * @param {Object} params 查询参数
 */
export function getMyCollections(params = {}) {
  return authGet('/learning/my-collections', params)
}

/**
 * 评分学习内容
 * @param {string|number} id 内容ID
 * @param {Object} data 评分数据
 * @param {number} data.rating 评分 1-5
 * @param {string} data.comment 评价内容（可选）
 */
export function rateContent(id, data) {
  return authPost(`/learning/${id}/rate`, data)
}

/**
 * 获取推荐内容
 * @param {Object} params 查询参数
 * @param {string} params.contentId 基于某个内容推荐
 * @param {string} params.type 推荐类型
 */
export function getRecommendations(params = {}) {
  return guestGet('/learning/recommendations', params)
}

/**
 * 获取学习分类
 */
export function getCategories() {
  return guestGet('/learning/categories')
}

/**
 * 搜索学习内容
 * @param {Object} params 搜索参数
 * @param {string} params.keyword 关键词
 * @param {string} params.category 分类
 */
export function searchContent(params = {}) {
  return guestGet('/learning/search', params)
}

// 导出学习API对象
const learningApi = {
  getList,
  getDetail,
  recordProgress,
  getLearningRecord,
  collectContent,
  uncollectContent,
  getMyCollections,
  rateContent,
  getRecommendations,
  getCategories,
  searchContent
}

export default learningApi 