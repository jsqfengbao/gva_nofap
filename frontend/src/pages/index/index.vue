<template>
  <view class="home-page">
    <!-- 状态栏占位 -->
    <view class="status-bar-spacer"></view>
    
    <!-- 头部问候区域 - 毛玻璃效果 -->
    <view class="header-section">
      <view class="header-bg"></view>
      <view class="greeting-area">
        <view class="greeting-text">
          <text class="greeting">{{ greeting }}, {{ userName }} 👋</text>
          <text class="motto">「自律给你自由，坚持改变人生」</text>
        </view>
        <view class="header-actions">
          <view class="notification-btn" @click="goToNotifications">
            <text class="icon">🔔</text>
            <view class="notification-dot" v-if="hasNotification"></view>
          </view>
          <view class="avatar-circle" @click="goToProfile">
            <image v-if="userAvatar" class="avatar-img" :src="userAvatar" mode="aspectFill"></image>
            <text v-else class="level-text">L{{ userLevel }}</text>
          </view>
        </view>
      </view>
      
      <!-- 统计概览 -->
      <view class="stats-overview">
        <view class="stat-item">
          <text class="stat-number">{{ userStats.totalDays }}</text>
          <text class="stat-label">总天数</text>
        </view>
        <view class="stat-divider"></view>
        <view class="stat-item">
          <text class="stat-number">{{ userStats.successRate }}%</text>
          <text class="stat-label">成功率</text>
        </view>
        <view class="stat-divider"></view>
        <view class="stat-item">
          <text class="stat-number">{{ userStats.longestStreak }}</text>
          <text class="stat-label">最长纪录</text>
        </view>
      </view>
    </view>

    <!-- 主要内容区域 -->
    <view class="main-content">
      <!-- 连续坚持卡片 - 主视觉 -->
      <view class="streak-card" @click="goToCheckin">
        <view class="streak-card-bg"></view>
        <view class="streak-content">
          <view class="streak-header">
            <view class="streak-title-wrap">
              <text class="streak-icon">🔥</text>
              <text class="streak-title">当前连续坚持</text>
            </view>
            <view class="streak-badge" v-if="streakDays > 0">
              <text class="streak-badge-text">坚持中</text>
            </view>
          </view>
          
          <view class="streak-main">
            <text class="streak-number">{{ streakDays }}</text>
            <text class="streak-unit">天</text>
          </view>
          
          <view class="milestone-progress" v-if="daysToMilestone > 0">
            <view class="progress-info">
              <text class="progress-label">距离 {{ nextMilestone }} 天里程碑</text>
              <text class="progress-days">还需 {{ daysToMilestone }} 天</text>
            </view>
            <view class="progress-bar">
              <view class="progress-fill" :style="{ width: milestoneProgress + '%' }"></view>
            </view>
          </view>
          
          <view class="checkin-button" :class="{ 'checked': hasCheckedToday }" @click.stop="handleCheckin">
            <text class="checkin-text">
              {{ hasCheckedToday ? '✅ 今日已打卡' : '👉 点击今日打卡' }}
            </text>
          </view>
        </view>
      </view>

      <!-- 快捷功能网格 -->
      <view class="quick-grid">
        <view class="grid-item item-emerald" @click="goToCheckin">
          <view class="grid-icon">✅</view>
          <text class="grid-title">每日打卡</text>
          <text class="grid-desc">记录坚持</text>
        </view>
        
        <view class="grid-item item-blue" @click="goToEmergency">
          <view class="grid-icon">🛑</view>
          <text class="grid-title">紧急刹车</text>
          <text class="grid-desc">立即止损</text>
        </view>
        
        <view class="grid-item item-amber" @click="goToAchievement">
          <view class="grid-icon">🏆</view>
          <text class="grid-title">成就墙</text>
          <text class="grid-desc">收集徽章</text>
        </view>
        
        <view class="grid-item item-purple" @click="goToLearning">
          <view class="grid-icon">📚</view>
          <text class="grid-title">成长学院</text>
          <text class="grid-desc">知识改变认知</text>
        </view>
        
        <view class="grid-item item-orange" @click="goToAssessment">
          <view class="grid-icon">📊</view>
          <text class="grid-title">色隐测试</text>
          <text class="grid-desc">评估上瘾程度</text>
        </view>
        
        <view class="grid-item item-pink" @click="goToCommunity">
          <view class="grid-icon">👥</view>
          <text class="grid-title">互助社区</text>
          <text class="grid-desc">同行不孤独</text>
        </view>
      </view>

      <!-- 等级进度卡片 -->
      <view class="level-card" @click="goToAchievement">
        <view class="level-header">
          <view class="level-left">
            <view class="level-icon">
              <text class="icon">⭐</text>
            </view>
            <view class="level-info">
              <text class="level-title">等级 {{ userLevel }} · {{ levelTitle }}</text>
              <text class="exp-text">{{ currentExp }} / {{ nextLevelExp }} 经验值</text>
            </view>
          </view>
          <view class="level-badge">
            <text class="level-badge-text">{{ levelProgress }}%</text>
          </view>
        </view>
        <view class="level-progress-bar">
          <view class="level-progress-fill" :style="{ width: levelProgress + '%' }"></view>
        </view>
      </view>

      <!-- 每日格言 -->
      <view class="quote-card">
        <view class="quote-header">
          <view class="quote-icon">💭</view>
          <text class="quote-title">今日格言</text>
        </view>
        <text class="quote-content">{{ dailyQuote }}</text>
        <view class="quote-footer">
          <text class="quote-author">— 戒色修行路</text>
        </view>
      </view>

      <!-- 今日成就 -->
      <view class="achievements-card" v-if="todayAchievements.length > 0">
        <view class="card-header">
          <text class="card-title">🎯 今日解锁</text>
          <text class="card-more" @click="goToAchievement">查看全部 &gt;</text>
        </view>
        <view class="achievements-list">
          <view class="achievement-item" v-for="achievement in todayAchievements" :key="achievement.id">
            <view class="achievement-icon">
              <text class="icon">⭐</text>
            </view>
            <view class="achievement-details">
              <text class="achievement-name">{{ achievement.name }}</text>
              <text class="achievement-reward">+{{ achievement.rewards }} EXP</text>
            </view>
            <view class="achievement-plus">
              <text class="plus-text">+{{ achievement.rewards }}</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 最近成就 -->
      <view class="recent-card" v-if="recentAchievements.length > 0">
        <view class="card-header">
          <text class="card-title">🏆 最近获得</text>
        </view>
        <view class="recent-list">
          <view class="recent-item" v-for="achievement in recentAchievements" :key="achievement.id">
            <view class="recent-icon">
              <text class="icon">🎖️</text>
            </view>
            <view class="recent-info">
              <text class="recent-name">{{ achievement.name }}</text>
              <text class="recent-time">{{ formatTime(achievement.unlockedAt) }}</text>
            </view>
          </view>
        </view>
      </view>
      
      <!-- 底部安全间距 -->
      <view class="bottom-spacer"></view>
    </view>
  </view>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { onShow } from '@dcloudio/uni-app'
import homeApi from '@/apis/home.js'

// 响应式数据
const currentTime = ref('')
const userLevel = ref(1)
const streakDays = ref(0)
const hasCheckedToday = ref(false)
const currentExp = ref(0)
const nextLevelExp = ref(100)
const todayAchievements = ref([])
const recentAchievements = ref([])

const userStats = reactive({
  level: 1,
  experience: 0,
  currentStreak: 0,
  longestStreak: 0,
  totalDays: 0,
  successRate: 0
})

const gameStats = reactive({
  currentLevel: 1,
  currentExp: 0,
  nextLevelExp: 100,
  levelProgress: 0,
  totalAchievements: 0,
  unlockedAchievements: 0,
  achievementRate: 0,
  recentAchievements: []
})

// 计算属性
const greeting = computed(() => {
  const hour = new Date().getHours()
  if (hour < 12) return '早上好'
  if (hour < 18) return '下午好'
  return '晚上好'
})

const levelTitle = computed(() => {
  const level = userLevel.value
  if (level >= 50) return "传奇大师"
  if (level >= 40) return "宗师"
  if (level >= 30) return "专家"
  if (level >= 25) return "资深导师"
  if (level >= 20) return "高级导师"
  if (level >= 15) return "导师"
  if (level >= 10) return "进阶行者"
  if (level >= 5) return "初级学徒"
  if (level >= 1) return "新手"
  return "未知"
})

const levelProgress = computed(() => {
  if (nextLevelExp.value <= 0) return 0
  return Math.round((currentExp.value / nextLevelExp.value) * 100)
})

const daysToMilestone = computed(() => {
  const milestones = [7, 14, 21, 30, 60, 90, 180, 365]
  const current = streakDays.value
  for (let milestone of milestones) {
    if (current < milestone) {
      return milestone - current
    }
  }
  return 30 // 默认下一个30天里程碑
})

const milestoneProgress = computed(() => {
  const milestones = [7, 14, 21, 30, 60, 90, 180, 365]
  const current = streakDays.value
  let previousMilestone = 0
  
  for (let milestone of milestones) {
    if (current < milestone) {
      const progress = ((current - previousMilestone) / (milestone - previousMilestone)) * 100
      return Math.max(0, Math.min(100, progress))
    }
    previousMilestone = milestone
  }
  return 100
})

const dailyTip = ref('当感到冲动时，尝试做10个深呼吸或者快速做20个俯卧撑。身体运动能有效转移注意力并释放内啡肽。')

// 方法
const updateTime = () => {
  const now = new Date()
  const hours = now.getHours().toString().padStart(2, '0')
  const minutes = now.getMinutes().toString().padStart(2, '0')
  currentTime.value = `${hours}:${minutes}`
}

/**
 * 加载用户数据
 * 使用统一的 API 模块获取数据，避免硬编码域名
 */
const loadUserData = async () => {
  try {
    const token = uni.getStorageSync('token')
    if (!token) {
      console.log('用户未登录')
      // #ifdef H5
      // H5环境下跳转到登录页面
      uni.navigateTo({
        url: '/pages/auth/login'
      })
      // #endif
      return
    }

    // 使用优化的并发请求获取所有首页数据
    const result = await homeApi.getHomeData()
    
    if (result.code === 0) {
      const { userStats: stats, todayStatus, gameStats: game } = result.data
      
      // 更新用户统计数据
      if (stats) {
        Object.assign(userStats, stats)
        streakDays.value = stats.currentStreak || 0
        userLevel.value = stats.level || 1
      }
      
      // 更新今日打卡状态
      if (todayStatus) {
        hasCheckedToday.value = todayStatus.hasChecked || false
      }
      
      // 更新游戏化数据
      if (game) {
        Object.assign(gameStats, game)
        userLevel.value = game.currentLevel || 1
        currentExp.value = game.currentExp || 0
        nextLevelExp.value = game.nextLevelExp || 100
        recentAchievements.value = game.recentAchievements || []
      }
    } else {
      console.error('获取首页数据失败:', result.message)
      uni.showToast({
        title: '数据加载失败',
        icon: 'none'
      })
    }

  } catch (error) {
    console.error('加载用户数据失败:', error)
    uni.showToast({
      title: '网络错误，请重试',
      icon: 'none'
    })
  }
}

const formatTime = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  const now = new Date()
  const diffTime = now - date
  const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays === 0) return '今天获得'
  if (diffDays === 1) return '昨天获得'
  if (diffDays < 7) return `${diffDays}天前获得`
  return `${Math.floor(diffDays / 7)}周前获得`
}

const handleCheckin = () => {
  if (hasCheckedToday.value) {
    goToCheckin()
  } else {
    goToCheckin()
  }
}

// 导航方法
const goToCheckin = () => {
  uni.switchTab({ url: '/pages/checkin/index' })
}

const goToAchievement = () => {
  uni.navigateTo({ url: '/pages/achievement/index' })
}

const goToAssessment = () => {
  uni.navigateTo({ url: '/pages/assessment/index' })
}

const goToCommunity = () => {
  uni.switchTab({ url: '/pages/community/index' })
}

const goToLearning = () => {
  uni.switchTab({ url: '/pages/learning/index' })
}

const goToEmergency = () => {
  uni.navigateTo({ url: '/pages/emergency/index' })
}

const goToHistory = () => {
  uni.navigateTo({ url: '/pages/checkin/history' })
}

const goToProgress = () => {
  uni.navigateTo({ url: '/pages/progress/index' })
}

const goToProfile = () => {
  uni.switchTab({ url: '/pages/profile/index' })
}

const goToNotifications = () => {
  uni.showToast({
    title: '通知功能开发中',
    icon: 'none'
  })
}

// 生命周期
onMounted(() => {
  updateTime()
  setInterval(updateTime, 60000) // 每分钟更新一次时间
  loadUserData()
})

onShow(() => {
  loadUserData()
})
</script>

<style lang="scss" scoped>
.home-page {
  min-height: 100vh;
  background: #F8FAFC;
}

.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16rpx 32rpx;
  font-size: 24rpx;
  color: #1f2937;
  background: transparent;
  
  .status-time {
    font-weight: 500;
  }
  
  .status-icons {
    display: flex;
    align-items: center;
    gap: 8rpx;
    
    .icon {
      font-size: 24rpx;
    }
    
    .battery {
      width: 48rpx;
      height: 24rpx;
      border: 2rpx solid #1f2937;
      border-radius: 4rpx;
      position: relative;
      
      .battery-fill {
        width: 32rpx;
        height: 12rpx;
        background-color: #10b981;
        border-radius: 2rpx;
        margin: 4rpx;
      }
    }
  }
}

.header-section {
  background: #FFFFFF;
  padding: 32rpx 48rpx 48rpx;
  
  .greeting-area {
    display: flex;
    align-items: center;
    justify-content: space-between;
    
    .greeting-text {
      .greeting {
        display: block;
        font-size: 48rpx;
        font-weight: bold;
        color: #1f2937;
        margin-bottom: 8rpx;
      }
      
      .motto {
        font-size: 28rpx;
        color: #6b7280;
      }
    }
    
    .header-actions {
      display: flex;
      align-items: center;
      gap: 24rpx;
      
      .notification-btn {
        width: 80rpx;
        height: 80rpx;
        background: #f3f4f6;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all 0.3s ease;
        
        &:active {
          background: #e5e7eb;
          transform: scale(0.95);
        }
        
        .icon {
          font-size: 28rpx;
        }
      }
      
      .level-avatar {
        width: 80rpx;
        height: 80rpx;
        background: linear-gradient(135deg, #34D399 0%, #10B981 100%);
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all 0.3s ease;
        
        &:active {
          transform: scale(0.95);
        }
        
        .level-text {
          color: white;
          font-weight: bold;
          font-size: 28rpx;
        }
      }
    }
  }
}

/* 统计概览样式 */
.stats-overview {
  display: flex;
  align-items: center;
  justify-content: space-around;
  background: linear-gradient(135deg, #34D399 0%, #10B981 100%);
  border-radius: 32rpx;
  padding: 32rpx 24rpx;
  margin: 0 48rpx 48rpx;
  box-shadow: 0 8rpx 32rpx rgba(52, 211, 153, 0.2);
}

.stats-overview .stat-item {
  flex: 1;
  text-align: center;
}

.stats-overview .stat-item .stat-number {
  display: block;
  font-size: 48rpx;
  font-weight: bold;
  color: white;
  margin-bottom: 8rpx;
}

.stats-overview .stat-item .stat-label {
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.9);
}

.stats-overview .stat-divider {
  width: 2rpx;
  height: 60rpx;
  background: rgba(255, 255, 255, 0.3);
}

.main-content {
  padding: 0 48rpx 160rpx;
}

.streak-card {
  background: linear-gradient(135deg, #34D399 0%, #10B981 100%);
  border-radius: 60rpx;
  padding: 48rpx;
  margin-bottom: 48rpx;
  box-shadow: 0 0 60rpx rgba(52, 211, 153, 0.3);
  
  .streak-content {
    text-align: center;
    color: white;
    
    .streak-icon {
      font-size: 80rpx;
      margin-bottom: 16rpx;
    }
    
    .streak-number {
      display: block;
      font-size: 72rpx;
      font-weight: bold;
      margin-bottom: 8rpx;
    }
    
    .streak-label {
      font-size: 28rpx;
      color: rgba(255, 255, 255, 0.9);
      margin-bottom: 32rpx;
    }
    
    .milestone-progress {
      background: rgba(255, 255, 255, 0.2);
      border-radius: 32rpx;
      padding: 32rpx;
      margin-bottom: 32rpx;
      
      .progress-info {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 16rpx;
        
        .progress-label {
          font-size: 28rpx;
        }
        
        .progress-days {
          font-size: 28rpx;
          font-weight: 600;
        }
      }
      
      .progress-bar {
        width: 100%;
        height: 16rpx;
        background: rgba(255, 255, 255, 0.3);
        border-radius: 8rpx;
        overflow: hidden;
        
        .progress-fill {
          height: 100%;
          background: white;
          border-radius: 8rpx;
          transition: width 0.3s ease;
        }
      }
    }
    
    .checkin-button {
      width: 100%;
      background: white;
      color: #34D399;
      font-weight: bold;
      padding: 24rpx;
      border-radius: 32rpx;
      transition: all 0.3s ease;
      
      &:active {
        transform: scale(0.98);
      }
      
      &.checked {
        background: rgba(255, 255, 255, 0.9);
        color: #10B981;
      }
      
      .checkin-text {
        font-size: 32rpx;
      }
    }
  }
}

.level-card {
  background: #FFFFFF;
  border-radius: 32rpx;
  padding: 32rpx;
  margin-bottom: 48rpx;
  box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.05);
  
  .level-info {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 24rpx;
    
    .level-icon {
      width: 96rpx;
      height: 96rpx;
      background: linear-gradient(135deg, #F59E0B 0%, #F97316 100%);
      border-radius: 24rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 24rpx;
      
      .icon {
        font-size: 36rpx;
      }
    }
    
    .level-details {
      flex: 1;
      
      .level-title {
        display: block;
        font-size: 32rpx;
        font-weight: 600;
        color: #1f2937;
        margin-bottom: 8rpx;
      }
      
      .exp-text {
        font-size: 28rpx;
        color: #6b7280;
      }
    }
    
    .game-emoji {
      font-size: 48rpx;
    }
  }
  
  .level-progress-bar {
    width: 100%;
    height: 24rpx;
    background: #e5e7eb;
    border-radius: 12rpx;
    overflow: hidden;
    
    .level-progress-fill {
      height: 100%;
      background: linear-gradient(90deg, #34D399 0%, #10B981 100%);
      border-radius: 12rpx;
      transition: width 0.3s ease;
    }
  }
}

.quick-actions {
  display: flex;
  gap: 32rpx;
  margin-bottom: 48rpx;
  
  .action-card {
    flex: 1;
    background: #FFFFFF;
    border-radius: 32rpx;
    padding: 32rpx;
    text-align: center;
    box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.05);
    transition: all 0.3s ease;
    
    &:active {
      transform: scale(0.98);
    }
    
    .action-icon {
      width: 96rpx;
      height: 96rpx;
      border-radius: 24rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      margin: 0 auto 24rpx;
      
      .icon {
        font-size: 40rpx;
      }
    }
    
    &.emergency .action-icon {
      background: rgba(239, 68, 68, 0.1);
    }
    
    &.community .action-icon {
      background: rgba(6, 182, 212, 0.1);
    }
    
    .action-title {
      display: block;
      font-size: 32rpx;
      font-weight: 600;
      color: #1f2937;
      margin-bottom: 8rpx;
    }
    
    .action-desc {
      font-size: 24rpx;
      color: #6b7280;
    }
  }
}

.achievements-card, .tips-card, .recent-achievements {
  background: #FFFFFF;
  border-radius: 32rpx;
  padding: 32rpx;
  margin-bottom: 48rpx;
  box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.05);
  
  .card-title {
    display: block;
    font-size: 32rpx;
    font-weight: 600;
    color: #1f2937;
    margin-bottom: 32rpx;
  }
}

.achievements-list {
  .achievement-item {
    display: flex;
    align-items: center;
    padding: 24rpx;
    background: #f0fdf4;
    border-radius: 24rpx;
    margin-bottom: 24rpx;
    
    &:last-child {
      margin-bottom: 0;
    }
    
    .achievement-icon {
      width: 64rpx;
      height: 64rpx;
      background: rgba(34, 197, 94, 0.2);
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 24rpx;
      
      .icon {
        font-size: 28rpx;
      }
    }
    
    .achievement-details {
      flex: 1;
      
      .achievement-name {
        display: block;
        font-size: 32rpx;
        font-weight: 500;
        color: #1f2937;
        margin-bottom: 4rpx;
      }
      
      .achievement-reward {
        font-size: 24rpx;
        color: #6b7280;
      }
    }
    
    .reward-value {
      font-size: 32rpx;
      font-weight: bold;
      color: #34D399;
    }
  }
}

.tips-card {
  background: linear-gradient(90deg, #faf5ff 0%, #fdf2f8 100%);
  border: 2rpx solid #e879f9;
  
  .tips-header {
    display: flex;
    align-items: flex-start;
    margin-bottom: 16rpx;
    
    .tips-icon {
      width: 64rpx;
      height: 64rpx;
      background: #a855f7;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 24rpx;
      
      .icon {
        font-size: 28rpx;
      }
    }
    
    .tips-title {
      font-size: 32rpx;
      font-weight: 600;
      color: #1f2937;
      margin-top: 8rpx;
    }
  }
  
  .tips-content {
    font-size: 28rpx;
    color: #374151;
    line-height: 1.6;
    margin-left: 88rpx;
  }
}

.recent-list {
  .recent-item {
    display: flex;
    align-items: center;
    margin-bottom: 24rpx;
    
    &:last-child {
      margin-bottom: 0;
    }
    
    .recent-icon {
      width: 80rpx;
      height: 80rpx;
      background: linear-gradient(135deg, #fbbf24 0%, #f59e0b 100%);
      border-radius: 24rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 24rpx;
      
      .icon {
        font-size: 32rpx;
      }
    }
    
    .recent-details {
      .recent-name {
        display: block;
        font-size: 32rpx;
        font-weight: 500;
        color: #1f2937;
        margin-bottom: 4rpx;
      }
      
      .recent-time {
        font-size: 24rpx;
        color: #6b7280;
      }
    }
  }
}
</style> 