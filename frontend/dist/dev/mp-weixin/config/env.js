"use strict";
const ENV_TYPES = {
  DEVELOPMENT: "development",
  PRODUCTION: "production",
  TESTING: "testing"
};
const CURRENT_ENV = "development";
const API_DOMAINS = {
  [ENV_TYPES.DEVELOPMENT]: {
    BASE_URL: "http://192.168.1.140:8888",
    WS_URL: "ws://192.168.1.140:8888"
  },
  [ENV_TYPES.PRODUCTION]: {
    BASE_URL: "https://api.nofap-app.com",
    // 生产环境API域名
    WS_URL: "wss://api.nofap-app.com"
  },
  [ENV_TYPES.TESTING]: {
    BASE_URL: "https://test-api.nofap-app.com",
    // 测试环境API域名
    WS_URL: "wss://test-api.nofap-app.com"
  }
};
const API_PREFIXES = {
  [ENV_TYPES.DEVELOPMENT]: "/api/v1/miniprogram",
  // 本地环境也需要 /api 前缀
  [ENV_TYPES.PRODUCTION]: "/api/v1/miniprogram",
  // 生产环境保留 /api
  [ENV_TYPES.TESTING]: "/api/v1/miniprogram"
  // 测试环境保留 /api
};
const STATIC_DOMAINS = {
  [ENV_TYPES.DEVELOPMENT]: "http://192.168.1.140:8888",
  [ENV_TYPES.PRODUCTION]: "https://cdn.nofap-app.com",
  [ENV_TYPES.TESTING]: "https://test-cdn.nofap-app.com"
};
const CDN_CONFIG = {
  [ENV_TYPES.DEVELOPMENT]: {
    IMAGES: "http://192.168.1.140:8888/static/images",
    VIDEOS: "http://192.168.1.140:8888/static/videos",
    DOCUMENTS: "http://192.168.1.140:8888/static/documents"
  },
  [ENV_TYPES.PRODUCTION]: {
    IMAGES: "https://cdn.nofap-app.com/images",
    VIDEOS: "https://cdn.nofap-app.com/videos",
    DOCUMENTS: "https://cdn.nofap-app.com/documents"
  },
  [ENV_TYPES.TESTING]: {
    IMAGES: "https://test-cdn.nofap-app.com/images",
    VIDEOS: "https://test-cdn.nofap-app.com/videos",
    DOCUMENTS: "https://test-cdn.nofap-app.com/documents"
  }
};
const THIRD_PARTY_CONFIG = {
  [ENV_TYPES.DEVELOPMENT]: {
    WECHAT_APP_ID: "wx07c9e8e4f105260b",
    // 开发环境微信小程序AppID
    ANALYTICS_ID: null,
    // 开发环境不启用统计
    ERROR_REPORT_URL: null
    // 开发环境不启用错误上报
  },
  [ENV_TYPES.PRODUCTION]: {
    WECHAT_APP_ID: "wx07c9e8e4f105260b",
    // 生产环境微信小程序AppID
    ANALYTICS_ID: "ga_tracking_id",
    // 生产环境统计ID
    ERROR_REPORT_URL: "https://api.nofap-app.com/error-report"
  },
  [ENV_TYPES.TESTING]: {
    WECHAT_APP_ID: "wx07c9e8e4f105260b",
    // 测试环境微信小程序AppID
    ANALYTICS_ID: "test_analytics_id",
    ERROR_REPORT_URL: "https://test-api.nofap-app.com/error-report"
  }
};
const TIMEOUT_CONFIG = {
  API_REQUEST: 3e4,
  // API请求超时时间 (30秒)
  UPLOAD_FILE: 6e4,
  // 文件上传超时时间 (60秒)
  WS_CONNECT: 1e4,
  // WebSocket连接超时时间 (10秒)
  LOGIN_TIMEOUT: 2e4
  // 登录超时时间 (20秒)
};
function getCurrentConfig() {
  return {
    env: CURRENT_ENV,
    domain: API_DOMAINS[CURRENT_ENV],
    apiPrefix: API_PREFIXES[CURRENT_ENV],
    staticDomain: STATIC_DOMAINS[CURRENT_ENV],
    cdn: CDN_CONFIG[CURRENT_ENV],
    thirdParty: THIRD_PARTY_CONFIG[CURRENT_ENV],
    timeout: TIMEOUT_CONFIG
  };
}
function buildApiUrl(path = "") {
  const config = getCurrentConfig();
  const cleanPath = path.startsWith("/") ? path : `/${path}`;
  return `${config.domain.BASE_URL}${config.apiPrefix}${cleanPath}`;
}
function buildStaticUrl(path = "", type = "default") {
  const config = getCurrentConfig();
  if (type in config.cdn) {
    const cleanPath2 = path.startsWith("/") ? path.substring(1) : path;
    return `${config.cdn[type]}/${cleanPath2}`;
  }
  const cleanPath = path.startsWith("/") ? path : `/${path}`;
  return `${config.staticDomain}/static${cleanPath}`;
}
function getImageUrl(path) {
  return buildStaticUrl(path, "IMAGES");
}
exports.buildApiUrl = buildApiUrl;
exports.getCurrentConfig = getCurrentConfig;
exports.getImageUrl = getImageUrl;
