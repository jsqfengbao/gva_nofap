# 微信开发者工具导入指南

## 🎯 项目配置已优化

### ✅ 已完成的修复
1. **AppID统一配置**: 所有配置文件中的AppID已统一为 `wx07c9e8e4f105260b`
2. **配置文件迁移**: 关键配置文件已移动到frontend目录
3. **错误处理优化**: 登录失败时提供更详细的错误信息

## 📁 正确的项目结构

```
gva_NoFap/
├── frontend/                     ← 微信小程序源码目录
│   ├── project.config.json       ← 主配置文件 (AppID: wx07c9e8e4f105260b)
│   ├── project.private.config.json ← 私有配置文件
│   ├── sitemap.json              ← 页面索引配置
│   ├── manifest.json             ← uni-app配置
│   ├── src/                      ← 源代码
│   └── dist/dev/mp-weixin/       ← 构建输出 (导入此目录)
└── ...
```

## 🔧 微信开发者工具导入步骤

### 步骤 1: 构建项目
```bash
cd frontend
npm install
npm run dev:mp-weixin
```

### 步骤 2: 打开微信开发者工具
1. 启动微信开发者工具
2. 选择 "导入项目"

### 步骤 3: 配置项目信息
- **项目目录**: 选择 `frontend/dist/dev/mp-weixin/`
- **AppID**: `wx07c9e8e4f105260b`
- **项目名称**: `戒色助手` (可自定义)

### 步骤 4: 开发设置
在微信开发者工具中：
1. 点击右上角"详情"
2. 选择"本地设置"标签
3. 勾选以下选项：
   - ✅ 不校验合法域名、web-view(业务域名)、TLS版本以及HTTPS证书
   - ✅ 启用调试
   - ✅ 启用ES6转ES5
   - ✅ 启用增强编译

## 🐛 常见问题排查

### 问题1: "获取登录凭证失败"
**原因**: AppID配置问题或小程序未激活
**解决**: 
1. 确认AppID `wx07c9e8e4f105260b` 的有效性
2. 检查微信公众平台中小程序状态
3. 确认开发者权限

### 问题2: 项目导入失败
**原因**: 导入目录错误
**解决**: 确保导入的是 `frontend/dist/dev/mp-weixin/` 目录

### 问题3: 构建失败
**原因**: 依赖缺失
**解决**:
```bash
cd frontend
rm -rf node_modules
npm install
npm run dev:mp-weixin
```

## 🧪 测试登录功能

### 方法1: 控制台测试
在微信开发者工具的Console中运行：
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

### 方法2: 应用内测试
1. 启动小程序
2. 进入欢迎页面
3. 点击"微信授权登录"按钮
4. 查看控制台输出

### 方法3: 游客模式测试
如果微信登录暂时无法使用：
1. 点击"游客模式体验"按钮
2. 可以测试除登录外的其他功能

## 📝 配置文件说明

### project.config.json
- 微信小程序主配置文件
- 包含AppID、编译设置等
- AppID: `wx07c9e8e4f105260b`

### project.private.config.json  
- 私有配置文件
- 包含个人开发设置
- 不会被版本控制

### sitemap.json
- 页面索引配置
- 控制哪些页面可被微信搜索

### manifest.json
- uni-app框架配置
- 包含各平台的特定设置

## 🔄 开发流程

1. **修改源码**: 在 `frontend/src/` 目录下开发
2. **实时预览**: 运行 `npm run dev:mp-weixin` 保持构建
3. **查看效果**: 在微信开发者工具中实时查看
4. **调试问题**: 使用开发者工具的调试功能

## 📞 技术支持

如果遇到问题，请查看：
1. `frontend/WECHAT_LOGIN_DEBUG.md` - 详细的登录问题排查指南
2. 微信开发者工具的控制台输出
3. `frontend/src/utils/auth.js` - 登录相关代码

---

**注意**: 确保使用最新版本的微信开发者工具，建议版本 ≥ 1.06.x
