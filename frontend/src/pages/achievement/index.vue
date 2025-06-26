<template>
  <view class="achievement-page">
    <!-- 顶部导航 -->
    <NfNavBar title="成就系统" show-back />

    <!-- 游戏化统计概览 -->
    <view class="stats-overview">
      <view class="level-card">
        <view class="level-info">
          <text class="level-title">{{ gamificationService.GetLevelTitle(gameStats.currentLevel) }}</text>
          <text class="level-number">Lv.{{ gameStats.currentLevel }}</text>
        </view>
        <view class="exp-info">
          <view class="exp-bar">
            <view class="exp-progress" :style="{ width: gameStats.levelProgress + '%' }"></view>
          </view>
          <text class="exp-text">{{ gameStats.currentExp }} / {{ gameStats.nextLevelExp }}</text>
        </view>
      </view>

      <view class="achievement-overview">
        <view class="overview-item">
          <text class="overview-number">{{ gameStats.unlockedAchievements }}</text>
          <text class="overview-label">已解锁</text>
        </view>
        <view class="overview-item">
          <text class="overview-number">{{ gameStats.totalAchievements }}</text>
          <text class="overview-label">总成就</text>
        </view>
        <view class="overview-item">
          <text class="overview-number">{{ Math.round(gameStats.achievementRate) }}%</text>
          <text class="overview-label">完成率</text>
        </view>
      </view>
    </view>

    <!-- 最近解锁的成就 -->
    <view class="recent-section" v-if="gameStats.recentAchievements && gameStats.recentAchievements.length > 0">
      <view class="section-header">
        <text class="section-title">最近解锁</text>
      </view>
      <scroll-view class="recent-achievements" scroll-x>
        <view class="recent-item" v-for="achievement in gameStats.recentAchievements" :key="achievement.id">
          <view class="achievement-icon">
            <text class="icon-emoji">🏆</text>
            <view class="rarity-badge" :class="getRarityClass(achievement.rarity)">
              {{ getRarityText(achievement.rarity) }}
            </view>
          </view>
          <text class="achievement-name">{{ achievement.name }}</text>
          <text class="unlock-date">{{ formatDate(achievement.unlockedAt) }}</text>
        </view>
      </scroll-view>
    </view>

    <!-- 成就分类 -->
    <view class="achievement-categories">
      <view class="category-tabs">
        <view 
          class="tab-item" 
          :class="{ active: activeCategory === category }"
          v-for="(categoryName, category) in categoryNames" 
          :key="category"
          @click="switchCategory(category)"
        >
          <text>{{ categoryName }}</text>
        </view>
      </view>

      <!-- 成就列表 -->
      <view class="achievement-list" v-if="achievements.categories[activeCategory]">
        <view 
          class="achievement-item" 
          :class="{ unlocked: achievement.isUnlocked }"
          v-for="achievement in achievements.categories[activeCategory]" 
          :key="achievement.id"
          @click="showAchievementDetail(achievement)"
        >
          <view class="achievement-left">
            <view class="achievement-icon-wrapper">
              <text class="achievement-icon" v-if="achievement.isUnlocked">🏆</text>
              <text class="achievement-icon locked" v-else>🔒</text>
              <view class="rarity-indicator" :class="getRarityClass(achievement.rarity)"></view>
            </view>
            <view class="achievement-content">
              <text class="achievement-title">{{ achievement.name }}</text>
              <text class="achievement-desc">{{ achievement.description }}</text>
              <view class="achievement-meta">
                <text class="rewards">奖励: +{{ achievement.rewards }}经验</text>
                <text class="rarity">{{ getRarityText(achievement.rarity) }}</text>
              </view>
            </view>
          </view>
          <view class="achievement-right">
            <text class="unlock-status" v-if="achievement.isUnlocked">已解锁</text>
            <text class="unlock-status locked" v-else>未解锁</text>
          </view>
        </view>
      </view>
    </view>

    <!-- 成就详情弹窗 -->
    <view class="achievement-modal" v-show="showModal" @click="closeModal">
      <view class="modal-content" @click.stop>
        <view class="modal-header">
          <view class="modal-icon">
            <text v-if="selectedAchievement?.isUnlocked">🏆</text>
            <text v-else>🔒</text>
          </view>
          <view class="modal-title">
            <text class="title">{{ selectedAchievement?.name }}</text>
            <text class="rarity" :class="getRarityClass(selectedAchievement?.rarity)">
              {{ getRarityText(selectedAchievement?.rarity) }}
            </text>
          </view>
        </view>
        
        <view class="modal-body">
          <text class="description">{{ selectedAchievement?.description }}</text>
          
          <view class="achievement-details">
            <view class="detail-item">
              <text class="label">分类:</text>
              <text class="value">{{ getCategoryText(selectedAchievement?.category) }}</text>
            </view>
            <view class="detail-item">
              <text class="label">奖励:</text>
              <text class="value">+{{ selectedAchievement?.rewards }}经验值</text>
            </view>
            <view class="detail-item" v-if="selectedAchievement?.isUnlocked">
              <text class="label">解锁时间:</text>
              <text class="value">{{ formatDateTime(selectedAchievement?.unlockedAt) }}</text>
            </view>
          </view>
        </view>
        
        <view class="modal-footer">
          <button class="close-btn" @click="closeModal">关闭</button>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'

// 响应式数据
const loading = ref(false)
const activeCategory = ref(1) // 默认打卡类
const showModal = ref(false)
const selectedAchievement = ref(null)

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

const achievements = reactive({
  totalAchievements: 0,
  unlockedAchievements: 0,
  unlockRate: 0,
  categories: {
    1: [], // 打卡类
    2: [], // 等级类
    3: [], // 社区类
    4: [], // 学习类
    5: []  // 特殊类
  },
  recentUnlocked: []
})

// 常量数据
const categoryNames = {
  1: '打卡类',
  2: '等级类',
  3: '社区类',
  4: '学习类',
  5: '特殊类'
}

const rarityNames = {
  1: '普通',
  2: '稀有', 
  3: '史诗',
  4: '传说'
}

// 计算属性
const gamificationService = computed(() => {
  return {
    GetLevelTitle: (level) => {
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
    }
  }
})

// 方法
const switchCategory = (category) => {
  activeCategory.value = category
}

const getRarityClass = (rarity) => {
  const classes = {
    1: 'common',
    2: 'rare',
    3: 'epic',
    4: 'legendary'
  }
  return classes[rarity] || 'common'
}

const getRarityText = (rarity) => {
  return rarityNames[rarity] || '普通'
}

const getCategoryText = (category) => {
  return categoryNames[category] || '未知'
}

const formatDate = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}月${date.getDate()}日`
}

const formatDateTime = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')} ${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
}

const showAchievementDetail = (achievement) => {
  selectedAchievement.value = achievement
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  selectedAchievement.value = null
}

const loadGameStats = async () => {
  try {
    const token = uni.getStorageSync('token')
    if (!token) {
      uni.redirectTo({ url: '/pages/auth/login' })
      return
    }

    const res = await uni.request({
      url: 'http://localhost:8888/api/v1/miniprogram/achievement/game-stats',
      method: 'GET',
      header: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })

    if (res.data.code === 0) {
      Object.assign(gameStats, res.data.data)
    }
  } catch (error) {
    console.error('获取游戏统计失败:', error)
  }
}

const loadAchievements = async () => {
  try {
    const token = uni.getStorageSync('token')
    if (!token) {
      uni.redirectTo({ url: '/pages/auth/login' })
      return
    }

    const res = await uni.request({
      url: 'http://localhost:8888/api/v1/miniprogram/achievement/list',
      method: 'GET',
      header: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })

    if (res.data.code === 0) {
      Object.assign(achievements, res.data.data)
    }
  } catch (error) {
    console.error('获取成就列表失败:', error)
  }
}

const loadData = async () => {
  loading.value = true
  try {
    await Promise.all([
      loadGameStats(),
      loadAchievements()
    ])
  } finally {
    loading.value = false
  }
}

// 生命周期
onMounted(() => {
  loadData()
})
</script>

<style lang="scss" scoped>
.achievement-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding-bottom: 40rpx;
}

.stats-overview {
  padding: 40rpx 30rpx;
  
  .level-card {
    background: rgba(255, 255, 255, 0.95);
    border-radius: 20rpx;
    padding: 30rpx;
    margin-bottom: 30rpx;
    box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.1);
    
    .level-info {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20rpx;
      
      .level-title {
        font-size: 32rpx;
        font-weight: bold;
        color: #333;
      }
      
      .level-number {
        font-size: 28rpx;
        color: #667eea;
        font-weight: bold;
      }
    }
    
    .exp-info {
      .exp-bar {
        height: 12rpx;
        background: #f0f0f0;
        border-radius: 6rpx;
        overflow: hidden;
        margin-bottom: 10rpx;
        
        .exp-progress {
          height: 100%;
          background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
          transition: width 0.3s ease;
        }
      }
      
      .exp-text {
        font-size: 24rpx;
        color: #666;
      }
    }
  }
  
  .achievement-overview {
    display: flex;
    justify-content: space-around;
    background: rgba(255, 255, 255, 0.95);
    border-radius: 20rpx;
    padding: 30rpx;
    box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.1);
    
    .overview-item {
      text-align: center;
      
      .overview-number {
        display: block;
        font-size: 40rpx;
        font-weight: bold;
        color: #667eea;
        margin-bottom: 10rpx;
      }
      
      .overview-label {
        font-size: 24rpx;
        color: #666;
      }
    }
  }
}

.recent-section {
  padding: 0 30rpx 30rpx;
  
  .section-header {
    margin-bottom: 20rpx;
    
    .section-title {
      font-size: 32rpx;
      font-weight: bold;
      color: white;
    }
  }
  
  .recent-achievements {
    white-space: nowrap;
    
    .recent-item {
      display: inline-block;
      width: 160rpx;
      margin-right: 20rpx;
      background: rgba(255, 255, 255, 0.95);
      border-radius: 15rpx;
      padding: 20rpx;
      text-align: center;
      vertical-align: top;
      white-space: normal;
      
      .achievement-icon {
        position: relative;
        margin-bottom: 15rpx;
        
        .icon-emoji {
          font-size: 40rpx;
        }
        
        .rarity-badge {
          position: absolute;
          top: -5rpx;
          right: -5rpx;
          font-size: 16rpx;
          padding: 2rpx 8rpx;
          border-radius: 8rpx;
          color: white;
          
          &.common { background: #9ca3af; }
          &.rare { background: #3b82f6; }
          &.epic { background: #8b5cf6; }
          &.legendary { background: #f59e0b; }
        }
      }
      
      .achievement-name {
        display: block;
        font-size: 24rpx;
        font-weight: bold;
        color: #333;
        margin-bottom: 8rpx;
      }
      
      .unlock-date {
        font-size: 20rpx;
        color: #666;
      }
    }
  }
}

.achievement-categories {
  background: white;
  border-radius: 30rpx 30rpx 0 0;
  min-height: 60vh;
  
  .category-tabs {
    display: flex;
    padding: 30rpx 30rpx 0;
    border-bottom: 2rpx solid #f0f0f0;
    
    .tab-item {
      flex: 1;
      text-align: center;
      padding: 20rpx 0;
      font-size: 28rpx;
      color: #666;
      border-bottom: 4rpx solid transparent;
      transition: all 0.3s ease;
      
      &.active {
        color: #667eea;
        border-bottom-color: #667eea;
        font-weight: bold;
      }
    }
  }
  
  .achievement-list {
    padding: 30rpx;
    
    .achievement-item {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 30rpx;
      margin-bottom: 20rpx;
      background: #f8f9fa;
      border-radius: 15rpx;
      border: 2rpx solid transparent;
      transition: all 0.3s ease;
      
      &.unlocked {
        background: #f0f9ff;
        border-color: #e0f2fe;
      }
      
      .achievement-left {
        display: flex;
        align-items: center;
        flex: 1;
        
        .achievement-icon-wrapper {
          position: relative;
          margin-right: 20rpx;
          
          .achievement-icon {
            font-size: 48rpx;
            
            &.locked {
              opacity: 0.3;
            }
          }
          
          .rarity-indicator {
            position: absolute;
            bottom: -5rpx;
            right: -5rpx;
            width: 16rpx;
            height: 16rpx;
            border-radius: 50%;
            
            &.common { background: #9ca3af; }
            &.rare { background: #3b82f6; }
            &.epic { background: #8b5cf6; }
            &.legendary { background: #f59e0b; }
          }
        }
        
        .achievement-content {
          flex: 1;
          
          .achievement-title {
            display: block;
            font-size: 28rpx;
            font-weight: bold;
            color: #333;
            margin-bottom: 8rpx;
          }
          
          .achievement-desc {
            display: block;
            font-size: 24rpx;
            color: #666;
            margin-bottom: 8rpx;
          }
          
          .achievement-meta {
            display: flex;
            align-items: center;
            gap: 20rpx;
            
            .rewards {
              font-size: 22rpx;
              color: #10b981;
              font-weight: bold;
            }
            
            .rarity {
              font-size: 22rpx;
              color: #8b5cf6;
            }
          }
        }
      }
      
      .achievement-right {
        .unlock-status {
          font-size: 24rpx;
          font-weight: bold;
          color: #10b981;
          
          &.locked {
            color: #9ca3af;
          }
        }
      }
    }
  }
}

.achievement-modal {
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
  
  .modal-content {
    background: white;
    border-radius: 20rpx;
    max-width: 600rpx;
    width: 90%;
    max-height: 80%;
    overflow: hidden;
    
    .modal-header {
      padding: 40rpx;
      text-align: center;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      
      .modal-icon {
        font-size: 80rpx;
        margin-bottom: 20rpx;
      }
      
      .modal-title {
        .title {
          display: block;
          font-size: 32rpx;
          font-weight: bold;
          color: white;
          margin-bottom: 10rpx;
        }
        
        .rarity {
          font-size: 24rpx;
          padding: 8rpx 16rpx;
          border-radius: 12rpx;
          color: white;
          
          &.common { background: rgba(156, 163, 175, 0.5); }
          &.rare { background: rgba(59, 130, 246, 0.5); }
          &.epic { background: rgba(139, 92, 246, 0.5); }
          &.legendary { background: rgba(245, 158, 11, 0.5); }
        }
      }
    }
    
    .modal-body {
      padding: 40rpx;
      
      .description {
        display: block;
        font-size: 28rpx;
        color: #333;
        line-height: 1.6;
        margin-bottom: 30rpx;
      }
      
      .achievement-details {
        .detail-item {
          display: flex;
          justify-content: space-between;
          margin-bottom: 20rpx;
          
          .label {
            font-size: 26rpx;
            color: #666;
          }
          
          .value {
            font-size: 26rpx;
            color: #333;
            font-weight: bold;
          }
        }
      }
    }
    
    .modal-footer {
      padding: 30rpx 40rpx;
      border-top: 2rpx solid #f0f0f0;
      
      .close-btn {
        width: 100%;
        height: 80rpx;
        background: #667eea;
        color: white;
        border: none;
        border-radius: 10rpx;
        font-size: 28rpx;
        font-weight: bold;
      }
    }
  }
}
</style> 