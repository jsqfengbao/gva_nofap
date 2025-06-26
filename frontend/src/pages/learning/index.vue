<template>
  <view class="learning-page">
    <!-- 状态栏 -->
    <view class="status-bar">
      <text class="time">{{ currentTime }}</text>
      <view class="status-icons">
        <text class="signal">📶</text>
        <text class="wifi">📶</text>
        <view class="battery">
          <view class="battery-level"></view>
        </view>
      </view>
    </view>

    <!-- 头部导航 -->
    <view class="header">
      <view class="header-top">
        <text class="title">学习内容</text>
        <view class="search-btn" @click="showSearch">
          <text class="search-icon">🔍</text>
        </view>
      </view>

      <!-- 分类标签 -->
      <scroll-view class="category-tabs" scroll-x="true" show-scrollbar="false">
        <view class="tabs-container">
          <view 
            v-for="(tab, index) in categoryTabs" 
            :key="index"
            class="tab-item"
            :class="{ active: activeTabIndex === index }"
            @click="switchTab(index)"
          >
            <text class="tab-text">{{ tab.name }}</text>
          </view>
        </view>
      </scroll-view>
    </view>

    <!-- 主要内容 -->
    <scroll-view class="main-content" scroll-y="true" @scrolltolower="loadMore">
      <!-- 精品课程推荐 -->
      <view class="featured-course" v-if="activeTabIndex === 0">
        <view class="course-card">
          <image 
            class="course-image" 
            :src="featuredCourse.thumbnail"
            mode="aspectFill"
          />
          <view class="course-info">
            <text class="course-badge">精品课程</text>
            <text class="course-title">{{ featuredCourse.title }}</text>
            <text class="course-desc">{{ featuredCourse.summary }}</text>
            <view class="course-btn" @click="startLearning(featuredCourse)">
              <text class="btn-text">开始学习</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 学习进度统计 -->
      <view class="learning-progress" v-if="activeTabIndex === 0">
        <text class="section-title">学习进度</text>
        
        <view class="progress-content">
          <view class="weekly-progress">
            <view class="progress-header">
              <text class="progress-label">本周学习时长</text>
              <text class="progress-value">{{ learningStats.weeklyHours }}小时</text>
            </view>
            <view class="progress-bar">
              <view 
                class="progress-fill" 
                :style="{ width: learningStats.weeklyProgress + '%' }"
              ></view>
            </view>
            <text class="progress-goal">目标: {{ learningStats.weeklyGoal }}小时/周</text>
          </view>
          
          <view class="stats-grid">
            <view class="stat-item">
              <text class="stat-number">{{ learningStats.articlesRead }}</text>
              <text class="stat-label">文章已读</text>
            </view>
            <view class="stat-item">
              <text class="stat-number">{{ learningStats.videosWatched }}</text>
              <text class="stat-label">视频已看</text>
            </view>
            <view class="stat-item">
              <text class="stat-number">{{ learningStats.audiosListened }}</text>
              <text class="stat-label">音频已听</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 内容列表 -->
      <view class="content-sections">
        <!-- 文章区域 -->
        <view class="content-section" v-if="shouldShowSection('articles')">
          <view class="section-header">
            <text class="section-title">{{ getSectionTitle('articles') }}</text>
            <text class="more-btn" @click="viewMore('articles')">查看更多</text>
          </view>
          
          <view class="article-list">
            <view 
              v-for="article in currentArticles" 
              :key="article.id"
              class="article-item"
              @click="viewContent(article)"
            >
              <image 
                class="article-image" 
                :src="article.thumbnailUrl"
                mode="aspectFill"
              />
              <view class="article-info">
                <text class="article-title">{{ article.title }}</text>
                <text class="article-desc">{{ article.summary }}</text>
                <view class="article-meta">
                  <text class="meta-item">⏱️ {{ article.duration }}分钟</text>
                  <text class="meta-item">❤️ {{ article.likeCount }}</text>
                  <text class="meta-item">👁️ {{ formatNumber(article.viewCount) }}</text>
                </view>
              </view>
            </view>
          </view>
        </view>

        <!-- 视频区域 -->
        <view class="content-section" v-if="shouldShowSection('videos')">
          <view class="section-header">
            <text class="section-title">{{ getSectionTitle('videos') }}</text>
            <text class="more-btn" @click="viewMore('videos')">查看更多</text>
          </view>
          
          <view class="video-list">
            <view 
              v-for="video in currentVideos" 
              :key="video.id"
              class="video-item"
              @click="viewContent(video)"
            >
              <view class="video-thumbnail">
                <image 
                  class="video-image" 
                  :src="video.thumbnailUrl"
                  mode="aspectFill"
                />
                <view class="play-overlay">
                  <view class="play-button">
                    <text class="play-icon">▶️</text>
                  </view>
                </view>
                <text class="video-duration">{{ formatDuration(video.duration) }}</text>
              </view>
              <text class="video-title">{{ video.title }}</text>
              <text class="video-desc">{{ video.summary }}</text>
              <view class="video-meta">
                <text class="meta-item">▶️ {{ formatNumber(video.viewCount) }}</text>
                <text class="meta-item">👍 {{ video.likeCount }}</text>
                <text class="meta-item">💬 {{ video.commentCount }}</text>
              </view>
            </view>
          </view>
        </view>

        <!-- 音频区域 -->
        <view class="content-section" v-if="shouldShowSection('audios')">
          <view class="section-header">
            <text class="section-title">{{ getSectionTitle('audios') }}</text>
            <text class="more-btn" @click="viewMore('audios')">查看更多</text>
          </view>
          
          <view class="audio-list">
            <view 
              v-for="audio in currentAudios" 
              :key="audio.id"
              class="audio-item"
              @click="viewContent(audio)"
            >
              <view class="audio-icon">
                <text class="headphones">🎧</text>
              </view>
              <view class="audio-info">
                <text class="audio-title">{{ audio.title }}</text>
                <text class="audio-desc">{{ audio.summary }}</text>
                <view class="audio-meta">
                  <text class="meta-item">⏱️ {{ audio.duration }}分钟</text>
                  <text class="meta-item" v-if="audio.isDownloaded">📥 已下载</text>
                </view>
              </view>
              <view class="audio-play-btn" @click.stop="playAudio(audio)">
                <text class="play-icon">▶️</text>
              </view>
            </view>
          </view>
        </view>
      </view>

      <!-- 加载更多 -->
      <view class="load-more" v-if="hasMore">
        <text class="load-text">{{ loadingMore ? '加载中...' : '上拉加载更多' }}</text>
      </view>
    </scroll-view>

    <!-- 底部导航 -->
    <nf-tab-bar current="learning" />

    <!-- 搜索弹窗 -->
    <view class="search-modal" v-if="showSearchModal" @click="hideSearch">
      <view class="search-content" @click.stop>
        <view class="search-header">
          <input 
            class="search-input"
            v-model="searchKeyword"
            placeholder="搜索学习内容..."
            @input="onSearchInput"
          />
          <text class="search-cancel" @click="hideSearch">取消</text>
        </view>
        <view class="search-results" v-if="searchResults.length > 0">
          <view 
            v-for="result in searchResults" 
            :key="result.id"
            class="search-item"
            @click="viewContent(result)"
          >
            <text class="search-title">{{ result.title }}</text>
            <text class="search-desc">{{ result.summary }}</text>
          </view>
        </view>
        <view class="no-results" v-else-if="searchKeyword">
          <text class="no-results-text">暂无搜索结果</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { getCurrentTime, formatNumber } from '@/utils/format'
import NfTabBar from '@/components/ui/navigation/NfTabBar.vue'

export default {
  name: 'LearningIndex',
  components: {
    NfTabBar
  },
  data() {
    return {
      currentTime: getCurrentTime(),
      activeTabIndex: 0,
      categoryTabs: [
        { name: '推荐', type: 'recommend' },
        { name: '文章', type: 'article' },
        { name: '视频', type: 'video' },
        { name: '音频', type: 'audio' }
      ],
      
      // 精品课程
      featuredCourse: {
        id: 1,
        title: '21天自控力训练营',
        summary: '科学方法帮你建立持久的自控习惯',
        thumbnail: 'https://images.unsplash.com/photo-1499336315816-097655dcfbda?w=160&h=160&fit=crop',
        category: 2,
        difficulty: 2
      },
      
      // 学习统计
      learningStats: {
        weeklyHours: 3.5,
        weeklyGoal: 5,
        weeklyProgress: 70,
        articlesRead: 12,
        videosWatched: 5,
        audiosListened: 8
      },
      
      // 内容数据
      articles: [],
      videos: [],
      audios: [],
      
      // 分页
      currentPage: 1,
      pageSize: 10,
      hasMore: true,
      loadingMore: false,
      
      // 搜索
      showSearchModal: false,
      searchKeyword: '',
      searchResults: []
    }
  },
  
  computed: {
    currentArticles() {
      return this.activeTabIndex === 0 || this.activeTabIndex === 1 
        ? this.articles.slice(0, this.activeTabIndex === 0 ? 2 : this.articles.length)
        : []
    },
    
    currentVideos() {
      return this.activeTabIndex === 0 || this.activeTabIndex === 2
        ? this.videos.slice(0, this.activeTabIndex === 0 ? 1 : this.videos.length)
        : []
    },
    
    currentAudios() {
      return this.activeTabIndex === 0 || this.activeTabIndex === 3
        ? this.audios.slice(0, this.activeTabIndex === 0 ? 1 : this.audios.length)
        : []
    }
  },
  
  onLoad() {
    this.initPage()
  },
  
  onShow() {
    this.updateTime()
  },
  
  methods: {
    // 初始化页面
    async initPage() {
      // 初始化页面数据
      this.updateTime()
      this.loadMockData()
      
      // 尝试加载真实数据（如果API可用）
      try {
        await this.loadLearningStats()
      } catch (error) {
        console.log('使用模拟数据')
      }
    },
    
    // 加载模拟数据
    loadMockData() {
      // 文章数据
      this.articles = [
        {
          id: 1,
          title: '如何建立健康的生活习惯',
          summary: '从小习惯开始，逐步建立健康的生活方式...',
          thumbnailUrl: 'https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=160&h=120&fit=crop',
          duration: 8,
          likeCount: 234,
          viewCount: 1200,
          contentType: 1,
          category: 2,
          createdAt: '2024-06-20T10:00:00Z'
        },
        {
          id: 2,
          title: '理解大脑的奖励机制',
          summary: '科学解释成瘾的神经学原理和应对方法...',
          thumbnailUrl: 'https://images.unsplash.com/photo-1559757148-5c350d0d3c56?w=160&h=120&fit=crop',
          duration: 12,
          likeCount: 189,
          viewCount: 856,
          contentType: 1,
          category: 1,
          createdAt: '2024-06-19T14:30:00Z'
        },
        {
          id: 3,
          title: '正念冥想入门指南',
          summary: '学会用冥想管理情绪和冲动，建立内心平静...',
          thumbnailUrl: 'https://images.unsplash.com/photo-1545389336-cf090694435e?w=160&h=120&fit=crop',
          duration: 15,
          likeCount: 156,
          viewCount: 743,
          contentType: 1,
          category: 3,
          createdAt: '2024-06-18T09:15:00Z'
        }
      ]
      
      // 视频数据
      this.videos = [
        {
          id: 4,
          title: '正念冥想入门指南',
          summary: '学会用冥想管理情绪和冲动',
          thumbnailUrl: 'https://images.unsplash.com/photo-1545389336-cf090694435e?w=300&h=200&fit=crop',
          duration: 15,
          likeCount: 127,
          viewCount: 3200,
          commentCount: 23,
          contentType: 2,
          category: 3,
          createdAt: '2024-06-17T16:45:00Z',
          videoUrl: 'https://example.com/video/meditation.mp4'
        },
        {
          id: 5,
          title: '呼吸练习技巧讲解',
          summary: '通过正确的呼吸方法缓解压力和焦虑',
          thumbnailUrl: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300&h=200&fit=crop',
          duration: 10,
          likeCount: 89,
          viewCount: 1800,
          commentCount: 15,
          contentType: 2,
          category: 2,
          createdAt: '2024-06-16T11:20:00Z',
          videoUrl: 'https://example.com/video/breathing.mp4'
        }
      ]
      
      // 音频数据
      this.audios = [
        {
          id: 6,
          title: '睡前放松冥想',
          summary: '帮助改善睡眠质量的引导冥想',
          thumbnailUrl: 'https://images.unsplash.com/photo-1542662565-7e4b16f20bfb?w=160&h=120&fit=crop',
          duration: 20,
          likeCount: 67,
          viewCount: 890,
          contentType: 3,
          category: 3,
          createdAt: '2024-06-15T20:00:00Z',
          audioUrl: 'https://example.com/audio/sleep-meditation.mp3',
          isDownloaded: true
        },
        {
          id: 7,
          title: '专注力训练音频',
          summary: '提升注意力和专注能力的训练课程',
          thumbnailUrl: 'https://images.unsplash.com/photo-1499209974431-9dddcece7f88?w=160&h=120&fit=crop',
          duration: 25,
          likeCount: 45,
          viewCount: 612,
          contentType: 3,
          category: 2,
          createdAt: '2024-06-14T08:30:00Z',
          audioUrl: 'https://example.com/audio/focus-training.mp3',
          isDownloaded: false
        }
      ]
    },
    
    // 加载学习统计
    async loadLearningStats() {
      try {
        const res = await uni.request({
          url: '/api/v1/miniprogram/learning/stats',
          method: 'GET',
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.learningStats = {
            weeklyHours: (res.data.data.totalLearningTime / 60).toFixed(1),
            weeklyGoal: 5,
            weeklyProgress: Math.min((res.data.data.totalLearningTime / 300) * 100, 100),
            articlesRead: res.data.data.completedContents || 0,
            videosWatched: res.data.data.likedContents || 0,
            audiosListened: res.data.data.collectedContents || 0
          }
        }
      } catch (error) {
        console.error('加载学习统计失败:', error)
      }
    },
    
    // 加载内容列表
    async loadContent(loadMore = false) {
      if (this.loadingMore) return
      
      this.loadingMore = true
      
      try {
        const contentType = this.getContentType()
        const res = await uni.request({
          url: '/api/v1/miniprogram/learning/contents',
          method: 'GET',
          data: {
            page: loadMore ? this.currentPage + 1 : 1,
            pageSize: this.pageSize,
            contentType: contentType || undefined,
            sortBy: 'view_count',
            order: 'desc'
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          const newContent = res.data.data.list || []
          
          if (loadMore) {
            this.currentPage++
            this.appendContent(newContent)
          } else {
            this.currentPage = 1
            this.setContent(newContent)
          }
          
          this.hasMore = newContent.length === this.pageSize
        }
      } catch (error) {
        console.error('加载内容失败:', error)
        uni.showToast({
          title: '加载失败',
          icon: 'error'
        })
      } finally {
        this.loadingMore = false
      }
    },
    
    // 获取当前内容类型
    getContentType() {
      const typeMap = {
        0: null, // 推荐
        1: 1,    // 文章
        2: 2,    // 视频
        3: 3     // 音频
      }
      return typeMap[this.activeTabIndex]
    },
    
    // 设置内容数据
    setContent(content) {
      content.forEach(item => {
        if (item.contentType === 1) {
          if (!this.articles.find(a => a.id === item.id)) {
            this.articles.push(item)
          }
        } else if (item.contentType === 2) {
          if (!this.videos.find(v => v.id === item.id)) {
            this.videos.push(item)
          }
        } else if (item.contentType === 3) {
          if (!this.audios.find(a => a.id === item.id)) {
            this.audios.push(item)
          }
        }
      })
    },
    
    // 追加内容数据
    appendContent(content) {
      this.setContent(content)
    },
    
    // 切换标签
    async switchTab(index) {
      if (this.activeTabIndex === index) return
      
      this.activeTabIndex = index
      this.currentPage = 1
      this.hasMore = true
      
      // 重新加载内容
      await this.loadContent()
    },
    
    // 判断是否显示区域
    shouldShowSection(type) {
      if (this.activeTabIndex === 0) return true // 推荐页显示所有
      
      const typeMap = {
        'articles': 1,
        'videos': 2, 
        'audios': 3
      }
      
      return this.activeTabIndex === typeMap[type]
    },
    
    // 获取区域标题
    getSectionTitle(type) {
      const titleMap = {
        'articles': this.activeTabIndex === 0 ? '热门文章' : '文章列表',
        'videos': this.activeTabIndex === 0 ? '推荐视频' : '视频列表',
        'audios': this.activeTabIndex === 0 ? '音频内容' : '音频列表'
      }
      
      return titleMap[type]
    },
    
    // 查看更多
    viewMore(type) {
      const typeMap = {
        'articles': 1,
        'videos': 2,
        'audios': 3
      }
      
      this.switchTab(typeMap[type])
    },
    
    // 查看内容详情
    viewContent(content) {
      uni.navigateTo({
        url: `/pages/learning/detail?id=${content.id}`
      })
    },
    
    // 开始学习
    startLearning(content) {
      this.viewContent(content)
    },
    
    // 播放音频
    async playAudio(audio) {
      try {
        // 开始学习记录
        await uni.request({
          url: '/api/v1/miniprogram/learning/start',
          method: 'POST',
          data: {
            contentId: audio.id
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        // 跳转到音频播放页面
        uni.navigateTo({
          url: `/pages/learning/audio?id=${audio.id}`
        })
      } catch (error) {
        console.error('播放音频失败:', error)
        uni.showToast({
          title: '播放失败',
          icon: 'error'
        })
      }
    },
    
    // 显示搜索
    showSearch() {
      this.showSearchModal = true
    },
    
    // 隐藏搜索
    hideSearch() {
      this.showSearchModal = false
      this.searchKeyword = ''
      this.searchResults = []
    },
    
    // 搜索输入
    async onSearchInput() {
      if (!this.searchKeyword.trim()) {
        this.searchResults = []
        return
      }
      
      try {
        const res = await uni.request({
          url: '/api/v1/miniprogram/learning/contents',
          method: 'GET',
          data: {
            page: 1,
            pageSize: 10,
            keyword: this.searchKeyword
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.searchResults = res.data.data.list || []
        }
      } catch (error) {
        console.error('搜索失败:', error)
      }
    },
    
    // 加载更多
    loadMore() {
      if (this.hasMore && !this.loadingMore) {
        this.loadContent(true)
      }
    },
    
    // 更新时间
    updateTime() {
      this.currentTime = getCurrentTime()
    },
    
    // 格式化数字
    formatNumber,
    
    // 格式化时长
    formatDuration(minutes) {
      const hours = Math.floor(minutes / 60)
      const mins = minutes % 60
      
      if (hours > 0) {
        return `${hours}:${mins.toString().padStart(2, '0')}`
      }
      return `${mins}:00`
    }
  }
}
</script>

<style lang="scss">
.learning-page {
  min-height: 100vh;
  background: #F8FAFC;
  padding-bottom: 140rpx;
}

/* 状态栏 */
.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20rpx 32rpx;
  font-size: 24rpx;
  color: #1F2937;
  
  .time {
    font-weight: 500;
  }
  
  .status-icons {
    display: flex;
    align-items: center;
    gap: 8rpx;
    
    .signal, .wifi {
      font-size: 20rpx;
    }
    
    .battery {
      width: 48rpx;
      height: 24rpx;
      border: 2rpx solid #1F2937;
      border-radius: 4rpx;
      position: relative;
      
      .battery-level {
        width: 32rpx;
        height: 12rpx;
        background: #10B981;
        border-radius: 2rpx;
        margin: 4rpx;
      }
    }
  }
}

/* 头部 */
.header {
  background: white;
  padding: 32rpx 48rpx;
  border-bottom: 2rpx solid #F3F4F6;
  
  .header-top {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 32rpx;
    
    .title {
      font-size: 48rpx;
      font-weight: bold;
      color: #1F2937;
    }
    
    .search-btn {
      width: 80rpx;
      height: 80rpx;
      background: #F3F4F6;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      
      .search-icon {
        font-size: 32rpx;
        color: #6B7280;
      }
    }
  }
  
  .category-tabs {
    white-space: nowrap;
    
    .tabs-container {
      display: flex;
      gap: 16rpx;
      
      .tab-item {
        flex-shrink: 0;
        padding: 16rpx 32rpx;
        background: #F3F4F6;
        border-radius: 50rpx;
        
        .tab-text {
          font-size: 28rpx;
          font-weight: 500;
          color: #6B7280;
        }
        
        &.active {
          background: #34D399;
          
          .tab-text {
            color: white;
          }
        }
      }
    }
  }
}

/* 主要内容 */
.main-content {
  padding: 48rpx;
  height: calc(100vh - 300rpx);
}

/* 精品课程 */
.featured-course {
  margin-bottom: 48rpx;
  
  .course-card {
    background: linear-gradient(135deg, #F3E8FF 0%, #E0E7FF 100%);
    border-radius: 32rpx;
    padding: 48rpx;
    border: 2rpx solid #C4B5FD;
    display: flex;
    gap: 32rpx;
    
    .course-image {
      width: 160rpx;
      height: 160rpx;
      border-radius: 24rpx;
      flex-shrink: 0;
    }
    
    .course-info {
      flex: 1;
      
      .course-badge {
        display: inline-block;
        padding: 8rpx 16rpx;
        background: #DDD6FE;
        color: #7C3AED;
        font-size: 24rpx;
        border-radius: 20rpx;
        margin-bottom: 16rpx;
      }
      
      .course-title {
        display: block;
        font-size: 32rpx;
        font-weight: bold;
        color: #1F2937;
        margin-bottom: 8rpx;
      }
      
      .course-desc {
        display: block;
        font-size: 28rpx;
        color: #6B7280;
        margin-bottom: 24rpx;
        line-height: 1.5;
      }
      
      .course-btn {
        background: #7C3AED;
        color: white;
        padding: 16rpx 32rpx;
        border-radius: 24rpx;
        display: inline-block;
        
        .btn-text {
          font-size: 28rpx;
          font-weight: 500;
        }
      }
    }
  }
}

/* 学习进度 */
.learning-progress {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 48rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
  
  .section-title {
    font-size: 36rpx;
    font-weight: 600;
    color: #1F2937;
    margin-bottom: 32rpx;
    display: block;
  }
  
  .progress-content {
    .weekly-progress {
      margin-bottom: 32rpx;
      
      .progress-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 16rpx;
        
        .progress-label {
          font-size: 28rpx;
          font-weight: 500;
          color: #1F2937;
        }
        
        .progress-value {
          font-size: 28rpx;
          font-weight: 600;
          color: #34D399;
        }
      }
      
      .progress-bar {
        width: 100%;
        height: 16rpx;
        background: #E5E7EB;
        border-radius: 8rpx;
        margin-bottom: 8rpx;
        
        .progress-fill {
          height: 100%;
          background: #34D399;
          border-radius: 8rpx;
          transition: width 0.3s ease;
        }
      }
      
      .progress-goal {
        font-size: 24rpx;
        color: #6B7280;
      }
    }
    
    .stats-grid {
      display: flex;
      justify-content: space-around;
      padding-top: 32rpx;
      border-top: 2rpx solid #F3F4F6;
      
      .stat-item {
        text-align: center;
        
        .stat-number {
          display: block;
          font-size: 40rpx;
          font-weight: bold;
          color: #10B981;
          margin-bottom: 8rpx;
        }
        
        .stat-label {
          font-size: 24rpx;
          color: #6B7280;
        }
      }
    }
  }
}

/* 内容区域 */
.content-sections {
  .content-section {
    margin-bottom: 48rpx;
    
    .section-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 32rpx;
      
      .section-title {
        font-size: 36rpx;
        font-weight: 600;
        color: #1F2937;
      }
      
      .more-btn {
        font-size: 28rpx;
        font-weight: 500;
        color: #34D399;
      }
    }
  }
}

/* 文章列表 */
.article-list {
  .article-item {
    background: white;
    border-radius: 32rpx;
    padding: 32rpx;
    margin-bottom: 32rpx;
    box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
    display: flex;
    gap: 32rpx;
    
    .article-image {
      width: 160rpx;
      height: 120rpx;
      border-radius: 24rpx;
      flex-shrink: 0;
    }
    
    .article-info {
      flex: 1;
      
      .article-title {
        font-size: 32rpx;
        font-weight: 600;
        color: #1F2937;
        margin-bottom: 8rpx;
        display: block;
        line-height: 1.4;
      }
      
      .article-desc {
        font-size: 28rpx;
        color: #6B7280;
        margin-bottom: 16rpx;
        display: block;
        line-height: 1.5;
      }
      
      .article-meta {
        display: flex;
        gap: 32rpx;
        
        .meta-item {
          font-size: 24rpx;
          color: #9CA3AF;
        }
      }
    }
  }
}

/* 视频列表 */
.video-list {
  .video-item {
    background: white;
    border-radius: 32rpx;
    padding: 32rpx;
    margin-bottom: 32rpx;
    box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
    
    .video-thumbnail {
      position: relative;
      margin-bottom: 24rpx;
      
      .video-image {
        width: 100%;
        height: 320rpx;
        border-radius: 24rpx;
      }
      
      .play-overlay {
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: rgba(0, 0, 0, 0.2);
        border-radius: 24rpx;
        display: flex;
        align-items: center;
        justify-content: center;
        
        .play-button {
          width: 96rpx;
          height: 96rpx;
          background: rgba(255, 255, 255, 0.9);
          border-radius: 50%;
          display: flex;
          align-items: center;
          justify-content: center;
          
          .play-icon {
            font-size: 32rpx;
            color: #1F2937;
            margin-left: 8rpx;
          }
        }
      }
      
      .video-duration {
        position: absolute;
        bottom: 16rpx;
        right: 16rpx;
        background: rgba(0, 0, 0, 0.7);
        color: white;
        padding: 8rpx 16rpx;
        border-radius: 8rpx;
        font-size: 24rpx;
      }
    }
    
    .video-title {
      font-size: 32rpx;
      font-weight: 600;
      color: #1F2937;
      margin-bottom: 8rpx;
      display: block;
      line-height: 1.4;
    }
    
    .video-desc {
      font-size: 28rpx;
      color: #6B7280;
      margin-bottom: 16rpx;
      display: block;
      line-height: 1.5;
    }
    
    .video-meta {
      display: flex;
      gap: 32rpx;
      
      .meta-item {
        font-size: 24rpx;
        color: #9CA3AF;
      }
    }
  }
}

/* 音频列表 */
.audio-list {
  .audio-item {
    background: white;
    border-radius: 32rpx;
    padding: 32rpx;
    margin-bottom: 32rpx;
    box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
    display: flex;
    align-items: center;
    gap: 32rpx;
    
    .audio-icon {
      width: 128rpx;
      height: 128rpx;
      background: linear-gradient(135deg, #34D399 0%, #06B6D4 100%);
      border-radius: 24rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-shrink: 0;
      
      .headphones {
        font-size: 40rpx;
      }
    }
    
    .audio-info {
      flex: 1;
      
      .audio-title {
        font-size: 32rpx;
        font-weight: 600;
        color: #1F2937;
        margin-bottom: 8rpx;
        display: block;
        line-height: 1.4;
      }
      
      .audio-desc {
        font-size: 28rpx;
        color: #6B7280;
        margin-bottom: 16rpx;
        display: block;
        line-height: 1.5;
      }
      
      .audio-meta {
        display: flex;
        gap: 32rpx;
        
        .meta-item {
          font-size: 24rpx;
          color: #9CA3AF;
        }
      }
    }
    
    .audio-play-btn {
      width: 80rpx;
      height: 80rpx;
      background: rgba(52, 211, 153, 0.1);
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-shrink: 0;
      
      .play-icon {
        font-size: 28rpx;
        color: #34D399;
      }
    }
  }
}

/* 加载更多 */
.load-more {
  text-align: center;
  padding: 32rpx;
  
  .load-text {
    color: #9CA3AF;
    font-size: 28rpx;
  }
}

/* 搜索弹窗 */
.search-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 1000;
  
  .search-content {
    background: white;
    margin: 100rpx 32rpx;
    border-radius: 32rpx;
    padding: 32rpx;
    max-height: 80vh;
    overflow: hidden;
    
    .search-header {
      display: flex;
      align-items: center;
      gap: 16rpx;
      margin-bottom: 32rpx;
      
      .search-input {
        flex: 1;
        height: 80rpx;
        background: #F3F4F6;
        border-radius: 20rpx;
        padding: 0 24rpx;
        font-size: 28rpx;
        border: none;
        outline: none;
      }
      
      .search-cancel {
        color: #34D399;
        font-size: 28rpx;
        font-weight: 500;
      }
    }
    
    .search-results {
      max-height: 60vh;
      overflow-y: auto;
      
      .search-item {
        padding: 24rpx 0;
        border-bottom: 2rpx solid #F3F4F6;
        
        &:last-child {
          border-bottom: none;
        }
        
        .search-title {
          font-size: 30rpx;
          font-weight: 600;
          color: #1F2937;
          margin-bottom: 8rpx;
          display: block;
        }
        
        .search-desc {
          font-size: 26rpx;
          color: #6B7280;
          line-height: 1.5;
        }
      }
    }
    
    .no-results {
      text-align: center;
      padding: 80rpx 0;
      
      .no-results-text {
        color: #9CA3AF;
        font-size: 28rpx;
      }
    }
  }
}
</style> 