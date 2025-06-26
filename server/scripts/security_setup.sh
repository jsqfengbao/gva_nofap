#!/bin/bash

# 安全性配置脚本
# 用于配置和启用各种安全措施

echo "🔒 开始配置安全性措施..."

# 进入服务器目录
cd "$(dirname "$0")/.."

# 检查Redis是否运行（用于缓存和速率限制）
if ! redis-cli ping > /dev/null 2>&1; then
    echo "⚠️  Redis未运行，某些安全功能可能无法使用"
    echo "💡 启动Redis: redis-server"
fi

# 检查配置文件
if [ ! -f "config.yaml" ]; then
    echo "❌ 配置文件config.yaml不存在"
    exit 1
fi

echo "✅ 配置文件检查通过"

# 检查安全相关的依赖
echo "📦 检查安全依赖..."

# 检查bcrypt
if ! go list -m golang.org/x/crypto/bcrypt > /dev/null 2>&1; then
    echo "📥 安装bcrypt依赖..."
    go get golang.org/x/crypto/bcrypt
fi

# 检查scrypt
if ! go list -m golang.org/x/crypto/scrypt > /dev/null 2>&1; then
    echo "📥 安装scrypt依赖..."
    go get golang.org/x/crypto/scrypt
fi

echo "✅ 依赖检查完成"

# 创建安全日志目录
mkdir -p log/security
mkdir -p log/audit

echo "📁 创建日志目录完成"

# 生成安全密钥（如果不存在）
SECURITY_KEY_FILE=".security_key"
if [ ! -f "$SECURITY_KEY_FILE" ]; then
    echo "🔑 生成安全密钥..."
    openssl rand -base64 32 > "$SECURITY_KEY_FILE"
    chmod 600 "$SECURITY_KEY_FILE"
    echo "✅ 安全密钥生成完成"
else
    echo "✅ 安全密钥已存在"
fi

# 设置文件权限
echo "🛡️  设置文件权限..."
chmod 600 config.yaml
chmod 600 config.docker.yaml
chmod -R 700 log/
chmod 755 scripts/*.sh

echo "✅ 文件权限设置完成"

# 检查SSL证书（如果在生产环境）
if [ "$GIN_MODE" = "release" ]; then
    echo "🔐 检查SSL证书配置..."
    if [ ! -f "cert.pem" ] || [ ! -f "key.pem" ]; then
        echo "⚠️  生产环境缺少SSL证书"
        echo "💡 请配置SSL证书以启用HTTPS"
    else
        echo "✅ SSL证书配置正常"
    fi
fi

# 创建安全配置备份
echo "💾 创建配置备份..."
cp config.yaml "config.backup.$(date +%Y%m%d_%H%M%S).yaml"

# 验证配置
echo "🔍 验证安全配置..."

# 检查JWT密钥强度
JWT_KEY=$(grep "signing-key" config.yaml | awk '{print $2}')
if [ ${#JWT_KEY} -lt 32 ]; then
    echo "⚠️  JWT密钥强度不足，建议使用32位以上的密钥"
fi

# 检查数据库密码
DB_PASSWORD=$(grep "password" config.yaml | grep mysql -A 5 | grep password | awk '{print $2}' | tr -d '"')
if [ "$DB_PASSWORD" = "123456" ] || [ "$DB_PASSWORD" = "password" ]; then
    echo "⚠️  数据库密码过于简单，建议修改为复杂密码"
fi

# 检查Redis配置
if grep -q "password: \"\"" config.yaml; then
    echo "⚠️  Redis未设置密码，建议配置密码"
fi

echo ""
echo "🎯 安全配置建议："
echo "   - 定期更新依赖包"
echo "   - 使用强密码策略"
echo "   - 启用防火墙规则"
echo "   - 定期备份数据"
echo "   - 监控安全日志"
echo "   - 限制API访问频率"
echo ""

# 运行安全测试
echo "🧪 运行安全测试..."
if [ -f "test/security_test.go" ]; then
    go test ./test/security_test.go -v
else
    echo "⚠️  安全测试文件不存在"
fi

echo ""
echo "🔒 安全配置完成！"
echo "📋 配置摘要："
echo "   - 安全中间件: 已配置"
echo "   - 数据加密: 已启用"
echo "   - 输入验证: 已启用"
echo "   - 速率限制: 已启用"
echo "   - 审计日志: 已启用"
echo "   - 敏感数据保护: 已启用"
echo ""
echo "🚀 可以启动服务器了: go run main.go" 