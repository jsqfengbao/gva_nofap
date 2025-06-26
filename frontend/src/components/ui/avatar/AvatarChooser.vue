<template>
  <button 
    class="avatar-chooser"
    open-type="chooseAvatar"
    @chooseavatar="onChooseAvatar"
  >
    <view class="avatar-container">
      <image 
        v-if="avatarUrl" 
        :src="avatarUrl" 
        class="avatar-image"
      />
      <view v-else class="avatar-placeholder">
        <text class="avatar-emoji">🌱</text>
      </view>
      <view class="avatar-mask">
        <text class="camera-icon">📷</text>
      </view>
    </view>
  </button>
</template>

<script>
export default {
  name: 'AvatarChooser',
  props: {
    avatarUrl: {
      type: String,
      default: ''
    }
  },
  emits: ['choose'],
  methods: {
    onChooseAvatar(e) {
      console.log('选择头像:', e.detail)
      if (e.detail.avatarUrl) {
        this.$emit('choose', e.detail.avatarUrl)
      }
    }
  }
}
</script>

<style scoped>
.avatar-chooser {
  border: none;
  background: none;
  padding: 0;
  position: relative;
}

.avatar-container {
  width: 128rpx;
  height: 128rpx;
  border-radius: 50%;
  overflow: hidden;
  position: relative;
  background: rgba(255, 255, 255, 0.2);
}

.avatar-image {
  width: 100%;
  height: 100%;
  border-radius: 50%;
}

.avatar-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-emoji {
  font-size: 48rpx;
}

.avatar-mask {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.avatar-chooser:hover .avatar-mask {
  opacity: 1;
}

.camera-icon {
  font-size: 32rpx;
  color: white;
}
</style> 