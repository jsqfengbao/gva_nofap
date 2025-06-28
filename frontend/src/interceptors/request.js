/**
 * 统一请求拦截器
 * 处理API请求的认证、错误处理和重试逻辑
 */
import { buildApiUrl } from '@/config/env.js'
import { getToken, isLoggedIn, clearUserInfo, wxLogin } from '@/utils/auth.js'

// 自定义请求选项类型
export const REQUEST_TYPES = {
  NORMAL: 'normal',      // 普通请求
  AUTH_REQUIRED: 'auth', // 需要认证的请求
  GUEST_ALLOWED: 'guest' // 允许游客访问的请求
}

// 错误重试配置
const RETRY_CONFIG = {
  maxRetries: 2,         // 最大重试次数
  retryDelay: 1000,      // 重试延迟(ms)
  retryableErrors: [     // 可重试的错误码
    'NETWORK_ERROR',
    'TIMEOUT',
    'SERVER_ERROR'
  ]
}

// 请求队列管理
const requestQueue = new Map()
let isRefreshingToken = false

/**
 * HTTP拦截器核心逻辑
 */
const httpInterceptor = {
  async invoke(options) {
    try {
      // 1. 处理查询参数
      if (options.query) {
        const queryStr = Object.keys(options.query)
          .filter(key => options.query[key] !== undefined && options.query[key] !== null)
          .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(options.query[key])}`)
          .join('&')
        
        if (queryStr) {
          if (options.url.includes('?')) {
            options.url += `&${queryStr}`
          } else {
            options.url += `?${queryStr}`
          }
        }
      }

      // 2. 处理URL拼接
      if (!options.url.startsWith('http')) {
        // #ifdef H5
        if (import.meta.env.VITE_APP_PROXY === 'true') {
          options.url = import.meta.env.VITE_APP_PROXY_PREFIX + options.url
        } else {
          options.url = buildApiUrl(options.url)
        }
        // #endif
        // #ifndef H5
        options.url = buildApiUrl(options.url)
        // #endif
      }

      // 3. 设置基础配置
      options.timeout = options.timeout || 30000
      options.header = {
        'Content-Type': 'application/json',
        'platform': getplatform(),
        ...options.header
      }

      // 4. 处理认证
      await handleAuthentication(options)

      // 5. 添加请求标识（用于去重）
      const requestId = generateRequestId(options)
      options.requestId = requestId

      console.log(`📤 API请求 [${requestId}]:`, {
        url: options.url,
        method: options.method || 'GET',
        hasAuth: !!options.header.Authorization
      })

    } catch (error) {
      console.error('❌ 请求拦截器处理失败:', error)
      throw error
    }
  }
}

/**
 * 处理认证逻辑
 */
async function handleAuthentication(options) {
  const requestType = options.requestType || REQUEST_TYPES.NORMAL
  const token = getToken()

  // 如果是需要认证的请求但没有token
  if (requestType === REQUEST_TYPES.AUTH_REQUIRED && !token) {
    throw new Error('需要登录后才能访问')
  }

  // 添加认证头
  if (token && token !== 'guest_token') {
    // 验证token格式
    if (isValidJWT(token)) {
      options.header.Authorization = `Bearer ${token}`
      options.header['x-token'] = token
    } else {
      console.warn('⚠️ 无效的token格式，清除token')
      clearUserInfo()
      if (requestType === REQUEST_TYPES.AUTH_REQUIRED) {
        throw new Error('登录信息已失效，请重新登录')
      }
    }
  }
}

/**
 * 验证JWT token格式
 */
function isValidJWT(token) {
  if (!token || typeof token !== 'string') return false
  
  const parts = token.split('.')
  if (parts.length !== 3) return false
  
  try {
    for (let part of parts) {
      while (part.length % 4) {
        part += '='
      }
      atob(part.replace(/-/g, '+').replace(/_/g, '/'))
    }
    return true
  } catch (e) {
    return false
  }
}

/**
 * 获取平台标识
 */
function getplatform() {
  // #ifdef MP-WEIXIN
  return 'mp-weixin'
  // #endif
  // #ifdef H5
  return 'h5'
  // #endif
  // #ifdef APP-PLUS
  return 'app'
  // #endif
  return 'unknown'
}

/**
 * 生成请求ID
 */
function generateRequestId(options) {
  const method = options.method || 'GET'
  const url = options.url
  const timestamp = Date.now()
  return `${method}_${url.split('?')[0].split('/').pop()}_${timestamp}`
}

/**
 * 响应拦截器
 */
const responseInterceptor = {
  success: async (response, options) => {
    const requestId = options?.requestId || 'unknown'
    console.log(`📥 API响应 [${requestId}]:`, {
      status: response.statusCode,
      success: response.statusCode === 200
    })

    // 处理认证错误
    if (isAuthError(response)) {
      return handleAuthError(response, options)
    }

    // 处理业务错误
    if (isBusinessError(response)) {
      return handleBusinessError(response, options)
    }

    return response
  },

  fail: async (error, options) => {
    const requestId = options?.requestId || 'unknown'
    console.error(`❌ API请求失败 [${requestId}]:`, error)

    // 处理网络错误
    if (isNetworkError(error)) {
      return handleNetworkError(error, options)
    }

    // 处理超时错误
    if (isTimeoutError(error)) {
      return handleTimeoutError(error, options)
    }

    throw error
  }
}

/**
 * 判断是否为认证错误
 */
function isAuthError(response) {
  return response.statusCode === 401 || 
         (response.data?.code === 401) ||
         (response.data?.code === 7 && response.data?.msg?.includes('token'))
}

/**
 * 处理认证错误
 */
async function handleAuthError(response, options) {
  console.log('🔒 检测到认证错误，处理中...')
  
  // 清除无效token
  clearUserInfo()
  
  // 如果不是必须登录的请求，直接返回错误
  if (options?.requestType !== REQUEST_TYPES.AUTH_REQUIRED) {
    return response
  }

  // 显示登录提示
  uni.showToast({
    title: '登录已过期，请重新登录',
    icon: 'none',
    duration: 2000
  })

  // 延迟跳转到登录页
  setTimeout(() => {
    uni.reLaunch({
      url: '/pages/welcome/welcome'
    })
  }, 2000)

  throw new Error('认证失败')
}

/**
 * 判断是否为业务错误
 */
function isBusinessError(response) {
  return response.statusCode === 200 && response.data?.code !== 0
}

/**
 * 处理业务错误
 */
function handleBusinessError(response, options) {
  const errorMsg = response.data?.msg || '请求失败'
  
  // 如果不隐藏错误提示，则显示toast
  if (!options?.hideErrorToast) {
    uni.showToast({
      title: errorMsg,
      icon: 'none',
      duration: 2000
    })
  }

  return response
}

/**
 * 判断是否为网络错误
 */
function isNetworkError(error) {
  return error.errMsg?.includes('fail') && !error.errMsg?.includes('timeout')
}

/**
 * 处理网络错误
 */
function handleNetworkError(error, options) {
  console.error('🌐 网络错误:', error)
  
  if (!options?.hideErrorToast) {
    uni.showToast({
      title: '网络连接失败，请检查网络设置',
      icon: 'none',
      duration: 3000
    })
  }

  throw error
}

/**
 * 判断是否为超时错误
 */
function isTimeoutError(error) {
  return error.errMsg?.includes('timeout')
}

/**
 * 处理超时错误
 */
function handleTimeoutError(error, options) {
  console.error('⏰ 请求超时:', error)
  
  if (!options?.hideErrorToast) {
    uni.showToast({
      title: '请求超时，请稍后重试',
      icon: 'none',
      duration: 2000
    })
  }

  throw error
}

/**
 * 增强的uni.request方法
 */
export function enhancedRequest(options) {
  return new Promise(async (resolve, reject) => {
    try {
      // 应用请求拦截器
      await httpInterceptor.invoke(options)
      
      // 发起请求
      uni.request({
        ...options,
        success: async (response) => {
          try {
            const processedResponse = await responseInterceptor.success(response, options)
            resolve(processedResponse)
          } catch (error) {
            reject(error)
          }
        },
        fail: async (error) => {
          try {
            await responseInterceptor.fail(error, options)
          } catch (processedError) {
            reject(processedError)
          }
        }
      })
    } catch (error) {
      reject(error)
    }
  })
}

/**
 * 请求拦截器安装器
 */
export const requestInterceptor = {
  install() {
    // 拦截 uni.request
    uni.addInterceptor('request', httpInterceptor)
    
    // 拦截 uni.uploadFile
    uni.addInterceptor('uploadFile', {
      invoke: async (options) => {
        await handleAuthentication(options)
        console.log('📤 文件上传:', options.url)
      }
    })
    
    // 拦截 uni.downloadFile  
    uni.addInterceptor('downloadFile', {
      invoke: async (options) => {
        await handleAuthentication(options)
        console.log('📥 文件下载:', options.url)
      }
    })

    console.log('✅ 请求拦截器已安装')
  }
} 