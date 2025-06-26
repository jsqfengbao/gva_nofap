#!/bin/bash

# API性能测试脚本
# 用于测试API接口的性能指标

echo "🚀 开始API性能测试..."

# 进入服务器目录
cd "$(dirname "$0")/.."

# 检查服务是否运行
if ! curl -s http://localhost:8888/health > /dev/null; then
    echo "⚠️  服务未运行，请先启动服务器"
    echo "💡 运行命令: go run main.go"
    exit 1
fi

echo "✅ 服务器运行正常"

# 运行Go基准测试
echo ""
echo "📊 运行Go基准测试..."
go test ./test/... -bench=. -benchmem -count=3

# 使用wrk进行压力测试（如果安装了wrk）
if command -v wrk &> /dev/null; then
    echo ""
    echo "🔥 运行wrk压力测试..."
    
    # 测试微信登录接口
    echo "测试微信登录接口..."
    wrk -t12 -c400 -d30s -s scripts/wrk_scripts/login_test.lua http://localhost:8888/miniprogram/auth/wx-login
    
    # 测试用户资料接口
    echo "测试用户资料接口..."
    wrk -t12 -c400 -d30s http://localhost:8888/miniprogram/user/profile
    
    # 测试打卡接口
    echo "测试打卡接口..."
    wrk -t12 -c400 -d30s -s scripts/wrk_scripts/checkin_test.lua http://localhost:8888/miniprogram/checkin/daily
    
else
    echo "⚠️  wrk未安装，跳过压力测试"
    echo "💡 安装wrk: brew install wrk (macOS) 或参考官方文档"
fi

# 使用ab进行压力测试（如果安装了ab）
if command -v ab &> /dev/null; then
    echo ""
    echo "🔥 运行Apache Bench压力测试..."
    
    # 测试健康检查接口
    echo "测试健康检查接口..."
    ab -n 10000 -c 100 http://localhost:8888/health
    
else
    echo "⚠️  Apache Bench未安装，跳过ab测试"
fi

echo ""
echo "📋 性能测试完成！"
echo "📈 查看详细报告请关注上述输出"
echo ""
echo "🎯 性能优化建议："
echo "   - 响应时间应 < 100ms (P95)"
echo "   - 吞吐量应 > 1000 QPS"
echo "   - 内存使用应稳定增长"
echo "   - CPU使用率应 < 80%" 