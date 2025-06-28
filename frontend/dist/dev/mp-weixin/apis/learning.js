"use strict";
const utils_request = require("../utils/request.js");
function getList(params = {}) {
  return utils_request.guestGet("/learning/list", params);
}
function getDetail(id) {
  return utils_request.guestGet(`/learning/${id}`);
}
function recordProgress(id, data) {
  return utils_request.authPost(`/learning/${id}/progress`, data);
}
function getLearningRecord(params = {}) {
  return utils_request.authGet("/learning/record", params);
}
function collectContent(id) {
  return utils_request.authPost(`/learning/${id}/collect`);
}
function uncollectContent(id) {
  return utils_request.authPost(`/learning/${id}/uncollect`);
}
function getMyCollections(params = {}) {
  return utils_request.authGet("/learning/my-collections", params);
}
function rateContent(id, data) {
  return utils_request.authPost(`/learning/${id}/rate`, data);
}
function getRecommendations(params = {}) {
  return utils_request.guestGet("/learning/recommendations", params);
}
function getCategories() {
  return utils_request.guestGet("/learning/categories");
}
function searchContent(params = {}) {
  return utils_request.guestGet("/learning/search", params);
}
const learningApi = {
  getList,
  getDetail,
  recordProgress,
  getLearningRecord,
  collectContent,
  uncollectContent,
  getMyCollections,
  rateContent,
  getRecommendations,
  getCategories,
  searchContent
};
exports.learningApi = learningApi;
