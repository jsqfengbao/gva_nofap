<template>
  <view class="notification-container">
    <!-- Status Bar -->
    <view class="status-bar">
      <text class="time">9:41</text>
      <view class="status-icons">
        <text class="icon">📶</text>
        <text class="icon">📶</text>
        <view class="battery">
          <view class="battery-fill"></view>
        </view>
      </view>
    </view>

    <!-- Header -->
    <view class="header">
      <view class="back-btn" @click="goBack">
        <text class="back-icon">←</text>
      </view>
      <text class="page-title">通知设置</text>
      <view class="placeholder"></view>
    </view>

    <!-- Main Content -->
    <view class="main-content">
      <!-- Settings List -->
      <view class="settings-section">
        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">打卡提醒</text>
            <text class="setting-desc">每日提醒您进行打卡</text>
          </view>
          <switch 
            :checked="settings.checkinReminder" 
            @change="onCheckinReminderChange"
            color="#34D399"
          />
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">社区回复提醒</text>
            <text class="setting-desc">收到社区回复时通知</text>
          </view>
          <switch 
            :checked="settings.communityReply" 
            @change="onCommunityReplyChange"
            color="#34D399"
          />
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">成就解锁提醒</text>
            <text class="setting-desc">获得新成就时通知</text>
          </view>
          <switch 
            :checked="settings.achievementUnlock" 
            @change="onAchievementUnlockChange"
            color="#34D399"
          />
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">周报提醒</text>
            <text class="setting-desc">每周发送进度报告</text>
          </view>
          <switch 
            :checked="settings.weeklyReport" 
            @change="onWeeklyReportChange"
            color="#34D399"
          />
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">紧急求助提醒</text>
            <text class="setting-desc">收到紧急求助时通知</text>
          </view>
          <switch 
            :checked="settings.emergencyAlert" 
            @change="onEmergencyAlertChange"
            color="#34D399"
          />
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">学习提醒</text>
            <text class="setting-desc">推荐学习内容时通知</text>
          </view>
          <switch 
            :checked="settings.learningReminder" 
            @change="onLearningReminderChange"
            color="#34D399"
          />
        </view>
      </view>

      <!-- Save Button -->
      <view class="save-section">
        <button class="save-btn" @click="saveSettings" :disabled="saving">
          {{ saving ? '保存中...' : '保存设置' }}
        </button>
      </view>
    </view>

    <!-- Loading Toast -->
    <view v-if="loading" class="loading-toast">
      <view class="loading-content">
        <view class="loading-spinner"></view>
        <text class="loading-text">{{ loadingText }}</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// 响应式数据
const loading = ref(false)
const saving = ref(false)
const loadingText = ref('加载中...')

// 通知设置
const settings = ref({
  checkinReminder: true,
  communityReply: true,
  achievementUnlock: true,
  weeklyReport: true,
  emergencyAlert: true,
  learningReminder: true
})

// 生命周期
onMounted(() => {
  loadSettings()
})

// 方法
const goBack = () => {
  uni.navigateBack()
}

const loadSettings = async () => {
  try {
    loading.value = true
    loadingText.value = '加载设置...'
    
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 800))
    
    // 这里应该调用实际的API
    // const response = await uni.request({
    //   url: '/api/v1/miniprogram/profile/settings',
    //   method: 'GET',
    //   header: {
    //     'Authorization': 'Bearer ' + uni.getStorageSync('token')
    //   }
    // })
    // settings.value = response.data.notificationSettings
    
    loading.value = false
  } catch (error) {
    console.error('加载设置失败:', error)
    loading.value = false
    uni.showToast({
      title: '加载失败',
      icon: 'error'
    })
  }
}

const saveSettings = async () => {
  try {
    saving.value = true
    
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // 这里应该调用实际的API
    // const response = await uni.request({
    //   url: '/api/v1/miniprogram/profile/notification',
    //   method: 'PUT',
    //   data: settings.value,
    //   header: {
    //     'Authorization': 'Bearer ' + uni.getStorageSync('token')
    //   }
    // })
    
    saving.value = false
    uni.showToast({
      title: '保存成功',
      icon: 'success'
    })
  } catch (error) {
    console.error('保存设置失败:', error)
    saving.value = false
    uni.showToast({
      title: '保存失败',
      icon: 'error'
    })
  }
}

// 开关变更事件
const onCheckinReminderChange = (e) => {
  settings.value.checkinReminder = e.detail.value
}

const onCommunityReplyChange = (e) => {
  settings.value.communityReply = e.detail.value
}

const onAchievementUnlockChange = (e) => {
  settings.value.achievementUnlock = e.detail.value
}

const onWeeklyReportChange = (e) => {
  settings.value.weeklyReport = e.detail.value
}

const onEmergencyAlertChange = (e) => {
  settings.value.emergencyAlert = e.detail.value
}

const onLearningReminderChange = (e) => {
  settings.value.learningReminder = e.detail.value
}
</script>

<style scoped>
.notification-container {
  min-height: 100vh;
  background: #F8FAFC;
}

/* Status Bar */
.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8rpx 32rpx;
  font-size: 24rpx;
  color: #1F2937;
}

.status-icons {
  display: flex;
  align-items: center;
  gap: 8rpx;
}

.icon {
  font-size: 20rpx;
}

.battery {
  width: 48rpx;
  height: 24rpx;
  border: 2rpx solid #1F2937;
  border-radius: 4rpx;
  position: relative;
}

.battery-fill {
  width: 32rpx;
  height: 16rpx;
  background: #10B981;
  border-radius: 2rpx;
  margin: 2rpx;
}

/* Header */
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 32rpx 48rpx;
  background: white;
  border-bottom: 1rpx solid #E5E7EB;
}

.back-btn {
  width: 64rpx;
  height: 64rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.back-icon {
  font-size: 32rpx;
  color: #34D399;
}

.page-title {
  font-size: 36rpx;
  font-weight: 600;
  color: #1F2937;
}

.placeholder {
  width: 64rpx;
}

/* Main Content */
.main-content {
  padding: 48rpx;
}

/* Settings Section */
.settings-section {
  background: white;
  border-radius: 32rpx;
  overflow: hidden;
  margin-bottom: 48rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.setting-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 32rpx 48rpx;
  border-bottom: 1rpx solid #F3F4F6;
}

.setting-item:last-child {
  border-bottom: none;
}

.setting-info {
  flex: 1;
}

.setting-title {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 8rpx;
}

.setting-desc {
  display: block;
  font-size: 24rpx;
  color: #6B7280;
}

/* Save Section */
.save-section {
  padding: 32rpx 0;
}

.save-btn {
  width: 100%;
  padding: 32rpx;
  background: #34D399;
  color: white;
  border: none;
  border-radius: 24rpx;
  font-size: 32rpx;
  font-weight: 600;
}

.save-btn:disabled {
  background: #D1D5DB;
  color: #9CA3AF;
}

/* Loading Toast */
.loading-toast {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: rgba(0, 0, 0, 0.8);
  border-radius: 16rpx;
  padding: 32rpx 48rpx;
  z-index: 2000;
}

.loading-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16rpx;
}

.loading-spinner {
  width: 64rpx;
  height: 64rpx;
  border: 4rpx solid rgba(255, 255, 255, 0.3);
  border-top: 4rpx solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

.loading-text {
  color: white;
  font-size: 28rpx;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style> 