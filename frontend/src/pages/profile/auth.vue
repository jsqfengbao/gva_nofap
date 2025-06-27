<template>
  <view class="auth-container">
    <!-- 顶部说明 -->
    <view class="header">
      <text class="title">完善个人信息</text>
      <text class="desc">为了提供更好的服务，请设置您的头像和昵称</text>
    </view>

    <!-- 头像选择区域 -->
    <view class="avatar-section">
      <text class="section-title">头像设置</text>
      <view class="avatar-container">
        <!-- 使用微信官方的头像选择API -->
        <button 
          class="avatar-btn"
          open-type="chooseAvatar"
          @chooseavatar="onChooseAvatar"
        >
          <image 
            v-if="userInfo.avatarUrl" 
            :src="userInfo.avatarUrl" 
            class="avatar-image"
            mode="aspectFill"
          />
          <view v-else class="avatar-placeholder">
            <text class="camera-icon">📷</text>
            <text class="placeholder-text">点击选择头像</text>
          </view>
        </button>
      </view>
      <text class="hint-text">点击上方按钮选择微信头像</text>
    </view>

    <!-- 昵称输入区域 -->
    <view class="nickname-section">
      <text class="section-title">昵称设置</text>
      <view class="nickname-container">
        <!-- 使用 type="nickname" 获取微信昵称建议 -->
        <input 
          v-model="userInfo.nickname"
          type="nickname"
          class="nickname-input"
          placeholder="请输入昵称"
          maxlength="20"
          @input="onNicknameInput"
          @blur="onNicknameBlur"
        />
        <text class="input-tip">输入时键盘上方会显示微信昵称建议</text>
      </view>
    </view>

    <!-- 调试信息 -->
    <view class="debug-section" v-if="true">
      <text class="debug-title">调试信息：</text>
      <text class="debug-text">头像URL: {{ userInfo.avatarUrl || '未设置' }}</text>
      <text class="debug-text">昵称: {{ userInfo.nickname || '未设置' }}</text>
      <text class="debug-text">是否可保存: {{ canSave ? '是' : '否' }}</text>
    </view>

    <!-- 操作按钮 -->
    <view class="actions">
      <button 
        class="save-btn"
        :disabled="!canSave"
        @click="saveUserInfo"
      >
        保存信息
      </button>
      <button 
        class="skip-btn"
        @click="skipAuth"
      >
        暂时跳过
      </button>
    </view>

    <!-- 权限说明 -->
    <view class="permission-note">
      <text class="note-text">
        • 头像信息将用于个人资料展示
        • 昵称将作为您在社区中的显示名称
        • 您可以随时在设置中修改这些信息
      </text>
    </view>
  </view>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { userApi } from '@/utils/api.js'
import { setUserInfo, getUserInfo } from '@/utils/auth.js'

// 响应式数据
const userInfo = ref({
  avatarUrl: '',
  nickname: ''
})

// 计算属性
const canSave = computed(() => {
  return userInfo.value.nickname.trim().length > 0
})

// 生命周期
onMounted(() => {
  // 加载已有的用户信息
  const existingUser = getUserInfo()
  if (existingUser) {
    userInfo.value.nickname = existingUser.nickname || ''
    userInfo.value.avatarUrl = existingUser.avatarUrl || ''
  }
})

// 头像选择回调
const onChooseAvatar = (e) => {
  console.log('用户选择头像:', e.detail)
  if (e.detail.avatarUrl) {
    userInfo.value.avatarUrl = e.detail.avatarUrl
    console.log('头像URL设置为:', userInfo.value.avatarUrl)
    uni.showToast({
      title: '头像选择成功',
      icon: 'success',
      duration: 1500
    })
  } else {
    console.error('头像选择失败，没有获取到avatarUrl')
    uni.showToast({
      title: '头像选择失败',
      icon: 'error'
    })
  }
}

// 昵称输入处理
const onNicknameInput = (e) => {
  userInfo.value.nickname = e.detail.value
  console.log('昵称更新为:', userInfo.value.nickname)
}

// 昵称输入失去焦点处理
const onNicknameBlur = () => {
  console.log('昵称失去焦点:', userInfo.value.nickname)
}

// 保存用户信息
const saveUserInfo = async () => {
  if (!canSave.value) {
    uni.showToast({
      title: '请输入昵称',
      icon: 'none'
    })
    return
  }

  try {
    uni.showLoading({
      title: '保存中...'
    })

    console.log('准备保存用户信息:', userInfo.value)

    // 处理头像上传
    let finalAvatarUrl = userInfo.value.avatarUrl
    if (finalAvatarUrl && finalAvatarUrl.includes('wxfile://')) {
      console.log('检测到微信临时头像，准备上传:', finalAvatarUrl)
      
      try {
        // 微信临时头像需要上传到服务器
        const uploadResponse = await userApi.saveWxAvatar({
          tempUrl: finalAvatarUrl
        })
        
        if (uploadResponse.data.code === 0) {
          finalAvatarUrl = uploadResponse.data.data.url
          console.log('头像上传成功，新URL:', finalAvatarUrl)
        } else {
          console.error('头像上传失败:', uploadResponse.data.msg)
          // 如果上传失败，仍然使用临时URL
        }
      } catch (uploadError) {
        console.error('头像上传异常:', uploadError)
        // 上传失败不影响昵称保存
      }
    }

    // 更新用户信息
    console.log('调用API更新用户信息:', {
      nickname: userInfo.value.nickname.trim(),
      avatarUrl: finalAvatarUrl
    })
    
    const updateResponse = await userApi.updateUserInfo({
      nickname: userInfo.value.nickname.trim(),
      avatarUrl: finalAvatarUrl
    })

    if (updateResponse.data.code !== 0) {
      throw new Error(updateResponse.data.msg || '更新失败')
    }

    // 更新本地存储
    const currentUser = getUserInfo()
    const updatedUser = {
      ...currentUser,
      nickname: userInfo.value.nickname.trim(),
      avatarUrl: finalAvatarUrl
    }
    setUserInfo(updatedUser)
    console.log('本地用户信息更新完成:', updatedUser)

    uni.hideLoading()
    uni.showToast({
      title: '保存成功',
      icon: 'success'
    })

    // 返回个人中心
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)

  } catch (error) {
    console.error('保存用户信息失败:', error)
    uni.hideLoading()
    uni.showToast({
      title: error.message || '保存失败，请重试',
      icon: 'none'
    })
  }
}

// 跳过授权
const skipAuth = () => {
  uni.showModal({
    title: '确认跳过',
    content: '跳过后可在个人中心随时完善信息',
    success: (res) => {
      if (res.confirm) {
        uni.switchTab({
          url: '/pages/profile/index'
        })
      }
    }
  })
}
</script>

<style scoped>
.auth-container {
  min-height: 100vh;
  padding: 48rpx;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.header {
  text-align: center;
  margin-bottom: 80rpx;
}

.title {
  display: block;
  font-size: 48rpx;
  font-weight: bold;
  margin-bottom: 16rpx;
}

.desc {
  font-size: 28rpx;
  opacity: 0.8;
  line-height: 1.5;
}

.avatar-section, .nickname-section {
  margin-bottom: 60rpx;
}

.section-title {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 32rpx;
}

.avatar-container {
  display: flex;
  justify-content: center;
}

.avatar-btn {
  width: 160rpx;
  height: 160rpx;
  border: none;
  background: none;
  padding: 0;
  border-radius: 50%;
}

.avatar-image {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  border: 4rpx solid rgba(255, 255, 255, 0.3);
}

.avatar-placeholder {
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.2);
  border: 4rpx dashed rgba(255, 255, 255, 0.5);
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.camera-icon {
  font-size: 40rpx;
  margin-bottom: 8rpx;
}

.placeholder-text {
  font-size: 20rpx;
  text-align: center;
}

.nickname-container {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.nickname-input {
  height: 88rpx;
  background: rgba(255, 255, 255, 0.2);
  border: 2rpx solid rgba(255, 255, 255, 0.3);
  border-radius: 16rpx;
  padding: 0 32rpx;
  font-size: 32rpx;
  color: white;
}

.nickname-input::placeholder {
  color: rgba(255, 255, 255, 0.6);
}

.input-tip {
  font-size: 24rpx;
  opacity: 0.7;
  text-align: center;
}

.hint-text {
  font-size: 24rpx;
  opacity: 0.7;
  text-align: center;
}

.debug-section {
  margin-bottom: 60rpx;
}

.debug-title {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 32rpx;
}

.debug-text {
  font-size: 24rpx;
  line-height: 1.6;
  opacity: 0.8;
}

.actions {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
  margin-bottom: 60rpx;
}

.save-btn, .skip-btn {
  height: 88rpx;
  border-radius: 16rpx;
  font-size: 32rpx;
  font-weight: 600;
  border: none;
}

.save-btn {
  background: white;
  color: #667eea;
}

.save-btn:disabled {
  background: rgba(255, 255, 255, 0.3);
  color: rgba(255, 255, 255, 0.6);
}

.skip-btn {
  background: transparent;
  color: white;
  border: 2rpx solid rgba(255, 255, 255, 0.5);
}

.permission-note {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 16rpx;
  padding: 32rpx;
}

.note-text {
  font-size: 24rpx;
  line-height: 1.6;
  opacity: 0.8;
}
</style> 