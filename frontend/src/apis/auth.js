/**
 * 认证相关API
 */
import { post, guestPost } from '@/utils/request.js'

/**
 * 微信登录
 * @param {Object} data 登录数据
 * @param {string} data.code 微信登录code
 * @param {string} data.encryptedData 加密用户信息（可选）
 * @param {string} data.iv 初始向量（可选）
 */
export function wxLogin(data) {
  return guestPost('/auth/wx-login', data)
}

/**
 * 游客登录
 */
export function guestLogin() {
  return guestPost('/auth/guest-login')
}

/**
 * 刷新token
 * @param {Object} data 刷新数据
 * @param {string} data.refreshToken 刷新token
 */
export function refreshToken(data) {
  return post('/auth/refresh-token', data)
}

/**
 * 登出
 */
export function logout() {
  return post('/auth/logout')
}

/**
 * 检查登录状态
 */
export function checkLoginStatus() {
  return post('/auth/check-status')
}

// 导出认证API对象
const authApi = {
  wxLogin,
  guestLogin,
  refreshToken,
  logout,
  checkLoginStatus
}

export default authApi 