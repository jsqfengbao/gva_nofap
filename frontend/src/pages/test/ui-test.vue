<template>
  <view class="page">
    <!-- 顶部导航 -->
    <NfNavBar 
      title="UI组件测试"
      :show-back="true"
      right-icon="fa-cog"
      @back="handleBack"
      @right-click="handleSettings"
    />
    
    <view class="page-content space-y-6">
      <!-- 按钮组件测试 -->
      <NfCard title="按钮组件" icon="fa-mouse-pointer">
        <view class="space-y-4">
          <view class="space-y-2">
            <text class="font-semibold">Primary按钮：</text>
            <NfButton 
              type="primary" 
              label="主要按钮"
              @click="showToast('Primary按钮点击')"
            />
            <NfButton 
              type="primary" 
              size="large"
              label="大按钮"
              icon-left="fa-plus"
              full-width
              @click="showToast('大按钮点击')"
            />
          </view>
          
          <view class="space-y-2">
            <text class="font-semibold">Secondary按钮：</text>
            <NfButton 
              type="secondary"
              label="次要按钮"
              icon-right="fa-arrow-right"
              @click="showToast('Secondary按钮点击')"
            />
          </view>
          
          <view class="space-y-2">
            <text class="font-semibold">图标按钮：</text>
            <view class="flex space-x-3">
              <NfButton 
                type="icon"
                icon-left="fa-heart"
                @click="showToast('心形图标点击')"
              />
              <NfButton 
                type="icon"
                size="large"
                icon-left="fa-share"
                @click="showToast('分享图标点击')"
              />
              <NfButton 
                type="icon"
                icon-left="fa-bookmark"
                disabled
              />
            </view>
          </view>
        </view>
      </NfCard>
      
      <!-- 卡片组件测试 -->
      <NfCard title="卡片组件" icon="fa-square">
        <view class="space-y-4">
          <NfCard 
            type="basic"
            title="基础卡片"
            icon="fa-info-circle"
            content="这是一个基础卡片的示例内容。"
          />
          
          <NfCard 
            type="highlight"
            title="强调卡片"
            icon="fa-star"
            content="这是一个强调卡片，有特殊的边框样式。"
          />
          
          <NfCard 
            type="gradient"
            title="渐变卡片"
            icon="fa-crown"
            content="这是一个渐变背景的卡片，用于重要内容展示。"
          />
        </view>
      </NfCard>
      
      <!-- 表单组件测试 -->
      <NfCard title="表单组件" icon="fa-edit">
        <view class="space-y-4">
          <NfInput
            v-model="formData.username"
            label="用户名"
            placeholder="请输入用户名"
            icon-left="fa-user"
            required
          />
          
          <NfInput
            v-model="formData.password"
            type="password"
            label="密码"
            placeholder="请输入密码"
            icon-left="fa-lock"
            required
          />
          
          <NfInput
            v-model="formData.email"
            type="email"
            label="邮箱"
            placeholder="请输入邮箱地址"
            icon-left="fa-envelope"
            :error="emailError"
            error-message="请输入有效的邮箱地址"
            help-text="用于找回密码和接收通知"
          />
          
          <NfButton 
            type="primary"
            label="提交表单"
            full-width
            @click="submitForm"
          />
        </view>
      </NfCard>

      <!-- 社区功能测试 -->
      <view class="test-section">
        <text class="section-title">社区功能测试</text>
        
        <view class="test-item">
          <text class="test-label">社区首页</text>
          <nf-button 
            type="primary" 
            size="small"
            @click="navigateToPage('/pages/community/index')"
          >
            测试社区首页
          </nf-button>
        </view>

        <view class="test-item">
          <text class="test-label">发布帖子</text>
          <nf-button 
            type="secondary" 
            size="small"
            @click="navigateToPage('/pages/community/post')"
          >
            测试发帖页面
          </nf-button>
        </view>

        <view class="test-item">
          <text class="test-label">帖子详情</text>
          <nf-button 
            type="primary" 
            size="small"
            @click="navigateToPage('/pages/community/detail?id=1')"
          >
            测试详情页面
          </nf-button>
        </view>
      </view>
    </view>
    
    <!-- 底部导航 -->
    <NfTabBar 
      :active-index="activeTabIndex"
      @change="handleTabChange"
    />
  </view>
</template>

<script>
import NfTabBar from '@/components/ui/navigation/NfTabBar.vue'

export default {
  name: 'UITest',
  components: {
    NfTabBar
  },
  data() {
    return {
      activeTabIndex: 0,
      emailError: false,
      formData: {
        username: '',
        password: '',
        email: ''
      }
    }
  },
  methods: {
    showToast(message) {
      uni.showToast({
        title: message,
        duration: 2000
      })
    },
    handleBack() {
      uni.navigateBack()
    },
    handleSettings() {
      this.showToast('设置按钮点击')
    },
    handleTabChange(index) {
      this.activeTabIndex = index
      this.showToast(`切换到标签页 ${index}`)
    },
    submitForm() {
      // 简单的邮箱验证
      const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      this.emailError = this.formData.email && !emailPattern.test(this.formData.email)
      
      if (!this.emailError) {
        this.showToast('表单提交成功')
        console.log('表单数据:', this.formData)
      }
    },
    navigateToPage(url) {
      uni.navigateTo({
        url: url
      })
    }
  }
}
</script>

<style scoped>
.page-content {
  padding-top: 60px; /* 为顶部导航留空间 */
}
</style> 