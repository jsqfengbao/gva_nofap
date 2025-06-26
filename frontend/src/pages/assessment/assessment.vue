<template>
  <view class="assessment-page">
    <!-- 顶部导航 -->
    <NfNavBar 
      title="色隐指数评估"
      :show-back="true"
      @back="goBack"
    />
    
    <view class="page-content">
      <!-- 评估介绍卡片 -->
      <NfCard type="highlight" class="intro-card">
        <view class="intro-header">
          <view class="intro-icon">
            <text class="nf-icon fa-brain"></text>
          </view>
          <text class="intro-title">专业心理评估</text>
          <text class="intro-subtitle">了解您的当前状况，制定个性化康复方案</text>
        </view>
      </NfCard>
      
      <!-- 评估说明 -->
      <NfCard title="评估说明" icon="fa-info-circle">
        <view class="description-content">
          <view class="feature-item">
            <text class="nf-icon fa-check-circle feature-icon"></text>
            <view class="feature-text">
              <text class="feature-title">科学专业</text>
              <text class="feature-desc">基于国际认可的心理学评估标准</text>
            </view>
          </view>
          
          <view class="feature-item">
            <text class="nf-icon fa-shield-alt feature-icon"></text>
            <view class="feature-text">
              <text class="feature-title">隐私保护</text>
              <text class="feature-desc">所有信息严格保密，仅用于个人评估</text>
            </view>
          </view>
          
          <view class="feature-item">
            <text class="nf-icon fa-chart-line feature-icon"></text>
            <view class="feature-text">
              <text class="feature-title">个性化结果</text>
              <text class="feature-desc">获得专属的评估报告和康复建议</text>
            </view>
          </view>
          
          <view class="feature-item">
            <text class="nf-icon fa-clock feature-icon"></text>
            <view class="feature-text">
              <text class="feature-title">快速便捷</text>
              <text class="feature-desc">仅需5-10分钟完成全部评估</text>
            </view>
          </view>
        </view>
      </NfCard>
      
      <!-- 评估内容 -->
      <NfCard title="评估内容" icon="fa-list-check">
        <view class="assessment-categories">
          <view class="category-item">
            <text class="category-icon">🕐</text>
            <text class="category-name">行为频率评估</text>
          </view>
          
          <view class="category-item">
            <text class="category-icon">🎯</text>
            <text class="category-name">自控能力评估</text>
          </view>
          
          <view class="category-item">
            <text class="category-icon">💼</text>
            <text class="category-name">生活影响评估</text>
          </view>
          
          <view class="category-item">
            <text class="category-icon">🧠</text>
            <text class="category-name">心理状态评估</text>
          </view>
          
          <view class="category-item">
            <text class="category-icon">👥</text>
            <text class="category-name">社交关系评估</text>
          </view>
          
          <view class="category-item">
            <text class="category-icon">❤️</text>
            <text class="category-name">身体健康评估</text>
          </view>
          
          <view class="category-item">
            <text class="category-icon">💡</text>
            <text class="category-name">认知功能评估</text>
          </view>
        </view>
      </NfCard>
      
      <!-- 操作按钮 -->
      <view class="action-section">
        <NfButton 
          type="primary"
          size="large"
          label="开始评估"
          full-width
          icon-left="fa-play"
          @click="startAssessment"
        />
        
        <NfButton 
          v-if="hasHistory"
          type="secondary"
          size="medium"
          label="查看历史评估"
          full-width
          icon-left="fa-history"
          @click="viewHistory"
        />
      </view>
      
      <!-- 温馨提示 -->
      <view class="notice-section">
        <NfCard>
          <view class="notice-content">
            <view class="notice-header">
              <text class="nf-icon fa-lightbulb notice-icon"></text>
              <text class="notice-title">温馨提示</text>
            </view>
            <view class="notice-list">
              <text class="notice-item">• 请在安静的环境中进行评估，确保能够专心思考</text>
              <text class="notice-item">• 诚实回答每个问题，这将帮助您获得更准确的结果</text>
              <text class="notice-item">• 评估过程中可以随时暂停，您的进度会被保存</text>
              <text class="notice-item">• 评估结果仅供参考，如需专业帮助请咨询心理健康专家</text>
            </view>
          </view>
        </NfCard>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'AssessmentPage',
  data() {
    return {
      hasHistory: false
    }
  },
  methods: {
    // 返回
    goBack() {
      uni.navigateBack()
    },
    
    // 开始评估
    startAssessment() {
      // 显示确认对话框
      uni.showModal({
        title: '开始评估',
        content: '评估大约需要5-10分钟时间，请确保您有足够的时间完成。准备好开始了吗？',
        confirmText: '开始评估',
        cancelText: '稍后再说',
        success: (res) => {
          if (res.confirm) {
            // 跳转到问卷页面
            uni.navigateTo({
              url: '/pages/assessment/questionnaire'
            })
          }
        }
      })
    },
    
    // 查看历史评估
    viewHistory() {
      uni.showToast({
        title: '功能开发中...',
        icon: 'none'
      })
    },
    
    // 检查是否有历史评估记录
    checkHistory() {
      try {
        const history = uni.getStorageSync('assessmentHistory') || []
        this.hasHistory = history.length > 0
      } catch (error) {
        console.error('检查历史记录失败:', error)
        this.hasHistory = false
      }
    }
  },
  
  onLoad() {
    this.checkHistory()
  },
  
  onShow() {
    // 页面显示时重新检查历史记录
    this.checkHistory()
  }
}
</script>

<style scoped>
.page-content {
  padding: 80px 16px 16px;
  min-height: 100vh;
  background-color: var(--background);
}

/* 介绍卡片 */
.intro-card {
  margin-bottom: 24px;
}

.intro-header {
  text-align: center;
  color: white;
}

.intro-icon {
  width: 80px;
  height: 80px;
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
}

.intro-icon .nf-icon {
  font-size: 36px;
  color: white;
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.intro-title {
  font-size: 24px;
  font-weight: 700;
  display: block;
  margin-bottom: 8px;
  color: white;
}

.intro-subtitle {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.9);
  line-height: 1.5;
}

/* 功能特点 */
.description-content {
  /* 由全局样式处理 */
}

.feature-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 16px;
}

.feature-item:last-child {
  margin-bottom: 0;
}

.feature-icon {
  font-size: 20px;
  color: var(--primary);
  margin-right: 12px;
  margin-top: 2px;
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.feature-text {
  flex: 1;
}

.feature-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  display: block;
  margin-bottom: 4px;
}

.feature-desc {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.4;
}

/* 评估分类 */
.assessment-categories {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.category-item {
  display: flex;
  align-items: center;
  padding: 12px;
  background-color: var(--surface);
  border-radius: 8px;
  border: 1px solid #e5e7eb;
}

.category-icon {
  font-size: 20px;
  margin-right: 8px;
}

.category-name {
  font-size: 14px;
  color: var(--text-primary);
  font-weight: 500;
}

/* 操作区域 */
.action-section {
  margin-top: 32px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

/* 温馨提示 */
.notice-section {
  margin-top: 24px;
}

.notice-content {
  padding: 8px 0;
}

.notice-header {
  display: flex;
  align-items: center;
  margin-bottom: 16px;
}

.notice-icon {
  font-size: 18px;
  color: var(--accent);
  margin-right: 8px;
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.notice-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
}

.notice-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.notice-item {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.5;
}

/* FontAwesome图标 */
.nf-icon {
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}
</style> 