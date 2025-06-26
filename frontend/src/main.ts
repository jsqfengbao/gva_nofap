import { createSSRApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
// 导入UI组件库
import UIComponents from './components/ui/index.js'

export function createApp() {
  const app = createSSRApp(App)
  const pinia = createPinia()
  
  app.use(pinia)
  // 注册UI组件库
  app.use(UIComponents)
  
  return {
    app
  }
} 