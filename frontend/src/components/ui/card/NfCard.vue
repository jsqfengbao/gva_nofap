<template>
  <view :class="cardClasses">
    <!-- 卡片头部 -->
    <view v-if="title || $slots.header" class="nf-card__header">
      <slot name="header">
        <view v-if="icon" class="nf-card__icon">
          <text :class="`nf-icon ${icon}`"></text>
        </view>
        <text v-if="title" class="nf-card__title">{{ title }}</text>
      </slot>
    </view>
    
    <!-- 卡片内容 -->
    <view v-if="content || $slots.default" class="nf-card__content">
      <slot>
        <text v-if="content">{{ content }}</text>
      </slot>
    </view>
    
    <!-- 卡片底部 -->
    <view v-if="$slots.footer" class="nf-card__footer">
      <slot name="footer"></slot>
    </view>
  </view>
</template>

<script>
export default {
  name: 'NfCard',
  props: {
    type: {
      type: String,
      default: 'basic',
      validator: (value) => ['basic', 'highlight', 'gradient'].includes(value)
    },
    title: {
      type: String,
      default: ''
    },
    content: {
      type: String,
      default: ''
    },
    icon: {
      type: String,
      default: ''
    },
    padding: {
      type: String,
      default: 'medium',
      validator: (value) => ['small', 'medium', 'large'].includes(value)
    },
    shadow: {
      type: Boolean,
      default: true
    },
    border: {
      type: Boolean,
      default: true
    }
  },
  computed: {
    cardClasses() {
      const classes = ['nf-card']
      
      // 添加类型样式
      classes.push(`nf-card--${this.type}`)
      
      // 添加内边距样式
      classes.push(`nf-card--padding-${this.padding}`)
      
      // 添加阴影
      if (this.shadow) {
        classes.push('nf-card--shadow')
      }
      
      // 添加边框
      if (this.border) {
        classes.push('nf-card--border')
      }
      
      return classes.join(' ')
    }
  }
}
</script>

<style scoped>
/* 引入FontAwesome图标 */
/* 注意：微信小程序不支持外部CSS导入，FontAwesome图标需要通过其他方式引入 */
/* @import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css'); */

/* 基础卡片样式 */
.nf-card {
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.2s ease;
}

/* 基础卡片类型 */
.nf-card--basic {
  background-color: #FFFFFF;
  color: #1F2937;
}

/* 高亮卡片类型 */
.nf-card--highlight {
  background-color: #FFFFFF;
  color: #1F2937;
  border: 2px solid #34D399;
}

/* 渐变卡片类型 */
.nf-card--gradient {
  background: linear-gradient(135deg, #34D399 0%, #10B981 100%);
  color: white;
}

/* 内边距样式 */
.nf-card--padding-small {
  padding: 12px;
}

.nf-card--padding-medium {
  padding: 16px;
}

.nf-card--padding-large {
  padding: 24px;
}

/* 阴影样式 */
.nf-card--shadow {
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
}

/* 边框样式 */
.nf-card--border {
  border: 1px solid #E5E7EB;
}

/* 卡片头部 */
.nf-card__header {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.nf-card__icon {
  margin-right: 8px;
  font-size: 18px;
}

.nf-card__title {
  font-size: 18px;
  font-weight: 600;
  line-height: 1.4;
}

/* 卡片内容 */
.nf-card__content {
  font-size: 14px;
  line-height: 1.5;
  color: inherit;
}

/* 卡片底部 */
.nf-card__footer {
  margin-top: 16px;
  padding-top: 12px;
  border-top: 1px solid rgba(0, 0, 0, 0.1);
}

/* 渐变卡片的特殊样式 */
.nf-card--gradient .nf-card__footer {
  border-top-color: rgba(255, 255, 255, 0.2);
}

/* 图标样式 */
.nf-icon {
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

/* 悬浮效果 */
.nf-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.nf-card--gradient:hover {
  box-shadow: 0 10px 15px -3px rgba(52, 211, 153, 0.3), 0 4px 6px -2px rgba(52, 211, 153, 0.1);
}
</style> 