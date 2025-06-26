<template>
  <view class="nf-input-wrapper">
    <!-- 输入框标签 -->
    <view v-if="label" class="nf-input__label">
      <text>{{ label }}</text>
      <text v-if="required" class="nf-input__required">*</text>
    </view>
    
    <!-- 输入框容器 -->
    <view :class="inputContainerClasses">
      <!-- 左侧图标 -->
      <view v-if="iconLeft" class="nf-input__icon nf-input__icon--left">
        <text :class="`nf-icon ${iconLeft}`"></text>
      </view>
      
      <!-- 输入框 -->
      <input
        :class="inputClasses"
        :type="type"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :maxlength="maxlength"
        @input="handleInput"
        @focus="handleFocus"
        @blur="handleBlur"
      />
      
      <!-- 右侧图标 -->
      <view v-if="iconRight" class="nf-input__icon nf-input__icon--right">
        <text :class="`nf-icon ${iconRight}`"></text>
      </view>
    </view>
    
    <!-- 错误提示 -->
    <view v-if="error && errorMessage" class="nf-input__error">
      <text>{{ errorMessage }}</text>
    </view>
    
    <!-- 帮助文本 -->
    <view v-if="helpText" class="nf-input__help">
      <text>{{ helpText }}</text>
    </view>
  </view>
</template>

<script>
export default {
  name: 'NfInput',
  props: {
    modelValue: {
      type: [String, Number],
      default: ''
    },
    type: {
      type: String,
      default: 'text',
      validator: (value) => ['text', 'password', 'number', 'email', 'tel'].includes(value)
    },
    label: {
      type: String,
      default: ''
    },
    placeholder: {
      type: String,
      default: ''
    },
    required: {
      type: Boolean,
      default: false
    },
    disabled: {
      type: Boolean,
      default: false
    },
    error: {
      type: Boolean,
      default: false
    },
    errorMessage: {
      type: String,
      default: ''
    },
    helpText: {
      type: String,
      default: ''
    },
    iconLeft: {
      type: String,
      default: ''
    },
    iconRight: {
      type: String,
      default: ''
    },
    size: {
      type: String,
      default: 'medium',
      validator: (value) => ['small', 'medium', 'large'].includes(value)
    },
    maxlength: {
      type: Number,
      default: -1
    }
  },
  emits: ['update:modelValue', 'focus', 'blur', 'input'],
  data() {
    return {
      focused: false
    }
  },
  computed: {
    inputContainerClasses() {
      const classes = ['nf-input__container']
      
      // 添加尺寸样式
      classes.push(`nf-input__container--${this.size}`)
      
      // 添加状态样式
      if (this.focused) {
        classes.push('nf-input__container--focused')
      }
      
      if (this.error) {
        classes.push('nf-input__container--error')
      }
      
      if (this.disabled) {
        classes.push('nf-input__container--disabled')
      }
      
      return classes.join(' ')
    },
    inputClasses() {
      const classes = ['nf-input__field']
      
      // 添加图标间距
      if (this.iconLeft) {
        classes.push('nf-input__field--with-left-icon')
      }
      
      if (this.iconRight) {
        classes.push('nf-input__field--with-right-icon')
      }
      
      return classes.join(' ')
    }
  },
  methods: {
    handleInput(event) {
      this.$emit('update:modelValue', event.target.value)
      this.$emit('input', event)
    },
    handleFocus(event) {
      this.focused = true
      this.$emit('focus', event)
    },
    handleBlur(event) {
      this.focused = false
      this.$emit('blur', event)
    }
  }
}
</script>

<style scoped>
/* 输入框包装器 */
.nf-input-wrapper {
  width: 100%;
}

/* 输入框标签 */
.nf-input__label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
  margin-bottom: 6px;
}

.nf-input__required {
  color: #EF4444;
  margin-left: 2px;
}

/* 输入框容器 */
.nf-input__container {
  position: relative;
  display: flex;
  align-items: center;
  background-color: #FFFFFF;
  border: 2px solid #E5E7EB;
  border-radius: 12px;
  transition: all 0.2s ease;
}

.nf-input__container--small {
  padding: 8px 12px;
}

.nf-input__container--medium {
  padding: 12px 16px;
}

.nf-input__container--large {
  padding: 16px 20px;
}

/* 聚焦状态 */
.nf-input__container--focused {
  border-color: #34D399;
  box-shadow: 0 0 0 3px rgba(52, 211, 153, 0.1);
}

/* 错误状态 */
.nf-input__container--error {
  border-color: #EF4444;
}

.nf-input__container--error.nf-input__container--focused {
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

/* 禁用状态 */
.nf-input__container--disabled {
  background-color: #F9FAFB;
  border-color: #E5E7EB;
  cursor: not-allowed;
}

/* 输入框字段 */
.nf-input__field {
  flex: 1;
  border: none;
  outline: none;
  background: transparent;
  font-size: 16px;
  color: #1F2937;
}

.nf-input__field::placeholder {
  color: #9CA3AF;
}

.nf-input__field:disabled {
  color: #9CA3AF;
  cursor: not-allowed;
}

/* 图标间距 */
.nf-input__field--with-left-icon {
  margin-left: 8px;
}

.nf-input__field--with-right-icon {
  margin-right: 8px;
}

/* 图标样式 */
.nf-input__icon {
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9CA3AF;
  font-size: 16px;
}

.nf-input__icon--left {
  margin-right: 0;
}

.nf-input__icon--right {
  margin-left: 0;
}

.nf-icon {
  font-family: "Font Awesome 6 Free";
  font-weight: 900;
}

/* 错误信息 */
.nf-input__error {
  margin-top: 6px;
  font-size: 12px;
  color: #EF4444;
}

/* 帮助文本 */
.nf-input__help {
  margin-top: 6px;
  font-size: 12px;
  color: #6B7280;
}

/* 悬浮效果 */
.nf-input__container:hover:not(.nf-input__container--disabled) {
  border-color: #34D399;
}
</style> 