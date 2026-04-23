/**
 * 认证相关工具函数
 */
import { buildApiUrl, getImageUrl, getCurrentConfig } from '../config/env.js'

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
  const token = uni.getStorageSync('token') || ''
  // 验证token格式 - 简化版本，只检查三段式结构
  if (token && token !== 'guest_token') {
    const parts = token.split('.')
    if (parts.length !== 3) {
      console.warn('检测到无效token，自动清除:', token)
      clearUserInfo()
      return ''
    }
  }
  return token
}

// 验证JWT token格式 - 简化版本
function isValidJWT(token) {
  if (!token || typeof token !== 'string') {
    return false
  }
  
  // 只检查是否是三段式结构，避免base64解码在小程序环境的问题
  return token.split('.').length === 3
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
  // 不跳转，留在当前页面
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

// 检查微信小程序配置
function checkWechatConfig() {
  // 检查当前环境的AppID配置
  const config = getCurrentConfig()
  const appId = config.thirdParty.WECHAT_APP_ID
  
  // wx07c9e8e4f105260b 是真实的AppID，不是示例AppID
  if (!appId || appId === 'your_wechat_app_id_here' || appId === 'test_app_id') {
    console.warn('⚠️ 使用的是示例AppID，可能导致登录失败')
    return false
  }
  
  return true
}

// 微信授权登录
export async function wxLogin(encryptedData = '', iv = '', retryCount = 0) {
  const maxRetries = 2 // 最大重试次数
  
  return new Promise((resolve, reject) => {
    console.log(`开始微信登录流程... (尝试次数: ${retryCount + 1})`)
    
    // 检查配置
    if (!checkWechatConfig()) {
      reject(new Error('微信小程序配置错误，请联系开发者'))
      return
    }
    
    uni.login({
      provider: 'weixin',
      success: async (loginRes) => {
        try {
          console.log('微信登录成功，获取到code:', loginRes.code)
          console.log('code长度:', loginRes.code?.length)
          console.log('完整登录响应:', loginRes)
          
          // 立即发送请求，减少code过期的可能性
          const startTime = Date.now()
          
          const requestData = {
            code: loginRes.code
          }
          
          // 如果有加密用户信息，则包含进去
          if (encryptedData && iv) {
            requestData.encryptedData = encryptedData
            requestData.iv = iv
            console.log('包含用户加密信息进行登录')
          } else {
            console.log('基础登录，仅使用code')
          }
          
          console.log('准备发送登录请求，请求数据:', requestData)
          console.log('请求URL:', buildApiUrl('/auth/wx-login'))
          
          const res = await uni.request({
            url: buildApiUrl('/auth/wx-login'),
            method: 'POST',
            header: {
              'Content-Type': 'application/json'
            },
            data: requestData,
            timeout: 20000
          })

          const endTime = Date.now()
          console.log(`登录请求耗时: ${endTime - startTime}ms`)
          console.log('登录API完整响应:', res)
          console.log('响应状态码:', res.statusCode)
          console.log('响应数据:', res.data)

          if (res.statusCode !== 200) {
            throw new Error(`HTTP错误: ${res.statusCode}`)
          }

          if (res.data.code === 0) {
            console.log('登录成功，保存用户信息')
            // 保存token和用户信息
            uni.setStorageSync('token', res.data.data.token)
            if (res.data.data.user) {
              uni.setStorageSync('userInfo', res.data.data.user)
            }
            resolve(res.data.data)
          } else {
            console.error('登录失败，服务器返回错误:', res.data)
            
            // 如果是code相关错误，尝试重新登录一次
            if (res.data.msg && res.data.msg.includes('invalid code')) {
              console.log('检测到code错误，尝试重新登录...')
              setTimeout(() => {
                wxLogin(encryptedData, iv).then(resolve).catch(reject)
              }, 1000)
              return
            }
            
            reject(new Error(res.data.msg || '登录失败'))
          }
        } catch (error) {
          console.error('登录请求异常:', error)
          
          // 网络错误时的处理
          if (error.errMsg && error.errMsg.includes('timeout')) {
            reject(new Error('网络超时，请检查网络连接'))
          } else if (error.errMsg && error.errMsg.includes('fail')) {
            reject(new Error('网络请求失败，请稍后重试'))
          } else {
            reject(error)
          }
        }
      },
      fail: (err) => {
        console.error('微信登录失败详细信息:', err)
        let errorMsg = '获取登录凭证失败'
        
        // 处理不同类型的错误
        if (err.errMsg) {
          console.log('微信登录错误信息:', err.errMsg)
          
          // 检查是否是access_token missing错误
          if (err.errMsg.includes('access_token missing')) {
            errorMsg = 'AppID配置错误或小程序未激活，请联系开发者检查微信小程序配置'
          } else if (err.errMsg.includes('需要重新登录')) {
            errorMsg = '登录凭证已失效，请重新尝试登录'
          } else if (err.errMsg.includes('invalid code')) {
            errorMsg = '登录凭证无效，请重新尝试'
          } else if (err.errMsg.includes('timeout')) {
            errorMsg = '网络超时，请检查网络连接后重试'
          }
        }
        
        // 根据错误码提供更具体的错误信息
        if (err.errCode) {
          console.log('微信登录错误码:', err.errCode)
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
    // 确保token是有效的JWT格式
    if (token.split('.').length === 3) {
      config.header.Authorization = `Bearer ${token}`
      config.header['x-token'] = token
    } else {
      console.warn('❌ 无效的token格式，跳过认证:', token)
      // 清除无效token
      clearUserInfo()
    }
  }
  
  return new Promise((resolve, reject) => {
    uni.request({
      ...config,
      success: (res) => {
        // 检查是否为认证错误
        if (res.statusCode === 401 || 
            (res.data && res.data.code === 401) ||
            (res.data && res.data.code === 7 && res.data.msg?.includes('token'))) {
          
          console.error('🔒 认证失败，清除token并重新登录')
          
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
        console.error('❌ API请求失败:', err.errMsg)
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
    return null  // 返回null，让UI组件显示默认占位符（emoji等）
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

// 移除调试函数，只保留基本功能

 