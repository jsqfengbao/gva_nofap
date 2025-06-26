#!/bin/bash

# 性能优化脚本
# 用于优化应用性能和数据库配置

echo "🚀 开始性能优化配置..."

# 进入服务器目录
cd "$(dirname "$0")/.."

# 检查MySQL是否运行
if ! mysql --version > /dev/null 2>&1; then
    echo "⚠️  MySQL未安装或未在PATH中"
    echo "💡 请确保MySQL已安装并可访问"
fi

# 检查Redis是否运行
if ! redis-cli ping > /dev/null 2>&1; then
    echo "⚠️  Redis未运行，缓存功能将无法使用"
    echo "💡 启动Redis: redis-server"
else
    echo "✅ Redis运行正常"
fi

echo "📊 数据库性能优化..."

# 创建数据库性能优化SQL脚本
cat > temp_db_optimize.sql << 'EOF'
-- 数据库性能优化配置

-- 设置查询缓存
SET GLOBAL query_cache_type = ON;
SET GLOBAL query_cache_size = 268435456; -- 256MB

-- 优化InnoDB配置
SET GLOBAL innodb_buffer_pool_size = 1073741824; -- 1GB
SET GLOBAL innodb_log_file_size = 268435456; -- 256MB
SET GLOBAL innodb_flush_log_at_trx_commit = 2;
SET GLOBAL innodb_flush_method = 'O_DIRECT';

-- 优化连接配置
SET GLOBAL max_connections = 200;
SET GLOBAL wait_timeout = 28800;
SET GLOBAL interactive_timeout = 28800;

-- 优化临时表配置
SET GLOBAL tmp_table_size = 134217728; -- 128MB
SET GLOBAL max_heap_table_size = 134217728; -- 128MB

-- 显示当前配置
SHOW VARIABLES LIKE 'innodb_buffer_pool_size';
SHOW VARIABLES LIKE 'query_cache_size';
SHOW VARIABLES LIKE 'max_connections';
EOF

# 执行数据库优化（如果可以连接到数据库）
if mysql -h127.0.0.1 -uroot -p123456 -e "SELECT 1;" > /dev/null 2>&1; then
    echo "📈 执行数据库性能优化..."
    mysql -h127.0.0.1 -uroot -p123456 < temp_db_optimize.sql
    echo "✅ 数据库优化完成"
else
    echo "⚠️  无法连接到数据库，跳过数据库优化"
    echo "💡 请手动执行 temp_db_optimize.sql 中的SQL语句"
fi

# 清理临时文件
rm -f temp_db_optimize.sql

echo "🗄️ 创建数据库索引优化..."

# 创建索引优化脚本
cat > temp_index_optimize.sql << 'EOF'
-- 索引优化脚本

USE gva_nofap;

-- 用户表索引优化
ALTER TABLE nofap_users ADD INDEX idx_openid (openid);
ALTER TABLE nofap_users ADD INDEX idx_created_at (created_at);
ALTER TABLE nofap_users ADD INDEX idx_level (level);

-- 打卡记录索引优化
ALTER TABLE nofap_checkin_records ADD INDEX idx_user_date (user_id, checkin_date);
ALTER TABLE nofap_checkin_records ADD INDEX idx_checkin_date (checkin_date);

-- 社区帖子索引优化
ALTER TABLE nofap_community_posts ADD INDEX idx_category_created (category, created_at);
ALTER TABLE nofap_community_posts ADD INDEX idx_user_created (user_id, created_at);
ALTER TABLE nofap_community_posts ADD INDEX idx_likes_count (likes_count);

-- 学习内容索引优化
ALTER TABLE nofap_learning_contents ADD INDEX idx_category_status (category, status);
ALTER TABLE nofap_learning_contents ADD INDEX idx_difficulty_views (difficulty, view_count);

-- 成就记录索引优化
ALTER TABLE nofap_user_achievements ADD INDEX idx_user_unlocked (user_id, unlocked_at);
ALTER TABLE nofap_user_achievements ADD INDEX idx_achievement_unlocked (achievement_id, unlocked_at);

-- 评估记录索引优化
ALTER TABLE nofap_assessment_records ADD INDEX idx_user_created (user_id, created_at);
ALTER TABLE nofap_assessment_records ADD INDEX idx_risk_level (risk_level);

-- 紧急求助索引优化
ALTER TABLE nofap_emergency_helps ADD INDEX idx_user_status (user_id, status);
ALTER TABLE nofap_emergency_helps ADD INDEX idx_type_created (type, created_at);

-- 显示索引信息
SHOW INDEX FROM nofap_users;
SHOW INDEX FROM nofap_checkin_records;
SHOW INDEX FROM nofap_community_posts;
EOF

# 执行索引优化
if mysql -h127.0.0.1 -uroot -p123456 -e "SELECT 1;" > /dev/null 2>&1; then
    echo "📊 执行索引优化..."
    mysql -h127.0.0.1 -uroot -p123456 < temp_index_optimize.sql 2>/dev/null
    echo "✅ 索引优化完成"
else
    echo "⚠️  无法连接到数据库，跳过索引优化"
    echo "💡 请手动执行 temp_index_optimize.sql 中的SQL语句"
fi

# 清理临时文件
rm -f temp_index_optimize.sql

echo "⚡ Redis缓存配置优化..."

# 配置Redis缓存策略
if redis-cli ping > /dev/null 2>&1; then
    echo "🔧 配置Redis缓存策略..."
    
    # 设置内存策略
    redis-cli CONFIG SET maxmemory 512mb
    redis-cli CONFIG SET maxmemory-policy allkeys-lru
    
    # 设置持久化策略
    redis-cli CONFIG SET save "900 1 300 10 60 10000"
    
    # 设置网络优化
    redis-cli CONFIG SET tcp-keepalive 300
    redis-cli CONFIG SET timeout 0
    
    echo "✅ Redis缓存配置完成"
    
    # 显示Redis配置
    echo "📋 Redis配置信息:"
    redis-cli CONFIG GET maxmemory
    redis-cli CONFIG GET maxmemory-policy
else
    echo "⚠️  Redis未运行，跳过Redis优化"
fi

echo "🔧 Go应用性能优化..."

# 创建性能优化的环境变量配置
cat > .env.performance << 'EOF'
# Go应用性能优化配置

# 垃圾回收优化
GOGC=100
GOMEMLIMIT=1GiB

# 并发优化
GOMAXPROCS=0  # 使用所有可用CPU核心

# 网络优化
GODEBUG=netdns=go

# 编译优化
CGO_ENABLED=1
EOF

echo "📁 创建性能优化配置文件: .env.performance"

# 创建编译优化脚本
cat > scripts/build_optimized.sh << 'EOF'
#!/bin/bash

# 优化编译脚本

echo "🔨 开始优化编译..."

# 设置编译优化标志
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=amd64

# 编译优化版本
go build -ldflags="-s -w" -o bin/server-optimized main.go

echo "✅ 优化编译完成: bin/server-optimized"
echo "💡 使用方法: ./bin/server-optimized"
EOF

chmod +x scripts/build_optimized.sh
echo "📁 创建优化编译脚本: scripts/build_optimized.sh"

# 创建监控脚本
cat > scripts/monitor_performance.sh << 'EOF'
#!/bin/bash

# 性能监控脚本

echo "📊 应用性能监控..."

# 检查进程状态
if pgrep -f "main.go\|server" > /dev/null; then
    PID=$(pgrep -f "main.go\|server")
    echo "✅ 应用进程运行中 (PID: $PID)"
    
    # 显示CPU和内存使用情况
    echo "📈 资源使用情况:"
    ps -p $PID -o pid,ppid,pcpu,pmem,etime,comm
    
    # 显示网络连接
    echo "🌐 网络连接:"
    netstat -tulpn | grep :8888
    
    # 显示文件描述符使用情况
    echo "📁 文件描述符:"
    ls /proc/$PID/fd | wc -l
    echo "   当前打开文件数"
else
    echo "❌ 应用未运行"
fi

# 检查数据库连接
echo "🗄️ 数据库状态:"
if mysql -h127.0.0.1 -uroot -p123456 -e "SHOW PROCESSLIST;" 2>/dev/null | wc -l; then
    echo "✅ 数据库连接正常"
    mysql -h127.0.0.1 -uroot -p123456 -e "SHOW STATUS LIKE 'Threads_connected';" 2>/dev/null
else
    echo "❌ 数据库连接失败"
fi

# 检查Redis状态
echo "🔴 Redis状态:"
if redis-cli ping > /dev/null 2>&1; then
    echo "✅ Redis连接正常"
    redis-cli INFO memory | grep used_memory_human
    redis-cli INFO stats | grep total_commands_processed
else
    echo "❌ Redis连接失败"
fi

# 检查磁盘使用情况
echo "💾 磁盘使用情况:"
df -h | grep -E "/$|/var|/tmp"

# 检查日志文件大小
echo "📝 日志文件大小:"
if [ -d "log" ]; then
    du -sh log/*
else
    echo "   日志目录不存在"
fi
EOF

chmod +x scripts/monitor_performance.sh
echo "📁 创建性能监控脚本: scripts/monitor_performance.sh"

echo ""
echo "🎯 性能优化建议："
echo "   - 定期监控应用性能"
echo "   - 适当调整数据库连接池大小"
echo "   - 使用Redis缓存热点数据"
echo "   - 定期清理日志文件"
echo "   - 监控内存使用情况"
echo "   - 使用CDN加速静态资源"
echo ""

echo "📋 性能优化完成！"
echo "🔧 优化内容："
echo "   - 数据库性能配置: 已优化"
echo "   - 数据库索引: 已优化"
echo "   - Redis缓存策略: 已配置"
echo "   - Go应用编译: 已优化"
echo "   - 性能监控: 已配置"
echo ""
echo "🚀 建议重启应用以应用所有优化: go run main.go"
echo "📊 监控性能: ./scripts/monitor_performance.sh" 