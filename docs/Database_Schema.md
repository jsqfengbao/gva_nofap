# 戒色助手微信小程序数据库设计文档

## 数据库概述

本项目使用MySQL 8.0+作为主数据库，采用GORM进行ORM映射。数据库设计遵循第三范式，包含12个核心业务表，支持用户管理、戒色记录、社区互动、学习内容、成就系统等完整功能。

## 表结构设计

### 1. 用户相关表

#### nofap_wx_users (微信用户表)
存储微信小程序用户的基本信息和隐私设置。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| openid | varchar(100) | UNIQUE, NOT NULL | 微信openid |
| unionid | varchar(100) | | 微信unionid |
| nickname | varchar(50) | | 用户昵称 |
| avatar_url | text | | 头像URL |
| gender | int | DEFAULT 0 | 性别: 0未知,1男,2女 |
| city | varchar(50) | | 城市 |
| province | varchar(50) | | 省份 |
| country | varchar(50) | | 国家 |
| privacy_level | int | DEFAULT 1 | 隐私级别: 1低,2中,3高 |
| status | int | DEFAULT 1 | 状态: 0禁用,1正常 |
| last_login_at | timestamp | | 最后登录时间 |
| created_at | timestamp | NOT NULL | 创建时间 |
| updated_at | timestamp | NOT NULL | 更新时间 |

**索引设计:**
- PRIMARY KEY (id)
- UNIQUE INDEX idx_wx_users_openid (openid)
- INDEX idx_wx_users_status (status)

### 2. 戒色记录相关表

#### nofap_abstinence_records (戒色记录表)
存储用户的戒色总体数据和进度信息。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| user_id | uint | NOT NULL, FK | 用户ID |
| start_date | timestamp | NOT NULL | 开始戒色日期 |
| current_streak | int | DEFAULT 0 | 当前连续天数 |
| longest_streak | int | DEFAULT 0 | 最长连续天数 |
| total_days | int | DEFAULT 0 | 总戒色天数 |
| success_rate | decimal(5,2) | DEFAULT 0 | 成功率百分比 |
| level | int | DEFAULT 1 | 等级(1-50) |
| experience | int | DEFAULT 0 | 经验值 |
| last_relapse_at | timestamp | | 最后复发时间 |
| status | int | DEFAULT 1 | 状态: 0结束,1进行中 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_abstinence_records_user_id (user_id)
- INDEX idx_abstinence_records_user_status (user_id, status)
- INDEX idx_abstinence_records_level (level)

#### daily_checkins (每日打卡表)
记录用户每日的打卡状态和心情数据。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| user_id | uint | NOT NULL, FK | 用户ID |
| checkin_date | date | NOT NULL | 打卡日期 |
| mood_level | int | NOT NULL | 情绪等级(1-5) |
| notes | text | | 打卡备注 |
| rewards | int | DEFAULT 0 | 获得奖励经验值 |
| is_success | boolean | DEFAULT true | 是否成功戒色 |
| streak | int | DEFAULT 0 | 当前连续天数 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_daily_checkins_user_id (user_id)
- INDEX idx_daily_checkins_user_date (user_id, checkin_date)
- INDEX idx_daily_checkins_date (checkin_date)

#### assessment_results (评估结果表)
存储用户的色隐指数评估测试结果。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| user_id | uint | NOT NULL, FK | 用户ID |
| total_score | int | NOT NULL | 总分 |
| risk_level | int | NOT NULL | 风险等级: 1正常,2轻度,3中度,4重度,5严重 |
| answers | text | | 答案JSON |
| test_date | timestamp | NOT NULL | 测试日期 |
| test_type | int | DEFAULT 1 | 测试类型: 1初次,2复评 |
| duration | int | | 测试耗时(秒) |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_assessment_results_user_id (user_id)
- INDEX idx_assessment_results_user_date (user_id, test_date)
- INDEX idx_assessment_results_risk_level (risk_level)

### 3. 社区相关表

#### community_posts (社区动态表)
存储用户发布的社区动态内容。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| user_id | uint | NOT NULL, FK | 用户ID |
| title | varchar(200) | NOT NULL | 标题 |
| content | text | NOT NULL | 内容 |
| category | int | NOT NULL | 分类: 1经验分享,2求助求鼓励,3日常打卡,4成功故事 |
| is_anonymous | boolean | DEFAULT false | 是否匿名 |
| view_count | int | DEFAULT 0 | 查看次数 |
| like_count | int | DEFAULT 0 | 点赞数 |
| comment_count | int | DEFAULT 0 | 评论数 |
| status | int | DEFAULT 1 | 状态: 0删除,1正常,2待审核,3已拒绝 |
| audit_at | timestamp | | 审核时间 |
| audit_by | uint | | 审核人ID |
| audit_reason | varchar(500) | | 审核原因 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_community_posts_user_id (user_id)
- INDEX idx_community_posts_category_status (category, status)
- INDEX idx_community_posts_created_at (created_at)
- INDEX idx_community_posts_like_count (like_count)

#### community_comments (社区评论表)
存储对社区动态的评论信息。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| post_id | uint | NOT NULL, FK | 动态ID |
| user_id | uint | NOT NULL, FK | 用户ID |
| content | text | NOT NULL | 评论内容 |
| parent_id | uint | FK | 父评论ID,0为顶级评论 |
| is_anonymous | boolean | DEFAULT false | 是否匿名 |
| like_count | int | DEFAULT 0 | 点赞数 |
| status | int | DEFAULT 1 | 状态: 0删除,1正常,2待审核 |
| audit_at | timestamp | | 审核时间 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_community_comments_post_id (post_id)
- INDEX idx_community_comments_user_id (user_id)
- INDEX idx_community_comments_post_status (post_id, status)
- INDEX idx_community_comments_parent_id (parent_id)

#### community_likes (社区点赞表)
记录用户对动态和评论的点赞行为。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| post_id | uint | NOT NULL, FK | 动态ID |
| user_id | uint | NOT NULL, FK | 用户ID |
| like_type | int | DEFAULT 1 | 点赞类型: 1动态,2评论 |
| target_id | uint | NOT NULL | 目标ID(动态ID或评论ID) |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_community_likes_post_id (post_id)
- INDEX idx_community_likes_user_id (user_id)
- UNIQUE INDEX idx_community_likes_unique (user_id, like_type, target_id)

### 4. 紧急求助相关表

#### emergency_helps (紧急求助表)
存储用户发布的紧急求助信息。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| user_id | uint | NOT NULL, FK | 求助用户ID |
| title | varchar(200) | NOT NULL | 求助标题 |
| content | text | NOT NULL | 求助内容 |
| urgency_level | int | NOT NULL | 紧急程度: 1一般,2紧急,3非常紧急 |
| is_anonymous | boolean | DEFAULT true | 是否匿名 |
| status | int | DEFAULT 1 | 状态: 1待响应,2进行中,3已解决,4已关闭 |
| response_count | int | DEFAULT 0 | 响应次数 |
| resolved_at | timestamp | | 解决时间 |
| closed_at | timestamp | | 关闭时间 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_emergency_helps_user_id (user_id)
- INDEX idx_emergency_helps_status_urgency (status, urgency_level)
- INDEX idx_emergency_helps_created_at (created_at)

#### help_responses (求助响应表)
存储对紧急求助的响应信息。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| help_id | uint | NOT NULL, FK | 求助ID |
| responder_id | uint | NOT NULL, FK | 响应者ID |
| content | text | NOT NULL | 响应内容 |
| response_type | int | DEFAULT 1 | 响应类型: 1文字回复,2申请私聊,3专业建议 |
| is_volunteer | boolean | DEFAULT false | 是否志愿者 |
| status | int | DEFAULT 1 | 状态: 1正常,2已删除 |
| helpful_count | int | DEFAULT 0 | 有用评价数 |
| response_at | timestamp | NOT NULL | 响应时间 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_help_responses_help_id (help_id)
- INDEX idx_help_responses_responder_id (responder_id)

### 5. 学习内容相关表

#### learning_contents (学习内容表)
存储平台的学习内容资源。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| title | varchar(200) | NOT NULL | 标题 |
| summary | varchar(500) | | 摘要 |
| content | longtext | NOT NULL | 内容 |
| content_type | int | NOT NULL | 内容类型: 1文章,2视频,3音频 |
| category | int | NOT NULL | 分类: 1科普知识,2康复指导,3心理健康,4经验分享 |
| difficulty | int | DEFAULT 1 | 难度等级: 1入门,2初级,3中级,4高级 |
| duration | int | | 时长(分钟) |
| thumbnail_url | text | | 缩略图URL |
| media_url | text | | 媒体文件URL |
| author | varchar(100) | | 作者 |
| view_count | int | DEFAULT 0 | 观看次数 |
| like_count | int | DEFAULT 0 | 点赞数 |
| collect_count | int | DEFAULT 0 | 收藏数 |
| comment_count | int | DEFAULT 0 | 评论数 |
| status | int | DEFAULT 1 | 状态: 0删除,1正常,2草稿 |
| publish_at | timestamp | | 发布时间 |
| tags | varchar(500) | | 标签,逗号分隔 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_learning_contents_category_status (category, status)
- INDEX idx_learning_contents_type (content_type)
- INDEX idx_learning_contents_view_count (view_count)

#### user_learning_records (用户学习记录表)
记录用户的学习行为和进度。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| user_id | uint | NOT NULL, FK | 用户ID |
| content_id | uint | NOT NULL, FK | 内容ID |
| start_time | timestamp | NOT NULL | 开始学习时间 |
| end_time | timestamp | | 结束学习时间 |
| duration | int | DEFAULT 0 | 学习时长(秒) |
| progress | int | DEFAULT 0 | 学习进度(百分比) |
| is_completed | boolean | DEFAULT false | 是否完成 |
| is_liked | boolean | DEFAULT false | 是否点赞 |
| is_collected | boolean | DEFAULT false | 是否收藏 |
| rating | int | DEFAULT 0 | 评分(1-5) |
| notes | text | | 学习笔记 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_user_learning_records_user_id (user_id)
- INDEX idx_user_learning_records_content_id (content_id)
- INDEX idx_user_learning_records_user_content (user_id, content_id)
- INDEX idx_user_learning_records_completed (is_completed)

### 6. 成就系统相关表

#### achievements (成就系统表)
定义平台的所有成就项目。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| name | varchar(100) | NOT NULL | 成就名称 |
| description | varchar(500) | | 成就描述 |
| icon_url | text | | 成就图标URL |
| category | int | NOT NULL | 成就分类: 1打卡,2等级,3社区,4学习,5特殊 |
| type | int | NOT NULL | 成就类型: 1累计,2连续,3一次性 |
| condition | text | | 解锁条件JSON |
| rewards | int | DEFAULT 0 | 奖励经验值 |
| rarity | int | DEFAULT 1 | 稀有度: 1普通,2稀有,3史诗,4传说 |
| display_order | int | DEFAULT 0 | 显示顺序 |
| is_active | boolean | DEFAULT true | 是否启用 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_achievements_category_active (category, is_active)
- INDEX idx_achievements_rarity (rarity)

#### user_achievements (用户成就表)
记录用户获得的成就信息。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | uint | PK, AUTO_INCREMENT | 主键ID |
| user_id | uint | NOT NULL, FK | 用户ID |
| achievement_id | uint | NOT NULL, FK | 成就ID |
| unlocked_at | timestamp | NOT NULL | 解锁时间 |
| progress | int | DEFAULT 0 | 完成进度 |
| is_notified | boolean | DEFAULT false | 是否已通知 |

**索引设计:**
- PRIMARY KEY (id)
- INDEX idx_user_achievements_user_id (user_id)
- INDEX idx_user_achievements_achievement_id (achievement_id)
- INDEX idx_user_achievements_user_unlocked (user_id, unlocked_at)
- INDEX idx_user_achievements_notified (is_notified)

## 数据关系图

```
wx_users (用户表)
├── abstinence_records (戒色记录) [1:1]
├── daily_checkins (每日打卡) [1:N]
├── assessment_results (评估结果) [1:N]
├── community_posts (社区动态) [1:N]
├── community_comments (社区评论) [1:N]
├── community_likes (社区点赞) [1:N]
├── emergency_helps (紧急求助) [1:N]
├── help_responses (求助响应) [1:N]
├── user_learning_records (学习记录) [1:N]
└── user_achievements (用户成就) [1:N]

community_posts (社区动态)
├── community_comments (评论) [1:N]
└── community_likes (点赞) [1:N]

emergency_helps (紧急求助)
└── help_responses (响应) [1:N]

learning_contents (学习内容)
└── user_learning_records (学习记录) [1:N]

achievements (成就系统)
└── user_achievements (用户成就) [1:N]
```

## 性能优化策略

### 1. 索引优化
- 为高频查询字段创建合适的索引
- 使用复合索引优化多条件查询
- 定期分析查询性能并调整索引

### 2. 分区策略
- 对于大数据量表(如daily_checkins、user_learning_records)考虑按时间分区
- 提高查询性能和维护效率

### 3. 缓存策略
- 用户基本信息、成就列表等相对稳定的数据使用Redis缓存
- 热点内容和统计数据进行缓存

### 4. 数据归档
- 定期归档历史数据，保持主表的查询性能
- 制定数据保留策略

## 数据安全考虑

### 1. 隐私保护
- 用户可设置隐私级别控制信息展示
- 支持匿名发帖和求助
- 敏感信息加密存储

### 2. 数据完整性
- 设置外键约束保证数据一致性
- 使用事务处理复杂业务操作

### 3. 备份策略
- 定期全量备份和增量备份
- 制定灾难恢复方案

这个数据库设计支持完整的戒色助手功能需求，具备良好的扩展性和性能表现。 