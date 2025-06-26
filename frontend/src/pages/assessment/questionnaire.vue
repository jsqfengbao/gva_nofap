<template>
  <view class="questionnaire-page">
    <!-- 状态栏 -->
    <view class="status-bar">
      <view class="flex justify-between items-center px-4 py-2 text-xs text-gray-800">
        <text class="font-medium">9:41</text>
        <view class="flex items-center space-x-1">
          <text class="nf-icon fa-signal text-xs"></text>
          <text class="nf-icon fa-wifi text-xs"></text>
          <view class="battery-indicator">
            <view class="battery-level"></view>
          </view>
        </view>
      </view>
    </view>

    <!-- 导航头部 -->
    <view class="nav-header bg-white shadow-sm border-b border-gray-100">
      <view class="flex items-center justify-between px-4 py-3">
        <button class="nav-btn" @click="handleBack">
          <text class="nf-icon fa-arrow-left text-gray-600 text-lg"></text>
        </button>
        <text class="nav-title">评估测试</text>
        <button class="nav-btn" @click="showHelp">
          <text class="nf-icon fa-question-circle text-gray-400 text-lg"></text>
        </button>
      </view>
      
      <!-- 进度条 -->
      <view class="progress-section px-4 pb-3">
        <view class="flex items-center justify-between mb-2">
          <text class="progress-text">问题 {{ currentQuestionIndex + 1 }} / {{ totalQuestions }}</text>
          <text class="progress-text">{{ Math.round(((currentQuestionIndex + 1) / totalQuestions) * 100) }}%</text>
        </view>
        <view class="progress-track">
          <view 
            class="progress-bar" 
            :style="{ width: ((currentQuestionIndex + 1) / totalQuestions) * 100 + '%' }"
          ></view>
        </view>
      </view>
    </view>

    <!-- 主要内容 -->
    <view class="main-content flex-1 px-6 py-6">
      <!-- 问题卡片 -->
      <view class="question-card bg-white rounded-3xl shadow-sm border border-gray-100 p-6 mb-6">
        <view class="question-header text-center mb-6">
          <view class="category-icon">
            <text class="nf-icon" :class="currentQuestion.icon"></text>
          </view>
          <text class="category-title">{{ currentQuestion.categoryName }}</text>
          <text class="category-desc">请根据实际情况诚实回答</text>
        </view>
        
        <view class="question-content mb-8">
          <text class="question-text">{{ currentQuestion.question }}</text>
          <text class="privacy-note">*此信息仅用于个人评估，不会被分享给任何第三方</text>
        </view>

        <!-- 答案选项 -->
        <view class="answer-options space-y-3">
          <button 
            v-for="(option, index) in currentQuestion.options" 
            :key="index"
            class="option-btn"
            :class="{ 'option-selected': selectedAnswer === option.value }"
            @click="selectAnswer(option.value)"
          >
            <view class="flex items-center">
              <view class="option-radio" :class="{ 'radio-selected': selectedAnswer === option.value }">
                <view v-if="selectedAnswer === option.value" class="radio-dot"></view>
              </view>
              <text class="option-text">{{ option.text }}</text>
            </view>
          </button>
        </view>
      </view>

      <!-- 操作按钮 -->
      <view class="action-buttons flex space-x-4">
        <button 
          class="action-btn secondary-btn flex-1"
          :class="{ 'btn-disabled': currentQuestionIndex === 0 }"
          @click="previousQuestion"
          :disabled="currentQuestionIndex === 0"
        >
          上一题
        </button>
        <button 
          class="action-btn primary-btn flex-1"
          :class="{ 'btn-disabled': selectedAnswer === null }"
          @click="nextQuestion"
          :disabled="selectedAnswer === null"
        >
          {{ isLastQuestion ? '完成评估' : '下一题' }}
        </button>
      </view>

      <!-- 帮助说明 -->
      <view class="help-section mt-6 p-4 bg-blue-50 rounded-xl border border-blue-100">
        <view class="flex items-start space-x-3">
          <text class="nf-icon fa-info-circle text-blue-500 mt-0_5"></text>
          <view class="help-content">
            <text class="help-title">评估说明</text>
            <text class="help-text">
              本评估基于国际认可的心理健康评估标准，结果仅供参考。如需专业帮助，建议咨询心理健康专家。
            </text>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { assessmentQuestions } from '@/data/assessment-questions.js'

export default {
  name: 'QuestionnairePage',
  data() {
    return {
      questions: assessmentQuestions,
      currentQuestionIndex: 0,
      answers: {},
      selectedAnswer: null
    }
  },
  computed: {
    currentQuestion() {
      return this.questions[this.currentQuestionIndex] || {}
    },
    totalQuestions() {
      return this.questions.length
    },
    isLastQuestion() {
      return this.currentQuestionIndex === this.totalQuestions - 1
    }
  },
  methods: {
    selectAnswer(value) {
      this.selectedAnswer = value
    },
    
    nextQuestion() {
      if (this.selectedAnswer === null) return
      
      this.answers[this.currentQuestion.id] = this.selectedAnswer
      
      if (this.isLastQuestion) {
        this.completeAssessment()
      } else {
        this.currentQuestionIndex++
        this.loadQuestionAnswer()
      }
    },
    
    previousQuestion() {
      if (this.currentQuestionIndex > 0) {
        this.currentQuestionIndex--
        this.loadQuestionAnswer()
      }
    },
    
    loadQuestionAnswer() {
      const savedAnswer = this.answers[this.currentQuestion.id]
      this.selectedAnswer = savedAnswer !== undefined ? savedAnswer : null
    },
    
    completeAssessment() {
      this.answers[this.currentQuestion.id] = this.selectedAnswer
      
      uni.showLoading({ title: '正在分析结果...' })
      
      setTimeout(() => {
        uni.hideLoading()
        uni.navigateTo({
          url: `/pages/assessment/result?answers=${encodeURIComponent(JSON.stringify(this.answers))}`
        })
      }, 2000)
    },
    
    handleBack() {
      uni.navigateBack()
    },
    
    showHelp() {
      uni.showModal({
        title: '评估说明',
        content: '本评估采用专业的心理学量表，请诚实回答每个问题。所有信息严格保密。',
        showCancel: false,
        confirmText: '我知道了'
      })
    }
  },
  onLoad() {
    this.loadQuestionAnswer()
  }
}
</script>

<style scoped>
.questionnaire-page {
  height: 100vh;
  background-color: var(--background);
  display: flex;
  flex-direction: column;
}

.status-bar {
  background: transparent;
}

.battery-indicator {
  width: 24px;
  height: 12px;
  border: 1px solid #1f2937;
  border-radius: 2px;
  position: relative;
}

.battery-level {
  width: 16px;
  height: 6px;
  background-color: #10b981;
  border-radius: 1px;
  margin: 2px 0 0 2px;
}

.nav-header {
  background-color: white;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  border-bottom: 1px solid #e5e7eb;
}

.nav-btn {
  padding: 8px;
  border-radius: 50%;
  transition: all 0.2s;
}

.nav-btn:active {
  background-color: #f3f4f6;
  transform: scale(0.95);
}

.nav-title {
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
}

.progress-section {
  padding-bottom: 12px;
}

.progress-text {
  font-size: 14px;
  color: #6b7280;
}

.progress-track {
  width: 100%;
  height: 8px;
  background-color: #e5e7eb;
  border-radius: 4px;
  overflow: hidden;
}

.progress-bar {
  height: 100%;
  background: linear-gradient(90deg, #34D399 0%, #10B981 100%);
  border-radius: 4px;
  transition: width 0.3s ease;
}

.main-content {
  flex: 1;
  padding: 24px;
}

.question-card {
  background-color: white;
  border-radius: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  border: 1px solid #e5e7eb;
  padding: 24px;
  margin-bottom: 24px;
}

.question-header {
  text-align: center;
  margin-bottom: 24px;
}

.category-icon {
  width: 64px;
  height: 64px;
  background-color: rgba(6, 182, 212, 0.1);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 16px;
}

.category-icon .nf-icon {
  font-size: 24px;
  color: #06b6d4;
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.category-title {
  font-size: 20px;
  font-weight: 600;
  color: #1f2937;
  display: block;
  margin-bottom: 8px;
}

.category-desc {
  font-size: 14px;
  color: #6b7280;
  display: block;
}

.question-content {
  margin-bottom: 32px;
}

.question-text {
  font-size: 18px;
  font-weight: 500;
  color: #1f2937;
  line-height: 1.6;
  display: block;
  margin-bottom: 24px;
}

.privacy-note {
  font-size: 14px;
  color: #9ca3af;
  display: block;
}

.option-btn {
  width: 100%;
  padding: 16px;
  text-align: left;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  background-color: white;
  transition: all 0.2s;
  margin-bottom: 12px;
}

.option-btn:active {
  transform: scale(0.98);
}

.option-btn:hover {
  border-color: var(--primary);
  background-color: rgba(52, 211, 153, 0.05);
}

.option-selected {
  border-color: var(--primary) !important;
  background-color: rgba(52, 211, 153, 0.05) !important;
}

.option-radio {
  width: 20px;
  height: 20px;
  border: 2px solid #d1d5db;
  border-radius: 50%;
  margin-right: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.radio-selected {
  border-color: var(--primary);
}

.radio-dot {
  width: 8px;
  height: 8px;
  background-color: var(--primary);
  border-radius: 50%;
}

.option-text {
  color: #374151;
  font-size: 16px;
}

.action-buttons {
  display: flex;
  gap: 16px;
}

.action-btn {
  flex: 1;
  padding: 16px;
  font-size: 16px;
  font-weight: 600;
  border-radius: 16px;
  transition: all 0.2s;
}

.action-btn:active {
  transform: scale(0.95);
}

.secondary-btn {
  background-color: #f3f4f6;
  color: #6b7280;
}

.primary-btn {
  background-color: var(--primary);
  color: white;
}

.btn-disabled {
  opacity: 0.5;
  pointer-events: none;
}

.help-section {
  margin-top: 24px;
  padding: 16px;
  background-color: #eff6ff;
  border-radius: 12px;
  border: 1px solid #bfdbfe;
}

.help-content {
  flex: 1;
}

.help-title {
  font-size: 16px;
  font-weight: 500;
  color: #1e40af;
  display: block;
  margin-bottom: 4px;
}

.help-text {
  font-size: 14px;
  color: #1d4ed8;
  line-height: 1.5;
}

.nf-icon {
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.flex { display: flex; }
.flex-1 { flex: 1; }
.items-center { align-items: center; }
.justify-between { justify-content: space-between; }
/* 注意：微信小程序不支持通用选择器 *，使用具体选择器 */
.space-x-1 > view + view { margin-left: 4px; }
.space-x-3 > view + view { margin-left: 12px; }
.space-x-4 > view + view { margin-left: 16px; }
.space-y-3 > view + view { margin-top: 12px; }
.text-xs { font-size: 12px; }
.text-lg { font-size: 18px; }
.font-medium { font-weight: 500; }
.text-gray-800 { color: #1f2937; }
.text-gray-600 { color: #4b5563; }
.text-gray-400 { color: #9ca3af; }
.text-blue-500 { color: #3b82f6; }
.px-4 { padding-left: 16px; padding-right: 16px; }
.py-2 { padding-top: 8px; padding-bottom: 8px; }
.py-3 { padding-top: 12px; padding-bottom: 12px; }
.px-6 { padding-left: 24px; padding-right: 24px; }
.py-6 { padding-top: 24px; padding-bottom: 24px; }
.pb-3 { padding-bottom: 12px; }
.mb-2 { margin-bottom: 8px; }
.mb-6 { margin-bottom: 24px; }
.mb-8 { margin-bottom: 32px; }
.mt-6 { margin-top: 24px; }
.mt-0_5 { margin-top: 2px; }
.bg-white { background-color: white; }
.bg-blue-50 { background-color: #eff6ff; }
.border { border-width: 1px; }
.border-gray-100 { border-color: #f3f4f6; }
.border-blue-100 { border-color: #bfdbfe; }
.rounded-xl { border-radius: 12px; }
.rounded-3xl { border-radius: 24px; }
.shadow-sm { box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05); }
</style>
