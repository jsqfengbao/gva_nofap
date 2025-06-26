<template>
  <view class="post-page">
    <!-- 状态栏 -->
    <view class="status-bar">
      <text>{{ currentTime }}</text>
      <view class="status-icons">
        <text>📶</text>
        <text>📶</text>
        <view class="battery">
          <view class="battery-level"></view>
        </view>
      </view>
    </view>

    <!-- 头部 -->
    <view class="header">
      <view class="header-main">
        <view class="back-btn" @tap="goBack">
          <text class="icon">←</text>
        </view>
        <text class="title">发布帖子</text>
        <view class="publish-btn" :class="{ disabled: !canPublish }" @tap="publishPost">
          <text>发布</text>
        </view>
      </view>
    </view>

    <!-- 表单内容 -->
    <scroll-view scroll-y="true" class="form-container">
      <!-- 分类选择 -->
      <view class="form-section">
        <view class="section-title">
          <text>选择分类</text>
          <text class="required">*</text>
        </view>
        <view class="category-options">
          <view class="category-option" 
                v-for="category in categories" 
                :key="category.id"
                :class="{ active: selectedCategory === category.id }"
                @tap="selectCategory(category.id)">
            <text class="category-icon">{{ category.icon }}</text>
            <text class="category-name">{{ category.name }}</text>
          </view>
        </view>
      </view>

      <!-- 标题输入 -->
      <view class="form-section">
        <view class="section-title">
          <text>标题</text>
          <text class="required">*</text>
        </view>
        <view class="input-container">
          <input
            class="title-input"
            v-model="postTitle"
            placeholder="请输入帖子标题..."
            maxlength="100"
            @input="onTitleInput"
          />
          <text class="char-count">{{ postTitle.length }}/100</text>
        </view>
      </view>

      <!-- 内容输入 -->
      <view class="form-section">
        <view class="section-title">
          <text>内容</text>
          <text class="required">*</text>
        </view>
        <view class="textarea-container">
          <textarea
            class="content-textarea"
            v-model="postContent"
            placeholder="分享你的经历、感受或想法..."
            maxlength="2000"
            auto-height
            @input="onContentInput"
          />
          <text class="char-count">{{ postContent.length }}/2000</text>
        </view>
      </view>

      <!-- 匿名选项 -->
      <view class="form-section">
        <view class="anonymous-option" @tap="toggleAnonymous">
          <view class="option-left">
            <text class="option-icon">🔒</text>
            <view class="option-info">
              <text class="option-title">匿名发布</text>
              <text class="option-desc">其他用户将看不到你的真实昵称</text>
            </view>
          </view>
          <view class="toggle-switch" :class="{ active: isAnonymous }">
            <view class="toggle-button"></view>
          </view>
        </view>
      </view>
    </scroll-view>

    <!-- 加载遮罩 -->
    <view class="loading-mask" v-if="isPublishing">
      <view class="loading-content">
        <view class="spinner"></view>
        <text>发布中...</text>
      </view>
    </view>
  </view>
</template>

<script>
import { ref, computed, onMounted } from 'vue'

export default {
  name: 'CommunityPost',
  setup() {
    const currentTime = ref('9:41')
    const selectedCategory = ref(null)
    const postTitle = ref('')
    const postContent = ref('')
    const isAnonymous = ref(false)
    const isPublishing = ref(false)
    
    const categories = ref([
      { id: 1, name: '经验分享', icon: '💡' },
      { id: 2, name: '求助求鼓励', icon: '🤝' },
      { id: 3, name: '日常打卡', icon: '📅' },
      { id: 4, name: '成功故事', icon: '🎉' }
    ])

    const canPublish = computed(() => {
      return selectedCategory.value && 
             postTitle.value.trim().length > 0 && 
             postContent.value.trim().length > 0 &&
             !isPublishing.value
    })

    const updateTime = () => {
      const now = new Date()
      currentTime.value = `${now.getHours()}:${now.getMinutes().toString().padStart(2, '0')}`
    }

    const selectCategory = (categoryId) => {
      selectedCategory.value = categoryId
    }

    const toggleAnonymous = () => {
      isAnonymous.value = !isAnonymous.value
    }

    const onTitleInput = (e) => {
      postTitle.value = e.detail.value
    }

    const onContentInput = (e) => {
      postContent.value = e.detail.value
    }

    const goBack = () => {
      uni.navigateBack()
    }

    const publishPost = async () => {
      if (!canPublish.value) return
      
      isPublishing.value = true
      
      try {
        await new Promise(resolve => setTimeout(resolve, 2000))
        uni.showToast({ title: '发布成功', icon: 'success' })
        setTimeout(() => uni.navigateBack(), 1500)
      } catch (error) {
        uni.showToast({ title: '发布失败', icon: 'none' })
      } finally {
        isPublishing.value = false
      }
    }

    onMounted(() => {
      updateTime()
      setInterval(updateTime, 60000)
    })

    return {
      currentTime, selectedCategory, postTitle, postContent, isAnonymous, isPublishing,
      categories, canPublish, selectCategory, toggleAnonymous, onTitleInput, onContentInput,
      goBack, publishPost
    }
  }
}
</script>

<style scoped>
.post-page {
  min-height: 100vh;
  background: #F8FAFC;
}

.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
  font-size: 12px;
  color: #1F2937;
  background: #FFFFFF;
}

.status-icons {
  display: flex;
  align-items: center;
  gap: 4px;
}

.battery {
  width: 24px;
  height: 12px;
  border: 1px solid #1F2937;
  border-radius: 2px;
  position: relative;
}

.battery-level {
  width: 16px;
  height: 6px;
  background: #10B981;
  border-radius: 1px;
  margin: 2px;
}

.header {
  background: #FFFFFF;
  border-bottom: 1px solid #E5E7EB;
}

.header-main {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
}

.back-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: #F3F4F6;
}

.back-btn .icon {
  font-size: 18px;
  color: #6B7280;
}

.title {
  font-size: 18px;
  font-weight: 600;
  color: #1F2937;
}

.publish-btn {
  padding: 8px 16px;
  background: #22D3AA;
  border-radius: 8px;
  color: #FFFFFF;
  font-size: 14px;
  font-weight: 600;
}

.publish-btn.disabled {
  background: #D1D5DB;
  color: #9CA3AF;
}

.form-container {
  flex: 1;
  height: calc(100vh - 120px);
}

.form-section {
  background: #FFFFFF;
  margin-bottom: 12px;
  padding: 16px;
}

.section-title {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
  font-size: 16px;
  font-weight: 600;
  color: #1F2937;
}

.required {
  color: #EF4444;
  margin-left: 4px;
}

.category-options {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.category-option {
  flex: 1;
  min-width: 45%;
  padding: 12px;
  border: 2px solid #E5E7EB;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  transition: all 0.3s;
}

.category-option.active {
  border-color: #22D3AA;
  background: rgba(34, 211, 170, 0.05);
}

.category-icon {
  font-size: 24px;
}

.category-name {
  font-size: 14px;
  font-weight: 500;
  color: #374151;
}

.input-container {
  position: relative;
}

.title-input {
  width: 100%;
  padding: 12px;
  border: 1px solid #E5E7EB;
  border-radius: 8px;
  font-size: 16px;
  background: #FFFFFF;
}

.char-count {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 12px;
  color: #9CA3AF;
}

.textarea-container {
  position: relative;
}

.content-textarea {
  width: 100%;
  min-height: 120px;
  padding: 12px;
  border: 1px solid #E5E7EB;
  border-radius: 8px;
  font-size: 16px;
  background: #FFFFFF;
  line-height: 1.5;
}

.textarea-container .char-count {
  position: absolute;
  right: 12px;
  bottom: 12px;
  font-size: 12px;
  color: #9CA3AF;
}

.anonymous-option {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 0;
}

.option-left {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.option-icon {
  font-size: 20px;
}

.option-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.option-title {
  font-size: 16px;
  font-weight: 500;
  color: #1F2937;
}

.option-desc {
  font-size: 12px;
  color: #6B7280;
}

.toggle-switch {
  width: 48px;
  height: 28px;
  border-radius: 14px;
  background: #E5E7EB;
  position: relative;
  transition: background-color 0.3s;
}

.toggle-switch.active {
  background: #22D3AA;
}

.toggle-button {
  width: 24px;
  height: 24px;
  border-radius: 12px;
  background: #FFFFFF;
  position: absolute;
  top: 2px;
  left: 2px;
  transition: transform 0.3s;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.toggle-switch.active .toggle-button {
  transform: translateX(20px);
}

.loading-mask {
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

.loading-content {
  background: #FFFFFF;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.spinner {
  width: 32px;
  height: 32px;
  border: 3px solid #E5E7EB;
  border-top: 3px solid #22D3AA;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.loading-content text {
  font-size: 14px;
  color: #6B7280;
}
</style>
