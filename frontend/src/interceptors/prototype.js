/**
 * 原型拦截器
 * 解决低版本手机兼容性问题
 */

/**
 * 原型拦截器安装器
 */
export const prototypeInterceptor = {
  install() {
    console.log('🔧 安装原型拦截器...')
    
    // 解决低版本手机不识别 array.at() 导致运行报错的问题
    if (typeof Array.prototype.at !== 'function') {
      console.log('⚠️ 检测到Array.prototype.at不存在，添加polyfill')
      // eslint-disable-next-line no-extend-native
      Array.prototype.at = function (index) {
        if (index < 0) return this[this.length + index]
        if (index >= this.length) return undefined
        return this[index]
      }
    }

    // 解决低版本手机不识别 String.prototype.replaceAll() 的问题
    if (typeof String.prototype.replaceAll !== 'function') {
      console.log('⚠️ 检测到String.prototype.replaceAll不存在，添加polyfill')
      // eslint-disable-next-line no-extend-native
      String.prototype.replaceAll = function (search, replace) {
        return this.split(search).join(replace)
      }
    }

    // 解决低版本手机不识别 Object.fromEntries() 的问题
    if (typeof Object.fromEntries !== 'function') {
      console.log('⚠️ 检测到Object.fromEntries不存在，添加polyfill')
      Object.fromEntries = function (entries) {
        const obj = {}
        for (const [key, value] of entries) {
          obj[key] = value
        }
        return obj
      }
    }

    // 解决低版本手机不识别 Promise.allSettled() 的问题
    if (typeof Promise.allSettled !== 'function') {
      console.log('⚠️ 检测到Promise.allSettled不存在，添加polyfill')
      Promise.allSettled = function (promises) {
        return Promise.all(
          promises.map(promise =>
            Promise.resolve(promise)
              .then(value => ({ status: 'fulfilled', value }))
              .catch(reason => ({ status: 'rejected', reason }))
          )
        )
      }
    }

    console.log('✅ 原型拦截器安装完成')
  }
} 