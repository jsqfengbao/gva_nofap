/**
 * 社区相关API
 */
import { authGet, authPost, guestGet } from '@/utils/request.js'

/**
 * 获取帖子列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.category 分类
 * @param {string} params.sort 排序方式
 */
export function getPosts(params = {}) {
  return guestGet('/community/posts', params)
}

/**
 * 获取帖子详情
 * @param {string|number} id 帖子ID
 */
export function getPostDetail(id) {
  return guestGet(`/community/posts/${id}`)
}

/**
 * 发布帖子
 * @param {Object} data 帖子数据
 * @param {string} data.title 标题
 * @param {string} data.content 内容
 * @param {string} data.category 分类
 * @param {Array} data.tags 标签
 */
export function createPost(data) {
  return authPost('/community/posts', data)
}

/**
 * 回复帖子
 * @param {string|number} postId 帖子ID
 * @param {Object} data 回复数据
 * @param {string} data.content 回复内容
 * @param {string} data.replyToId 回复的评论ID（可选）
 */
export function replyPost(postId, data) {
  return authPost(`/community/posts/${postId}/replies`, data)
}

/**
 * 点赞帖子
 * @param {string|number} postId 帖子ID
 */
export function likePost(postId) {
  return authPost(`/community/posts/${postId}/like`)
}

/**
 * 取消点赞帖子
 * @param {string|number} postId 帖子ID
 */
export function unlikePost(postId) {
  return authPost(`/community/posts/${postId}/unlike`)
}

/**
 * 收藏帖子
 * @param {string|number} postId 帖子ID
 */
export function favoritePost(postId) {
  return authPost(`/community/posts/${postId}/favorite`)
}

/**
 * 取消收藏帖子
 * @param {string|number} postId 帖子ID
 */
export function unfavoritePost(postId) {
  return authPost(`/community/posts/${postId}/unfavorite`)
}

/**
 * 举报帖子
 * @param {string|number} postId 帖子ID
 * @param {Object} data 举报数据
 * @param {string} data.reason 举报原因
 * @param {string} data.description 详细描述
 */
export function reportPost(postId, data) {
  return authPost(`/community/posts/${postId}/report`, data)
}

/**
 * 获取我的帖子
 * @param {Object} params 查询参数
 */
export function getMyPosts(params = {}) {
  return authGet('/community/my-posts', params)
}

/**
 * 获取我的收藏
 * @param {Object} params 查询参数
 */
export function getMyFavorites(params = {}) {
  return authGet('/community/my-favorites', params)
}

/**
 * 获取社区分类
 */
export function getCategories() {
  return guestGet('/community/categories')
}

/**
 * 获取热门标签
 */
export function getPopularTags() {
  return guestGet('/community/popular-tags')
}

// 导出社区API对象
const communityApi = {
  getPosts,
  getPostDetail,
  createPost,
  replyPost,
  likePost,
  unlikePost,
  favoritePost,
  unfavoritePost,
  reportPost,
  getMyPosts,
  getMyFavorites,
  getCategories,
  getPopularTags
}

export default communityApi 