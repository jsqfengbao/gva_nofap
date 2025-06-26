<template>
  <view class="nf-navbar">
    <!-- 状态栏占位 -->
    <view :style="{ height: statusBarHeight + 'px' }"></view>
    
    <!-- 导航栏内容 -->
    <view class="nf-navbar__content">
      <!-- 左侧内容 -->
      <view class="nf-navbar__left">
        <slot name="left">
          <view v-if="showBack" class="nf-navbar__back" @click="handleBack">
            <text class="nf-icon fa-arrow-left"></text>
          </view>
        </slot>
      </view>
      
      <!-- 中间标题 -->
      <view class="nf-navbar__center">
        <slot name="center">
          <text class="nf-navbar__title">{{ title }}</text>
        </slot>
      </view>
      
      <!-- 右侧内容 -->
      <view class="nf-navbar__right">
        <slot name="right">
          <view v-if="rightIcon" class="nf-navbar__right-btn" @click="handleRightClick">
            <text :class="`nf-icon ${rightIcon}`"></text>
          </view>
          <view v-else-if="rightText" class="nf-navbar__right-text" @click="handleRightClick">
            <text class="right-text">{{ rightText }}</text>
          </view>
        </slot>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'NfNavBar',
  props: {
    title: {
      type: String,
      default: ''
    },
    showBack: {
      type: Boolean,
      default: false
    },
    rightIcon: {
      type: String,
      default: ''
    },
    rightText: {
      type: String,
      default: ''
    },
    backgroundColor: {
      type: String,
      default: '#FFFFFF'
    },
    textColor: {
      type: String,
      default: '#1F2937'
    },
    statusBarHeight: {
      type: Number,
      default: 0
    }
  },
  emits: ['back', 'rightClick', 'statusBarHeight'],
  mounted() {
    // 获取状态栏高度
    if (this.statusBarHeight === 0) {
      this.getSystemInfo()
    }
  },
  methods: {
    getSystemInfo() {
      uni.getSystemInfo({
        success: (res) => {
          this.$emit('statusBarHeight', res.statusBarHeight || 20)
        }
      })
    },
    handleBack() {
      this.$emit('back')
      
      // 默认返回逻辑
      const pages = getCurrentPages()
      if (pages.length > 1) {
        uni.navigateBack()
      } else {
        uni.switchTab({
          url: '/pages/index/index'
        })
      }
    },
    handleRightClick() {
      this.$emit('rightClick')
    }
  }
}
</script>

<style scoped>
/* 引入FontAwesome图标 */
/* 注意：微信小程序不支持外部CSS导入，FontAwesome图标需要通过其他方式引入 */
/* @import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css'); */

/* 导航栏容器 */
.nf-navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background-color: v-bind(backgroundColor);
  z-index: 1001;
}

/* 导航栏内容 */
.nf-navbar__content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 44px;
  padding: 0 16px;
  border-bottom: 1px solid #E5E7EB;
}

/* 左侧区域 */
.nf-navbar__left {
  display: flex;
  align-items: center;
  min-width: 60px;
  justify-content: flex-start;
}

/* 返回按钮 */
.nf-navbar__back {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.nf-navbar__back:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.nf-navbar__back:active {
  transform: scale(0.95);
}

.nf-navbar__back .nf-icon {
  font-size: 16px;
  color: v-bind(textColor);
}

/* 中间区域 */
.nf-navbar__center {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 20px;
}

/* 标题 */
.nf-navbar__title {
  font-size: 17px;
  font-weight: 600;
  color: v-bind(textColor);
  text-align: center;
  line-height: 1;
}

/* 右侧区域 */
.nf-navbar__right {
  display: flex;
  align-items: center;
  min-width: 60px;
  justify-content: flex-end;
}

/* 右侧按钮 */
.nf-navbar__right-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.nf-navbar__right-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.nf-navbar__right-btn:active {
  transform: scale(0.95);
}

.nf-navbar__right-btn .nf-icon {
  font-size: 16px;
  color: v-bind(textColor);
}

/* 右侧文字按钮 */
.nf-navbar__right-text {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 6px 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.nf-navbar__right-text:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.nf-navbar__right-text:active {
  transform: scale(0.95);
}

.right-text {
  font-size: 14px;
  color: v-bind(textColor);
  font-weight: 500;
}

/* 图标样式 */
.nf-icon {
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}
</style> 