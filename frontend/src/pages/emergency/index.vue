<template>
  <view class="emergency-page">
    <!-- Status Bar -->
    <view class="status-bar">
      <text class="time">{{ currentTime }}</text>
      <view class="status-icons">
        <text class="icon">📶</text>
        <text class="icon">📶</text>
        <view class="battery">
          <view class="battery-level"></view>
        </view>
      </view>
    </view>

    <!-- Header with Calm Gradient -->
    <view class="header-section">
      <view class="header">
        <view class="nav-button" @click="goBack">
          <text class="nav-icon">←</text>
        </view>
        <text class="header-title">紧急求助</text>
        <view class="nav-button" @click="callEmergency">
          <text class="nav-icon">📞</text>
        </view>
      </view>
      
      <view class="hero-section">
        <view class="breathing-circle" :class="{ breathing: isBreathing }">
          <text class="heart-icon">💖</text>
        </view>
        <text class="hero-title">你并不孤单</text>
        <text class="hero-subtitle">这种感觉会过去的，让我们一起度过</text>
      </view>
    </view>

    <!-- Main Content -->
    <scroll-view class="main-content" scroll-y>
      <!-- Immediate Actions -->
      <view class="section">
        <text class="section-title">立即行动</text>
        
        <!-- Breathing Exercise -->
        <view class="action-card" @click="startBreathingExercise">
          <view class="card-header">
            <view class="card-icon breathing-icon">
              <text class="icon-text">🫁</text>
            </view>
            <view class="card-info">
              <text class="card-title">深呼吸练习</text>
              <text class="card-subtitle">4-7-8呼吸法，立即平静心情</text>
            </view>
          </view>
          <view class="action-button breathing-button">
            <text class="button-text">开始 3 分钟呼吸练习</text>
          </view>
        </view>

        <!-- Physical Exercise -->
        <view class="action-card" @click="startPhysicalExercise">
          <view class="card-header">
            <view class="card-icon exercise-icon">
              <text class="icon-text">💪</text>
            </view>
            <view class="card-info">
              <text class="card-title">快速运动</text>
              <text class="card-subtitle">20个俯卧撑或深蹲，释放压力</text>
            </view>
          </view>
          <view class="action-button exercise-button">
            <text class="button-text">开始运动指导</text>
          </view>
        </view>
      </view>

      <!-- Distraction Activities -->
      <view class="section">
        <text class="section-title">转移注意力</text>
        
        <view class="activity-grid">
          <view class="activity-item" @click="openActivity('puzzle')">
            <view class="activity-icon puzzle-icon">
              <text class="icon-text">🧩</text>
            </view>
            <text class="activity-title">益智游戏</text>
            <text class="activity-subtitle">数独、拼图</text>
          </view>
          
          <view class="activity-item" @click="openActivity('meditation')">
            <view class="activity-icon meditation-icon">
              <text class="icon-text">🌱</text>
            </view>
            <text class="activity-title">正念冥想</text>
            <text class="activity-subtitle">5分钟冥想</text>
          </view>
          
          <view class="activity-item" @click="openActivity('music')">
            <view class="activity-icon music-icon">
              <text class="icon-text">🎵</text>
            </view>
            <text class="activity-title">舒缓音乐</text>
            <text class="activity-subtitle">放松播放列表</text>
          </view>
          
          <view class="activity-item" @click="openActivity('reading')">
            <view class="activity-icon reading-icon">
              <text class="icon-text">📖</text>
            </view>
            <text class="activity-title">励志文章</text>
            <text class="activity-subtitle">正能量内容</text>
          </view>
        </view>
      </view>

      <!-- Community Support -->
      <view class="community-section">
        <view class="community-header">
          <view class="community-icon">
            <text class="icon-text">👥</text>
          </view>
          <text class="community-title">寻求社区支持</text>
          <text class="community-subtitle">
            有 {{ onlineVolunteers }} 位伙伴正在线上，他们愿意为你提供支持和鼓励
          </text>
        </view>
        
        <view class="community-actions">
          <view class="primary-button" @click="requestHelp">
            <text class="button-text">匿名求助 - 立即连接</text>
          </view>
          <view class="secondary-button" @click="viewStories">
            <text class="button-text">查看励志故事</text>
          </view>
        </view>
      </view>

      <!-- Motivational Quote -->
      <view class="quote-section">
        <text class="quote-icon">💬</text>
        <text class="quote-text">
          "{{ currentQuote.text }}"
        </text>
        <text class="quote-author">— {{ currentQuote.author }}</text>
      </view>
    </scroll-view>

    <!-- Emergency Help Modal -->
    <view v-if="showHelpModal" class="modal-overlay" @click="closeHelpModal">
      <view class="help-modal" @click.stop>
        <view class="modal-header">
          <text class="modal-title">紧急求助</text>
          <text class="close-button" @click="closeHelpModal">✕</text>
        </view>
        <view class="modal-content">
          <text class="modal-text">请选择你遇到的情况：</text>
          <view class="help-types">
            <view class="help-type" @click="selectHelpType(1)">
              <text class="type-title">紧急冲动</text>
              <text class="type-desc">感到强烈的冲动，需要立即帮助</text>
            </view>
            <view class="help-type" @click="selectHelpType(2)">
              <text class="type-title">情绪低落</text>
              <text class="type-desc">心情沮丧，需要支持鼓励</text>
            </view>
            <view class="help-type" @click="selectHelpType(3)">
              <text class="type-title">复发担忧</text>
              <text class="type-desc">担心自己会复发，需要指导</text>
            </view>
            <view class="help-type" @click="selectHelpType(4)">
              <text class="type-title">其他情况</text>
              <text class="type-desc">其他需要帮助的情况</text>
            </view>
          </view>
        </view>
      </view>
    </view>

    <!-- Resource Modal -->
    <view v-if="showResourceModal" class="modal-overlay" @click="closeResourceModal">
      <view class="resource-modal" @click.stop>
        <view class="modal-header">
          <text class="modal-title">{{ selectedResource.title }}</text>
          <text class="close-button" @click="closeResourceModal">✕</text>
        </view>
        <view class="modal-content">
          <scroll-view class="resource-content" scroll-y>
            <text class="resource-text">{{ selectedResource.content }}</text>
          </scroll-view>
          <view class="resource-actions">
            <view class="resource-button" @click="useResource">
              <text class="button-text">开始使用</text>
            </view>
            <view class="resource-button secondary" @click="rateResource">
              <text class="button-text">评分</text>
            </view>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { ref, onMounted, reactive } from 'vue'

export default {
  name: 'EmergencyHelp',
  setup() {
    const currentTime = ref('9:41')
    const isBreathing = ref(true)
    const onlineVolunteers = ref(23)
    const showHelpModal = ref(false)
    const showResourceModal = ref(false)
    const selectedResource = reactive({
      title: '',
      content: '',
      id: 0
    })

    const currentQuote = ref({
      text: "你比你想象的更坚强，这个时刻会过去的。",
      author: "戒色社区"
    })

    const quotes = [
      {
        text: "你比你想象的更坚强，这个时刻会过去的。",
        author: "戒色社区"
      },
      {
        text: "每个选择都是新的开始，每一刻都有改变的可能。",
        author: "戒色助手"
      },
      {
        text: "成长需要时间，对自己要有耐心。",
        author: "励志语录"
      }
    ]

    const emergencyResources = ref([
      {
        id: 1,
        title: "4-7-8深呼吸练习",
        type: 1,
        content: "这是一种简单而有效的呼吸技巧，可以帮助您快速平静下来。\n\n步骤：\n1. 吸气4秒钟\n2. 屏住呼吸7秒钟\n3. 呼气8秒钟\n4. 重复3-4次\n\n这种呼吸方式可以激活副交感神经系统，帮助身体放松。",
        duration: 180
      },
      {
        id: 2,
        title: "5分钟正念冥想",
        type: 2,
        content: "正念冥想可以帮助您专注当下，减少负面情绪的影响。\n\n指导：\n1. 找一个安静的地方坐下\n2. 闭上眼睛，专注于呼吸\n3. 当思绪飘散时，温和地把注意力拉回呼吸\n4. 观察身体的感觉，不做判断\n5. 保持5分钟\n\n记住：没有'正确'或'错误'的冥想，只要保持观察即可。",
        duration: 300
      },
      {
        id: 5,
        title: "快速运动指导",
        type: 5,
        content: "运动可以释放内啡肽，帮助改善心情。这里有一套简单的运动，无需器械：\n\n**热身（1分钟）**\n• 原地踏步 30秒\n• 手臂绕圈 30秒\n\n**主要运动（3分钟）**\n• 俯卧撑 20个（可膝盖着地）\n• 深蹲 20个\n• 平板支撑 30秒\n• 开合跳 20个\n\n**放松（1分钟）**\n• 深呼吸 30秒\n• 拉伸手臂和腿部 30秒",
        duration: 300
      }
    ])

    const updateTime = () => {
      const now = new Date()
      const hours = now.getHours().toString().padStart(2, '0')
      const minutes = now.getMinutes().toString().padStart(2, '0')
      currentTime.value = `${hours}:${minutes}`
    }

    const rotateQuote = () => {
      const randomIndex = Math.floor(Math.random() * quotes.length)
      currentQuote.value = quotes[randomIndex]
    }

    const goBack = () => {
      uni.navigateBack()
    }

    const callEmergency = () => {
      uni.showModal({
        title: '紧急联系',
        content: '如果遇到生命危险，请立即拨打120或当地急救电话',
        confirmText: '知道了',
        showCancel: false
      })
    }

    const startBreathingExercise = () => {
      const resource = emergencyResources.value.find(r => r.type === 1)
      if (resource) {
        openResourceModal(resource)
      }
    }

    const startPhysicalExercise = () => {
      const resource = emergencyResources.value.find(r => r.type === 5)
      if (resource) {
        openResourceModal(resource)
      }
    }

    const openActivity = (type) => {
      let resource
      switch (type) {
        case 'meditation':
          resource = emergencyResources.value.find(r => r.type === 2)
          break
        case 'music':
          uni.showToast({
            title: '音乐功能开发中',
            icon: 'none'
          })
          return
        case 'reading':
          uni.navigateTo({
            url: '/pages/emergency/articles'
          })
          return
        case 'puzzle':
          uni.showToast({
            title: '游戏功能开发中',
            icon: 'none'
          })
          return
      }
      
      if (resource) {
        openResourceModal(resource)
      }
    }

    const openResourceModal = (resource) => {
      selectedResource.title = resource.title
      selectedResource.content = resource.content
      selectedResource.id = resource.id
      showResourceModal.value = true
    }

    const closeResourceModal = () => {
      showResourceModal.value = false
    }

    const useResource = () => {
      // 记录使用次数
      uni.showToast({
        title: '开始使用',
        icon: 'success'
      })
      closeResourceModal()
    }

    const rateResource = () => {
      uni.showToast({
        title: '感谢你的反馈',
        icon: 'success'
      })
    }

    const requestHelp = () => {
      showHelpModal.value = true
    }

    const closeHelpModal = () => {
      showHelpModal.value = false
    }

    const selectHelpType = (type) => {
      closeHelpModal()
      
      // 这里可以集成实际的求助功能
      uni.showModal({
        title: '求助已发送',
        content: '你的求助信息已发送给在线志愿者，他们会尽快回复你。',
        confirmText: '好的',
        showCancel: false,
        success: () => {
          // 可以跳转到聊天页面或等待页面
          uni.showToast({
            title: '正在为您匹配志愿者',
            icon: 'loading',
            duration: 2000
          })
        }
      })
    }

    const viewStories = () => {
      uni.navigateTo({
        url: '/pages/community/index?category=4' // 成功故事分类
      })
    }

    onMounted(() => {
      updateTime()
      setInterval(updateTime, 60000) // 每分钟更新时间
      
      // 每30秒切换励志语录
      setInterval(rotateQuote, 30000)
      
      // 模拟获取在线志愿者数量
      // 实际项目中应该调用API
    })

    return {
      currentTime,
      isBreathing,
      onlineVolunteers,
      currentQuote,
      showHelpModal,
      showResourceModal,
      selectedResource,
      goBack,
      callEmergency,
      startBreathingExercise,
      startPhysicalExercise,
      openActivity,
      requestHelp,
      closeHelpModal,
      selectHelpType,
      viewStories,
      closeResourceModal,
      useResource,
      rateResource
    }
  }
}
</script>

<style scoped>
.emergency-page {
  min-height: 100vh;
  background: #F8FAFC;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', system-ui, sans-serif;
}

/* Status Bar */
.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
  font-size: 12px;
  color: #1F2937;
  font-weight: 500;
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

/* Header Section */
.header-section {
  background: linear-gradient(135deg, #E0E7FF 0%, #DBEAFE 50%, #E0F2FE 100%);
  padding: 16px 24px 32px;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
}

.nav-button {
  width: 40px;
  height: 40px;
  background: rgba(255, 255, 255, 0.5);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(4px);
}

.nav-icon {
  font-size: 16px;
  color: #374151;
}

.header-title {
  font-size: 20px;
  font-weight: bold;
  color: #1F2937;
}

.hero-section {
  text-align: center;
}

.breathing-circle {
  width: 96px;
  height: 96px;
  background: linear-gradient(135deg, #8B5CF6 0%, #06B6D4 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 16px;
  transition: transform 0.3s ease;
}

.breathing-circle.breathing {
  animation: breathe 4s ease-in-out infinite;
}

@keyframes breathe {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}

.heart-icon {
  font-size: 32px;
  color: white;
}

.hero-title {
  font-size: 32px;
  font-weight: bold;
  color: #1F2937;
  margin-bottom: 8px;
  display: block;
}

.hero-subtitle {
  color: #4B5563;
  font-size: 16px;
  display: block;
}

/* Main Content */
.main-content {
  flex: 1;
  padding: 0 24px 80px;
}

.section {
  margin-bottom: 32px;
}

.section-title {
  font-size: 18px;
  font-weight: bold;
  color: #1F2937;
  margin-bottom: 16px;
  display: block;
}

/* Action Cards */
.action-card {
  background: white;
  border-radius: 16px;
  padding: 24px;
  margin-bottom: 16px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  align-items: flex-start;
  gap: 16px;
  margin-bottom: 16px;
}

.card-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.breathing-icon {
  background: rgba(139, 92, 246, 0.1);
}

.exercise-icon {
  background: rgba(52, 211, 153, 0.1);
}

.icon-text {
  font-size: 24px;
}

.card-info {
  flex: 1;
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 4px;
  display: block;
}

.card-subtitle {
  font-size: 14px;
  color: #6B7280;
  display: block;
}

.action-button {
  width: 100%;
  padding: 12px;
  border-radius: 12px;
  text-align: center;
}

.breathing-button {
  background: #8B5CF6;
}

.exercise-button {
  background: #34D399;
}

.button-text {
  color: white;
  font-weight: 600;
  font-size: 16px;
}

/* Activity Grid */
.activity-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.activity-item {
  background: white;
  border-radius: 16px;
  padding: 16px;
  text-align: center;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.activity-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 8px;
}

.puzzle-icon {
  background: rgba(147, 51, 234, 0.1);
}

.meditation-icon {
  background: rgba(16, 185, 129, 0.1);
}

.music-icon {
  background: rgba(59, 130, 246, 0.1);
}

.reading-icon {
  background: rgba(245, 158, 11, 0.1);
}

.activity-title {
  font-size: 14px;
  font-weight: 500;
  color: #1F2937;
  margin-bottom: 4px;
  display: block;
}

.activity-subtitle {
  font-size: 12px;
  color: #6B7280;
  display: block;
}

/* Community Support */
.community-section {
  background: linear-gradient(to right, #EFF6FF, #EEF2FF);
  border-radius: 16px;
  padding: 24px;
  margin-bottom: 24px;
  border: 1px solid #DBEAFE;
}

.community-header {
  text-align: center;
  margin-bottom: 16px;
}

.community-icon {
  width: 64px;
  height: 64px;
  background: #DBEAFE;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 12px;
}

.community-title {
  font-size: 16px;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 8px;
  display: block;
}

.community-subtitle {
  font-size: 14px;
  color: #6B7280;
  line-height: 1.5;
  display: block;
  margin-bottom: 16px;
}

.community-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.primary-button {
  background: #2563EB;
  color: white;
  padding: 12px;
  border-radius: 12px;
  text-align: center;
  font-weight: 600;
}

.secondary-button {
  border: 2px solid #2563EB;
  color: #2563EB;
  padding: 12px;
  border-radius: 12px;
  text-align: center;
  font-weight: 600;
}

/* Quote Section */
.quote-section {
  text-align: center;
  padding: 24px;
  background: linear-gradient(to right, #FEF3C7, #FED7AA);
  border-radius: 16px;
  border: 1px solid #FEF3C7;
}

.quote-icon {
  font-size: 32px;
  color: #D97706;
  margin-bottom: 12px;
  display: block;
}

.quote-text {
  font-size: 18px;
  font-weight: 500;
  color: #1F2937;
  font-style: italic;
  margin-bottom: 8px;
  display: block;
  line-height: 1.5;
}

.quote-author {
  font-size: 14px;
  color: #6B7280;
  display: block;
}

/* Modals */
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
  padding: 20px;
}

.help-modal,
.resource-modal {
  background: white;
  border-radius: 16px;
  width: 100%;
  max-width: 400px;
  max-height: 80vh;
  overflow: hidden;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  border-bottom: 1px solid #E5E7EB;
}

.modal-title {
  font-size: 18px;
  font-weight: bold;
  color: #1F2937;
}

.close-button {
  font-size: 20px;
  color: #6B7280;
  padding: 4px;
}

.modal-content {
  padding: 20px;
}

.modal-text {
  font-size: 16px;
  color: #4B5563;
  margin-bottom: 16px;
  display: block;
}

.help-types {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.help-type {
  padding: 16px;
  border: 2px solid #E5E7EB;
  border-radius: 12px;
  transition: border-color 0.2s;
}

.help-type:active {
  border-color: #3B82F6;
  background: #EFF6FF;
}

.type-title {
  font-size: 16px;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 4px;
  display: block;
}

.type-desc {
  font-size: 14px;
  color: #6B7280;
  display: block;
}

.resource-content {
  max-height: 300px;
  margin-bottom: 20px;
}

.resource-text {
  font-size: 14px;
  line-height: 1.6;
  color: #374151;
  white-space: pre-line;
}

.resource-actions {
  display: flex;
  gap: 12px;
}

.resource-button {
  flex: 1;
  padding: 12px;
  background: #3B82F6;
  color: white;
  border-radius: 8px;
  text-align: center;
  font-weight: 600;
}

.resource-button.secondary {
  background: #E5E7EB;
  color: #374151;
}
</style> 