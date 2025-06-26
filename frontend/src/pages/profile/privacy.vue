<template>
  <view class="privacy-container">
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
      <text class="page-title">隐私设置</text>
      <view class="placeholder"></view>
    </view>

    <!-- Main Content -->
    <view class="main-content">
      <!-- Privacy Settings -->
      <view class="settings-section">
        <view class="section-header">
          <text class="section-title">个人信息显示</text>
          <text class="section-desc">控制其他用户能看到您的哪些信息</text>
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">显示个人资料</text>
            <text class="setting-desc">允许其他用户查看您的基本信息</text>
          </view>
          <switch 
            :checked="settings.showProfile" 
            @change="onShowProfileChange"
            color="#34D399"
          />
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">显示统计数据</text>
            <text class="setting-desc">显示您的打卡天数、成就等数据</text>
          </view>
          <switch 
            :checked="settings.showStats" 
            @change="onShowStatsChange"
            color="#34D399"
          />
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">显示成就</text>
            <text class="setting-desc">在个人资料中显示获得的成就</text>
          </view>
          <switch 
            :checked="settings.showAchievements" 
            @change="onShowAchievementsChange"
            color="#34D399"
          />
        </view>
      </view>

      <!-- Social Settings -->
      <view class="settings-section">
        <view class="section-header">
          <text class="section-title">社交设置</text>
          <text class="section-desc">管理您的社交互动偏好</text>
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">允许好友申请</text>
            <text class="setting-desc">其他用户可以向您发送好友申请</text>
          </view>
          <switch 
            :checked="settings.allowFriendRequest" 
            @change="onAllowFriendRequestChange"
            color="#34D399"
          />
        </view>

        <view class="setting-item">
          <view class="setting-info">
            <text class="setting-title">显示在线状态</text>
            <text class="setting-desc">让其他用户知道您是否在线</text>
          </view>
          <switch 
            :checked="settings.showOnlineStatus" 
            @change="onShowOnlineStatusChange"
            color="#34D399"
          />
        </view>
      </view>

      <!-- Data Management -->
      <view class="data-section">
        <view class="section-header">
          <text class="section-title">数据管理</text>
          <text class="section-desc">管理您的个人数据</text>
        </view>

        <view class="data-item" @click="showDataExportInfo">
          <view class="data-icon export-icon">📊</view>
          <view class="data-info">
            <text class="data-title">数据导出</text>
            <text class="data-desc">导出您的所有数据</text>
          </view>
          <text class="data-arrow">></text>
        </view>

        <view class="data-item" @click="showDeleteAccountInfo">
          <view class="data-icon delete-icon">🗑️</view>
          <view class="data-info">
            <text class="data-title">删除账户</text>
            <text class="data-desc">永久删除您的账户和所有数据</text>
          </view>
          <text class="data-arrow">></text>
        </view>
      </view>

      <!-- Save Button -->
      <view class="save-section">
        <button class="save-btn" @click="saveSettings" :disabled="saving">
          {{ saving ? '保存中...' : '保存设置' }}
        </button>
      </view>
    </view>

    <!-- Data Export Info Modal -->
    <view v-if="showExportModal" class="modal-overlay" @click="closeExportModal">
      <view class="info-modal" @click.stop>
        <view class="modal-header">
          <text class="modal-title">数据导出</text>
          <view class="close-btn" @click="closeExportModal">×</view>
        </view>
        
        <view class="modal-content">
          <text class="modal-text">您可以导出以下数据：</text>
          <view class="export-list">
            <text class="export-item">• 打卡记录和统计数据</text>
            <text class="export-item">• 评估结果和历史</text>
            <text class="export-item">• 成就和等级信息</text>
            <text class="export-item">• 社区互动记录</text>
            <text class="export-item">• 学习进度和记录</text>
          </view>
          <text class="modal-note">导出的数据将以JSON或CSV格式提供，请在个人中心主页面进行导出操作。</text>
        </view>

        <view class="modal-footer">
          <button class="modal-btn secondary" @click="closeExportModal">
            我知道了
          </button>
          <button class="modal-btn primary" @click="goToProfile">
            去导出
          </button>
        </view>
      </view>
    </view>

    <!-- Delete Account Info Modal -->
    <view v-if="showDeleteModal" class="modal-overlay" @click="closeDeleteModal">
      <view class="info-modal" @click.stop>
        <view class="modal-header">
          <text class="modal-title">删除账户</text>
          <view class="close-btn" @click="closeDeleteModal">×</view>
        </view>
        
        <view class="modal-content">
          <text class="modal-text warning">⚠️ 删除账户是不可逆的操作</text>
          <view class="delete-warning">
            <text class="warning-item">• 您的所有数据将被永久删除</text>
            <text class="warning-item">• 包括打卡记录、成就、社区互动等</text>
            <text class="warning-item">• 删除后无法恢复</text>
            <text class="warning-item">• 需要重新注册才能使用</text>
          </view>
          <text class="modal-note">如果您确实需要删除账户，请联系客服处理。</text>
        </view>

        <view class="modal-footer">
          <button class="modal-btn secondary" @click="closeDeleteModal">
            取消
          </button>
          <button class="modal-btn danger" @click="contactSupport">
            联系客服
          </button>
        </view>
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
const showExportModal = ref(false)
const showDeleteModal = ref(false)

// 隐私设置
const settings = ref({
  showProfile: true,
  showStats: true,
  showAchievements: true,
  allowFriendRequest: true,
  showOnlineStatus: true
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
const onShowProfileChange = (e) => {
  settings.value.showProfile = e.detail.value
}

const onShowStatsChange = (e) => {
  settings.value.showStats = e.detail.value
}

const onShowAchievementsChange = (e) => {
  settings.value.showAchievements = e.detail.value
}

const onAllowFriendRequestChange = (e) => {
  settings.value.allowFriendRequest = e.detail.value
}

const onShowOnlineStatusChange = (e) => {
  settings.value.showOnlineStatus = e.detail.value
}

// 模态框控制
const showDataExportInfo = () => {
  showExportModal.value = true
}

const closeExportModal = () => {
  showExportModal.value = false
}

const showDeleteAccountInfo = () => {
  showDeleteModal.value = true
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
}

const goToProfile = () => {
  closeExportModal()
  uni.switchTab({
    url: '/pages/profile/index'
  })
}

const contactSupport = () => {
  closeDeleteModal()
  uni.showModal({
    title: '联系客服',
    content: '请发送邮件至 support@nofap-helper.com 或在社区发帖寻求帮助',
    showCancel: false,
    confirmText: '我知道了'
  })
}
</script>

<style scoped>
.privacy-container {
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

.section-header {
  padding: 32rpx 48rpx 16rpx;
  border-bottom: 1rpx solid #F3F4F6;
}

.section-title {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 8rpx;
}

.section-desc {
  display: block;
  font-size: 24rpx;
  color: #6B7280;
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
  font-size: 28rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 8rpx;
}

.setting-desc {
  display: block;
  font-size: 24rpx;
  color: #6B7280;
}

/* Data Section */
.data-section {
  background: white;
  border-radius: 32rpx;
  overflow: hidden;
  margin-bottom: 48rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.data-item {
  display: flex;
  align-items: center;
  gap: 24rpx;
  padding: 32rpx 48rpx;
  border-bottom: 1rpx solid #F3F4F6;
}

.data-item:last-child {
  border-bottom: none;
}

.data-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 24rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
}

.export-icon {
  background: #FED7AA;
}

.delete-icon {
  background: #FECACA;
}

.data-info {
  flex: 1;
}

.data-title {
  display: block;
  font-size: 28rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 8rpx;
}

.data-desc {
  display: block;
  font-size: 24rpx;
  color: #6B7280;
}

.data-arrow {
  font-size: 24rpx;
  color: #9CA3AF;
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

/* Modal Styles */
.modal-overlay {
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

.info-modal {
  background: white;
  border-radius: 32rpx;
  width: 640rpx;
  max-height: 80vh;
  overflow: hidden;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 32rpx 48rpx;
  border-bottom: 1rpx solid #E5E7EB;
}

.modal-title {
  font-size: 36rpx;
  font-weight: 600;
  color: #1F2937;
}

.close-btn {
  width: 64rpx;
  height: 64rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 48rpx;
  color: #6B7280;
}

.modal-content {
  padding: 32rpx 48rpx;
}

.modal-text {
  display: block;
  font-size: 28rpx;
  color: #1F2937;
  margin-bottom: 24rpx;
}

.modal-text.warning {
  color: #DC2626;
  font-weight: 600;
}

.export-list, .delete-warning {
  margin: 24rpx 0;
}

.export-item, .warning-item {
  display: block;
  font-size: 26rpx;
  color: #6B7280;
  margin-bottom: 12rpx;
  line-height: 1.5;
}

.warning-item {
  color: #DC2626;
}

.modal-note {
  display: block;
  font-size: 24rpx;
  color: #6B7280;
  line-height: 1.6;
  margin-top: 24rpx;
}

.modal-footer {
  display: flex;
  gap: 16rpx;
  padding: 32rpx 48rpx;
  border-top: 1rpx solid #E5E7EB;
}

.modal-btn {
  flex: 1;
  padding: 24rpx;
  border: none;
  border-radius: 16rpx;
  font-size: 28rpx;
  font-weight: 600;
}

.modal-btn.primary {
  background: #34D399;
  color: white;
}

.modal-btn.secondary {
  background: #F3F4F6;
  color: #6B7280;
}

.modal-btn.danger {
  background: #DC2626;
  color: white;
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