/**
 * 认证相关工具函数
 */
import { buildApiUrl, getImageUrl } from '../config/env.js'

// 检查是否已登录
export function isLoggedIn() {
  const token = uni.getStorageSync('token')
  return token && token !== '' && token !== 'guest_token'
}

// 检查是否为游客模式
export function isGuestMode() {
  const token = uni.getStorageSync('token')
  return token === 'guest_token'
}

// 获取用户信息
export function getUserInfo() {
  try {
    const userInfo = uni.getStorageSync('userInfo')
    return userInfo ? JSON.parse(typeof userInfo === 'string' ? userInfo : JSON.stringify(userInfo)) : null
  } catch (error) {
    console.error('获取用户信息失败:', error)
    return null
  }
}

// 获取token
export function getToken() {
  return uni.getStorageSync('token') || ''
}

// 设置用户信息
export function setUserInfo(userInfo) {
  try {
    uni.setStorageSync('userInfo', userInfo)
    uni.setStorageSync('token', userInfo.token || getToken())
  } catch (error) {
    console.error('保存用户信息失败:', error)
  }
}

// 清除用户信息
export function clearUserInfo() {
  uni.removeStorageSync('token')
  uni.removeStorageSync('userInfo')
}

// 登出
export function logout() {
  clearUserInfo()
  
  // 跳转到登录页
  uni.reLaunch({
    url: '/pages/welcome/welcome'
  })
}

// 检查是否需要登录
export function requireAuth() {
  if (!isLoggedIn()) {
    uni.showModal({
      title: '需要登录',
      content: '此功能需要登录后使用，是否进行微信登录？',
      success: (res) => {
        if (res.confirm) {
          // 直接调用微信登录
          wxLogin().then(() => {
            uni.showToast({
              title: '登录成功',
              icon: 'success'
            })
          }).catch((error) => {
            uni.showToast({
              title: error.message || '登录失败',
              icon: 'error'
            })
          })
        }
      }
    })
    return false
  }
  return true
}

// 微信授权登录
export async function wxLogin(encryptedData = '', iv = '') {
  return new Promise((resolve, reject) => {
    uni.login({
      provider: 'weixin',
      success: async (loginRes) => {
        try {
          const requestData = {
            code: loginRes.code
          }
          
          // 如果有加密用户信息，则包含进去
          if (encryptedData && iv) {
            requestData.encryptedData = encryptedData
            requestData.iv = iv
          }
          
          const res = await uni.request({
            url: buildApiUrl('/auth/wx-login'),  // 使用环境配置的URL
            method: 'POST',
            header: {
              'Content-Type': 'application/json'
            },
            data: requestData
          })

          if (res.data.code === 0) {
            // 保存token和用户信息
            uni.setStorageSync('token', res.data.data.token)
            uni.setStorageSync('userInfo', res.data.data.user)
            resolve(res.data.data)
          } else {
            reject(new Error(res.data.msg || '登录失败'))
          }
        } catch (error) {
          reject(error)
        }
      },
      fail: (err) => {
        console.error('微信登录失败详细信息:', err)
        let errorMsg = '获取登录凭证失败'
        
        // 根据错误码提供更具体的错误信息
        if (err.errCode) {
          switch (err.errCode) {
            case -1:
              errorMsg = '系统错误，请稍后重试'
              break
            case 40013:
              errorMsg = 'AppID无效，请检查AppID配置'
              break
            case 40125:
              errorMsg = 'AppID未激活，请在微信公众平台激活小程序'
              break
            case 40163:
              errorMsg = '登录态已过期，请重试'
              break
            default:
              errorMsg = `登录失败(错误码:${err.errCode}): ${err.errMsg || '未知错误'}`
          }
        }
        
        reject(new Error(errorMsg))
      }
    })
  })
}

// API请求封装（自动携带token）
export function request(options) {
  const token = getToken()
  
  // 设置默认配置
  const defaultOptions = {
    header: {
      'Content-Type': 'application/json'
    },
    timeout: 30000
  }
  
  // 合并配置
  const config = Object.assign({}, defaultOptions, options)
  
  // 添加token到请求头
  if (token && token !== 'guest_token') {
    config.header.Authorization = `Bearer ${token}`
    config.header['x-token'] = token
  }
  
  return new Promise((resolve, reject) => {
    uni.request({
      ...config,
      success: (res) => {
        // 检查是否为认证错误
        if (res.statusCode === 401 || 
            (res.data && res.data.code === 401)) {
          // token已过期或无效，清除并跳转登录
          clearUserInfo()
          uni.showToast({
            title: '登录已过期，请重新登录',
            icon: 'none'
          })
          setTimeout(() => {
            uni.reLaunch({
              url: '/pages/welcome/welcome'
            })
          }, 1500)
          reject(new Error('认证失败'))
          return
        }
        
        resolve(res)
      },
      fail: (err) => {
        reject(err)
      }
    })
  })
}

// 获取API基础URL
export function getApiBaseUrl() {
  return buildApiUrl('')  // 使用环境配置
}

// 格式化API URL
export function formatApiUrl(path) {
  return buildApiUrl(path)  // 使用环境配置
}

// 用户头像处理
export function getAvatarUrl(avatarUrl) {
  if (!avatarUrl) {
    return getImageUrl('default-avatar.png')  // 使用环境配置的图片URL
  }
  
  // 如果是微信头像URL，直接返回
  if (avatarUrl.startsWith('http')) {
    return avatarUrl
  }
  
  // 如果是相对路径，使用CDN或静态资源域名
  return getImageUrl(avatarUrl)
}

// 检查网络状态
export function checkNetworkStatus() {
  return new Promise((resolve) => {
    uni.getNetworkType({
      success: (res) => {
        resolve({
          isConnected: res.networkType !== 'none',
          networkType: res.networkType
        })
      },
      fail: () => {
        resolve({
          isConnected: false,
          networkType: 'unknown'
        })
      }
    })
  })
}

// 显示网络错误提示
export function showNetworkError() {
  uni.showToast({
    title: '网络连接失败，请检查网络设置',
    icon: 'none',
    duration: 3000
  })
}

// 防抖函数
export function debounce(func, wait) {
  let timeout
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout)
      func(...args)
    }
    clearTimeout(timeout)
    timeout = setTimeout(later, wait)
  }
}

// 节流函数
export function throttle(func, limit) {
  let inThrottle
  return function() {
    const args = arguments
    const context = this
    if (!inThrottle) {
      func.apply(context, args)
      inThrottle = true
      setTimeout(() => inThrottle = false, limit)
    }
  }
} 