"use strict";
const utils_request = require("../utils/request.js");
function getProfile() {
  return utils_request.authGet("/user/profile");
}
function updateUserInfo(data) {
  return utils_request.authPut("/user/info", data);
}
function saveWxAvatar(data) {
  return utils_request.authPost("/user/save-wx-avatar", data);
}
function getStats() {
  return utils_request.authGet("/user/stats");
}
function updatePrivacySettings(data) {
  return utils_request.authPut("/user/privacy-settings", data);
}
function updateNotificationSettings(data) {
  return utils_request.authPut("/user/notification-settings", data);
}
function createDataExport(data) {
  return utils_request.authPost("/user/export", data);
}
function getUserSettings() {
  return utils_request.authGet("/user/settings");
}
function uploadAvatar(filePath) {
  return utils_request.uploadFile("/user/upload-avatar", filePath);
}
const userApi = {
  getProfile,
  updateUserInfo,
  saveWxAvatar,
  getStats,
  updatePrivacySettings,
  updateNotificationSettings,
  createDataExport,
  getUserSettings,
  uploadAvatar
};
exports.userApi = userApi;
