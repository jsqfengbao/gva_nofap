<template>
  <view class="nf-tabbar">
    <view 
      v-for="(item, index) in tabs" 
      :key="index"
      :class="getTabClasses(index)"
      @click="handleTabClick(index)"
    >
      <view class="nf-tabbar__icon">
        <text class="nf-icon">{{ item.icon }}</text>
      </view>
      <text class="nf-tabbar__text">{{ item.text }}</text>
    </view>
  </view>
</template>

<script>
export default {
  name: 'NfTabBar',
  props: {
    current: {
      type: String,
      default: 'home'
    }
  },
  data() {
    return {
      tabs: [
        {
          icon: '🏠',
          text: '主页',
          name: 'home',
          pagePath: '/pages/index/index'
        },
        {
          icon: '📊',
          text: '进度',
          name: 'progress',
          pagePath: '/pages/progress/index'
        },
        {
          icon: '👥',
          text: '社区',
          name: 'community',
          pagePath: '/pages/community/index'
        },
        {
          icon: '📚',
          text: '学习',
          name: 'learning',
          pagePath: '/pages/learning/index'
        },
        {
          icon: '👤',
          text: '我的',
          name: 'profile',
          pagePath: '/pages/profile/index'
        }
      ]
    }
  },
  computed: {
    activeIndex() {
      return this.tabs.findIndex(tab => tab.name === this.current)
    }
  },
  methods: {
    getTabClasses(index) {
      const classes = ['nf-tabbar__item']
      
      if (index === this.activeIndex) {
        classes.push('nf-tabbar__item--active')
      }
      
      return classes.join(' ')
    },
    handleTabClick(index) {
      if (index !== this.activeIndex) {
        // 页面跳转
        const tab = this.tabs[index]
        if (tab.pagePath) {
          uni.switchTab({
            url: tab.pagePath
          })
        }
      }
    }
  }
}
</script>

<style scoped>
/* 底部导航栏 */
.nf-tabbar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  background-color: #FFFFFF;
  border-top: 1px solid #E5E7EB;
  padding: 8px 0;
  padding-bottom: calc(8px + env(safe-area-inset-bottom));
  z-index: 1000;
}

/* 导航项 */
.nf-tabbar__item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 6px 4px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.nf-tabbar__item:active {
  transform: scale(0.95);
}

/* 导航图标 */
.nf-tabbar__icon {
  margin-bottom: 4px;
  font-size: 20px;
  color: #9CA3AF;
  transition: color 0.2s ease;
}

/* 导航文字 */
.nf-tabbar__text {
  font-size: 10px;
  color: #9CA3AF;
  font-weight: 500;
  transition: color 0.2s ease;
}

/* 激活状态 */
.nf-tabbar__item--active .nf-tabbar__icon {
  color: #34D399;
}

.nf-tabbar__item--active .nf-tabbar__text {
  color: #34D399;
}

/* 图标样式 */
.nf-icon {
  font-size: 18px;
}

/* 适配安全区域 */
@supports (padding: max(0px)) {
  .nf-tabbar {
    padding-bottom: max(8px, env(safe-area-inset-bottom));
  }
}
</style> 