<template>
  <view class="setup-page">
    <ProfileSetup
      @save="handleSave"
      @skip="handleSkip"
    />
  </view>
</template>

<script setup>
import ProfileSetup from '@/components/ui/profile/ProfileSetup.vue'
import { userApi } from '@/utils/api.js'
import { setUserInfo, getUserInfo, getToken } from '@/utils/auth.js'
import { getApiUrl } from '@/config/index.js'

// 保存用户资料
const handleSave = async (profileData) => {
  try {
    uni.showLoading({
      title: '保存中...'
    })

    // 如果有头像，先上传头像
    let avatarUrl = profileData.avatarUrl
    if (avatarUrl && avatarUrl.startsWith('http://tmp/')) {
      // 这是微信临时头像，需要上传到服务器
      avatarUrl = await uploadTempAvatar(avatarUrl)
    }

    // 更新用户信息
    await userApi.updateUserInfo({
      nickname: profileData.nickname,
      avatarUrl: avatarUrl
    })

    // 更新本地存储的用户信息
    const currentUser = getUserInfo()
    const updatedUser = {
      ...currentUser,
      nickname: profileData.nickname,
      avatarUrl: avatarUrl
    }
    setUserInfo(updatedUser)

    uni.hideLoading()
    uni.showToast({
      title: '保存成功',
      icon: 'success'
    })

    // 返回个人中心
    setTimeout(() => {
      uni.switchTab({
        url: '/pages/profile/index'
      })
    }, 1500)

  } catch (error) {
    console.error('保存用户资料失败:', error)
    uni.hideLoading()
    uni.showToast({
      title: '保存失败',
      icon: 'none'
    })
  }
}

// 上传临时头像
const uploadTempAvatar = async (tempPath) => {
  try {
    // 如果是微信临时头像
    if (tempPath.includes('wxfile://')) {
      // 使用专门的API保存微信头像
      const response = await userApi.saveWxAvatar({
        tempUrl: tempPath
      })
      return response.data.url
    } else {
      // 普通文件上传
      return new Promise((resolve, reject) => {
        uni.uploadFile({
          url: getApiUrl('/user/upload-avatar'),
          filePath: tempPath,
          name: 'file',
          header: {
            'Authorization': `Bearer ${getToken()}`
          },
          success: (res) => {
            const data = JSON.parse(res.data)
            if (data.code === 0) {
              resolve(data.data.url)
            } else {
              reject(new Error(data.msg))
            }
          },
          fail: reject
        })
      })
    }
  } catch (error) {
    throw error
  }
}

// 跳过设置
const handleSkip = () => {
  uni.switchTab({
    url: '/pages/profile/index'
  })
}
</script>

<style scoped>
.setup-page {
  min-height: 100vh;
}
</style> 