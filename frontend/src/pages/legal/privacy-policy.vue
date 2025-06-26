<template>
  <view class="legal-page">
    <!-- 顶部导航 -->
    <NfNavBar 
      title="隐私政策"
      :show-back="true"
      @back="goBack"
    />
    
    <view class="page-content">
      <NfCard>
        <view class="legal-content">
          <view class="section">
            <text class="section-title">1. 信息收集</text>
            <text class="section-text">
              我们仅收集提供服务所必需的信息，包括微信授权信息、使用统计数据、评估结果等。我们承诺不收集敏感个人信息。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">2. 信息使用</text>
            <text class="section-text">
              收集的信息仅用于：提供个性化服务、改进产品功能、生成匿名统计报告。我们不会将您的信息用于营销或其他商业目的。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">3. 信息保护</text>
            <text class="section-text">
              我们采用行业标准的安全措施保护您的信息，包括加密存储、访问控制、定期安全审计等。所有数据传输均采用HTTPS加密。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">4. 匿名化处理</text>
            <text class="section-text">
              社区功能采用严格的匿名机制，其他用户无法获取您的真实身份信息。评估结果和统计数据均经过匿名化处理。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">5. 第三方服务</text>
            <text class="section-text">
              我们使用微信小程序平台提供服务，遵循微信的隐私政策。不会与其他第三方分享您的个人信息。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">6. 数据删除</text>
            <text class="section-text">
              您可以随时申请删除个人数据。停用账号后，我们将在合理期限内删除相关信息，但匿名统计数据可能保留。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">7. 政策更新</text>
            <text class="section-text">
              隐私政策如有更新，我们将通过应用内通知告知。继续使用服务即表示同意更新后的政策。
            </text>
          </view>
          
          <view class="section">
            <text class="section-title">8. 联系我们</text>
            <text class="section-text">
              如有隐私相关问题，请联系：privacy@nofap-helper.com
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
  name: 'PrivacyPolicyPage',
  methods: {
    goBack() {
      uni.navigateBack()
    },
    
    agreeAndContinue() {
      // 记录用户同意隐私政策
      uni.setStorageSync('privacyPolicyAccepted', true)
      
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