# 戒色助手 UI 组件库

基于设计规范 `design/specs/Design_Spec.md` 构建的可复用 UI 组件库。

## 安装和使用

### 自动注册（推荐）

在 `main.ts` 中已经自动注册了所有组件：

```typescript
import UIComponents from './components/ui/index.js'
app.use(UIComponents)
```

### 单独导入

```vue
<script>
import { NfButton, NfCard } from '@/components/ui/index.js'

export default {
  components: {
    NfButton,
    NfCard
  }
}
</script>
```

## 组件列表

### 1. NfButton - 按钮组件

支持多种类型和尺寸的按钮组件。

#### Props

| 属性名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| type | String | 'primary' | 按钮类型：'primary', 'secondary', 'icon' |
| size | String | 'medium' | 按钮尺寸：'small', 'medium', 'large' |
| disabled | Boolean | false | 是否禁用 |
| iconLeft | String | '' | 左侧图标（FontAwesome类名） |
| iconRight | String | '' | 右侧图标（FontAwesome类名） |
| fullWidth | Boolean | false | 是否全宽 |
| label | String | '' | 按钮文字 |

#### Events

| 事件名 | 参数 | 说明 |
|--------|------|------|
| click | event | 点击事件 |

#### 示例

```vue
<template>
  <!-- Primary 按钮 -->
  <NfButton 
    type="primary" 
    label="确认"
    @click="handleClick"
  />
  
  <!-- Secondary 按钮 -->
  <NfButton 
    type="secondary"
    label="取消"
    icon-left="fa-arrow-left"
  />
  
  <!-- 图标按钮 -->
  <NfButton 
    type="icon"
    icon-left="fa-heart"
    @click="toggleLike"
  />
  
  <!-- 全宽大按钮 -->
  <NfButton 
    type="primary"
    size="large"
    label="开始戒色"
    full-width
    icon-left="fa-play"
  />
</template>
```

### 2. NfCard - 卡片组件

灵活的卡片容器组件，支持多种样式。

#### Props

| 属性名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| type | String | 'basic' | 卡片类型：'basic', 'highlight', 'gradient' |
| title | String | '' | 卡片标题 |
| content | String | '' | 卡片内容 |
| icon | String | '' | 标题图标（FontAwesome类名） |
| padding | String | 'medium' | 内边距：'small', 'medium', 'large' |
| shadow | Boolean | true | 是否显示阴影 |
| border | Boolean | true | 是否显示边框 |

#### Slots

| 插槽名 | 说明 |
|--------|------|
| default | 卡片主要内容 |
| header | 卡片头部自定义内容 |
| footer | 卡片底部内容 |

#### 示例

```vue
<template>
  <!-- 基础卡片 -->
  <NfCard 
    title="今日进度"
    icon="fa-calendar-day"
    content="已坚持 15 天"
  />
  
  <!-- 渐变卡片 -->
  <NfCard 
    type="gradient"
    title="等级信息"
    icon="fa-star"
  >
    <view class="text-center">
      <text class="text-2xl font-bold">Level 3</text>
      <text class="text-sm">新手导师</text>
    </view>
    
    <template #footer>
      <NfButton type="secondary" label="查看详情" />
    </template>
  </NfCard>
  
  <!-- 自定义头部 -->
  <NfCard>
    <template #header>
      <view class="flex justify-between items-center">
        <text class="font-semibold">自定义标题</text>
        <NfButton type="icon" icon-left="fa-ellipsis" />
      </view>
    </template>
    
    这里是卡片内容...
  </NfCard>
</template>
```

### 3. NfInput - 输入框组件

功能完整的表单输入组件。

#### Props

| 属性名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| modelValue | String/Number | '' | v-model 绑定值 |
| type | String | 'text' | 输入类型：'text', 'password', 'number', 'email', 'tel' |
| label | String | '' | 输入框标签 |
| placeholder | String | '' | 占位符文字 |
| required | Boolean | false | 是否必填（显示红色*） |
| disabled | Boolean | false | 是否禁用 |
| error | Boolean | false | 是否显示错误状态 |
| errorMessage | String | '' | 错误提示信息 |
| helpText | String | '' | 帮助文字 |
| iconLeft | String | '' | 左侧图标 |
| iconRight | String | '' | 右侧图标 |
| size | String | 'medium' | 尺寸：'small', 'medium', 'large' |
| maxlength | Number | -1 | 最大输入长度 |

#### Events

| 事件名 | 参数 | 说明 |
|--------|------|------|
| update:modelValue | value | v-model 更新事件 |
| input | event | 输入事件 |
| focus | event | 获取焦点事件 |
| blur | event | 失去焦点事件 |

#### 示例

```vue
<template>
  <view class="space-y-4">
    <!-- 基础输入框 -->
    <NfInput
      v-model="username"
      label="用户名"
      placeholder="请输入用户名"
      icon-left="fa-user"
      required
    />
    
    <!-- 密码输入框 -->
    <NfInput
      v-model="password"
      type="password"
      label="密码"
      placeholder="请输入密码"
      icon-left="fa-lock"
      required
    />
    
    <!-- 错误状态 -->
    <NfInput
      v-model="email"
      type="email"
      label="邮箱"
      placeholder="请输入邮箱"
      icon-left="fa-envelope"
      :error="emailError"
      error-message="请输入有效的邮箱地址"
      help-text="用于找回密码"
    />
  </view>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
      email: '',
      emailError: false
    }
  }
}
</script>
```

### 4. NfTabBar - 底部导航

底部标签页导航组件。

#### Props

| 属性名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| activeIndex | Number | 0 | 当前激活的标签索引 |
| tabs | Array | 默认5个标签 | 标签配置数组 |

#### tabs 配置项

```typescript
interface Tab {
  icon: string      // FontAwesome图标类名
  text: string      // 标签文字
  pagePath: string  // 页面路径（用于 uni.switchTab）
}
```

#### Events

| 事件名 | 参数 | 说明 |
|--------|------|------|
| change | index | 标签切换事件 |

#### 示例

```vue
<template>
  <NfTabBar 
    :active-index="currentTab"
    :tabs="customTabs"
    @change="handleTabChange"
  />
</template>

<script>
export default {
  data() {
    return {
      currentTab: 0,
      customTabs: [
        { icon: 'fa-home', text: '首页', pagePath: '/pages/index/index' },
        { icon: 'fa-chart-line', text: '进度', pagePath: '/pages/progress/progress' },
        { icon: 'fa-users', text: '社区', pagePath: '/pages/community/community' },
        { icon: 'fa-book', text: '学习', pagePath: '/pages/learning/learning' },
        { icon: 'fa-user', text: '我的', pagePath: '/pages/profile/profile' }
      ]
    }
  },
  methods: {
    handleTabChange(index) {
      this.currentTab = index
    }
  }
}
</script>
```

### 5. NfNavBar - 顶部导航

顶部导航栏组件，支持返回按钮和功能按钮。

#### Props

| 属性名 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| title | String | '' | 导航标题 |
| showBack | Boolean | false | 是否显示返回按钮 |
| rightIcon | String | '' | 右侧功能按钮图标 |
| backgroundColor | String | '#FFFFFF' | 背景颜色 |
| textColor | String | '#1F2937' | 文字颜色 |
| statusBarHeight | Number | 0 | 状态栏高度（自动获取） |

#### Events

| 事件名 | 参数 | 说明 |
|--------|------|------|
| back | - | 返回按钮点击事件 |
| rightClick | - | 右侧按钮点击事件 |

#### Slots

| 插槽名 | 说明 |
|--------|------|
| left | 左侧内容自定义 |
| center | 中间标题自定义 |
| right | 右侧内容自定义 |

#### 示例

```vue
<template>
  <!-- 基础导航 -->
  <NfNavBar 
    title="个人资料"
    :show-back="true"
    right-icon="fa-edit"
    @back="goBack"
    @right-click="editProfile"
  />
  
  <!-- 自定义导航 -->
  <NfNavBar>
    <template #left>
      <NfButton type="icon" icon-left="fa-bars" />
    </template>
    
    <template #center>
      <image src="/logo.png" class="w-8 h-8" />
    </template>
    
    <template #right>
      <view class="flex space-x-2">
        <NfButton type="icon" icon-left="fa-search" />
        <NfButton type="icon" icon-left="fa-bell" />
      </view>
    </template>
  </NfNavBar>
</template>
```

## 全局样式

组件库包含完整的全局样式文件 `styles/globals.css`，提供：

### 设计系统变量

```css
:root {
  --primary: #34D399;      /* 主色调 */
  --secondary: #06B6D4;    /* 次要色 */
  --success: #10B981;      /* 成功色 */
  --danger: #EF4444;       /* 警告色 */
  --background: #F8FAFC;   /* 页面背景 */
  --surface: #FFFFFF;      /* 卡片背景 */
  --text-primary: #1F2937; /* 主要文字 */
  --text-secondary: #6B7280; /* 次要文字 */
}
```

### 工具类

- **布局**: `.flex`, `.items-center`, `.justify-between`, `.w-full` 等
- **间距**: `.p-4`, `.m-2`, `.px-6`, `.py-3`, `.space-y-4` 等
- **文字**: `.text-lg`, `.font-semibold`, `.text-center` 等
- **颜色**: `.bg-primary`, `.text-success`, `.border-gray-200` 等
- **圆角**: `.rounded-lg`, `.rounded-xl`, `.rounded-full` 等
- **阴影**: `.shadow-sm`, `.shadow-md`, `.shadow-lg` 等

### 特殊动画

- `.animate-breathe`: 呼吸动画（紧急求助页面）
- `.animate-pulse-green`: 绿色脉冲动画（重要按钮）
- `.animate-floating`: 浮动动画（装饰元素）
- `.animate-fade-in`: 淡入动画

### 响应式支持

- 移动端优先设计
- 平板和桌面端适配
- 安全区域适配（刘海屏、底部指示器）

## 使用指南

### 1. 颜色使用

按照设计规范使用颜色：
- **主色（Primary）**: 主要操作按钮、重要强调
- **次色（Secondary）**: 次要功能、信息展示  
- **成功色**: 完成状态、正向反馈
- **警告色**: 仅用于紧急情况、错误提示

### 2. 间距系统

使用统一的间距系统：
- `space-y-*`: 垂直间距
- `space-x-*`: 水平间距
- `p-*`, `m-*`: 内外边距
- 数值对应：1=4px, 2=8px, 3=12px, 4=16px, 6=24px, 8=32px

### 3. 字体层级

- `text-xs` (12px): 标签、时间戳
- `text-sm` (14px): 辅助信息
- `text-base` (16px): 正文内容
- `text-lg` (18px): 小标题
- `text-xl` (20px): 副标题
- `text-2xl` (24px): 主标题
- `text-3xl` (30px): 页面标题

### 4. 无障碍设计

- 所有交互元素最小点击区域 44x44px
- 合适的颜色对比度
- 语义化的组件结构
- 键盘导航支持

## 开发规范

### 1. 组件命名

- 统一使用 `Nf` 前缀（NoFap 缩写）
- 使用 PascalCase 命名法
- 文件名与组件名保持一致

### 2. 代码结构

```
components/
└── ui/
    ├── button/
    │   └── NfButton.vue
    ├── card/
    │   └── NfCard.vue
    ├── form/
    │   └── NfInput.vue
    ├── navigation/
    │   ├── NfTabBar.vue
    │   └── NfNavBar.vue
    └── index.js
```

### 3. 样式规范

- 使用 CSS 变量而非硬编码颜色
- BEM 命名法：`.nf-button__icon--active`
- 响应式设计优先
- 性能优化（避免深层嵌套选择器）

## 测试页面

组件库包含完整的测试页面 `/pages/test/ui-test.vue`，展示所有组件的使用方法和效果。

可以通过以下方式访问测试页面：

```javascript
uni.navigateTo({
  url: '/pages/test/ui-test'
})
```

## 更新日志

### v1.0.0 (2025-01-23)

- ✅ 初始版本发布
- ✅ 基础组件：NfButton, NfCard, NfInput, NfTabBar, NfNavBar
- ✅ 完整的设计系统和工具类
- ✅ FontAwesome 图标集成
- ✅ 响应式设计支持
- ✅ 测试页面和文档 