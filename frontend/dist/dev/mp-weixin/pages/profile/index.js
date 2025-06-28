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
const utils_auth = require("../../utils/auth.js");
const apis_user = require("../../apis/user.js");
const apis_achievement = require("../../apis/achievement.js");
require("../../data/learning-mock.js");
const _sfc_main = {
  __name: "index",
  setup(__props) {
    const currentTime = common_vendor.ref("9:41");
    const showExportModal = common_vendor.ref(false);
    const showLoading = common_vendor.ref(false);
    const loadingText = common_vendor.ref("加载中...");
    const exportFormat = common_vendor.ref("json");
    const exportFormatIndex = common_vendor.ref(0);
    const exportDataTypes = common_vendor.ref(["profile", "checkin"]);
    const dateRangeIndex = common_vendor.ref(0);
    const loggingIn = common_vendor.ref(false);
    const isUserLoggedIn = common_vendor.ref(false);
    const userProfile = common_vendor.ref({
      nickname: "",
      avatarUrl: "",
      level: 1,
      levelTitle: "新手",
      joinDays: 1,
      currentStreak: 0,
      longestStreak: 0,
      experience: 0,
      achievementCount: 0,
      helpCount: 0
    });
    const recentAchievements = common_vendor.ref([]);
    const dateRangeOptions = common_vendor.ref(["全部时间", "最近一个月", "最近三个月", "最近一年"]);
    const currentTimeDisplay = common_vendor.computed(() => {
      const now = /* @__PURE__ */ new Date();
      return `${now.getHours()}:${now.getMinutes().toString().padStart(2, "0")}`;
    });
    const isLoggedIn = common_vendor.computed(() => isUserLoggedIn.value);
    common_vendor.onMounted(() => {
      updateTime();
      checkLoginStatus();
      if (isUserLoggedIn.value) {
        loadUserProfile();
      }
    });
    const updateTime = () => {
      currentTime.value = currentTimeDisplay.value;
      setInterval(() => {
        currentTime.value = currentTimeDisplay.value;
      }, 6e4);
    };
    const checkLoginStatus = () => {
      isUserLoggedIn.value = utils_auth.isLoggedIn();
      console.log("检查登录状态:", isUserLoggedIn.value);
      if (isUserLoggedIn.value) {
        const userInfo = utils_auth.getUserInfo();
        console.log("本地存储的用户信息:", userInfo);
        if (userInfo) {
          userProfile.value.nickname = userInfo.nickname || "微信用户";
          userProfile.value.avatarUrl = utils_auth.getAvatarUrl(userInfo.avatarUrl);
          console.log("设置用户资料:", {
            nickname: userProfile.value.nickname,
            avatarUrl: userProfile.value.avatarUrl
          });
        }
      }
    };
    const loadUserProfile = () => __async(this, null, function* () {
      var _a, _b, _c, _d, _e, _f, _g, _h, _i, _j, _k;
      if (!isUserLoggedIn.value) return;
      try {
        showLoading.value = true;
        loadingText.value = "加载个人资料...";
        console.log("开始加载用户资料...");
        const res = yield apis_user.userApi.getProfile();
        console.log("用户资料API响应:", res.data);
        if (res.data.code === 0) {
          const data = res.data.data;
          console.log("解析的用户数据:", data);
          const updatedProfile = {
            nickname: ((_a = data.user) == null ? void 0 : _a.nickname) || userProfile.value.nickname || "微信用户",
            avatarUrl: utils_auth.getAvatarUrl((_b = data.user) == null ? void 0 : _b.avatarUrl) || userProfile.value.avatarUrl,
            level: ((_c = data.user) == null ? void 0 : _c.level) || 1,
            levelTitle: ((_d = data.user) == null ? void 0 : _d.levelTitle) || "新手",
            joinDays: ((_e = data.abstinenceRecord) == null ? void 0 : _e.totalDays) || 1,
            currentStreak: ((_f = data.abstinenceRecord) == null ? void 0 : _f.currentStreak) || 0,
            longestStreak: ((_g = data.abstinenceRecord) == null ? void 0 : _g.longestStreak) || 0,
            experience: ((_h = data.user) == null ? void 0 : _h.experience) || 0,
            achievementCount: ((_i = data.user) == null ? void 0 : _i.achievementCount) || 0,
            helpCount: ((_j = data.user) == null ? void 0 : _j.helpCount) || 0
          };
          console.log("更新后的用户资料:", updatedProfile);
          userProfile.value = updatedProfile;
          const currentUser = utils_auth.getUserInfo();
          if (currentUser) {
            const updatedUser = __spreadProps(__spreadValues({}, currentUser), {
              nickname: updatedProfile.nickname,
              avatarUrl: ((_k = data.user) == null ? void 0 : _k.avatarUrl) || currentUser.avatarUrl,
              level: updatedProfile.level,
              experience: updatedProfile.experience
            });
            utils_auth.setUserInfo(updatedUser);
            console.log("更新本地用户信息:", updatedUser);
          }
          yield loadRecentAchievements();
        } else {
          console.error("API返回错误:", res.data.msg);
          common_vendor.index.showToast({
            title: res.data.msg || "加载失败",
            icon: "error"
          });
        }
        showLoading.value = false;
      } catch (error) {
        console.error("加载用户资料失败:", error);
        showLoading.value = false;
        common_vendor.index.showToast({
          title: "加载失败",
          icon: "error"
        });
      }
    });
    const loadRecentAchievements = () => __async(this, null, function* () {
      try {
        const res = yield apis_achievement.achievementApi.getUserAchievements({ limit: 3, recent: true });
        if (res.data.code === 0) {
          recentAchievements.value = res.data.data.list || [];
        }
      } catch (error) {
        console.error("加载成就数据失败:", error);
      }
    });
    const handleAvatarClick = () => {
      if (!isUserLoggedIn.value) {
        handleWxLogin();
      } else {
        goToAuthPage();
      }
    };
    const handleWxLogin = () => __async(this, null, function* () {
      if (loggingIn.value) return;
      try {
        loggingIn.value = true;
        showLoading.value = true;
        loadingText.value = "微信登录中...";
        const result = yield utils_auth.wxLogin();
        isUserLoggedIn.value = true;
        checkLoginStatus();
        yield loadUserProfile();
        common_vendor.index.showToast({
          title: "登录成功",
          icon: "success"
        });
        setTimeout(() => {
          common_vendor.index.showModal({
            title: "完善个人信息",
            content: "是否要设置头像和昵称，获得更好的使用体验？",
            confirmText: "去设置",
            cancelText: "暂不",
            success: (res) => {
              if (res.confirm) {
                common_vendor.index.navigateTo({
                  url: "/pages/profile/auth"
                });
              }
            }
          });
        }, 1500);
      } catch (error) {
        console.error("微信登录失败:", error);
        common_vendor.index.showToast({
          title: error.message || "登录失败",
          icon: "error"
        });
      } finally {
        loggingIn.value = false;
        showLoading.value = false;
      }
    });
    const goToSettings = () => {
      common_vendor.index.navigateTo({
        url: "/pages/profile/privacy"
      });
    };
    const goToSetup = () => {
      common_vendor.index.navigateTo({
        url: "/pages/profile/setup"
      });
    };
    const goToAuthPage = () => {
      common_vendor.index.navigateTo({
        url: "/pages/profile/auth"
      });
    };
    const handleDataExport = () => {
      showExportModal.value = true;
    };
    const closeExportModal = () => {
      showExportModal.value = false;
      exportDataTypes.value = ["profile", "checkin"];
      exportFormat.value = "json";
      exportFormatIndex.value = 0;
      dateRangeIndex.value = 0;
    };
    const onExportFormatChange = (e) => {
      exportFormatIndex.value = e.detail.value;
      const formats = ["json", "csv", "excel"];
      exportFormat.value = formats[e.detail.value];
    };
    const onDataTypeChange = (e) => {
      const { value } = e.detail;
      exportDataTypes.value = value;
    };
    const onDateRangeChange = (e) => {
      dateRangeIndex.value = e.detail.value;
    };
    const startExport = () => __async(this, null, function* () {
      if (exportDataTypes.value.length === 0) {
        common_vendor.index.showToast({
          title: "请选择导出内容",
          icon: "none"
        });
        return;
      }
      try {
        showLoading.value = true;
        loadingText.value = "导出数据中...";
        const exportReq = {
          format: exportFormat.value,
          dataTypes: exportDataTypes.value,
          dateRange: dateRangeIndex.value
        };
        const res = yield apis_user.userApi.createDataExport(exportReq);
        if (res.data.code === 0) {
          closeExportModal();
          common_vendor.index.showToast({
            title: "导出成功",
            icon: "success"
          });
        } else {
          throw new Error(res.data.msg || "导出失败");
        }
      } catch (error) {
        console.error("数据导出失败:", error);
        common_vendor.index.showToast({
          title: error.message || "导出失败",
          icon: "error"
        });
      } finally {
        showLoading.value = false;
      }
    });
    const goToPrivacySettings = () => {
      common_vendor.index.navigateTo({
        url: "/pages/profile/privacy"
      });
    };
    const goToNotificationSettings = () => {
      common_vendor.index.navigateTo({
        url: "/pages/profile/notification"
      });
    };
    const goToHelpCenter = () => {
      common_vendor.index.navigateTo({
        url: "/pages/learning/help"
      });
    };
    const handleAuth = () => {
      if (isUserLoggedIn.value) {
        common_vendor.index.showModal({
          title: "确认退出",
          content: "确定要退出登录吗？",
          success: (res) => {
            if (res.confirm) {
              utils_auth.logout();
              isUserLoggedIn.value = false;
              userProfile.value = {
                nickname: "",
                avatarUrl: "",
                level: 1,
                levelTitle: "新手",
                joinDays: 1,
                currentStreak: 0,
                longestStreak: 0,
                experience: 0,
                achievementCount: 0,
                helpCount: 0
              };
              recentAchievements.value = [];
              common_vendor.index.showToast({
                title: "已退出登录",
                icon: "success"
              });
            }
          }
        });
      }
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.t(currentTime.value),
        b: common_vendor.o(goToSetup),
        c: common_vendor.o(goToSettings),
        d: isLoggedIn.value && userProfile.value.avatarUrl
      }, isLoggedIn.value && userProfile.value.avatarUrl ? {
        e: userProfile.value.avatarUrl
      } : isLoggedIn.value ? {} : {}, {
        f: isLoggedIn.value,
        g: isLoggedIn.value ? 1 : "",
        h: !isLoggedIn.value
      }, !isLoggedIn.value ? {} : {}, {
        i: common_vendor.o(handleAvatarClick),
        j: common_vendor.t(isLoggedIn.value ? userProfile.value.nickname || "用户" : "未登录"),
        k: isLoggedIn.value
      }, isLoggedIn.value ? {
        l: common_vendor.t(userProfile.value.level || 1),
        m: common_vendor.t(userProfile.value.levelTitle || "新手")
      } : {}, {
        n: isLoggedIn.value
      }, isLoggedIn.value ? {
        o: common_vendor.t(userProfile.value.joinDays || 1)
      } : {}, {
        p: common_vendor.t(isLoggedIn.value ? userProfile.value.currentStreak : "--"),
        q: common_vendor.t(isLoggedIn.value ? userProfile.value.longestStreak : "--"),
        r: common_vendor.t(isLoggedIn.value ? userProfile.value.experience : "--"),
        s: common_vendor.t(isLoggedIn.value ? userProfile.value.achievementCount : "--"),
        t: common_vendor.t(isLoggedIn.value ? userProfile.value.helpCount : "--"),
        v: isLoggedIn.value
      }, isLoggedIn.value ? {
        w: common_vendor.f(recentAchievements.value, (achievement, k0, i0) => {
          return {
            a: common_vendor.t(achievement.icon),
            b: common_vendor.n(achievement.rarity),
            c: common_vendor.t(achievement.title),
            d: common_vendor.t(achievement.daysAgo),
            e: achievement.id
          };
        })
      } : {}, {
        x: !isLoggedIn.value
      }, !isLoggedIn.value ? {
        y: common_vendor.t(loggingIn.value ? "登录中..." : "微信快速登录"),
        z: common_vendor.o(handleWxLogin),
        A: loggingIn.value
      } : {}, {
        B: isLoggedIn.value
      }, isLoggedIn.value ? {
        C: common_vendor.o(handleAuth)
      } : {}, {
        D: isLoggedIn.value
      }, isLoggedIn.value ? {
        E: common_vendor.o(handleDataExport)
      } : {}, {
        F: common_vendor.o(goToPrivacySettings),
        G: common_vendor.o(goToNotificationSettings),
        H: common_vendor.o(goToHelpCenter),
        I: showExportModal.value
      }, showExportModal.value ? {
        J: common_vendor.o(closeExportModal),
        K: common_vendor.t(["JSON格式", "CSV格式", "Excel格式"][exportFormatIndex.value]),
        L: ["JSON格式", "CSV格式", "Excel格式"],
        M: common_vendor.o(onExportFormatChange),
        N: common_vendor.o(onDataTypeChange),
        O: common_vendor.t(dateRangeOptions.value[dateRangeIndex.value]),
        P: dateRangeOptions.value,
        Q: common_vendor.o(onDateRangeChange),
        R: common_vendor.o(startExport),
        S: exportDataTypes.value.length === 0,
        T: common_vendor.o(() => {
        }),
        U: common_vendor.o(closeExportModal)
      } : {}, {
        V: showLoading.value
      }, showLoading.value ? {
        W: common_vendor.t(loadingText.value)
      } : {});
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-f97f9319"]]);
wx.createPage(MiniProgramPage);
