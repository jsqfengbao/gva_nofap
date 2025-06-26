# 戒色助手小程序安全文档

## 概述

本文档描述了戒色助手微信小程序的安全架构、安全措施和最佳实践。我们采用了多层次的安全防护策略，确保用户数据和系统安全。

## 安全架构

### 1. 网络安全层
- **HTTPS加密**: 所有API通信使用HTTPS协议
- **TLS 1.2+**: 强制使用TLS 1.2或更高版本
- **HSTS**: 启用HTTP严格传输安全
- **证书固定**: 防止中间人攻击

### 2. 应用安全层
- **JWT认证**: 基于JSON Web Token的用户认证
- **Token黑名单**: 支持Token撤销和黑名单机制
- **会话管理**: 安全的会话生命周期管理
- **权限控制**: 基于角色的访问控制(RBAC)

### 3. 数据安全层
- **数据加密**: 敏感数据使用AES-256-GCM加密
- **数据脱敏**: 敏感信息显示时自动脱敏
- **数据备份**: 定期加密备份
- **数据保留**: 符合数据保护法规的数据保留策略

## 安全措施详解

### 1. 认证与授权

#### JWT Token安全
```yaml
jwt:
  blacklist-enabled: true        # 启用Token黑名单
  refresh-threshold: 24h         # Token刷新阈值
  max-fail-attempts: 10          # 最大失败尝试次数
  lockout-duration: 1h           # 锁定时长
```

#### 微信小程序认证
- 使用微信官方code2session接口
- 验证sessionKey的有效性
- 用户信息解密验证
- 防止重放攻击

### 2. 输入验证与防护

#### XSS防护
- 输入内容过滤危险脚本
- HTML实体编码
- CSP内容安全策略
- 输出时自动转义

#### SQL注入防护
- 参数化查询
- 输入模式匹配检测
- 危险SQL关键词过滤
- 数据库权限最小化

#### 请求验证
```yaml
input-validation:
  max-request-size: 10485760     # 最大请求大小(10MB)
  field-limits:                  # 字段长度限制
    username: 50
    nickname: 100
    content: 2000
    description: 500
    title: 200
```

### 3. 速率限制

#### API速率限制
```yaml
rate-limit:
  enabled: true
  window-size: 60s
  max-requests: 100
  auth-api:
    window-size: 60s
    max-requests: 10             # 认证API更严格的限制
  miniprogram-api:
    window-size: 60s
    max-requests: 60
```

#### IP限制策略
- 基于IP的请求频率限制
- 恶意IP自动封禁
- 白名单和黑名单机制
- 分布式限流支持

### 4. 数据保护

#### 敏感数据加密
```go
// 需要加密的字段
encrypted-fields: 
  - phone
  - email
  - real_name
  - id_card
  - address
```

#### 数据脱敏规则
```go
// 脱敏映射
masked-fields:
  phone: "phone"          # 138****5678
  email: "email"          # te***t@example.com
  real_name: "name"       # 张*丰
  id_card: "idcard"       # 110101********1234
  address: "address"      # 北京市***区
```

#### 数据库安全
- 连接加密
- 最小权限原则
- 定期权限审计
- 敏感操作日志记录

### 5. 安全响应头

系统自动添加以下安全响应头：

```http
X-Frame-Options: DENY
X-Content-Type-Options: nosniff
X-XSS-Protection: 1; mode=block
Strict-Transport-Security: max-age=31536000; includeSubDomains
Content-Security-Policy: default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval' https://res.wx.qq.com
Referrer-Policy: strict-origin-when-cross-origin
Permissions-Policy: geolocation=(), microphone=(), camera=()
```

### 6. 审计日志

#### 日志配置
```yaml
audit-log:
  enabled: true
  log-level: "info"
  log-path: "./log/audit.log"
  max-size: 100              # MB
  max-backups: 30
  max-age: 90                # days
  compress: true
```

#### 记录的安全事件
- 用户登录/登出
- 认证失败
- 权限违规尝试
- 敏感数据访问
- 系统配置变更
- 异常API调用

## 安全配置

### 1. 环境配置

#### 生产环境要求
- 使用强密码策略
- 启用防火墙
- 定期安全更新
- SSL证书配置
- 数据库访问限制

#### 配置文件安全
```bash
# 设置配置文件权限
chmod 600 config.yaml
chmod 600 .env

# 设置日志目录权限
chmod -R 700 log/
```

### 2. 密钥管理

#### 加密密钥
```yaml
encryption:
  secret-key: "nofap-miniprogram-secret-key-32"  # 32字节密钥
  salt: "nofap-salt-2025"                        # 盐值
  algorithm: "AES-256-GCM"                       # 加密算法
```

#### JWT密钥
- 使用强随机密钥
- 定期轮换密钥
- 密钥分离存储
- 密钥访问控制

## 安全测试

### 1. 自动化测试

运行安全测试套件：
```bash
# 运行安全测试
go test ./test/security_test.go -v

# 运行性能基准测试
go test ./test/security_test.go -bench=. -benchmem
```

### 2. 渗透测试

定期进行以下安全测试：
- SQL注入测试
- XSS攻击测试
- CSRF攻击测试
- 认证绕过测试
- 权限提升测试
- 敏感信息泄露测试

### 3. 漏洞扫描

使用工具进行定期扫描：
- 静态代码分析
- 依赖漏洞扫描
- 容器安全扫描
- 网络端口扫描

## 事件响应

### 1. 安全事件分类

#### 高危事件
- 数据泄露
- 系统入侵
- 权限提升
- 恶意代码注入

#### 中危事件
- 异常登录
- 频繁失败尝试
- 可疑API调用
- 配置变更

#### 低危事件
- 正常认证失败
- 常规错误
- 性能异常

### 2. 响应流程

1. **检测**: 自动监控和告警
2. **分析**: 确定事件类型和影响范围
3. **遏制**: 立即阻止进一步损害
4. **根除**: 清除威胁源
5. **恢复**: 恢复正常服务
6. **总结**: 事后分析和改进

### 3. 联系方式

安全事件报告：
- 邮箱: security@nofap-app.com
- 电话: +86-xxx-xxxx-xxxx
- 24小时应急响应

## 合规要求

### 1. 数据保护法规
- 《网络安全法》
- 《数据安全法》
- 《个人信息保护法》
- 微信小程序安全规范

### 2. 行业标准
- ISO 27001信息安全管理
- OWASP安全开发指南
- 等保2.0合规要求

### 3. 隐私保护
- 用户授权机制
- 数据最小化原则
- 用户数据删除权
- 隐私政策透明化

## 安全最佳实践

### 1. 开发安全
- 安全编码规范
- 代码审查机制
- 安全测试集成
- 依赖安全管理

### 2. 运维安全
- 定期安全更新
- 访问权限管理
- 监控和告警
- 备份和恢复

### 3. 用户安全
- 安全使用指南
- 隐私设置教育
- 异常行为提醒
- 安全功能推广

## 更新记录

| 版本 | 日期 | 变更内容 |
|------|------|----------|
| v1.0.0 | 2025-01-23 | 初始版本，包含完整安全架构 |

---

**注意**: 本文档包含敏感的安全配置信息，请妥善保管，仅限授权人员查看。 