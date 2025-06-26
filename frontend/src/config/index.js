/**
 * 应用配置文件
 */
import { getCurrentConfig, buildApiUrl, buildWsUrl, buildStaticUrl, validateConfig } from './env.js'

// 获取当前环境配置
const envConfig = getCurrentConfig()

// 导出配置
export const config = {
  // 环境信息
  ENV: envConfig.env,
  IS_DEV: envConfig.env === 'development',
  IS_PROD: envConfig.env === 'production',
  IS_TEST: envConfig.env === 'testing',
  
  // API配置 - 使用环境配置
  API: {
    BASE_URL: envConfig.domain.BASE_URL,
    API_PREFIX: envConfig.apiPrefix,
    TIMEOUT: envConfig.timeout.API_REQUEST,
    WS_URL: envConfig.domain.WS_URL,
    UPLOAD_TIMEOUT: envConfig.timeout.UPLOAD_FILE,
    LOGIN_TIMEOUT: envConfig.timeout.LOGIN_TIMEOUT
  },
  
  // 静态资源配置
  STATIC: {
    DOMAIN: envConfig.staticDomain,
    CDN: envConfig.cdn
  },
  
  // 第三方服务配置
  THIRD_PARTY: envConfig.thirdParty,
  
  // 小程序配置
  MINIPROGRAM: {
    NAME: 'NoFap 戒色助手',
    VERSION: '1.0.0',
    DESCRIPTION: '帮助用户建立健康生活习惯的微信小程序',
    APP_ID: envConfig.thirdParty.WECHAT_APP_ID
  },
  
  // 存储键名
  STORAGE_KEYS: {
    TOKEN: 'token',
    USER_INFO: 'userInfo',
    HAS_LAUNCHED: 'hasLaunched',
    SETTINGS: 'settings',
    THEME: 'theme'
  },
  
  // 页面路径
  PAGES: {
    WELCOME: '/pages/welcome/welcome',
    INDEX: '/pages/index/index',
    PROFILE: '/pages/profile/index',
    CHECKIN: '/pages/checkin/index',
    COMMUNITY: '/pages/community/index',
    LEARNING: '/pages/learning/index',
    ASSESSMENT: '/pages/assessment/index',
    ACHIEVEMENT: '/pages/achievement/index',
    EMERGENCY: '/pages/emergency/index',
    PROGRESS: '/pages/progress/index'
  },
  
  // 主题配置
  THEME: {
    PRIMARY_COLOR: '#34d399',
    SECONDARY_COLOR: '#06b6d4',
    ACCENT_COLOR: '#f59e0b',
    SUCCESS_COLOR: '#10b981',
    WARNING_COLOR: '#f59e0b',
    DANGER_COLOR: '#ef4444',
    INFO_COLOR: '#3b82f6'
  },
  
  // 功能开关
  FEATURES: {
    GUEST_MODE: true,
    DARK_MODE: true,
    OFFLINE_MODE: false,
    PUSH_NOTIFICATION: true,
    ANALYTICS: envConfig.thirdParty.ANALYTICS_ID !== null,
    ERROR_REPORT: envConfig.thirdParty.ERROR_REPORT_URL !== null
  },
  
  // 限制配置
  LIMITS: {
    MAX_UPLOAD_SIZE: 10 * 1024 * 1024, // 10MB
    MAX_CONTENT_LENGTH: 2000,
    MAX_NICKNAME_LENGTH: 50,
    MIN_PASSWORD_LENGTH: 6,
    MAX_PASSWORD_LENGTH: 20
  },
  
  // 缓存配置
  CACHE: {
    USER_INFO_EXPIRES: 7 * 24 * 60 * 60 * 1000, // 7天
    API_CACHE_EXPIRES: 5 * 60 * 1000, // 5分钟
    IMAGE_CACHE_EXPIRES: 24 * 60 * 60 * 1000 // 24小时
  },
  
  // 默认设置
  DEFAULT_SETTINGS: {
    theme: 'auto', // auto, light, dark
    language: 'zh-CN',
    notifications: {
      checkinReminder: true,
      communityReply: true,
      achievementUnlock: true,
      weeklyReport: true,
      emergencyAlert: true,
      learningReminder: true
    },
    privacy: {
      showProfile: true,
      showStats: true,
      showAchievements: true,
      allowFriendRequest: true,
      showOnlineStatus: true
    }
  }
}

// 获取完整的API URL - 使用环境配置
export function getApiUrl(path = '') {
  return buildApiUrl(path)
}

// 获取WebSocket URL - 使用环境配置
export function getWsUrl(path = '') {
  return buildWsUrl(path)
}

// 获取静态资源URL - 使用环境配置
export function getStaticUrl(path = '', type = 'default') {
  return buildStaticUrl(path, type)
}

// 调试日志
export function debugLog(...args) {
  if (config.IS_DEV) {
    console.log('[NoFap Debug]', ...args)
  }
}

// 错误日志
export function errorLog(...args) {
  console.error('[NoFap Error]', ...args)
  
  // 如果配置了错误上报，发送错误报告
  if (config.FEATURES.ERROR_REPORT && envConfig.thirdParty.ERROR_REPORT_URL) {
    try {
      // 可以在这里添加错误上报逻辑
      console.log('Error reported to:', envConfig.thirdParty.ERROR_REPORT_URL)
    } catch (error) {
      console.error('Failed to report error:', error)
    }
  }
}

// 验证配置完整性
export function validateAppConfig() {
  const validation = validateConfig()
  
  if (!validation.valid) {
    console.error('Configuration validation failed:', validation.errors)
    return false
  }
  
  return true
}

// 获取当前环境信息
export function getEnvInfo() {
  return {
    environment: envConfig.env,
    apiBaseUrl: envConfig.domain.BASE_URL,
    staticDomain: envConfig.staticDomain,
    features: config.FEATURES,
    version: config.MINIPROGRAM.VERSION
  }
}

// 导出默认配置
export default config 