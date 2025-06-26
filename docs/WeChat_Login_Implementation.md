# 微信小程序登录实现说明

## 概述

本文档说明了NoFap戒色助手小程序的微信登录功能实现，从传统的账号密码登录方式改为微信官方登录方式。

## 实现内容

### 1. 前端页面改造

#### 1.1 登录页面 (`frontend/src/pages/auth/login.vue`)
- 移除了传统的手机号和密码输入框
- 添加了微信授权登录按钮，支持获取用户信息
- 添加了一键登录按钮作为兜底方案
- 保留了游客模式入口
- 优化了UI设计，增加了登录说明和用户协议链接

#### 1.2 欢迎页面 (`frontend/src/pages/welcome/welcome.vue`)
- 设计了精美的欢迎界面，包含功能介绍
- 集成了微信登录功能
- 添加了动态效果和加载状态

#### 1.3 配置文件更新
- 更新 `manifest.json`，添加了微信小程序所需的权限配置
- 添加了 `scope.userInfo` 权限和 `requiredPrivateInfos` 配置

### 2. 后端API支持

#### 2.1 微信登录API (`server/api/v1/miniprogram/auth.go`)
- `WxLogin` 方法：处理微信登录凭证，支持用户信息解密
- 已集成微信API调用，获取openid和session_key
- 支持用户信息解密（encryptedData + iv）
- 自动创建或更新用户信息

#### 2.2 认证服务 (`server/service/miniprogram/auth_service.go`)
- `GetWxSessionInfo`：调用微信API获取会话信息
- `DecryptWxUserInfo`：解密微信用户信息
- `GenerateToken`：生成JWT token
- 完整的加密解密处理

#### 2.3 用户服务 (`server/service/miniprogram/user_service.go`)
- `FindOrCreateWxUser`：查找或创建微信用户
- 自动为新用户创建戒色记录和默认设置
- 支持用户信息更新

### 3. 工具函数和配置

#### 3.1 认证工具 (`frontend/src/utils/auth.js`)
- `isLoggedIn()`：检查登录状态
- `wxLogin()`：微信登录封装
- `request()`：API请求封装，自动携带token
- `logout()`：登出功能
- 完整的token管理和用户信息处理

#### 3.2 应用配置 (`frontend/src/config/index.js`)
- 环境配置管理（开发/生产）
- API地址配置
- 存储键名统一管理
- 主题和功能开关配置

### 4. 中间件和安全

#### 4.1 JWT认证中间件 (`server/middleware/miniprogram_jwt.go`)
- 专门为小程序设计的JWT认证中间件
- 支持Bearer token和x-token两种方式
- IP失败次数限制
- token黑名单支持

## 使用流程

### 用户登录流程
1. 用户打开小程序，首次启动显示欢迎页面
2. 用户点击"微信授权登录"按钮
3. 小程序调用 `uni.login()` 获取登录凭证code
4. 如果用户同意授权，获取encryptedData和iv
5. 发送code（和用户信息）到后端API
6. 后端调用微信API验证code，获取openid
7. 后端解密用户信息（如果有），查找或创建用户
8. 返回JWT token和用户信息
9. 前端保存token，跳转到首页

### 兜底方案
- 如果用户拒绝授权，提供"一键登录"仅使用openid创建用户
- 提供"游客模式"供用户体验基本功能

## 配置要求

### 微信小程序配置
1. 在微信公众平台配置小程序信息
2. 获取AppID和AppSecret
3. 在 `server/config.yaml` 中配置：
```yaml
miniprogram:
    app-id: "your_real_app_id"
    app-secret: "your_real_app_secret"
```

### 服务器域名配置
需要在微信公众平台配置服务器域名：
- request合法域名：添加你的API服务器域名
- 如：`https://your-api-domain.com`

## 安全特性

1. **JWT Token**：使用安全的JWT token进行认证
2. **用户信息加密**：支持微信用户信息解密
3. **IP限制**：防止暴力破解，IP失败次数限制
4. **Token黑名单**：支持token撤销
5. **HTTPS**：生产环境强制使用HTTPS

## 测试建议

### 开发环境测试
1. 确保后端服务启动（localhost:8888）
2. 配置测试用的微信AppID和AppSecret
3. 在微信开发者工具中测试小程序

### 功能测试点
- [ ] 微信授权登录流程
- [ ] 用户信息获取和显示
- [ ] Token自动携带
- [ ] 登录状态检查
- [ ] 游客模式功能
- [ ] 登出功能
- [ ] 网络错误处理

## 后续优化建议

1. **缓存优化**：添加用户信息本地缓存
2. **离线支持**：基本功能离线可用
3. **推送通知**：集成微信模板消息
4. **数据同步**：多设备数据同步
5. **性能监控**：添加性能和错误监控

## 故障排查

### 常见问题
1. **登录失败**：检查AppID和AppSecret配置
2. **网络错误**：检查API服务器是否启动
3. **token失效**：检查JWT配置和过期时间
4. **用户信息为空**：检查用户授权流程

### 调试方法
1. 查看控制台日志输出
2. 检查网络请求和响应
3. 验证微信API调用结果
4. 检查数据库用户记录

## 部署注意事项

1. 生产环境需要配置真实的微信AppID和AppSecret
2. 确保API服务器域名已在微信平台配置
3. 使用HTTPS协议
4. 配置正确的环境变量
5. 数据库迁移和初始化

---

*该文档记录了微信小程序登录功能的完整实现，包括前端页面、后端API、安全机制等所有相关内容。* 