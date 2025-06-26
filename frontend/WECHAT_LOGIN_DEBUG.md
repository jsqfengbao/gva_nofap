# 微信登录问题排查指南

## 问题描述
前端登录请求报错：微信登录失败: Error: 获取登录凭证失败

## 已修复的问题

### 1. AppID配置不一致 ✅
**问题**: `project.config.json` 中的AppID与其他配置文件不一致
- 根目录 `project.config.json`: `wx2cda45f236177104` ❌
- 其他配置文件: `wx07c9e8e4f105260b` ✅

**解决方案**: 已将根目录 `project.config.json` 的AppID统一为 `wx07c9e8e4f105260b`

## 需要检查的问题

### 2. 微信开发者工具登录状态 🆕

**常见错误**: `login:fail 需要重新登录,access_token missing`

**解决方案**:
1. **完全退出微信开发者工具**
   - 关闭所有窗口
   - 确保进程完全退出
2. **重新启动并登录**
   - 重新打开微信开发者工具
   - 使用微信扫码重新登录
   - 确保登录成功
3. **使用测试账号功能**
   - 点击右上角头像
   - 选择"使用测试号"
   - 重新尝试

### 3. 微信开发者工具配置
请确保：
- [ ] 在微信开发者工具中打开的项目目录是 `frontend/dist/dev/mp-weixin/`
- [ ] AppID设置为 `wx07c9e8e4f105260b`
- [ ] 在微信开发者工具的"详情"→"本地设置"中：
  - [ ] 勾选"不校验合法域名、web-view(业务域名)、TLS版本以及HTTPS证书"
  - [ ] 勾选"启用调试"

### 4. AppID有效性检查
请确认：
- [ ] AppID `wx07c9e8e4f105260b` 是否为您拥有的有效小程序AppID
- [ ] 该小程序是否已在微信公众平台完成基本配置
- [ ] 开发者是否已被添加到该小程序的开发者列表中

### 5. 网络环境
请检查：
- [ ] 网络连接正常
- [ ] 防火墙是否阻止了微信开发者工具的网络请求
- [ ] 代理设置是否正确

## 调试步骤

### 步骤1: 重新构建项目
```bash
cd frontend
npm run dev:mp-weixin
```

### 步骤2: 在微信开发者工具中重新导入项目
1. 关闭微信开发者工具
2. 重新打开微信开发者工具
3. 选择"导入项目"
4. 项目目录选择：`frontend/dist/dev/mp-weixin/`
5. AppID填写：`wx07c9e8e4f105260b`

### 步骤3: 检查控制台输出
在微信开发者工具的"调试器"→"Console"中查看详细错误信息

### 步骤4: 测试基础登录功能

#### 方法1: 控制台测试
在微信开发者工具的"Console"中运行：
```javascript
wx.login({
  success: (res) => {
    console.log('登录成功:', res)
  },
  fail: (err) => {
    console.error('登录失败:', err)
  }
})
```

#### 方法2: 使用测试页面
1. 将 `frontend/test_wx_login.html` 复制到 `frontend/dist/dev/mp-weixin/` 目录
2. 在微信开发者工具中访问此页面
3. 点击"测试 wx.login"按钮
4. 查看详细的测试结果

## 常见解决方案

### 方案1: 更换为测试AppID
如果当前AppID有问题，可以临时使用微信提供的测试AppID：
```
测试AppID: wx07c9e8e4f105260b
```
**注意**: 这个AppID只能在开发环境使用，不能发布到生产环境。

### 方案2: 申请新的小程序AppID
1. 访问 [微信公众平台](https://mp.weixin.qq.com/)
2. 注册新的小程序账号
3. 获取新的AppID
4. 更新所有配置文件中的AppID

### 方案3: 使用游客模式进行测试
如果微信登录暂时无法解决，可以先使用游客模式测试其他功能：
- 在欢迎页面点击"游客模式体验"按钮
- 这不会调用微信登录接口，可以绕过登录问题

## 配置文件清单

需要确保以下文件中的AppID一致：
- ✅ `project.config.json` (根目录)
- ✅ `frontend/manifest.json`
- ✅ `frontend/src/manifest.json`
- ✅ `frontend/src/config/env.js`
- ✅ `frontend/dist/dev/mp-weixin/project.config.json`
- ✅ `server/config.yaml`

## 后续排查

如果以上步骤都无法解决问题，请提供：
1. 微信开发者工具的完整错误日志
2. 网络环境信息（是否使用代理、防火墙设置）
3. 确认的AppID所有权信息
4. 微信开发者工具版本号

## 临时绕过方案

在登录问题修复期间，可以修改 `frontend/src/pages/welcome/welcome.vue`，默认进入游客模式：

```javascript
// 在 onMounted 中添加
onMounted(() => {
  // 临时跳过登录检查，直接进入游客模式
  // checkLoginStatus()
  
  // 自动进入游客模式（仅用于调试）
  setTimeout(() => {
    guestMode()
  }, 1000)
})
```

**记住在修复登录问题后要恢复正常的登录流程！** 