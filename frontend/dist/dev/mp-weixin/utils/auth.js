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
  const token = common_vendor.index.getStorageSync("token") || "";
  if (token && token !== "guest_token" && !isValidJWT(token)) {
    console.warn("检测到无效token，自动清除:", token);
    clearUserInfo();
    return "";
  }
  return token;
}
function isValidJWT(token) {
  if (!token || typeof token !== "string") {
    return false;
  }
  const parts = token.split(".");
  if (parts.length !== 3) {
    return false;
  }
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
function checkWechatConfig() {
  const config = config_env.getCurrentConfig();
  const appId = config.thirdParty.WECHAT_APP_ID;
  console.log("当前微信小程序配置:", {
    appId,
    env: config.env
  });
  return true;
}
function wxLogin(encryptedData = "", iv = "", retryCount = 0) {
  return __async(this, null, function* () {
    return new Promise((resolve, reject) => {
      console.log(`开始微信登录流程... (尝试次数: ${retryCount + 1})`);
      if (!checkWechatConfig()) {
        reject(new Error("微信小程序配置错误，请联系开发者"));
        return;
      }
      common_vendor.index.login({
        provider: "weixin",
        success: (loginRes) => __async(this, null, function* () {
          var _a;
          try {
            console.log("微信登录成功，获取到code:", loginRes.code);
            console.log("code长度:", (_a = loginRes.code) == null ? void 0 : _a.length);
            console.log("完整登录响应:", loginRes);
            const startTime = Date.now();
            const requestData = {
              code: loginRes.code
            };
            if (encryptedData && iv) {
              requestData.encryptedData = encryptedData;
              requestData.iv = iv;
              console.log("包含用户加密信息进行登录");
            } else {
              console.log("基础登录，仅使用code");
            }
            console.log("准备发送登录请求，请求数据:", requestData);
            console.log("请求URL:", config_env.buildApiUrl("/auth/wx-login"));
            const res = yield common_vendor.index.request({
              url: config_env.buildApiUrl("/auth/wx-login"),
              method: "POST",
              header: {
                "Content-Type": "application/json"
              },
              data: requestData,
              timeout: 2e4
            });
            const endTime = Date.now();
            console.log(`登录请求耗时: ${endTime - startTime}ms`);
            console.log("登录API完整响应:", res);
            console.log("响应状态码:", res.statusCode);
            console.log("响应数据:", res.data);
            if (res.statusCode !== 200) {
              throw new Error(`HTTP错误: ${res.statusCode}`);
            }
            if (res.data.code === 0) {
              console.log("登录成功，保存用户信息");
              common_vendor.index.setStorageSync("token", res.data.data.token);
              if (res.data.data.user) {
                common_vendor.index.setStorageSync("userInfo", res.data.data.user);
              }
              resolve(res.data.data);
            } else {
              console.error("登录失败，服务器返回错误:", res.data);
              if (res.data.msg && res.data.msg.includes("invalid code")) {
                console.log("检测到code错误，尝试重新登录...");
                setTimeout(() => {
                  wxLogin(encryptedData, iv).then(resolve).catch(reject);
                }, 1e3);
                return;
              }
              reject(new Error(res.data.msg || "登录失败"));
            }
          } catch (error) {
            console.error("登录请求异常:", error);
            if (error.errMsg && error.errMsg.includes("timeout")) {
              reject(new Error("网络超时，请检查网络连接"));
            } else if (error.errMsg && error.errMsg.includes("fail")) {
              reject(new Error("网络请求失败，请稍后重试"));
            } else {
              reject(error);
            }
          }
        }),
        fail: (err) => {
          console.error("微信登录失败详细信息:", err);
          let errorMsg = "获取登录凭证失败";
          if (err.errMsg) {
            console.log("微信登录错误信息:", err.errMsg);
            if (err.errMsg.includes("access_token missing")) {
              errorMsg = "AppID配置错误或小程序未激活，请联系开发者检查微信小程序配置";
            } else if (err.errMsg.includes("需要重新登录")) {
              errorMsg = "登录凭证已失效，请重新尝试登录";
            } else if (err.errMsg.includes("invalid code")) {
              errorMsg = "登录凭证无效，请重新尝试";
            } else if (err.errMsg.includes("timeout")) {
              errorMsg = "网络超时，请检查网络连接后重试";
            }
          }
          if (err.errCode) {
            console.log("微信登录错误码:", err.errCode);
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
    if (token.split(".").length === 3) {
      config.header.Authorization = `Bearer ${token}`;
      config.header["x-token"] = token;
    } else {
      console.warn("❌ 无效的token格式，跳过认证:", token);
      clearUserInfo();
    }
  }
  return new Promise((resolve, reject) => {
    common_vendor.index.request(__spreadProps(__spreadValues({}, config), {
      success: (res) => {
        var _a;
        if (res.statusCode === 401 || res.data && res.data.code === 401 || res.data && res.data.code === 7 && ((_a = res.data.msg) == null ? void 0 : _a.includes("token"))) {
          console.error("🔒 认证失败，清除token并重新登录");
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
        console.error("❌ API请求失败:", err.errMsg);
        reject(err);
      }
    }));
  });
}
function getAvatarUrl(avatarUrl) {
  if (!avatarUrl) {
    return null;
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
