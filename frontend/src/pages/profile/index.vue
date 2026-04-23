<template>
  <view class="profile-container">
    <!-- Status Bar -->
    <view class="status-bar">
      <text class="time">{{ currentTime }}</text>
      <view class="status-icons">
        <text class="icon">📶</text>
        <text class="icon">📶</text>
        <view class="battery">
          <view class="battery-fill"></view>
        </view>
      </view>
    </view>

    <!-- Header with Profile -->
    <view class="header-section">
      <view class="header-top">
        <text class="page-title">个人中心</text>
        <view class="header-actions">
          <view class="edit-btn" @click="goToSetup">
            <text class="edit-icon">✏️</text>
          </view>
          <view class="settings-btn" @click="goToSettings">
            <text class="settings-icon">⚙️</text>
          </view>
        </view>
      </view>

      <!-- Profile Card -->
      <view class="profile-card">
        <view class="profile-info">
          <view class="avatar-section" @click="handleAvatarClick">
            <view class="avatar" :class="{ 'clickable': isLoggedIn }">
              <image 
                v-if="isLoggedIn && userProfile.avatarUrl" 
                :src="userProfile.avatarUrl" 
                class="avatar-image"
              />
              <text v-else-if="isLoggedIn" class="avatar-emoji">🌱</text>
              <text v-else class="login-icon">🔐</text>
            </view>
            <text v-if="!isLoggedIn" class="login-hint">点击微信登录</text>
            <text v-else class="avatar-hint">点击设置头像昵称</text>
          </view>
          <view class="user-details">
            <text class="username">{{ isLoggedIn ? (userProfile.nickname || '用户') : '未登录' }}</text>
            <text v-if="isLoggedIn" class="level-text">等级 {{ userProfile.level || 1 }} - {{ userProfile.levelTitle || '新手' }}</text>
            <text v-if="isLoggedIn" class="join-days">加入第 {{ userProfile.joinDays || 1 }} 天</text>
            <text v-else class="login-desc">登录后查看详细数据</text>
          </view>
        </view>
        
        <view class="stats-grid">
          <view class="stat-item">
            <text class="stat-number">{{ isLoggedIn ? userProfile.currentStreak : '--' }}</text>
            <text class="stat-label">当前连击</text>
          </view>
          <view class="stat-item">
            <text class="stat-number">{{ isLoggedIn ? userProfile.longestStreak : '--' }}</text>
            <text class="stat-label">最长连击</text>
          </view>
          <view class="stat-item">
            <text class="stat-number">{{ isLoggedIn ? userProfile.experience : '--' }}</text>
            <text class="stat-label">经验值</text>
          </view>
        </view>
      </view>
    </view>

    <!-- Main Content -->
    <view class="main-content">
      <!-- Quick Stats -->
      <view class="quick-stats">
        <view class="stat-card">
          <view class="stat-icon achievement-icon">🏆</view>
          <text class="stat-title">成就徽章</text>
          <text class="stat-value">{{ isLoggedIn ? userProfile.achievementCount : '--' }}</text>
        </view>
        
        <view class="stat-card">
          <view class="stat-icon help-icon">👥</view>
          <text class="stat-title">帮助他人</text>
          <text class="stat-value">{{ isLoggedIn ? userProfile.helpCount : '--' }}</text>
        </view>
      </view>

      <!-- Recent Achievements -->
      <view class="achievements-section" v-if="isLoggedIn">
        <text class="section-title">最新成就</text>
        
        <view class="achievements-list">
          <view 
            v-for="achievement in recentAchievements" 
            :key="achievement.id"
            class="achievement-item"
          >
            <view class="achievement-icon" :class="achievement.rarity">
              <text>{{ achievement.icon }}</text>
            </view>
            <view class="achievement-info">
              <text class="achievement-title">{{ achievement.title }}</text>
              <text class="achievement-time">{{ achievement.daysAgo }}天前获得</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 未登录提示 -->
      <view class="login-prompt" v-if="!isLoggedIn">
        <view class="prompt-icon">🔒</view>
        <text class="prompt-title">登录查看更多</text>
        <text class="prompt-desc">登录后可查看个人成就、进度统计等详细信息</text>
        <button class="wx-login-btn" @click="handleWxLogin" :disabled="loggingIn">
          <text class="wx-icon">🟢</text>
          <text class="wx-text">{{ loggingIn ? '登录中...' : '微信快速登录' }}</text>
        </button>
      </view>

      <!-- Menu Options -->
      <view class="menu-section">
        <view class="menu-item" @click="handleAuth" v-if="isLoggedIn">
          <view class="menu-icon auth-icon">🚪</view>
          <text class="menu-text">退出登录</text>
          <text class="menu-arrow">></text>
        </view>

        <view class="menu-item" @click="handleDataExport" v-if="isLoggedIn">
          <view class="menu-icon export-icon">📊</view>
          <text class="menu-text">数据导出</text>
          <text class="menu-arrow">></text>
        </view>

        <view class="menu-item" @click="goToPrivacySettings">
          <view class="menu-icon privacy-icon">🛡️</view>
          <text class="menu-text">隐私设置</text>
          <text class="menu-arrow">></text>
        </view>

        <view class="menu-item" @click="goToNotificationSettings">
          <view class="menu-icon notification-icon">🔔</view>
          <text class="menu-text">通知设置</text>
          <text class="menu-arrow">></text>
        </view>

        <view class="menu-item" @click="goToHelpCenter">
          <view class="menu-icon help-center-icon">❓</view>
          <text class="menu-text">帮助中心</text>
          <text class="menu-arrow">></text>
        </view>


      </view>
    </view>

    <!-- Data Export Modal -->
    <view v-if="showExportModal" class="modal-overlay" @click="closeExportModal">
      <view class="modal-content" @click.stop>
        <view class="modal-header">
          <text class="modal-title">数据导出</text>
          <view class="modal-close" @click="closeExportModal">
            <text class="close-icon">×</text>
          </view>
        </view>

        <view class="modal-body">
          <view class="export-options">
            <text class="option-label">导出格式</text>
            <picker 
              mode="selector" 
              :range="['JSON格式', 'CSV格式', 'Excel格式']" 
              @change="onExportFormatChange"
            >
              <view class="picker-input">
                {{ ['JSON格式', 'CSV格式', 'Excel格式'][exportFormatIndex] }}
              </view>
            </picker>
          </view>

          <view class="export-options">
            <text class="option-label">导出内容</text>
            <checkbox-group @change="onDataTypeChange">
              <label class="checkbox-item">
                <checkbox value="profile" checked />
                <text>个人资料</text>
              </label>
              <label class="checkbox-item">
                <checkbox value="checkin" checked />
                <text>打卡记录</text>
              </label>
              <label class="checkbox-item">
                <checkbox value="achievements" />
                <text>成就记录</text>
              </label>
            </checkbox-group>
          </view>

          <view class="export-options">
            <text class="option-label">时间范围</text>
            <picker 
              mode="selector" 
              :range="dateRangeOptions" 
              @change="onDateRangeChange"
            >
              <view class="picker-input">
                {{ dateRangeOptions[dateRangeIndex] }}
              </view>
            </picker>
          </view>
        </view>

        <view class="modal-footer">
          <button class="export-btn" @click="startExport" :disabled="exportDataTypes.length === 0">
            开始导出
          </button>
        </view>
      </view>
    </view>

    <!-- Loading Toast -->
    <view v-if="showLoading" class="loading-toast">
      <view class="loading-content">
        <view class="loading-spinner"></view>
        <text class="loading-text">{{ loadingText }}</text>
      </view>
    </view>
    
    <!-- 底部导航 -->
    <nf-tab-bar current="profile" />
  </view>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { isLoggedIn as checkIsLoggedIn, getUserInfo, wxLogin, logout, getAvatarUrl, setUserInfo, getToken } from '@/utils/auth'
import { userApi, achievementApi } from '@/apis'
import { buildApiUrl } from '@/config/env.js'
import NfTabBar from '@/components/ui/navigation/NfTabBar.vue'

// 响应式数据
const currentTime = ref('9:41')
const showExportModal = ref(false)
const showLoading = ref(false)
const loadingText = ref('加载中...')
const exportFormat = ref('json')
const exportFormatIndex = ref(0)
const exportDataTypes = ref(['profile', 'checkin'])
const dateRangeIndex = ref(0)
const loggingIn = ref(false)

// 登录状态
const isUserLoggedIn = ref(false)

// 用户资料数据
const userProfile = ref({
  nickname: '',
  avatarUrl: '',
  level: 1,
  levelTitle: '新手',
  joinDays: 1,
  currentStreak: 0,
  longestStreak: 0,
  experience: 0,
  achievementCount: 0,
  helpCount: 0
})

// 最近成就数据
const recentAchievements = ref([])

// 导出选项
const dateRangeOptions = ref(['全部时间', '最近一个月', '最近三个月', '最近一年'])

// 计算属性
const currentTimeDisplay = computed(() => {
  const now = new Date()
  return `${now.getHours()}:${now.getMinutes().toString().padStart(2, '0')}`
})

// 是否已登录
const isLoggedIn = computed(() => isUserLoggedIn.value)

// 生命周期
onMounted(() => {
  updateTime()
  checkLoginStatus()
  if (isUserLoggedIn.value) {
    loadUserProfile()
  }
})

// 方法
const updateTime = () => {
  currentTime.value = currentTimeDisplay.value
  setInterval(() => {
    currentTime.value = currentTimeDisplay.value
  }, 60000) // 每分钟更新一次
}

const checkLoginStatus = () => {
  isUserLoggedIn.value = checkIsLoggedIn()
  console.log('检查登录状态:', isUserLoggedIn.value)
  
  if (isUserLoggedIn.value) {
    const userInfo = getUserInfo()
    console.log('本地存储的用户信息:', userInfo)
    
    if (userInfo) {
      // 设置用户基本信息
      userProfile.value.nickname = userInfo.nickname || '微信用户'
      userProfile.value.avatarUrl = getAvatarUrl(userInfo.avatarUrl)
      console.log('设置用户资料:', {
        nickname: userProfile.value.nickname,
        avatarUrl: userProfile.value.avatarUrl
      })
    }
  }
}

const loadUserProfile = async () => {
  if (!isUserLoggedIn.value) return
  
  try {
    showLoading.value = true
    loadingText.value = '加载个人资料...'
    
    console.log('开始加载用户资料...')
    
    // 调用实际API获取用户资料
    const res = await userApi.getProfile()
    console.log('用户资料API响应:', res.data)
    
    if (res.data.code === 0) {
      const data = res.data.data
      console.log('解析的用户数据:', data)
      
      // 更新用户资料，优先使用API返回的数据
      // Go模型字段是 AvatarUrl，JSON序列化后是 avatarUrl
      const updatedProfile = {
        nickname: data.user?.nickname || userProfile.value.nickname || '微信用户',
        avatarUrl: getAvatarUrl(data.user?.avatarUrl) || userProfile.value.avatarUrl,
        level: data.user?.level || 1,
        levelTitle: data.user?.levelTitle || '新手',
        joinDays: data.abstinenceRecord?.totalDays || 1,
        currentStreak: data.abstinenceRecord?.currentStreak || 0,
        longestStreak: data.abstinenceRecord?.longestStreak || 0,
        experience: data.user?.experience || 0,
        achievementCount: data.user?.achievementCount || 0,
        helpCount: data.user?.helpCount || 0
      }
      
      console.log('更新后的用户资料:', updatedProfile)
      userProfile.value = updatedProfile
      
      // 更新本地存储
      const currentUser = getUserInfo()
      if (currentUser) {
        const updatedUser = {
          ...currentUser,
          nickname: updatedProfile.nickname,
          avatarUrl: data.user?.avatarUrl || currentUser.avatarUrl,
          level: updatedProfile.level,
          experience: updatedProfile.experience
        }
        setUserInfo(updatedUser)
        console.log('更新本地用户信息:', updatedUser)
      }
      
      // 加载最新成就
      await loadRecentAchievements()
    } else {
      console.error('API返回错误:', res.data.msg)
      uni.showToast({
        title: res.data.msg || '加载失败',
        icon: 'error'
      })
    }
    
    showLoading.value = false
  } catch (error) {
    console.error('加载用户资料失败:', error)
    showLoading.value = false
    uni.showToast({
      title: '加载失败',
      icon: 'error'
    })
  }
}

const loadRecentAchievements = async () => {
  try {
    const res = await achievementApi.getUserAchievements({ limit: 3, recent: true })
    
    if (res.data.code === 0) {
      recentAchievements.value = res.data.data.list || []
    }
  } catch (error) {
    console.error('加载成就数据失败:', error)
  }
}

const handleAvatarClick = () => {
  if (!isUserLoggedIn.value) {
    handleWxLogin()
  } else {
    // 跳转到用户信息授权页面
    goToAuthPage()
  }
}

const handleWxLogin = async () => {
  if (loggingIn.value) return
  
  try {
    loggingIn.value = true
    showLoading.value = true
    loadingText.value = '微信登录中...'
    
    // 直接进行基础登录（不获取敏感信息）
    const result = await wxLogin()
    
    // 登录成功，更新状态
    isUserLoggedIn.value = true
    checkLoginStatus()
    await loadUserProfile()
    
    uni.showToast({
      title: '登录成功',
      icon: 'success'
    })
    
    // 登录成功后，询问是否要完善个人信息
    setTimeout(() => {
      uni.showModal({
        title: '完善个人信息',
        content: '是否要设置头像和昵称，获得更好的使用体验？',
        confirmText: '去设置',
        cancelText: '暂不',
        success: (res) => {
          if (res.confirm) {
            uni.navigateTo({
              url: '/pages/profile/auth'
            })
          }
        }
      })
    }, 1500)
    
  } catch (error) {
    console.error('微信登录失败:', error)
    uni.showToast({
      title: error.message || '登录失败',
      icon: 'error'
    })
  } finally {
    loggingIn.value = false
    showLoading.value = false
  }
}

const goToSettings = () => {
  uni.navigateTo({
    url: '/pages/profile/privacy'
  })
}

const goToSetup = () => {
  uni.navigateTo({
    url: '/pages/profile/setup'
  })
}

const goToAuthPage = () => {
  uni.navigateTo({
    url: '/pages/profile/auth'
  })
}

const handleDataExport = () => {
  showExportModal.value = true
}

const closeExportModal = () => {
  showExportModal.value = false
  exportDataTypes.value = ['profile', 'checkin']
  exportFormat.value = 'json'
  exportFormatIndex.value = 0
  dateRangeIndex.value = 0
}

const onExportFormatChange = (e) => {
  exportFormatIndex.value = e.detail.value
  const formats = ['json', 'csv', 'excel']
  exportFormat.value = formats[e.detail.value]
}

const onDataTypeChange = (e) => {
  const { value } = e.detail
  exportDataTypes.value = value
}

const onDateRangeChange = (e) => {
  dateRangeIndex.value = e.detail.value
}

const startExport = async () => {
  if (exportDataTypes.value.length === 0) {
    uni.showToast({
      title: '请选择导出内容',
      icon: 'none'
    })
    return
  }

  try {
    showLoading.value = true
    loadingText.value = '导出数据中...'
    
    const exportReq = {
      format: exportFormat.value,
      dataTypes: exportDataTypes.value,
      dateRange: dateRangeIndex.value
    }
    
    const res = await userApi.createDataExport(exportReq)
    
    if (res.data.code === 0) {
      closeExportModal()
      uni.showToast({
        title: '导出成功',
        icon: 'success'
      })
    } else {
      throw new Error(res.data.msg || '导出失败')
    }
  } catch (error) {
    console.error('数据导出失败:', error)
    uni.showToast({
      title: error.message || '导出失败',
      icon: 'error'
    })
  } finally {
    showLoading.value = false
  }
}

const goToPrivacySettings = () => {
  uni.navigateTo({
    url: '/pages/profile/privacy'
  })
}

const goToNotificationSettings = () => {
  uni.navigateTo({
    url: '/pages/profile/notification'
  })
}

const goToHelpCenter = () => {
  uni.navigateTo({
    url: '/pages/learning/help'
  })
}

const handleAuth = () => {
  if (isUserLoggedIn.value) {
    // 退出登录
    uni.showModal({
      title: '确认退出',
      content: '确定要退出登录吗？',
      success: (res) => {
        if (res.confirm) {
          logout()
          isUserLoggedIn.value = false
          userProfile.value = {
            nickname: '',
            avatarUrl: '',
            level: 1,
            levelTitle: '新手',
            joinDays: 1,
            currentStreak: 0,
            longestStreak: 0,
            experience: 0,
            achievementCount: 0,
            helpCount: 0
          }
          recentAchievements.value = []
          uni.showToast({
            title: '已退出登录',
            icon: 'success'
          })
        }
      }
    })
  }
}

// 选择头像
const chooseAvatar = () => {
  uni.chooseMedia({
    count: 1,
    mediaType: ['image'],
    sourceType: ['album', 'camera'],
    maxDuration: 30,
    camera: 'back',
    success: (res) => {
      console.log('选择头像成功:', res)
      const tempFilePath = res.tempFiles[0].tempFilePath
      uploadAvatar(tempFilePath)
    },
    fail: (err) => {
      console.error('选择头像失败:', err)
      uni.showToast({
        title: '选择头像失败',
        icon: 'none'
      })
    }
  })
}

// 上传头像
const uploadAvatar = async (filePath) => {
  try {
    uni.showLoading({
      title: '上传中...'
    })

    // 这里需要实现头像上传逻辑
    // 可以上传到你的服务器或使用云存储
    const uploadResult = await uploadFile(filePath)
    
    if (uploadResult.success) {
      // 更新用户头像
      await updateUserAvatar(uploadResult.url)
      userProfile.value.avatarUrl = uploadResult.url
      
      uni.showToast({
        title: '头像更新成功',
        icon: 'success'
      })
    }
  } catch (error) {
    console.error('上传头像失败:', error)
    uni.showToast({
      title: '上传失败',
      icon: 'none'
    })
  } finally {
    uni.hideLoading()
  }
}

// 上传文件到服务器
const uploadFile = async (filePath) => {
  return new Promise((resolve, reject) => {
    uni.uploadFile({
      url: buildApiUrl('/user/upload-avatar'),
      filePath: filePath,
      name: 'file',
      header: {
        'Authorization': `Bearer ${getToken()}`
      },
      success: (res) => {
        const data = JSON.parse(res.data)
        if (data.code === 0) {
          resolve({
            success: true,
            url: data.data.url
          })
        } else {
          reject(new Error(data.msg))
        }
      },
      fail: reject
    })
  })
}

// 更新用户头像
const updateUserAvatar = async (avatarUrl) => {
  const response = await userApi.updateUserInfo({
    avatarUrl: avatarUrl
  })
  return response
}
</script>

<style scoped>
.profile-container {
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
  background: transparent;
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

/* Header Section */
.header-section {
  background: linear-gradient(135deg, #34D399 0%, #10B981 100%);
  padding: 32rpx 48rpx 64rpx;
}

.header-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 48rpx;
}

.page-title {
  font-size: 40rpx;
  font-weight: bold;
  color: white;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 16rpx;
}

.edit-btn {
  width: 64rpx;
  height: 64rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(10rpx);
}

.edit-icon {
  font-size: 32rpx;
  color: white;
}

.settings-btn {
  width: 64rpx;
  height: 64rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(10rpx);
}

.settings-icon {
  font-size: 32rpx;
  color: white;
}

/* Profile Card */
.profile-card {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20rpx);
  border-radius: 32rpx;
  padding: 48rpx;
  color: white;
}

.profile-info {
  display: flex;
  align-items: center;
  gap: 32rpx;
  margin-bottom: 32rpx;
}

.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8rpx;
}

.avatar {
  width: 128rpx;
  height: 128rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
}

.avatar.clickable {
  cursor: pointer;
  transition: transform 0.2s ease;
}

.avatar.clickable:hover {
  transform: scale(1.05);
}

.avatar-image {
  width: 100%;
  height: 100%;
  border-radius: 50%;
}

.avatar-emoji, .login-icon {
  font-size: 48rpx;
}

.login-icon {
  color: rgba(255, 255, 255, 0.8);
}

.login-hint {
  font-size: 20rpx;
  color: rgba(255, 255, 255, 0.8);
  text-align: center;
}

.avatar-hint {
  font-size: 20rpx;
  color: rgba(255, 255, 255, 0.8);
  text-align: center;
}

.user-details {
  flex: 1;
}

.username {
  display: block;
  font-size: 40rpx;
  font-weight: bold;
  margin-bottom: 8rpx;
}

.level-text, .login-desc {
  display: block;
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.8);
  margin-bottom: 4rpx;
}

.join-days {
  display: block;
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.7);
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 32rpx;
  text-align: center;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-number {
  font-size: 48rpx;
  font-weight: bold;
  margin-bottom: 8rpx;
}

.stat-label {
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
}

/* Main Content */
.main-content {
  padding: 48rpx;
  padding-bottom: 160rpx; /* 为底部导航留空间 */
}

/* Quick Stats */
.quick-stats {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 32rpx;
  margin-bottom: 48rpx;
}

.stat-card {
  background: white;
  border-radius: 32rpx;
  padding: 32rpx;
  text-align: center;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.stat-icon {
  width: 96rpx;
  height: 96rpx;
  border-radius: 24rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 16rpx;
  font-size: 40rpx;
}

.achievement-icon {
  background: #FEF3C7;
}

.help-icon {
  background: #E0E7FF;
}

.stat-title {
  display: block;
  font-size: 28rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 8rpx;
}

.stat-value {
  display: block;
  font-size: 48rpx;
  font-weight: bold;
  color: #34D399;
}

/* Login Prompt */
.login-prompt {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 48rpx;
  text-align: center;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.prompt-icon {
  font-size: 80rpx;
  margin-bottom: 24rpx;
}

.prompt-title {
  display: block;
  font-size: 36rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 16rpx;
}

.prompt-desc {
  display: block;
  font-size: 28rpx;
  color: #6B7280;
  margin-bottom: 32rpx;
  line-height: 1.5;
}

.wx-login-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16rpx;
  background: #07C160;
  color: white;
  border: none;
  border-radius: 24rpx;
  padding: 24rpx 48rpx;
  font-size: 32rpx;
  font-weight: 600;
}

.wx-login-btn:disabled {
  opacity: 0.6;
}

.wx-icon {
  font-size: 32rpx;
}

/* Achievements Section */
.achievements-section {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin-bottom: 48rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.section-title {
  font-size: 36rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 32rpx;
}

.achievements-list {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
}

.achievement-item {
  display: flex;
  align-items: center;
  gap: 24rpx;
}

.achievement-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 24rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
}

.achievement-icon.gold {
  background: linear-gradient(135deg, #F59E0B 0%, #D97706 100%);
}

.achievement-icon.purple {
  background: linear-gradient(135deg, #8B5CF6 0%, #7C3AED 100%);
}

.achievement-icon.green {
  background: linear-gradient(135deg, #10B981 0%, #059669 100%);
}

.achievement-info {
  flex: 1;
}

.achievement-title {
  display: block;
  font-size: 28rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 4rpx;
}

.achievement-time {
  display: block;
  font-size: 24rpx;
  color: #6B7280;
}

/* Menu Section */
.menu-section {
  display: flex;
  flex-direction: column;
  gap: 32rpx;
}

.menu-item {
  background: white;
  border-radius: 32rpx;
  padding: 32rpx;
  display: flex;
  align-items: center;
  gap: 24rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.menu-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 24rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
}

.auth-icon {
  background: #DDD6FE;
}

.export-icon {
  background: #FED7AA;
}

.privacy-icon {
  background: #DDD6FE;
}

.notification-icon {
  background: #BFDBFE;
}

.help-center-icon {
  background: #BBF7D0;
}

.menu-text {
  flex: 1;
  font-size: 28rpx;
  font-weight: 600;
  color: #1F2937;
}

.menu-arrow {
  font-size: 28rpx;
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

.modal-content {
  background: white;
  border-radius: 32rpx;
  padding: 48rpx;
  margin: 48rpx;
  max-height: 80vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32rpx;
}

.modal-title {
  font-size: 36rpx;
  font-weight: 600;
  color: #1F2937;
}

.modal-close {
  width: 64rpx;
  height: 64rpx;
  border-radius: 50%;
  background: #F3F4F6;
  display: flex;
  align-items: center;
  justify-content: center;
}

.close-icon {
  font-size: 32rpx;
  color: #6B7280;
}

.modal-body {
  margin-bottom: 32rpx;
}

.export-options {
  margin-bottom: 32rpx;
}

.option-label {
  display: block;
  font-size: 28rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 16rpx;
}

.picker-input {
  background: #F9FAFB;
  border: 2rpx solid #E5E7EB;
  border-radius: 16rpx;
  padding: 24rpx;
  font-size: 28rpx;
  color: #1F2937;
}

.checkbox-item {
  display: flex;
  align-items: center;
  gap: 16rpx;
  margin-bottom: 16rpx;
  font-size: 28rpx;
  color: #1F2937;
}

.modal-footer {
  display: flex;
  justify-content: center;
}

.export-btn {
  background: #34D399;
  color: white;
  border: none;
  border-radius: 24rpx;
  padding: 24rpx 48rpx;
  font-size: 32rpx;
  font-weight: 600;
}

.export-btn:disabled {
  opacity: 0.6;
}

/* Loading Toast */
.loading-toast {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
}

.loading-content {
  background: white;
  border-radius: 24rpx;
  padding: 48rpx;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 24rpx;
}

.loading-spinner {
  width: 64rpx;
  height: 64rpx;
  border: 4rpx solid #E5E7EB;
  border-top: 4rpx solid #34D399;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.loading-text {
  font-size: 28rpx;
  color: #1F2937;
}

/* 底部安全间距 */
.profile-page {
  padding-bottom: 120rpx;
}


</style> 