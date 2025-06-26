<template>
  <view class="profile-setup">
    <view class="header">
      <text class="title">完善个人资料</text>
      <text class="desc">请设置您的头像和昵称</text>
    </view>

    <view class="form-section">
      <!-- 头像选择 -->
      <view class="form-item">
        <text class="label">头像</text>
        <button 
          class="avatar-btn"
          open-type="chooseAvatar"
          @chooseavatar="onChooseAvatar"
        >
          <view class="avatar-container">
            <image 
              v-if="userInfo.avatarUrl" 
              :src="userInfo.avatarUrl" 
              class="avatar-image"
            />
            <view v-else class="avatar-placeholder">
              <text class="avatar-text">点击选择头像</text>
            </view>
          </view>
        </button>
      </view>

      <!-- 昵称输入 -->
      <view class="form-item">
        <text class="label">昵称</text>
        <input 
          v-model="userInfo.nickname"
          type="nickname" 
          class="nickname-input"
          placeholder="请输入昵称"
          maxlength="20"
        />
      </view>
    </view>

    <view class="actions">
      <button 
        class="save-btn"
        :disabled="!canSave"
        @click="saveProfile"
      >
        保存资料
      </button>
      <button 
        class="skip-btn"
        @click="skipSetup"
      >
        暂时跳过
      </button>
    </view>
  </view>
</template>

<script setup>
import { ref, computed } from 'vue'

const emit = defineEmits(['save', 'skip'])

const userInfo = ref({
  avatarUrl: '',
  nickname: ''
})

const canSave = computed(() => {
  return userInfo.value.nickname.trim().length > 0
})

// 选择头像
const onChooseAvatar = (e) => {
  console.log('选择头像:', e.detail)
  if (e.detail.avatarUrl) {
    userInfo.value.avatarUrl = e.detail.avatarUrl
  }
}

// 保存资料
const saveProfile = () => {
  if (!canSave.value) {
    uni.showToast({
      title: '请输入昵称',
      icon: 'none'
    })
    return
  }

  emit('save', {
    avatarUrl: userInfo.value.avatarUrl,
    nickname: userInfo.value.nickname.trim()
  })
}

// 跳过设置
const skipSetup = () => {
  emit('skip')
}
</script>

<style scoped>
.profile-setup {
  padding: 48rpx;
  min-height: 100vh;
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
}

.form-section {
  margin-bottom: 80rpx;
}

.form-item {
  margin-bottom: 48rpx;
}

.label {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 24rpx;
}

.avatar-btn {
  background: none;
  border: none;
  padding: 0;
}

.avatar-container {
  width: 160rpx;
  height: 160rpx;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.2);
  border: 4rpx dashed rgba(255, 255, 255, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
  transition: all 0.3s ease;
}

.avatar-container:active {
  transform: scale(0.95);
}

.avatar-image {
  width: 100%;
  height: 100%;
  border-radius: 50%;
}

.avatar-placeholder {
  text-align: center;
}

.avatar-text {
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
}

.nickname-input {
  width: 100%;
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

.actions {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
}

.save-btn, .skip-btn {
  height: 88rpx;
  border-radius: 16rpx;
  font-size: 32rpx;
  font-weight: 600;
}

.save-btn {
  background: white;
  color: #667eea;
  border: none;
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
</style> 