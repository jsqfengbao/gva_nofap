"use strict";
var __defProp = Object.defineProperty;
var __defProps = Object.defineProperties;
var __getOwnPropDescs = Object.getOwnPropertyDescriptors;
var __getOwnPropSymbols = Object.getOwnPropertySymbols;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __propIsEnum = Object.prototype.propertyIsEnumerable;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp.call(b, prop))
      __defNormalProp(a, prop, b[prop]);
  if (__getOwnPropSymbols)
    for (var prop of __getOwnPropSymbols(b)) {
      if (__propIsEnum.call(b, prop))
        __defNormalProp(a, prop, b[prop]);
    }
  return a;
};
var __spreadProps = (a, b) => __defProps(a, __getOwnPropDescs(b));
var __async = (__this, __arguments, generator) => {
  return new Promise((resolve, reject) => {
    var fulfilled = (value) => {
      try {
        step(generator.next(value));
      } catch (e) {
        reject(e);
      }
    };
    var rejected = (value) => {
      try {
        step(generator.throw(value));
      } catch (e) {
        reject(e);
      }
    };
    var step = (x) => x.done ? resolve(x.value) : Promise.resolve(x.value).then(fulfilled, rejected);
    step((generator = generator.apply(__this, __arguments)).next());
  });
};
const common_vendor = require("../common/vendor.js");
const config_env = require("../config/env.js");
const utils_auth = require("../utils/auth.js");
const REQUEST_TYPES = {
  NORMAL: "normal",
  // 普通请求
  AUTH_REQUIRED: "auth",
  // 需要认证的请求
  GUEST_ALLOWED: "guest"
  // 允许游客访问的请求
};
const httpInterceptor = {
  invoke(options) {
    return __async(this, null, function* () {
      try {
        if (options.query) {
          const queryStr = Object.keys(options.query).filter((key) => options.query[key] !== void 0 && options.query[key] !== null).map((key) => `${encodeURIComponent(key)}=${encodeURIComponent(options.query[key])}`).join("&");
          if (queryStr) {
            if (options.url.includes("?")) {
              options.url += `&${queryStr}`;
            } else {
              options.url += `?${queryStr}`;
            }
          }
        }
        if (!options.url.startsWith("http")) {
          options.url = config_env.buildApiUrl(options.url);
        }
        options.timeout = options.timeout || 3e4;
        options.header = __spreadValues({
          "Content-Type": "application/json",
          "platform": getplatform()
        }, options.header);
        yield handleAuthentication(options);
        const requestId = generateRequestId(options);
        options.requestId = requestId;
        console.log(`📤 API请求 [${requestId}]:`, {
          url: options.url,
          method: options.method || "GET",
          hasAuth: !!options.header.Authorization
        });
      } catch (error) {
        console.error("❌ 请求拦截器处理失败:", error);
        throw error;
      }
    });
  }
};
function handleAuthentication(options) {
  return __async(this, null, function* () {
    const requestType = options.requestType || REQUEST_TYPES.NORMAL;
    const token = utils_auth.getToken();
    if (requestType === REQUEST_TYPES.AUTH_REQUIRED && !token) {
      throw new Error("需要登录后才能访问");
    }
    if (token && token !== "guest_token") {
      if (isValidJWT(token)) {
        options.header.Authorization = `Bearer ${token}`;
        options.header["x-token"] = token;
      } else {
        console.warn("⚠️ 无效的token格式，清除token");
        utils_auth.clearUserInfo();
        if (requestType === REQUEST_TYPES.AUTH_REQUIRED) {
          throw new Error("登录信息已失效，请重新登录");
        }
      }
    }
  });
}
function isValidJWT(token) {
  if (!token || typeof token !== "string") return false;
  const parts = token.split(".");
  if (parts.length !== 3) return false;
  try {
    for (let part of parts) {
      while (part.length % 4) {
        part += "=";
      }
      atob(part.replace(/-/g, "+").replace(/_/g, "/"));
    }
    return true;
  } catch (e) {
    return false;
  }
}
function getplatform() {
  return "mp-weixin";
}
function generateRequestId(options) {
  const method = options.method || "GET";
  const url = options.url;
  const timestamp = Date.now();
  return `${method}_${url.split("?")[0].split("/").pop()}_${timestamp}`;
}
const responseInterceptor = {
  success: (response, options) => __async(exports, null, function* () {
    const requestId = (options == null ? void 0 : options.requestId) || "unknown";
    console.log(`📥 API响应 [${requestId}]:`, {
      status: response.statusCode,
      success: response.statusCode === 200
    });
    if (isAuthError(response)) {
      return handleAuthError(response, options);
    }
    if (isBusinessError(response)) {
      return handleBusinessError(response, options);
    }
    return response;
  }),
  fail: (error, options) => __async(exports, null, function* () {
    const requestId = (options == null ? void 0 : options.requestId) || "unknown";
    console.error(`❌ API请求失败 [${requestId}]:`, error);
    if (isNetworkError(error)) {
      return handleNetworkError(error, options);
    }
    if (isTimeoutError(error)) {
      return handleTimeoutError(error, options);
    }
    throw error;
  })
};
function isAuthError(response) {
  var _a, _b, _c, _d;
  return response.statusCode === 401 || ((_a = response.data) == null ? void 0 : _a.code) === 401 || ((_b = response.data) == null ? void 0 : _b.code) === 7 && ((_d = (_c = response.data) == null ? void 0 : _c.msg) == null ? void 0 : _d.includes("token"));
}
function handleAuthError(response, options) {
  return __async(this, null, function* () {
    console.log("🔒 检测到认证错误，处理中...");
    utils_auth.clearUserInfo();
    if ((options == null ? void 0 : options.requestType) !== REQUEST_TYPES.AUTH_REQUIRED) {
      return response;
    }
    common_vendor.index.showToast({
      title: "登录已过期，请重新登录",
      icon: "none",
      duration: 2e3
    });
    setTimeout(() => {
      common_vendor.index.reLaunch({
        url: "/pages/welcome/welcome"
      });
    }, 2e3);
    throw new Error("认证失败");
  });
}
function isBusinessError(response) {
  var _a;
  return response.statusCode === 200 && ((_a = response.data) == null ? void 0 : _a.code) !== 0;
}
function handleBusinessError(response, options) {
  var _a;
  const errorMsg = ((_a = response.data) == null ? void 0 : _a.msg) || "请求失败";
  if (!(options == null ? void 0 : options.hideErrorToast)) {
    common_vendor.index.showToast({
      title: errorMsg,
      icon: "none",
      duration: 2e3
    });
  }
  return response;
}
function isNetworkError(error) {
  var _a, _b;
  return ((_a = error.errMsg) == null ? void 0 : _a.includes("fail")) && !((_b = error.errMsg) == null ? void 0 : _b.includes("timeout"));
}
function handleNetworkError(error, options) {
  console.error("🌐 网络错误:", error);
  if (!(options == null ? void 0 : options.hideErrorToast)) {
    common_vendor.index.showToast({
      title: "网络连接失败，请检查网络设置",
      icon: "none",
      duration: 3e3
    });
  }
  throw error;
}
function isTimeoutError(error) {
  var _a;
  return (_a = error.errMsg) == null ? void 0 : _a.includes("timeout");
}
function handleTimeoutError(error, options) {
  console.error("⏰ 请求超时:", error);
  if (!(options == null ? void 0 : options.hideErrorToast)) {
    common_vendor.index.showToast({
      title: "请求超时，请稍后重试",
      icon: "none",
      duration: 2e3
    });
  }
  throw error;
}
function enhancedRequest(options) {
  return new Promise((resolve, reject) => __async(this, null, function* () {
    try {
      yield httpInterceptor.invoke(options);
      common_vendor.index.request(__spreadProps(__spreadValues({}, options), {
        success: (response) => __async(this, null, function* () {
          try {
            const processedResponse = yield responseInterceptor.success(response, options);
            resolve(processedResponse);
          } catch (error) {
            reject(error);
          }
        }),
        fail: (error) => __async(this, null, function* () {
          try {
            yield responseInterceptor.fail(error, options);
          } catch (processedError) {
            reject(processedError);
          }
        })
      }));
    } catch (error) {
      reject(error);
    }
  }));
}
exports.REQUEST_TYPES = REQUEST_TYPES;
exports.enhancedRequest = enhancedRequest;
