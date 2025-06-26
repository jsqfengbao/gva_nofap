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
function isLoggedIn() {
  const token = common_vendor.index.getStorageSync("token");
  return token && token !== "" && token !== "guest_token";
}
function getUserInfo() {
  try {
    const userInfo = common_vendor.index.getStorageSync("userInfo");
    return userInfo ? JSON.parse(typeof userInfo === "string" ? userInfo : JSON.stringify(userInfo)) : null;
  } catch (error) {
    console.error("获取用户信息失败:", error);
    return null;
  }
}
function getToken() {
  return common_vendor.index.getStorageSync("token") || "";
}
function setUserInfo(userInfo) {
  try {
    common_vendor.index.setStorageSync("userInfo", userInfo);
    common_vendor.index.setStorageSync("token", userInfo.token || getToken());
  } catch (error) {
    console.error("保存用户信息失败:", error);
  }
}
function clearUserInfo() {
  common_vendor.index.removeStorageSync("token");
  common_vendor.index.removeStorageSync("userInfo");
}
function logout() {
  clearUserInfo();
  common_vendor.index.reLaunch({
    url: "/pages/welcome/welcome"
  });
}
function wxLogin(encryptedData = "", iv = "") {
  return __async(this, null, function* () {
    return new Promise((resolve, reject) => {
      common_vendor.index.login({
        provider: "weixin",
        success: (loginRes) => __async(this, null, function* () {
          try {
            const requestData = {
              code: loginRes.code
            };
            if (encryptedData && iv) {
              requestData.encryptedData = encryptedData;
              requestData.iv = iv;
            }
            const res = yield common_vendor.index.request({
              url: config_env.buildApiUrl("/auth/wx-login"),
              // 使用环境配置的URL
              method: "POST",
              header: {
                "Content-Type": "application/json"
              },
              data: requestData
            });
            if (res.data.code === 0) {
              common_vendor.index.setStorageSync("token", res.data.data.token);
              common_vendor.index.setStorageSync("userInfo", res.data.data.user);
              resolve(res.data.data);
            } else {
              reject(new Error(res.data.msg || "登录失败"));
            }
          } catch (error) {
            reject(error);
          }
        }),
        fail: (err) => {
          console.error("微信登录失败详细信息:", err);
          let errorMsg = "获取登录凭证失败";
          if (err.errCode) {
            switch (err.errCode) {
              case -1:
                errorMsg = "系统错误，请稍后重试";
                break;
              case 40013:
                errorMsg = "AppID无效，请检查AppID配置";
                break;
              case 40125:
                errorMsg = "AppID未激活，请在微信公众平台激活小程序";
                break;
              case 40163:
                errorMsg = "登录态已过期，请重试";
                break;
              default:
                errorMsg = `登录失败(错误码:${err.errCode}): ${err.errMsg || "未知错误"}`;
            }
          }
          reject(new Error(errorMsg));
        }
      });
    });
  });
}
function request(options) {
  const token = getToken();
  const defaultOptions = {
    header: {
      "Content-Type": "application/json"
    },
    timeout: 3e4
  };
  const config = Object.assign({}, defaultOptions, options);
  if (token && token !== "guest_token") {
    config.header.Authorization = `Bearer ${token}`;
    config.header["x-token"] = token;
  }
  return new Promise((resolve, reject) => {
    common_vendor.index.request(__spreadProps(__spreadValues({}, config), {
      success: (res) => {
        if (res.statusCode === 401 || res.data && res.data.code === 401) {
          clearUserInfo();
          common_vendor.index.showToast({
            title: "登录已过期，请重新登录",
            icon: "none"
          });
          setTimeout(() => {
            common_vendor.index.reLaunch({
              url: "/pages/welcome/welcome"
            });
          }, 1500);
          reject(new Error("认证失败"));
          return;
        }
        resolve(res);
      },
      fail: (err) => {
        reject(err);
      }
    }));
  });
}
function getAvatarUrl(avatarUrl) {
  if (!avatarUrl) {
    return config_env.getImageUrl("default-avatar.png");
  }
  if (avatarUrl.startsWith("http")) {
    return avatarUrl;
  }
  return config_env.getImageUrl(avatarUrl);
}
exports.getAvatarUrl = getAvatarUrl;
exports.getToken = getToken;
exports.getUserInfo = getUserInfo;
exports.isLoggedIn = isLoggedIn;
exports.logout = logout;
exports.request = request;
exports.setUserInfo = setUserInfo;
exports.wxLogin = wxLogin;
