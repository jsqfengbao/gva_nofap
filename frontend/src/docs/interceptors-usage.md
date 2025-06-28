# 统一拦截器系统使用指南

## 概述

本项目实现了统一的拦截器系统，包括：
- **请求拦截器** - 处理API请求的认证、错误处理和重试逻辑
- **路由拦截器** - 处理登录拦截和页面访问权限控制
- **原型拦截器** - 解决低版本手机兼容性问题

## 快速开始

### 1. 自动安装

拦截器系统在应用启动时自动安装，无需手动配置：

```javascript
// main.ts 中已自动安装
import { installInterceptors } from './interceptors/index.js'
installInterceptors()
```

### 2. 使用新的API接口

#### 基础用法

```javascript
// 导入API
import { checkinApi, userApi, communityApi } from '@/apis/index.js'

// 或者导入具体的请求方法
import { authGet, authPost, guestGet } from '@/utils/request.js'

// 使用API（自动处理认证和错误）
const res = await checkinApi.getTodayStatus()
const userProfile = await userApi.getProfile()
```

#### 请求类型

系统支持三种请求类型：

1. **普通请求** (`REQUEST_TYPES.NORMAL`) - 默认类型
2. **需要认证** (`REQUEST_TYPES.AUTH_REQUIRED`) - 需要登录token
3. **允许游客** (`REQUEST_TYPES.GUEST_ALLOWED`) - 游客可访问

```javascript
// 需要认证的请求
const profile = await authGet('/user/profile')

// 允许游客的请求  
const articles = await guestGet('/learning/articles')

// 普通请求
const config = await get('/app/config')
```

## 各模块改造示例

### 首页模块

```javascript
// 原来的写法
const loadUserData = async () => {
  const token = uni.getStorageSync('token')
  const res = await uni.request({
    url: 'http://localhost:8888/api/v1/miniprogram/user/profile',
    method: 'GET',
    header: { 'Authorization': `Bearer ${token}` }
  })
}

// 新的写法
const loadUserData = async () => {
  try {
    const res = await userApi.getProfile()
    // 认证、错误处理、URL构建都由拦截器自动处理
  } catch (error) {
    // 错误已由拦截器处理，这里只需记录日志
    console.error('加载用户数据失败:', error)
  }
}
```

### 打卡模块

```javascript
// 原来的写法
const submitCheckin = async () => {
  const token = uni.getStorageSync('token')
  const res = await uni.request({
    url: 'http://localhost:8888/api/v1/miniprogram/checkin/daily',
    method: 'POST',
    header: { 'Authorization': `Bearer ${token}` },
    data: { moodLevel: selectedMood.value }
  })
}

// 新的写法
const submitCheckin = async () => {
  try {
    const res = await checkinApi.dailyCheckin({
      moodLevel: selectedMood.value,
      notes: checkinNotes.value
    })
    // 自动处理认证、URL构建、错误提示
  } catch (error) {
    // 错误已处理，可添加额外逻辑
  }
}
```

### 社区模块

```javascript
// 原来的写法
const loadPosts = async () => {
  const token = uni.getStorageSync('token')
  if (!token) {
    uni.showToast({ title: '请先登录' })
    return
  }
  
  const res = await uni.request({
    url: 'http://localhost:8888/api/v1/miniprogram/community/posts',
    method: 'GET',
    header: { 'Authorization': `Bearer ${token}` }
  })
}

// 新的写法
const loadPosts = async () => {
  try {
    const res = await communityApi.getPosts()
    // 登录检查、认证、错误处理都自动完成
  } catch (error) {
    // 如果未登录，拦截器会自动处理登录流程
  }
}
```

### 学习模块

```javascript
// 新的写法 - 支持游客访问
const loadLearningList = async () => {
  try {
    // 游客可以浏览学习内容列表
    const res = await learningApi.getList()
  } catch (error) {
    console.error('加载学习内容失败:', error)
  }
}

// 需要登录的功能
const recordProgress = async (contentId, progress) => {
  try {
    // 自动检查登录状态，未登录会提示登录
    const res = await learningApi.recordProgress(contentId, { progress })
  } catch (error) {
    // 错误已处理
  }
}
```

## 路由拦截配置

### 需要登录的页面

在 `frontend/src/interceptors/route.js` 中配置：

```javascript
const NEED_LOGIN_PAGES = [
  '/pages/profile/auth',
  '/pages/profile/setup', 
  '/pages/profile/privacy',
  '/pages/checkin/index',
  '/pages/checkin/history',
  '/pages/community/index',
  '/pages/community/post',
  '/pages/learning/detail',
  '/pages/progress/index'
]
```

### 游客可访问的页面

```javascript
const GUEST_ALLOWED_PAGES = [
  '/pages/index/index',
  '/pages/welcome/welcome',
  '/pages/learning/index',
  '/pages/emergency/index',
  '/pages/about/index'
]
```

## 错误处理

### 自动处理的错误

1. **认证错误** (401) - 自动清除token，跳转登录页
2. **网络错误** - 显示网络错误提示
3. **超时错误** - 显示超时提示
4. **业务错误** - 显示服务器返回的错误信息

### 自定义错误处理

```javascript
// 隐藏默认错误提示
const res = await authGet('/api/data', {}, { 
  hideErrorToast: true 
})

// 自定义处理
if (res.data.code !== 0) {
  // 自定义错误处理逻辑
  uni.showModal({
    title: '操作失败',
    content: res.data.msg
  })
}
```

## 文件上传/下载

```javascript
import { uploadFile, downloadFile } from '@/utils/request.js'

// 文件上传（自动处理认证）
const uploadAvatar = async (filePath) => {
  try {
    const res = await uploadFile('/user/upload-avatar', filePath)
    return res.data.url
  } catch (error) {
    console.error('上传失败:', error)
  }
}

// 文件下载
const downloadDocument = async (url) => {
  try {
    const res = await downloadFile(url)
    return res.tempFilePath
  } catch (error) {
    console.error('下载失败:', error)
  }
}
```

## 调试和日志

拦截器系统提供详细的日志输出：

```
🚀 开始安装拦截器...
✅ 原型拦截器安装完成
✅ 请求拦截器安装完成
🚀 初始化启动拦截器...
✅ 用户已登录
✅ 路由拦截器安装完成
🎉 所有拦截器安装完成

📤 API请求 [GET_today_1640995200000]: {
  url: "http://192.168.1.140:8888/api/v1/miniprogram/checkin/today",
  method: "GET",
  hasAuth: true
}

📥 API响应 [GET_today_1640995200000]: {
  status: 200,
  success: true
}
```

## 迁移指南

### 1. 替换导入

```javascript
// 旧的导入
import { request } from '@/utils/auth'
import { userApi } from '@/utils/api'

// 新的导入
import { userApi, checkinApi } from '@/apis/index.js'
```

### 2. 移除手动认证

```javascript
// 旧的代码 - 需要手动处理token
const token = uni.getStorageSync('token')
const res = await uni.request({
  url: buildApiUrl('/user/profile'),
  method: 'GET',
  header: { 'Authorization': `Bearer ${token}` }
})

// 新的代码 - 自动处理认证
const res = await userApi.getProfile()
```

### 3. 简化错误处理

```javascript
// 旧的代码 - 需要手动处理各种错误
try {
  const res = await uni.request(...)
  if (res.statusCode === 401) {
    // 处理认证错误
  } else if (res.statusCode !== 200) {
    // 处理网络错误
  } else if (res.data.code !== 0) {
    // 处理业务错误
  }
} catch (error) {
  // 处理异常
}

// 新的代码 - 错误自动处理
try {
  const res = await userApi.getProfile()
  // 只需处理成功的情况
} catch (error) {
  // 错误已自动处理，这里只需记录日志
}
```

## 注意事项

1. **向后兼容** - 旧的API调用方式仍然可用，但建议逐步迁移
2. **错误处理** - 大部分错误已自动处理，减少重复代码
3. **登录拦截** - 访问需要登录的页面会自动提示登录
4. **性能优化** - 请求去重、自动重试等功能提升用户体验
5. **调试信息** - 开发环境下提供详细的请求日志

## 常见问题

### Q: 如何禁用某个请求的自动错误提示？
A: 在请求选项中设置 `hideErrorToast: true`

### Q: 如何添加新的需要登录的页面？
A: 在 `frontend/src/interceptors/route.js` 的 `NEED_LOGIN_PAGES` 数组中添加页面路径

### Q: 如何处理特殊的业务错误？
A: 使用 `hideErrorToast: true` 禁用自动提示，然后自定义处理逻辑

### Q: 拦截器会影响现有代码吗？
A: 不会，拦截器是增强性的，现有代码可以正常运行 