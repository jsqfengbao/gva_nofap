// API测试脚本
const http = require('http');

// 测试配置
const API_BASE = 'http://192.168.1.140:8888';
const API_PREFIX = '/api/v1/miniprogram';

// 测试函数
function testAPI(path, method = 'GET', data = null, headers = {}) {
  return new Promise((resolve, reject) => {
    const url = API_BASE + API_PREFIX + path;
    console.log(`🧪 测试: ${method} ${url}`);
    
    const options = {
      hostname: '192.168.1.140',
      port: 8888,
      path: API_PREFIX + path,
      method: method,
      headers: {
        'Content-Type': 'application/json',
        ...headers
      }
    };

    const req = http.request(options, (res) => {
      let responseBody = '';
      
      res.on('data', (chunk) => {
        responseBody += chunk;
      });
      
      res.on('end', () => {
        try {
          const jsonData = JSON.parse(responseBody);
          console.log(`✅ 响应 (${res.statusCode}):`, jsonData);
          resolve({ statusCode: res.statusCode, data: jsonData });
        } catch (error) {
          console.log(`�� 原始响应 (${res.statusCode}):`, responseBody);
          resolve({ statusCode: res.statusCode, raw: responseBody });
        }
      });
    });

    req.on('error', (error) => {
      console.error('❌ 请求失败:', error.message);
      reject(error);
    });

    if (data) {
      req.write(JSON.stringify(data));
    }
    
    req.end();
  });
}

// 主测试函数
async function runTests() {
  console.log('🚀 开始API测试...\n');

  try {
    console.log('\n--- 测试: 用户info接口（需要认证，应该返回401或7） ---');
    // 测试用户info接口（这是报错的接口）
    await testAPI('/user/info');
    
    console.log('\n🎉 测试完成！');
    console.log('\n📋 分析结果:');
    console.log('- 如果看到"未提供认证token"错误(code: 7)，说明JWT中间件正常工作');
    console.log('- 如果看到其他错误，可能存在路由配置问题');
    
  } catch (error) {
    console.error('💥 测试过程中出现错误:', error);
  }
}

// 运行测试
runTests();
