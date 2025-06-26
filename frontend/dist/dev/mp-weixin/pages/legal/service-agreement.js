"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  name: "ServiceAgreementPage",
  methods: {
    goBack() {
      common_vendor.index.navigateBack();
    },
    agreeAndContinue() {
      common_vendor.index.setStorageSync("serviceAgreementAccepted", true);
      common_vendor.index.showToast({
        title: "感谢您的同意",
        icon: "success"
      });
      setTimeout(() => {
        common_vendor.index.navigateBack();
      }, 1e3);
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
      title: "服务协议",
      ["show-back"]: true
    }),
    c: common_vendor.o($options.agreeAndContinue),
    d: common_vendor.p({
      type: "primary",
      size: "large",
      label: "我已阅读并同意",
      ["full-width"]: true
    })
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-fd48dad1"]]);
wx.createPage(MiniProgramPage);
