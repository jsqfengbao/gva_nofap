# 戒色助手 App 设计规范文档

## 1. 文档信息

**项目名称**: 戒色助手 (NoFap Helper)  
**版本**: v1.0  
**创建日期**: 2024-01-15  
**设计平台**: iOS 优先，兼容 Android/Web/小程序  
**技术栈**: HTML5 + Tailwind CSS + FontAwesome  

## 2. 设计理念与原则

### 2.1 核心设计理念
- **温暖友好**: 使用温暖的色彩和友好的界面语言，减少用户的心理压力
- **私密安全**: 界面设计体现隐私保护，让用户感到安全和信任
- **游戏化视觉**: 通过游戏化元素增强用户参与度和成就感
- **简洁明了**: 信息层次清晰，操作路径简单，减少认知负担

### 2.2 设计原则
1. **以用户为中心**: 所有设计决策都以用户需求和体验为出发点
2. **一致性**: 保持整个应用的视觉和交互一致性
3. **可访问性**: 确保不同能力用户都能顺利使用
4. **响应式**: 适配不同屏幕尺寸和设备类型
5. **渐进增强**: 核心功能优先，逐步增加高级特性

## 3. 色彩系统

### 3.1 主色调
```css
primary: '#34D399'    /* 温暖的翠绿色 - 主要CTA按钮、强调元素 */
secondary: '#06B6D4'  /* 清新的天蓝色 - 次要按钮、信息提示 */
accent: '#F59E0B'     /* 活力的琥珀色 - 等级、奖励系统 */
success: '#10B981'    /* 成功绿色 - 成功状态、正向反馈 */
danger: '#EF4444'     /* 警告红色 - 紧急求助、错误状态 */
```

### 3.2 中性色调
```css
background: '#F8FAFC'  /* 页面背景色 */
surface: '#FFFFFF'     /* 卡片、弹窗背景色 */
text-primary: '#1F2937'    /* 主要文字颜色 */
text-secondary: '#6B7280'  /* 次要文字颜色 */
```

### 3.3 渐变色方案
- **主要渐变**: `from-primary to-success` (翠绿到深绿)
- **次要渐变**: `from-secondary to-blue-600` (天蓝到深蓝)
- **等级渐变**: `from-accent to-orange-500` (琥珀到橙色)
- **警告渐变**: `from-purple-400 to-pink-500` (紫到粉)

### 3.4 色彩使用规范
- **主色**: 用于主要行动按钮、导航激活状态、进度指示
- **次色**: 用于次要功能、信息展示、辅助图标
- **强调色**: 用于游戏化元素、等级显示、奖励系统
- **成功色**: 用于完成状态、正向反馈、成就展示
- **警告色**: 仅用于紧急情况、错误提示、重要警告

## 4. 字体系统

### 4.1 字体家族
```css
font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', system-ui, sans-serif;
```

### 4.2 字体层级
- **超大标题**: text-3xl (30px) - 页面主标题
- **大标题**: text-2xl (24px) - 卡片标题、重要信息
- **中标题**: text-xl (20px) - 副标题、导航标题
- **小标题**: text-lg (18px) - 内容标题、分组标题
- **正文**: text-base (16px) - 正常文本内容
- **小文本**: text-sm (14px) - 辅助信息、说明文字
- **超小文本**: text-xs (12px) - 标签、时间戳

### 4.3 字重规范
- **粗体**: font-bold (700) - 标题、重要信息
- **中粗**: font-semibold (600) - 次级标题、强调文字
- **常规**: font-medium (500) - 按钮文字、导航项
- **正常**: font-normal (400) - 正文内容

## 5. 间距系统

### 5.1 间距规范
- **超小间距**: 1 (4px) - 紧密元素间距
- **小间距**: 2-3 (8-12px) - 相关元素间距
- **常规间距**: 4-6 (16-24px) - 组件间距、内边距
- **大间距**: 8-12 (32-48px) - 模块间距、外边距
- **超大间距**: 16-20 (64-80px) - 页面级间距

### 5.2 布局间距
- **页面边距**: px-6 (24px) - 页面水平边距
- **卡片内边距**: p-4/p-6 (16px/24px) - 卡片内容间距
- **组件间距**: space-y-4/space-y-6 (16px/24px) - 垂直组件间距
- **按钮内边距**: py-3 px-4 (12px 16px) - 按钮内部间距

## 6. 圆角系统

### 6.1 圆角规范
- **小圆角**: rounded-lg (8px) - 按钮、输入框
- **中圆角**: rounded-xl (12px) - 卡片内元素、图标容器
- **大圆角**: rounded-2xl (16px) - 主要卡片、弹窗
- **超大圆角**: rounded-3xl (24px) - 特殊卡片、主要容器
- **圆形**: rounded-full - 头像、图标按钮、进度指示器

## 7. 阴影系统

### 7.1 阴影层级
- **基础阴影**: shadow-sm - 卡片、输入框
- **中等阴影**: shadow-md - 悬浮元素、弹窗
- **强阴影**: shadow-lg - 重要弹窗、模态框
- **特殊阴影**: 自定义渐变阴影用于特殊效果

### 7.2 阴影使用场景
- **卡片容器**: 使用 shadow-sm 提供轻微层次感
- **悬浮按钮**: 使用 shadow-lg 强调可交互性
- **弹窗模态**: 使用 shadow-xl 突出重要性
- **特效元素**: 使用自定义阴影营造氛围

## 8. 图标系统

### 8.1 图标库
**主要图标库**: FontAwesome 6.4.0  
**CDN链接**: https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css

### 8.2 图标分类与用法
#### 导航图标
- `fa-home` - 主页
- `fa-chart-line` - 进度追踪
- `fa-users` - 社区互助
- `fa-book` - 学习内容
- `fa-user` - 个人中心

#### 功能图标
- `fa-fire` - 连击天数、热门内容
- `fa-trophy` - 成就、奖励
- `fa-heart` - 点赞、收藏
- `fa-comment` - 评论、交流
- `fa-share` - 分享功能

#### 状态图标
- `fa-check` - 完成、成功
- `fa-times` - 取消、错误
- `fa-exclamation-triangle` - 警告、紧急
- `fa-info-circle` - 信息提示
- `fa-question-circle` - 帮助、疑问

#### 游戏化图标
- `fa-star` - 等级、评级
- `fa-medal` - 成就徽章
- `fa-crown` - 高级用户、导师
- `fa-gamepad` - 游戏化元素

### 8.3 图标尺寸规范
- **小图标**: text-sm (14px) - 内联图标、状态指示
- **常规图标**: text-base (16px) - 按钮图标、导航图标
- **中图标**: text-lg (18px) - 卡片标题图标
- **大图标**: text-xl/text-2xl (20px/24px) - 主要功能图标
- **超大图标**: text-3xl/text-4xl (30px/36px) - 页面主要元素

## 9. 组件设计规范

### 9.1 按钮组件
#### 主要按钮 (Primary Button)
```css
/* 样式 */
bg-primary text-white font-semibold py-3 px-6 rounded-2xl 
active:scale-95 transition-transform

/* 使用场景 */
- 主要操作按钮（确认、提交、开始）
- CTA按钮
- 重要功能入口
```

#### 次要按钮 (Secondary Button)
```css
/* 样式 */
border-2 border-primary text-primary font-semibold py-3 px-6 rounded-2xl 
active:scale-95 transition-transform

/* 使用场景 */
- 次要操作（取消、返回）
- 可选功能
- 与主按钮配对使用
```

#### 图标按钮 (Icon Button)
```css
/* 样式 */
p-2 bg-gray-100 rounded-full hover:bg-gray-200 
active:scale-95 transition-all

/* 使用场景 */
- 工具栏按钮
- 快捷操作
- 空间受限的界面
```

### 9.2 卡片组件
#### 基础卡片
```css
/* 样式 */
bg-surface rounded-2xl p-4 shadow-sm border border-gray-100

/* 结构 */
- 统一的圆角和阴影
- 适当的内边距
- 清晰的边框分割
```

#### 强调卡片
```css
/* 样式 */
bg-gradient-to-br from-primary to-success rounded-3xl p-6 
shadow-lg text-white

/* 使用场景 */
- 重要数据展示
- 主要功能卡片
- 需要突出的内容
```

### 9.3 导航组件
#### 底部导航
```css
/* 样式 */
fixed bottom-0 left-0 right-0 bg-surface border-t border-gray-100 
px-6 py-3

/* 结构 */
- 5个主要功能入口
- 图标 + 文字标签
- 当前页面高亮显示
```

#### 顶部导航
```css
/* 样式 */
bg-surface px-6 pt-4 pb-6 border-b border-gray-100

/* 结构 */
- 页面标题居中
- 左侧返回/菜单按钮
- 右侧功能按钮
```

### 9.4 表单组件
#### 输入框
```css
/* 样式 */
w-full p-4 bg-surface border-2 border-gray-200 rounded-xl 
focus:border-primary focus:outline-none

/* 状态 */
- 默认态：淡灰色边框
- 焦点态：主色边框
- 错误态：红色边框
```

#### 选择按钮 (Radio/Checkbox)
```css
/* 样式 */
w-full p-4 border-2 border-gray-200 rounded-xl 
hover:border-primary hover:bg-primary/5 
active:scale-98 transition-all

/* 选中状态 */
border-primary bg-primary/5
```

## 10. 状态栏与设备模拟

### 10.1 iOS 状态栏
```css
/* 样式 */
flex justify-between items-center px-4 py-2 text-xs text-gray-800

/* 内容 */
- 左侧：时间显示 (9:41)
- 右侧：信号、WiFi、电池图标
- 整体高度：约32px
```

### 10.2 iPhone 设备框架
```css
/* 外框样式 */
background: linear-gradient(135deg, #1e1e1e 0%, #3a3a3a 100%);
border-radius: 3rem;
padding: 0.5rem;

/* 屏幕样式 */
background: #000;
border-radius: 2.5rem;
padding: 0.25rem;

/* 刘海设计 */
position: absolute;
top: 0.25rem;
left: 50%;
transform: translateX(-50%);
width: 120px;
height: 28px;
background: #1e1e1e;
border-radius: 0 0 1rem 1rem;
```

## 11. 动画与交互

### 11.1 基础动画
```css
/* 按钮点击 */
active:scale-95 transition-transform

/* 悬浮效果 */
hover:bg-gray-100 transition-colors

/* 焦点状态 */
focus:ring-2 focus:ring-primary focus:ring-offset-2
```

### 11.2 特殊动画
#### 呼吸动画 (紧急求助页面)
```css
@keyframes breathe {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.1); }
}
.breathing-circle {
    animation: breathe 4s ease-in-out infinite;
}
```

#### 脉冲动画 (重要按钮)
```css
@keyframes pulse-green {
    0%, 100% { box-shadow: 0 0 0 0 rgba(52, 211, 153, 0.7); }
    70% { box-shadow: 0 0 0 10px rgba(52, 211, 153, 0); }
}
.pulse-animation {
    animation: pulse-green 2s infinite;
}
```

#### 浮动动画 (装饰元素)
```css
@keyframes floating {
    0%, 100% { transform: translateY(0px); }
    50% { transform: translateY(-10px); }
}
.floating-animation {
    animation: floating 3s ease-in-out infinite;
}
```

## 12. 响应式设计

### 12.1 断点系统
- **移动端**: < 768px (主要目标)
- **平板端**: 768px - 1024px
- **桌面端**: > 1024px

### 12.2 适配策略
- **移动端优先**: 主要设计目标，完整功能
- **平板适配**: 增加边距，调整卡片布局
- **桌面适配**: 多列布局，增大内容宽度限制

## 13. 无障碍设计

### 13.1 颜色对比度
- 所有文字颜色对比度 ≥ 4.5:1
- 重要信息对比度 ≥ 7:1
- 不仅依赖颜色传达信息

### 13.2 触摸目标
- 最小触摸目标：44px × 44px
- 重要按钮推荐：48px × 48px
- 按钮间距至少8px

### 13.3 语义化标签
- 使用适当的HTML语义标签
- 为图片提供alt属性
- 表单元素关联label

## 14. 图片资源规范

### 14.1 图片来源
- **主要来源**: Unsplash (https://unsplash.com)
- **备选来源**: Pexels (https://pexels.com)
- **官方资源**: Apple Human Interface Guidelines

### 14.2 图片规格
- **卡片配图**: 400×200px, 16:9比例
- **头像图片**: 80×80px, 1:1比例
- **横幅图片**: 300×160px, 16:9比例
- **图标配图**: 80×60px, 4:3比例

### 14.3 图片优化
- 格式：WebP优先，fallback为JPEG
- 质量：80%压缩质量
- 加载：lazy loading延迟加载
- 尺寸：根据显示尺寸提供合适规格

## 15. 内容策略

### 15.1 文案基调
- **温暖友好**: 使用鼓励性、支持性的语言
- **专业可信**: 避免医疗诊断语言，提供科学依据
- **简洁明了**: 避免复杂术语，使用日常语言
- **积极正面**: 强调成长和改变，而非问题和困难

### 15.2 关键词汇
- 推荐使用：坚持、成长、改变、支持、陪伴、进步
- 谨慎使用：戒除、依赖、问题、治疗、病态
- 避免使用：羞耻性词汇、医疗诊断用词

## 16. 品牌视觉

### 16.1 App图标设计
- **主要元素**: 🌱 (幼苗) 象征成长和新生
- **颜色**: 主色调渐变 (primary to success)
- **形状**: 圆角正方形，符合iOS规范
- **背景**: 白色或浅色背景，确保可读性

### 16.2 Logo应用
- **文字Logo**: "戒色助手" 使用系统字体中粗体
- **图标Logo**: 🌱 + 圆形背景
- **组合Logo**: 图标 + 文字水平排列
- **颜色版本**: 彩色版、单色版、反色版

## 17. 开发实现说明

### 17.1 技术要求
- **CSS框架**: Tailwind CSS v3.0+
- **图标库**: FontAwesome v6.4.0
- **浏览器兼容**: iOS Safari 12+, Chrome 80+, Firefox 75+
- **响应式**: 移动端优先，渐进增强

### 17.2 代码规范
- 使用语义化HTML标签
- Tailwind class顺序：布局 > 尺寸 > 颜色 > 效果
- 自定义CSS最小化，优先使用Tailwind utilities
- 保持代码可读性和可维护性

### 17.3 性能优化
- CSS bundle size < 50KB (gzipped)
- 图片懒加载和WebP格式
- 减少不必要的JavaScript
- 使用CDN加载外部资源

---

**文档版本**: v1.0  
**最后更新**: 2024-01-15  
**维护者**: UI/UX 设计团队  

此设计规范文档是"戒色助手"产品UI/UX设计的权威指南，所有设计决策应遵循此规范，确保产品体验的一致性和专业性。 