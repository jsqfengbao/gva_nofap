# 微信小程序登录功能说明

## 问题背景

由于微信小程序隐私政策的变化，从2022年开始：

1. **`uni.getUserProfile()` 已被弃用**
2. **`getUserInfo` 的 open-type 也不再推荐使用**
3. **用户头像和昵称需要通过专门的授权组件获取**

## 当前解决方案

### 1. 基础登录（只获取openid）

用户点击登录按钮后：
- 调用 `uni.login()` 获取临时凭证 `code`
- 将 `code` 发送到后端，后端调用微信API获取 `openid` 和 `session_key`
- 后端创建用户记录，生成JWT token
- 前端保存token，用户登录成功

**特点：**
- ✅ 符合微信最新政策
- ✅ 用户无需额外授权
- ⚠️ 用户显示为默认昵称和头像

### 2. 完善个人信息（可选）

登录成功后，系统会询问用户是否完善个人信息：

#### 头像设置
```vue
<button 
  class="avatar-btn"
  open-type="chooseAvatar"
  @chooseavatar="onChooseAvatar"
>
  <!-- 头像显示区域 -->
</button>
```

#### 昵称设置
```vue
<input 
  v-model="userInfo.nickname"
  type="nickname"
  placeholder="请输入昵称"
  maxlength="20"
/>
```

**特点：**
- ✅ 符合微信最新API规范
- ✅ 用户完全自主选择
- ✅ 键盘会显示微信昵称建议

## 技术实现

### 前端流程

1. **基础登录**：
   ```javascript
   // 获取微信登录凭证
   const result = await wxLogin()
   // 登录成功，显示默认信息
   ```

2. **完善信息**（可选）：
   ```javascript
   // 用户选择头像
   const onChooseAvatar = (e) => {
     userInfo.avatarUrl = e.detail.avatarUrl
   }
   
   // 保存到服务器
   await userApi.updateUserInfo({
     nickname: userInfo.nickname,
     avatarUrl: finalAvatarUrl
   })
   ```

### 后端处理

1. **微信登录API** (`/miniprogram/auth/wx-login`)：
   - 接收前端传来的 `code`
   - 调用微信API获取 `openid` 和 `session_key`
   - 查找或创建用户记录
   - 返回JWT token和用户信息

2. **用户信息更新API** (`/miniprogram/user/info`)：
   - 接收昵称和头像信息
   - 验证用户身份
   - 更新数据库中的用户信息

3. **微信头像保存API** (`/miniprogram/user/save-wx-avatar`)：
   - 接收微信临时头像文件
   - 上传到服务器存储
   - 返回永久头像URL

## 用户体验流程

### 场景1：快速体验用户
1. 用户点击"微信快速登录"
2. 微信授权登录（无需额外操作）
3. 登录成功，显示默认头像和"微信用户"昵称
4. 系统询问是否完善个人信息
5. 用户选择"暂不"，直接使用应用

### 场景2：完整体验用户
1. 用户点击"微信快速登录"
2. 微信授权登录成功
3. 系统询问是否完善个人信息
4. 用户选择"去设置"
5. 在设置页面选择头像、输入昵称
6. 保存成功，返回个人中心显示真实信息

## 配置要求

### 1. 微信小程序配置

在 `app.json` 中正确配置权限：

```json
{
  "permission": {
    "scope.userLocation": {
      "desc": "你的位置信息将用于展示附近的相关内容"
    }
  },
  "requiredPrivateInfos": ["getLocation"],
  "lazyCodeLoading": "requiredComponents"
}
```

**注意：** `requiredPrivateInfos` 只能包含地理位置相关的API，不包括 `chooseAvatar`。

### 2. 网络配置

确保API地址配置正确：

```javascript
// frontend/src/config/env.js
export const API_DOMAINS = {
  [ENV_TYPES.DEVELOPMENT]: {
    BASE_URL: 'http://192.168.1.140:8888',  // 真机调试用局域网IP
    WS_URL: 'ws://192.168.1.140:8888'
  }
}
```

### 3. 后端CORS配置

在 `server/config.yaml` 中添加局域网IP的CORS支持：

```yaml
cors:
  mode: allow-all
  whitelist:
    - allow-origin: http://192.168.1.140:3000
      allow-methods: POST, GET, PUT, DELETE, OPTIONS
      allow-headers: Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id
      expose-headers: Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type
      allow-credentials: true
```

## 常见问题

### Q1: 用户登录后显示"微信用户"和默认头像？
**A:** 这是正常现象。微信现在只允许通过专门的授权组件获取用户信息，基础登录只能获取openid。

### Q2: 点击登录没有弹出授权窗口？
**A:** 现在的微信登录不会弹出获取用户信息的授权窗口，只会进行身份验证。

### Q3: 如何获取用户真实头像和昵称？
**A:** 需要用户主动在"完善个人信息"页面进行设置，使用 `open-type="chooseAvatar"` 和 `type="nickname"` 组件。

### Q4: 真机调试无法连接服务器？
**A:** 运行 `./scripts/update-network-config.sh` 脚本自动更新网络配置。

## 更新日志

- **2025-06-26**: 修复微信登录API路径问题，统一使用 `/miniprogram/*` 前缀
- **2025-06-26**: 移除 `app.json` 中无效的 `chooseAvatar` 配置
- **2025-06-26**: 更新登录流程，符合微信最新政策
- **2025-06-26**: 添加局域网IP自动配置脚本 