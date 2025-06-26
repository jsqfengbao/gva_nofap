<template>
  <view 
    :class="buttonClasses"
    @click="handleClick"
  >
    <text v-if="iconLeft" :class="`nf-icon ${iconLeft} mr-2`"></text>
    <text v-if="label || $slots.default" class="button-text">
      <slot>{{ label }}</slot>
    </text>
    <text v-if="iconRight" :class="`nf-icon ${iconRight} ml-2`"></text>
  </view>
</template>

<script>
export default {
  name: 'NfButton',
  props: {
    type: {
      type: String,
      default: 'primary',
      validator: (value) => ['primary', 'secondary', 'icon'].includes(value)
    },
    size: {
      type: String,
      default: 'medium',
      validator: (value) => ['small', 'medium', 'large'].includes(value)
    },
    disabled: {
      type: Boolean,
      default: false
    },
    iconLeft: {
      type: String,
      default: ''
    },
    iconRight: {
      type: String,
      default: ''
    },
    fullWidth: {
      type: Boolean,
      default: false
    },
    label: {
      type: String,
      default: ''
    }
  },
  emits: ['click'],
  computed: {
    buttonClasses() {
      const classes = ['nf-button', 'flex', 'items-center', 'justify-center', 'font-semibold']
      
      // 添加类型样式
      classes.push(`nf-button--${this.type}`)
      
      // 添加尺寸样式
      classes.push(`nf-button--${this.size}`)
      
      // 添加宽度样式
      if (this.fullWidth) {
        classes.push('w-full')
      }
      
      // 添加禁用样式
      if (this.disabled) {
        classes.push('nf-button--disabled')
      }
      
      return classes.join(' ')
    }
  },
  methods: {
    handleClick(event) {
      if (!this.disabled) {
        this.$emit('click', event)
      }
    }
  }
}
</script>

<style scoped>
/* 注意：微信小程序不支持外部CSS导入，FontAwesome图标需要通过其他方式引入 */
/* @import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css'); */

/* 基础按钮样式 */
.nf-button {
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  transition: all 0.2s ease;
  border: none;
  outline: none;
  cursor: pointer;
  active: transform scale(0.95);
}

.nf-button:active {
  transform: scale(0.95);
}

/* Primary按钮样式 */
.nf-button--primary {
  background-color: #34D399;
  color: white;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}

.nf-button--primary:hover {
  background-color: #10B981;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

/* Secondary按钮样式 */
.nf-button--secondary {
  background-color: transparent;
  color: #34D399;
  border: 2px solid #34D399;
}

.nf-button--secondary:hover {
  background-color: rgba(52, 211, 153, 0.05);
}

/* Icon按钮样式 */
.nf-button--icon {
  background-color: #F3F4F6;
  color: #6B7280;
  border-radius: 50%;
}

.nf-button--icon:hover {
  background-color: #E5E7EB;
}

/* 尺寸样式 */
.nf-button--small {
  padding: 8px 16px;
  font-size: 14px;
  border-radius: 12px;
}

.nf-button--medium {
  padding: 12px 24px;
  font-size: 16px;
  border-radius: 16px;
}

.nf-button--large {
  padding: 16px 32px;
  font-size: 18px;
  border-radius: 16px;
}

/* Icon按钮特殊尺寸 */
.nf-button--icon.nf-button--small {
  padding: 8px;
  border-radius: 50%;
}

.nf-button--icon.nf-button--medium {
  padding: 12px;
  border-radius: 50%;
}

.nf-button--icon.nf-button--large {
  padding: 16px;
  border-radius: 50%;
}

/* 禁用状态 */
.nf-button--disabled {
  opacity: 0.5;
  cursor: not-allowed;
  pointer-events: none;
}

/* 全宽样式 */
.w-full {
  width: 100%;
}

/* 图标样式 */
.nf-icon {
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

.mr-2 {
  margin-right: 8px;
}

.ml-2 {
  margin-left: 8px;
}

/* Flexbox utilities */
.flex {
  display: flex;
}

.items-center {
  align-items: center;
}

.justify-center {
  justify-content: center;
}

.font-semibold {
  font-weight: 600;
}
</style> 