<template>
  <view class="learning-detail">
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
      <view class="back-btn" @click="goBack">
        <text class="back-icon">←</text>
      </view>
      <text class="header-title">学习详情</text>
      <view class="actions">
        <view class="action-btn" @click="toggleCollect">
          <text class="action-icon">{{ content.isCollected ? '★' : '☆' }}</text>
        </view>
        <view class="action-btn" @click="shareContent">
          <text class="action-icon">📤</text>
        </view>
      </view>
    </view>

    <!-- 主要内容 -->
    <scroll-view class="main-content" scroll-y="true">
      <!-- 内容头部信息 -->
      <view class="content-header">
        <view class="category-badge" :class="'category-' + content.category">
          <text class="badge-text">{{ getCategoryName(content.category) }}</text>
        </view>
        
        <text class="content-title">{{ content.title }}</text>
        
        <view class="content-meta">
          <text class="meta-item">⏱️ {{ content.duration }}分钟</text>
          <text class="meta-item">👁️ {{ formatNumber(content.viewCount) }}阅读</text>
          <text class="meta-item">❤️ {{ content.likeCount }}点赞</text>
          <text class="meta-item">📅 {{ formatDate(content.createdAt) }}</text>
        </view>
        
        <text class="content-summary">{{ content.summary }}</text>
      </view>

      <!-- 视频播放器 -->
      <view class="video-player" v-if="content.contentType === 2">
        <video 
          :src="content.videoUrl"
          :poster="content.thumbnailUrl"
          class="video-element"
          controls
          @play="onVideoPlay"
          @pause="onVideoPause"
          @timeupdate="onVideoTimeUpdate"
        />
      </view>

      <!-- 音频播放器 -->
      <view class="audio-player" v-if="content.contentType === 3">
        <view class="audio-controls">
          <view class="audio-cover">
            <image 
              class="cover-image" 
              :src="content.thumbnailUrl || '/static/images/audio-placeholder.png'"
              mode="aspectFill"
            />
            <view class="play-btn" @click="toggleAudioPlay">
              <text class="play-icon">{{ isAudioPlaying ? '⏸️' : '▶️' }}</text>
            </view>
          </view>
          
          <view class="audio-info">
            <text class="audio-title">{{ content.title }}</text>
            <view class="progress-container">
              <text class="time-current">{{ formatTime(audioCurrentTime) }}</text>
              <view class="progress-bar" @click="seekAudio">
                <view 
                  class="progress-fill" 
                  :style="{ width: audioProgress + '%' }"
                ></view>
              </view>
              <text class="time-total">{{ formatTime(audioDuration) }}</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 文章内容 -->
      <view class="article-content" v-if="content.contentType === 1">
        <rich-text class="content-body" :nodes="content.content"></rich-text>
      </view>

      <!-- 文本内容（音频/视频的补充文字） -->
      <view class="text-content" v-if="content.contentType !== 1 && content.textContent">
        <text class="section-title">内容简介</text>
        <rich-text class="content-text" :nodes="content.textContent"></rich-text>
      </view>

      <!-- 学习进度 -->
      <view class="learning-progress" v-if="learningRecord">
        <text class="section-title">学习进度</text>
        <view class="progress-info">
          <view class="progress-header">
            <text class="progress-label">完成进度</text>
            <text class="progress-value">{{ Math.round(learningRecord.progress) }}%</text>
          </view>
          <view class="progress-bar">
            <view 
              class="progress-fill" 
              :style="{ width: learningRecord.progress + '%' }"
            ></view>
          </view>
          <view class="progress-details">
            <text class="detail-item">开始时间: {{ formatDateTime(learningRecord.startTime) }}</text>
            <text class="detail-item" v-if="learningRecord.completedAt">
              完成时间: {{ formatDateTime(learningRecord.completedAt) }}
            </text>
            <text class="detail-item">学习时长: {{ Math.round(learningRecord.learningTime / 60) }}分钟</text>
          </view>
        </view>
      </view>

      <!-- 互动区域 -->
      <view class="interaction-section">
        <view class="action-buttons">
          <view class="action-item" @click="toggleLike">
            <text class="action-icon" :class="{ liked: content.isLiked }">
              {{ content.isLiked ? '❤️' : '🤍' }}
            </text>
            <text class="action-text">{{ content.isLiked ? '已点赞' : '点赞' }}</text>
          </view>
          
          <view class="action-item" @click="toggleCollect">
            <text class="action-icon" :class="{ collected: content.isCollected }">
              {{ content.isCollected ? '★' : '☆' }}
            </text>
            <text class="action-text">{{ content.isCollected ? '已收藏' : '收藏' }}</text>
          </view>
          
          <view class="action-item" @click="shareContent">
            <text class="action-icon">📤</text>
            <text class="action-text">分享</text>
          </view>
          
          <view class="action-item" @click="rateContent">
            <text class="action-icon">⭐</text>
            <text class="action-text">评分</text>
          </view>
        </view>
      </view>

      <!-- 相关推荐 -->
      <view class="related-content" v-if="relatedContents.length > 0">
        <text class="section-title">相关推荐</text>
        <view class="related-list">
          <view 
            v-for="item in relatedContents" 
            :key="item.id"
            class="related-item"
            @click="viewRelated(item)"
          >
            <image 
              class="related-image" 
              :src="item.thumbnailUrl"
              mode="aspectFill"
            />
            <view class="related-info">
              <text class="related-title">{{ item.title }}</text>
              <view class="related-meta">
                <text class="meta-text">{{ getCategoryName(item.category) }}</text>
                <text class="meta-text">{{ item.duration }}分钟</text>
              </view>
            </view>
          </view>
        </view>
      </view>
    </scroll-view>

    <!-- 底部操作栏 -->
    <view class="bottom-actions">
      <view class="progress-action" v-if="!learningRecord || !learningRecord.isCompleted">
        <nf-button 
          type="primary"
          class="complete-btn"
          @click="markAsCompleted"
        >
          标记为已完成
        </nf-button>
      </view>
      <view class="completed-action" v-else>
        <text class="completed-text">✅ 已完成学习</text>
        <nf-button 
          type="secondary"
          class="restart-btn"
          @click="restartLearning"
        >
          重新学习
        </nf-button>
      </view>
    </view>

    <!-- 评分弹窗 -->
    <view class="rating-modal" v-if="showRatingModal" @click="hideRating">
      <view class="rating-content" @click.stop>
        <text class="rating-title">为内容评分</text>
        <view class="stars">
          <text 
            v-for="star in 5" 
            :key="star"
            class="star"
            :class="{ active: star <= currentRating }"
            @click="setRating(star)"
          >
            ⭐
          </text>
        </view>
        <view class="rating-actions">
          <nf-button type="secondary" @click="hideRating">取消</nf-button>
          <nf-button type="primary" @click="submitRating">确定</nf-button>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'LearningDetail',
  data() {
    return {
      currentTime: '9:41',
      contentId: null,
      content: {},
      learningRecord: null,
      relatedContents: [],
      
      // 音频播放状态
      isAudioPlaying: false,
      audioCurrentTime: 0,
      audioDuration: 0,
      audioProgress: 0,
      audioContext: null,
      
      // 评分
      showRatingModal: false,
      currentRating: 0
    }
  },
  
  onLoad(options) {
    if (options.id) {
      this.contentId = parseInt(options.id)
      this.initPage()
    }
  },
  
  onShow() {
    this.updateTime()
  },
  
  onUnload() {
    // 页面卸载时停止音频播放
    if (this.audioContext) {
      this.audioContext.stop()
    }
  },
  
  methods: {
    async initPage() {
      uni.showLoading({ title: '加载中...' })
      
      try {
        await Promise.all([
          this.loadContent(),
          this.loadLearningRecord(),
          this.loadRelatedContent()
        ])
        
        // 记录开始学习
        await this.startLearningRecord()
      } catch (error) {
        console.error('页面初始化失败:', error)
        uni.showToast({
          title: '加载失败',
          icon: 'error'
        })
      } finally {
        uni.hideLoading()
      }
    },
    
    // 加载内容详情
    async loadContent() {
      try {
        const res = await uni.request({
          url: `/api/v1/miniprogram/learning/contents/${this.contentId}`,
          method: 'GET',
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.content = res.data.data
        }
      } catch (error) {
        console.error('加载内容失败:', error)
        throw error
      }
    },
    
    // 加载学习记录
    async loadLearningRecord() {
      try {
        const res = await uni.request({
          url: `/api/v1/miniprogram/learning/record/${this.contentId}`,
          method: 'GET',
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.learningRecord = res.data.data
        }
      } catch (error) {
        console.error('加载学习记录失败:', error)
      }
    },
    
    // 加载相关内容
    async loadRelatedContent() {
      try {
        const res = await uni.request({
          url: '/api/v1/miniprogram/learning/recommendations',
          method: 'GET',
          data: {
            contentId: this.contentId,
            type: 'similar',
            limit: 3
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.relatedContents = res.data.data.contents || []
        }
      } catch (error) {
        console.error('加载相关内容失败:', error)
      }
    },
    
    // 开始学习记录
    async startLearningRecord() {
      try {
        await uni.request({
          url: '/api/v1/miniprogram/learning/start',
          method: 'POST',
          data: {
            contentId: this.contentId
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
      } catch (error) {
        console.error('开始学习记录失败:', error)
      }
    },
    
    // 返回上一页
    goBack() {
      uni.navigateBack()
    },
    
    // 切换收藏状态
    async toggleCollect() {
      try {
        const res = await uni.request({
          url: '/api/v1/miniprogram/learning/collect',
          method: 'POST',
          data: {
            contentId: this.contentId
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.content.isCollected = !this.content.isCollected
          uni.showToast({
            title: this.content.isCollected ? '收藏成功' : '取消收藏',
            icon: 'success'
          })
        }
      } catch (error) {
        console.error('收藏操作失败:', error)
        uni.showToast({
          title: '操作失败',
          icon: 'error'
        })
      }
    },
    
    // 切换点赞状态
    async toggleLike() {
      try {
        const res = await uni.request({
          url: '/api/v1/miniprogram/learning/like',
          method: 'POST',
          data: {
            contentId: this.contentId
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.content.isLiked = !this.content.isLiked
          this.content.likeCount += this.content.isLiked ? 1 : -1
          uni.showToast({
            title: this.content.isLiked ? '点赞成功' : '取消点赞',
            icon: 'success'
          })
        }
      } catch (error) {
        console.error('点赞操作失败:', error)
        uni.showToast({
          title: '操作失败',
          icon: 'error'
        })
      }
    },
    
    // 分享内容
    shareContent() {
      uni.share({
        provider: 'weixin',
        type: 0,
        title: this.content.title,
        summary: this.content.summary,
        imageUrl: this.content.thumbnailUrl,
        success: () => {
          uni.showToast({
            title: '分享成功',
            icon: 'success'
          })
        }
      })
    },
    
    // 标记为已完成
    async markAsCompleted() {
      try {
        const res = await uni.request({
          url: '/api/v1/miniprogram/learning/complete',
          method: 'POST',
          data: {
            contentId: this.contentId,
            progress: 100
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.learningRecord = {
            ...this.learningRecord,
            isCompleted: true,
            progress: 100,
            completedAt: new Date().toISOString()
          }
          
          uni.showToast({
            title: '恭喜完成学习！',
            icon: 'success'
          })
        }
      } catch (error) {
        console.error('完成学习失败:', error)
        uni.showToast({
          title: '操作失败',
          icon: 'error'
        })
      }
    },
    
    // 重新学习
    async restartLearning() {
      try {
        await this.startLearningRecord()
        this.learningRecord.isCompleted = false
        this.learningRecord.progress = 0
        
        uni.showToast({
          title: '开始重新学习',
          icon: 'success'
        })
      } catch (error) {
        console.error('重新学习失败:', error)
      }
    },
    
    // 音频播放控制
    toggleAudioPlay() {
      if (this.isAudioPlaying) {
        this.pauseAudio()
      } else {
        this.playAudio()
      }
    },
    
    playAudio() {
      if (!this.audioContext) {
        this.audioContext = uni.createInnerAudioContext()
        this.audioContext.src = this.content.audioUrl
        
        this.audioContext.onPlay(() => {
          this.isAudioPlaying = true
        })
        
        this.audioContext.onPause(() => {
          this.isAudioPlaying = false
        })
        
        this.audioContext.onTimeUpdate(() => {
          this.audioCurrentTime = this.audioContext.currentTime
          this.audioDuration = this.audioContext.duration
          this.audioProgress = (this.audioCurrentTime / this.audioDuration) * 100
        })
        
        this.audioContext.onEnded(() => {
          this.isAudioPlaying = false
          this.markAsCompleted()
        })
      }
      
      this.audioContext.play()
    },
    
    pauseAudio() {
      if (this.audioContext) {
        this.audioContext.pause()
      }
    },
    
    // 视频播放事件
    onVideoPlay() {
      console.log('视频开始播放')
    },
    
    onVideoPause() {
      console.log('视频暂停')
    },
    
    onVideoTimeUpdate(e) {
      const { currentTime, duration } = e.detail
      if (duration > 0) {
        const progress = (currentTime / duration) * 100
        this.updateLearningProgress(progress)
      }
    },
    
    // 更新学习进度
    async updateLearningProgress(progress) {
      try {
        await uni.request({
          url: '/api/v1/miniprogram/learning/progress',
          method: 'POST',
          data: {
            contentId: this.contentId,
            progress: progress
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (this.learningRecord) {
          this.learningRecord.progress = progress
        }
      } catch (error) {
        console.error('更新学习进度失败:', error)
      }
    },
    
    // 显示评分弹窗
    rateContent() {
      this.showRatingModal = true
      this.currentRating = 0
    },
    
    // 隐藏评分弹窗
    hideRating() {
      this.showRatingModal = false
    },
    
    // 设置评分
    setRating(rating) {
      this.currentRating = rating
    },
    
    // 提交评分
    async submitRating() {
      if (this.currentRating === 0) {
        uni.showToast({
          title: '请选择评分',
          icon: 'none'
        })
        return
      }
      
      try {
        const res = await uni.request({
          url: '/api/v1/miniprogram/learning/rate',
          method: 'POST',
          data: {
            contentId: this.contentId,
            rating: this.currentRating
          },
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          }
        })
        
        if (res.data.code === 0) {
          this.hideRating()
          uni.showToast({
            title: '评分成功',
            icon: 'success'
          })
        }
      } catch (error) {
        console.error('评分失败:', error)
        uni.showToast({
          title: '评分失败',
          icon: 'error'
        })
      }
    },
    
    // 查看相关内容
    viewRelated(item) {
      uni.redirectTo({
        url: `/pages/learning/detail?id=${item.id}`
      })
    },
    
    // 获取分类名称
    getCategoryName(category) {
      const categoryMap = {
        1: '科普知识',
        2: '康复指导', 
        3: '心理健康',
        4: '经验分享'
      }
      return categoryMap[category] || '其他'
    },
    
    // 格式化数字
    formatNumber(num) {
      if (num >= 1000) {
        return (num / 1000).toFixed(1) + 'k'
      }
      return num.toString()
    },
    
    // 格式化日期
    formatDate(dateStr) {
      const date = new Date(dateStr)
      return `${date.getMonth() + 1}月${date.getDate()}日`
    },
    
    // 格式化日期时间
    formatDateTime(dateStr) {
      const date = new Date(dateStr)
      return `${date.getMonth() + 1}月${date.getDate()}日 ${date.getHours()}:${date.getMinutes().toString().padStart(2, '0')}`
    },
    
    // 格式化时间
    formatTime(seconds) {
      const mins = Math.floor(seconds / 60)
      const secs = Math.floor(seconds % 60)
      return `${mins}:${secs.toString().padStart(2, '0')}`
    },
    
    // 更新时间
    updateTime() {
      const now = new Date()
      this.currentTime = `${now.getHours()}:${now.getMinutes().toString().padStart(2, '0')}`
    }
  }
}
</script>

<style lang="scss">
.learning-detail {
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

/* 头部导航 */
.header {
  background: white;
  padding: 32rpx 48rpx;
  border-bottom: 2rpx solid #F3F4F6;
  display: flex;
  align-items: center;
  justify-content: space-between;
  
  .back-btn {
    width: 80rpx;
    height: 80rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    
    .back-icon {
      font-size: 40rpx;
      color: #1F2937;
    }
  }
  
  .header-title {
    font-size: 36rpx;
    font-weight: 600;
    color: #1F2937;
  }
  
  .actions {
    display: flex;
    gap: 16rpx;
    
    .action-btn {
      width: 80rpx;
      height: 80rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      
      .action-icon {
        font-size: 32rpx;
        color: #6B7280;
      }
    }
  }
}

/* 主要内容 */
.main-content {
  padding: 48rpx;
  height: calc(100vh - 400rpx);
}

/* 内容头部 */
.content-header {
  margin-bottom: 32rpx;
  
  .category-badge {
    display: inline-block;
    padding: 8rpx 16rpx;
    border-radius: 20rpx;
    font-size: 24rpx;
    margin-bottom: 16rpx;
    
    &.category-1 {
      background: #DBEAFE;
      color: #1D4ED8;
    }
    
    &.category-2 {
      background: #D1FAE5;
      color: #059669;
    }
    
    &.category-3 {
      background: #FEF3C7;
      color: #D97706;
    }
    
    &.category-4 {
      background: #FCE7F3;
      color: #BE185D;
    }
  }
  
  .content-title {
    display: block;
    font-size: 48rpx;
    font-weight: bold;
    color: #1F2937;
    line-height: 1.3;
    margin-bottom: 24rpx;
  }
  
  .content-meta {
    display: flex;
    flex-wrap: wrap;
    gap: 24rpx;
    margin-bottom: 24rpx;
    
    .meta-item {
      font-size: 26rpx;
      color: #6B7280;
    }
  }
  
  .content-summary {
    font-size: 30rpx;
    color: #4B5563;
    line-height: 1.6;
    display: block;
  }
}

/* 视频播放器 */
.video-player {
  margin-bottom: 32rpx;
  border-radius: 24rpx;
  overflow: hidden;
  
  .video-element {
    width: 100%;
    height: 400rpx;
  }
}

/* 音频播放器 */
.audio-player {
  background: white;
  border-radius: 32rpx;
  padding: 32rpx;
  margin-bottom: 32rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
  
  .audio-controls {
    display: flex;
    gap: 32rpx;
    align-items: center;
    
    .audio-cover {
      position: relative;
      flex-shrink: 0;
      
      .cover-image {
        width: 160rpx;
        height: 160rpx;
        border-radius: 24rpx;
      }
      
      .play-btn {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 80rpx;
        height: 80rpx;
        background: rgba(0, 0, 0, 0.7);
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        
        .play-icon {
          font-size: 32rpx;
          color: white;
        }
      }
    }
    
    .audio-info {
      flex: 1;
      
      .audio-title {
        font-size: 32rpx;
        font-weight: 600;
        color: #1F2937;
        margin-bottom: 24rpx;
        display: block;
      }
      
      .progress-container {
        display: flex;
        align-items: center;
        gap: 16rpx;
        
        .time-current,
        .time-total {
          font-size: 24rpx;
          color: #6B7280;
          width: 80rpx;
        }
        
        .progress-bar {
          flex: 1;
          height: 8rpx;
          background: #E5E7EB;
          border-radius: 4rpx;
          position: relative;
          
          .progress-fill {
            height: 100%;
            background: #34D399;
            border-radius: 4rpx;
            transition: width 0.1s ease;
          }
        }
      }
    }
  }
}

/* 文章内容 */
.article-content {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 32rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
  
  .content-body {
    font-size: 32rpx;
    line-height: 1.8;
    color: #1F2937;
  }
}

/* 文本内容 */
.text-content {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 32rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
  
  .section-title {
    font-size: 36rpx;
    font-weight: 600;
    color: #1F2937;
    margin-bottom: 24rpx;
    display: block;
  }
  
  .content-text {
    font-size: 30rpx;
    line-height: 1.6;
    color: #4B5563;
  }
}

/* 学习进度 */
.learning-progress {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 32rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
  
  .section-title {
    font-size: 36rpx;
    font-weight: 600;
    color: #1F2937;
    margin-bottom: 32rpx;
    display: block;
  }
  
  .progress-info {
    .progress-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 16rpx;
      
      .progress-label {
        font-size: 28rpx;
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
      margin-bottom: 24rpx;
      
      .progress-fill {
        height: 100%;
        background: #34D399;
        border-radius: 8rpx;
        transition: width 0.3s ease;
      }
    }
    
    .progress-details {
      .detail-item {
        display: block;
        font-size: 26rpx;
        color: #6B7280;
        margin-bottom: 8rpx;
      }
    }
  }
}

/* 互动区域 */
.interaction-section {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 32rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
  
  .action-buttons {
    display: flex;
    justify-content: space-around;
    
    .action-item {
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 8rpx;
      
      .action-icon {
        font-size: 48rpx;
        
        &.liked {
          animation: heartbeat 0.6s ease-in-out;
        }
        
        &.collected {
          color: #F59E0B;
        }
      }
      
      .action-text {
        font-size: 24rpx;
        color: #6B7280;
      }
    }
  }
}

/* 相关推荐 */
.related-content {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 32rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
  
  .section-title {
    font-size: 36rpx;
    font-weight: 600;
    color: #1F2937;
    margin-bottom: 32rpx;
    display: block;
  }
  
  .related-list {
    .related-item {
      display: flex;
      gap: 24rpx;
      padding: 24rpx 0;
      border-bottom: 2rpx solid #F3F4F6;
      
      &:last-child {
        border-bottom: none;
      }
      
      .related-image {
        width: 120rpx;
        height: 90rpx;
        border-radius: 16rpx;
        flex-shrink: 0;
      }
      
      .related-info {
        flex: 1;
        
        .related-title {
          font-size: 28rpx;
          font-weight: 500;
          color: #1F2937;
          margin-bottom: 8rpx;
          display: block;
          line-height: 1.4;
        }
        
        .related-meta {
          display: flex;
          gap: 16rpx;
          
          .meta-text {
            font-size: 24rpx;
            color: #9CA3AF;
          }
        }
      }
    }
  }
}

/* 底部操作栏 */
.bottom-actions {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: white;
  padding: 32rpx 48rpx;
  border-top: 2rpx solid #F3F4F6;
  
  .progress-action {
    .complete-btn {
      width: 100%;
    }
  }
  
  .completed-action {
    display: flex;
    align-items: center;
    justify-content: space-between;
    
    .completed-text {
      font-size: 28rpx;
      color: #10B981;
      font-weight: 500;
    }
    
    .restart-btn {
      width: 200rpx;
    }
  }
}

/* 评分弹窗 */
.rating-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
  
  .rating-content {
    background: white;
    border-radius: 32rpx;
    padding: 48rpx;
    margin: 0 48rpx;
    text-align: center;
    
    .rating-title {
      font-size: 36rpx;
      font-weight: 600;
      color: #1F2937;
      margin-bottom: 32rpx;
      display: block;
    }
    
    .stars {
      display: flex;
      justify-content: center;
      gap: 16rpx;
      margin-bottom: 48rpx;
      
      .star {
        font-size: 60rpx;
        color: #E5E7EB;
        transition: color 0.2s ease;
        
        &.active {
          color: #F59E0B;
        }
      }
    }
    
    .rating-actions {
      display: flex;
      gap: 24rpx;
      
      button {
        flex: 1;
      }
    }
  }
}

/* 动画 */
@keyframes heartbeat {
  0% { transform: scale(1); }
  50% { transform: scale(1.2); }
  100% { transform: scale(1); }
}
</style> 