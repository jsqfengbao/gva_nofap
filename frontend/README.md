# 戒色助手微信小程序前端

基于 uni-app + Vue 3 + TypeScript 开发的微信小程序前端项目。

## 开发环境

- Node.js >= 16.0.0
- HBuilderX 最新版本
- 微信开发者工具

## 项目结构

```
frontend/
├── src/                    # 源代码目录
│   ├── components/         # 组件目录
│   ├── pages/             # 页面目录
│   ├── static/            # 静态资源目录
│   ├── styles/            # 样式目录
│   ├── utils/             # 工具函数目录
│   ├── App.vue            # 应用入口
│   ├── main.ts            # 主入口文件
│   ├── manifest.json      # 应用配置
│   └── pages.json         # 页面配置
├── dist/                  # 编译输出目录
├── .hbuilderx/           # HBuilderX 配置
├── vite.config.ts        # Vite 配置
├── vue.config.js         # Vue/Webpack 配置
└── package.json          # 项目依赖
```

## 在 HBuilder 中的使用方法

### 1. 导入项目

1. 打开 HBuilderX
2. 点击 `文件` -> `导入` -> `从本地目录导入`
3. 选择 `frontend` 目录
4. 点击 `导入`

### 2. 安装依赖

在 HBuilderX 终端中运行：

```bash
npm install
```

### 3. 运行到微信小程序

#### 方法一：使用 HBuilderX 界面

1. 在项目管理器中右键点击项目
2. 选择 `运行` -> `运行到小程序模拟器` -> `微信开发者工具`
3. 首次运行需要配置微信开发者工具路径

#### 方法二：使用终端命令

```bash
# 开发环境
npm run dev:mp-weixin

# 或者使用专门的 HBuilder 命令
npm run dev:mp-weixin:hbuilder

# 开发环境（带文件监听，文件变化时自动重新构建）
npm run watch:mp-weixin

# 生产环境构建
npm run build:mp-weixin:hbuilder
```

### 4. 在微信开发者工具中查看

1. 打开微信开发者工具
2. 点击 `导入项目`
3. 选择 `frontend/dist/dev/mp-weixin` 目录
4. 填写 AppID（测试阶段可以使用测试号）
5. 点击 `导入`

## 可用脚本

```bash
# 开发命令
npm run dev:mp-weixin          # 微信小程序开发模式
npm run dev:h5                 # H5 开发模式
npm run dev:app                # App 开发模式

# 构建命令
npm run build:mp-weixin        # 微信小程序生产构建
npm run build:h5               # H5 生产构建
npm run build:app              # App 生产构建

# HBuilder 专用命令
npm run dev:mp-weixin:hbuilder    # HBuilder 微信小程序开发
npm run build:mp-weixin:hbuilder  # HBuilder 微信小程序构建

# 监听模式
npm run watch:mp-weixin           # 监听文件变化，自动重新构建

# 清理命令
npm run clean                     # 清理所有构建文件
npm run clean:mp-weixin          # 清理微信小程序构建文件
```

## 配置说明

### vite.config.ts 主要配置

- **服务器配置**：端口 3000，支持热重载
- **别名配置**：`@` 指向 `src` 目录
- **构建优化**：针对微信小程序优化的构建配置
- **资源处理**：图片、字体等静态资源的处理规则

### vue.config.js 配置

- **copy-webpack-plugin**：自动复制静态资源和配置文件
- **资源处理**：图片压缩和优化
- **微信小程序特殊配置**：针对小程序平台的特殊处理

### manifest.json 配置

- **微信小程序设置**：ES6 转换、代码压缩等
- **权限配置**：位置信息等权限申请
- **优化配置**：分包优化等性能配置

## 注意事项

1. **微信小程序限制**：
   - 不支持某些 CSS 特性（如通用选择器 `*`）
   - 不支持外部资源引用
   - 包大小限制（主包 2MB，分包 2MB）

2. **开发建议**：
   - 使用 `rpx` 单位进行布局
   - 遵循微信小程序的组件和API规范
   - 注意页面栈深度限制（最多10层）

3. **调试技巧**：
   - 使用微信开发者工具的调试面板
   - 查看 Console 输出定位问题
   - 使用网络面板检查API调用

## 常见问题

### Q: WXSS 编译错误怎么解决？
A: 检查 CSS 语法是否符合微信小程序规范，避免使用不支持的 CSS 特性。

### Q: 静态资源找不到？
A: 确保资源放在 `src/static` 目录中，构建时会自动复制到输出目录。

### Q: 热重载不生效？
A: 检查端口是否被占用，或者重启开发服务器。

### Q: HBuilderX报错"node_modules缺少编译器模块"？
A: 请参考 [HBuilderX配置指南](./HBUILDER_SETUP.md) 获得详细的解决方案。

## 技术栈

- **框架**：uni-app
- **前端框架**：Vue 3
- **编程语言**：TypeScript
- **样式处理**：SCSS
- **状态管理**：Pinia
- **构建工具**：Vite
- **开发工具**：HBuilderX

## 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交代码
4. 推送到分支
5. 创建 Pull Request 