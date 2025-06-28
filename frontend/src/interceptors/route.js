/**
 * 路由拦截器
 * 处理登录拦截和页面访问权限控制
 */
import { isLoggedIn, getUserInfo, wxLogin } from '@/utils/auth.js'

// 登录页面路径
const LOGIN_ROUTE = '/pages/welcome/welcome'

// 需要登录的页面列表（黑名单模式）
const NEED_LOGIN_PAGES = [
  '/pages/profile/auth',
  '/pages/profile/setup', 
  '/pages/profile/privacy',
  '/pages/checkin/index',
  '/pages/checkin/history',
  '/pages/checkin/weekly',
  '/pages/checkin/calendar',
  '/pages/community/index',
  '/pages/community/post',
  '/pages/community/detail',
  '/pages/learning/detail',
  '/pages/progress/index'
]

// 游客模式可访问的页面（白名单模式）
const GUEST_ALLOWED_PAGES = [
  '/pages/index/index',
  '/pages/welcome/welcome',
  '/pages/learning/index',
  '/pages/emergency/index',
  '/pages/about/index',
  '/pages/legal/privacy',
  '/pages/legal/terms'
]

/**
 * 检查是否已登录
 */
function checkIsLoggedIn() {
  return isLoggedIn()
}

/**
 * 初始启动拦截器
 * 应用启动时自动尝试登录
 */
export async function initLaunchInterceptor() {
  console.log('🚀 初始化启动拦截器...')
  
  const hasLogin = checkIsLoggedIn()
  if (!hasLogin) {
    console.log('📱 用户未登录，尝试自动登录...')
    
    try {
      // 尝试自动微信登录（仅获取基础信息）
      const result = await wxLogin()
      if (result) {
        console.log('✅ 自动登录成功')
        return true
      }
    } catch (error) {
      console.log('⚠️ 自动登录失败，将在用户主动操作时再次尝试:', error.message)
    }
  } else {
    console.log('✅ 用户已登录')
  }
  
  return hasLogin
}

/**
 * 导航拦截器逻辑
 */
const navigateInterceptor = {
  invoke({ url }) {
    const path = url.split('?')[0] // 移除查询参数
    
    console.log(`🧭 路由拦截检查: ${path}`)
    
    // 检查是否需要登录
    const needLogin = NEED_LOGIN_PAGES.includes(path)
    
    if (!needLogin) {
      console.log(`✅ 页面 ${path} 无需登录验证`)
      return true
    }
    
    // 检查登录状态
    const hasLogin = checkIsLoggedIn()
    
    if (hasLogin) {
      console.log(`✅ 用户已登录，允许访问 ${path}`)
      return true
    }
    
    // 用户未登录，需要拦截
    console.log(`🔒 用户未登录，拦截访问 ${path}`)
    
    // 显示登录提示
    uni.showModal({
      title: '需要登录',
      content: '此功能需要登录后使用，是否进行微信登录？',
      confirmText: '立即登录',
      cancelText: '取消',
      success: async (res) => {
        if (res.confirm) {
          try {
            uni.showLoading({
              title: '登录中...'
            })
            
            // 执行微信登录
            await wxLogin()
            
            uni.hideLoading()
            uni.showToast({
              title: '登录成功',
              icon: 'success'
            })
            
            // 登录成功后跳转到原目标页面
            setTimeout(() => {
              uni.navigateTo({ url })
            }, 1500)
            
          } catch (error) {
            uni.hideLoading()
            console.error('登录失败:', error)
            
            uni.showToast({
              title: error.message || '登录失败',
              icon: 'none'
            })
            
            // 登录失败，跳转到登录页面
            setTimeout(() => {
              const redirectUrl = encodeURIComponent(url)
              uni.navigateTo({
                url: `${LOGIN_ROUTE}?redirect=${redirectUrl}`
              })
            }, 2000)
          }
        } else {
          // 用户取消登录，跳转到首页
          uni.switchTab({
            url: '/pages/index/index'
          })
        }
      }
    })
    
    // 阻止原始导航
    return false
  }
}

/**
 * Tab切换拦截器
 * 处理底部Tab页面的登录拦截
 */
const tabSwitchInterceptor = {
  invoke({ url }) {
    const path = url.split('?')[0]
    
    console.log(`📱 Tab切换检查: ${path}`)
    
    // Tab页面通常都允许访问，但可能需要特殊处理
    // 比如个人中心页面，未登录时显示登录引导
    if (path === '/pages/profile/index') {
      const hasLogin = checkIsLoggedIn()
      if (!hasLogin) {
        console.log('👤 个人中心页面，用户未登录，显示登录引导')
        // 允许访问，但页面内会显示登录引导
      }
    }
    
    return true
  }
}

/**
 * 页面重定向拦截器
 */
const redirectInterceptor = {
  invoke({ url }) {
    const path = url.split('?')[0]
    
    console.log(`🔄 重定向检查: ${path}`)
    
    // 重定向时的登录检查逻辑
    const needLogin = NEED_LOGIN_PAGES.includes(path)
    
    if (needLogin && !checkIsLoggedIn()) {
      console.log(`🔒 重定向被拦截: ${path}`)
      
      // 重定向到登录页面
      const redirectUrl = encodeURIComponent(url)
      uni.redirectTo({
        url: `${LOGIN_ROUTE}?redirect=${redirectUrl}`
      })
      
      return false
    }
    
    return true
  }
}

/**
 * 页面重启拦截器
 */
const reLaunchInterceptor = {
  invoke({ url }) {
    const path = url.split('?')[0]
    
    console.log(`🚀 重启应用检查: ${path}`)
    
    // reLaunch通常用于登录/登出后的页面跳转，一般不需要拦截
    // 但可以在这里做一些清理工作
    
    return true
  }
}

/**
 * 获取需要登录的页面列表
 * 开发环境下动态获取，生产环境使用静态列表
 */
export function getNeedLoginPages() {
  // 开发环境下可以动态检查pages.json
  if (process.env.NODE_ENV === 'development') {
    try {
      // 这里可以添加动态检查逻辑
      return NEED_LOGIN_PAGES
    } catch (error) {
      console.warn('获取页面配置失败，使用默认配置')
      return NEED_LOGIN_PAGES
    }
  }
  
  return NEED_LOGIN_PAGES
}

/**
 * 检查页面是否允许游客访问
 */
export function isGuestAllowed(path) {
  return GUEST_ALLOWED_PAGES.includes(path)
}

/**
 * 路由拦截器安装器
 */
export const routeInterceptor = {
  async install() {
    console.log('🧭 安装路由拦截器...')
    
    try {
      // 初始化启动拦截器
      await initLaunchInterceptor()
      
      // 安装各种导航拦截器
      uni.addInterceptor('navigateTo', navigateInterceptor)
      uni.addInterceptor('switchTab', tabSwitchInterceptor)
      uni.addInterceptor('redirectTo', redirectInterceptor)
      uni.addInterceptor('reLaunch', reLaunchInterceptor)
      
      console.log('✅ 路由拦截器安装完成')
    } catch (error) {
      console.error('❌ 路由拦截器安装失败:', error)
      throw error
    }
  }
} 