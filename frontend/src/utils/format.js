/**
 * 格式化数字显示
 * @param {number} num 
 * @returns {string}
 */
export function formatNumber(num) {
  if (!num) return '0'
  
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  }
  
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'k'
  }
  
  return num.toString()
}

/**
 * 格式化时长（分钟）
 * @param {number} minutes 
 * @returns {string}
 */
export function formatDuration(minutes) {
  if (!minutes) return '0:00'
  
  const hours = Math.floor(minutes / 60)
  const mins = minutes % 60
  
  if (hours > 0) {
    return `${hours}:${mins.toString().padStart(2, '0')}`
  }
  
  return `${mins}:00`
}

/**
 * 格式化时间（秒）
 * @param {number} seconds 
 * @returns {string}
 */
export function formatTime(seconds) {
  if (!seconds) return '0:00'
  
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60)
  
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

/**
 * 格式化日期
 * @param {string|Date} date 
 * @returns {string}
 */
export function formatDate(date) {
  if (!date) return ''
  
  const d = new Date(date)
  return `${d.getMonth() + 1}月${d.getDate()}日`
}

/**
 * 格式化日期时间
 * @param {string|Date} date 
 * @returns {string}
 */
export function formatDateTime(date) {
  if (!date) return ''
  
  const d = new Date(date)
  return `${d.getMonth() + 1}月${d.getDate()}日 ${d.getHours()}:${d.getMinutes().toString().padStart(2, '0')}`
}

/**
 * 获取当前时间字符串
 * @returns {string}
 */
export function getCurrentTime() {
  const now = new Date()
  return `${now.getHours()}:${now.getMinutes().toString().padStart(2, '0')}`
}

/**
 * 获取相对时间
 * @param {string|Date} date 
 * @returns {string}
 */
export function getRelativeTime(date) {
  if (!date) return ''
  
  const now = new Date()
  const target = new Date(date)
  const diff = now - target
  
  const seconds = Math.floor(diff / 1000)
  const minutes = Math.floor(seconds / 60)
  const hours = Math.floor(minutes / 60)
  const days = Math.floor(hours / 24)
  
  if (days > 0) {
    return `${days}天前`
  } else if (hours > 0) {
    return `${hours}小时前`
  } else if (minutes > 0) {
    return `${minutes}分钟前`
  } else {
    return '刚刚'
  }
} 