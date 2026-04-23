<template>
  <view class="app">
    <!-- uni-app X 小程序应用根组件 -->
  </view>
</template>

<script setup lang="ts">
import { onLaunch, onShow, onHide } from '@dcloudio/uni-app'

onLaunch(() => {
  console.log('App Launch')
})

onShow(() => {
  console.log('App Show')
  // 应用显示时的逻辑
})

onHide(() => {
  console.log('App Hide')
  // 应用隐藏时的逻辑
})
</script>

<style lang="scss">
/* 引入全局样式 */
@import '@/styles/globals.css';

.app {
  width: 100%;
  height: 100vh;
}

/* 全局样式 */
page {
  background-color: var(--background);
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', system-ui, sans-serif;
  font-size: 16px;
  line-height: 1.5;
  color: var(--text-primary);
}

/* 重置样式 */
/* 注意：微信小程序不支持通用选择器 *，使用具体选择器 */
view, text, button, input, textarea, image, page {
  box-sizing: border-box;
}

/* uni-app 组件样式重置 */
view,
text,
button,
input,
textarea,
image {
  margin: 0;
  padding: 0;
  border: none;
  outline: none;
}

button {
  background: none;
  color: inherit;
  font: inherit;
  cursor: pointer;
}

/* 通用工具类 */
.flex {
  display: flex;
}

.flex-1 {
  flex: 1;
}

.items-center {
  align-items: center;
}

.justify-center {
  justify-content: center;
}

.justify-between {
  justify-content: space-between;
}

.text-center {
  text-align: center;
}

.w-full {
  width: 100%;
}

.h-full {
  height: 100%;
}

.bg-primary {
  background-color: #34d399;
}

.text-primary {
  color: #34d399;
}

.bg-secondary {
  background-color: #06b6d4;
}

.text-secondary {
  color: #06b6d4;
}

.bg-accent {
  background-color: #f59e0b;
}

.text-accent {
  color: #f59e0b;
}

.bg-success {
  background-color: #10b981;
}

.text-success {
  color: #10b981;
}

.bg-danger {
  background-color: #ef4444;
}

.text-danger {
  color: #ef4444;
}

.rounded-lg {
  border-radius: 8px;
}

.rounded-xl {
  border-radius: 12px;
}

.rounded-2xl {
  border-radius: 16px;
}

.rounded-3xl {
  border-radius: 24px;
}

.rounded-full {
  border-radius: 50%;
}

.shadow-sm {
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
}

.shadow-md {
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}

.shadow-lg {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.p-2 {
  padding: 8px;
}

.p-3 {
  padding: 12px;
}

.p-4 {
  padding: 16px;
}

.p-6 {
  padding: 24px;
}

.px-4 {
  padding-left: 16px;
  padding-right: 16px;
}

.px-6 {
  padding-left: 24px;
  padding-right: 24px;
}

.py-2 {
  padding-top: 8px;
  padding-bottom: 8px;
}

.py-3 {
  padding-top: 12px;
  padding-bottom: 12px;
}

.py-4 {
  padding-top: 16px;
  padding-bottom: 16px;
}

.m-2 {
  margin: 8px;
}

.m-4 {
  margin: 16px;
}

.mt-2 {
  margin-top: 8px;
}

.mt-4 {
  margin-top: 16px;
}

.mb-2 {
  margin-bottom: 8px;
}

.mb-4 {
  margin-bottom: 16px;
}

.text-sm {
  font-size: 14px;
}

.text-base {
  font-size: 16px;
}

.text-lg {
  font-size: 18px;
}

.text-xl {
  font-size: 20px;
}

.text-2xl {
  font-size: 24px;
}

.text-3xl {
  font-size: 30px;
}

.font-medium {
  font-weight: 500;
}

.font-semibold {
  font-weight: 600;
}

.font-bold {
  font-weight: 700;
}

.text-gray-600 {
  color: #6b7280;
}

.text-gray-800 {
  color: #1f2937;
}

.text-white {
  color: #ffffff;
}

.bg-white {
  background-color: #ffffff;
}

.bg-gray-100 {
  background-color: #f3f4f6;
}

.bg-gray-200 {
  background-color: #e5e7eb;
}

.border {
  border-width: 1px;
}

.border-gray-200 {
  border-color: #e5e7eb;
}

.border-gray-300 {
  border-color: #d1d5db;
}
</style> 