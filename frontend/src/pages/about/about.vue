<template>
  <view class="about-page">
    <!-- 顶部导航 -->
    <NfNavBar 
      title="了解更多"
      :show-back="true"
      @back="goBack"
    />
    
    <view class="page-content">
      <!-- 应用介绍 -->
      <NfCard type="gradient" icon="fa-seedling" title="戒色助手">
        <view class="app-intro">
          <text class="intro-text">
            戒色助手是一款专为戒除色情依赖而设计的综合性健康应用。通过科学的评估系统、游戏化的激励机制和温暖的社区支持，帮助用户建立健康的生活方式，重获内心的平静与力量。
          </text>
        </view>
      </NfCard>
      
      <!-- 核心功能 -->
      <NfCard title="核心功能" icon="fa-star">
        <view class="features-list space-y-4">
          <view class="feature-item">
            <view class="feature-header">
              <text class="nf-icon fa-chart-line feature-icon"></text>
              <text class="feature-name">科学评估系统</text>
            </view>
            <text class="feature-description">
              基于专业心理学理论，提供50道精心设计的评估问题，科学判断依赖程度，为个性化康复方案提供数据支持。
            </text>
          </view>
          
          <view class="feature-item">
            <view class="feature-header">
              <text class="nf-icon fa-gamepad feature-icon"></text>
              <text class="feature-name">游戏化激励体验</text>
            </view>
            <text class="feature-description">
              设计50级成长体系和100+成就徽章，让戒色过程充满成就感。每日打卡、连续记录、等级升级，让改变变得有趣而持久。
            </text>
          </view>
          
          <view class="feature-item">
            <view class="feature-header">
              <text class="nf-icon fa-users feature-icon"></text>
              <text class="feature-name">匿名社区支持</text>
            </view>
            <text class="feature-description">
              匿名的交流环境，分享经验、互相鼓励。专业的内容审核机制确保社区的正面氛围，让每个人都能获得理解和支持。
            </text>
          </view>
          
          <view class="feature-item">
            <view class="feature-header">
              <text class="nf-icon fa-heart feature-icon"></text>
              <text class="feature-name">紧急求助系统</text>
            </view>
            <text class="feature-description">
              一键求助功能，在关键时刻提供即时支持。志愿者响应机制和专业的危机干预指导，确保没有人独自面对困难。
            </text>
          </view>
          
          <view class="feature-item">
            <view class="feature-header">
              <text class="nf-icon fa-book feature-icon"></text>
              <text class="feature-name">学习成长内容</text>
            </view>
            <text class="feature-description">
              精选的心理健康知识、自我管理技巧和正能量文章。个性化推荐算法确保每个人都能获得最适合的成长内容。
            </text>
          </view>
        </view>
      </NfCard>
      
      <!-- 设计理念 -->
      <NfCard title="设计理念" icon="fa-lightbulb">
        <view class="principles space-y-3">
          <view class="principle-item">
            <text class="principle-title">🌱 温暖友好</text>
            <text class="principle-desc">用鼓励代替指责，用理解代替批判</text>
          </view>
          <view class="principle-item">
            <text class="principle-title">🔒 隐私保护</text>
            <text class="principle-desc">严格的匿名机制，保护每个用户的隐私</text>
          </view>
          <view class="principle-item">
            <text class="principle-title">🎮 游戏化</text>
            <text class="principle-desc">让改变过程充满乐趣和成就感</text>
          </view>
          <view class="principle-item">
            <text class="principle-title">🤝 社区支持</text>
            <text class="principle-desc">没有人是孤岛，共同成长更有力量</text>
          </view>
        </view>
      </NfCard>
      
      <!-- 开始使用 -->
      <view class="action-section">
        <NfButton 
          type="primary"
          size="large"
          label="开始我的康复之旅"
          full-width
          icon-left="fa-play"
          @click="startJourney"
        />
        
        <view class="support-text">
          <text>需要帮助？联系我们：support@nofap-helper.com</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'AboutPage',
  methods: {
    goBack() {
      uni.navigateBack()
    },
    
    startJourney() {
      uni.navigateTo({
        url: '/pages/assessment/assessment'
      })
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

/* 应用介绍 */
.app-intro {
  text-align: center;
}

.intro-text {
  font-size: 16px;
  line-height: 1.6;
  color: white;
}

/* 功能列表 */
.features-list {
  /* space-y-4 通过全局样式实现 */
}

.feature-item {
  padding-bottom: 16px;
  border-bottom: 1px solid #E5E7EB;
}

.feature-item:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.feature-header {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
}

.feature-icon {
  font-size: 18px;
  color: var(--primary);
  margin-right: 12px;
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.feature-name {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
}

.feature-description {
  font-size: 14px;
  line-height: 1.5;
  color: var(--text-secondary);
  padding-left: 30px;
}

/* 设计原则 */
.principles {
  /* space-y-3 通过全局样式实现 */
}

.principle-item {
  display: flex;
  flex-direction: column;
  padding: 12px;
  background-color: #F9FAFB;
  border-radius: 12px;
}

.principle-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.principle-desc {
  font-size: 14px;
  color: var(--text-secondary);
}

/* 操作区域 */
.action-section {
  margin-top: 24px;
  text-align: center;
}

.support-text {
  margin-top: 16px;
}

.support-text text {
  font-size: 12px;
  color: var(--text-muted);
}

/* FontAwesome图标 */
.nf-icon {
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

/* 间距工具类 */
/* 注意：微信小程序不支持通用选择器 *，使用具体选择器 */
.space-y-3 > view + view {
  margin-top: 12px;
}

.space-y-4 > view + view {
  margin-top: 16px;
}
</style> 