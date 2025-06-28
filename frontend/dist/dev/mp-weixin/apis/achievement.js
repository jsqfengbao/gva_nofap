"use strict";
const utils_auth = require("../utils/auth.js");
const config_env = require("../config/env.js");
const achievementApi = {
  // 获取成就列表
  getList: (params = {}) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/achievement/list"),
      method: "GET",
      data: params
    });
  },
  // 获取用户成就
  getUserAchievements: (params = {}) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/achievement/user"),
      method: "GET",
      data: params
    });
  },
  // 获取游戏化统计数据（首页使用）
  getGameStats: () => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/achievement/game-stats"),
      method: "GET"
    });
  }
};
exports.achievementApi = achievementApi;
