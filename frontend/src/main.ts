import { createSSRApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
// 导入UI组件库
import UIComponents from './components/ui/index.js'
// 导入拦截器系统
import { installInterceptors } from './interceptors/index.js'

export function createApp() {
  const app = createSSRApp(App)
  const pinia = createPinia()
  
  app.use(pinia)
  // 注册UI组件库
  app.use(UIComponents)
  
  // 安装拦截器系统
  installInterceptors().catch(error => {
    console.error('拦截器安装失败:', error)
  })
  
  return {
    app
  }
} 