#!/bin/bash

# API文档生成脚本
# 用于生成Swagger API文档

echo "开始生成API文档..."

# 检查swag是否安装
if ! command -v swag &> /dev/null; then
    echo "swag工具未安装，正在安装..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# 进入服务器目录
cd "$(dirname "$0")/.."

# 生成Swagger文档
echo "正在生成Swagger文档..."
swag init -g main.go -o docs

if [ $? -eq 0 ]; then
    echo "✅ API文档生成成功！"
    echo "📖 文档位置: server/docs/"
    echo "🌐 启动服务后可访问: http://localhost:8888/swagger/index.html"
else
    echo "❌ API文档生成失败！"
    exit 1
fi

# 运行API测试
echo ""
echo "开始运行API测试..."
go test ./test/... -v

if [ $? -eq 0 ]; then
    echo "✅ API测试通过！"
else
    echo "⚠️  部分API测试失败，请检查相关接口"
fi

echo ""
echo "📋 文档生成完成！"
echo "📁 生成的文件："
echo "   - docs/swagger.json"
echo "   - docs/swagger.yaml" 
echo "   - docs/docs.go" 