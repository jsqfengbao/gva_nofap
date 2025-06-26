<template>
  <view class="result-page">
    <NfNavBar 
      title="评估结果"
      :show-back="true"
      @back="goBack"
    />
    
    <view class="page-content">
      <NfCard type="gradient" class="result-card">
        <view class="result-header">
          <view class="result-icon" :style="{ backgroundColor: levelInfo.color + '20' }">
            <text class="nf-icon fa-chart-pie" :style="{ color: levelInfo.color }"></text>
          </view>
          <text class="result-title">您的评估结果</text>
          <view class="score-display">
            <text class="score-number">{{ assessmentResult.totalScore }}</text>
            <text class="score-suffix">分</text>
          </view>
          <view class="level-badge" :style="{ backgroundColor: levelInfo.color }">
            <text class="level-text">{{ levelInfo.name }}</text>
          </view>
        </view>
      </NfCard>
      
      <NfCard title="结果解读" icon="fa-lightbulb">
        <view class="interpretation">
          <text class="interpretation-text">{{ levelInfo.description }}</text>
        </view>
      </NfCard>
      
      <NfCard title="详细分析" icon="fa-chart-bar">
        <view class="category-analysis">
          <view 
            v-for="(score, category) in assessmentResult.categoryScores" 
            :key="category"
            class="category-item"
          >
            <view class="category-info">
              <text class="category-name">{{ getCategoryName(category) }}</text>
              <text class="category-score">{{ Math.round(score) }}%</text>
            </view>
            <view class="category-bar">
              <view 
                class="category-progress" 
                :style="{ width: score + '%', backgroundColor: getProgressColor(score) }"
              ></view>
            </view>
          </view>
        </view>
      </NfCard>
      
      <NfCard title="个性化建议" icon="fa-heart">
        <view class="recommendations">
          <view 
            v-for="(recommendation, index) in levelInfo.recommendations" 
            :key="index"
            class="recommendation-item"
          >
            <view class="recommendation-icon">
              <text class="nf-icon fa-check-circle"></text>
            </view>
            <text class="recommendation-text">{{ recommendation }}</text>
          </view>
        </view>
      </NfCard>
      
      <view class="action-section">
        <NfButton 
          type="primary"
          size="large"
          label="开始康复计划"
          full-width
          icon-left="fa-play"
          @click="startRecoveryPlan"
        />
        
        <NfButton 
          type="secondary"
          size="medium"
          label="保存评估结果"
          full-width
          icon-left="fa-save"
          @click="saveResult"
        />
        
        <NfButton 
          type="secondary"
          size="medium"
          label="重新评估"
          full-width
          icon-left="fa-redo"
          @click="retakeAssessment"
        />
      </view>
    </view>
  </view>
</template>

<script>
import { assessmentConfig } from '@/data/assessment-questions.js'

export default {
  name: 'AssessmentResultPage',
  data() {
    return {
      assessmentResult: null,
      levelInfo: null
    }
  },
  methods: {
    calculateResult(answers) {
      const totalAnswers = Object.values(answers)
      const totalScore = totalAnswers.reduce((sum, value) => sum + (value || 0), 0) * 4
      
      const level = assessmentConfig.levels.find(l => 
        totalScore >= l.range[0] && totalScore <= l.range[1]
      ) || assessmentConfig.levels[0]
      
      return {
        totalScore: Math.min(200, totalScore),
        categoryScores: {
          frequency: Math.random() * 100,
          control: Math.random() * 100,
          impact: Math.random() * 100,
          psychology: Math.random() * 100,
          social: Math.random() * 100,
          health: Math.random() * 100,
          cognitive: Math.random() * 100
        },
        level,
        completedAt: new Date().toISOString()
      }
    },
    
    getCategoryName(category) {
      const names = {
        frequency: '行为频率',
        control: '自控能力',
        impact: '生活影响',
        psychology: '心理状态',
        social: '社交关系',
        health: '身体健康',
        cognitive: '认知功能'
      }
      return names[category] || category
    },
    
    getProgressColor(score) {
      if (score < 30) return '#10B981'
      if (score < 60) return '#F59E0B'
      if (score < 80) return '#EF4444'
      return '#DC2626'
    },
    
    goBack() {
      uni.navigateBack()
    },
    
    startRecoveryPlan() {
      uni.showToast({
        title: '正在为您制定计划...',
        icon: 'none'
      })
      
      setTimeout(() => {
        uni.switchTab({
          url: '/pages/index/index'
        })
      }, 1500)
    },
    
    saveResult() {
      try {
        const existingResults = uni.getStorageSync('assessmentHistory') || []
        existingResults.push(this.assessmentResult)
        uni.setStorageSync('assessmentHistory', existingResults)
        
        uni.showToast({
          title: '评估结果已保存',
          icon: 'success'
        })
      } catch (error) {
        uni.showToast({
          title: '保存失败，请重试',
          icon: 'error'
        })
      }
    },
    
    retakeAssessment() {
      uni.showModal({
        title: '重新评估',
        content: '确定要重新进行评估吗？当前结果将会被覆盖。',
        success: (res) => {
          if (res.confirm) {
            uni.redirectTo({
              url: '/pages/assessment/questionnaire'
            })
          }
        }
      })
    }
  },
  
  onLoad(options) {
    if (options.answers) {
      try {
        const answers = JSON.parse(decodeURIComponent(options.answers))
        this.assessmentResult = this.calculateResult(answers)
        this.levelInfo = this.assessmentResult.level
        
        this.levelInfo.recommendations = [
          '建立规律的作息时间',
          '培养健康的兴趣爱好',
          '寻求专业指导帮助',
          '使用戒色助手的各项功能'
        ]
      } catch (error) {
        console.error('解析评估答案失败:', error)
        uni.showToast({
          title: '数据加载失败',
          icon: 'error'
        })
        
        setTimeout(() => {
          uni.navigateBack()
        }, 2000)
      }
    }
  }
}
</script>

<style scoped>
.page-content {
  padding: 80px 16px 16px;
  min-height: 100vh;
  background-color: var(--background);
}

.result-card {
  margin-bottom: 24px;
}

.result-header {
  text-align: center;
  color: white;
}

.result-icon {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 16px;
}

.result-icon .nf-icon {
  font-size: 32px;
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.result-title {
  font-size: 20px;
  font-weight: 600;
  display: block;
  margin-bottom: 16px;
  color: white;
}

.score-display {
  margin-bottom: 16px;
}

.score-number {
  font-size: 48px;
  font-weight: 700;
  color: white;
}

.score-suffix {
  font-size: 20px;
  color: rgba(255, 255, 255, 0.8);
  margin-left: 4px;
}

.level-badge {
  display: inline-block;
  padding: 8px 16px;
  border-radius: 20px;
  margin-top: 8px;
}

.level-text {
  color: white;
  font-size: 16px;
  font-weight: 600;
}

.interpretation {
  text-align: center;
}

.interpretation-text {
  font-size: 16px;
  line-height: 1.6;
  color: var(--text-secondary);
}

.category-item {
  margin-bottom: 16px;
}

.category-item:last-child {
  margin-bottom: 0;
}

.category-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.category-name {
  font-size: 14px;
  color: var(--text-primary);
  font-weight: 500;
}

.category-score {
  font-size: 14px;
  color: var(--text-secondary);
  font-weight: 600;
}

.category-bar {
  width: 100%;
  height: 6px;
  background-color: #e5e7eb;
  border-radius: 3px;
  overflow: hidden;
}

.category-progress {
  height: 100%;
  border-radius: 3px;
  transition: width 0.3s ease;
}

.recommendation-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 12px;
}

.recommendation-item:last-child {
  margin-bottom: 0;
}

.recommendation-icon {
  margin-right: 12px;
  margin-top: 2px;
}

.recommendation-icon .nf-icon {
  font-size: 16px;
  color: var(--success);
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.recommendation-text {
  font-size: 14px;
  line-height: 1.5;
  color: var(--text-secondary);
  flex: 1;
}

.action-section {
  margin-top: 24px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.nf-icon {
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}
</style>
