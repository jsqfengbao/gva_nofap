<template>
  <view class="community-page">
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
        <text class="title">社区互助</text>
        <view class="header-actions">
          <view class="action-btn search-btn" @tap="openSearch">
            <text>🔍</text>
          </view>
          <view class="action-btn post-btn" @tap="openPostModal">
            <text>✏️</text>
          </view>
        </view>
      </view>

      <!-- 分类标签 -->
      <scroll-view scroll-x="true" class="category-tabs">
        <view class="tab-item" 
              :class="{ active: selectedCategory === 0 }" 
              @tap="selectCategory(0)">
          <text>全部</text>
        </view>
        <view class="tab-item" 
              :class="{ active: selectedCategory === 1 }" 
              @tap="selectCategory(1)">
          <text>经验分享</text>
        </view>
        <view class="tab-item" 
              :class="{ active: selectedCategory === 2 }" 
              @tap="selectCategory(2)">
          <text>求助求鼓励</text>
        </view>
        <view class="tab-item" 
              :class="{ active: selectedCategory === 3 }" 
              @tap="selectCategory(3)">
          <text>日常打卡</text>
        </view>
        <view class="tab-item" 
              :class="{ active: selectedCategory === 4 }" 
              @tap="selectCategory(4)">
          <text>成功故事</text>
        </view>
      </scroll-view>
    </view>

    <!-- 帖子列表 -->
    <scroll-view scroll-y="true" class="posts-container" @scrolltolower="loadMorePosts">
      <view class="posts-list">
        <view class="post-item" v-for="post in posts" :key="post.id" @tap="viewPostDetail(post.id)">
          <!-- 用户信息 -->
          <view class="post-header">
            <view class="user-info">
              <view class="avatar" :class="getAvatarClass(post.userNickname)">
                <text class="avatar-text">{{ getAvatarText(post.userNickname) }}</text>
              </view>
              <view class="user-details">
                <text class="username">{{ post.userNickname }}</text>
                <text class="post-time">{{ formatTime(post.createdAt) }}</text>
              </view>
            </view>
            <view class="category-tag" :class="getCategoryClass(post.category)">
              <text>{{ post.categoryName }}</text>
            </view>
          </view>

          <!-- 帖子内容 -->
          <view class="post-content">
            <text class="post-title">{{ post.title }}</text>
            <text class="post-text">{{ post.content }}</text>
          </view>

          <!-- 互动区域 -->
          <view class="post-actions">
            <view class="action-group">
              <view class="action-item" @tap.stop="toggleLike(post)">
                <text class="icon" :class="{ liked: post.isLiked }">❤️</text>
                <text class="count">{{ post.likeCount }}</text>
              </view>
              <view class="action-item" @tap.stop="viewPostDetail(post.id)">
                <text class="icon">💬</text>
                <text class="count">{{ post.commentCount }}</text>
              </view>
              <view class="action-item">
                <text class="icon">👁️</text>
                <text class="count">{{ post.viewCount }}</text>
              </view>
            </view>
            <!-- 鼓励按钮 (仅求助类帖子显示) -->
            <view class="encourage-btn" v-if="post.category === 2" @tap.stop="encourageUser(post)">
              <text>鼓励TA</text>
            </view>
          </view>
        </view>

        <!-- 加载更多 -->
        <view class="load-more" v-if="hasMore">
          <text>{{ loading ? '加载中...' : '上拉加载更多' }}</text>
        </view>

        <!-- 没有更多数据 -->
        <view class="no-more" v-if="!hasMore && posts.length > 0">
          <text>没有更多内容了</text>
        </view>
      </view>
    </scroll-view>

    <!-- 底部导航 -->
    <view class="bottom-nav">
      <nf-tab-bar current="community" />
    </view>
  </view>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import NfTabBar from '@/components/ui/navigation/NfTabBar.vue'

export default {
  name: 'CommunityIndex',
  components: {
    NfTabBar
  },
  setup() {
    // 响应式数据
    const currentTime = ref('9:41')
    const selectedCategory = ref(0)
    const posts = ref([])
    const loading = ref(false)
    const hasMore = ref(true)
    const page = ref(1)
    
    // 方法
    const updateTime = () => {
      const now = new Date()
      currentTime.value = `${now.getHours()}:${now.getMinutes().toString().padStart(2, '0')}`
    }

    const selectCategory = (category) => {
      selectedCategory.value = category
      page.value = 1
      posts.value = []
      loadPosts()
    }

    const loadPosts = async () => {
      if (loading.value) return
      
      loading.value = true
      try {
        // 模拟数据，后续连接真实API
        const mockPosts = [
          {
            id: 1,
            title: '🎉 100天里程碑达成！',
            content: '想跟大家分享一些这段路程的心得...',
            category: 4,
            categoryName: '成功故事',
            userNickname: '坚持者_阳光',
            isAnonymous: false,
            viewCount: 156,
            likeCount: 128,
            commentCount: 45,
            isLiked: false,
            createdAt: new Date(Date.now() - 180000).toISOString()
          },
          {
            id: 2,
            title: '感觉很困难，需要大家的鼓励 😔',
            content: '第7天了，总是想要放弃...',
            category: 2,
            categoryName: '求助求鼓励',
            userNickname: '新手_求助',
            isAnonymous: false,
            viewCount: 89,
            likeCount: 32,
            commentCount: 18,
            isLiked: false,
            createdAt: new Date(Date.now() - 900000).toISOString()
          },
          {
            id: 3,
            title: '第30天打卡 ✅',
            content: '今天是第30天，感觉状态不错！',
            category: 3,
            categoryName: '日常打卡',
            userNickname: '努力的小伙',
            isAnonymous: false,
            viewCount: 45,
            likeCount: 28,
            commentCount: 12,
            isLiked: true,
            createdAt: new Date(Date.now() - 1800000).toISOString()
          },
          {
            id: 4,
            title: '分享一些有效的转移注意力方法',
            content: '运动、阅读、冥想都是很好的方式...',
            category: 1,
            categoryName: '经验分享',
            userNickname: '经验分享者',
            isAnonymous: false,
            viewCount: 234,
            likeCount: 89,
            commentCount: 36,
            isLiked: false,
            createdAt: new Date(Date.now() - 3600000).toISOString()
          }
        ]

        if (page.value === 1) {
          posts.value = mockPosts
        }
        hasMore.value = false
      } catch (error) {
        console.error('加载帖子失败:', error)
        uni.showToast({
          title: '网络错误',
          icon: 'none'
        })
      } finally {
        loading.value = false
      }
    }

    const loadMorePosts = () => {
      if (hasMore.value && !loading.value) {
        page.value++
        loadPosts()
      }
    }

    const toggleLike = async (post) => {
      post.isLiked = !post.isLiked
      post.likeCount += post.isLiked ? 1 : -1
      
      uni.showToast({
        title: post.isLiked ? '点赞成功' : '取消点赞',
        icon: 'success'
      })
    }

    const viewPostDetail = (postId) => {
      uni.navigateTo({
        url: `/pages/community/detail?id=${postId}`
      })
    }

    const openPostModal = () => {
      uni.navigateTo({
        url: '/pages/community/post'
      })
    }

    const openSearch = () => {
      uni.showToast({
        title: '搜索功能开发中',
        icon: 'none'
      })
    }

    const encourageUser = (post) => {
      uni.showToast({
        title: '已给TA发送鼓励',
        icon: 'success'
      })
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
      
      if (diff < 60000) { // 1分钟内
        return '刚刚'
      } else if (diff < 3600000) { // 1小时内
        return `${Math.floor(diff / 60000)}分钟前`
      } else if (diff < 86400000) { // 1天内
        return `${Math.floor(diff / 3600000)}小时前`
      } else {
        return date.toLocaleDateString()
      }
    }

    // 生命周期
    onMounted(() => {
      updateTime()
      setInterval(updateTime, 60000) // 每分钟更新时间
      loadPosts()
    })

    return {
      currentTime,
      selectedCategory,
      posts,
      loading,
      hasMore,
      selectCategory,
      loadMorePosts,
      toggleLike,
      viewPostDetail,
      openPostModal,
      openSearch,
      encourageUser,
      getAvatarClass,
      getAvatarText,
      getCategoryClass,
      formatTime
    }
  }
}
</script>

<style scoped>
.community-page {
  min-height: 100vh;
  background: #F8FAFC;
  padding-bottom: 80px;
}

/* 状态栏 */
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

/* 头部 */
.header {
  background: #FFFFFF;
  padding: 16px 24px;
  border-bottom: 1px solid #E5E7EB;
}

.header-main {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.title {
  font-size: 24px;
  font-weight: bold;
  color: #1F2937;
}

.header-actions {
  display: flex;
  gap: 12px;
}

.action-btn {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

.search-btn {
  background: #F3F4F6;
  color: #6B7280;
}

.post-btn {
  background: rgba(34, 211, 153, 0.1);
  color: #22D3AA;
}

/* 分类标签 */
.category-tabs {
  white-space: nowrap;
}

.tab-item {
  display: inline-block;
  padding: 8px 16px;
  margin-right: 8px;
  border-radius: 20px;
  font-size: 14px;
  font-weight: 500;
  background: #F3F4F6;
  color: #6B7280;
  transition: all 0.3s;
}

.tab-item.active {
  background: #22D3AA;
  color: #FFFFFF;
}

/* 帖子列表 */
.posts-container {
  flex: 1;
  height: calc(100vh - 200px);
}

.posts-list {
  padding: 16px;
}

.post-item {
  background: #FFFFFF;
  border-radius: 16px;
  padding: 16px;
  margin-bottom: 16px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

/* 帖子头部 */
.post-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 14px;
  color: #FFFFFF;
}

.avatar.anonymous {
  background: linear-gradient(135deg, #6B7280, #9CA3AF);
}

.avatar.avatar-0 {
  background: linear-gradient(135deg, #F59E0B, #F97316);
}

.avatar.avatar-1 {
  background: linear-gradient(135deg, #3B82F6, #8B5CF6);
}

.avatar.avatar-2 {
  background: linear-gradient(135deg, #10B981, #22D3AA);
}

.avatar.avatar-3 {
  background: linear-gradient(135deg, #EF4444, #F97316);
}

.avatar.avatar-4 {
  background: linear-gradient(135deg, #8B5CF6, #EC4899);
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.username {
  font-weight: 600;
  color: #1F2937;
  font-size: 14px;
}

.post-time {
  font-size: 12px;
  color: #6B7280;
}

.category-tag {
  padding: 4px 8px;
  border-radius: 8px;
  font-size: 12px;
  font-weight: 500;
}

.category-tag.experience {
  background: rgba(59, 130, 246, 0.1);
  color: #3B82F6;
}

.category-tag.help {
  background: rgba(239, 68, 68, 0.1);
  color: #EF4444;
}

.category-tag.checkin {
  background: rgba(16, 185, 129, 0.1);
  color: #10B981;
}

.category-tag.success {
  background: rgba(245, 158, 11, 0.1);
  color: #F59E0B;
}

/* 帖子内容 */
.post-content {
  margin-bottom: 12px;
}

.post-title {
  font-weight: 600;
  color: #1F2937;
  font-size: 16px;
  margin-bottom: 8px;
  display: block;
}

.post-text {
  color: #4B5563;
  font-size: 14px;
  line-height: 1.5;
  display: block;
}

/* 互动区域 */
.post-actions {
  padding-top: 12px;
  border-top: 1px solid #F3F4F6;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.action-group {
  display: flex;
  gap: 24px;
}

.action-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #6B7280;
}

.action-item .icon {
  font-size: 16px;
}

.action-item .icon.liked {
  color: #EF4444;
}

.action-item .count {
  font-size: 14px;
}

/* 鼓励按钮 */
.encourage-btn {
  padding: 6px 12px;
  background: rgba(34, 211, 170, 0.1);
  color: #22D3AA;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

/* 加载更多 */
.load-more, .no-more {
  text-align: center;
  padding: 20px;
  color: #6B7280;
  font-size: 14px;
}

/* 底部导航 */
.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 100;
}
</style> 