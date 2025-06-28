"use strict";
const utils_request = require("../utils/request.js");
function getList(params = {}) {
  return utils_request.guestGet("/achievement/list", params);
}
function getUserAchievements(params = {}) {
  return utils_request.authGet("/achievement/user", params);
}
function getAchievementDetail(id) {
  return utils_request.guestGet(`/achievement/${id}`);
}
function getAchievementProgress(id) {
  return utils_request.authGet(`/achievement/${id}/progress`);
}
function getGameStats() {
  return utils_request.authGet("/achievement/game-stats");
}
function getLeaderboard(params = {}) {
  return utils_request.guestGet("/achievement/leaderboard", params);
}
function getCategories() {
  return utils_request.guestGet("/achievement/categories");
}
function getUserLevel() {
  return utils_request.authGet("/achievement/user-level");
}
function getExpHistory(params = {}) {
  return utils_request.authGet("/achievement/exp-history", params);
}
const achievementApi = {
  getList,
  getUserAchievements,
  getAchievementDetail,
  getAchievementProgress,
  getGameStats,
  getLeaderboard,
  getCategories,
  getUserLevel,
  getExpHistory
};
exports.achievementApi = achievementApi;
