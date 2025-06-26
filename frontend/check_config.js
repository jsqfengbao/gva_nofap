#!/usr/bin/env node

/**
 * 微信小程序配置检查脚本
 * 验证所有配置文件中的AppID是否一致
 */

const fs = require('fs')
const path = require('path')

const EXPECTED_APPID = 'wx07c9e8e4f105260b'

function checkFile(filePath, property = 'appid') {
  try {
    if (!fs.existsSync(filePath)) {
      return { exists: false, error: '文件不存在' }
    }
    
    // JS文件特殊处理
    if (filePath.includes('env.js')) {
      return { exists: true, needsManualCheck: true }
    }
    
    const content = fs.readFileSync(filePath, 'utf8')
    
    // 移除注释的简单方法（仅针对JSON文件中的注释）
    const cleanContent = content.replace(/\/\*[\s\S]*?\*\//g, '').replace(/\/\/.*$/gm, '')
    
    const data = JSON.parse(cleanContent)
    
    // 根据文件类型查找AppID的位置
    let appid
    if (filePath.includes('manifest.json')) {
      appid = data['mp-weixin']?.appid
    } else {
      appid = data[property]
    }
    
    return {
      exists: true,
      appid: appid,
      correct: appid === EXPECTED_APPID,
      content: data
    }
  } catch (error) {
    return { exists: true, error: error.message }
  }
}

function checkConfigFiles() {
  console.log('🔍 检查微信小程序配置文件...\n')
  
  const files = [
    { path: 'project.config.json', name: '主配置文件' },
    { path: 'project.private.config.json', name: '私有配置文件', optional: true },
    { path: 'manifest.json', name: 'uni-app配置文件' },
    { path: 'src/manifest.json', name: '源码配置文件', optional: true },
    { path: 'src/config/env.js', name: '环境配置文件', needsManualCheck: true },
    { path: 'dist/dev/mp-weixin/project.config.json', name: '构建输出配置', optional: true }
  ]
  
  let allCorrect = true
  const issues = []
  
  files.forEach(file => {
    const result = checkFile(file.path)
    
    console.log(`📄 ${file.name} (${file.path})`)
    
    if (!result.exists) {
      if (file.optional) {
        console.log('   ⚠️  文件不存在（可选）')
      } else {
        console.log('   ❌ 文件不存在')
        allCorrect = false
        issues.push(`${file.name}: 文件不存在`)
      }
    } else if (result.error) {
      console.log(`   ❌ 读取错误: ${result.error}`)
      allCorrect = false
      issues.push(`${file.name}: ${result.error}`)
    } else if (result.needsManualCheck) {
      console.log('   🔧 需要手动检查 JS 文件中的 WECHAT_APP_ID')
    } else if (result.correct) {
      console.log(`   ✅ AppID 正确: ${result.appid}`)
    } else {
      console.log(`   ❌ AppID 错误: ${result.appid} (期望: ${EXPECTED_APPID})`)
      allCorrect = false
      issues.push(`${file.name}: AppID不匹配`)
    }
    
    console.log('')
  })
  
  // 检查关键目录
  console.log('📁 检查关键目录结构...')
  const dirs = ['src', 'dist', 'dist/dev', 'dist/dev/mp-weixin']
  dirs.forEach(dir => {
    if (fs.existsSync(dir)) {
      console.log(`   ✅ ${dir}`)
    } else {
      console.log(`   ❌ ${dir} (不存在)`)
      if (dir === 'dist/dev/mp-weixin') {
        issues.push('构建输出目录不存在，请运行 npm run dev:mp-weixin')
      }
    }
  })
  
  console.log('\n' + '='.repeat(50))
  
  if (allCorrect && issues.length === 0) {
    console.log('🎉 所有配置检查通过！')
    console.log('\n📝 下一步操作：')
    console.log('1. 确保运行 npm run dev:mp-weixin')
    console.log('2. 在微信开发者工具中导入 dist/dev/mp-weixin/')
    console.log(`3. AppID 设置为: ${EXPECTED_APPID}`)
  } else {
    console.log('❌ 发现配置问题：')
    issues.forEach(issue => console.log(`   - ${issue}`))
    console.log('\n🔧 建议解决方案：')
    console.log('1. 检查文件内容和AppID配置')
    console.log('2. 重新运行构建命令')
    console.log('3. 参考 WECHAT_LOGIN_DEBUG.md 排查指南')
  }
  
  return allCorrect
}

// 手动检查 env.js 文件
function checkEnvFile() {
  console.log('\n🔧 手动检查环境配置文件...')
  
  const envFile = 'src/config/env.js'
  try {
    const content = fs.readFileSync(envFile, 'utf8')
    const matches = content.match(/WECHAT_APP_ID:\s*['"](.*?)['"],?/g)
    
    if (matches) {
      console.log('📱 找到的 WECHAT_APP_ID 配置：')
      matches.forEach((match, index) => {
        const appid = match.match(/['"](.*?)['"]/) ?.[1]
        if (appid === EXPECTED_APPID) {
          console.log(`   ✅ 配置 ${index + 1}: ${appid}`)
        } else {
          console.log(`   ❌ 配置 ${index + 1}: ${appid} (期望: ${EXPECTED_APPID})`)
        }
      })
    } else {
      console.log('   ⚠️  未找到 WECHAT_APP_ID 配置')
    }
  } catch (error) {
    console.log(`   ❌ 读取失败: ${error.message}`)
  }
}

// 主函数
function main() {
  console.log('微信小程序配置检查工具\n')
  
  // 检查是否在正确的目录
  if (!fs.existsSync('package.json')) {
    console.log('❌ 请在 frontend 目录下运行此脚本')
    process.exit(1)
  }
  
  const success = checkConfigFiles()
  checkEnvFile()
  
  process.exit(success ? 0 : 1)
}

if (require.main === module) {
  main()
} 