// UI组件库统一导出
import NfButton from './button/NfButton.vue'
import NfCard from './card/NfCard.vue'
import NfInput from './form/NfInput.vue'
import NfTabBar from './navigation/NfTabBar.vue'
import NfNavBar from './navigation/NfNavBar.vue'

// 组件映射
const components = {
  NfButton,
  NfCard,
  NfInput,
  NfTabBar,
  NfNavBar
}

// 批量注册函数
const install = (app) => {
  Object.keys(components).forEach(key => {
    app.component(key, components[key])
  })
}

export {
  // 单独导出组件
  NfButton,
  NfCard,
  NfInput,
  NfTabBar,
  NfNavBar,
  
  // 批量安装函数
  install
}

export default {
  install
} 