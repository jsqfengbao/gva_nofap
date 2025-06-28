"use strict";
var __defProp = Object.defineProperty;
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
const utils_auth = require("../utils/auth.js");
const config_env = require("../config/env.js");
const learningApi = {
  // 获取学习统计
  getStats: () => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/learning/stats"),
      method: "GET"
    });
  },
  // 获取学习内容列表
  getContents: (params = {}) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/learning/contents"),
      method: "GET",
      data: params
    });
  },
  // 获取学习内容列表 (兼容旧接口)
  getList: (params = {}) => {
    return learningApi.getContents(params);
  },
  // 获取学习内容详情
  getDetail: (id) => {
    return utils_auth.request({
      url: config_env.buildApiUrl(`/learning/${id}`),
      method: "GET"
    });
  },
  // 开始学习记录
  startLearning: (contentId) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/learning/start"),
      method: "POST",
      data: { contentId }
    });
  },
  // 记录学习进度
  recordProgress: (id, data) => {
    return utils_auth.request({
      url: config_env.buildApiUrl(`/learning/${id}/progress`),
      method: "POST",
      data
    });
  },
  // 完成学习
  completeLearning: (contentId, data = {}) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/learning/complete"),
      method: "POST",
      data: __spreadValues({ contentId }, data)
    });
  },
  // 点赞内容
  likeContent: (contentId) => {
    return utils_auth.request({
      url: config_env.buildApiUrl(`/learning/${contentId}/like`),
      method: "POST"
    });
  },
  // 收藏内容
  collectContent: (contentId) => {
    return utils_auth.request({
      url: config_env.buildApiUrl(`/learning/${contentId}/collect`),
      method: "POST"
    });
  },
  // 搜索内容
  searchContents: (params = {}) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/learning/search"),
      method: "GET",
      data: params
    });
  },
  // 获取推荐内容
  getRecommendations: (params = {}) => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/learning/recommendations"),
      method: "GET",
      data: params
    });
  },
  // 获取分类统计
  getCategoryStats: () => {
    return utils_auth.request({
      url: config_env.buildApiUrl("/learning/category-stats"),
      method: "GET"
    });
  }
};
exports.learningApi = learningApi;
