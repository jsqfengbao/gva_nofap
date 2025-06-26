"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  name: "AboutPage",
  methods: {
    goBack() {
      common_vendor.index.navigateBack();
    },
    startJourney() {
      common_vendor.index.navigateTo({
        url: "/pages/assessment/assessment"
      });
    }
  }
};
if (!Array) {
  const _component_NfNavBar = common_vendor.resolveComponent("NfNavBar");
  const _component_NfCard = common_vendor.resolveComponent("NfCard");
  const _component_NfButton = common_vendor.resolveComponent("NfButton");
  (_component_NfNavBar + _component_NfCard + _component_NfButton)();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.o($options.goBack),
    b: common_vendor.p({
      title: "了解更多",
      ["show-back"]: true
    }),
    c: common_vendor.p({
      type: "gradient",
      icon: "fa-seedling",
      title: "戒色助手"
    }),
    d: common_vendor.p({
      title: "核心功能",
      icon: "fa-star"
    }),
    e: common_vendor.p({
      title: "设计理念",
      icon: "fa-lightbulb"
    }),
    f: common_vendor.o($options.startJourney),
    g: common_vendor.p({
      type: "primary",
      size: "large",
      label: "开始我的康复之旅",
      ["full-width"]: true,
      ["icon-left"]: "fa-play"
    })
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-b5177f87"]]);
wx.createPage(MiniProgramPage);
