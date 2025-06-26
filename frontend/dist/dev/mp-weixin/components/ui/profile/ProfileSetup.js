"use strict";
const common_vendor = require("../../../common/vendor.js");
const _sfc_main = {
  __name: "ProfileSetup",
  emits: ["save", "skip"],
  setup(__props, { emit: __emit }) {
    const emit = __emit;
    const userInfo = common_vendor.ref({
      avatarUrl: "",
      nickname: ""
    });
    const canSave = common_vendor.computed(() => {
      return userInfo.value.nickname.trim().length > 0;
    });
    const onChooseAvatar = (e) => {
      console.log("选择头像:", e.detail);
      if (e.detail.avatarUrl) {
        userInfo.value.avatarUrl = e.detail.avatarUrl;
      }
    };
    const saveProfile = () => {
      if (!canSave.value) {
        common_vendor.index.showToast({
          title: "请输入昵称",
          icon: "none"
        });
        return;
      }
      emit("save", {
        avatarUrl: userInfo.value.avatarUrl,
        nickname: userInfo.value.nickname.trim()
      });
    };
    const skipSetup = () => {
      emit("skip");
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: userInfo.value.avatarUrl
      }, userInfo.value.avatarUrl ? {
        b: userInfo.value.avatarUrl
      } : {}, {
        c: common_vendor.o(onChooseAvatar),
        d: userInfo.value.nickname,
        e: common_vendor.o(($event) => userInfo.value.nickname = $event.detail.value),
        f: !canSave.value,
        g: common_vendor.o(saveProfile),
        h: common_vendor.o(skipSetup)
      });
    };
  }
};
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-21a4da30"]]);
wx.createComponent(Component);
