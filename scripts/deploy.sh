#!/bin/bash

# 戒色助手生产环境部署脚本
# NoFap Assistant Production Deployment Script

set -e  # 遇到错误立即退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查是否为root用户
check_root() {
    if [[ $EUID -eq 0 ]]; then
        log_error "请不要使用root用户运行此脚本"
        exit 1
    fi
}

# 检查系统要求
check_system_requirements() {
    log_info "检查系统要求..."
    
    # 检查Docker
    if ! command -v docker &> /dev/null; then
        log_error "Docker 未安装，请先安装 Docker"
        exit 1
    fi
    
    # 检查Docker Compose
    if ! command -v docker-compose &> /dev/null; then
        log_error "Docker Compose 未安装，请先安装 Docker Compose"
        exit 1
    fi
    
    # 检查Git
    if ! command -v git &> /dev/null; then
        log_error "Git 未安装，请先安装 Git"
        exit 1
    fi
    
    # 检查磁盘空间（至少需要5GB）
    available_space=$(df / | awk 'NR==2{print $4}')
    required_space=5242880  # 5GB in KB
    
    if [[ $available_space -lt $required_space ]]; then
        log_error "磁盘空间不足，至少需要5GB可用空间"
        exit 1
    fi
    
    log_success "系统要求检查通过"
}

# 检查环境变量文件
check_env_file() {
    log_info "检查环境变量配置..."
    
    if [[ ! -f "deploy/production/.env" ]]; then
        if [[ -f "deploy/production/env.example" ]]; then
            log_warning ".env 文件不存在，从 env.example 复制"
            cp deploy/production/env.example deploy/production/.env
            log_error "请编辑 deploy/production/.env 文件，配置正确的环境变量"
            exit 1
        else
            log_error "环境变量配置文件不存在"
            exit 1
        fi
    fi
    
    # 检查关键环境变量
    source deploy/production/.env
    
    required_vars=(
        "DB_PASSWORD"
        "REDIS_PASSWORD" 
        "JWT_SECRET"
        "AES_KEY"
        "WECHAT_APPID"
        "WECHAT_APPSECRET"
    )
    
    for var in "${required_vars[@]}"; do
        if [[ -z "${!var}" || "${!var}" == "your_"* ]]; then
            log_error "环境变量 $var 未正确配置"
            exit 1
        fi
    done
    
    log_success "环境变量配置检查通过"
}

# 创建必要的目录
create_directories() {
    log_info "创建必要的目录..."
    
    directories=(
        "deploy/production/mysql/conf.d"
        "deploy/production/mysql/init"
        "deploy/production/redis"
        "deploy/production/nginx/conf.d"
        "deploy/production/ssl"
        "deploy/production/prometheus"
        "deploy/production/grafana/provisioning"
        "deploy/production/loki"
        "deploy/production/promtail"
        "logs"
        "backups"
    )
    
    for dir in "${directories[@]}"; do
        mkdir -p "$dir"
        log_info "创建目录: $dir"
    done
    
    log_success "目录创建完成"
}

# 生成SSL证书（自签名）
generate_ssl_certificates() {
    log_info "生成SSL证书..."
    
    ssl_dir="deploy/production/ssl"
    
    # 检查是否已存在证书
    if [[ -f "$ssl_dir/admin.yourdomain.com.crt" && -f "$ssl_dir/api.yourdomain.com.crt" ]]; then
        log_info "SSL证书已存在，跳过生成"
        return
    fi
    
    # 生成管理端证书
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
        -keyout "$ssl_dir/admin.yourdomain.com.key" \
        -out "$ssl_dir/admin.yourdomain.com.crt" \
        -subj "/C=CN/ST=Beijing/L=Beijing/O=NoFap/CN=admin.yourdomain.com"
    
    # 生成API证书
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
        -keyout "$ssl_dir/api.yourdomain.com.key" \
        -out "$ssl_dir/api.yourdomain.com.crt" \
        -subj "/C=CN/ST=Beijing/L=Beijing/O=NoFap/CN=api.yourdomain.com"
    
    # 设置权限
    chmod 600 "$ssl_dir"/*.key
    chmod 644 "$ssl_dir"/*.crt
    
    log_success "SSL证书生成完成"
    log_warning "生产环境请使用有效的SSL证书替换自签名证书"
}

# 配置数据库
setup_database() {
    log_info "配置数据库..."
    
    # 创建数据库初始化脚本
    cat > deploy/production/mysql/init/01-init.sql << EOF
-- 创建数据库
CREATE DATABASE IF NOT EXISTS nofap_production CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建用户
CREATE USER IF NOT EXISTS 'nofap_user'@'%' IDENTIFIED BY '${DB_PASSWORD}';
GRANT ALL PRIVILEGES ON nofap_production.* TO 'nofap_user'@'%';
FLUSH PRIVILEGES;

-- 设置时区
SET time_zone = '+08:00';
EOF
    
    # 创建MySQL配置文件
    cat > deploy/production/mysql/conf.d/mysql.cnf << EOF
[mysqld]
# 基础配置
default-authentication-plugin=mysql_native_password
character-set-server=utf8mb4
collation-server=utf8mb4_unicode_ci
default-time-zone='+08:00'

# 性能优化
innodb_buffer_pool_size=1G
innodb_log_file_size=256M
innodb_flush_log_at_trx_commit=2
innodb_flush_method=O_DIRECT

# 连接配置
max_connections=200
max_connect_errors=100000
wait_timeout=28800
interactive_timeout=28800

# 查询缓存
query_cache_type=1
query_cache_size=64M
query_cache_limit=2M

# 慢查询日志
slow_query_log=1
slow_query_log_file=/var/lib/mysql/slow.log
long_query_time=2

# 二进制日志
log-bin=mysql-bin
binlog_format=ROW
expire_logs_days=7

# 安全设置
skip-name-resolve
sql_mode=STRICT_TRANS_TABLES,NO_ZERO_DATE,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO

[mysql]
default-character-set=utf8mb4

[client]
default-character-set=utf8mb4
EOF
    
    log_success "数据库配置完成"
}

# 配置Redis
setup_redis() {
    log_info "配置Redis..."
    
    cat > deploy/production/redis/redis.conf << EOF
# Redis 生产配置
bind 0.0.0.0
port 6379
requirepass ${REDIS_PASSWORD}

# 内存配置
maxmemory 512mb
maxmemory-policy allkeys-lru

# 持久化配置
save 900 1
save 300 10
save 60 10000
appendonly yes
appendfsync everysec

# 安全配置
protected-mode yes
timeout 300

# 日志配置
loglevel notice
logfile ""

# 网络配置
tcp-keepalive 300
tcp-backlog 511

# 性能配置
databases 16
EOF
    
    log_success "Redis配置完成"
}

# 配置监控
setup_monitoring() {
    log_info "配置监控系统..."
    
    # Prometheus配置
    cat > deploy/production/prometheus/prometheus.yml << EOF
global:
  scrape_interval: 15s
  evaluation_interval: 15s

rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']

  - job_name: 'nofap-backend'
    static_configs:
      - targets: ['backend:8888']
    metrics_path: '/api/metrics'

  - job_name: 'mysql'
    static_configs:
      - targets: ['mysql:3306']

  - job_name: 'redis'
    static_configs:
      - targets: ['redis:6379']

  - job_name: 'nginx'
    static_configs:
      - targets: ['frontend:80']
EOF
    
    # Loki配置
    cat > deploy/production/loki/loki-config.yml << EOF
auth_enabled: false

server:
  http_listen_port: 3100

ingester:
  lifecycler:
    address: 127.0.0.1
    ring:
      kvstore:
        store: inmemory
      replication_factor: 1
    final_sleep: 0s
  chunk_idle_period: 1h
  max_chunk_age: 1h
  chunk_target_size: 1048576
  chunk_retain_period: 30s

schema_config:
  configs:
    - from: 2020-10-24
      store: boltdb-shipper
      object_store: filesystem
      schema: v11
      index:
        prefix: index_
        period: 24h

storage_config:
  boltdb_shipper:
    active_index_directory: /loki/boltdb-shipper-active
    cache_location: /loki/boltdb-shipper-cache
    cache_ttl: 24h
    shared_store: filesystem
  filesystem:
    directory: /loki/chunks

limits_config:
  enforce_metric_name: false
  reject_old_samples: true
  reject_old_samples_max_age: 168h

chunk_store_config:
  max_look_back_period: 0s

table_manager:
  retention_deletes_enabled: false
  retention_period: 0s
EOF
    
    # Promtail配置
    cat > deploy/production/promtail/promtail-config.yml << EOF
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

clients:
  - url: http://loki:3100/loki/api/v1/push

scrape_configs:
  - job_name: nofap-backend
    static_configs:
      - targets:
          - localhost
        labels:
          job: nofap-backend
          __path__: /var/log/nofap/*.log

  - job_name: nginx
    static_configs:
      - targets:
          - localhost
        labels:
          job: nginx
          __path__: /var/log/nginx/*.log

  - job_name: docker
    static_configs:
      - targets:
          - localhost
        labels:
          job: docker
          __path__: /var/lib/docker/containers/*/*.log
EOF
    
    log_success "监控系统配置完成"
}

# 构建和启动服务
deploy_services() {
    log_info "构建和启动服务..."
    
    cd deploy/production
    
    # 拉取最新镜像
    docker-compose -f docker-compose.prod.yaml pull
    
    # 构建自定义镜像
    docker-compose -f docker-compose.prod.yaml build --no-cache
    
    # 启动服务
    docker-compose -f docker-compose.prod.yaml up -d
    
    cd ../..
    
    log_success "服务启动完成"
}

# 等待服务就绪
wait_for_services() {
    log_info "等待服务就绪..."
    
    # 等待数据库
    log_info "等待数据库启动..."
    timeout=60
    while ! docker exec nofap-mysql-prod mysqladmin ping -h localhost --silent && [ $timeout -gt 0 ]; do
        sleep 2
        timeout=$((timeout-2))
    done
    
    if [ $timeout -le 0 ]; then
        log_error "数据库启动超时"
        exit 1
    fi
    
    # 等待Redis
    log_info "等待Redis启动..."
    timeout=30
    while ! docker exec nofap-redis-prod redis-cli ping && [ $timeout -gt 0 ]; do
        sleep 2
        timeout=$((timeout-2))
    done
    
    if [ $timeout -le 0 ]; then
        log_error "Redis启动超时"
        exit 1
    fi
    
    # 等待后端服务
    log_info "等待后端服务启动..."
    timeout=60
    while ! curl -f http://localhost:8888/api/health &>/dev/null && [ $timeout -gt 0 ]; do
        sleep 2
        timeout=$((timeout-2))
    done
    
    if [ $timeout -le 0 ]; then
        log_error "后端服务启动超时"
        exit 1
    fi
    
    # 等待前端服务
    log_info "等待前端服务启动..."
    timeout=30
    while ! curl -f http://localhost/health &>/dev/null && [ $timeout -gt 0 ]; do
        sleep 2
        timeout=$((timeout-2))
    done
    
    if [ $timeout -le 0 ]; then
        log_error "前端服务启动超时"
        exit 1
    fi
    
    log_success "所有服务已就绪"
}

# 运行数据库迁移
run_migrations() {
    log_info "运行数据库迁移..."
    
    # 等待一下确保数据库完全就绪
    sleep 10
    
    # 执行数据库迁移
    docker exec nofap-backend-prod ./main migrate
    
    log_success "数据库迁移完成"
}

# 健康检查
health_check() {
    log_info "执行健康检查..."
    
    services=(
        "http://localhost:8888/api/health:后端API"
        "http://localhost/health:前端服务"
        "http://localhost:9090:Prometheus"
        "http://localhost:3000:Grafana"
    )
    
    for service in "${services[@]}"; do
        url="${service%:*}"
        name="${service#*:}"
        
        if curl -f "$url" &>/dev/null; then
            log_success "$name 健康检查通过"
        else
            log_error "$name 健康检查失败"
        fi
    done
}

# 显示部署信息
show_deployment_info() {
    log_success "🎉 部署完成！"
    echo
    echo "=== 服务访问地址 ==="
    echo "管理后台: https://admin.yourdomain.com"
    echo "API接口: https://api.yourdomain.com"
    echo "Grafana监控: http://localhost:3000 (admin/${GRAFANA_PASSWORD})"
    echo "Prometheus: http://localhost:9090"
    echo
    echo "=== 重要提醒 ==="
    echo "1. 请将域名 admin.yourdomain.com 和 api.yourdomain.com 解析到此服务器"
    echo "2. 生产环境请使用有效的SSL证书替换自签名证书"
    echo "3. 请定期备份数据库和上传文件"
    echo "4. 请监控服务状态和系统资源使用情况"
    echo
    echo "=== 常用命令 ==="
    echo "查看服务状态: docker-compose -f deploy/production/docker-compose.prod.yaml ps"
    echo "查看日志: docker-compose -f deploy/production/docker-compose.prod.yaml logs -f [service]"
    echo "重启服务: docker-compose -f deploy/production/docker-compose.prod.yaml restart [service]"
    echo "停止服务: docker-compose -f deploy/production/docker-compose.prod.yaml down"
    echo
}

# 主函数
main() {
    log_info "开始部署戒色助手生产环境..."
    
    check_root
    check_system_requirements
    check_env_file
    create_directories
    generate_ssl_certificates
    setup_database
    setup_redis
    setup_monitoring
    deploy_services
    wait_for_services
    run_migrations
    health_check
    show_deployment_info
    
    log_success "部署流程全部完成！"
}

# 脚本入口
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi 