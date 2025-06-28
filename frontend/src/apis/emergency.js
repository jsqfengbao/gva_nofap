/**
 * 紧急求助相关API
 */
import { authGet, authPost, guestGet } from '@/utils/request.js'

/**
 * 获取紧急资源
 * @param {Object} params 查询参数
 * @param {string} params.type 资源类型
 * @param {string} params.category 分类
 */
export function getResources(params = {}) {
  return guestGet('/emergency/resources', params)
}

/**
 * 获取励志文章
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 */
export function getArticles(params = {}) {
  return guestGet('/emergency/articles', params)
}

/**
 * 记录求助记录
 * @param {Object} data 求助数据
 * @param {string} data.type 求助类型
 * @param {string} data.description 描述
 * @param {string} data.mood 当前心情
 */
export function recordHelp(data) {
  return authPost('/emergency/help', data)
}

/**
 * 获取求助热线
 */
export function getHotlines() {
  return guestGet('/emergency/hotlines')
}

/**
 * 获取紧急联系人
 */
export function getEmergencyContacts() {
  return authGet('/emergency/contacts')
}

/**
 * 添加紧急联系人
 * @param {Object} data 联系人数据
 * @param {string} data.name 姓名
 * @param {string} data.phone 电话
 * @param {string} data.relationship 关系
 */
export function addEmergencyContact(data) {
  return authPost('/emergency/contacts', data)
}

/**
 * 删除紧急联系人
 * @param {string|number} id 联系人ID
 */
export function removeEmergencyContact(id) {
  return authPost(`/emergency/contacts/${id}/remove`)
}

/**
 * 获取应急指南
 * @param {Object} params 查询参数
 */
export function getGuides(params = {}) {
  return guestGet('/emergency/guides', params)
}

/**
 * 获取放松练习
 * @param {Object} params 查询参数
 */
export function getRelaxationExercises(params = {}) {
  return guestGet('/emergency/relaxation', params)
}

/**
 * 记录练习使用
 * @param {string|number} exerciseId 练习ID
 * @param {Object} data 使用数据
 */
export function recordExerciseUsage(exerciseId, data) {
  return authPost(`/emergency/relaxation/${exerciseId}/usage`, data)
}

// 导出紧急求助API对象
const emergencyApi = {
  getResources,
  getArticles,
  recordHelp,
  getHotlines,
  getEmergencyContacts,
  addEmergencyContact,
  removeEmergencyContact,
  getGuides,
  getRelaxationExercises,
  recordExerciseUsage
}

export default emergencyApi 