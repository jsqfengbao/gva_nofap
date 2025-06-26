"use strict";
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
const _sfc_main = {
  __name: "privacy",
  setup(__props) {
    const loading = common_vendor.ref(false);
    const saving = common_vendor.ref(false);
    const loadingText = common_vendor.ref("加载中...");
    const showExportModal = common_vendor.ref(false);
    const showDeleteModal = common_vendor.ref(false);
    const settings = common_vendor.ref({
      showProfile: true,
      showStats: true,
      showAchievements: true,
      allowFriendRequest: true,
      showOnlineStatus: true
    });
    common_vendor.onMounted(() => {
      loadSettings();
    });
    const goBack = () => {
      common_vendor.index.navigateBack();
    };
    const loadSettings = () => __async(this, null, function* () {
      try {
        loading.value = true;
        loadingText.value = "加载设置...";
        yield new Promise((resolve) => setTimeout(resolve, 800));
        loading.value = false;
      } catch (error) {
        console.error("加载设置失败:", error);
        loading.value = false;
        common_vendor.index.showToast({
          title: "加载失败",
          icon: "error"
        });
      }
    });
    const saveSettings = () => __async(this, null, function* () {
      try {
        saving.value = true;
        yield new Promise((resolve) => setTimeout(resolve, 1e3));
        saving.value = false;
        common_vendor.index.showToast({
          title: "保存成功",
          icon: "success"
        });
      } catch (error) {
        console.error("保存设置失败:", error);
        saving.value = false;
        common_vendor.index.showToast({
          title: "保存失败",
          icon: "error"
        });
      }
    });
    const onShowProfileChange = (e) => {
      settings.value.showProfile = e.detail.value;
    };
    const onShowStatsChange = (e) => {
      settings.value.showStats = e.detail.value;
    };
    const onShowAchievementsChange = (e) => {
      settings.value.showAchievements = e.detail.value;
    };
    const onAllowFriendRequestChange = (e) => {
      settings.value.allowFriendRequest = e.detail.value;
    };
    const onShowOnlineStatusChange = (e) => {
      settings.value.showOnlineStatus = e.detail.value;
    };
    const showDataExportInfo = () => {
      showExportModal.value = true;
    };
    const closeExportModal = () => {
      showExportModal.value = false;
    };
    const showDeleteAccountInfo = () => {
      showDeleteModal.value = true;
    };
    const closeDeleteModal = () => {
      showDeleteModal.value = false;
    };
    const goToProfile = () => {
      closeExportModal();
      common_vendor.index.switchTab({
        url: "/pages/profile/index"
      });
    };
    const contactSupport = () => {
      closeDeleteModal();
      common_vendor.index.showModal({
        title: "联系客服",
        content: "请发送邮件至 support@nofap-helper.com 或在社区发帖寻求帮助",
        showCancel: false,
        confirmText: "我知道了"
      });
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.o(goBack),
        b: settings.value.showProfile,
        c: common_vendor.o(onShowProfileChange),
        d: settings.value.showStats,
        e: common_vendor.o(onShowStatsChange),
        f: settings.value.showAchievements,
        g: common_vendor.o(onShowAchievementsChange),
        h: settings.value.allowFriendRequest,
        i: common_vendor.o(onAllowFriendRequestChange),
        j: settings.value.showOnlineStatus,
        k: common_vendor.o(onShowOnlineStatusChange),
        l: common_vendor.o(showDataExportInfo),
        m: common_vendor.o(showDeleteAccountInfo),
        n: common_vendor.t(saving.value ? "保存中..." : "保存设置"),
        o: common_vendor.o(saveSettings),
        p: saving.value,
        q: showExportModal.value
      }, showExportModal.value ? {
        r: common_vendor.o(closeExportModal),
        s: common_vendor.o(closeExportModal),
        t: common_vendor.o(goToProfile),
        v: common_vendor.o(() => {
        }),
        w: common_vendor.o(closeExportModal)
      } : {}, {
        x: showDeleteModal.value
      }, showDeleteModal.value ? {
        y: common_vendor.o(closeDeleteModal),
        z: common_vendor.o(closeDeleteModal),
        A: common_vendor.o(contactSupport),
        B: common_vendor.o(() => {
        }),
        C: common_vendor.o(closeDeleteModal)
      } : {}, {
        D: loading.value
      }, loading.value ? {
        E: common_vendor.t(loadingText.value)
      } : {});
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-984500fe"]]);
wx.createPage(MiniProgramPage);
