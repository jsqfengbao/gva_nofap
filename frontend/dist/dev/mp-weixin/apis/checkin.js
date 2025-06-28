"use strict";
const utils_auth = require("../utils/auth.js");
const config_env = require("../config/env.js");
const checkinApi = {
  // 今日打卡
  checkin: (data) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/checkin"),
      method: "POST",
      data
    });
  },
  // 获取打卡历史
  getHistory: (params = {}) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/checkin/history"),
      method: "GET",
      data: params
    });
  },
  // 获取打卡统计
  getStats: (params = {}) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/checkin/stats"),
      method: "GET",
      data: params
    });
  },
  // 获取打卡统计数据（首页使用）
  getStatistics: () => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/checkin/statistics"),
      method: "GET"
    });
  },
  // 获取今日打卡状态（首页使用）
  getTodayStatus: () => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/checkin/today"),
      method: "GET"
    });
  }
};
exports.checkinApi = checkinApi;
