<template>
  <view class="progress-page">
    <!-- 状态栏 -->
    <view class="status-bar">
      <text class="status-time">{{ currentTime }}</text>
      <view class="status-icons">
        <text class="icon">📶</text>
        <text class="icon">📶</text>
        <view class="battery">
          <view class="battery-fill"></view>
        </view>
      </view>
    </view>

    <!-- 头部标题区域 -->
    <view class="header-section">
      <view class="header-content">
        <text class="page-title">进度追踪</text>
        <view class="export-btn" @click="exportData">
          <text class="icon">📥</text>
        </view>
      </view>
    </view>

    <!-- 主要内容区域 -->
    <view class="main-content">
      <!-- 概览统计卡片 -->
      <view class="overview-stats">
        <view class="stat-card current-streak">
          <view class="stat-content">
            <text class="stat-icon">🔥</text>
            <text class="stat-number">{{ currentStreak }}</text>
            <text class="stat-label">当前连击</text>
          </view>
        </view>
        
        <view class="stat-card longest-streak">
          <view class="stat-content">
            <text class="stat-icon">🏆</text>
            <text class="stat-number">{{ longestStreak }}</text>
            <text class="stat-label">最长连击</text>
          </view>
        </view>
      </view>

      <!-- 本周进度图表 -->
      <view class="weekly-chart-card">
        <text class="card-title">本周进度</text>
        
        <!-- 图表容器 -->
        <view class="chart-container">
          <view class="chart-bars">
            <view 
              class="chart-bar" 
              v-for="(day, index) in weeklyData" 
              :key="index"
            >
              <view 
                class="bar-fill" 
                :class="{ 'active': day.hasChecked, 'inactive': !day.hasChecked }"
                :style="{ height: day.height + '%' }"
              ></view>
              <text class="bar-label">{{ day.label }}</text>
            </view>
          </view>
        </view>
        
        <view class="completion-rate">
          <text class="rate-text">本周完成率: </text>
          <text class="rate-value">{{ weeklyCompletionRate }}%</text>
        </view>
      </view>

      <!-- 成就进度 -->
      <view class="achievement-progress-card">
        <text class="card-title">成就进度</text>
        
        <view class="achievement-list">
          <view 
            class="achievement-item" 
            v-for="achievement in progressAchievements" 
            :key="achievement.id"
          >
            <view class="achievement-icon" :class="achievement.category">
              <text class="icon">{{ achievement.icon }}</text>
            </view>
            <view class="achievement-details">
              <text class="achievement-name">{{ achievement.name }}</text>
              <view class="progress-container">
                <view class="progress-bar">
                  <view 
                    class="progress-fill" 
                    :class="achievement.category"
                    :style="{ width: achievement.progress + '%' }"
                  ></view>
                </view>
                <text class="progress-text">{{ achievement.current }}/{{ achievement.target }}</text>
              </view>
            </view>
          </view>
        </view>
      </view>

      <!-- 本月总结 -->
      <view class="monthly-summary-card">
        <text class="card-title">本月总结</text>
        
        <view class="summary-stats">
          <view class="summary-item">
            <text class="summary-value success-rate">{{ monthlyStats.successRate }}%</text>
            <text class="summary-label">成功率</text>
          </view>
          <view class="summary-item">
            <text class="summary-value learning-days">{{ monthlyStats.learningDays }}</text>
            <text class="summary-label">学习天数</text>
          </view>
          <view class="summary-item">
            <text class="summary-value experience">{{ monthlyStats.experience }}</text>
            <text class="summary-label">经验值</text>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { onShow } from '@dcloudio/uni-app'

// 响应式数据
const currentTime = ref('')
const currentStreak = ref(0)
const longestStreak = ref(0)

const weeklyData = ref([
  { label: '周一', hasChecked: true, height: 75 },
  { label: '周二', hasChecked: true, height: 62 },
  { label: '周三', hasChecked: true, height: 87 },
  { label: '周四', hasChecked: true, height: 50 },
  { label: '周五', hasChecked: true, height: 100 },
  { label: '周六', hasChecked: false, height: 25 },
  { label: '周日', hasChecked: false, height: 25 }
])

const progressAchievements = ref([
  {
    id: 1,
    name: '30天挑战',
    icon: '🏅',
    category: 'streak',
    current: 23,
    target: 30,
    progress: 77
  },
  {
    id: 2,
    name: '社区贡献者',
    icon: '⭐',
    category: 'community',
    current: 6,
    target: 10,
    progress: 60
  },
  {
    id: 3,
    name: '学习达人',
    icon: '📚',
    category: 'learning',
    current: 8,
    target: 20,
    progress: 40
  }
])

const monthlyStats = reactive({
  successRate: 85,
  learningDays: 12,
  experience: 156
})

const userStats = reactive({
  currentStreak: 0,
  longestStreak: 0,
  totalDays: 0,
  successRate: 0
})

// 计算属性
const weeklyCompletionRate = computed(() => {
  const checkedDays = weeklyData.value.filter(day => day.hasChecked).length
  return Math.round((checkedDays / weeklyData.value.length) * 100)
})

// 方法
const updateTime = () => {
  const now = new Date()
  const hours = now.getHours().toString().padStart(2, '0')
  const minutes = now.getMinutes().toString().padStart(2, '0')
  currentTime.value = `${hours}:${minutes}`
}

const loadProgressData = async () => {
  try {
    const token = uni.getStorageSync('token')
    if (!token) {
      console.log('用户未登录')
      return
    }

    // 加载打卡统计数据
    const statsRes = await uni.request({
      url: 'http://localhost:8888/api/v1/miniprogram/checkin/statistics',
      method: 'GET',
      header: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })

    if (statsRes.data.code === 0) {
      const data = statsRes.data.data
      Object.assign(userStats, data)
      currentStreak.value = data.currentStreak || 0
      longestStreak.value = data.longestStreak || 0
    }

    // 加载本周进度数据
    const weeklyRes = await uni.request({
      url: 'http://localhost:8888/api/v1/miniprogram/checkin/weekly-progress',
      method: 'GET',
      header: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })

    if (weeklyRes.data.code === 0) {
      const weeklyProgressData = weeklyRes.data.data
      updateWeeklyChart(weeklyProgressData)
    }

    // 加载成就进度数据
    const achievementsRes = await uni.request({
      url: 'http://localhost:8888/api/v1/miniprogram/achievement/progress',
      method: 'GET',
      header: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })

    if (achievementsRes.data.code === 0) {
      const achievementsData = achievementsRes.data.data
      updateAchievementProgress(achievementsData)
    }

  } catch (error) {
    console.error('加载进度数据失败:', error)
  }
}

const updateWeeklyChart = (data) => {
  if (!data || !data.weeklyCheckins) return
  
  const days = ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
  weeklyData.value = data.weeklyCheckins.map((checkin, index) => ({
    label: days[index],
    hasChecked: checkin.hasChecked || false,
    height: checkin.hasChecked ? Math.max(50, checkin.moodLevel * 20) : 25
  }))
}

const updateAchievementProgress = (achievements) => {
  if (!achievements || !achievements.length) return
  
  progressAchievements.value = achievements.map(achievement => ({
    id: achievement.id,
    name: achievement.name,
    icon: getAchievementIcon(achievement.category),
    category: achievement.category,
    current: achievement.progress || 0,
    target: achievement.targetValue || 100,
    progress: Math.round(((achievement.progress || 0) / (achievement.targetValue || 100)) * 100)
  }))
}

const getAchievementIcon = (category) => {
  const icons = {
    'checkin': '🏅',
    'streak': '🔥',
    'community': '⭐',
    'learning': '📚',
    'level': '🎮'
  }
  return icons[category] || '🏆'
}

const exportData = () => {
  uni.showToast({
    title: '数据导出功能开发中',
    icon: 'none'
  })
}

// 生命周期
onMounted(() => {
  updateTime()
  setInterval(updateTime, 60000) // 每分钟更新一次时间
  loadProgressData()
})

onShow(() => {
  loadProgressData()
})
</script>

<style lang="scss" scoped>
.progress-page {
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
  
  .header-content {
    display: flex;
    align-items: center;
    justify-content: space-between;
    
    .page-title {
      font-size: 48rpx;
      font-weight: bold;
      color: #1f2937;
    }
    
    .export-btn {
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
  }
}

.main-content {
  padding: 0 48rpx 160rpx;
}

.overview-stats {
  display: flex;
  gap: 32rpx;
  margin-bottom: 48rpx;
  
  .stat-card {
    flex: 1;
    border-radius: 32rpx;
    padding: 32rpx;
    
    &.current-streak {
      background: linear-gradient(135deg, #34D399 0%, #10B981 100%);
    }
    
    &.longest-streak {
      background: linear-gradient(135deg, #06B6D4 0%, #0284c7 100%);
    }
    
    .stat-content {
      text-align: center;
      color: white;
      
      .stat-icon {
        display: block;
        font-size: 48rpx;
        margin-bottom: 16rpx;
      }
      
      .stat-number {
        display: block;
        font-size: 48rpx;
        font-weight: bold;
        margin-bottom: 8rpx;
      }
      
      .stat-label {
        font-size: 28rpx;
        opacity: 0.9;
      }
    }
  }
}

.weekly-chart-card, .achievement-progress-card, .monthly-summary-card {
  background: #FFFFFF;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 48rpx;
  box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.05);
  
  .card-title {
    display: block;
    font-size: 36rpx;
    font-weight: 600;
    color: #1f2937;
    margin-bottom: 32rpx;
  }
}

.chart-container {
  margin-bottom: 32rpx;
  
  .chart-bars {
    display: flex;
    align-items: flex-end;
    gap: 16rpx;
    height: 256rpx;
    
    .chart-bar {
      flex: 1;
      display: flex;
      flex-direction: column;
      align-items: center;
      height: 100%;
      
      .bar-fill {
        width: 100%;
        border-radius: 8rpx 8rpx 0 0;
        transition: all 0.3s ease;
        
        &.active {
          background: #34D399;
        }
        
        &.inactive {
          background: #e5e7eb;
        }
      }
      
      .bar-label {
        font-size: 24rpx;
        color: #6b7280;
        margin-top: 16rpx;
      }
    }
  }
}

.completion-rate {
  text-align: center;
  
  .rate-text {
    font-size: 28rpx;
    color: #6b7280;
  }
  
  .rate-value {
    font-size: 28rpx;
    font-weight: 600;
    color: #34D399;
  }
}

.achievement-list {
  .achievement-item {
    display: flex;
    align-items: center;
    margin-bottom: 32rpx;
    
    &:last-child {
      margin-bottom: 0;
    }
    
    .achievement-icon {
      width: 96rpx;
      height: 96rpx;
      border-radius: 24rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 32rpx;
      
      &.streak {
        background: #fef3c7;
      }
      
      &.community {
        background: #dbeafe;
      }
      
      &.learning {
        background: #dcfce7;
      }
      
      .icon {
        font-size: 40rpx;
      }
    }
    
    .achievement-details {
      flex: 1;
      
      .achievement-name {
        display: block;
        font-size: 32rpx;
        font-weight: 600;
        color: #1f2937;
        margin-bottom: 16rpx;
      }
      
      .progress-container {
        display: flex;
        align-items: center;
        gap: 16rpx;
        
        .progress-bar {
          flex: 1;
          height: 16rpx;
          background: #e5e7eb;
          border-radius: 8rpx;
          overflow: hidden;
          
          .progress-fill {
            height: 100%;
            border-radius: 8rpx;
            transition: width 0.3s ease;
            
            &.streak {
              background: #34D399;
            }
            
            &.community {
              background: #06B6D4;
            }
            
            &.learning {
              background: #10B981;
            }
          }
        }
        
        .progress-text {
          font-size: 28rpx;
          color: #6b7280;
          white-space: nowrap;
        }
      }
    }
  }
}

.summary-stats {
  display: flex;
  justify-content: space-around;
  text-align: center;
  
  .summary-item {
    .summary-value {
      display: block;
      font-size: 48rpx;
      font-weight: bold;
      margin-bottom: 8rpx;
      
      &.success-rate {
        color: #34D399;
      }
      
      &.learning-days {
        color: #06B6D4;
      }
      
      &.experience {
        color: #F59E0B;
      }
    }
    
    .summary-label {
      font-size: 28rpx;
      color: #6b7280;
    }
  }
}
</style> 