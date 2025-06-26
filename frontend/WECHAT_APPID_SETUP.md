# 微信小程序AppID申请和配置指南

## 🚨 当前问题

**错误信息**: `invalid appid, rid: 685c06f8-560fd6dd-5a99c0f5`

**根本原因**: 当前使用的AppID `wx07c9e8e4f105260b` 是一个示例/测试AppID，不是真实有效的微信小程序AppID。

## 📋 解决步骤

### 步骤1: 申请微信小程序账号

1. **访问微信公众平台**
   - 打开 https://mp.weixin.qq.com/
   - 点击"立即注册"

2. **选择账号类型**
   - 选择"小程序"
   - 填写邮箱和密码
   - 验证邮箱

3. **完善信息**
   - 填写小程序基本信息
   - 主体信息认证（个人/企业）
   - 完成微信认证

4. **获取开发信息**
   - 登录小程序后台
   - 进入"开发" → "开发管理" → "开发设置"
   - 记录下：
     - **AppID** (小程序ID)
     - **AppSecret** (小程序密钥)

### 步骤2: 配置服务器

更新 `server/config.yaml` 文件：

```yaml
# 微信小程序配置
miniprogram:
    app-id: "你的真实AppID"           # 替换为从微信公众平台获取的AppID
    app-secret: "你的真实AppSecret"   # 替换为从微信公众平台获取的AppSecret
```

### 步骤3: 配置前端

1. **更新前端配置文件**：

```javascript
// frontend/src/config/env.js
export const THIRD_PARTY_CONFIG = {
  [ENV_TYPES.DEVELOPMENT]: {
    WECHAT_APP_ID: '你的真实AppID',  // 替换
    // ... 其他配置
  },
  [ENV_TYPES.PRODUCTION]: {
    WECHAT_APP_ID: '你的真实AppID',  // 替换
    // ... 其他配置
  },
  [ENV_TYPES.TESTING]: {
    WECHAT_APP_ID: '你的真实AppID',  // 替换
    // ... 其他配置
  }
}
```

2. **更新manifest.json**：

```json
{
  "mp-weixin": {
    "appid": "你的真实AppID"
  }
}
```

### 步骤4: 配置小程序开发者工具

1. **重新导入项目**：
   - 项目目录: `frontend/dist/dev/mp-weixin/`
   - AppID: 填写你的真实AppID
   - 项目名称: 戒色助手

2. **配置服务器域名**：
   在微信公众平台 → 开发 → 开发管理 → 服务器域名中添加：
   ```
   request合法域名: https://你的域名.com
   或开发时使用: http://localhost:8888 (仅开发环境)
   ```

## 🧪 临时测试方案

如果暂时无法申请真实AppID，可以使用以下方法进行开发测试：

### 方案1: 使用游客模式
在前端直接进入游客模式，跳过微信登录：

```javascript
// 在 frontend/src/pages/welcome/welcome.vue 中
onMounted(() => {
  // 自动进入游客模式（仅用于开发测试）
  setTimeout(() => {
    guestMode()
  }, 1000)
})
```

### 方案2: 模拟登录后端
暂时修改后端服务，返回模拟的登录成功响应：

```go
// 在 server/api/v1/miniprogram/auth.go 中添加测试模式
func (a *AuthApi) WxLogin(c *gin.Context) {
    // 开发环境下的模拟登录
    if os.Getenv("APP_ENV") == "development" {
        // 返回模拟的登录成功响应
        response.OkWithData(gin.H{
            "token": "dev_token_12345",
            "user": gin.H{
                "id": 1,
                "nickname": "测试用户",
                "avatarUrl": "",
            },
        }, c)
        return
    }
    
    // 正常的微信登录逻辑...
}
```

## 📝 配置检查清单

使用真实AppID后，确保以下文件都已更新：

- [ ] `server/config.yaml` - 后端配置
- [ ] `frontend/src/config/env.js` - 前端环境配置
- [ ] `frontend/manifest.json` - uni-app配置
- [ ] `frontend/src/manifest.json` - 源码配置
- [ ] `frontend/project.config.json` - 开发者工具配置
- [ ] `frontend/project.private.config.json` - 私有配置

## 🔄 重启服务

配置更新后需要重启：

1. **后端服务器**：
```bash
# 停止当前服务器
# 重新启动
cd server
go run main.go
```

2. **前端构建**：
```bash
cd frontend
npm run dev:mp-weixin
```

3. **微信开发者工具**：
   - 关闭项目
   - 重新导入项目（使用新的AppID）

## ⚠️ 注意事项

1. **AppSecret安全性**：
   - 不要将AppSecret提交到公开的Git仓库
   - 生产环境使用环境变量管理敏感配置

2. **域名配置**：
   - 生产环境必须在微信公众平台配置合法域名
   - 开发环境可以勾选"不校验合法域名"

3. **开发者权限**：
   - 确保开发者账号已被添加到小程序的开发者列表

## 🆘 故障排除

如果配置真实AppID后仍然失败：

1. **检查AppSecret是否正确**
2. **确认小程序状态是否正常**
3. **验证开发者权限**
4. **查看微信公众平台的错误日志**

---

**重要**: 这是开发环境的必要步骤。没有真实的AppID和AppSecret，微信登录功能无法正常工作。 