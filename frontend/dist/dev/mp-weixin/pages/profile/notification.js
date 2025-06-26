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
  __name: "notification",
  setup(__props) {
    const loading = common_vendor.ref(false);
    const saving = common_vendor.ref(false);
    const loadingText = common_vendor.ref("加载中...");
    const settings = common_vendor.ref({
      checkinReminder: true,
      communityReply: true,
      achievementUnlock: true,
      weeklyReport: true,
      emergencyAlert: true,
      learningReminder: true
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
    const onCheckinReminderChange = (e) => {
      settings.value.checkinReminder = e.detail.value;
    };
    const onCommunityReplyChange = (e) => {
      settings.value.communityReply = e.detail.value;
    };
    const onAchievementUnlockChange = (e) => {
      settings.value.achievementUnlock = e.detail.value;
    };
    const onWeeklyReportChange = (e) => {
      settings.value.weeklyReport = e.detail.value;
    };
    const onEmergencyAlertChange = (e) => {
      settings.value.emergencyAlert = e.detail.value;
    };
    const onLearningReminderChange = (e) => {
      settings.value.learningReminder = e.detail.value;
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.o(goBack),
        b: settings.value.checkinReminder,
        c: common_vendor.o(onCheckinReminderChange),
        d: settings.value.communityReply,
        e: common_vendor.o(onCommunityReplyChange),
        f: settings.value.achievementUnlock,
        g: common_vendor.o(onAchievementUnlockChange),
        h: settings.value.weeklyReport,
        i: common_vendor.o(onWeeklyReportChange),
        j: settings.value.emergencyAlert,
        k: common_vendor.o(onEmergencyAlertChange),
        l: settings.value.learningReminder,
        m: common_vendor.o(onLearningReminderChange),
        n: common_vendor.t(saving.value ? "保存中..." : "保存设置"),
        o: common_vendor.o(saveSettings),
        p: saving.value,
        q: loading.value
      }, loading.value ? {
        r: common_vendor.t(loadingText.value)
      } : {});
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-52942290"]]);
wx.createPage(MiniProgramPage);
