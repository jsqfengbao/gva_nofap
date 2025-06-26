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
  name: "CheckinIndex",
  setup() {
    const currentTime = common_vendor.ref("");
    const hasCheckedToday = common_vendor.ref(false);
    const selectedMood = common_vendor.ref(0);
    const checkinNotes = common_vendor.ref("");
    const isSubmitting = common_vendor.ref(false);
    const showSuccessModal = common_vendor.ref(false);
    const todayCheckin = common_vendor.reactive({
      checkinTime: "",
      moodLevel: 0,
      notes: "",
      rewards: 0
    });
    const statistics = common_vendor.reactive({
      currentStreak: 0,
      totalDays: 0,
      level: 1,
      experience: 0
    });
    const checkinResult = common_vendor.reactive({
      streak: 0,
      rewards: 0,
      levelUp: false,
      newLevel: 0
    });
    const moodOptions = [
      { level: 1, emoji: "😰", label: "很糟糕" },
      { level: 2, emoji: "😟", label: "不太好" },
      { level: 3, emoji: "😐", label: "一般" },
      { level: 4, emoji: "😊", label: "不错" },
      { level: 5, emoji: "😄", label: "很棒" }
    ];
    const currentDate = common_vendor.computed(() => {
      const now = /* @__PURE__ */ new Date();
      const year = now.getFullYear();
      const month = String(now.getMonth() + 1).padStart(2, "0");
      const day = String(now.getDate()).padStart(2, "0");
      return `${year}-${month}-${day}`;
    });
    const currentWeekday = common_vendor.computed(() => {
      const weekdays = ["周日", "周一", "周二", "周三", "周四", "周五", "周六"];
      return weekdays[(/* @__PURE__ */ new Date()).getDay()];
    });
    const updateTime = () => {
      const now = /* @__PURE__ */ new Date();
      const hours = now.getHours().toString().padStart(2, "0");
      const minutes = now.getMinutes().toString().padStart(2, "0");
      currentTime.value = `${hours}:${minutes}`;
    };
    const selectMood = (level) => {
      selectedMood.value = level;
    };
    const getMoodEmoji = (level) => {
      const mood = moodOptions.find((m) => m.level === level);
      return mood ? mood.emoji : "😐";
    };
    const getMoodLabel = (level) => {
      const mood = moodOptions.find((m) => m.level === level);
      return mood ? mood.label : "一般";
    };
    const loadTodayStatus = () => __async(this, null, function* () {
      try {
        const token = common_vendor.index.getStorageSync("token");
        if (!token) {
          console.log("用户未登录");
          return;
        }
        const res = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/checkin/today",
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (res.data.code === 0) {
          const data = res.data.data;
          hasCheckedToday.value = data.hasChecked;
          if (data.hasChecked) {
            todayCheckin.checkinTime = data.checkinTime ? new Date(data.checkinTime).toLocaleTimeString("zh-CN", {
              hour: "2-digit",
              minute: "2-digit"
            }) : "";
            todayCheckin.moodLevel = data.moodLevel;
            todayCheckin.notes = data.notes || "";
            todayCheckin.rewards = data.rewards || 0;
          }
        }
      } catch (error) {
        console.error("获取今日状态失败:", error);
      }
    });
    const loadStatistics = () => __async(this, null, function* () {
      try {
        const token = common_vendor.index.getStorageSync("token");
        const res = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/checkin/statistics",
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (res.data.code === 0) {
          const data = res.data.data;
          Object.assign(statistics, data);
        }
      } catch (error) {
        console.error("获取统计数据失败:", error);
      }
    });
    const submitCheckin = () => __async(this, null, function* () {
      if (!selectedMood.value || isSubmitting.value) return;
      isSubmitting.value = true;
      try {
        const token = common_vendor.index.getStorageSync("token");
        const res = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/checkin/daily",
          method: "POST",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          },
          data: {
            moodLevel: selectedMood.value,
            notes: checkinNotes.value
          }
        });
        if (res.data.code === 0) {
          const data = res.data.data;
          Object.assign(checkinResult, {
            streak: data.streak,
            rewards: data.rewards,
            levelUp: data.levelUp,
            newLevel: data.newLevel
          });
          showSuccessModal.value = true;
          yield loadTodayStatus();
          yield loadStatistics();
        } else {
          common_vendor.index.showToast({
            title: res.data.msg || "打卡失败",
            icon: "none"
          });
        }
      } catch (error) {
        console.error("打卡失败:", error);
        common_vendor.index.showToast({
          title: "网络错误，请重试",
          icon: "none"
        });
      } finally {
        isSubmitting.value = false;
      }
    });
    const closeSuccessModal = () => {
      showSuccessModal.value = false;
    };
    const viewHistory = () => {
      common_vendor.index.navigateTo({ url: "/pages/checkin/history" });
    };
    const goToHistory = () => {
      common_vendor.index.navigateTo({ url: "/pages/checkin/history" });
    };
    const goToWeekly = () => {
      common_vendor.index.navigateTo({ url: "/pages/checkin/weekly" });
    };
    const goToCalendar = () => {
      common_vendor.index.navigateTo({ url: "/pages/checkin/calendar" });
    };
    common_vendor.onMounted(() => {
      updateTime();
      setInterval(updateTime, 6e4);
      loadTodayStatus();
      loadStatistics();
    });
    common_vendor.onShow(() => {
      loadTodayStatus();
      loadStatistics();
    });
    return {
      // 数据
      currentTime,
      hasCheckedToday,
      selectedMood,
      checkinNotes,
      isSubmitting,
      showSuccessModal,
      todayCheckin,
      statistics,
      checkinResult,
      moodOptions,
      // 计算属性
      currentDate,
      currentWeekday,
      // 方法
      updateTime,
      selectMood,
      getMoodEmoji,
      getMoodLabel,
      submitCheckin,
      closeSuccessModal,
      viewHistory,
      goToHistory,
      goToWeekly,
      goToCalendar
    };
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
    a: common_vendor.t($setup.currentTime),
    b: common_vendor.o($setup.viewHistory),
    c: common_vendor.p({
      title: "每日打卡",
      ["show-back"]: false,
      ["right-text"]: "历史"
    }),
    d: common_vendor.t($setup.currentDate),
    e: common_vendor.t($setup.currentWeekday),
    f: !$setup.hasCheckedToday
  }, !$setup.hasCheckedToday ? {} : {
    g: common_vendor.t($setup.todayCheckin.checkinTime)
  }, {
    h: common_vendor.p({
      type: "gradient"
    }),
    i: common_vendor.t($setup.statistics.currentStreak),
    j: common_vendor.t($setup.statistics.totalDays),
    k: common_vendor.t($setup.statistics.level),
    l: !$setup.hasCheckedToday
  }, !$setup.hasCheckedToday ? {
    m: common_vendor.f($setup.moodOptions, (mood, k0, i0) => {
      return {
        a: common_vendor.t(mood.emoji),
        b: common_vendor.t(mood.label),
        c: $setup.selectedMood === mood.level ? 1 : "",
        d: mood.level,
        e: common_vendor.o(($event) => $setup.selectMood(mood.level), mood.level)
      };
    })
  } : {}, {
    n: !$setup.hasCheckedToday
  }, !$setup.hasCheckedToday ? {
    o: $setup.checkinNotes,
    p: common_vendor.o(($event) => $setup.checkinNotes = $event.detail.value),
    q: common_vendor.t($setup.checkinNotes.length)
  } : {}, {
    r: !$setup.hasCheckedToday
  }, !$setup.hasCheckedToday ? common_vendor.e({
    s: !$setup.isSubmitting
  }, !$setup.isSubmitting ? {
    t: common_vendor.t($setup.selectedMood ? "立即打卡" : "请先选择心情")
  } : {}, {
    v: common_vendor.o($setup.submitCheckin),
    w: common_vendor.p({
      type: "primary",
      size: "large",
      ["full-width"]: true,
      disabled: !$setup.selectedMood || $setup.isSubmitting
    })
  }) : common_vendor.e({
    x: common_vendor.t($setup.getMoodEmoji($setup.todayCheckin.moodLevel)),
    y: common_vendor.t($setup.getMoodLabel($setup.todayCheckin.moodLevel)),
    z: $setup.todayCheckin.notes
  }, $setup.todayCheckin.notes ? {
    A: common_vendor.t($setup.todayCheckin.notes)
  } : {}, {
    B: common_vendor.t($setup.todayCheckin.rewards)
  }), {
    C: common_vendor.o((...args) => $setup.goToHistory && $setup.goToHistory(...args)),
    D: common_vendor.o((...args) => $setup.goToWeekly && $setup.goToWeekly(...args)),
    E: common_vendor.o((...args) => $setup.goToCalendar && $setup.goToCalendar(...args)),
    F: $setup.showSuccessModal
  }, $setup.showSuccessModal ? common_vendor.e({
    G: common_vendor.t($setup.checkinResult.streak),
    H: common_vendor.t($setup.checkinResult.rewards),
    I: $setup.checkinResult.levelUp
  }, $setup.checkinResult.levelUp ? {
    J: common_vendor.t($setup.checkinResult.newLevel)
  } : {}, {
    K: common_vendor.o($setup.closeSuccessModal),
    L: common_vendor.p({
      type: "primary"
    }),
    M: common_vendor.o(() => {
    }),
    N: common_vendor.o((...args) => $setup.closeSuccessModal && $setup.closeSuccessModal(...args))
  }) : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-a8aeee56"]]);
wx.createPage(MiniProgramPage);
