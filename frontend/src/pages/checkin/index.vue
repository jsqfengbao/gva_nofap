<template>
  <view class="checkin-page">
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
      title="每日打卡"
      :show-back="false"
      right-text="历史"
      @right-click="viewHistory"
    />

    <!-- 主要内容 -->
    <view class="page-content">
      <!-- 今日状态卡片 -->
      <NfCard type="gradient" class="status-card">
        <template #content>
          <view class="status-content">
            <view class="date-info">
              <text class="date">{{ currentDate }}</text>
              <text class="weekday">{{ currentWeekday }}</text>
            </view>
            
            <view class="checkin-status" v-if="!hasCheckedToday">
              <view class="status-icon">
                <i class="fas fa-calendar-check"></i>
              </view>
              <text class="status-text">今日未打卡</text>
              <text class="status-desc">坚持打卡，记录你的成长</text>
            </view>
            
            <view class="checkin-status checked" v-else>
              <view class="status-icon checked">
                <i class="fas fa-check-circle"></i>
              </view>
              <text class="status-text">今日已打卡</text>
              <text class="status-desc">{{ todayCheckin.checkinTime }}</text>
            </view>
          </view>
        </template>
      </NfCard>

      <!-- 连续天数统计 -->
      <view class="stats-section">
        <view class="stats-grid">
          <view class="stat-item">
            <text class="stat-number">{{ statistics.currentStreak }}</text>
            <text class="stat-label">连续天数</text>
          </view>
          <view class="stat-item">
            <text class="stat-number">{{ statistics.totalDays }}</text>
            <text class="stat-label">累计天数</text>
          </view>
          <view class="stat-item">
            <text class="stat-number">{{ statistics.level }}</text>
            <text class="stat-label">当前等级</text>
          </view>
        </view>
      </view>

      <!-- 心情选择 -->
      <view class="mood-section" v-if="!hasCheckedToday">
        <view class="section-title">
          <text>今日心情如何？</text>
        </view>
        <view class="mood-selector">
          <view 
            class="mood-item"
            :class="{ active: selectedMood === mood.level }"
            v-for="mood in moodOptions"
            :key="mood.level"
            @click="selectMood(mood.level)"
          >
            <text class="mood-emoji">{{ mood.emoji }}</text>
            <text class="mood-label">{{ mood.label }}</text>
          </view>
        </view>
      </view>

      <!-- 打卡备注 -->
      <view class="notes-section" v-if="!hasCheckedToday">
        <view class="section-title">
          <text>今日感想（可选）</text>
        </view>
        <textarea 
          v-model="checkinNotes"
          class="notes-input"
          placeholder="记录今天的想法、感受或收获..."
          maxlength="200"
        />
        <text class="notes-count">{{ checkinNotes.length }}/200</text>
      </view>

      <!-- 打卡按钮 -->
      <view class="action-section">
        <NfButton 
          v-if="!hasCheckedToday"
          type="primary" 
          size="large" 
          full-width
          :disabled="!selectedMood || isSubmitting"
          @click="submitCheckin"
        >
          <text v-if="!isSubmitting">{{ selectedMood ? '立即打卡' : '请先选择心情' }}</text>
          <text v-else>提交中...</text>
        </NfButton>
        
        <view v-else class="checked-section">
          <view class="checked-info">
            <view class="checked-mood">
              <text class="mood-emoji">{{ getMoodEmoji(todayCheckin.moodLevel) }}</text>
              <text class="mood-text">心情：{{ getMoodLabel(todayCheckin.moodLevel) }}</text>
            </view>
            <view class="checked-notes" v-if="todayCheckin.notes">
              <text class="notes-label">今日感想：</text>
              <text class="notes-text">{{ todayCheckin.notes }}</text>
            </view>
            <view class="rewards-info">
              <text class="rewards-text">获得经验：+{{ todayCheckin.rewards }}</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 快速导航 -->
      <view class="quick-nav">
        <view class="nav-item" @click="goToHistory">
          <i class="fas fa-history"></i>
          <text>打卡历史</text>
        </view>
        <view class="nav-item" @click="goToWeekly">
          <i class="fas fa-chart-line"></i>
          <text>本周进度</text>
        </view>
        <view class="nav-item" @click="goToCalendar">
          <i class="fas fa-calendar"></i>
          <text>月度日历</text>
        </view>
      </view>
    </view>

    <!-- 成功提示弹窗 -->
    <view class="success-modal" v-if="showSuccessModal" @click="closeSuccessModal">
      <view class="modal-content" @click.stop>
        <view class="success-icon">
          <i class="fas fa-check-circle"></i>
        </view>
        <text class="success-title">打卡成功！</text>
        <view class="success-details">
          <text>连续打卡：{{ checkinResult.streak }} 天</text>
          <text>获得经验：+{{ checkinResult.rewards }}</text>
          <text v-if="checkinResult.levelUp">恭喜升级到 Lv.{{ checkinResult.newLevel }}！</text>
        </view>
        <NfButton type="primary" @click="closeSuccessModal">继续</NfButton>
      </view>
    </view>
  </view>
</template>

<script>
import { ref, reactive, onMounted, computed } from 'vue'
import { onShow } from '@dcloudio/uni-app'

export default {
  name: 'CheckinIndex',
  setup() {
    // 响应式数据
    const currentTime = ref('')
    const hasCheckedToday = ref(false)
    const selectedMood = ref(0)
    const checkinNotes = ref('')
    const isSubmitting = ref(false)
    const showSuccessModal = ref(false)
    
    const todayCheckin = reactive({
      checkinTime: '',
      moodLevel: 0,
      notes: '',
      rewards: 0
    })
    
    const statistics = reactive({
      currentStreak: 0,
      totalDays: 0,
      level: 1,
      experience: 0
    })
    
    const checkinResult = reactive({
      streak: 0,
      rewards: 0,
      levelUp: false,
      newLevel: 0
    })

    // 心情选项
    const moodOptions = [
      { level: 1, emoji: '😰', label: '很糟糕' },
      { level: 2, emoji: '😟', label: '不太好' },
      { level: 3, emoji: '😐', label: '一般' },
      { level: 4, emoji: '😊', label: '不错' },
      { level: 5, emoji: '😄', label: '很棒' }
    ]

    // 计算属性
    const currentDate = computed(() => {
      const now = new Date()
      const year = now.getFullYear()
      const month = String(now.getMonth() + 1).padStart(2, '0')
      const day = String(now.getDate()).padStart(2, '0')
      return `${year}-${month}-${day}`
    })

    const currentWeekday = computed(() => {
      const weekdays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
      return weekdays[new Date().getDay()]
    })

    // 方法
    const updateTime = () => {
      const now = new Date()
      const hours = now.getHours().toString().padStart(2, '0')
      const minutes = now.getMinutes().toString().padStart(2, '0')
      currentTime.value = `${hours}:${minutes}`
    }

    const selectMood = (level) => {
      selectedMood.value = level
    }

    const getMoodEmoji = (level) => {
      const mood = moodOptions.find(m => m.level === level)
      return mood ? mood.emoji : '😐'
    }

    const getMoodLabel = (level) => {
      const mood = moodOptions.find(m => m.level === level)
      return mood ? mood.label : '一般'
    }

    const loadTodayStatus = async () => {
      try {
        const token = uni.getStorageSync('token')
        if (!token) {
          console.log('用户未登录')
          return
        }

        const res = await uni.request({
          url: 'http://localhost:8888/api/v1/miniprogram/checkin/today',
          method: 'GET',
          header: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        })

        if (res.data.code === 0) {
          const data = res.data.data
          hasCheckedToday.value = data.hasChecked
          
          if (data.hasChecked) {
            todayCheckin.checkinTime = data.checkinTime ? 
              new Date(data.checkinTime).toLocaleTimeString('zh-CN', { 
                hour: '2-digit', 
                minute: '2-digit' 
              }) : ''
            todayCheckin.moodLevel = data.moodLevel
            todayCheckin.notes = data.notes || ''
            todayCheckin.rewards = data.rewards || 0
          }
        }
      } catch (error) {
        console.error('获取今日状态失败:', error)
      }
    }

    const loadStatistics = async () => {
      try {
        const token = uni.getStorageSync('token')
        const res = await uni.request({
          url: 'http://localhost:8888/api/v1/miniprogram/checkin/statistics',
          method: 'GET',
          header: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        })

        if (res.data.code === 0) {
          const data = res.data.data
          Object.assign(statistics, data)
        }
      } catch (error) {
        console.error('获取统计数据失败:', error)
      }
    }

    const submitCheckin = async () => {
      if (!selectedMood.value || isSubmitting.value) return
      
      isSubmitting.value = true
      
      try {
        const token = uni.getStorageSync('token')
        const res = await uni.request({
          url: 'http://localhost:8888/api/v1/miniprogram/checkin/daily',
          method: 'POST',
          header: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          },
          data: {
            moodLevel: selectedMood.value,
            notes: checkinNotes.value
          }
        })

        if (res.data.code === 0) {
          const data = res.data.data
          
          // 更新结果数据
          Object.assign(checkinResult, {
            streak: data.streak,
            rewards: data.rewards,
            levelUp: data.levelUp,
            newLevel: data.newLevel
          })
          
          // 显示成功弹窗
          showSuccessModal.value = true
          
          // 重新加载数据
          await loadTodayStatus()
          await loadStatistics()
          
        } else {
          uni.showToast({
            title: res.data.msg || '打卡失败',
            icon: 'none'
          })
        }
      } catch (error) {
        console.error('打卡失败:', error)
        uni.showToast({
          title: '网络错误，请重试',
          icon: 'none'
        })
      } finally {
        isSubmitting.value = false
      }
    }

    const closeSuccessModal = () => {
      showSuccessModal.value = false
    }

    const viewHistory = () => {
      uni.navigateTo({ url: '/pages/checkin/history' })
    }

    const goToHistory = () => {
      uni.navigateTo({ url: '/pages/checkin/history' })
    }

    const goToWeekly = () => {
      uni.navigateTo({ url: '/pages/checkin/weekly' })
    }

    const goToCalendar = () => {
      uni.navigateTo({ url: '/pages/checkin/calendar' })
    }

    // 生命周期
    onMounted(() => {
      updateTime()
      setInterval(updateTime, 60000) // 每分钟更新一次时间
      loadTodayStatus()
      loadStatistics()
    })

    onShow(() => {
      loadTodayStatus()
      loadStatistics()
    })

    return {
      // 数据
      currentTime,
      hasCheckedToday,
      selectedMood,
      checkinNotes,
      isSubmitting,
      showSuccessModal,
      todayCheckin,
      statistics,
      checkinResult,
      moodOptions,
      
      // 计算属性
      currentDate,
      currentWeekday,
      
      // 方法
      updateTime,
      selectMood,
      getMoodEmoji,
      getMoodLabel,
      submitCheckin,
      closeSuccessModal,
      viewHistory,
      goToHistory,
      goToWeekly,
      goToCalendar
    }
  }
}
</script>

<style scoped>
/* 页面容器 */
.checkin-page {
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
  padding: 80px 16px 100px;
  overflow-y: auto;
}

/* 今日状态卡片 */
.status-card {
  margin-bottom: 24px;
}

.status-content {
  text-align: center;
  padding: 20px 0;
}

.date-info {
  margin-bottom: 20px;
}

.date {
  display: block;
  font-size: 24px;
  font-weight: bold;
  color: #333;
  margin-bottom: 4px;
}

.weekday {
  font-size: 14px;
  color: #666;
}

.checkin-status {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.status-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: #f0f0f0;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 12px;
  font-size: 24px;
  color: #ccc;
}

.status-icon.checked {
  background: #4CAF50;
  color: white;
}

.status-text {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin-bottom: 4px;
}

.status-desc {
  font-size: 14px;
  color: #666;
}

/* 连续天数统计 */
.stats-section {
  padding: 0 20px 20px;
}

.stats-grid {
  display: flex;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 12px;
  overflow: hidden;
}

.stat-item {
  flex: 1;
  padding: 20px;
  text-align: center;
  border-right: 1px solid #eee;
}

.stat-item:last-child {
  border-right: none;
}

.stat-number {
  display: block;
  font-size: 24px;
  font-weight: bold;
  color: var(--primary-color);
  margin-bottom: 4px;
}

.stat-label {
  font-size: 12px;
  color: #666;
}

/* 心情选择 */
.mood-section,
.notes-section {
  margin: 20px;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 12px;
  padding: 20px;
}

.section-title {
  margin-bottom: 16px;
}

.section-title text {
  font-size: 16px;
  font-weight: bold;
  color: #333;
}

.mood-selector {
  display: flex;
  justify-content: space-between;
  gap: 8px;
}

.mood-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 12px 8px;
  border-radius: 8px;
  border: 2px solid #eee;
  background: #f9f9f9;
  transition: all 0.3s ease;
}

.mood-item.active {
  border-color: var(--primary-color);
  background: var(--primary-light);
}

.mood-emoji {
  font-size: 24px;
  margin-bottom: 4px;
}

.mood-label {
  font-size: 12px;
  color: #666;
}

.notes-input {
  width: 100%;
  min-height: 80px;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 14px;
  line-height: 1.4;
  background: #fff;
  resize: none;
}

.notes-count {
  display: block;
  text-align: right;
  font-size: 12px;
  color: #999;
  margin-top: 8px;
}

/* 打卡按钮 */
.action-section {
  padding: 20px;
}

.checked-section {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 12px;
  padding: 20px;
}

.checked-info {
  text-align: center;
}

.checked-mood {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 12px;
}

.checked-mood .mood-emoji {
  font-size: 20px;
  margin-right: 8px;
}

.mood-text {
  font-size: 16px;
  color: #333;
}

.checked-notes {
  margin-bottom: 12px;
  padding: 12px;
  background: #f5f5f5;
  border-radius: 8px;
  text-align: left;
}

.notes-label {
  display: block;
  font-size: 14px;
  color: #666;
  margin-bottom: 4px;
}

.notes-text {
  font-size: 14px;
  color: #333;
  line-height: 1.4;
}

.rewards-info {
  font-size: 14px;
  color: var(--primary-color);
  font-weight: bold;
}

/* 快速导航 */
.quick-nav {
  display: flex;
  padding: 0 20px 20px;
  gap: 12px;
}

.nav-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 16px;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 12px;
  transition: transform 0.2s ease;
}

.nav-item:active {
  transform: scale(0.95);
}

.nav-item i {
  font-size: 20px;
  color: var(--primary-color);
  margin-bottom: 8px;
}

.nav-item text {
  font-size: 12px;
  color: #333;
}

/* 成功提示弹窗 */
.success-modal {
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

.modal-content {
  background: white;
  border-radius: 12px;
  padding: 30px;
  margin: 20px;
  text-align: center;
  min-width: 280px;
}

.success-icon {
  font-size: 48px;
  color: #4CAF50;
  margin-bottom: 16px;
}

.success-title {
  display: block;
  font-size: 20px;
  font-weight: bold;
  color: #333;
  margin-bottom: 16px;
}

.success-details {
  margin-bottom: 24px;
}

.success-details text {
  display: block;
  font-size: 14px;
  color: #666;
  margin-bottom: 4px;
}
</style> 