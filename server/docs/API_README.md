# 戒色助手小程序 API 文档

## 概述

本文档描述了戒色助手微信小程序的后端API接口。所有API都遵循RESTful设计原则，使用JSON格式进行数据交换。

## 基础信息

- **Base URL**: `http://localhost:8888`
- **API版本**: v1
- **数据格式**: JSON
- **字符编码**: UTF-8

## 认证方式

大部分API需要用户认证，认证方式为JWT Token。

### 获取Token

通过微信登录接口获取：

```http
POST /miniprogram/auth/wx-login
```

### 使用Token

在请求头中添加：

```http
Authorization: Bearer <your_token>
```

或者使用：

```http
x-token: <your_token>
```

## API分组

### 1. 认证模块 (Auth)

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 微信登录 | POST | `/miniprogram/auth/wx-login` | 微信小程序登录 |
| 刷新Token | POST | `/miniprogram/auth/refresh-token` | 刷新用户token |

### 2. 用户模块 (User)

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 获取用户资料 | GET | `/miniprogram/user/profile` | 获取用户详细资料 |
| 更新用户信息 | PUT | `/miniprogram/user/info` | 更新用户基本信息 |
| 更新隐私级别 | PUT | `/miniprogram/user/privacy-level` | 更新用户隐私级别 |
| 获取统计数据 | GET | `/miniprogram/user/stats` | 获取个人中心统计数据 |
| 更新通知设置 | PUT | `/miniprogram/user/notification-settings` | 更新通知设置 |
| 更新隐私设置 | PUT | `/miniprogram/user/privacy-settings` | 更新隐私设置 |
| 创建数据导出 | POST | `/miniprogram/user/data-export` | 创建数据导出任务 |
| 获取导出状态 | GET | `/miniprogram/user/data-export/:exportId` | 获取数据导出状态 |
| 下载导出文件 | GET | `/miniprogram/user/data-export/:exportId/download` | 下载导出文件 |

### 3. 打卡模块 (Checkin)

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 每日打卡 | POST | `/miniprogram/checkin/daily` | 用户每日打卡 |
| 获取打卡历史 | GET | `/miniprogram/checkin/history` | 获取打卡历史记录 |
| 获取本周进度 | GET | `/miniprogram/checkin/weekly-progress` | 获取本周打卡进度 |
| 获取月度日历 | GET | `/miniprogram/checkin/monthly-calendar` | 获取月度打卡日历 |

### 4. 成就模块 (Achievement)

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 获取成就列表 | GET | `/miniprogram/achievement/list` | 获取用户成就列表 |
| 获取成就进度 | GET | `/miniprogram/achievement/progress` | 获取成就解锁进度 |
| 获取成就详情 | GET | `/miniprogram/achievement/:id` | 获取成就详细信息 |

### 5. 社区模块 (Community)

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 获取帖子列表 | GET | `/miniprogram/community/posts` | 获取社区帖子列表 |
| 创建帖子 | POST | `/miniprogram/community/posts` | 创建新帖子 |
| 获取帖子详情 | GET | `/miniprogram/community/posts/:id` | 获取帖子详细信息 |
| 点赞帖子 | POST | `/miniprogram/community/posts/:id/like` | 点赞/取消点赞帖子 |
| 创建评论 | POST | `/miniprogram/community/posts/:id/comments` | 创建帖子评论 |
| 获取评论列表 | GET | `/miniprogram/community/posts/:id/comments` | 获取帖子评论列表 |

### 6. 学习模块 (Learning)

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 获取学习内容 | GET | `/miniprogram/learning/content` | 获取学习内容列表 |
| 获取内容详情 | GET | `/miniprogram/learning/content/:id` | 获取学习内容详情 |
| 开始学习 | POST | `/miniprogram/learning/start` | 开始学习某个内容 |
| 更新学习进度 | PUT | `/miniprogram/learning/progress` | 更新学习进度 |
| 完成学习 | POST | `/miniprogram/learning/complete` | 完成学习内容 |
| 获取学习统计 | GET | `/miniprogram/learning/stats` | 获取学习统计数据 |

### 7. 评估模块 (Assessment)

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 获取评估问卷 | GET | `/miniprogram/assessment/questionnaire` | 获取评估问卷 |
| 提交评估 | POST | `/miniprogram/assessment/submit` | 提交评估答案 |
| 获取评估结果 | GET | `/miniprogram/assessment/result/:id` | 获取评估结果 |
| 获取评估历史 | GET | `/miniprogram/assessment/history` | 获取评估历史记录 |

### 8. 紧急求助模块 (Emergency)

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 创建求助 | POST | `/miniprogram/emergency/help` | 创建紧急求助 |
| 获取求助历史 | GET | `/miniprogram/emergency/help/history` | 获取求助历史 |
| 获取辅助资源 | GET | `/miniprogram/emergency/resources` | 获取紧急辅助资源 |
| 志愿者注册 | POST | `/miniprogram/emergency/volunteer/register` | 注册成为志愿者 |
| 响应求助 | POST | `/miniprogram/emergency/help/:id/respond` | 志愿者响应求助 |

## 通用响应格式

所有API响应都采用统一格式：

```json
{
  "code": 0,
  "data": {},
  "msg": "success"
}
```

### 响应字段说明

- `code`: 响应状态码，0表示成功，非0表示失败
- `data`: 响应数据，具体结构根据接口而定
- `msg`: 响应消息，成功时为"success"，失败时为错误描述

### 常见状态码

| 状态码 | 说明 |
|--------|------|
| 0 | 成功 |
| 7 | 参数错误 |
| 1001 | 用户未登录 |
| 1002 | Token过期 |
| 1003 | 权限不足 |
| 5000 | 服务器内部错误 |

## 分页参数

列表类接口支持分页，使用以下参数：

```json
{
  "page": 1,
  "pageSize": 10
}
```

分页响应格式：

```json
{
  "code": 0,
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "pageSize": 10
  },
  "msg": "success"
}
```

## 错误处理

### 客户端错误 (4xx)

- **400 Bad Request**: 请求参数错误
- **401 Unauthorized**: 未授权，需要登录
- **403 Forbidden**: 权限不足
- **404 Not Found**: 资源不存在
- **429 Too Many Requests**: 请求频率过高

### 服务器错误 (5xx)

- **500 Internal Server Error**: 服务器内部错误
- **502 Bad Gateway**: 网关错误
- **503 Service Unavailable**: 服务不可用

## 测试环境

### 运行测试

```bash
# 单元测试
go test ./test/... -v

# 性能测试
go test ./test/... -bench=. -benchmem

# 生成API文档
./scripts/generate_docs.sh

# 性能压力测试
./scripts/performance_test.sh
```

### Swagger文档

启动服务后访问：`http://localhost:8888/swagger/index.html`

## 版本历史

| 版本 | 日期 | 变更内容 |
|------|------|----------|
| v1.0.0 | 2025-01-23 | 初始版本，包含所有核心功能 |

## 联系方式

如有问题请联系开发团队或查看项目文档。 