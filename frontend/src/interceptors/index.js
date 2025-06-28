/**
 * 统一拦截器入口
 * 整合请求拦截器、路由拦截器和原型拦截器
 */
export { requestInterceptor } from './request.js'
export { routeInterceptor } from './route.js'
export { prototypeInterceptor } from './prototype.js'

/**
 * 安装所有拦截器
 */
export async function installInterceptors() {
  console.log('🚀 开始安装拦截器...')
  
  try {
    // 1. 安装原型拦截器（解决兼容性问题）
    prototypeInterceptor.install()
    console.log('✅ 原型拦截器安装完成')
    
    // 2. 安装请求拦截器（处理API请求）
    requestInterceptor.install()
    console.log('✅ 请求拦截器安装完成')
    
    // 3. 安装路由拦截器（处理登录拦截）
    await routeInterceptor.install()
    console.log('✅ 路由拦截器安装完成')
    
    console.log('🎉 所有拦截器安装完成')
  } catch (error) {
    console.error('❌ 拦截器安装失败:', error)
    throw error
  }
} 