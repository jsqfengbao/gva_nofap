"use strict";
const utils_auth = require("../utils/auth.js");
const config_env = require("../config/env.js");
const userApi = {
  // 获取用户资料
  getProfile: () => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/user/profile"),
      method: "GET"
    });
  },
  // 更新用户信息
  updateUserInfo: (data) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/user/info"),
      method: "PUT",
      data
    });
  },
  // 保存微信头像
  saveWxAvatar: (data) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/user/save-wx-avatar"),
      method: "POST",
      data
    });
  },
  // 获取用户统计数据
  getStats: () => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/user/stats"),
      method: "GET"
    });
  },
  // 更新隐私设置
  updatePrivacySettings: (data) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/user/privacy-settings"),
      method: "PUT",
      data
    });
  },
  // 更新通知设置
  updateNotificationSettings: (data) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/user/notification-settings"),
      method: "PUT",
      data
    });
  },
  // 创建数据导出
  createDataExport: (data) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/user/export"),
      method: "POST",
      data
    });
  },
  // 获取用户设置
  getUserSettings: () => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/user/settings"),
      method: "GET"
    });
  }
};
exports.userApi = userApi;
