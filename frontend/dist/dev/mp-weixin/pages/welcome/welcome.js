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
const config_index = require("../../config/index.js");
const utils_auth = require("../../utils/auth.js");
const _sfc_main = {
  __name: "welcome",
  setup(__props) {
    const isLoading = common_vendor.ref(false);
    const checkLoginStatus = () => {
      const token = common_vendor.index.getStorageSync("token");
      if (token && token !== "guest_token") {
        common_vendor.index.switchTab({
          url: "/pages/index/index"
        });
      }
    };
    const enableDevelopmentMode = () => {
      {
        console.log("🚀 开发模式：自动进入游客模式");
        setTimeout(() => {
          guestMode();
        }, 2e3);
      }
    };
    const handleWxLogin = () => __async(this, null, function* () {
      if (isLoading.value) return;
      isLoading.value = true;
      try {
        const loginRes = yield new Promise((resolve, reject) => {
          common_vendor.index.login({
            provider: "weixin",
            success: resolve,
            fail: reject
          });
        });
        console.log("获取微信登录凭证成功:", loginRes);
        yield performWxLogin(loginRes.code);
      } catch (error) {
        console.error("微信登录失败:", error);
        let errorTitle = "登录失败，请重试";
        if (error.errCode === 40013) {
          errorTitle = "AppID配置错误，请联系开发者";
        } else if (error.errCode === 40125) {
          errorTitle = "小程序未激活，请联系开发者";
        }
        common_vendor.index.showToast({
          title: errorTitle,
          icon: "none",
          duration: 3e3
        });
        isLoading.value = false;
      }
    });
    const performWxLogin = (code) => __async(this, null, function* () {
      try {
        const requestData = { code };
        console.log("发送登录请求:", requestData);
        const res = yield common_vendor.index.request({
          url: config_index.getApiUrl("/auth/wx-login"),
          method: "POST",
          header: {
            "Content-Type": "application/json"
          },
          data: requestData
        });
        console.log("登录响应:", res.data);
        if (res.data.code === 0) {
          const { token, user } = res.data.data;
          utils_auth.setUserInfo(__spreadProps(__spreadValues({}, user), { token }));
          common_vendor.index.showToast({
            title: "登录成功",
            icon: "success"
          });
          const needsSetup = !user.nickname || !user.avatarUrl;
          setTimeout(() => {
            if (needsSetup) {
              common_vendor.index.navigateTo({
                url: "/pages/profile/setup"
              });
            } else {
              common_vendor.index.switchTab({
                url: "/pages/index/index"
              });
            }
          }, 1500);
        } else {
          common_vendor.index.showToast({
            title: res.data.msg || "登录失败",
            icon: "none"
          });
        }
      } catch (error) {
        console.error("微信登录失败:", error);
        common_vendor.index.showToast({
          title: "网络错误，请重试",
          icon: "none"
        });
      } finally {
        isLoading.value = false;
      }
    });
    const guestMode = () => {
      common_vendor.index.setStorageSync("token", "guest_token");
      common_vendor.index.setStorageSync("userInfo", {
        id: 0,
        nickname: "游客用户",
        isGuest: true
      });
      common_vendor.index.showToast({
        title: "进入游客模式",
        icon: "success"
      });
      setTimeout(() => {
        common_vendor.index.switchTab({
          url: "/pages/index/index"
        });
      }, 1500);
    };
    const showAgreement = (type) => {
      const title = type === "user" ? "用户协议" : "隐私政策";
      common_vendor.index.showModal({
        title,
        content: `这里显示${title}的内容...`,
        showCancel: false,
        confirmText: "我知道了"
      });
    };
    common_vendor.onMounted(() => {
      checkLoginStatus();
      enableDevelopmentMode();
    });
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: !isLoading.value
      }, !isLoading.value ? {} : {}, {
        b: common_vendor.t(isLoading.value ? "登录中..." : "微信快速登录"),
        c: isLoading.value ? 1 : "",
        d: common_vendor.o(handleWxLogin),
        e: isLoading.value,
        f: common_vendor.o(guestMode),
        g: isLoading.value,
        h: common_vendor.o(($event) => showAgreement("user")),
        i: common_vendor.o(($event) => showAgreement("privacy"))
      });
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-520b4d47"]]);
wx.createPage(MiniProgramPage);
