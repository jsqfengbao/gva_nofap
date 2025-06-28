/**
 * 社区相关 API
 */
import { request } from '@/utils/auth'
import { buildApiUrl } from '@/config/env.js'

export const communityApi = {
  // 获取帖子列表
  getPosts: (params = {}) => {
    return request({
      url: buildApiUrl('/community/posts'),
      method: 'GET',
      data: params
    })
  },

  // 获取帖子详情
  getPostDetail: (id) => {
    return request({
      url: buildApiUrl(`/community/posts/${id}`),
      method: 'GET'
    })
  },

  // 发布帖子
  createPost: (data) => {
    return request({
      url: buildApiUrl('/community/posts'),
      method: 'POST',
      data
    })
  },

  // 回复帖子
  replyPost: (postId, data) => {
    return request({
      url: buildApiUrl(`/community/posts/${postId}/replies`),
      method: 'POST',
      data
    })
  }
} 