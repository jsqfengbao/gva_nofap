/**
 * 评估相关API
 */
import { authGet, authPost, guestGet } from '@/utils/request.js'

/**
 * 获取评估问题
 * @param {string} type 评估类型
 */
export function getQuestions(type) {
  return guestGet(`/assessment/questions/${type}`)
}

/**
 * 提交评估结果
 * @param {Object} data 评估数据
 * @param {string} data.type 评估类型
 * @param {Array} data.answers 答案列表
 */
export function submitResult(data) {
  return authPost('/assessment/submit', data)
}

/**
 * 获取评估历史
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.type 评估类型
 */
export function getHistory(params = {}) {
  return authGet('/assessment/history', params)
}

/**
 * 获取评估报告
 * @param {string|number} id 评估记录ID
 */
export function getReport(id) {
  return authGet(`/assessment/report/${id}`)
}

/**
 * 获取评估统计
 */
export function getStats() {
  return authGet('/assessment/stats')
}

/**
 * 获取评估类型列表
 */
export function getTypes() {
  return guestGet('/assessment/types')
}

/**
 * 重新评估
 * @param {string} type 评估类型
 */
export function retakeAssessment(type) {
  return authPost(`/assessment/retake/${type}`)
}

// 导出评估API对象
const assessmentApi = {
  getQuestions,
  submitResult,
  getHistory,
  getReport,
  getStats,
  getTypes,
  retakeAssessment
}

export default assessmentApi 