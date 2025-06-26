# 戒色助手部署指南
# NoFap Assistant Deployment Guide

本文档提供戒色助手微信小程序的完整部署指南，包括生产环境部署、域名配置、SSL证书配置等详细步骤。

## 目录

- [系统要求](#系统要求)
- [部署准备](#部署准备)
- [快速部署](#快速部署)
- [手动部署](#手动部署)
- [域名配置](#域名配置)
- [SSL证书配置](#ssl证书配置)
- [监控配置](#监控配置)
- [备份策略](#备份策略)
- [故障排除](#故障排除)
- [维护指南](#维护指南)

## 系统要求

### 硬件要求
- **CPU**: 2核心以上
- **内存**: 4GB以上（推荐8GB）
- **存储**: 50GB以上可用空间
- **网络**: 稳定的互联网连接

### 软件要求
- **操作系统**: Ubuntu 20.04+ / CentOS 8+ / Debian 11+
- **Docker**: 20.10+
- **Docker Compose**: 2.0+
- **Git**: 2.0+
- **OpenSSL**: 1.1.1+

### 端口要求
确保以下端口可用：
- `80`: HTTP服务
- `443`: HTTPS服务
- `3306`: MySQL数据库
- `6379`: Redis缓存
- `8888`: 后端API服务
- `3000`: Grafana监控
- `9090`: Prometheus监控

## 部署准备

### 1. 获取源代码
```bash
git clone https://github.com/yourusername/nofap-assistant.git
cd nofap-assistant
```

### 2. 安装Docker和Docker Compose
```bash
# Ubuntu/Debian
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER

# 安装Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/v2.20.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# 重新登录以使用户组生效
newgrp docker
```

### 3. 配置环境变量
```bash
# 复制环境变量模板
cp deploy/production/env.example deploy/production/.env

# 编辑环境变量
nano deploy/production/.env
```

### 4. 配置微信小程序
在微信公众平台配置以下信息：
- 服务器域名：`https://api.yourdomain.com`
- 业务域名：`https://admin.yourdomain.com`
- 获取AppID和AppSecret并填入环境变量

## 快速部署

使用自动化部署脚本：

```bash
# 设置脚本执行权限
chmod +x scripts/deploy.sh

# 运行部署脚本
./scripts/deploy.sh
```

部署脚本将自动完成：
1. 系统要求检查
2. 环境变量验证
3. 目录创建
4. SSL证书生成
5. 数据库配置
6. 服务启动
7. 健康检查

## 手动部署

如果需要手动控制部署过程：

### 1. 创建必要目录
```bash
mkdir -p deploy/production/{mysql/{conf.d,init},redis,nginx/conf.d,ssl,prometheus,grafana/provisioning,loki,promtail}
mkdir -p logs backups
```

### 2. 生成SSL证书
```bash
# 自签名证书（测试用）
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout deploy/production/ssl/admin.yourdomain.com.key \
    -out deploy/production/ssl/admin.yourdomain.com.crt \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=NoFap/CN=admin.yourdomain.com"

openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout deploy/production/ssl/api.yourdomain.com.key \
    -out deploy/production/ssl/api.yourdomain.com.crt \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=NoFap/CN=api.yourdomain.com"

# 设置权限
chmod 600 deploy/production/ssl/*.key
chmod 644 deploy/production/ssl/*.crt
```

### 3. 启动服务
```bash
cd deploy/production
docker-compose -f docker-compose.prod.yaml up -d
```

### 4. 等待服务启动
```bash
# 检查服务状态
docker-compose -f docker-compose.prod.yaml ps

# 查看日志
docker-compose -f docker-compose.prod.yaml logs -f
```

### 5. 运行数据库迁移
```bash
# 等待数据库完全启动后执行
sleep 30
docker exec nofap-backend-prod ./main migrate
```

## 域名配置

### 1. DNS解析配置
将以下域名解析到服务器IP：
- `admin.yourdomain.com` → 服务器IP
- `api.yourdomain.com` → 服务器IP

### 2. 域名配置示例
```bash
# 添加DNS记录（以阿里云为例）
A记录: admin.yourdomain.com → 1.2.3.4
A记录: api.yourdomain.com → 1.2.3.4
```

### 3. 验证域名解析
```bash
nslookup admin.yourdomain.com
nslookup api.yourdomain.com
```

## SSL证书配置

### 1. 使用Let's Encrypt（推荐）
```bash
# 安装Certbot
sudo apt-get update
sudo apt-get install certbot

# 获取证书
sudo certbot certonly --standalone \
    -d admin.yourdomain.com \
    -d api.yourdomain.com

# 复制证书到部署目录
sudo cp /etc/letsencrypt/live/admin.yourdomain.com/fullchain.pem \
    deploy/production/ssl/admin.yourdomain.com.crt
sudo cp /etc/letsencrypt/live/admin.yourdomain.com/privkey.pem \
    deploy/production/ssl/admin.yourdomain.com.key

sudo cp /etc/letsencrypt/live/api.yourdomain.com/fullchain.pem \
    deploy/production/ssl/api.yourdomain.com.crt
sudo cp /etc/letsencrypt/live/api.yourdomain.com/privkey.pem \
    deploy/production/ssl/api.yourdomain.com.key

# 设置权限
sudo chmod 644 deploy/production/ssl/*.crt
sudo chmod 600 deploy/production/ssl/*.key
sudo chown $USER:$USER deploy/production/ssl/*
```

### 2. 自动续期配置
```bash
# 添加cron任务
sudo crontab -e

# 添加以下行（每月1号凌晨2点检查续期）
0 2 1 * * /usr/bin/certbot renew --quiet && docker-compose -f /path/to/deploy/production/docker-compose.prod.yaml restart frontend
```

### 3. 使用自有证书
如果有自己的SSL证书：
```bash
# 复制证书文件
cp your-admin-cert.crt deploy/production/ssl/admin.yourdomain.com.crt
cp your-admin-key.key deploy/production/ssl/admin.yourdomain.com.key
cp your-api-cert.crt deploy/production/ssl/api.yourdomain.com.crt
cp your-api-key.key deploy/production/ssl/api.yourdomain.com.key

# 设置权限
chmod 644 deploy/production/ssl/*.crt
chmod 600 deploy/production/ssl/*.key
```

## 监控配置

### 1. 访问监控面板
- **Grafana**: http://your-server-ip:3000
  - 用户名: admin
  - 密码: 在.env文件中的GRAFANA_PASSWORD

- **Prometheus**: http://your-server-ip:9090

### 2. 配置Grafana数据源
1. 登录Grafana
2. 添加Prometheus数据源
3. URL: http://prometheus:9090
4. 导入预设仪表板

### 3. 设置告警
```bash
# 编辑Prometheus告警规则
nano deploy/production/prometheus/alert_rules.yml

# 配置告警通知
nano deploy/production/grafana/provisioning/notifiers.yml
```

## 备份策略

### 1. 自动备份配置
```bash
# 设置备份脚本权限
chmod +x scripts/backup.sh

# 添加定时备份任务
crontab -e

# 每天凌晨2点执行备份
0 2 * * * /path/to/nofap-assistant/scripts/backup.sh backup
```

### 2. 手动备份
```bash
# 执行完整备份
./scripts/backup.sh backup

# 列出可用备份
./scripts/backup.sh list

# 恢复备份
./scripts/backup.sh restore 20240123_140530
```

### 3. 云端备份配置
在.env文件中配置AWS S3：
```bash
BACKUP_S3_BUCKET=your-backup-bucket
BACKUP_S3_REGION=us-east-1
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
```

## 故障排除

### 1. 常见问题

#### 服务无法启动
```bash
# 检查Docker服务
sudo systemctl status docker

# 检查端口占用
sudo netstat -tlnp | grep :80
sudo netstat -tlnp | grep :443

# 查看容器日志
docker-compose -f deploy/production/docker-compose.prod.yaml logs [service_name]
```

#### 数据库连接失败
```bash
# 检查MySQL容器状态
docker ps | grep mysql

# 查看MySQL日志
docker logs nofap-mysql-prod

# 测试数据库连接
docker exec -it nofap-mysql-prod mysql -u root -p
```

#### 前端无法访问
```bash
# 检查Nginx配置
docker exec nofap-frontend-prod nginx -t

# 查看Nginx日志
docker logs nofap-frontend-prod

# 检查SSL证书
openssl x509 -in deploy/production/ssl/admin.yourdomain.com.crt -text -noout
```

### 2. 性能问题

#### 数据库性能优化
```bash
# 查看MySQL进程
docker exec nofap-mysql-prod mysqladmin processlist -u root -p

# 查看慢查询
docker exec nofap-mysql-prod mysql -u root -p -e "SHOW VARIABLES LIKE 'slow_query_log';"

# 优化表
docker exec nofap-mysql-prod mysql -u root -p -e "OPTIMIZE TABLE nofap_users;"
```

#### 内存使用监控
```bash
# 查看容器资源使用
docker stats

# 查看系统内存
free -h

# 查看磁盘使用
df -h
```

### 3. 日志分析
```bash
# 查看应用日志
docker exec nofap-backend-prod tail -f /var/log/nofap/app.log

# 查看访问日志
docker exec nofap-frontend-prod tail -f /var/log/nginx/access.log

# 查看错误日志
docker exec nofap-frontend-prod tail -f /var/log/nginx/error.log
```

## 维护指南

### 1. 定期维护任务

#### 每日检查
- [ ] 检查服务状态
- [ ] 查看错误日志
- [ ] 验证备份完成
- [ ] 监控资源使用

#### 每周检查
- [ ] 更新系统安全补丁
- [ ] 检查磁盘空间
- [ ] 清理旧日志文件
- [ ] 验证SSL证书有效期

#### 每月检查
- [ ] 数据库性能优化
- [ ] 备份恢复测试
- [ ] 安全漏洞扫描
- [ ] 监控配置更新

### 2. 更新部署
```bash
# 拉取最新代码
git pull origin main

# 重新构建镜像
docker-compose -f deploy/production/docker-compose.prod.yaml build --no-cache

# 滚动更新服务
docker-compose -f deploy/production/docker-compose.prod.yaml up -d
```

### 3. 扩容指南

#### 垂直扩容（增加资源）
```yaml
# 修改docker-compose.prod.yaml
services:
  backend:
    deploy:
      resources:
        limits:
          memory: 2G
          cpus: '2.0'
```

#### 水平扩容（增加实例）
```yaml
# 修改docker-compose.prod.yaml
services:
  backend:
    scale: 3  # 运行3个后端实例
```

### 4. 安全维护

#### 定期安全检查
```bash
# 扫描容器漏洞
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
  aquasec/trivy image nofap-backend-prod

# 检查开放端口
nmap -sT -O localhost

# 查看登录日志
sudo grep "Failed password" /var/log/auth.log
```

#### 更新密码
```bash
# 更新数据库密码
docker exec -it nofap-mysql-prod mysql -u root -p
# ALTER USER 'nofap_user'@'%' IDENTIFIED BY 'new_password';

# 更新Redis密码
# 修改deploy/production/.env中的REDIS_PASSWORD
# 重启Redis服务
```

## 小程序发布

### 1. 微信开发者工具配置
1. 下载并安装微信开发者工具
2. 导入frontend项目
3. 配置项目AppID
4. 设置服务器域名

### 2. 域名白名单配置
在微信公众平台配置：
- request合法域名：`https://api.yourdomain.com`
- socket合法域名：`wss://api.yourdomain.com`
- uploadFile合法域名：`https://api.yourdomain.com`
- downloadFile合法域名：`https://api.yourdomain.com`

### 3. 提交审核
1. 在开发者工具中上传代码
2. 在微信公众平台提交审核
3. 填写版本描述和功能说明
4. 等待审核结果

### 4. 发布上线
审核通过后：
1. 在微信公众平台点击发布
2. 设置版本号和发布说明
3. 确认发布

## 联系支持

如遇到部署问题，请：
1. 查看本文档的故障排除部分
2. 检查GitHub Issues
3. 联系技术支持：support@yourdomain.com

---

**注意**: 本部署指南适用于生产环境。在正式部署前，请在测试环境中充分验证所有功能。 