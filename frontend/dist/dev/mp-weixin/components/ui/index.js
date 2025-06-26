"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main$4 = {
  name: "NfButton",
  props: {
    type: {
      type: String,
      default: "primary",
      validator: (value) => ["primary", "secondary", "icon"].includes(value)
    },
    size: {
      type: String,
      default: "medium",
      validator: (value) => ["small", "medium", "large"].includes(value)
    },
    disabled: {
      type: Boolean,
      default: false
    },
    iconLeft: {
      type: String,
      default: ""
    },
    iconRight: {
      type: String,
      default: ""
    },
    fullWidth: {
      type: Boolean,
      default: false
    },
    label: {
      type: String,
      default: ""
    }
  },
  emits: ["click"],
  computed: {
    buttonClasses() {
      const classes = ["nf-button", "flex", "items-center", "justify-center", "font-semibold"];
      classes.push(`nf-button--${this.type}`);
      classes.push(`nf-button--${this.size}`);
      if (this.fullWidth) {
        classes.push("w-full");
      }
      if (this.disabled) {
        classes.push("nf-button--disabled");
      }
      return classes.join(" ");
    }
  },
  methods: {
    handleClick(event) {
      if (!this.disabled) {
        this.$emit("click", event);
      }
    }
  }
};
function _sfc_render$4(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: $props.iconLeft
  }, $props.iconLeft ? {
    b: common_vendor.n(`nf-icon ${$props.iconLeft} mr-2`)
  } : {}, {
    c: $props.label || _ctx.$slots.default
  }, $props.label || _ctx.$slots.default ? {
    d: common_vendor.t($props.label)
  } : {}, {
    e: $props.iconRight
  }, $props.iconRight ? {
    f: common_vendor.n(`nf-icon ${$props.iconRight} ml-2`)
  } : {}, {
    g: common_vendor.n($options.buttonClasses),
    h: common_vendor.o((...args) => $options.handleClick && $options.handleClick(...args))
  });
}
const Component$3 = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main$4, [["render", _sfc_render$4], ["__scopeId", "data-v-62d418e1"]]);
const _sfc_main$3 = {
  name: "NfCard",
  props: {
    type: {
      type: String,
      default: "basic",
      validator: (value) => ["basic", "highlight", "gradient"].includes(value)
    },
    title: {
      type: String,
      default: ""
    },
    content: {
      type: String,
      default: ""
    },
    icon: {
      type: String,
      default: ""
    },
    padding: {
      type: String,
      default: "medium",
      validator: (value) => ["small", "medium", "large"].includes(value)
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
      const classes = ["nf-card"];
      classes.push(`nf-card--${this.type}`);
      classes.push(`nf-card--padding-${this.padding}`);
      if (this.shadow) {
        classes.push("nf-card--shadow");
      }
      if (this.border) {
        classes.push("nf-card--border");
      }
      return classes.join(" ");
    }
  }
};
function _sfc_render$3(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: $props.title || _ctx.$slots.header
  }, $props.title || _ctx.$slots.header ? common_vendor.e({
    b: $props.icon
  }, $props.icon ? {
    c: common_vendor.n(`nf-icon ${$props.icon}`)
  } : {}, {
    d: $props.title
  }, $props.title ? {
    e: common_vendor.t($props.title)
  } : {}) : {}, {
    f: $props.content || _ctx.$slots.default
  }, $props.content || _ctx.$slots.default ? common_vendor.e({
    g: $props.content
  }, $props.content ? {
    h: common_vendor.t($props.content)
  } : {}) : {}, {
    i: _ctx.$slots.footer
  }, _ctx.$slots.footer ? {} : {}, {
    j: common_vendor.n($options.cardClasses)
  });
}
const Component$2 = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main$3, [["render", _sfc_render$3], ["__scopeId", "data-v-c0c434ff"]]);
const _sfc_main$2 = {
  name: "NfInput",
  props: {
    modelValue: {
      type: [String, Number],
      default: ""
    },
    type: {
      type: String,
      default: "text",
      validator: (value) => ["text", "password", "number", "email", "tel"].includes(value)
    },
    label: {
      type: String,
      default: ""
    },
    placeholder: {
      type: String,
      default: ""
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
      default: ""
    },
    helpText: {
      type: String,
      default: ""
    },
    iconLeft: {
      type: String,
      default: ""
    },
    iconRight: {
      type: String,
      default: ""
    },
    size: {
      type: String,
      default: "medium",
      validator: (value) => ["small", "medium", "large"].includes(value)
    },
    maxlength: {
      type: Number,
      default: -1
    }
  },
  emits: ["update:modelValue", "focus", "blur", "input"],
  data() {
    return {
      focused: false
    };
  },
  computed: {
    inputContainerClasses() {
      const classes = ["nf-input__container"];
      classes.push(`nf-input__container--${this.size}`);
      if (this.focused) {
        classes.push("nf-input__container--focused");
      }
      if (this.error) {
        classes.push("nf-input__container--error");
      }
      if (this.disabled) {
        classes.push("nf-input__container--disabled");
      }
      return classes.join(" ");
    },
    inputClasses() {
      const classes = ["nf-input__field"];
      if (this.iconLeft) {
        classes.push("nf-input__field--with-left-icon");
      }
      if (this.iconRight) {
        classes.push("nf-input__field--with-right-icon");
      }
      return classes.join(" ");
    }
  },
  methods: {
    handleInput(event) {
      this.$emit("update:modelValue", event.target.value);
      this.$emit("input", event);
    },
    handleFocus(event) {
      this.focused = true;
      this.$emit("focus", event);
    },
    handleBlur(event) {
      this.focused = false;
      this.$emit("blur", event);
    }
  }
};
function _sfc_render$2(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: $props.label
  }, $props.label ? common_vendor.e({
    b: common_vendor.t($props.label),
    c: $props.required
  }, $props.required ? {} : {}) : {}, {
    d: $props.iconLeft
  }, $props.iconLeft ? {
    e: common_vendor.n(`nf-icon ${$props.iconLeft}`)
  } : {}, {
    f: common_vendor.n($options.inputClasses),
    g: $props.type,
    h: $props.modelValue,
    i: $props.placeholder,
    j: $props.disabled,
    k: $props.maxlength,
    l: common_vendor.o((...args) => $options.handleInput && $options.handleInput(...args)),
    m: common_vendor.o((...args) => $options.handleFocus && $options.handleFocus(...args)),
    n: common_vendor.o((...args) => $options.handleBlur && $options.handleBlur(...args)),
    o: $props.iconRight
  }, $props.iconRight ? {
    p: common_vendor.n(`nf-icon ${$props.iconRight}`)
  } : {}, {
    q: common_vendor.n($options.inputContainerClasses),
    r: $props.error && $props.errorMessage
  }, $props.error && $props.errorMessage ? {
    s: common_vendor.t($props.errorMessage)
  } : {}, {
    t: $props.helpText
  }, $props.helpText ? {
    v: common_vendor.t($props.helpText)
  } : {});
}
const NfInput = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main$2, [["render", _sfc_render$2], ["__scopeId", "data-v-205f98aa"]]);
const _sfc_main$1 = {
  name: "NfTabBar",
  props: {
    current: {
      type: String,
      default: "home"
    }
  },
  data() {
    return {
      tabs: [
        {
          icon: "🏠",
          text: "主页",
          name: "home",
          pagePath: "/pages/index/index"
        },
        {
          icon: "📊",
          text: "进度",
          name: "progress",
          pagePath: "/pages/progress/index"
        },
        {
          icon: "👥",
          text: "社区",
          name: "community",
          pagePath: "/pages/community/index"
        },
        {
          icon: "📚",
          text: "学习",
          name: "learning",
          pagePath: "/pages/learning/index"
        },
        {
          icon: "👤",
          text: "我的",
          name: "profile",
          pagePath: "/pages/profile/index"
        }
      ]
    };
  },
  computed: {
    activeIndex() {
      return this.tabs.findIndex((tab) => tab.name === this.current);
    }
  },
  methods: {
    getTabClasses(index) {
      const classes = ["nf-tabbar__item"];
      if (index === this.activeIndex) {
        classes.push("nf-tabbar__item--active");
      }
      return classes.join(" ");
    },
    handleTabClick(index) {
      if (index !== this.activeIndex) {
        const tab = this.tabs[index];
        if (tab.pagePath) {
          common_vendor.index.switchTab({
            url: tab.pagePath
          });
        }
      }
    }
  }
};
function _sfc_render$1(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.f($data.tabs, (item, index, i0) => {
      return {
        a: common_vendor.t(item.icon),
        b: common_vendor.t(item.text),
        c: index,
        d: common_vendor.n($options.getTabClasses(index)),
        e: common_vendor.o(($event) => $options.handleTabClick(index), index)
      };
    })
  };
}
const Component$1 = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main$1, [["render", _sfc_render$1], ["__scopeId", "data-v-d6f9864e"]]);
const _sfc_main = {
  name: "NfNavBar",
  props: {
    title: {
      type: String,
      default: ""
    },
    showBack: {
      type: Boolean,
      default: false
    },
    rightIcon: {
      type: String,
      default: ""
    },
    rightText: {
      type: String,
      default: ""
    },
    backgroundColor: {
      type: String,
      default: "#FFFFFF"
    },
    textColor: {
      type: String,
      default: "#1F2937"
    },
    statusBarHeight: {
      type: Number,
      default: 0
    }
  },
  emits: ["back", "rightClick", "statusBarHeight"],
  mounted() {
    if (this.statusBarHeight === 0) {
      this.getSystemInfo();
    }
  },
  methods: {
    getSystemInfo() {
      common_vendor.index.getSystemInfo({
        success: (res) => {
          this.$emit("statusBarHeight", res.statusBarHeight || 20);
        }
      });
    },
    handleBack() {
      this.$emit("back");
      const pages = getCurrentPages();
      if (pages.length > 1) {
        common_vendor.index.navigateBack();
      } else {
        common_vendor.index.switchTab({
          url: "/pages/index/index"
        });
      }
    },
    handleRightClick() {
      this.$emit("rightClick");
    }
  }
};
const __injectCSSVars__ = () => {
  common_vendor.useCssVars((_ctx) => ({
    "6652b3ff": _ctx.backgroundColor,
    "40a4da40": _ctx.textColor
  }));
};
const __setup__ = _sfc_main.setup;
_sfc_main.setup = __setup__ ? (props, ctx) => {
  __injectCSSVars__();
  return __setup__(props, ctx);
} : __injectCSSVars__;
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: $props.statusBarHeight + "px",
    b: $props.showBack
  }, $props.showBack ? {
    c: common_vendor.o((...args) => $options.handleBack && $options.handleBack(...args))
  } : {}, {
    d: common_vendor.t($props.title),
    e: $props.rightIcon
  }, $props.rightIcon ? {
    f: common_vendor.n(`nf-icon ${$props.rightIcon}`),
    g: common_vendor.o((...args) => $options.handleRightClick && $options.handleRightClick(...args))
  } : $props.rightText ? {
    i: common_vendor.t($props.rightText),
    j: common_vendor.o((...args) => $options.handleRightClick && $options.handleRightClick(...args))
  } : {}, {
    h: $props.rightText,
    k: common_vendor.s(_ctx.__cssVars())
  });
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-71c93a46"]]);
const components = {
  NfButton: Component$3,
  NfCard: Component$2,
  NfInput,
  NfTabBar: Component$1,
  NfNavBar: Component
};
const install = (app) => {
  Object.keys(components).forEach((key) => {
    app.component(key, components[key]);
  });
};
const UIComponents = {
  install
};
exports.Component = Component;
exports.Component$1 = Component$2;
exports.Component$2 = Component$3;
exports.Component$3 = Component$1;
exports.UIComponents = UIComponents;
