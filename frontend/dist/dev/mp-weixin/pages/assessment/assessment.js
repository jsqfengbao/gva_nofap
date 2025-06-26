"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  name: "AssessmentPage",
  data() {
    return {
      hasHistory: false
    };
  },
  methods: {
    // 返回
    goBack() {
      common_vendor.index.navigateBack();
    },
    // 开始评估
    startAssessment() {
      common_vendor.index.showModal({
        title: "开始评估",
        content: "评估大约需要5-10分钟时间，请确保您有足够的时间完成。准备好开始了吗？",
        confirmText: "开始评估",
        cancelText: "稍后再说",
        success: (res) => {
          if (res.confirm) {
            common_vendor.index.navigateTo({
              url: "/pages/assessment/questionnaire"
            });
          }
        }
      });
    },
    // 查看历史评估
    viewHistory() {
      common_vendor.index.showToast({
        title: "功能开发中...",
        icon: "none"
      });
    },
    // 检查是否有历史评估记录
    checkHistory() {
      try {
        const history = common_vendor.index.getStorageSync("assessmentHistory") || [];
        this.hasHistory = history.length > 0;
      } catch (error) {
        console.error("检查历史记录失败:", error);
        this.hasHistory = false;
      }
    }
  },
  onLoad() {
    this.checkHistory();
  },
  onShow() {
    this.checkHistory();
  }
};
if (!Array) {
  const _component_NfNavBar = common_vendor.resolveComponent("NfNavBar");
  const _component_NfCard = common_vendor.resolveComponent("NfCard");
  const _component_NfButton = common_vendor.resolveComponent("NfButton");
  (_component_NfNavBar + _component_NfCard + _component_NfButton)();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.o($options.goBack),
    b: common_vendor.p({
      title: "色隐指数评估",
      ["show-back"]: true
    }),
    c: common_vendor.p({
      type: "highlight"
    }),
    d: common_vendor.p({
      title: "评估说明",
      icon: "fa-info-circle"
    }),
    e: common_vendor.p({
      title: "评估内容",
      icon: "fa-list-check"
    }),
    f: common_vendor.o($options.startAssessment),
    g: common_vendor.p({
      type: "primary",
      size: "large",
      label: "开始评估",
      ["full-width"]: true,
      ["icon-left"]: "fa-play"
    }),
    h: $data.hasHistory
  }, $data.hasHistory ? {
    i: common_vendor.o($options.viewHistory),
    j: common_vendor.p({
      type: "secondary",
      size: "medium",
      label: "查看历史评估",
      ["full-width"]: true,
      ["icon-left"]: "fa-history"
    })
  } : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-3c273270"]]);
wx.createPage(MiniProgramPage);
