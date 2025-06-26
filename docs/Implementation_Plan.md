# 戒色助手微信小程序实现方案

## 1. 项目概述

### 1.1 项目基本信息
- **项目名称**: 戒色助手微信小程序 (NoFap Helper Mini Program)
- **产品定位**: 专注于帮助年轻人戒除色情内容依赖的游戏化健康管理小程序
- **目标平台**: 微信小程序
- **开发周期**: 3-4个月
- **版本规划**: MVP v1.0 -> 功能完善版 v2.0 -> 商业化版本 v3.0

### 1.2 技术栈选型

#### 前端技术栈
- **框架**: uni-app X (Vue 3 + TypeScript)
- **UI组件库**: Element Plus (适配小程序)
- **CSS框架**: Tailwind CSS
- **状态管理**: Pinia
- **HTTP客户端**: uni.request + 封装
- **图表库**: uCharts (小程序专用)

#### 后端技术栈
- **语言**: Go 1.22+
- **Web框架**: Gin
- **数据库**: MySQL 8.0+ + Redis
- **ORM**: GORM
- **认证**: JWT
- **权限**: Casbin
- **消息队列**: RabbitMQ
- **文件存储**: 腾讯云COS
- **推送**: 微信小程序订阅消息

## 2. 系统架构设计

### 2.1 整体架构图

```
┌─────────────────────────────────────────────────┐
│                微信小程序端                        │
├─────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │
│  │   首页模块   │  │   社区模块   │  │   个人中心   │ │
│  └─────────────┘  └─────────────┘  └─────────────┘ │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │
│  │   评估模块   │  │   紧急求助   │  │   学习模块   │ │
│  └─────────────┘  └─────────────┘  └─────────────┘ │
├─────────────────────────────────────────────────┤
│                  API网关层                        │
├─────────────────────────────────────────────────┤
│                  后端服务层                        │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │
│  │  用户服务    │  │  评估服务    │  │  社区服务    │ │
│  └─────────────┘  └─────────────┘  └─────────────┘ │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │
│  │  游戏化服务  │  │  内容服务    │  │  通知服务    │ │
│  └─────────────┘  └─────────────┘  └─────────────┘ │
├─────────────────────────────────────────────────┤
│                  数据存储层                        │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │
│  │   MySQL     │  │    Redis     │  │  腾讯云COS   │ │
│  └─────────────┘  └─────────────┘  └─────────────┘ │
└─────────────────────────────────────────────────┘
```

### 2.2 核心模块设计

#### 2.2.1 用户管理模块
- 微信授权登录
- 用户信息管理
- 隐私设置
- 数据同步

#### 2.2.2 色隐指数评估系统
- 评估问卷系统
- 评分算法引擎
- 风险等级判定
- 评估历史记录
- 复评提醒机制

#### 2.2.3 每日打卡功能模块
- 每日签到打卡
- 情绪状态记录
- 连续天数统计
- 打卡奖励机制
- 打卡历史查看

#### 2.2.4 社区互动功能
- 匿名发帖系统
- 内容分类管理
- 点赞评论功能
- 内容审核机制
- 举报处理系统
- 用户互动统计

#### 2.2.5 游戏化激励机制
- 等级系统(50级)
- 经验值计算
- 成就徽章系统(100+)
- 连续奖励机制
- 里程碑奖励
- 排行榜系统

#### 2.2.6 寻求社区紧急互助系统
- 一键紧急求助
- 在线志愿者响应
- 匿名求助发布
- 实时互助聊天
- 专业资源推荐
- 危机干预记录

#### 2.2.7 学习模块
- **推荐系统**: 基于用户画像的个性化推荐
- **文章模块**: 科普文章、康复指导、心理健康
- **视频模块**: 专家讲座、康复经验分享
- **音频模块**: 冥想指导、正念练习、舒缓音乐
- **学习进度**: 学习记录、完成度统计
- **内容收藏**: 收藏管理、离线下载

## 3. 数据库设计

### 3.1 核心数据表结构

#### 用户表 (users)
```sql
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    openid VARCHAR(100) UNIQUE NOT NULL COMMENT '微信openid',
    unionid VARCHAR(100) COMMENT '微信unionid',
    nickname VARCHAR(50) COMMENT '用户昵称',
    avatar_url TEXT COMMENT '头像URL',
    gender TINYINT DEFAULT 0 COMMENT '性别:0未知,1男,2女',
    city VARCHAR(50) COMMENT '城市',
    province VARCHAR(50) COMMENT '省份',
    country VARCHAR(50) COMMENT '国家',
    privacy_level TINYINT DEFAULT 1 COMMENT '隐私级别:1低,2中,3高',
    status TINYINT DEFAULT 1 COMMENT '状态:0禁用,1正常',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
```

#### 戒色记录表 (abstinence_records)
```sql
CREATE TABLE abstinence_records (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    start_date DATE NOT NULL COMMENT '开始日期',
    current_streak INT DEFAULT 0 COMMENT '当前连续天数',
    longest_streak INT DEFAULT 0 COMMENT '最长连续天数',
    total_attempts INT DEFAULT 1 COMMENT '总尝试次数',
    level INT DEFAULT 1 COMMENT '当前等级',
    experience_points INT DEFAULT 0 COMMENT '经验值',
    total_checkins INT DEFAULT 0 COMMENT '总打卡次数',
    status TINYINT DEFAULT 1 COMMENT '状态:0停止,1进行中',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='戒色记录表';
```

#### 每日打卡表 (daily_checkins)
```sql
CREATE TABLE daily_checkins (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    checkin_date DATE NOT NULL COMMENT '打卡日期',
    mood TINYINT DEFAULT 3 COMMENT '心情评分:1-5',
    note TEXT COMMENT '打卡备注',
    experience_gained INT DEFAULT 10 COMMENT '获得经验值',
    is_continuous BOOLEAN DEFAULT FALSE COMMENT '是否连续打卡',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_date (user_id, checkin_date),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='每日打卡表';
```

#### 评估结果表 (assessment_results)
```sql
CREATE TABLE assessment_results (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    addiction_score INT NOT NULL COMMENT '色隐指数总分',
    detailed_scores JSON COMMENT '详细评分数据',
    risk_level TINYINT NOT NULL COMMENT '风险等级:1正常,2轻度,3中度,4重度,5严重',
    assessment_date DATE NOT NULL COMMENT '评估日期',
    version VARCHAR(10) DEFAULT 'v1.0' COMMENT '评估版本',
    recommendations TEXT COMMENT '建议内容',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评估结果表';
```

#### 社区动态表 (community_posts)
```sql
CREATE TABLE community_posts (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    title VARCHAR(200) COMMENT '标题',
    content TEXT NOT NULL COMMENT '内容',
    category TINYINT NOT NULL COMMENT '分类:1经验分享,2求助求鼓励,3日常打卡,4成功故事',
    images JSON COMMENT '图片URLs',
    like_count INT DEFAULT 0 COMMENT '点赞数',
    comment_count INT DEFAULT 0 COMMENT '评论数',
    view_count INT DEFAULT 0 COMMENT '浏览数',
    is_anonymous BOOLEAN DEFAULT TRUE COMMENT '是否匿名',
    is_top BOOLEAN DEFAULT FALSE COMMENT '是否置顶',
    status TINYINT DEFAULT 0 COMMENT '状态:0待审核,1已发布,2已删除',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区动态表';
```

#### 社区评论表 (community_comments)
```sql
CREATE TABLE community_comments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    post_id BIGINT NOT NULL COMMENT '动态ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    parent_id BIGINT DEFAULT 0 COMMENT '父评论ID',
    content TEXT NOT NULL COMMENT '评论内容',
    like_count INT DEFAULT 0 COMMENT '点赞数',
    is_anonymous BOOLEAN DEFAULT TRUE COMMENT '是否匿名',
    status TINYINT DEFAULT 1 COMMENT '状态:0删除,1正常',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES community_posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区评论表';
```

#### 紧急求助表 (emergency_helps)
```sql
CREATE TABLE emergency_helps (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT '求助用户ID',
    title VARCHAR(200) COMMENT '求助标题',
    content TEXT NOT NULL COMMENT '求助内容',
    urgency_level TINYINT DEFAULT 2 COMMENT '紧急程度:1低,2中,3高',
    status TINYINT DEFAULT 1 COMMENT '状态:1待响应,2已响应,3已解决',
    helper_count INT DEFAULT 0 COMMENT '响应人数',
    is_resolved BOOLEAN DEFAULT FALSE COMMENT '是否已解决',
    resolved_at TIMESTAMP NULL COMMENT '解决时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='紧急求助表';
```

#### 求助响应表 (help_responses)
```sql
CREATE TABLE help_responses (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    help_id BIGINT NOT NULL COMMENT '求助ID',
    helper_id BIGINT NOT NULL COMMENT '响应者ID',
    content TEXT NOT NULL COMMENT '响应内容',
    response_type TINYINT DEFAULT 1 COMMENT '响应类型:1文字,2语音,3视频',
    is_helpful BOOLEAN DEFAULT FALSE COMMENT '是否有帮助',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (help_id) REFERENCES emergency_helps(id) ON DELETE CASCADE,
    FOREIGN KEY (helper_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='求助响应表';
```

#### 学习内容表 (learning_contents)
```sql
CREATE TABLE learning_contents (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(200) NOT NULL COMMENT '标题',
    content TEXT COMMENT '内容',
    content_type TINYINT NOT NULL COMMENT '内容类型:1文章,2视频,3音频',
    category TINYINT NOT NULL COMMENT '分类:1科普,2康复指导,3心理健康,4冥想',
    author VARCHAR(100) COMMENT '作者',
    cover_url VARCHAR(500) COMMENT '封面图片',
    file_url VARCHAR(500) COMMENT '文件URL',
    duration INT DEFAULT 0 COMMENT '时长(秒)',
    difficulty TINYINT DEFAULT 1 COMMENT '难度:1初级,2中级,3高级',
    view_count INT DEFAULT 0 COMMENT '浏览数',
    like_count INT DEFAULT 0 COMMENT '点赞数',
    collect_count INT DEFAULT 0 COMMENT '收藏数',
    tags VARCHAR(500) COMMENT '标签',
    is_free BOOLEAN DEFAULT TRUE COMMENT '是否免费',
    sort_order INT DEFAULT 0 COMMENT '排序',
    status TINYINT DEFAULT 1 COMMENT '状态:0下线,1上线',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学习内容表';
```

#### 用户学习记录表 (user_learning_records)
```sql
CREATE TABLE user_learning_records (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    content_id BIGINT NOT NULL COMMENT '内容ID',
    progress INT DEFAULT 0 COMMENT '学习进度(百分比)',
    duration INT DEFAULT 0 COMMENT '学习时长(秒)',
    is_completed BOOLEAN DEFAULT FALSE COMMENT '是否完成',
    is_liked BOOLEAN DEFAULT FALSE COMMENT '是否点赞',
    is_collected BOOLEAN DEFAULT FALSE COMMENT '是否收藏',
    last_position INT DEFAULT 0 COMMENT '上次播放位置(秒)',
    completed_at TIMESTAMP NULL COMMENT '完成时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_content (user_id, content_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (content_id) REFERENCES learning_contents(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户学习记录表';
```

#### 成就系统表 (achievements)
```sql
CREATE TABLE achievements (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '成就名称',
    description TEXT COMMENT '成就描述',
    icon_url VARCHAR(500) COMMENT '图标URL',
    category TINYINT NOT NULL COMMENT '分类:1时间类,2行为类,3社区类',
    condition_type TINYINT NOT NULL COMMENT '条件类型:1连续天数,2总天数,3发帖数,4帮助数',
    condition_value INT NOT NULL COMMENT '条件数值',
    reward_exp INT DEFAULT 0 COMMENT '奖励经验值',
    reward_title VARCHAR(100) COMMENT '奖励称号',
    rarity TINYINT DEFAULT 1 COMMENT '稀有度:1普通,2稀有,3史诗,4传说',
    sort_order INT DEFAULT 0 COMMENT '排序',
    status TINYINT DEFAULT 1 COMMENT '状态:0禁用,1启用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成就系统表';
```

#### 用户成就表 (user_achievements)
```sql
CREATE TABLE user_achievements (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    achievement_id BIGINT NOT NULL COMMENT '成就ID',
    progress INT DEFAULT 0 COMMENT '完成进度',
    is_completed BOOLEAN DEFAULT FALSE COMMENT '是否完成',
    completed_at TIMESTAMP NULL COMMENT '完成时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_achievement (user_id, achievement_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (achievement_id) REFERENCES achievements(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户成就表';
```

### 3.2 索引优化策略
```sql
-- 用户相关索引
CREATE INDEX idx_users_openid ON users(openid);
CREATE INDEX idx_users_status ON users(status);

-- 戒色记录索引
CREATE INDEX idx_abstinence_user_id ON abstinence_records(user_id);
CREATE INDEX idx_abstinence_start_date ON abstinence_records(start_date);
CREATE INDEX idx_abstinence_level ON abstinence_records(level);

-- 每日打卡索引
CREATE INDEX idx_checkins_user_id ON daily_checkins(user_id);
CREATE INDEX idx_checkins_date ON daily_checkins(checkin_date);
CREATE INDEX idx_checkins_user_date ON daily_checkins(user_id, checkin_date);

-- 评估结果索引
CREATE INDEX idx_assessment_user_id ON assessment_results(user_id);
CREATE INDEX idx_assessment_date ON assessment_results(assessment_date);
CREATE INDEX idx_assessment_risk_level ON assessment_results(risk_level);

-- 社区动态索引
CREATE INDEX idx_posts_category ON community_posts(category);
CREATE INDEX idx_posts_status ON community_posts(status);
CREATE INDEX idx_posts_created_at ON community_posts(created_at);
CREATE INDEX idx_posts_user_id ON community_posts(user_id);
CREATE INDEX idx_posts_category_status ON community_posts(category, status);

-- 社区评论索引
CREATE INDEX idx_comments_post_id ON community_comments(post_id);
CREATE INDEX idx_comments_user_id ON community_comments(user_id);
CREATE INDEX idx_comments_parent_id ON community_comments(parent_id);

-- 紧急求助索引
CREATE INDEX idx_emergency_user_id ON emergency_helps(user_id);
CREATE INDEX idx_emergency_status ON emergency_helps(status);
CREATE INDEX idx_emergency_urgency ON emergency_helps(urgency_level);
CREATE INDEX idx_emergency_created_at ON emergency_helps(created_at);

-- 求助响应索引
CREATE INDEX idx_responses_help_id ON help_responses(help_id);
CREATE INDEX idx_responses_helper_id ON help_responses(helper_id);

-- 学习内容索引
CREATE INDEX idx_learning_category ON learning_contents(category);
CREATE INDEX idx_learning_type ON learning_contents(content_type);
CREATE INDEX idx_learning_status ON learning_contents(status);
CREATE INDEX idx_learning_sort ON learning_contents(sort_order);

-- 学习记录索引
CREATE INDEX idx_user_learning_user_id ON user_learning_records(user_id);
CREATE INDEX idx_user_learning_content_id ON user_learning_records(content_id);
CREATE INDEX idx_user_learning_completed ON user_learning_records(is_completed);

-- 成就相关索引
CREATE INDEX idx_achievements_category ON achievements(category);
CREATE INDEX idx_achievements_status ON achievements(status);
CREATE INDEX idx_user_achievements_user_id ON user_achievements(user_id);
CREATE INDEX idx_user_achievements_completed ON user_achievements(is_completed);
```

## 4. API接口设计

### 4.1 接口规范

#### 基础响应格式
```json
{
    "code": 200,
    "message": "success",
    "data": {},
    "timestamp": 1640995200
}
```

#### 错误码定义
```go
const (
    SUCCESS           = 200
    ERROR            = 500
    INVALID_PARAMS   = 400
    UNAUTHORIZED     = 401
    FORBIDDEN        = 403
    NOT_FOUND        = 404
    
    // 业务错误码
    USER_NOT_EXIST   = 1001
    ASSESSMENT_FAILED = 2001
    POST_AUDIT_FAILED = 3001
)
```

### 4.2 核心接口列表

#### 4.2.1 用户管理接口
```go
// 微信登录
POST /api/v1/auth/wx-login
{
    "code": "wx_code",
    "iv": "encrypt_iv",
    "encrypted_data": "encrypted_user_info"
}

// 获取用户信息
GET /api/v1/user/profile

// 更新用户信息
PUT /api/v1/user/profile
{
    "nickname": "用户昵称",
    "avatar_url": "头像URL",
    "privacy_level": 1
}
```

#### 4.2.2 评估系统接口
```go
// 获取评估问题
GET /api/v1/assessment/questions

// 提交评估答案
POST /api/v1/assessment/submit
{
    "answers": [1, 2, 3, 4, 5],
    "version": "v1.0"
}

// 获取评估历史
GET /api/v1/assessment/history?page=1&size=10
```

#### 4.2.3 每日打卡接口
```go
// 每日打卡
POST /api/v1/checkin/daily
{
    "mood": 3,
    "note": "今天状态不错"
}

// 获取打卡历史
GET /api/v1/checkin/history?page=1&size=30

// 获取打卡统计
GET /api/v1/checkin/statistics

// 检查今日是否已打卡
GET /api/v1/checkin/today
```

#### 4.2.4 戒色记录接口
```go
// 获取戒色统计
GET /api/v1/record/statistics

// 重新开始
POST /api/v1/record/restart
{
    "reason": "意外破戒"
}

// 获取等级信息
GET /api/v1/record/level

// 获取成就列表
GET /api/v1/record/achievements
```

#### 4.2.5 社区接口
```go
// 获取动态列表
GET /api/v1/community/posts?category=1&page=1&size=20

// 发布动态
POST /api/v1/community/posts
{
    "title": "分享标题",
    "content": "分享内容",
    "category": 1,
    "images": ["url1", "url2"],
    "is_anonymous": true
}

// 点赞/取消点赞
POST /api/v1/community/posts/{id}/like

// 评论
POST /api/v1/community/posts/{id}/comments
{
    "content": "评论内容",
    "parent_id": 0
}

// 获取评论列表
GET /api/v1/community/posts/{id}/comments?page=1&size=20
```

#### 4.2.6 紧急求助接口
```go
// 发布紧急求助
POST /api/v1/emergency/help
{
    "title": "需要帮助",
    "content": "遇到困难，需要支持",
    "urgency_level": 2
}

// 获取求助列表
GET /api/v1/emergency/helps?status=1&page=1&size=20

// 响应求助
POST /api/v1/emergency/helps/{id}/response
{
    "content": "我来帮助你",
    "response_type": 1
}

// 获取我的求助记录
GET /api/v1/emergency/my-helps?page=1&size=20

// 获取我的响应记录
GET /api/v1/emergency/my-responses?page=1&size=20
```

#### 4.2.7 学习模块接口
```go
// 获取学习内容列表
GET /api/v1/learning/contents?category=1&type=1&page=1&size=20

// 获取推荐内容
GET /api/v1/learning/recommend?size=10

// 获取内容详情
GET /api/v1/learning/contents/{id}

// 记录学习进度
POST /api/v1/learning/progress
{
    "content_id": 123,
    "progress": 50,
    "duration": 300,
    "last_position": 150
}

// 收藏/取消收藏
POST /api/v1/learning/contents/{id}/collect

// 点赞/取消点赞
POST /api/v1/learning/contents/{id}/like

// 获取我的学习记录
GET /api/v1/learning/my-records?page=1&size=20

// 获取我的收藏
GET /api/v1/learning/my-collections?page=1&size=20
```

## 5. 前端实现方案

### 5.1 项目结构
```
src/
├── components/          # 公共组件
│   ├── common/         # 通用组件
│   ├── charts/         # 图表组件
│   └── forms/          # 表单组件
├── pages/              # 页面
│   ├── index/          # 首页
│   ├── assessment/     # 评估
│   ├── checkin/        # 每日打卡
│   ├── community/      # 社区
│   ├── profile/        # 个人中心
│   ├── emergency/      # 紧急求助
│   └── learning/       # 学习内容
├── store/              # 状态管理
│   ├── modules/        # 模块
│   └── index.js        # 入口
├── utils/              # 工具函数
│   ├── request.js      # 请求封装
│   ├── auth.js         # 认证相关
│   └── common.js       # 通用工具
├── static/             # 静态资源
└── App.vue             # 应用入口
```

### 5.2 关键组件实现

#### 5.2.1 底部导航组件
```vue
<template>
  <view class="tabbar">
    <view 
      v-for="(item, index) in tabList" 
      :key="index"
      class="tabbar-item"
      :class="{ active: current === index }"
      @click="switchTab(index, item.pagePath)"
    >
      <view class="tabbar-icon">
        <text :class="item.iconFont"></text>
      </view>
      <text class="tabbar-text">{{ item.text }}</text>
    </view>
  </view>
</template>

<script setup>
import { ref } from 'vue'

const current = ref(0)
const tabList = [
  { text: '首页', pagePath: '/pages/index/index', iconFont: 'icon-home' },
  { text: '打卡', pagePath: '/pages/checkin/index', iconFont: 'icon-check' },
  { text: '社区', pagePath: '/pages/community/index', iconFont: 'icon-community' },
  { text: '学习', pagePath: '/pages/learning/index', iconFont: 'icon-book' },
  { text: '我的', pagePath: '/pages/profile/index', iconFont: 'icon-user' }
]

const switchTab = (index, pagePath) => {
  current.value = index
  uni.switchTab({ url: pagePath })
}
</script>
```

#### 5.2.2 等级进度组件
```vue
<template>
  <view class="level-progress">
    <view class="level-info">
      <text class="level-text">Lv.{{ level }}</text>
      <text class="level-title">{{ levelTitle }}</text>
    </view>
    <view class="progress-bar">
      <view 
        class="progress-fill"
        :style="{ width: progressPercent + '%' }"
      ></view>
    </view>
    <view class="exp-info">
      <text>{{ currentExp }}/{{ nextLevelExp }} EXP</text>
    </view>
  </view>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  level: Number,
  currentExp: Number,
  nextLevelExp: Number
})

const levelTitle = computed(() => {
  const titles = {
    1: '初心者', 5: '坚持者', 10: '挑战者',
    15: '勇士', 20: '专家', 25: '大师'
  }
  return titles[props.level] || '修行者'
})

const progressPercent = computed(() => {
  return (props.currentExp / props.nextLevelExp * 100).toFixed(1)
})
</script>
```

### 5.3 状态管理设计

#### 5.3.1 用户状态管理
```javascript
// store/modules/user.js
import { defineStore } from 'pinia'
import { login, getUserInfo } from '@/api/user'

export const useUserStore = defineStore('user', {
  state: () => ({
    userInfo: null,
    token: uni.getStorageSync('token') || '',
    isLogin: false
  }),

  getters: {
    userId: (state) => state.userInfo?.id,
    nickname: (state) => state.userInfo?.nickname || '匿名用户',
    avatarUrl: (state) => state.userInfo?.avatar_url || '/static/default-avatar.png'
  },

  actions: {
    async wxLogin(code) {
      try {
        const res = await login({ code })
        this.token = res.data.token
        this.userInfo = res.data.user
        this.isLogin = true
        
        uni.setStorageSync('token', this.token)
        return res
      } catch (error) {
        throw error
      }
    },

    async fetchUserInfo() {
      try {
        const res = await getUserInfo()
        this.userInfo = res.data
        return res
      } catch (error) {
        throw error
      }
    },

    logout() {
      this.token = ''
      this.userInfo = null
      this.isLogin = false
      uni.removeStorageSync('token')
    }
  }
})
```

#### 5.3.2 戒色记录状态管理
```javascript
// store/modules/record.js
import { defineStore } from 'pinia'
import { getRecordStats, checkin } from '@/api/record'

export const useRecordStore = defineStore('record', {
  state: () => ({
    currentStreak: 0,
    longestStreak: 0,
    totalDays: 0,
    level: 1,
    experience: 0,
    lastCheckinDate: null
  }),

  getters: {
    canCheckin: (state) => {
      const today = new Date().toDateString()
      const lastCheckin = new Date(state.lastCheckinDate).toDateString()
      return today !== lastCheckin
    },

    nextLevelExp: (state) => {
      return state.level * 100 // 简化的经验值计算
    }
  },

  actions: {
    async fetchStats() {
      try {
        const res = await getRecordStats()
        Object.assign(this, res.data)
        return res
      } catch (error) {
        throw error
      }
    },

    async doCheckin(data) {
      try {
        const res = await checkin(data)
        await this.fetchStats() // 重新获取统计数据
        return res
      } catch (error) {
        throw error
      }
    }
  }
})
```

## 6. 后端实现方案

### 6.1 项目结构
```
server/
├── api/                # API控制器
│   └── v1/
│       ├── auth.go     # 认证相关
│       ├── user.go     # 用户管理
│       ├── assessment.go # 评估系统
│       ├── checkin.go  # 每日打卡
│       ├── record.go   # 戒色记录
│       ├── community.go # 社区功能
│       ├── emergency.go # 紧急求助
│       └── learning.go # 学习模块
├── config/             # 配置
├── core/               # 核心功能
├── global/             # 全局变量
├── initialize/         # 初始化
├── middleware/         # 中间件
├── model/              # 数据模型
├── router/             # 路由
├── service/            # 业务逻辑
└── utils/              # 工具函数
```

### 6.2 核心服务实现

#### 6.2.1 微信认证服务
```go
// service/auth_service.go
package service

import (
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "your-project/global"
    "your-project/model"
)

type AuthService struct{}

func (s *AuthService) WxLogin(code string) (*model.User, string, error) {
    // 1. 通过code获取openid
    openid, sessionKey, err := s.getWxUserInfo(code)
    if err != nil {
        return nil, "", err
    }

    // 2. 查找或创建用户
    var user model.User
    err = global.GVA_DB.Where("openid = ?", openid).First(&user).Error
    if err != nil {
        // 创建新用户
        user = model.User{
            Openid: openid,
            Status: 1,
        }
        global.GVA_DB.Create(&user)
    }

    // 3. 生成JWT token
    token, err := s.generateToken(user.ID)
    if err != nil {
        return nil, "", err
    }

    return &user, token, nil
}

func (s *AuthService) getWxUserInfo(code string) (string, string, error) {
    url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
        global.GVA_CONFIG.Wechat.AppID,
        global.GVA_CONFIG.Wechat.AppSecret,
        code)

    // HTTP请求微信API
    // ... 实现HTTP请求逻辑
    
    return openid, sessionKey, nil
}
```

#### 6.2.2 评估系统服务
```go
// service/assessment_service.go
package service

import (
    "your-project/global"
    "your-project/model"
)

type AssessmentService struct{}

func (s *AssessmentService) SubmitAssessment(userID uint, answers []int) (*model.AssessmentResult, error) {
    // 1. 计算评估分数
    score := s.calculateScore(answers)
    
    // 2. 确定风险等级
    riskLevel := s.getRiskLevel(score)
    
    // 3. 保存评估结果
    result := model.AssessmentResult{
        UserID:        userID,
        AddictionScore: score,
        DetailedScores: answers,
        RiskLevel:     riskLevel,
        AssessmentDate: time.Now(),
    }
    
    err := global.GVA_DB.Create(&result).Error
    if err != nil {
        return nil, err
    }
    
    return &result, nil
}

func (s *AssessmentService) calculateScore(answers []int) int {
    total := 0
    weights := []int{2, 3, 2, 4, 3, 2, 4, 3, 2, 3} // 题目权重
    
    for i, answer := range answers {
        if i < len(weights) {
            total += answer * weights[i]
        }
    }
    
    return total
}

func (s *AssessmentService) getRiskLevel(score int) int {
    switch {
    case score <= 20:
        return 1 // 正常
    case score <= 40:
        return 2 // 轻度
    case score <= 60:
        return 3 // 中度
    case score <= 80:
        return 4 // 重度
    default:
        return 5 // 严重
    }
}
```

#### 6.2.3 游戏化系统服务
```go
// service/gamification_service.go
package service

import (
    "your-project/global"
    "your-project/model"
)

type GamificationService struct{}

func (s *GamificationService) ProcessCheckin(userID uint, mood int, note string) error {
    // 1. 更新戒色记录
    var record model.AbstinenceRecord
    err := global.GVA_DB.Where("user_id = ?", userID).First(&record).Error
    if err != nil {
        // 创建新记录
        record = model.AbstinenceRecord{
            UserID:     userID,
            StartDate:  time.Now(),
            CurrentStreak: 1,
            LongestStreak: 1,
            Level:      1,
            ExperiencePoints: 10,
        }
    } else {
        // 更新现有记录
        record.CurrentStreak++
        if record.CurrentStreak > record.LongestStreak {
            record.LongestStreak = record.CurrentStreak
        }
        
        // 计算经验值奖励
        expReward := s.calculateExpReward(record.CurrentStreak)
        record.ExperiencePoints += expReward
        
        // 检查是否升级
        newLevel := s.calculateLevel(record.ExperiencePoints)
        if newLevel > record.Level {
            record.Level = newLevel
            // 触发升级奖励
            s.triggerLevelUpReward(userID, newLevel)
        }
    }
    
    err = global.GVA_DB.Save(&record).Error
    if err != nil {
        return err
    }
    
    // 2. 记录打卡日志
    checkinLog := model.CheckinLog{
        UserID: userID,
        Mood:   mood,
        Note:   note,
        Date:   time.Now(),
    }
    global.GVA_DB.Create(&checkinLog)
    
    return nil
}

func (s *GamificationService) calculateExpReward(streak int) int {
    baseExp := 10
    bonusExp := 0
    
    // 连续天数奖励
    switch {
    case streak >= 30:
        bonusExp = 50
    case streak >= 7:
        bonusExp = 20
    case streak >= 3:
        bonusExp = 10
    }
    
    return baseExp + bonusExp
}

func (s *GamificationService) calculateLevel(exp int) int {
    // 简化的等级计算公式
    level := 1
    requiredExp := 100
    
    for exp >= requiredExp {
        exp -= requiredExp
        level++
        requiredExp = level * 100 // 每级所需经验递增
    }
    
    return level
}
```

## 7. 部署方案

### 7.1 服务器环境
- **云服务商**: 腾讯云/阿里云
- **服务器配置**: 2核4G内存起步
- **操作系统**: Ubuntu 20.04 LTS
- **容器化**: Docker + Docker Compose

### 7.2 部署架构
```yaml
# docker-compose.yml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
    depends_on:
      - mysql
      - redis
    volumes:
      - ./config:/app/config

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_DATABASE: nofap_helper
      MYSQL_USER: ${DB_USER}
      MYSQL_PASSWORD: ${DB_PASSWORD}
      MYSQL_ROOT_PASSWORD: ${DB_ROOT_PASSWORD}
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - "3306:3306"
    command: --default-authentication-plugin=mysql_native_password

  redis:
    image: redis:6-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - app

volumes:
  mysql_data:
  redis_data:
```

### 7.3 CI/CD流程
```yaml
# .github/workflows/deploy.yml
name: Deploy to Production

on:
  push:
    branches: [ main ]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    
    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.22
    
    - name: Build
      run: go build -v ./...
    
    - name: Test
      run: go test -v ./...
    
    - name: Deploy to server
      uses: appleboy/ssh-action@v0.1.4
      with:
        host: ${{ secrets.HOST }}
        username: ${{ secrets.USERNAME }}
        key: ${{ secrets.KEY }}
        script: |
          cd /app/nofap-helper
          git pull origin main
          docker-compose down
          docker-compose up -d --build
```

## 8. 开发计划与里程碑

### 8.1 开发阶段规划

#### Phase 1: 基础功能开发 (4-6周)
- **Week 1-2**: 项目搭建、用户认证、基础UI框架
- **Week 3-4**: 色隐指数评估系统、每日打卡功能
- **Week 5-6**: 游戏化激励机制、基础测试

#### Phase 2: 核心功能完善 (4-6周)
- **Week 7-8**: 社区互动功能开发
- **Week 9-10**: 寻求社区紧急互助系统
- **Week 11-12**: 学习模块(推荐、文章、视频、音频)

#### Phase 3: 优化与上线 (2-4周)
- **Week 13-14**: 性能优化、安全加固、内容审核
- **Week 15-16**: 全面测试、部署、小程序审核

### 8.2 关键里程碑
- **M1**: 用户管理模块 + 色隐指数评估系统完成
- **M2**: 每日打卡功能 + 游戏化激励机制完成  
- **M3**: 社区互动功能 + 寻求社区紧急互助系统完成
- **M4**: 学习模块完成，系统集成测试通过
- **M5**: 小程序提审通过，正式上线运营

## 9. 风险评估与应对策略

### 9.1 技术风险
- **风险**: 微信小程序API限制
- **应对**: 提前调研API限制，设计降级方案

- **风险**: 数据库性能瓶颈
- **应对**: 数据库优化、读写分离、缓存策略

### 9.2 业务风险
- **风险**: 内容审核不过关
- **应对**: 严格的内容审核机制、敏感词过滤

- **风险**: 用户隐私保护
- **应对**: 数据加密、匿名化处理、合规性审查

### 9.3 合规风险
- **风险**: 微信小程序审核不通过
- **应对**: 严格遵循微信小程序开发规范

- **风险**: 涉及敏感内容被下架
- **应对**: 内容健康化、专业化包装

## 10. 总结

本实现方案基于PRD文档和原型设计，采用现代化的技术栈，设计了完整的微信小程序解决方案。方案涵盖了从技术架构到部署上线的全流程，为戒色助手产品的成功开发和运营提供了详细的技术指导。

关键成功因素：
1. 严格的内容审核机制确保合规性
2. 优秀的用户体验设计提升留存率
3. 科学的游戏化机制增强用户粘性
4. 完善的技术架构保障系统稳定性
5. 渐进式开发策略降低项目风险

该方案将为用户提供一个专业、安全、有效的戒色辅助工具，帮助他们重建健康的生活方式。 