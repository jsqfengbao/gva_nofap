"use strict";
var __defProp = Object.defineProperty;
var __defProps = Object.defineProperties;
var __getOwnPropDescs = Object.getOwnPropertyDescriptors;
var __getOwnPropSymbols = Object.getOwnPropertySymbols;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __propIsEnum = Object.prototype.propertyIsEnumerable;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp.call(b, prop))
      __defNormalProp(a, prop, b[prop]);
  if (__getOwnPropSymbols)
    for (var prop of __getOwnPropSymbols(b)) {
      if (__propIsEnum.call(b, prop))
        __defNormalProp(a, prop, b[prop]);
    }
  return a;
};
var __spreadProps = (a, b) => __defProps(a, __getOwnPropDescs(b));
var __async = (__this, __arguments, generator) => {
  return new Promise((resolve, reject) => {
    var fulfilled = (value) => {
      try {
        step(generator.next(value));
      } catch (e) {
        reject(e);
      }
    };
    var rejected = (value) => {
      try {
        step(generator.throw(value));
      } catch (e) {
        reject(e);
      }
    };
    var step = (x) => x.done ? resolve(x.value) : Promise.resolve(x.value).then(fulfilled, rejected);
    step((generator = generator.apply(__this, __arguments)).next());
  });
};
const common_vendor = require("../../common/vendor.js");
const utils_api = require("../../utils/api.js");
const utils_auth = require("../../utils/auth.js");
const _sfc_main = {
  __name: "auth",
  setup(__props) {
    const userInfo = common_vendor.ref({
      avatarUrl: "",
      nickname: ""
    });
    const canSave = common_vendor.computed(() => {
      return userInfo.value.nickname.trim().length > 0;
    });
    common_vendor.onMounted(() => {
      const existingUser = utils_auth.getUserInfo();
      if (existingUser) {
        userInfo.value.nickname = existingUser.nickname || "";
        userInfo.value.avatarUrl = existingUser.avatarUrl || "";
      }
    });
    const onChooseAvatar = (e) => {
      console.log("用户选择头像:", e.detail);
      if (e.detail.avatarUrl) {
        userInfo.value.avatarUrl = e.detail.avatarUrl;
        console.log("头像URL设置为:", userInfo.value.avatarUrl);
        common_vendor.index.showToast({
          title: "头像选择成功",
          icon: "success",
          duration: 1500
        });
      } else {
        console.error("头像选择失败，没有获取到avatarUrl");
        common_vendor.index.showToast({
          title: "头像选择失败",
          icon: "error"
        });
      }
    };
    const onNicknameInput = (e) => {
      userInfo.value.nickname = e.detail.value;
      console.log("昵称更新为:", userInfo.value.nickname);
    };
    const onNicknameBlur = () => {
      console.log("昵称失去焦点:", userInfo.value.nickname);
    };
    const saveUserInfo = () => __async(this, null, function* () {
      if (!canSave.value) {
        common_vendor.index.showToast({
          title: "请输入昵称",
          icon: "none"
        });
        return;
      }
      try {
        common_vendor.index.showLoading({
          title: "保存中..."
        });
        console.log("准备保存用户信息:", userInfo.value);
        let finalAvatarUrl = userInfo.value.avatarUrl;
        if (finalAvatarUrl && finalAvatarUrl.includes("wxfile://")) {
          console.log("检测到微信临时头像，准备上传:", finalAvatarUrl);
          try {
            const uploadResponse = yield utils_api.userApi.saveWxAvatar({
              tempUrl: finalAvatarUrl
            });
            if (uploadResponse.data.code === 0) {
              finalAvatarUrl = uploadResponse.data.data.url;
              console.log("头像上传成功，新URL:", finalAvatarUrl);
            } else {
              console.error("头像上传失败:", uploadResponse.data.msg);
            }
          } catch (uploadError) {
            console.error("头像上传异常:", uploadError);
          }
        }
        console.log("调用API更新用户信息:", {
          nickname: userInfo.value.nickname.trim(),
          avatarUrl: finalAvatarUrl
        });
        const updateResponse = yield utils_api.userApi.updateUserInfo({
          nickname: userInfo.value.nickname.trim(),
          avatarUrl: finalAvatarUrl
        });
        if (updateResponse.data.code !== 0) {
          throw new Error(updateResponse.data.msg || "更新失败");
        }
        const currentUser = utils_auth.getUserInfo();
        const updatedUser = __spreadProps(__spreadValues({}, currentUser), {
          nickname: userInfo.value.nickname.trim(),
          avatarUrl: finalAvatarUrl
        });
        utils_auth.setUserInfo(updatedUser);
        console.log("本地用户信息更新完成:", updatedUser);
        common_vendor.index.hideLoading();
        common_vendor.index.showToast({
          title: "保存成功",
          icon: "success"
        });
        setTimeout(() => {
          common_vendor.index.navigateBack();
        }, 1500);
      } catch (error) {
        console.error("保存用户信息失败:", error);
        common_vendor.index.hideLoading();
        common_vendor.index.showToast({
          title: error.message || "保存失败，请重试",
          icon: "none"
        });
      }
    });
    const skipAuth = () => {
      common_vendor.index.showModal({
        title: "确认跳过",
        content: "跳过后可在个人中心随时完善信息",
        success: (res) => {
          if (res.confirm) {
            common_vendor.index.switchTab({
              url: "/pages/profile/index"
            });
          }
        }
      });
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: userInfo.value.avatarUrl
      }, userInfo.value.avatarUrl ? {
        b: userInfo.value.avatarUrl
      } : {}, {
        c: common_vendor.o(onChooseAvatar),
        d: common_vendor.o([($event) => userInfo.value.nickname = $event.detail.value, onNicknameInput]),
        e: common_vendor.o(onNicknameBlur),
        f: userInfo.value.nickname
      }, {
        g: common_vendor.t(userInfo.value.avatarUrl || "未设置"),
        h: common_vendor.t(userInfo.value.nickname || "未设置"),
        i: common_vendor.t(canSave.value ? "是" : "否")
      }, {
        j: !canSave.value,
        k: common_vendor.o(saveUserInfo),
        l: common_vendor.o(skipAuth)
      });
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-8a01fc08"]]);
wx.createPage(MiniProgramPage);
