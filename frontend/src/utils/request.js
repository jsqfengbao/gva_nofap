/**
 * 统一请求工具
 * 使用拦截器系统，提供统一的API调用接口
 */
import { enhancedRequest, REQUEST_TYPES } from '@/interceptors/request.js'
import { buildApiUrl } from '@/config/env.js'

/**
 * 基础请求方法
 * @param {Object} options 请求选项
 * @param {string} options.url 请求URL
 * @param {string} options.method 请求方法
 * @param {Object} options.data 请求数据
 * @param {Object} options.query 查询参数
 * @param {Object} options.header 请求头
 * @param {string} options.requestType 请求类型 (normal|auth|guest)
 * @param {boolean} options.hideErrorToast 是否隐藏错误提示
 * @returns {Promise} 请求Promise
 */
export function request(options = {}) {
  const config = {
    method: 'GET',
    requestType: REQUEST_TYPES.NORMAL,
    hideErrorToast: false,
    ...options
  }

  // 如果URL不是完整URL，则使用buildApiUrl构建
  if (!config.url.startsWith('http')) {
    config.url = buildApiUrl(config.url)
  }

  return enhancedRequest(config)
}

/**
 * GET请求
 * @param {string} url 请求URL
 * @param {Object} params 查询参数
 * @param {Object} options 其他选项
 */
export function get(url, params = {}, options = {}) {
  return request({
    url,
    method: 'GET',
    query: params,
    ...options
  })
}

/**
 * POST请求
 * @param {string} url 请求URL
 * @param {Object} data 请求数据
 * @param {Object} options 其他选项
 */
export function post(url, data = {}, options = {}) {
  return request({
    url,
    method: 'POST',
    data,
    ...options
  })
}

/**
 * PUT请求
 * @param {string} url 请求URL
 * @param {Object} data 请求数据
 * @param {Object} options 其他选项
 */
export function put(url, data = {}, options = {}) {
  return request({
    url,
    method: 'PUT',
    data,
    ...options
  })
}

/**
 * DELETE请求
 * @param {string} url 请求URL
 * @param {Object} params 查询参数
 * @param {Object} options 其他选项
 */
export function del(url, params = {}, options = {}) {
  return request({
    url,
    method: 'DELETE',
    query: params,
    ...options
  })
}

/**
 * 需要认证的GET请求
 * @param {string} url 请求URL
 * @param {Object} params 查询参数
 * @param {Object} options 其他选项
 */
export function authGet(url, params = {}, options = {}) {
  return get(url, params, {
    requestType: REQUEST_TYPES.AUTH_REQUIRED,
    ...options
  })
}

/**
 * 需要认证的POST请求
 * @param {string} url 请求URL
 * @param {Object} data 请求数据
 * @param {Object} options 其他选项
 */
export function authPost(url, data = {}, options = {}) {
  return post(url, data, {
    requestType: REQUEST_TYPES.AUTH_REQUIRED,
    ...options
  })
}

/**
 * 需要认证的PUT请求
 * @param {string} url 请求URL
 * @param {Object} data 请求数据
 * @param {Object} options 其他选项
 */
export function authPut(url, data = {}, options = {}) {
  return put(url, data, {
    requestType: REQUEST_TYPES.AUTH_REQUIRED,
    ...options
  })
}

/**
 * 需要认证的DELETE请求
 * @param {string} url 请求URL
 * @param {Object} params 查询参数
 * @param {Object} options 其他选项
 */
export function authDel(url, params = {}, options = {}) {
  return del(url, params, {
    requestType: REQUEST_TYPES.AUTH_REQUIRED,
    ...options
  })
}

/**
 * 允许游客访问的GET请求
 * @param {string} url 请求URL
 * @param {Object} params 查询参数
 * @param {Object} options 其他选项
 */
export function guestGet(url, params = {}, options = {}) {
  return get(url, params, {
    requestType: REQUEST_TYPES.GUEST_ALLOWED,
    ...options
  })
}

/**
 * 允许游客访问的POST请求
 * @param {string} url 请求URL
 * @param {Object} data 请求数据
 * @param {Object} options 其他选项
 */
export function guestPost(url, data = {}, options = {}) {
  return post(url, data, {
    requestType: REQUEST_TYPES.GUEST_ALLOWED,
    ...options
  })
}

/**
 * 文件上传
 * @param {string} url 上传URL
 * @param {string} filePath 文件路径
 * @param {Object} options 其他选项
 */
export function uploadFile(url, filePath, options = {}) {
  return new Promise((resolve, reject) => {
    const config = {
      url: url.startsWith('http') ? url : buildApiUrl(url),
      filePath,
      name: 'file',
      requestType: REQUEST_TYPES.AUTH_REQUIRED,
      ...options
    }

    uni.uploadFile({
      ...config,
      success: (res) => {
        try {
          const data = JSON.parse(res.data)
          resolve({
            statusCode: res.statusCode,
            data
          })
        } catch (error) {
          resolve({
            statusCode: res.statusCode,
            data: res.data
          })
        }
      },
      fail: reject
    })
  })
}

/**
 * 文件下载
 * @param {string} url 下载URL
 * @param {Object} options 其他选项
 */
export function downloadFile(url, options = {}) {
  return new Promise((resolve, reject) => {
    const config = {
      url: url.startsWith('http') ? url : buildApiUrl(url),
      ...options
    }

    uni.downloadFile({
      ...config,
      success: resolve,
      fail: reject
    })
  })
}

// 导出默认配置
export default {
  request,
  get,
  post,
  put,
  del,
  authGet,
  authPost,
  authPut,
  authDel,
  guestGet,
  guestPost,
  uploadFile,
  downloadFile
} 