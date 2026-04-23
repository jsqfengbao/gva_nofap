const ci = require('miniprogram-ci');
const fs = require('fs');
const path = require('path');

// 从 package.json 读取版本号
const pkg = require('./package.json');

// 配置信息
const config = {
  appid: 'wx07c9e8e4f105260b',
  type: 'miniProgram',
  projectPath: path.resolve(__dirname, './dist/build/mp-weixin'),
  privateKeyPath: path.resolve(__dirname, './private-key.key'),
  ignores: ['node_modules/**/*'],
};

// 上传配置 - 优先从命令行读取，否则从 package.json 读取
const uploadConfig = {
  version: process.argv[2] || pkg.version || '1.0.0',
  desc: process.argv[3] || '自律助手小程序 - 自动发布',
  robot: 1, // 使用第1个机器人
};

async function main() {
  if (!fs.existsSync(config.privateKeyPath)) {
    console.error('❌ 未找到私钥文件，请先下载 private.key 放到当前目录');
    console.error('   下载地址：微信公众平台 -> 开发 -> 开发设置 -> 小程序代码上传 -> 下载私钥');
    process.exit(1);
  }

  if (!fs.existsSync(config.projectPath)) {
    console.error('❌ 未找到编译产物，请先编译项目', config.projectPath);
    process.exit(1);
  }

  console.log('🚀 开始上传小程序...');
  console.log('AppID:', config.appid);
  console.log('版本:', uploadConfig.version);
  console.log('描述:', uploadConfig.desc);

  const project = new ci.Project({
    ...config,
  });

  try {
    const result = await ci.upload({
      project,
      ...uploadConfig,
      onProgressInfo: (info) => {
        console.log(`📊 进度: ${info.message || JSON.stringify(info)}`);
      },
    });

    console.log('\n✅ 上传成功！');
    console.log('预览链接:', result.previewUrl || '无');
    console.log('分包信息:', JSON.stringify(result.subpackageInfo || {}, null, 2));
    console.log('\n📝 下一步：');
    console.log('1. 请到微信公众平台后台查看上传版本');
    console.log('2. 提交审核 -> 审核通过后发布上线');
  } catch (error) {
    console.error('\n❌ 上传失败:', error.message);
    process.exit(1);
  }
}

main();
