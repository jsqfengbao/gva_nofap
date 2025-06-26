# 数据库迁移测试说明

## 测试目标
验证小程序数据库表结构能够正确创建，索引正确设置，基础数据正确初始化。

## 测试步骤

### 1. 启动服务验证表创建
```bash
cd server
go run main.go
```

服务启动时会自动执行数据库迁移，检查日志输出：
- ✅ "小程序数据表初始化成功!"
- ✅ "数据库索引创建完成!"
- ✅ "小程序基础数据初始化完成!"

### 2. 验证表结构创建
登录MySQL数据库检查表是否正确创建：

```sql
-- 查看所有小程序相关表
SHOW TABLES LIKE '%wx_users%';
SHOW TABLES LIKE '%abstinence_records%';
SHOW TABLES LIKE '%daily_checkins%';
SHOW TABLES LIKE '%assessment_results%';
SHOW TABLES LIKE '%community_%';
SHOW TABLES LIKE '%emergency_helps%';
SHOW TABLES LIKE '%help_responses%';
SHOW TABLES LIKE '%learning_%';
SHOW TABLES LIKE '%achievements%';
SHOW TABLES LIKE '%user_achievements%';
```

### 3. 验证表结构字段
```sql
-- 检查关键表结构
DESC wx_users;
DESC abstinence_records;
DESC daily_checkins;
DESC community_posts;
DESC achievements;
```

### 4. 验证索引创建
```sql
-- 查看关键表的索引
SHOW INDEX FROM wx_users;
SHOW INDEX FROM daily_checkins;
SHOW INDEX FROM community_posts;
SHOW INDEX FROM achievements;
```

### 5. 验证基础数据初始化
```sql
-- 检查成就系统基础数据
SELECT COUNT(*) as achievement_count FROM achievements;
SELECT name, category, type, rarity FROM achievements ORDER BY display_order LIMIT 10;
```

## 预期结果

### 表创建验证
应该创建12个核心业务表：
1. `wx_users` - 微信用户表
2. `abstinence_records` - 戒色记录表  
3. `daily_checkins` - 每日打卡表
4. `assessment_results` - 评估结果表
5. `community_posts` - 社区动态表
6. `community_comments` - 社区评论表
7. `community_likes` - 社区点赞表
8. `emergency_helps` - 紧急求助表
9. `help_responses` - 求助响应表
10. `learning_contents` - 学习内容表
11. `user_learning_records` - 用户学习记录表
12. `achievements` - 成就系统表
13. `user_achievements` - 用户成就表

### 索引验证
关键索引应正确创建：
- `wx_users`: openid唯一索引, status索引
- `daily_checkins`: (user_id, checkin_date)复合索引
- `community_posts`: (category, status)复合索引
- `achievements`: (category, is_active)复合索引

### 基础数据验证
- 成就系统应初始化20个基础成就
- 成就分类包含：打卡(1)、等级(2)、社区(3)、学习(4)、特殊(5)
- 成就稀有度包含：普通(1)、稀有(2)、史诗(3)、传说(4)

## 常见问题排查

### 1. 表创建失败
- 检查数据库连接配置
- 确认MySQL版本 >= 8.0
- 检查数据库用户权限

### 2. 索引创建失败
- 检查字段名是否正确
- 确认字段类型支持索引
- 检查索引名是否重复

### 3. 基础数据初始化失败
- 检查外键关联是否正确
- 确认JSON字段格式正确
- 检查必填字段是否有值

## 数据库性能验证

### 查询性能测试
```sql
-- 测试用户查询性能
EXPLAIN SELECT * FROM wx_users WHERE openid = 'test_openid';

-- 测试打卡记录查询性能  
EXPLAIN SELECT * FROM daily_checkins WHERE user_id = 1 AND checkin_date >= '2024-01-01';

-- 测试社区动态查询性能
EXPLAIN SELECT * FROM community_posts WHERE category = 1 AND status = 1 ORDER BY created_at DESC LIMIT 20;
```

预期所有查询都应使用索引，type为ref或range，不应出现全表扫描。 