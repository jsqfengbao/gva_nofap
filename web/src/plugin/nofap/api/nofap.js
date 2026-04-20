// NOPEMON 戒色小程序管理端 API
import request from "@/utils/request";

// 用户管理
export function getNofapUserList(params) {
  return request({
    url: "/v1/miniprogram/admin/users",
    method: "get",
    params,
  });
}

export function getNofapUserDetail(id) {
  return request({
    url: `/v1/miniprogram/admin/user/${id}`,
    method: "get",
  });
}

export function updateNofapUserStatus(id, data) {
  return request({
    url: `/v1/miniprogram/admin/user/${id}/status`,
    method: "put",
    data,
  });
}

// 统计数据
export function getNofapStatistics() {
  return request({
    url: "/v1/miniprogram/admin/statistics",
    method: "get",
  });
}

// 内容管理 - 学习内容
export function getLearningContentList(params) {
  return request({
    url: "/v1/miniprogram/admin/learning",
    method: "get",
    params,
  });
}

export function createLearningContent(data) {
  return request({
    url: "/v1/miniprogram/admin/learning",
    method: "post",
    data,
  });
}

export function updateLearningContent(data) {
  return request({
    url: "/v1/miniprogram/admin/learning",
    method: "put",
    data,
  });
}

export function deleteLearningContent(id) {
  return request({
    url: `/v1/miniprogram/admin/learning/${id}`,
    method: "delete",
  });
}

// 紧急求助资源管理
export function getEmergencyResourceList(params) {
  return request({
    url: "/v1/miniprogram/admin/resources",
    method: "get",
    params,
  });
}

export function createEmergencyResource(data) {
  return request({
    url: "/v1/miniprogram/admin/resources",
    method: "post",
    data,
  });
}

export function updateEmergencyResource(data) {
  return request({
    url: "/v1/miniprogram/admin/resources",
    method: "put",
    data,
  });
}

export function deleteEmergencyResource(id) {
  return request({
    url: `/v1/miniprogram/admin/resources/${id}`,
    method: "delete",
  });
}

// 社区管理 - 帖子
export function getCommunityPostList(params) {
  return request({
    url: "/v1/miniprogram/admin/posts",
    method: "get",
    params,
  });
}

export function updateCommunityPostStatus(id, data) {
  return request({
    url: `/v1/miniprogram/admin/posts/${id}/status`,
    method: "put",
    data,
  });
}

export function deleteCommunityPost(id) {
  return request({
    url: `/v1/miniprogram/admin/posts/${id}`,
    method: "delete",
  });
}

// 社区管理 - 评论
export function getCommunityCommentList(params) {
  return request({
    url: "/v1/miniprogram/admin/comments",
    method: "get",
    params,
  });
}

export function updateCommunityCommentStatus(id, data) {
  return request({
    url: `/v1/miniprogram/admin/comments/${id}/status`,
    method: "put",
    data,
  });
}

export function deleteCommunityComment(id) {
  return request({
    url: `/v1/miniprogram/admin/comments/${id}`,
    method: "delete",
  });
}

// 成就管理
export function getAchievementList(params) {
  return request({
    url: "/v1/miniprogram/admin/achievements",
    method: "get",
    params,
  });
}

export function createAchievement(data) {
  return request({
    url: "/v1/miniprogram/admin/achievements",
    method: "post",
    data,
  });
}

export function updateAchievement(data) {
  return request({
    url: "/v1/miniprogram/admin/achievements",
    method: "put",
    data,
  });
}

export function deleteAchievement(id) {
  return request({
    url: `/v1/miniprogram/admin/achievements/${id}`,
    method: "delete",
  });
}
