<template>
  <view class="history-page">
    <!-- 导航栏 -->
    <NfNavBar title="打卡历史" :showBack="true" />
    
    <!-- 筛选区域 -->
    <view class="filter-section">
      <view class="month-selector">
        <button class="month-btn" @click="selectMonth('')" :class="{ active: selectedMonth === '' }">
          全部
        </button>
        <button 
          class="month-btn" 
          v-for="month in availableMonths" 
          :key="month"
          @click="selectMonth(month)"
          :class="{ active: selectedMonth === month }"
        >
          {{ formatMonth(month) }}
        </button>
      </view>
    </view>
    
    <!-- 统计信息 -->
    <view class="stats-overview">
      <NfCard type="basic">
        <template #content>
          <view class="stats-grid">
            <view class="stat-item">
              <text class="stat-number">{{ historyData.total }}</text>
              <text class="stat-label">总打卡数</text>
            </view>
            <view class="stat-item">
              <text class="stat-number">{{ currentMonthCount }}</text>
              <text class="stat-label">{{ selectedMonth ? '当月' : '本月' }}打卡</text>
            </view>
            <view class="stat-item">
              <text class="stat-number">{{ averageMood }}</text>
              <text class="stat-label">平均心情</text>
            </view>
          </view>
        </template>
      </NfCard>
    </view>
    
    <!-- 历史记录列表 -->
    <view class="history-list">
      <view v-if="loading" class="loading-state">
        <text>加载中...</text>
      </view>
      
      <view v-else-if="historyData.list.length === 0" class="empty-state">
        <i class="fas fa-calendar-times empty-icon"></i>
        <text class="empty-text">{{ selectedMonth ? '该月份暂无打卡记录' : '暂无打卡记录' }}</text>
        <text class="empty-hint">开始你的第一次打卡吧！</text>
      </view>
      
      <view v-else class="checkin-list">
        <view 
          v-for="checkin in historyData.list" 
          :key="checkin.id"
          class="checkin-item"
        >
          <view class="checkin-date">
            <text class="date-text">{{ formatDate(checkin.checkinDate) }}</text>
            <text class="weekday">{{ getWeekday(checkin.checkinDate) }}</text>
          </view>
          
          <view class="checkin-content">
            <view class="mood-info">
              <text class="mood-emoji">{{ getMoodEmoji(checkin.moodLevel) }}</text>
              <text class="mood-text">{{ getMoodLabel(checkin.moodLevel) }}</text>
              <view class="mood-stars">
                <text 
                  v-for="star in 5" 
                  :key="star"
                  class="mood-star"
                  :class="{ 'active': star <= checkin.moodLevel }"
                >★</text>
              </view>
            </view>
            
            <view v-if="checkin.notes" class="checkin-notes">
              <text class="notes-text">{{ checkin.notes }}</text>
            </view>
            
            <view class="checkin-meta">
              <text class="time-text">{{ formatTime(checkin.checkinDate) }}</text>
              <text class="rewards-text">+{{ checkin.rewards }} 经验</text>
            </view>
          </view>
        </view>
      </view>
      
      <!-- 加载更多 -->
      <view v-if="hasMore && !loading" class="load-more">
        <button class="load-more-btn" @click="loadMore">加载更多</button>
      </view>
      
      <view v-if="!hasMore && historyData.list.length > 0" class="no-more">
        <text>已显示全部记录</text>
      </view>
    </view>
  </view>
</template>

<script>
import { ref, reactive, computed, onMounted } from 'vue'

export default {
  name: 'CheckinHistory',
  setup() {
    // 响应式数据
    const loading = ref(false)
    const selectedMonth = ref('')
    const currentPage = ref(1)
    const pageSize = 20
    
    const historyData = reactive({
      list: [],
      total: 0,
      page: 1,
      pageSize: 20
    })
    
    const availableMonths = ref([])

    // 心情选项
    const moodOptions = [
      { level: 1, emoji: '😰', label: '很糟糕' },
      { level: 2, emoji: '😟', label: '不太好' },
      { level: 3, emoji: '😐', label: '一般' },
      { level: 4, emoji: '😊', label: '不错' },
      { level: 5, emoji: '😄', label: '很棒' }
    ]

    // 计算属性
    const hasMore = computed(() => {
      return historyData.list.length < historyData.total
    })

    const currentMonthCount = computed(() => {
      if (selectedMonth.value) {
        return historyData.list.length
      }
      const currentMonth = new Date().toISOString().slice(0, 7)
      return historyData.list.filter(item => 
        item.checkinDate.startsWith(currentMonth)
      ).length
    })

    const averageMood = computed(() => {
      if (historyData.list.length === 0) return '0.0'
      const sum = historyData.list.reduce((acc, item) => acc + item.moodLevel, 0)
      return (sum / historyData.list.length).toFixed(1)
    })

    // 方法
    const getMoodEmoji = (level) => {
      const mood = moodOptions.find(m => m.level === level)
      return mood ? mood.emoji : '😐'
    }

    const getMoodLabel = (level) => {
      const mood = moodOptions.find(m => m.level === level)
      return mood ? mood.label : '一般'
    }

    const formatDate = (dateString) => {
      const date = new Date(dateString)
      const today = new Date()
      const diffTime = today - date
      const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
      
      if (diffDays === 0) return '今天'
      if (diffDays === 1) return '昨天'
      if (diffDays === 2) return '前天'
      
      return `${date.getMonth() + 1}月${date.getDate()}日`
    }

    const getWeekday = (dateString) => {
      const date = new Date(dateString)
      const weekdays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
      return weekdays[date.getDay()]
    }

    const formatTime = (dateString) => {
      const date = new Date(dateString)
      return date.toLocaleTimeString('zh-CN', { 
        hour: '2-digit', 
        minute: '2-digit' 
      })
    }

    const formatMonth = (monthString) => {
      const [year, month] = monthString.split('-')
      return `${year}年${month}月`
    }

    const generateAvailableMonths = () => {
      const months = []
      const now = new Date()
      
      // 生成最近6个月
      for (let i = 0; i < 6; i++) {
        const date = new Date(now.getFullYear(), now.getMonth() - i, 1)
        const monthString = `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}`
        months.push(monthString)
      }
      
      availableMonths.value = months
    }

    const selectMonth = (month) => {
      selectedMonth.value = month
      currentPage.value = 1
      historyData.list = []
      loadHistoryData()
    }

    const loadHistoryData = async () => {
      if (loading.value) return
      
      loading.value = true
      
      try {
        const token = uni.getStorageSync('token')
        if (!token) {
          uni.redirectTo({ url: '/pages/auth/login' })
          return
        }

        const params = {
          page: currentPage.value,
          pageSize: pageSize
        }
        
        if (selectedMonth.value) {
          params.month = selectedMonth.value
        }

        const queryString = Object.keys(params)
          .map(key => `${key}=${encodeURIComponent(params[key])}`)
          .join('&')

        const res = await uni.request({
          url: `http://localhost:8888/api/v1/miniprogram/checkin/history?${queryString}`,
          method: 'GET',
          header: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        })

        if (res.data.code === 0) {
          const data = res.data.data
          
          if (currentPage.value === 1) {
            historyData.list = data.list || []
          } else {
            historyData.list.push(...(data.list || []))
          }
          
          historyData.total = data.total || 0
          historyData.page = data.page || 1
          historyData.pageSize = data.pageSize || pageSize
          
        } else {
          uni.showToast({
            title: res.data.msg || '获取数据失败',
            icon: 'none'
          })
        }
      } catch (error) {
        console.error('获取打卡历史失败:', error)
        uni.showToast({
          title: '网络错误，请重试',
          icon: 'none'
        })
      } finally {
        loading.value = false
      }
    }

    const loadMore = () => {
      if (!hasMore.value || loading.value) return
      currentPage.value++
      loadHistoryData()
    }

    // 生命周期
    onMounted(() => {
      generateAvailableMonths()
      loadHistoryData()
    })

    return {
      // 数据
      loading,
      selectedMonth,
      historyData,
      availableMonths,
      
      // 计算属性
      hasMore,
      currentMonthCount,
      averageMood,
      
      // 方法
      getMoodEmoji,
      getMoodLabel,
      formatDate,
      getWeekday,
      formatTime,
      formatMonth,
      selectMonth,
      loadMore
    }
  }
}
</script>

<style scoped>
.history-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding-bottom: var(--safe-area-inset-bottom);
}

.filter-section {
  padding: 20px;
  padding-bottom: 10px;
}

.month-selector {
  display: flex;
  gap: 8px;
  overflow-x: auto;
  padding-bottom: 4px;
}

.month-btn {
  white-space: nowrap;
  padding: 8px 16px;
  background: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 20px;
  font-size: 14px;
  color: #333;
  transition: all 0.3s ease;
}

.month-btn.active {
  background: var(--primary-color);
  border-color: var(--primary-color);
  color: white;
}

.month-btn:active {
  transform: scale(0.95);
}

.stats-overview {
  margin: 10px 20px 20px;
}

.stats-grid {
  display: flex;
  justify-content: space-around;
  padding: 16px 0;
}

.stat-item {
  text-align: center;
}

.stat-number {
  display: block;
  font-size: 20px;
  font-weight: bold;
  color: var(--primary-color);
  margin-bottom: 4px;
}

.stat-label {
  font-size: 12px;
  color: #666;
}

.history-list {
  padding: 0 20px 20px;
}

.loading-state, .empty-state, .no-more {
  text-align: center;
  padding: 40px 20px;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 12px;
  margin-bottom: 20px;
}

.empty-icon {
  font-size: 48px;
  color: #ccc;
  margin-bottom: 16px;
  display: block;
}

.empty-text {
  font-size: 16px;
  color: #666;
  margin-bottom: 8px;
  display: block;
}

.empty-hint {
  font-size: 14px;
  color: #999;
}

.checkin-list {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 12px;
  overflow: hidden;
}

.checkin-item {
  display: flex;
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
  transition: background-color 0.2s ease;
}

.checkin-item:last-child {
  border-bottom: none;
}

.checkin-item:active {
  background-color: #f5f5f5;
}

.checkin-date {
  min-width: 80px;
  margin-right: 16px;
  text-align: center;
}

.date-text {
  font-size: 14px;
  font-weight: bold;
  color: #333;
  display: block;
  margin-bottom: 2px;
}

.weekday {
  font-size: 12px;
  color: #666;
}

.checkin-content {
  flex: 1;
}

.mood-info {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
  gap: 8px;
}

.mood-emoji {
  font-size: 20px;
}

.mood-text {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.mood-stars {
  display: flex;
  gap: 2px;
  margin-left: auto;
}

.mood-star {
  font-size: 14px;
  color: #e0e0e0;
}

.mood-star.active {
  color: #ffd700;
}

.checkin-notes {
  margin-bottom: 8px;
  padding: 8px 12px;
  background: #f8f9fa;
  border-radius: 6px;
  border-left: 3px solid var(--primary-color);
}

.notes-text {
  font-size: 13px;
  color: #555;
  line-height: 1.4;
}

.checkin-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.time-text {
  font-size: 12px;
  color: #999;
}

.rewards-text {
  font-size: 12px;
  color: var(--primary-color);
  font-weight: bold;
}

.load-more {
  text-align: center;
  margin-top: 20px;
}

.load-more-btn {
  padding: 12px 24px;
  background: rgba(255, 255, 255, 0.9);
  border: 1px solid var(--primary-color);
  border-radius: 20px;
  color: var(--primary-color);
  font-size: 14px;
  transition: all 0.3s ease;
}

.load-more-btn:active {
  transform: scale(0.95);
  background: var(--primary-color);
  color: white;
}

.no-more {
  padding: 20px;
  font-size: 14px;
  color: #999;
}

/* 滚动条样式 */
.month-selector::-webkit-scrollbar {
  height: 4px;
}

.month-selector::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.3);
  border-radius: 2px;
}

.month-selector::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.6);
  border-radius: 2px;
}
</style> 