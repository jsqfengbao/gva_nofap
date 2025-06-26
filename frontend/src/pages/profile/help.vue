<template>
  <view class="help-container">
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
      <text class="page-title">帮助中心</text>
      <view class="placeholder"></view>
    </view>

    <!-- Main Content -->
    <view class="main-content">
      <!-- Search Bar -->
      <view class="search-section">
        <view class="search-box">
          <text class="search-icon">🔍</text>
          <input 
            class="search-input" 
            placeholder="搜索帮助内容..."
            v-model="searchText"
            @input="onSearchInput"
          />
        </view>
      </view>

      <!-- Quick Actions -->
      <view class="quick-actions">
        <view class="action-item" @click="contactSupport">
          <view class="action-icon support-icon">💬</view>
          <text class="action-text">联系客服</text>
        </view>
        
        <view class="action-item" @click="reportBug">
          <view class="action-icon bug-icon">🐛</view>
          <text class="action-text">反馈问题</text>
        </view>
        
        <view class="action-item" @click="suggestFeature">
          <view class="action-icon feature-icon">💡</view>
          <text class="action-text">功能建议</text>
        </view>
      </view>

      <!-- FAQ Categories -->
      <view class="faq-section">
        <text class="section-title">常见问题</text>
        
        <view class="category-list">
          <view 
            v-for="category in faqCategories" 
            :key="category.id"
            class="category-item"
            @click="toggleCategory(category.id)"
          >
            <view class="category-header">
              <view class="category-icon" :class="category.iconClass">
                <text>{{ category.icon }}</text>
              </view>
              <text class="category-title">{{ category.title }}</text>
              <text class="expand-icon" :class="{ expanded: category.expanded }">
                {{ category.expanded ? '−' : '+' }}
              </text>
            </view>
            
            <view v-if="category.expanded" class="category-content">
              <view 
                v-for="faq in category.faqs" 
                :key="faq.id"
                class="faq-item"
                @click="showFaqDetail(faq)"
              >
                <text class="faq-question">{{ faq.question }}</text>
                <text class="faq-arrow">></text>
              </view>
            </view>
          </view>
        </view>
      </view>

      <!-- Contact Info -->
      <view class="contact-section">
        <text class="section-title">联系我们</text>
        
        <view class="contact-list">
          <view class="contact-item">
            <view class="contact-icon email-icon">📧</view>
            <view class="contact-info">
              <text class="contact-title">邮箱支持</text>
              <text class="contact-desc">support@nofap-helper.com</text>
            </view>
          </view>
          
          <view class="contact-item">
            <view class="contact-icon community-icon">👥</view>
            <view class="contact-info">
              <text class="contact-title">社区求助</text>
              <text class="contact-desc">在社区发帖寻求帮助</text>
            </view>
          </view>
          
          <view class="contact-item">
            <view class="contact-icon time-icon">⏰</view>
            <view class="contact-info">
              <text class="contact-title">服务时间</text>
              <text class="contact-desc">周一至周五 9:00-18:00</text>
            </view>
          </view>
        </view>
      </view>
    </view>

    <!-- FAQ Detail Modal -->
    <view v-if="showDetailModal" class="modal-overlay" @click="closeDetailModal">
      <view class="detail-modal" @click.stop>
        <view class="modal-header">
          <text class="modal-title">{{ selectedFaq?.question }}</text>
          <view class="close-btn" @click="closeDetailModal">×</view>
        </view>
        
        <view class="modal-content">
          <text class="faq-answer">{{ selectedFaq?.answer }}</text>
          
          <view v-if="selectedFaq?.steps" class="steps-section">
            <text class="steps-title">解决步骤：</text>
            <view class="steps-list">
              <view 
                v-for="(step, index) in selectedFaq.steps" 
                :key="index"
                class="step-item"
              >
                <text class="step-number">{{ index + 1 }}.</text>
                <text class="step-text">{{ step }}</text>
              </view>
            </view>
          </view>
        </view>

        <view class="modal-footer">
          <button class="modal-btn helpful" @click="markHelpful">
            👍 有帮助
          </button>
          <button class="modal-btn not-helpful" @click="markNotHelpful">
            👎 没帮助
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
const loadingText = ref('加载中...')
const searchText = ref('')
const showDetailModal = ref(false)
const selectedFaq = ref(null)

// FAQ分类数据
const faqCategories = ref([
  {
    id: 1,
    title: '账户与登录',
    icon: '👤',
    iconClass: 'account-icon',
    expanded: false,
    faqs: [
      {
        id: 1,
        question: '如何注册账户？',
        answer: '本应用使用微信登录，无需单独注册。首次使用时，点击"微信登录"按钮，授权后即可自动创建账户。',
        steps: [
          '打开应用，点击"微信登录"',
          '在微信中确认授权',
          '完成用户信息设置',
          '开始使用应用功能'
        ]
      },
      {
        id: 2,
        question: '忘记了登录密码怎么办？',
        answer: '本应用使用微信登录，无需记住密码。如果无法登录，请检查微信是否正常，或重新授权登录。'
      },
      {
        id: 3,
        question: '如何更换绑定的微信账号？',
        answer: '目前暂不支持更换绑定的微信账号。如需使用其他微信账号，请联系客服处理。'
      }
    ]
  },
  {
    id: 2,
    title: '打卡与记录',
    icon: '✅',
    iconClass: 'checkin-icon',
    expanded: false,
    faqs: [
      {
        id: 4,
        question: '如何进行每日打卡？',
        answer: '在首页或打卡页面点击"今日打卡"按钮，选择当前心情状态（1-5级），添加备注（可选），然后确认打卡。',
        steps: [
          '进入打卡页面',
          '选择心情等级（1-5）',
          '添加今日感想（可选）',
          '点击确认打卡'
        ]
      },
      {
        id: 5,
        question: '错过了打卡时间怎么办？',
        answer: '如果错过了当天打卡，连续天数会重置为0。建议设置打卡提醒，避免错过打卡时间。'
      },
      {
        id: 6,
        question: '可以补签之前的打卡记录吗？',
        answer: '为了保证记录的真实性，暂不支持补签功能。请坚持每日打卡，养成良好习惯。'
      }
    ]
  },
  {
    id: 3,
    title: '成就与等级',
    icon: '🏆',
    iconClass: 'achievement-icon',
    expanded: false,
    faqs: [
      {
        id: 7,
        question: '如何获得成就徽章？',
        answer: '通过完成特定条件可以获得成就徽章，如连续打卡、帮助他人、学习内容等。成就会自动解锁。'
      },
      {
        id: 8,
        question: '等级是如何计算的？',
        answer: '等级基于经验值计算，通过打卡、获得成就、社区互动等方式可以获得经验值。等级越高，解锁的功能越多。'
      },
      {
        id: 9,
        question: '成就徽章有什么用？',
        answer: '成就徽章是对您坚持努力的认可，同时也能在社区中展示您的成长历程，激励自己和他人。'
      }
    ]
  },
  {
    id: 4,
    title: '社区与互助',
    icon: '👥',
    iconClass: 'community-icon',
    expanded: false,
    faqs: [
      {
        id: 10,
        question: '如何在社区发帖？',
        answer: '进入社区页面，点击右下角的"+"按钮，选择帖子分类，填写标题和内容，选择是否匿名发布。'
      },
      {
        id: 11,
        question: '社区有哪些发帖规则？',
        answer: '请遵守社区规则：内容积极正面、禁止发布不当内容、尊重他人、分享真实经验。违规内容会被删除。'
      },
      {
        id: 12,
        question: '如何举报不当内容？',
        answer: '在帖子或评论右上角点击"..."菜单，选择"举报"，选择举报原因并提交。我们会及时处理。'
      }
    ]
  }
])

// 生命周期
onMounted(() => {
  // 初始化数据
})

// 方法
const goBack = () => {
  uni.navigateBack()
}

const onSearchInput = () => {
  // 实现搜索功能
  console.log('搜索内容:', searchText.value)
}

const contactSupport = () => {
  uni.showModal({
    title: '联系客服',
    content: '请发送邮件至 support@nofap-helper.com 或在社区发帖寻求帮助',
    showCancel: false,
    confirmText: '我知道了'
  })
}

const reportBug = () => {
  uni.showModal({
    title: '反馈问题',
    content: '请详细描述遇到的问题，并发送至 bug@nofap-helper.com',
    showCancel: false,
    confirmText: '我知道了'
  })
}

const suggestFeature = () => {
  uni.showModal({
    title: '功能建议',
    content: '欢迎分享您的想法！请发送建议至 feature@nofap-helper.com',
    showCancel: false,
    confirmText: '我知道了'
  })
}

const toggleCategory = (categoryId) => {
  const category = faqCategories.value.find(c => c.id === categoryId)
  if (category) {
    category.expanded = !category.expanded
  }
}

const showFaqDetail = (faq) => {
  selectedFaq.value = faq
  showDetailModal.value = true
}

const closeDetailModal = () => {
  showDetailModal.value = false
  selectedFaq.value = null
}

const markHelpful = () => {
  uni.showToast({
    title: '感谢反馈',
    icon: 'success'
  })
  closeDetailModal()
}

const markNotHelpful = () => {
  uni.showToast({
    title: '我们会改进',
    icon: 'success'
  })
  closeDetailModal()
}
</script>

<style scoped>
.help-container {
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
  padding-bottom: 160rpx;
}

/* Search Section */
.search-section {
  margin-bottom: 48rpx;
}

.search-box {
  background: white;
  border-radius: 24rpx;
  padding: 24rpx 32rpx;
  display: flex;
  align-items: center;
  gap: 16rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.search-icon {
  font-size: 32rpx;
  color: #6B7280;
}

.search-input {
  flex: 1;
  font-size: 28rpx;
  color: #1F2937;
}

/* Quick Actions */
.quick-actions {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24rpx;
  margin-bottom: 48rpx;
}

.action-item {
  background: white;
  border-radius: 24rpx;
  padding: 32rpx 24rpx;
  text-align: center;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.action-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 20rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 16rpx;
  font-size: 32rpx;
}

.support-icon {
  background: #DBEAFE;
}

.bug-icon {
  background: #FEE2E2;
}

.feature-icon {
  background: #FEF3C7;
}

.action-text {
  font-size: 24rpx;
  font-weight: 600;
  color: #1F2937;
}

/* FAQ Section */
.faq-section {
  margin-bottom: 48rpx;
}

.section-title {
  font-size: 36rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 32rpx;
}

.category-list {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
}

.category-item {
  background: white;
  border-radius: 24rpx;
  overflow: hidden;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.category-header {
  display: flex;
  align-items: center;
  gap: 24rpx;
  padding: 32rpx 48rpx;
}

.category-icon {
  width: 64rpx;
  height: 64rpx;
  border-radius: 16rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28rpx;
}

.account-icon {
  background: #DBEAFE;
}

.checkin-icon {
  background: #D1FAE5;
}

.achievement-icon {
  background: #FEF3C7;
}

.community-icon {
  background: #E0E7FF;
}

.category-title {
  flex: 1;
  font-size: 32rpx;
  font-weight: 600;
  color: #1F2937;
}

.expand-icon {
  font-size: 32rpx;
  color: #6B7280;
  transition: transform 0.3s ease;
}

.expand-icon.expanded {
  transform: rotate(180deg);
}

.category-content {
  border-top: 1rpx solid #F3F4F6;
}

.faq-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 24rpx 48rpx;
  border-bottom: 1rpx solid #F3F4F6;
}

.faq-item:last-child {
  border-bottom: none;
}

.faq-question {
  flex: 1;
  font-size: 28rpx;
  color: #1F2937;
}

.faq-arrow {
  font-size: 24rpx;
  color: #9CA3AF;
}

/* Contact Section */
.contact-section {
  margin-bottom: 48rpx;
}

.contact-list {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
}

.contact-item {
  background: white;
  border-radius: 24rpx;
  padding: 32rpx;
  display: flex;
  align-items: center;
  gap: 24rpx;
  box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
}

.contact-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 20rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
}

.email-icon {
  background: #DBEAFE;
}

.community-icon {
  background: #E0E7FF;
}

.time-icon {
  background: #FEF3C7;
}

.contact-info {
  flex: 1;
}

.contact-title {
  display: block;
  font-size: 28rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 8rpx;
}

.contact-desc {
  display: block;
  font-size: 24rpx;
  color: #6B7280;
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

.detail-modal {
  background: white;
  border-radius: 32rpx;
  width: 640rpx;
  max-height: 80vh;
  overflow: hidden;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 32rpx 48rpx;
  border-bottom: 1rpx solid #E5E7EB;
}

.modal-title {
  flex: 1;
  font-size: 32rpx;
  font-weight: 600;
  color: #1F2937;
  line-height: 1.4;
  margin-right: 16rpx;
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
  max-height: 60vh;
  overflow-y: auto;
}

.faq-answer {
  display: block;
  font-size: 28rpx;
  color: #1F2937;
  line-height: 1.6;
  margin-bottom: 24rpx;
}

.steps-section {
  margin-top: 24rpx;
}

.steps-title {
  display: block;
  font-size: 28rpx;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 16rpx;
}

.steps-list {
  display: flex;
  flex-direction: column;
  gap: 12rpx;
}

.step-item {
  display: flex;
  gap: 16rpx;
}

.step-number {
  font-size: 24rpx;
  font-weight: 600;
  color: #34D399;
  min-width: 32rpx;
}

.step-text {
  flex: 1;
  font-size: 24rpx;
  color: #6B7280;
  line-height: 1.5;
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

.modal-btn.helpful {
  background: #34D399;
  color: white;
}

.modal-btn.not-helpful {
  background: #F3F4F6;
  color: #6B7280;
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