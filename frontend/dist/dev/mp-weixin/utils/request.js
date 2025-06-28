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
const common_vendor = require("../common/vendor.js");
const interceptors_request = require("../interceptors/request.js");
const config_env = require("../config/env.js");
function request(options = {}) {
  const config = __spreadValues({
    method: "GET",
    requestType: interceptors_request.REQUEST_TYPES.NORMAL,
    hideErrorToast: false
  }, options);
  if (!config.url.startsWith("http")) {
    config.url = config_env.buildApiUrl(config.url);
  }
  return interceptors_request.enhancedRequest(config);
}
function get(url, params = {}, options = {}) {
  return request(__spreadValues({
    url,
    method: "GET",
    query: params
  }, options));
}
function post(url, data = {}, options = {}) {
  return request(__spreadValues({
    url,
    method: "POST",
    data
  }, options));
}
function put(url, data = {}, options = {}) {
  return request(__spreadValues({
    url,
    method: "PUT",
    data
  }, options));
}
function authGet(url, params = {}, options = {}) {
  return get(url, params, __spreadValues({
    requestType: interceptors_request.REQUEST_TYPES.AUTH_REQUIRED
  }, options));
}
function authPost(url, data = {}, options = {}) {
  return post(url, data, __spreadValues({
    requestType: interceptors_request.REQUEST_TYPES.AUTH_REQUIRED
  }, options));
}
function authPut(url, data = {}, options = {}) {
  return put(url, data, __spreadValues({
    requestType: interceptors_request.REQUEST_TYPES.AUTH_REQUIRED
  }, options));
}
function guestGet(url, params = {}, options = {}) {
  return get(url, params, __spreadValues({
    requestType: interceptors_request.REQUEST_TYPES.GUEST_ALLOWED
  }, options));
}
function uploadFile(url, filePath, options = {}) {
  return new Promise((resolve, reject) => {
    const config = __spreadValues({
      url: url.startsWith("http") ? url : config_env.buildApiUrl(url),
      filePath,
      name: "file",
      requestType: interceptors_request.REQUEST_TYPES.AUTH_REQUIRED
    }, options);
    common_vendor.index.uploadFile(__spreadProps(__spreadValues({}, config), {
      success: (res) => {
        try {
          const data = JSON.parse(res.data);
          resolve({
            statusCode: res.statusCode,
            data
          });
        } catch (error) {
          resolve({
            statusCode: res.statusCode,
            data: res.data
          });
        }
      },
      fail: reject
    }));
  });
}
exports.authGet = authGet;
exports.authPost = authPost;
exports.authPut = authPut;
exports.guestGet = guestGet;
exports.uploadFile = uploadFile;
