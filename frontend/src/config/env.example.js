/**
 * 环境配置示例文件
 * 
 * 请将此文件复制为 env.js 并修改为您的实际配置
 * 
 * 配置说明：
 * 1. 本地开发环境 (DEVELOPMENT)：通常使用 localhost 或内网IP
 * 2. 测试环境 (TESTING)：使用测试服务器域名
 * 3. 生产环境 (PRODUCTION)：使用正式域名
 */

// 环境类型
export const ENV_TYPES = {
  DEVELOPMENT: 'development',
  PRODUCTION: 'production',
  TESTING: 'testing'
}

// 当前环境
export const CURRENT_ENV = process.env.NODE_ENV || ENV_TYPES.DEVELOPMENT

// API域名配置
export const API_DOMAINS = {
  [ENV_TYPES.DEVELOPMENT]: {
    BASE_URL: 'http://localhost:8888',  // 本地开发服务器
    WS_URL: 'ws://localhost:8888'
  },
  [ENV_TYPES.PRODUCTION]: {
    BASE_URL: 'https://api.yourapp.com',  // ⚠️ 请修改为您的生产环境API域名
    WS_URL: 'wss://api.yourapp.com'
  },
  [ENV_TYPES.TESTING]: {
    BASE_URL: 'https://test-api.yourapp.com',  // ⚠️ 请修改为您的测试环境API域名
    WS_URL: 'wss://test-api.yourapp.com'
  }
}

// API路径前缀配置
export const API_PREFIXES = {
  [ENV_TYPES.DEVELOPMENT]: '/v1/miniprogram',      // 本地环境去掉 /api
  [ENV_TYPES.PRODUCTION]: '/api/v1/miniprogram',   // 生产环境保留 /api
  [ENV_TYPES.TESTING]: '/api/v1/miniprogram'       // 测试环境保留 /api
}

// 静态资源域名配置
export const STATIC_DOMAINS = {
  [ENV_TYPES.DEVELOPMENT]: 'http://localhost:8888',
  [ENV_TYPES.PRODUCTION]: 'https://cdn.yourapp.com',      // ⚠️ 请修改为您的CDN域名
  [ENV_TYPES.TESTING]: 'https://test-cdn.yourapp.com'     // ⚠️ 请修改为您的测试CDN域名
}

// CDN配置
export const CDN_CONFIG = {
  [ENV_TYPES.DEVELOPMENT]: {
    IMAGES: 'http://localhost:8888/static/images',
    VIDEOS: 'http://localhost:8888/static/videos',
    DOCUMENTS: 'http://localhost:8888/static/documents'
  },
  [ENV_TYPES.PRODUCTION]: {
    IMAGES: 'https://cdn.yourapp.com/images',       // ⚠️ 请修改为您的图片CDN
    VIDEOS: 'https://cdn.yourapp.com/videos',       // ⚠️ 请修改为您的视频CDN
    DOCUMENTS: 'https://cdn.yourapp.com/documents'  // ⚠️ 请修改为您的文档CDN
  },
  [ENV_TYPES.TESTING]: {
    IMAGES: 'https://test-cdn.yourapp.com/images',
    VIDEOS: 'https://test-cdn.yourapp.com/videos',
    DOCUMENTS: 'https://test-cdn.yourapp.com/documents'
  }
}

// 第三方服务配置
export const THIRD_PARTY_CONFIG = {
  [ENV_TYPES.DEVELOPMENT]: {
    WECHAT_APP_ID: 'wx07c9e8e4f105260b',        // ⚠️ 请修改为您的开发环境微信小程序AppID
    ANALYTICS_ID: null,                     // 开发环境不启用统计
    ERROR_REPORT_URL: null                  // 开发环境不启用错误上报
  },
  [ENV_TYPES.PRODUCTION]: {
    WECHAT_APP_ID: 'wx07c9e8e4f105260b',       // ⚠️ 请修改为您的生产环境微信小程序AppID
    ANALYTICS_ID: 'ga_tracking_id',        // ⚠️ 请修改为您的统计ID
    ERROR_REPORT_URL: 'https://api.yourapp.com/error-report'  // ⚠️ 请修改为您的错误上报URL
  },
  [ENV_TYPES.TESTING]: {
    WECHAT_APP_ID: 'wx07c9e8e4f105260b',       // ⚠️ 请修改为您的测试环境微信小程序AppID
    ANALYTICS_ID: 'test_analytics_id',
    ERROR_REPORT_URL: 'https://test-api.yourapp.com/error-report'
  }
}

// 超时配置（单位：毫秒）
export const TIMEOUT_CONFIG = {
  API_REQUEST: 30000,   // API请求超时时间 (30秒)
  UPLOAD_FILE: 60000,   // 文件上传超时时间 (60秒)
  WS_CONNECT: 10000,    // WebSocket连接超时时间 (10秒)
  LOGIN_TIMEOUT: 20000  // 登录超时时间 (20秒)
}

// 获取当前环境的配置
export function getCurrentConfig() {
  return {
    env: CURRENT_ENV,
    domain: API_DOMAINS[CURRENT_ENV],
    apiPrefix: API_PREFIXES[CURRENT_ENV],
    staticDomain: STATIC_DOMAINS[CURRENT_ENV],
    cdn: CDN_CONFIG[CURRENT_ENV],
    thirdParty: THIRD_PARTY_CONFIG[CURRENT_ENV],
    timeout: TIMEOUT_CONFIG
  }
}

// 获取完整的API URL
export function buildApiUrl(path = '') {
  const config = getCurrentConfig()
  const cleanPath = path.startsWith('/') ? path : `/${path}`
  return `${config.domain.BASE_URL}${config.apiPrefix}${cleanPath}`
}

// 获取WebSocket URL
export function buildWsUrl(path = '') {
  const config = getCurrentConfig()
  const cleanPath = path.startsWith('/') ? path : `/${path}`
  return `${config.domain.WS_URL}${cleanPath}`
}

// 获取静态资源URL
export function buildStaticUrl(path = '', type = 'default') {
  const config = getCurrentConfig()
  
  // 如果是CDN资源类型
  if (type in config.cdn) {
    const cleanPath = path.startsWith('/') ? path.substring(1) : path
    return `${config.cdn[type]}/${cleanPath}`
  }
  
  // 默认静态资源
  const cleanPath = path.startsWith('/') ? path : `/${path}`
  return `${config.staticDomain}/static${cleanPath}`
}

// 获取图片URL
export function getImageUrl(path) {
  return buildStaticUrl(path, 'IMAGES')
}

// 获取视频URL
export function getVideoUrl(path) {
  return buildStaticUrl(path, 'VIDEOS')
}

// 获取文档URL
export function getDocumentUrl(path) {
  return buildStaticUrl(path, 'DOCUMENTS')
}

// 验证环境配置
export function validateConfig() {
  const config = getCurrentConfig()
  const errors = []
  
  if (!config.domain.BASE_URL) {
    errors.push('API Base URL is not configured')
  }
  
  if (!config.apiPrefix) {
    errors.push('API Prefix is not configured')
  }
  
  if (!config.staticDomain) {
    errors.push('Static Domain is not configured')
  }
  
  if (config.domain.BASE_URL.includes('yourapp.com')) {
    errors.push('Please replace "yourapp.com" with your actual domain')
  }
  
  if (config.thirdParty.WECHAT_APP_ID.includes('_app_id')) {
    errors.push('Please configure your actual WeChat Mini Program App ID')
  }
  
  return {
    valid: errors.length === 0,
    errors
  }
}

// 导出当前环境配置
export default getCurrentConfig()

/* 
配置检查清单：
□ 修改 API_DOMAINS 中的生产环境和测试环境域名
□ 修改 STATIC_DOMAINS 中的CDN域名
□ 修改 CDN_CONFIG 中的各类资源CDN地址
□ 配置 WECHAT_APP_ID 为实际的微信小程序AppID
□ 配置 ANALYTICS_ID 为实际的统计服务ID
□ 配置 ERROR_REPORT_URL 为实际的错误上报地址
□ 根据需要调整 TIMEOUT_CONFIG 中的超时时间
□ 确保所有占位符（如 yourapp.com）都已替换为实际值
*/ 