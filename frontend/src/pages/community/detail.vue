<template>
  <view class="detail-page">
    <!-- 状态栏 -->
    <view class="status-bar">
      <text>{{ currentTime }}</text>
      <view class="status-icons">
        <text>📶</text>
        <text>📶</text>
        <view class="battery">
          <view class="battery-level"></view>
        </view>
      </view>
    </view>

    <!-- 头部 -->
    <view class="header">
      <view class="header-main">
        <view class="back-btn" @tap="goBack">
          <text class="icon">←</text>
        </view>
        <text class="title">帖子详情</text>
        <view class="more-btn" @tap="showMoreOptions">
          <text class="icon">⋯</text>
        </view>
      </view>
    </view>

    <!-- 内容区域 -->
    <scroll-view scroll-y="true" class="content-container">
      <!-- 帖子内容 -->
      <view class="post-content">
        <!-- 用户信息 -->
        <view class="post-header">
          <view class="user-info">
            <view class="avatar" :class="getAvatarClass(post.userNickname)">
              <text class="avatar-text">{{ getAvatarText(post.userNickname) }}</text>
            </view>
            <view class="user-details">
              <text class="username">{{ post.userNickname }}</text>
              <view class="post-meta">
                <text class="post-time">{{ formatTime(post.createdAt) }}</text>
                <view class="category-tag" :class="getCategoryClass(post.category)">
                  <text>{{ post.categoryName }}</text>
                </view>
              </view>
            </view>
          </view>
          <view class="view-count">
            <text class="icon">👁️</text>
            <text class="count">{{ post.viewCount }}</text>
          </view>
        </view>

        <!-- 帖子标题 -->
        <view class="post-title">
          <text>{{ post.title }}</text>
        </view>

        <!-- 帖子正文 -->
        <view class="post-text">
          <text>{{ post.content }}</text>
        </view>

        <!-- 互动按钮 -->
        <view class="post-actions">
          <view class="action-group">
            <view class="action-item" @tap="toggleLike">
              <text class="icon" :class="{ liked: post.isLiked }">❤️</text>
              <text class="count">{{ post.likeCount }}</text>
            </view>
            <view class="action-item" @tap="focusCommentInput">
              <text class="icon">💬</text>
              <text class="count">{{ post.commentCount }}</text>
            </view>
            <view class="action-item" @tap="sharePost">
              <text class="icon">📤</text>
              <text>分享</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 评论列表 -->
      <view class="comments-section">
        <view class="section-header">
          <text class="section-title">评论 ({{ comments.length }})</text>
          <view class="sort-options">
            <text class="sort-item" :class="{ active: sortType === 'time' }" @tap="setSortType('time')">最新</text>
            <text class="sort-item" :class="{ active: sortType === 'hot' }" @tap="setSortType('hot')">最热</text>
          </view>
        </view>

        <view class="comments-list">
          <view class="comment-item" v-for="comment in comments" :key="comment.id">
            <!-- 评论者信息 -->
            <view class="comment-header">
              <view class="commenter-info">
                <view class="avatar" :class="getAvatarClass(comment.userNickname)">
                  <text class="avatar-text">{{ getAvatarText(comment.userNickname) }}</text>
                </view>
                <view class="commenter-details">
                  <text class="username">{{ comment.userNickname }}</text>
                  <text class="comment-time">{{ formatTime(comment.createdAt) }}</text>
                </view>
              </view>
              <view class="comment-actions">
                <view class="action-item" @tap="toggleCommentLike(comment)">
                  <text class="icon" :class="{ liked: comment.isLiked }">❤️</text>
                  <text class="count" v-if="comment.likeCount > 0">{{ comment.likeCount }}</text>
                </view>
                <view class="action-item" @tap="replyToComment(comment)">
                  <text class="icon">💬</text>
                  <text>回复</text>
                </view>
              </view>
            </view>

            <!-- 评论内容 -->
            <view class="comment-content">
              <text>{{ comment.content }}</text>
            </view>

            <!-- 子评论 -->
            <view class="sub-comments" v-if="comment.replies && comment.replies.length > 0">
              <view class="sub-comment" v-for="reply in comment.replies" :key="reply.id">
                <view class="sub-comment-header">
                  <view class="avatar small" :class="getAvatarClass(reply.userNickname)">
                    <text class="avatar-text">{{ getAvatarText(reply.userNickname) }}</text>
                  </view>
                  <text class="username">{{ reply.userNickname }}</text>
                  <text class="reply-time">{{ formatTime(reply.createdAt) }}</text>
                </view>
                <view class="sub-comment-content">
                  <text>{{ reply.content }}</text>
                </view>
              </view>
            </view>
          </view>

          <!-- 加载更多评论 -->
          <view class="load-more" v-if="hasMoreComments" @tap="loadMoreComments">
            <text>{{ loadingComments ? '加载中...' : '查看更多评论' }}</text>
          </view>
        </view>
      </view>
    </scroll-view>

    <!-- 评论输入框 -->
    <view class="comment-input-container">
      <view class="input-wrapper">
        <input
          ref="commentInput"
          class="comment-input"
          v-model="commentText"
          placeholder="写下你的评论..."
          maxlength="500"
          @confirm="submitComment"
        />
        <view class="send-btn" :class="{ active: commentText.trim() }" @tap="submitComment">
          <text>发送</text>
        </view>
      </view>
    </view>

    <!-- 加载遮罩 -->
    <view class="loading-mask" v-if="loading">
      <view class="loading-content">
        <view class="spinner"></view>
        <text>加载中...</text>
      </view>
    </view>
  </view>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: 'CommunityDetail',
  setup() {
    const currentTime = ref('9:41')
    const loading = ref(false)
    const loadingComments = ref(false)
    const hasMoreComments = ref(true)
    const commentText = ref('')
    const sortType = ref('time')
    
    // 帖子数据
    const post = ref({
      id: 1,
      title: '🎉 100天里程碑达成！',
      content: '想跟大家分享一些这段路程的心得和体验。首先，前50天是最困难的，特别是前2周，身体和心理都会有很大的反应。我的建议是：\n\n1. 建立良好的作息时间\n2. 多运动，转移注意力\n3. 寻找新的兴趣爱好\n4. 及时寻求社区支持\n\n现在100天了，感觉整个人的精神状态都有了很大改善，希望能继续坚持下去，也希望能帮助到更多的朋友！',
      category: 4,
      categoryName: '成功故事',
      userNickname: '坚持者_阳光',
      isAnonymous: false,
      viewCount: 156,
      likeCount: 128,
      commentCount: 45,
      isLiked: false,
      createdAt: new Date(Date.now() - 180000).toISOString()
    })

    // 评论数据
    const comments = ref([
      {
        id: 1,
        content: '太厉害了！100天真的不容易，请问你是怎么度过最困难的前期的？',
        userNickname: '新手_加油',
        likeCount: 12,
        isLiked: false,
        createdAt: new Date(Date.now() - 60000).toISOString(),
        replies: [
          {
            id: 11,
            content: '我也想知道这个问题的答案',
            userNickname: '同路人',
            createdAt: new Date(Date.now() - 30000).toISOString()
          }
        ]
      },
      {
        id: 2,
        content: '恭喜恭喜！你的分享很有用，收藏了！',
        userNickname: '学习者_努力',
        likeCount: 8,
        isLiked: true,
        createdAt: new Date(Date.now() - 120000).toISOString(),
        replies: []
      }
    ])

    const updateTime = () => {
      const now = new Date()
      currentTime.value = `${now.getHours()}:${now.getMinutes().toString().padStart(2, '0')}`
    }

    const goBack = () => {
      uni.navigateBack()
    }

    const showMoreOptions = () => {
      uni.showActionSheet({
        itemList: ['举报', '分享', '收藏'],
        success: (res) => {
          console.log('选中了第' + (res.tapIndex + 1) + '个按钮')
        }
      })
    }

    const toggleLike = () => {
      post.value.isLiked = !post.value.isLiked
      post.value.likeCount += post.value.isLiked ? 1 : -1
      uni.showToast({
        title: post.value.isLiked ? '点赞成功' : '取消点赞',
        icon: 'success'
      })
    }

    const toggleCommentLike = (comment) => {
      comment.isLiked = !comment.isLiked
      comment.likeCount += comment.isLiked ? 1 : -1
    }

    const focusCommentInput = () => {
      // 聚焦到评论输入框
      uni.showToast({
        title: '评论功能开发中',
        icon: 'none'
      })
    }

    const sharePost = () => {
      uni.showToast({
        title: '分享功能开发中',
        icon: 'none'
      })
    }

    const replyToComment = (comment) => {
      commentText.value = `@${comment.userNickname} `
    }

    const submitComment = () => {
      if (!commentText.value.trim()) return
      
      uni.showToast({
        title: '评论发布成功',
        icon: 'success'
      })
      commentText.value = ''
    }

    const setSortType = (type) => {
      sortType.value = type
    }

    const loadMoreComments = () => {
      loadingComments.value = true
      setTimeout(() => {
        loadingComments.value = false
        hasMoreComments.value = false
      }, 1000)
    }

    // 工具方法
    const getAvatarClass = (nickname) => {
      if (nickname === '匿名用户') return 'anonymous'
      const hash = nickname.charCodeAt(0) % 5
      return `avatar-${hash}`
    }

    const getAvatarText = (nickname) => {
      if (nickname === '匿名用户') return '匿'
      return nickname.substring(0, 1)
    }

    const getCategoryClass = (category) => {
      const classes = ['', 'experience', 'help', 'checkin', 'success']
      return classes[category] || ''
    }

    const formatTime = (dateStr) => {
      const date = new Date(dateStr)
      const now = new Date()
      const diff = now - date
      
      if (diff < 60000) {
        return '刚刚'
      } else if (diff < 3600000) {
        return `${Math.floor(diff / 60000)}分钟前`
      } else if (diff < 86400000) {
        return `${Math.floor(diff / 3600000)}小时前`
      } else {
        return date.toLocaleDateString()
      }
    }

    onMounted(() => {
      updateTime()
      setInterval(updateTime, 60000)
      
      // 增加浏览量
      post.value.viewCount++
    })

    return {
      currentTime, loading, loadingComments, hasMoreComments, commentText, sortType,
      post, comments, goBack, showMoreOptions, toggleLike, toggleCommentLike,
      focusCommentInput, sharePost, replyToComment, submitComment, setSortType,
      loadMoreComments, getAvatarClass, getAvatarText, getCategoryClass, formatTime
    }
  }
}
</script>

<style scoped>
.detail-page {
  min-height: 100vh;
  background: #F8FAFC;
  display: flex;
  flex-direction: column;
}

.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
  font-size: 12px;
  color: #1F2937;
  background: #FFFFFF;
}

.status-icons {
  display: flex;
  align-items: center;
  gap: 4px;
}

.battery {
  width: 24px;
  height: 12px;
  border: 1px solid #1F2937;
  border-radius: 2px;
  position: relative;
}

.battery-level {
  width: 16px;
  height: 6px;
  background: #10B981;
  border-radius: 1px;
  margin: 2px;
}

.header {
  background: #FFFFFF;
  border-bottom: 1px solid #E5E7EB;
}

.header-main {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
}

.back-btn, .more-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: #F3F4F6;
}

.back-btn .icon, .more-btn .icon {
  font-size: 18px;
  color: #6B7280;
}

.title {
  font-size: 18px;
  font-weight: 600;
  color: #1F2937;
}

.content-container {
  flex: 1;
  height: calc(100vh - 180px);
}

.post-content {
  background: #FFFFFF;
  padding: 16px;
  margin-bottom: 12px;
}

.post-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 16px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 16px;
  color: #FFFFFF;
}

.avatar.small {
  width: 32px;
  height: 32px;
  font-size: 12px;
}

.avatar.anonymous { background: linear-gradient(135deg, #6B7280, #9CA3AF); }
.avatar.avatar-0 { background: linear-gradient(135deg, #F59E0B, #F97316); }
.avatar.avatar-1 { background: linear-gradient(135deg, #3B82F6, #8B5CF6); }
.avatar.avatar-2 { background: linear-gradient(135deg, #10B981, #22D3AA); }
.avatar.avatar-3 { background: linear-gradient(135deg, #EF4444, #F97316); }
.avatar.avatar-4 { background: linear-gradient(135deg, #8B5CF6, #EC4899); }

.user-details {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.username {
  font-weight: 600;
  color: #1F2937;
  font-size: 16px;
}

.post-meta {
  display: flex;
  align-items: center;
  gap: 8px;
}

.post-time {
  font-size: 12px;
  color: #6B7280;
}

.category-tag {
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 10px;
  font-weight: 500;
}

.category-tag.experience { background: rgba(59, 130, 246, 0.1); color: #3B82F6; }
.category-tag.help { background: rgba(239, 68, 68, 0.1); color: #EF4444; }
.category-tag.checkin { background: rgba(16, 185, 129, 0.1); color: #10B981; }
.category-tag.success { background: rgba(245, 158, 11, 0.1); color: #F59E0B; }

.view-count {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #6B7280;
}

.post-title {
  font-size: 20px;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 12px;
  line-height: 1.4;
}

.post-text {
  font-size: 16px;
  color: #4B5563;
  line-height: 1.6;
  margin-bottom: 20px;
  white-space: pre-wrap;
}

.post-actions {
  border-top: 1px solid #F3F4F6;
  padding-top: 16px;
}

.action-group {
  display: flex;
  justify-content: space-around;
}

.action-item {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 14px;
  color: #6B7280;
  transition: all 0.3s;
}

.action-item:active {
  background: #F3F4F6;
}

.action-item .icon {
  font-size: 16px;
}

.action-item .icon.liked {
  color: #EF4444;
}

.comments-section {
  background: #FFFFFF;
  padding: 16px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  color: #1F2937;
}

.sort-options {
  display: flex;
  gap: 16px;
}

.sort-item {
  font-size: 14px;
  color: #6B7280;
  padding: 4px 8px;
  border-radius: 6px;
}

.sort-item.active {
  color: #22D3AA;
  background: rgba(34, 211, 170, 0.1);
}

.comment-item {
  padding: 16px 0;
  border-bottom: 1px solid #F3F4F6;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.commenter-info {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.commenter-details {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.comment-time {
  font-size: 12px;
  color: #9CA3AF;
}

.comment-actions {
  display: flex;
  gap: 16px;
}

.comment-content {
  font-size: 15px;
  color: #374151;
  line-height: 1.5;
  margin-left: 44px;
}

.sub-comments {
  margin-left: 44px;
  margin-top: 12px;
  background: #F9FAFB;
  border-radius: 8px;
  padding: 12px;
}

.sub-comment {
  margin-bottom: 12px;
}

.sub-comment:last-child {
  margin-bottom: 0;
}

.sub-comment-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 6px;
}

.reply-time {
  font-size: 11px;
  color: #9CA3AF;
}

.sub-comment-content {
  font-size: 14px;
  color: #374151;
  line-height: 1.4;
}

.load-more {
  text-align: center;
  padding: 16px;
  color: #6B7280;
  font-size: 14px;
}

.comment-input-container {
  background: #FFFFFF;
  border-top: 1px solid #E5E7EB;
  padding: 12px 16px;
}

.input-wrapper {
  display: flex;
  align-items: center;
  gap: 12px;
}

.comment-input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #E5E7EB;
  border-radius: 20px;
  font-size: 14px;
  background: #F9FAFB;
}

.send-btn {
  padding: 8px 16px;
  background: #D1D5DB;
  color: #9CA3AF;
  border-radius: 16px;
  font-size: 14px;
  font-weight: 500;
}

.send-btn.active {
  background: #22D3AA;
  color: #FFFFFF;
}

.loading-mask {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.loading-content {
  background: #FFFFFF;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.spinner {
  width: 32px;
  height: 32px;
  border: 3px solid #E5E7EB;
  border-top: 3px solid #22D3AA;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
