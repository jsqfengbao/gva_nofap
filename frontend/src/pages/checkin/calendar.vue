<template>
  <view class="calendar-page">
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
      title="月度日历"
      :show-back="true"
      @back="goBack"
    />

    <!-- 主要内容 -->
    <view class="page-content">
      <!-- 月份选择器 -->
      <view class="month-selector">
        <view class="month-nav" @click="previousMonth">
          <i class="fas fa-chevron-left"></i>
        </view>
        <view class="month-display">
          <text class="year">{{ calendarData.year }}年</text>
          <text class="month">{{ calendarData.month }}月</text>
        </view>
        <view class="month-nav" @click="nextMonth">
          <i class="fas fa-chevron-right"></i>
        </view>
      </view>

      <!-- 月度统计卡片 -->
      <NfCard type="gradient" class="summary-card">
        <template #content>
          <view class="summary-content">
            <view class="summary-title">
              <text>本月概览</text>
            </view>
            
            <view class="stats-grid">
              <view class="stat-item">
                <text class="stat-value">{{ calendarData.summary.checkedDays }}</text>
                <text class="stat-label">已打卡</text>
              </view>
              <view class="stat-item">
                <text class="stat-value">{{ calendarData.summary.successRate }}%</text>
                <text class="stat-label">成功率</text>
              </view>
              <view class="stat-item">
                <text class="stat-value">{{ calendarData.summary.averageMood }}</text>
                <text class="stat-label">平均心情</text>
              </view>
              <view class="stat-item">
                <text class="stat-value">{{ calendarData.summary.bestStreak }}</text>
                <text class="stat-label">最长连击</text>
              </view>
            </view>
          </view>
        </template>
      </NfCard>

      <!-- 日历区域 -->
      <view class="calendar-section">
        <NfCard class="calendar-card">
          <template #content>
            <!-- 星期标题 -->
            <view class="weekdays-header">
              <text v-for="weekday in weekdays" :key="weekday" class="weekday-label">{{ weekday }}</text>
            </view>
            
            <!-- 日历网格 -->
            <view class="calendar-grid">
              <view 
                v-for="(day, index) in calendarDays" 
                :key="index"
                class="calendar-day"
                :class="{
                  'today': day.isToday,
                  'checked': day.hasChecked,
                  'other-month': !day.isThisMonth
                }"
                @click="showDayDetail(day)"
              >
                <text class="day-number">{{ day.day }}</text>
                <view v-if="day.hasChecked && day.isThisMonth" class="day-indicator">
                  <view class="mood-dot" :class="'mood-' + day.moodLevel"></view>
                </view>
                <view v-else-if="day.isThisMonth" class="day-indicator">
                  <view class="empty-dot"></view>
                </view>
              </view>
            </view>
          </template>
        </NfCard>
      </view>
    </view>
  </view>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import NfNavBar from '@/components/ui/navigation/NfNavBar.vue'
import NfCard from '@/components/ui/card/NfCard.vue'

export default {
  name: 'CalendarPage',
  components: {
    NfNavBar,
    NfCard
  },
  setup() {
    // 响应式数据
    const calendarData = ref({
      year: new Date().getFullYear(),
      month: new Date().getMonth() + 1,
      days: [],
      summary: {
        checkedDays: 0,
        totalDays: 0,
        successRate: 0,
        averageMood: 0,
        bestStreak: 0
      }
    })
    
    const loading = ref(false)

    // 静态数据
    const weekdays = ['日', '一', '二', '三', '四', '五', '六']

    // 计算属性
    const currentTime = computed(() => {
      return new Date().toLocaleTimeString('zh-CN', { 
        hour: '2-digit', 
        minute: '2-digit' 
      })
    })

    const calendarDays = computed(() => {
      const year = calendarData.value.year
      const month = calendarData.value.month
      
      // 获取本月第一天和最后一天
      const firstDay = new Date(year, month - 1, 1)
      const lastDay = new Date(year, month, 0)
      
      // 获取本月第一天是星期几
      const startWeekday = firstDay.getDay()
      
      // 获取上个月的最后几天
      const prevMonth = month === 1 ? 12 : month - 1
      const prevYear = month === 1 ? year - 1 : year
      const prevMonthLastDay = new Date(prevYear, prevMonth, 0).getDate()
      
      // 获取下个月的前几天
      const nextMonth = month === 12 ? 1 : month + 1
      const nextYear = month === 12 ? year + 1 : year
      
      const days = []
      
      // 上个月的日期
      for (let i = startWeekday - 1; i >= 0; i--) {
        const day = prevMonthLastDay - i
        days.push({
          day,
          date: `${prevYear}-${String(prevMonth).padStart(2, '0')}-${String(day).padStart(2, '0')}`,
          hasChecked: false,
          moodLevel: 0,
          isToday: false,
          isThisMonth: false
        })
      }
      
      // 本月的日期
      const today = new Date().toISOString().split('T')[0]
      for (let day = 1; day <= lastDay.getDate(); day++) {
        const dateStr = `${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`
        const dayData = calendarData.value.days.find(d => d.date === dateStr) || {}
        
        days.push({
          day,
          date: dateStr,
          hasChecked: dayData.hasChecked || false,
          moodLevel: dayData.moodLevel || 0,
          isToday: dateStr === today,
          isThisMonth: true
        })
      }
      
      // 下个月的日期（填满6行）
      const remainingDays = 42 - days.length
      for (let day = 1; day <= remainingDays; day++) {
        days.push({
          day,
          date: `${nextYear}-${String(nextMonth).padStart(2, '0')}-${String(day).padStart(2, '0')}`,
          hasChecked: false,
          moodLevel: 0,
          isToday: false,
          isThisMonth: false
        })
      }
      
      return days
    })

    // 方法
    const showDayDetail = (day) => {
      if (!day.isThisMonth) return
      console.log('Day detail:', day)
    }

    const previousMonth = () => {
      if (calendarData.value.month === 1) {
        calendarData.value.year--
        calendarData.value.month = 12
      } else {
        calendarData.value.month--
      }
      loadCalendarData()
    }

    const nextMonth = () => {
      if (calendarData.value.month === 12) {
        calendarData.value.year++
        calendarData.value.month = 1
      } else {
        calendarData.value.month++
      }
      loadCalendarData()
    }

    const goBack = () => {
      uni.navigateBack()
    }

    const loadCalendarData = async () => {
      if (loading.value) return
      
      loading.value = true
      
      try {
        const token = uni.getStorageSync('token')
        if (!token) {
          uni.redirectTo({ url: '/pages/auth/login' })
          return
        }

        const res = await uni.request({
          url: `http://localhost:8888/api/v1/miniprogram/checkin/calendar?year=${calendarData.value.year}&month=${calendarData.value.month}`,
          method: 'GET',
          header: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        })

        if (res.data.code === 0) {
          Object.assign(calendarData.value, res.data.data)
        } else {
          uni.showToast({
            title: res.data.msg || '获取数据失败',
            icon: 'none'
          })
        }
      } catch (error) {
        console.error('获取日历数据失败:', error)
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
      loadCalendarData()
    })

    return {
      // 数据
      calendarData,
      loading,
      weekdays,
      
      // 计算属性
      currentTime,
      calendarDays,
      
      // 方法
      showDayDetail,
      previousMonth,
      nextMonth,
      goBack
    }
  }
}
</script>

<style scoped>
/* 基础样式省略... */
.calendar-page {
  height: 100vh;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  display: flex;
  flex-direction: column;
}

.page-content {
  flex: 1;
  padding: 80px 16px 32px;
  overflow-y: auto;
}

.calendar-grid {
  display: flex;
  flex-wrap: wrap;
  padding: 8px;
}

.calendar-day {
  width: calc(100% / 7);
  aspect-ratio: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
  border-radius: 8px;
  margin: 2px 0;
}

.mood-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.mood-dot.mood-5 { background: #22c55e; }
.mood-dot.mood-4 { background: #84cc16; }
.mood-dot.mood-3 { background: #eab308; }
.mood-dot.mood-2 { background: #f97316; }
.mood-dot.mood-1 { background: #ef4444; }
</style>
