<template>
  <view class="legal-page">
    <!-- 顶部导航 -->
    <NfNavBar 
      title="服务协议"
      :show-back="true"
      @back="goBack"
    />
    
    <view class="page-content">
      <NfCard>
        <view class="legal-content">
          <view class="section">
            <text class="section-title">1. 服务说明</text>
            <text class="section-text">
              戒色助手是一款健康生活辅助应用，旨在通过科学方法帮助用户建立健康的生活习惯。我们提供评估、打卡、社区交流、学习等功能服务。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">2. 用户责任</text>
            <text class="section-text">
              用户应真实、准确地填写个人信息，遵守社区规范，不发布违法违规内容。用户对自己的账号和密码安全负责。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">3. 内容规范</text>
            <text class="section-text">
              禁止发布色情、暴力、违法或有害信息。我们保留审核、删除不当内容的权利。严重违规行为将导致账号封禁。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">4. 免责声明</text>
            <text class="section-text">
              本应用仅提供信息和工具支持，不替代专业医疗建议。如有严重心理健康问题，请寻求专业医疗帮助。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">5. 服务变更</text>
            <text class="section-text">
              我们保留随时修改或终止服务的权利，重要变更将通过应用内通知的方式告知用户。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">6. 联系我们</text>
            <text class="section-text">
              如对本协议有任何疑问，请联系：support@nofap-helper.com
            </text>
          </view>
          
          <view class="update-info">
            <text class="update-text">最后更新：2024年1月1日</text>
          </view>
        </view>
      </NfCard>
      
      <view class="action-section">
        <NfButton 
          type="primary"
          size="large"
          label="我已阅读并同意"
          full-width
          @click="agreeAndContinue"
        />
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'ServiceAgreementPage',
  methods: {
    goBack() {
      uni.navigateBack()
    },
    
    agreeAndContinue() {
      // 记录用户同意协议
      uni.setStorageSync('serviceAgreementAccepted', true)
      
      uni.showToast({
        title: '感谢您的同意',
        icon: 'success'
      })
      
      setTimeout(() => {
        uni.navigateBack()
      }, 1000)
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

.legal-content {
  padding: 8px 0;
}

.section {
  margin-bottom: 24px;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  display: block;
  margin-bottom: 8px;
}

.section-text {
  font-size: 14px;
  line-height: 1.6;
  color: var(--text-secondary);
  display: block;
}

.update-info {
  margin-top: 32px;
  padding-top: 16px;
  border-top: 1px solid #E5E7EB;
  text-align: center;
}

.update-text {
  font-size: 12px;
  color: var(--text-muted);
}

.action-section {
  margin-top: 24px;
}
</style> 