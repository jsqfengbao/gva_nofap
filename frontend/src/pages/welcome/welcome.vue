<template>
  <view class="welcome-page">
    <!-- 背景装饰 -->
    <view class="bg-decoration">
      <view class="circle circle-1"></view>
      <view class="circle circle-2"></view>
      <view class="circle circle-3"></view>
    </view>

    <!-- 主要内容 -->
    <view class="welcome-content">
      <!-- Logo区域 -->
      <view class="logo-section">
        <view class="logo-container">
          <view class="logo">
            <text class="logo-emoji">🌟</text>
          </view>
          <view class="logo-glow"></view>
        </view>
        <text class="app-name">自律助手</text>
        <text class="app-slogan">重新定义自律 · 塑造更好的自己</text>
      </view>

      <!-- 功能介绍 -->
      <view class="features-section">
        <view class="feature-item">
          <text class="feature-icon">📊</text>
          <text class="feature-text">科学数据追踪</text>
        </view>
        <view class="feature-item">
          <text class="feature-icon">🎯</text>
          <text class="feature-text">目标管理系统</text>
        </view>
        <view class="feature-item">
          <text class="feature-icon">👥</text>
          <text class="feature-text">社区互助支持</text>
        </view>
        <view class="feature-item">
          <text class="feature-icon">🏆</text>
          <text class="feature-text">成就激励机制</text>
        </view>
      </view>

      <!-- 操作区域 -->
      <view class="login-section">
        <view class="login-intro">
          <text class="intro-title">开始您的自律之旅</text>
          <text class="intro-desc">登录后即可使用所有功能</text>
        </view>

        <!-- 微信登录按钮 -->
        <button 
          class="wx-login-btn"
          :class="{ 'loading': isLoading }"
          @click="handleWxLogin"
          :disabled="isLoading"
        >
          <view class="btn-content">
            <text class="wx-logo" v-if="!isLoading">🔑</text>
            <view class="loading-spinner" v-else>
              <view class="spinner"></view>
            </view>
            <text class="btn-text">{{ isLoading ? '登录中...' : '微信登录' }}</text>
          </view>
        </button>

        <!-- 隐私政策勾选 - 必须用户主动勾选 -->
        <view class="privacy-section">
          <view class="checkbox-container" @click="togglePrivacy">
            <view class="checkbox" :class="{ checked: agreedToPrivacy }">
              <text v-if="agreedToPrivacy">✓</text>
            </view>
            <text class="privacy-text">
              我已阅读并同意
              <text class="link" @click.stop="showAgreement('user')">《用户服务协议》</text>
              和
              <text class="link" @click.stop="showAgreement('privacy')">《隐私政策》</text>
            </text>
          </view>
        </view>
      </view>

      <!-- 版本信息 -->
      <view class="version-info">
        <text class="version-text">Version 1.0.0</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getApiUrl } from '@/config/index.js'
import { setUserInfo } from '@/utils/auth.js'

// 响应式数据
const isLoading = ref(false)
const agreedToPrivacy = ref(false) // 用户必须主动勾选

// 勾选隐私政策
const togglePrivacy = () => {
  agreedToPrivacy.value = !agreedToPrivacy.value
}

// 检查是否已登录
const checkLoginStatus = () => {
  const token = uni.getStorageSync('token')
  if (token && token !== 'guest_token') {
    // 已登录，跳转到首页
    uni.switchTab({
      url: '/pages/index/index'
    })
  }
}

// 临时开发模式：自动进入游客模式（生产环境关闭）
const enableDevelopmentMode = () => {
  const isDevelopment = process.env.NODE_ENV === 'development'
  const autoGuestMode = false // 审核期间关闭，让审核人员看到完整页面
  
  if (isDevelopment && autoGuestMode) {
    console.log('🚀 开发模式：自动进入游客模式')
    setTimeout(() => {
      guestMode()
    }, 2000) // 2秒后自动进入游客模式
  }
}

// 执行微信登录
const handleWxLogin = async () => {
  if (isLoading.value) return
  
  // 必须先勾选隐私政策
  if (!agreedToPrivacy.value) {
    uni.showToast({
      title: '请先阅读并同意用户协议和隐私政策',
      icon: 'none',
      duration: 2500
    })
    return
  }
  
  isLoading.value = true
  
  try {
    const loginRes = await new Promise((resolve, reject) => {
      uni.login({
        provider: 'weixin',
        success: resolve,
        fail: reject
      })
    })
    
    console.log('获取微信登录凭证成功:', loginRes)
    await performWxLogin(loginRes.code)
    
  } catch (error) {
    console.error('微信登录失败:', error)
    
    let errorTitle = '登录失败，请重试'
    if (error.errCode === 40013) {
      errorTitle = 'AppID配置错误，请联系开发者'
    } else if (error.errCode === 40125) {
      errorTitle = '小程序未激活，请联系开发者'
    }
    
    uni.showToast({
      title: errorTitle,
      icon: 'none',
      duration: 3000
    })
    
    isLoading.value = false
  }
}

// 执行微信登录请求
const performWxLogin = async (code) => {
  try {
    const requestData = { code }
    
    console.log('发送登录请求:', requestData)
    
    const res = await uni.request({
      url: getApiUrl('/auth/wx-login'),
      method: 'POST',
      header: {
        'Content-Type': 'application/json'
      },
      data: requestData
    })

    console.log('登录响应:', res.data)

    if (res.data.code === 0) {
      // 保存token和用户信息
      const { token, user } = res.data.data
      setUserInfo({ ...user, token })
      
      uni.showToast({
        title: '登录成功',
        icon: 'success'
      })
      
      // 检查是否需要完善资料
      const needsSetup = !user.nickname || !user.avatarUrl
      
      setTimeout(() => {
        if (needsSetup) {
          // 跳转到资料完善页面
          uni.navigateTo({
            url: '/pages/profile/setup'
          })
        } else {
          // 直接跳转到首页
          uni.switchTab({
            url: '/pages/index/index'
          })
        }
      }, 1500)
      
    } else {
      uni.showToast({
        title: res.data.msg || '登录失败',
        icon: 'none'
      })
    }
  } catch (error) {
    console.error('微信登录失败:', error)
    uni.showToast({
      title: '网络错误，请重试',
      icon: 'none'
    })
  } finally {
    isLoading.value = false
  }
}

// 游客模式


// 显示协议
const showAgreement = (type) => {
  const title = type === 'user' ? '用户协议' : '隐私政策'
  uni.showModal({
    title: title,
    content: `这里显示${title}的内容...`,
    showCancel: false,
    confirmText: '我知道了'
  })
}

// 生命周期
onMounted(() => {
  checkLoginStatus()
  
  // 启用开发模式（临时）
  enableDevelopmentMode()
})
</script>

<style lang="scss" scoped>
.welcome-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  position: relative;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.bg-decoration {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  
  .circle {
    position: absolute;
    border-radius: 50%;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10rpx);
  }
  
  .circle-1 {
    width: 300rpx;
    height: 300rpx;
    top: -150rpx;
    right: -100rpx;
    animation: float 6s ease-in-out infinite;
  }
  
  .circle-2 {
    width: 200rpx;
    height: 200rpx;
    bottom: 100rpx;
    left: -50rpx;
    animation: float 8s ease-in-out infinite reverse;
  }
  
  .circle-3 {
    width: 150rpx;
    height: 150rpx;
    top: 50%;
    right: 50rpx;
    animation: float 10s ease-in-out infinite;
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0rpx);
  }
  50% {
    transform: translateY(-40rpx);
  }
}

.welcome-content {
  flex: 1;
  padding: 100rpx 60rpx 60rpx;
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: relative;
  z-index: 1;
}

.logo-section {
  text-align: center;
  margin-bottom: 80rpx;
  
  .logo-container {
    position: relative;
    display: inline-block;
    margin-bottom: 40rpx;
    
    .logo {
      width: 200rpx;
      height: 200rpx;
      background: rgba(255, 255, 255, 0.2);
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      backdrop-filter: blur(10rpx);
      border: 2rpx solid rgba(255, 255, 255, 0.3);
      position: relative;
      z-index: 2;
      
      .logo-emoji {
        font-size: 100rpx;
      }
    }
    
    .logo-glow {
      position: absolute;
      top: -20rpx;
      left: -20rpx;
      right: -20rpx;
      bottom: -20rpx;
      background: linear-gradient(45deg, rgba(255, 255, 255, 0.3), rgba(255, 255, 255, 0.1));
      border-radius: 50%;
      filter: blur(20rpx);
      animation: pulse 3s ease-in-out infinite;
    }
  }
  
  .app-name {
    display: block;
    font-size: 56rpx;
    font-weight: bold;
    color: white;
    margin-bottom: 20rpx;
    text-shadow: 0 2rpx 10rpx rgba(0, 0, 0, 0.2);
  }
  
  .app-slogan {
    display: block;
    font-size: 28rpx;
    color: rgba(255, 255, 255, 0.9);
    line-height: 1.4;
  }
}

@keyframes pulse {
  0%, 100% {
    opacity: 0.7;
    transform: scale(1);
  }
  50% {
    opacity: 1;
    transform: scale(1.1);
  }
}

.features-section {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 40rpx;
  margin-bottom: 80rpx;
  
  .feature-item {
    background: rgba(255, 255, 255, 0.15);
    border-radius: 24rpx;
    padding: 40rpx 32rpx;
    text-align: center;
    backdrop-filter: blur(10rpx);
    border: 1rpx solid rgba(255, 255, 255, 0.2);
    
    .feature-icon {
      display: block;
      font-size: 48rpx;
      margin-bottom: 16rpx;
    }
    
    .feature-text {
      font-size: 28rpx;
      color: white;
      font-weight: 500;
    }
  }
}

.login-section {
  .login-intro {
    text-align: center;
    margin-bottom: 60rpx;
    
    .intro-title {
      display: block;
      font-size: 40rpx;
      font-weight: bold;
      color: white;
      margin-bottom: 20rpx;
    }
    
    .intro-desc {
      font-size: 28rpx;
      color: rgba(255, 255, 255, 0.8);
      line-height: 1.4;
    }
  }
  
  .guest-btn {
    width: 100%;
    height: 84rpx;
    background: rgba(255, 255, 255, 0.95);
    border: none;
    border-radius: 20rpx;
    
    &:active {
      background: white;
      transform: scale(0.98);
    }
  }
  
  .guest-text {
    color: #667eea;
    font-size: 32rpx;
    font-weight: 600;
  }
  
  .quick-actions {
    margin-top: 32rpx;
    
    .divider-text {
      display: block;
      text-align: center;
      color: rgba(255, 255, 255, 0.6);
      font-size: 24rpx;
      margin-bottom: 32rpx;
    }
    
    .wx-login-btn {
      width: 100%;
      height: 96rpx;
      background: linear-gradient(135deg, #1aad19 0%, #2dc653 100%);
      border: none;
      border-radius: 24rpx;
      color: white;
      font-size: 36rpx;
      font-weight: 600;
      transition: all 0.3s ease;
      box-shadow: 0 8rpx 24rpx rgba(26, 173, 25, 0.3);
      
      &:active {
        transform: scale(0.98);
      }
      
      &.loading {
        opacity: 0.8;
        transform: none;
      }
      
      .btn-content {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 20rpx;
      }
      
      .wx-logo {
        font-size: 36rpx;
      }
      
      .loading-spinner {
        width: 36rpx;
        height: 36rpx;
        
        .spinner {
          width: 100%;
          height: 100%;
          border: 4rpx solid rgba(255, 255, 255, 0.3);
          border-top: 4rpx solid white;
          border-radius: 50%;
          animation: spin 1s linear infinite;
        }
      }
      
      .btn-text {
        color: white;
      }
    }
  }
}

/* 隐私政策勾选区域 */
.privacy-section {
  margin-top: 40rpx;
  padding: 24rpx 32rpx;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 16rpx;
  
  .checkbox-container {
    display: flex;
    align-items: flex-start;
    gap: 20rpx;
    
    .checkbox {
      width: 36rpx;
      height: 36rpx;
      border: 2rpx solid rgba(255, 255, 255, 0.5);
      border-radius: 6rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-shrink: 0;
      margin-top: 4rpx;
      
      &.checked {
        background: #10b981;
        border-color: #10b981;
        color: white;
        font-size: 24rpx;
        font-weight: bold;
      }
    }
    
    .privacy-text {
      font-size: 24rpx;
      color: rgba(255, 255, 255, 0.9);
      line-height: 1.6;
      
      .link {
        color: #93c5fd;
        text-decoration: underline;
      }
    }
  }
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.version-info {
  text-align: center;
  margin-top: 40rpx;
  
  .version-text {
    font-size: 22rpx;
    color: rgba(255, 255, 255, 0.5);
  }
}
</style> 