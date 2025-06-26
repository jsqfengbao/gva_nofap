# 任务3: 用户认证系统 - API使用说明

## 📋 概述
本文档介绍微信小程序用户认证系统的API接口使用方法和测试案例。

## 🔐 认证流程

### 1. 微信小程序登录
```http
POST /api/v1/miniprogram/auth/login
Content-Type: application/json

{
  "code": "微信wx.login()获取的code",
  "encryptedData": "加密的用户信息(可选)",
  "iv": "加密算法初始向量(可选)"
}
```

**响应:**
```json
{
  "code": 0,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "openid": "oXXXXXXXXXXXXXXXXXXX",
      "nickname": "用户昵称",
      "avatarUrl": "头像URL",
      "gender": 1,
      "city": "城市",
      "province": "省份",
      "country": "国家",
      "privacyLevel": 1,
      "status": 1,
      "lastLoginAt": "2025-01-23T20:00:00Z"
    }
  },
  "msg": "登录成功"
}
```

### 2. Token刷新
```http
POST /api/v1/miniprogram/auth/refresh
Content-Type: application/json

{
  "token": "当前的JWT token"
}
```

**响应:**
```json
{
  "code": 0,
  "data": {
    "token": "新的JWT token"
  },
  "msg": "token刷新成功"
}
```

## 🔒 需要认证的接口

所有需要认证的接口都需要在请求头中包含JWT token：

```http
Authorization: Bearer your_jwt_token_here
```

### 3. 获取用户资料
```http
GET /api/v1/miniprogram/auth/profile
Authorization: Bearer your_jwt_token_here
```

**响应:**
```json
{
  "code": 0,
  "data": {
    "user": {
      "id": 1,
      "openid": "oXXXXXXXXXXXXXXXXXXX",
      "nickname": "用户昵称",
      "avatarUrl": "头像URL",
      "gender": 1,
      "privacyLevel": 1,
      "status": 1
    },
    "abstinenceRecord": {
      "id": 1,
      "userId": 1,
      "startDate": "2025-01-23T20:00:00Z",
      "currentStreak": 0,
      "longestStreak": 0,
      "totalDays": 0,
      "successRate": 0.00,
      "level": 1,
      "experience": 0,
      "status": 1
    }
  },
  "msg": "获取成功"
}
```

### 4. 更新用户信息
```http
PUT /api/v1/miniprogram/auth/profile
Authorization: Bearer your_jwt_token_here
Content-Type: application/json

{
  "nickname": "新昵称",
  "avatarUrl": "新头像URL",
  "gender": 2
}
```

### 5. 更新隐私级别
```http
PUT /api/v1/miniprogram/auth/privacy
Authorization: Bearer your_jwt_token_here
Content-Type: application/json

{
  "privacyLevel": 2
}
```

## 🧪 测试案例

### 使用Postman/curl测试

#### 1. 测试登录（模拟数据）
```bash
curl -X POST http://localhost:8888/api/v1/miniprogram/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "code": "test_wx_code_123456"
  }'
```

#### 2. 测试获取用户资料
```bash
# 首先从登录响应中获取token，然后：
curl -X GET http://localhost:8888/api/v1/miniprogram/auth/profile \
  -H "Authorization: Bearer your_jwt_token_here"
```

#### 3. 测试更新用户信息
```bash
curl -X PUT http://localhost:8888/api/v1/miniprogram/auth/profile \
  -H "Authorization: Bearer your_jwt_token_here" \
  -H "Content-Type: application/json" \
  -d '{
    "nickname": "测试用户",
    "gender": 1
  }'
```

## ⚠️ 注意事项

### 1. 微信配置
在`server/config.yaml`中配置真实的微信小程序AppID和AppSecret：
```yaml
miniprogram:
    app-id: "你的真实AppID"
    app-secret: "你的真实AppSecret"
```

### 2. 测试环境
目前使用测试配置，微信API调用会返回模拟数据。生产环境需要：
- 配置真实的微信小程序AppID和AppSecret
- 确保服务器能访问微信API (api.weixin.qq.com)
- 配置正确的网络环境

### 3. JWT安全
- JWT密钥在生产环境应该从配置文件读取
- Token过期时间设置为7天
- 支持token刷新机制

### 4. 错误处理
API会返回标准的错误响应格式：
```json
{
  "code": -1,
  "data": null,
  "msg": "错误信息描述"
}
```

常见错误码：
- `code: 0` - 成功
- `code: -1` - 一般错误
- HTTP 401 - 认证失败
- HTTP 403 - 权限不足

## 🔄 认证流程图

```
小程序端                    后端服务
    |                         |
    |-- wx.login() ---------->|
    |<-- code ----------------|
    |                         |
    |-- POST /auth/login ---->|
    |    { code }              |-- 调用微信API -->|
    |                         |<-- session_key --|
    |<-- { token, user } -----|-- 创建/更新用户 ->|
    |                         |<-- 生成JWT ------|
    |                         |
    |-- 后续API请求 ---------->|
    |   Header: Bearer token   |-- 验证JWT ----->|
    |<-- 业务数据 ------------|<-- 返回数据 -----|
```

## 📝 开发指南

### 前端集成
参考`web/src/api/miniprogram.js`中的认证相关接口封装。

### 中间件使用
在需要认证的路由中使用`MiniprogramJWTAuth()`中间件。

### 用户信息获取
在认证后的API中，可以通过`middleware.GetUserID(c)`获取当前用户ID。 