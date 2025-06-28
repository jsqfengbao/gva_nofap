"use strict";
const utils_request = require("../utils/request.js");
function dailyCheckin(data) {
  return utils_request.authPost("/checkin/daily", data);
}
function getTodayStatus() {
  return utils_request.authGet("/checkin/today");
}
function getHistory(params = {}) {
  return utils_request.authGet("/checkin/history", params);
}
function getStatistics(params = {}) {
  return utils_request.authGet("/checkin/statistics", params);
}
function getWeeklyProgress() {
  return utils_request.authGet("/checkin/weekly-progress");
}
function getCalendarData(params = {}) {
  return utils_request.authGet("/checkin/calendar", params);
}
function getStreakRecord() {
  return utils_request.authGet("/checkin/streak");
}
function makeupCheckin(data) {
  return utils_request.authPost("/checkin/makeup", data);
}
function getLeaderboard(params = {}) {
  return utils_request.guestGet("/checkin/leaderboard", params);
}
const checkinApi = {
  dailyCheckin,
  getTodayStatus,
  getHistory,
  getStatistics,
  getWeeklyProgress,
  getCalendarData,
  getStreakRecord,
  makeupCheckin,
  getLeaderboard,
  // 兼容旧的方法名
  checkin: dailyCheckin,
  getStats: getStatistics
};
exports.checkinApi = checkinApi;
