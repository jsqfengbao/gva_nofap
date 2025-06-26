<template>
  <view class="weekly-page">
    <!-- 状态栏 -->
    <view class="status-bar">
      <text class="status-time">{{ currentTime }}</text>
      <view class="status-icons">
        <text class="nf-icon fa-signal"></text>
        <text class="nf-icon fa-wifi"></text>
        <view class="battery">
          <view class="battery-fill"></view>
        </view>
      </view>
    </view>

    <!-- 导航头部 -->
    <NfNavBar 
      title="本周进度"
      :show-back="true"
      @back="goBack"
    />

    <!-- 主要内容 -->
    <view class="page-content">
      <!-- 本周概览卡片 -->
      <NfCard type="gradient" class="summary-card">
        <template #content>
          <view class="summary-content">
            <view class="week-range">
              <text class="range-text">{{ weekRange }}</text>
            </view>
            
            <view class="stats-grid">
              <view class="stat-item">
                <text class="stat-value">{{ weeklyData.summary.checkedDays }}</text>
                <text class="stat-label">已打卡</text>
              </view>
              <view class="stat-item">
                <text class="stat-value">{{ weeklyData.summary.successRate }}%</text>
                <text class="stat-label">成功率</text>
              </view>
              <view class="stat-item">
                <text class="stat-value">{{ weeklyData.summary.averageMood }}</text>
                <text class="stat-label">平均心情</text>
              </view>
            </view>
          </view>
        </template>
      </NfCard>

      <!-- 每日详情 -->
      <view class="daily-section">
        <view class="section-title">
          <text>每日详情</text>
        </view>
        
        <view class="weekdays-container">
          <view 
            v-for="(day, index) in weeklyData.weekDays" 
            :key="index"
            class="weekday-item"
            :class="{ 
              'checked': day.hasChecked, 
              'today': day.isToday,
              'unchecked': !day.hasChecked 
            }"
            @click="showDayDetail(day)"
          >
            <view class="day-header">
              <text class="weekday">{{ day.weekday }}</text>
              <text class="date">{{ formatDate(day.date) }}</text>
            </view>
            
            <view class="checkin-indicator">
              <view class="indicator-circle" :class="{ 'checked': day.hasChecked }">
                <i v-if="day.hasChecked" class="fas fa-check"></i>
                <i v-else class="fas fa-times"></i>
              </view>
            </view>
            
            <view class="mood-section" v-if="day.hasChecked">
              <text class="mood-emoji">{{ getMoodEmoji(day.moodLevel) }}</text>
              <text class="mood-label">{{ getMoodLabel(day.moodLevel) }}</text>
            </view>
            
            <view class="empty-mood" v-else>
              <text class="empty-text">未打卡</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 周统计图表 -->
      <view class="chart-section">
        <view class="section-title">
          <text>心情趋势</text>
        </view>
        
        <NfCard class="chart-card">
          <template #content>
            <view class="chart-container">
              <view class="chart-area">
                <view class="y-axis">
                  <text v-for="level in [5,4,3,2,1]" :key="level" class="y-label">{{ level }}</text>
                </view>
                <view class="chart-content">
                  <view class="chart-grid">
                    <view v-for="i in 5" :key="i" class="grid-line"></view>
                  </view>
                  <view class="chart-bars">
                    <view 
                      v-for="(day, index) in weeklyData.weekDays" 
                      :key="index"
                      class="bar-container"
                    >
                      <view 
                        class="mood-bar" 
                        :style="{ height: getMoodBarHeight(day.moodLevel) }"
                        :class="{ 'has-data': day.hasChecked }"
                      ></view>
                      <text class="bar-label">{{ getShortWeekday(day.weekday) }}</text>
                    </view>
                  </view>
                </view>
              </view>
            </view>
          </template>
        </NfCard>
      </view>

      <!-- 成就提示 -->
      <view class="achievement-section" v-if="weeklyData.summary.successRate >= 80">
        <NfCard type="success" class="achievement-card">
          <template #content>
            <view class="achievement-content">
              <view class="achievement-icon">
                <i class="fas fa-trophy"></i>
              </view>
              <view class="achievement-text">
                <text class="achievement-title">本周表现优秀！</text>
                <text class="achievement-desc">成功率达到{{ weeklyData.summary.successRate }}%，继续保持！</text>
              </view>
            </view>
          </template>
        </NfCard>
      </view>
    </view>

    <!-- 日详情弹窗 -->
    <view class="modal-overlay" v-if="showDetailModal" @click="closeDetailModal">
      <view class="modal-content" @click.stop>
        <view class="modal-header">
          <text class="modal-title">{{ selectedDay?.weekday }} 详情</text>
          <view class="modal-close" @click="closeDetailModal">
            <i class="fas fa-times"></i>
          </view>
        </view>
        
        <view class="modal-body">
          <view class="detail-date">
            <text>{{ selectedDay?.date }}</text>
          </view>
          
          <view v-if="selectedDay?.hasChecked" class="detail-content">
            <view class="detail-item">
              <text class="detail-label">心情指数</text>
              <view class="detail-value">
                <text class="mood-emoji">{{ getMoodEmoji(selectedDay.moodLevel) }}</text>
                <text class="mood-text">{{ getMoodLabel(selectedDay.moodLevel) }}</text>
              </view>
            </view>
            
            <view class="detail-item">
              <text class="detail-label">打卡状态</text>
              <text class="detail-value success">已完成</text>
            </view>
          </view>
          
          <view v-else class="detail-content">
            <view class="detail-item">
              <text class="detail-label">打卡状态</text>
              <text class="detail-value failed">未打卡</text>
            </view>
            
            <view class="detail-tip">
              <text>{{ selectedDay?.isToday ? '今天还可以去打卡哦！' : '这天没有打卡记录' }}</text>
            </view>
          </view>
        </view>
        
        <view class="modal-footer" v-if="selectedDay?.isToday && !selectedDay?.hasChecked">
          <NfButton type="primary" size="medium" @click="goToCheckin">
            去打卡
          </NfButton>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import NfNavBar from '@/components/ui/navigation/NfNavBar.vue'
import NfCard from '@/components/ui/card/NfCard.vue'
import NfButton from '@/components/ui/button/NfButton.vue'

export default {
  name: 'WeeklyProgress',
  components: {
    NfNavBar,
    NfCard,
    NfButton
  },
  setup() {
    // 响应式数据
    const weeklyData = ref({
      weekDays: [],
      summary: {
        checkedDays: 0,
        totalDays: 7,
        successRate: 0,
        averageMood: 0
      }
    })
    
    const loading = ref(false)
    const showDetailModal = ref(false)
    const selectedDay = ref(null)

    // 计算属性
    const currentTime = computed(() => {
      return new Date().toLocaleTimeString('zh-CN', { 
        hour: '2-digit', 
        minute: '2-digit' 
      })
    })

    const weekRange = computed(() => {
      if (weeklyData.value.weekDays.length === 0) return ''
      
      const firstDay = weeklyData.value.weekDays[0]?.date
      const lastDay = weeklyData.value.weekDays[6]?.date
      
      if (!firstDay || !lastDay) return ''
      
      const startDate = new Date(firstDay)
      const endDate = new Date(lastDay)
      
      return `${formatDateRange(startDate)} - ${formatDateRange(endDate)}`
    })

    // 方法
    const formatDate = (dateStr) => {
      const date = new Date(dateStr)
      return `${date.getMonth() + 1}/${date.getDate()}`
    }

    const formatDateRange = (date) => {
      return `${date.getMonth() + 1}月${date.getDate()}日`
    }

    const getShortWeekday = (weekday) => {
      const map = {
        '周一': '一',
        '周二': '二',
        '周三': '三',
        '周四': '四',
        '周五': '五',
        '周六': '六',
        '周日': '日'
      }
      return map[weekday] || weekday
    }

    const getMoodEmoji = (level) => {
      const emojis = ['', '😢', '😕', '😐', '😊', '😄']
      return emojis[level] || '😐'
    }

    const getMoodLabel = (level) => {
      const labels = ['', '很糟糕', '不太好', '一般', '不错', '很开心']
      return labels[level] || '一般'
    }

    const getMoodBarHeight = (moodLevel) => {
      if (!moodLevel) return '0%'
      return `${(moodLevel / 5) * 100}%`
    }

    const showDayDetail = (day) => {
      selectedDay.value = day
      showDetailModal.value = true
    }

    const closeDetailModal = () => {
      showDetailModal.value = false
      selectedDay.value = null
    }

    const goBack = () => {
      uni.navigateBack()
    }

    const goToCheckin = () => {
      closeDetailModal()
      uni.navigateTo({ url: '/pages/checkin/index' })
    }

    const loadWeeklyData = async () => {
      if (loading.value) return
      
      loading.value = true
      
      try {
        const token = uni.getStorageSync('token')
        if (!token) {
          uni.redirectTo({ url: '/pages/auth/login' })
          return
        }

        const res = await uni.request({
          url: 'http://localhost:8888/api/v1/miniprogram/checkin/weekly',
          method: 'GET',
          header: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        })

        if (res.data.code === 0) {
          weeklyData.value = res.data.data
        } else {
          uni.showToast({
            title: res.data.msg || '获取数据失败',
            icon: 'none'
          })
        }
      } catch (error) {
        console.error('获取本周进度失败:', error)
        uni.showToast({
          title: '网络错误，请重试',
          icon: 'none'
        })
      } finally {
        loading.value = false
      }
    }

    // 生命周期
    onMounted(() => {
      loadWeeklyData()
    })

    return {
      // 数据
      weeklyData,
      loading,
      showDetailModal,
      selectedDay,
      
      // 计算属性
      currentTime,
      weekRange,
      
      // 方法
      formatDate,
      getShortWeekday,
      getMoodEmoji,
      getMoodLabel,
      getMoodBarHeight,
      showDayDetail,
      closeDetailModal,
      goBack,
      goToCheckin
    }
  }
}
</script>

<style scoped>
/* 页面容器 */
.weekly-page {
  height: 100vh;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  display: flex;
  flex-direction: column;
}

/* 状态栏 */
.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
  font-size: 12px;
  color: #1f2937;
  background: transparent;
}

.status-time {
  font-weight: 500;
}

.status-icons {
  display: flex;
  align-items: center;
  gap: 4px;
}

.battery {
  width: 24px;
  height: 12px;
  border: 1px solid #1f2937;
  border-radius: 2px;
  position: relative;
}

.battery-fill {
  width: 16px;
  height: 6px;
  background-color: #10b981;
  border-radius: 1px;
  margin: 2px;
}

/* 主要内容 */
.page-content {
  flex: 1;
  padding: 80px 16px 32px;
  overflow-y: auto;
}

/* 概览卡片 */
.summary-card {
  margin-bottom: 24px;
}

.summary-content {
  text-align: center;
  padding: 20px 0;
}

.week-range {
  margin-bottom: 20px;
}

.range-text {
  font-size: 16px;
  color: #333;
  font-weight: 500;
}

.stats-grid {
  display: flex;
  justify-content: space-around;
  gap: 20px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #2563eb;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 12px;
  color: #666;
}

/* 每日详情 */
.daily-section {
  margin-bottom: 24px;
}

.section-title {
  margin-bottom: 16px;
}

.section-title text {
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
}

.weekdays-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.weekday-item {
  background: white;
  border-radius: 12px;
  padding: 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  transition: all 0.3s ease;
}

.weekday-item.today {
  border: 2px solid #2563eb;
  background: #eff6ff;
}

.weekday-item.checked {
  background: #f0fdf4;
  border: 1px solid #22c55e;
}

.weekday-item.unchecked {
  background: #fef2f2;
  border: 1px solid #ef4444;
}

.day-header {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.weekday {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.date {
  font-size: 12px;
  color: #666;
}

.checkin-indicator {
  margin: 0 16px;
}

.indicator-circle {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f3f4f6;
  color: #9ca3af;
}

.indicator-circle.checked {
  background: #22c55e;
  color: white;
}

.mood-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.mood-emoji {
  font-size: 20px;
}

.mood-label {
  font-size: 12px;
  color: #666;
}

.empty-mood {
  display: flex;
  align-items: center;
}

.empty-text {
  font-size: 12px;
  color: #9ca3af;
}

/* 图表区域 */
.chart-section {
  margin-bottom: 24px;
}

.chart-card {
  padding: 0;
}

.chart-container {
  padding: 20px;
}

.chart-area {
  display: flex;
  height: 200px;
}

.y-axis {
  display: flex;
  flex-direction: column-reverse;
  justify-content: space-between;
  width: 20px;
  margin-right: 10px;
}

.y-label {
  font-size: 12px;
  color: #666;
  text-align: center;
}

.chart-content {
  flex: 1;
  position: relative;
}

.chart-grid {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.grid-line {
  height: 1px;
  background: #e5e7eb;
}

.chart-bars {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  height: 100%;
  padding-top: 10px;
}

.bar-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  flex: 1;
  margin: 0 2px;
}

.mood-bar {
  width: 20px;
  background: #e5e7eb;
  border-radius: 2px;
  transition: all 0.3s ease;
}

.mood-bar.has-data {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%);
}

.bar-label {
  font-size: 12px;
  color: #666;
}

/* 成就区域 */
.achievement-section {
  margin-bottom: 24px;
}

.achievement-content {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 20px 0;
}

.achievement-icon {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: #fbbf24;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: white;
}

.achievement-text {
  flex: 1;
}

.achievement-title {
  display: block;
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 4px;
}

.achievement-desc {
  font-size: 14px;
  color: #666;
}

/* 弹窗样式 */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 16px;
  width: 90%;
  max-width: 400px;
  max-height: 80vh;
  overflow: hidden;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  border-bottom: 1px solid #e5e7eb;
}

.modal-title {
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
}

.modal-close {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
}

.modal-body {
  padding: 20px;
}

.detail-date {
  text-align: center;
  margin-bottom: 20px;
  font-size: 16px;
  color: #666;
}

.detail-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.detail-label {
  font-size: 14px;
  color: #666;
}

.detail-value {
  font-size: 14px;
  color: #1f2937;
  display: flex;
  align-items: center;
  gap: 8px;
}

.detail-value.success {
  color: #22c55e;
  font-weight: 600;
}

.detail-value.failed {
  color: #ef4444;
  font-weight: 600;
}

.detail-tip {
  text-align: center;
  padding: 16px;
  background: #f8fafc;
  border-radius: 8px;
  font-size: 14px;
  color: #666;
}

.modal-footer {
  padding: 20px;
  border-top: 1px solid #e5e7eb;
  text-align: center;
}
</style> 