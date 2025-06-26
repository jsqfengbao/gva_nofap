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
const config_index = require("../../config/index.js");
if (!Math) {
  ProfileSetup();
}
const ProfileSetup = () => "../../components/ui/profile/ProfileSetup.js";
const _sfc_main = {
  __name: "setup",
  setup(__props) {
    const handleSave = (profileData) => __async(this, null, function* () {
      try {
        common_vendor.index.showLoading({
          title: "保存中..."
        });
        let avatarUrl = profileData.avatarUrl;
        if (avatarUrl && avatarUrl.startsWith("http://tmp/")) {
          avatarUrl = yield uploadTempAvatar(avatarUrl);
        }
        yield utils_api.userApi.updateUserInfo({
          nickname: profileData.nickname,
          avatarUrl
        });
        const currentUser = utils_auth.getUserInfo();
        const updatedUser = __spreadProps(__spreadValues({}, currentUser), {
          nickname: profileData.nickname,
          avatarUrl
        });
        utils_auth.setUserInfo(updatedUser);
        common_vendor.index.hideLoading();
        common_vendor.index.showToast({
          title: "保存成功",
          icon: "success"
        });
        setTimeout(() => {
          common_vendor.index.switchTab({
            url: "/pages/profile/index"
          });
        }, 1500);
      } catch (error) {
        console.error("保存用户资料失败:", error);
        common_vendor.index.hideLoading();
        common_vendor.index.showToast({
          title: "保存失败",
          icon: "none"
        });
      }
    });
    const uploadTempAvatar = (tempPath) => __async(this, null, function* () {
      try {
        if (tempPath.includes("wxfile://")) {
          const response = yield utils_api.userApi.saveWxAvatar({
            tempUrl: tempPath
          });
          return response.data.url;
        } else {
          return new Promise((resolve, reject) => {
            common_vendor.index.uploadFile({
              url: config_index.getApiUrl("/user/upload-avatar"),
              filePath: tempPath,
              name: "file",
              header: {
                "Authorization": `Bearer ${utils_auth.getToken()}`
              },
              success: (res) => {
                const data = JSON.parse(res.data);
                if (data.code === 0) {
                  resolve(data.data.url);
                } else {
                  reject(new Error(data.msg));
                }
              },
              fail: reject
            });
          });
        }
      } catch (error) {
        throw error;
      }
    });
    const handleSkip = () => {
      common_vendor.index.switchTab({
        url: "/pages/profile/index"
      });
    };
    return (_ctx, _cache) => {
      return {
        a: common_vendor.o(handleSave),
        b: common_vendor.o(handleSkip)
      };
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-15137599"]]);
wx.createPage(MiniProgramPage);
