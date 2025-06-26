#!/usr/bin/env node

/**
 * 微信小程序快速启动脚本
 * 检查配置 -> 构建项目 -> 提供导入指南
 */

const { spawn } = require('child_process')
const fs = require('fs')

async function runCommand(command, args = []) {
  return new Promise((resolve, reject) => {
    console.log(`\n🚀 运行: ${command} ${args.join(' ')}`)
    
    const process = spawn(command, args, {
      stdio: 'inherit',
      shell: true
    })
    
    process.on('close', (code) => {
      if (code === 0) {
        resolve()
      } else {
        reject(new Error(`命令执行失败，退出码: ${code}`))
      }
    })
  })
}

async function main() {
  console.log('🎯 微信小程序快速启动工具\n')
  
  try {
    // 步骤1: 检查配置
    console.log('📋 步骤 1: 检查配置文件...')
    await runCommand('node', ['check_config.js'])
    
    // 步骤2: 确保依赖已安装
    console.log('\n📦 步骤 2: 检查依赖...')
    if (!fs.existsSync('node_modules')) {
      console.log('正在安装依赖...')
      await runCommand('npm', ['install'])
    } else {
      console.log('✅ 依赖已安装')
    }
    
    // 步骤3: 构建项目
    console.log('\n🔨 步骤 3: 构建项目...')
    console.log('⚠️  构建将在后台持续运行，请保持此进程')
    
    // 显示构建命令但不等待（因为它是持续运行的）
    console.log('\n运行构建命令: npm run dev:mp-weixin')
    console.log('（请在另一个终端中运行，或者按 Ctrl+C 退出此脚本后手动运行）')
    
    // 步骤4: 显示导入指南
    console.log('\n' + '='.repeat(60))
    console.log('🎉 配置检查完成！现在可以导入微信开发者工具了')
    console.log('='.repeat(60))
    
    console.log('\n📱 微信开发者工具导入步骤：')
    console.log('1. 打开微信开发者工具')
    console.log('2. 选择"导入项目"')
    console.log('3. 项目目录选择: frontend/dist/dev/mp-weixin/')
    console.log('4. AppID 填写: wx07c9e8e4f105260b')
    console.log('5. 项目名称: 戒色助手')
    
    console.log('\n⚙️  开发者工具设置：')
    console.log('- 详情 → 本地设置 → 勾选"不校验合法域名"')
    console.log('- 详情 → 本地设置 → 勾选"启用调试"')
    
    console.log('\n🧪 测试登录功能：')
    console.log('- 在开发者工具Console中运行: wx.login({success: console.log, fail: console.error})')
    console.log('- 或者在小程序中点击"微信授权登录"按钮')
    
    console.log('\n📚 更多帮助：')
    console.log('- 详细导入指南: frontend/WECHAT_IMPORT_GUIDE.md')
    console.log('- 登录问题排查: frontend/WECHAT_LOGIN_DEBUG.md')
    
    console.log('\n🔄 下次快速启动：')
    console.log('cd frontend && node start.js')
    
  } catch (error) {
    console.error('\n❌ 启动失败:', error.message)
    console.log('\n🔧 故障排除：')
    console.log('1. 确保在 frontend 目录下运行')
    console.log('2. 检查 Node.js 和 npm 是否已安装')
    console.log('3. 查看错误信息并参考文档')
    process.exit(1)
  }
}

if (require.main === module) {
  main()
} 