#!/bin/bash

# 戒色助手数据备份脚本
# NoFap Assistant Data Backup Script

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# 配置变量
BACKUP_DIR="/var/backups/nofap"
RETENTION_DAYS=30
DATE=$(date +%Y%m%d_%H%M%S)
COMPOSE_FILE="deploy/production/docker-compose.prod.yaml"

# 从环境变量文件加载配置
if [[ -f "deploy/production/.env" ]]; then
    source deploy/production/.env
fi

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

# 创建备份目录
create_backup_dir() {
    local backup_path="$BACKUP_DIR/$DATE"
    mkdir -p "$backup_path"
    echo "$backup_path"
}

# 备份MySQL数据库
backup_mysql() {
    local backup_path="$1"
    log_info "备份MySQL数据库..."
    
    # 检查容器是否运行
    if ! docker ps | grep -q "nofap-mysql-prod"; then
        log_error "MySQL容器未运行"
        return 1
    fi
    
    # 备份数据库
    docker exec nofap-mysql-prod mysqldump \
        --single-transaction \
        --routines \
        --triggers \
        --all-databases \
        -u root -p"${DB_ROOT_PASSWORD}" > "$backup_path/mysql_backup.sql"
    
    # 压缩备份文件
    gzip "$backup_path/mysql_backup.sql"
    
    log_success "MySQL数据库备份完成: mysql_backup.sql.gz"
}

# 备份Redis数据
backup_redis() {
    local backup_path="$1"
    log_info "备份Redis数据..."
    
    # 检查容器是否运行
    if ! docker ps | grep -q "nofap-redis-prod"; then
        log_error "Redis容器未运行"
        return 1
    fi
    
    # 创建Redis备份
    docker exec nofap-redis-prod redis-cli --rdb /data/backup.rdb
    
    # 复制备份文件
    docker cp nofap-redis-prod:/data/backup.rdb "$backup_path/redis_backup.rdb"
    
    # 压缩备份文件
    gzip "$backup_path/redis_backup.rdb"
    
    log_success "Redis数据备份完成: redis_backup.rdb.gz"
}

# 备份上传文件
backup_uploads() {
    local backup_path="$1"
    log_info "备份上传文件..."
    
    # 检查上传目录是否存在
    if docker exec nofap-backend-prod test -d /var/uploads/nofap; then
        # 创建上传文件的tar包
        docker exec nofap-backend-prod tar -czf /tmp/uploads_backup.tar.gz -C /var/uploads nofap
        
        # 复制备份文件
        docker cp nofap-backend-prod:/tmp/uploads_backup.tar.gz "$backup_path/uploads_backup.tar.gz"
        
        # 清理临时文件
        docker exec nofap-backend-prod rm -f /tmp/uploads_backup.tar.gz
        
        log_success "上传文件备份完成: uploads_backup.tar.gz"
    else
        log_warning "上传目录不存在，跳过文件备份"
    fi
}

# 备份配置文件
backup_configs() {
    local backup_path="$1"
    log_info "备份配置文件..."
    
    # 创建配置目录
    mkdir -p "$backup_path/configs"
    
    # 备份主要配置文件
    files_to_backup=(
        "server/config.prod.yaml"
        "deploy/production/.env"
        "deploy/production/docker-compose.prod.yaml"
        "web/nginx.prod.conf"
    )
    
    for file in "${files_to_backup[@]}"; do
        if [[ -f "$file" ]]; then
            cp "$file" "$backup_path/configs/"
            log_info "备份配置文件: $file"
        fi
    done
    
    # 备份Docker Compose配置目录
    if [[ -d "deploy/production" ]]; then
        tar -czf "$backup_path/configs/docker_configs.tar.gz" -C deploy/production \
            --exclude='.env' \
            --exclude='*.log' \
            --exclude='data' .
    fi
    
    log_success "配置文件备份完成"
}

# 备份日志文件
backup_logs() {
    local backup_path="$1"
    log_info "备份日志文件..."
    
    # 创建日志备份目录
    mkdir -p "$backup_path/logs"
    
    # 备份应用日志
    if docker exec nofap-backend-prod test -d /var/log/nofap; then
        docker exec nofap-backend-prod tar -czf /tmp/app_logs.tar.gz -C /var/log nofap
        docker cp nofap-backend-prod:/tmp/app_logs.tar.gz "$backup_path/logs/app_logs.tar.gz"
        docker exec nofap-backend-prod rm -f /tmp/app_logs.tar.gz
    fi
    
    # 备份Nginx日志
    if docker exec nofap-frontend-prod test -d /var/log/nginx; then
        docker exec nofap-frontend-prod tar -czf /tmp/nginx_logs.tar.gz -C /var/log nginx
        docker cp nofap-frontend-prod:/tmp/nginx_logs.tar.gz "$backup_path/logs/nginx_logs.tar.gz"
        docker exec nofap-frontend-prod rm -f /tmp/nginx_logs.tar.gz
    fi
    
    log_success "日志文件备份完成"
}

# 创建备份清单
create_manifest() {
    local backup_path="$1"
    log_info "创建备份清单..."
    
    cat > "$backup_path/MANIFEST.txt" << EOF
NoFap Assistant Backup Manifest
=====================================
Backup Date: $(date)
Backup Path: $backup_path
Retention: $RETENTION_DAYS days

Files Included:
EOF
    
    # 列出所有备份文件
    find "$backup_path" -type f -exec ls -lh {} \; >> "$backup_path/MANIFEST.txt"
    
    # 计算总大小
    total_size=$(du -sh "$backup_path" | cut -f1)
    echo "Total Backup Size: $total_size" >> "$backup_path/MANIFEST.txt"
    
    log_success "备份清单创建完成"
}

# 清理旧备份
cleanup_old_backups() {
    log_info "清理 $RETENTION_DAYS 天前的备份..."
    
    if [[ -d "$BACKUP_DIR" ]]; then
        find "$BACKUP_DIR" -type d -name "20*" -mtime +$RETENTION_DAYS -exec rm -rf {} \; 2>/dev/null || true
        
        # 统计剩余备份数量
        remaining_backups=$(find "$BACKUP_DIR" -type d -name "20*" | wc -l)
        log_success "清理完成，剩余 $remaining_backups 个备份"
    fi
}

# 上传到云存储（可选）
upload_to_cloud() {
    local backup_path="$1"
    
    # 检查是否配置了AWS S3
    if [[ -n "$BACKUP_S3_BUCKET" && -n "$AWS_ACCESS_KEY_ID" && -n "$AWS_SECRET_ACCESS_KEY" ]]; then
        log_info "上传备份到 S3..."
        
        # 检查AWS CLI是否安装
        if command -v aws &> /dev/null; then
            # 创建压缩包
            backup_archive="$BACKUP_DIR/nofap_backup_$DATE.tar.gz"
            tar -czf "$backup_archive" -C "$BACKUP_DIR" "$DATE"
            
            # 上传到S3
            aws s3 cp "$backup_archive" "s3://$BACKUP_S3_BUCKET/backups/" \
                --region "$BACKUP_S3_REGION" \
                --storage-class STANDARD_IA
            
            # 删除本地压缩包
            rm -f "$backup_archive"
            
            log_success "备份已上传到 S3"
        else
            log_warning "AWS CLI 未安装，跳过云端备份"
        fi
    fi
}

# 发送通知（可选）
send_notification() {
    local backup_path="$1"
    local total_size=$(du -sh "$backup_path" | cut -f1)
    
    # 如果配置了邮件，发送备份通知
    if [[ -n "$EMAIL_TO" && -n "$EMAIL_FROM" ]]; then
        local subject="NoFap Assistant Backup Completed - $DATE"
        local body="Backup completed successfully.

Backup Details:
- Date: $(date)
- Location: $backup_path
- Size: $total_size
- Retention: $RETENTION_DAYS days

Please verify the backup integrity."
        
        # 这里可以集成邮件发送逻辑
        log_info "备份通知已准备，可手动发送邮件通知"
    fi
}

# 验证备份完整性
verify_backup() {
    local backup_path="$1"
    log_info "验证备份完整性..."
    
    local errors=0
    
    # 检查MySQL备份
    if [[ -f "$backup_path/mysql_backup.sql.gz" ]]; then
        if gzip -t "$backup_path/mysql_backup.sql.gz"; then
            log_success "MySQL备份文件完整"
        else
            log_error "MySQL备份文件损坏"
            errors=$((errors + 1))
        fi
    fi
    
    # 检查Redis备份
    if [[ -f "$backup_path/redis_backup.rdb.gz" ]]; then
        if gzip -t "$backup_path/redis_backup.rdb.gz"; then
            log_success "Redis备份文件完整"
        else
            log_error "Redis备份文件损坏"
            errors=$((errors + 1))
        fi
    fi
    
    # 检查上传文件备份
    if [[ -f "$backup_path/uploads_backup.tar.gz" ]]; then
        if tar -tzf "$backup_path/uploads_backup.tar.gz" >/dev/null; then
            log_success "上传文件备份完整"
        else
            log_error "上传文件备份损坏"
            errors=$((errors + 1))
        fi
    fi
    
    if [[ $errors -eq 0 ]]; then
        log_success "所有备份文件验证通过"
        return 0
    else
        log_error "发现 $errors 个备份文件有问题"
        return 1
    fi
}

# 主备份函数
main_backup() {
    log_info "开始备份 NoFap Assistant 数据..."
    
    # 创建备份目录
    local backup_path=$(create_backup_dir)
    log_info "备份目录: $backup_path"
    
    # 执行各项备份
    backup_mysql "$backup_path"
    backup_redis "$backup_path"
    backup_uploads "$backup_path"
    backup_configs "$backup_path"
    backup_logs "$backup_path"
    
    # 创建备份清单
    create_manifest "$backup_path"
    
    # 验证备份
    if verify_backup "$backup_path"; then
        log_success "备份验证通过"
        
        # 上传到云存储
        upload_to_cloud "$backup_path"
        
        # 发送通知
        send_notification "$backup_path"
        
        # 清理旧备份
        cleanup_old_backups
        
        log_success "备份任务完成: $backup_path"
    else
        log_error "备份验证失败，请检查备份文件"
        exit 1
    fi
}

# 恢复功能
restore_backup() {
    local backup_date="$1"
    
    if [[ -z "$backup_date" ]]; then
        log_error "请指定要恢复的备份日期 (格式: YYYYMMDD_HHMMSS)"
        exit 1
    fi
    
    local restore_path="$BACKUP_DIR/$backup_date"
    
    if [[ ! -d "$restore_path" ]]; then
        log_error "备份目录不存在: $restore_path"
        exit 1
    fi
    
    log_warning "警告：恢复操作将覆盖当前数据！"
    read -p "确认要恢复备份 $backup_date 吗？(yes/no): " confirm
    
    if [[ "$confirm" != "yes" ]]; then
        log_info "恢复操作已取消"
        exit 0
    fi
    
    log_info "开始恢复备份: $backup_date"
    
    # 停止服务
    log_info "停止服务..."
    docker-compose -f "$COMPOSE_FILE" down
    
    # 恢复MySQL
    if [[ -f "$restore_path/mysql_backup.sql.gz" ]]; then
        log_info "恢复MySQL数据..."
        docker-compose -f "$COMPOSE_FILE" up -d mysql
        sleep 30
        
        gunzip -c "$restore_path/mysql_backup.sql.gz" | \
            docker exec -i nofap-mysql-prod mysql -u root -p"${DB_ROOT_PASSWORD}"
    fi
    
    # 恢复Redis
    if [[ -f "$restore_path/redis_backup.rdb.gz" ]]; then
        log_info "恢复Redis数据..."
        docker-compose -f "$COMPOSE_FILE" up -d redis
        sleep 10
        
        gunzip -c "$restore_path/redis_backup.rdb.gz" > /tmp/restore.rdb
        docker cp /tmp/restore.rdb nofap-redis-prod:/data/dump.rdb
        docker restart nofap-redis-prod
        rm -f /tmp/restore.rdb
    fi
    
    # 恢复上传文件
    if [[ -f "$restore_path/uploads_backup.tar.gz" ]]; then
        log_info "恢复上传文件..."
        docker-compose -f "$COMPOSE_FILE" up -d backend
        sleep 10
        
        docker cp "$restore_path/uploads_backup.tar.gz" nofap-backend-prod:/tmp/
        docker exec nofap-backend-prod tar -xzf /tmp/uploads_backup.tar.gz -C /var/uploads/
        docker exec nofap-backend-prod rm -f /tmp/uploads_backup.tar.gz
    fi
    
    # 重启所有服务
    log_info "重启所有服务..."
    docker-compose -f "$COMPOSE_FILE" up -d
    
    log_success "备份恢复完成"
}

# 列出可用备份
list_backups() {
    log_info "可用的备份列表:"
    
    if [[ -d "$BACKUP_DIR" ]]; then
        for backup in $(find "$BACKUP_DIR" -type d -name "20*" | sort -r); do
            local backup_name=$(basename "$backup")
            local backup_size=$(du -sh "$backup" | cut -f1)
            local backup_date=$(echo "$backup_name" | sed 's/_/ /')
            
            echo "  $backup_name ($backup_size) - $backup_date"
        done
    else
        log_warning "备份目录不存在: $BACKUP_DIR"
    fi
}

# 使用说明
usage() {
    echo "使用方法: $0 [选项]"
    echo
    echo "选项:"
    echo "  backup              执行完整备份"
    echo "  restore <date>      恢复指定日期的备份"
    echo "  list               列出可用的备份"
    echo "  cleanup            清理过期备份"
    echo "  help               显示此帮助信息"
    echo
    echo "示例:"
    echo "  $0 backup"
    echo "  $0 restore 20240123_140530"
    echo "  $0 list"
}

# 主函数
main() {
    case "${1:-backup}" in
        "backup")
            main_backup
            ;;
        "restore")
            restore_backup "$2"
            ;;
        "list")
            list_backups
            ;;
        "cleanup")
            cleanup_old_backups
            ;;
        "help"|"-h"|"--help")
            usage
            ;;
        *)
            log_error "未知选项: $1"
            usage
            exit 1
            ;;
    esac
}

# 脚本入口
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi 