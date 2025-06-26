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
const NfNavBar = () => "../../components/ui/navigation/NfNavBar.js";
const NfCard = () => "../../components/ui/card/NfCard.js";
const NfButton = () => "../../components/ui/button/NfButton.js";
const _sfc_main = {
  name: "WeeklyProgress",
  components: {
    NfNavBar,
    NfCard,
    NfButton
  },
  setup() {
    const weeklyData = common_vendor.ref({
      weekDays: [],
      summary: {
        checkedDays: 0,
        totalDays: 7,
        successRate: 0,
        averageMood: 0
      }
    });
    const loading = common_vendor.ref(false);
    const showDetailModal = common_vendor.ref(false);
    const selectedDay = common_vendor.ref(null);
    const currentTime = common_vendor.computed(() => {
      return (/* @__PURE__ */ new Date()).toLocaleTimeString("zh-CN", {
        hour: "2-digit",
        minute: "2-digit"
      });
    });
    const weekRange = common_vendor.computed(() => {
      var _a, _b;
      if (weeklyData.value.weekDays.length === 0) return "";
      const firstDay = (_a = weeklyData.value.weekDays[0]) == null ? void 0 : _a.date;
      const lastDay = (_b = weeklyData.value.weekDays[6]) == null ? void 0 : _b.date;
      if (!firstDay || !lastDay) return "";
      const startDate = new Date(firstDay);
      const endDate = new Date(lastDay);
      return `${formatDateRange(startDate)} - ${formatDateRange(endDate)}`;
    });
    const formatDate = (dateStr) => {
      const date = new Date(dateStr);
      return `${date.getMonth() + 1}/${date.getDate()}`;
    };
    const formatDateRange = (date) => {
      return `${date.getMonth() + 1}月${date.getDate()}日`;
    };
    const getShortWeekday = (weekday) => {
      const map = {
        "周一": "一",
        "周二": "二",
        "周三": "三",
        "周四": "四",
        "周五": "五",
        "周六": "六",
        "周日": "日"
      };
      return map[weekday] || weekday;
    };
    const getMoodEmoji = (level) => {
      const emojis = ["", "😢", "😕", "😐", "😊", "😄"];
      return emojis[level] || "😐";
    };
    const getMoodLabel = (level) => {
      const labels = ["", "很糟糕", "不太好", "一般", "不错", "很开心"];
      return labels[level] || "一般";
    };
    const getMoodBarHeight = (moodLevel) => {
      if (!moodLevel) return "0%";
      return `${moodLevel / 5 * 100}%`;
    };
    const showDayDetail = (day) => {
      selectedDay.value = day;
      showDetailModal.value = true;
    };
    const closeDetailModal = () => {
      showDetailModal.value = false;
      selectedDay.value = null;
    };
    const goBack = () => {
      common_vendor.index.navigateBack();
    };
    const goToCheckin = () => {
      closeDetailModal();
      common_vendor.index.navigateTo({ url: "/pages/checkin/index" });
    };
    const loadWeeklyData = () => __async(this, null, function* () {
      if (loading.value) return;
      loading.value = true;
      try {
        const token = common_vendor.index.getStorageSync("token");
        if (!token) {
          common_vendor.index.redirectTo({ url: "/pages/auth/login" });
          return;
        }
        const res = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/checkin/weekly",
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (res.data.code === 0) {
          weeklyData.value = res.data.data;
        } else {
          common_vendor.index.showToast({
            title: res.data.msg || "获取数据失败",
            icon: "none"
          });
        }
      } catch (error) {
        console.error("获取本周进度失败:", error);
        common_vendor.index.showToast({
          title: "网络错误，请重试",
          icon: "none"
        });
      } finally {
        loading.value = false;
      }
    });
    common_vendor.onMounted(() => {
      loadWeeklyData();
    });
    return {
      // 数据
      weeklyData,
      loading,
      showDetailModal,
      selectedDay,
      // 计算属性
      currentTime,
      weekRange,
      // 方法
      formatDate,
      getShortWeekday,
      getMoodEmoji,
      getMoodLabel,
      getMoodBarHeight,
      showDayDetail,
      closeDetailModal,
      goBack,
      goToCheckin
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
  var _a, _b, _c, _d, _e, _f, _g, _h, _i;
  return common_vendor.e({
    a: common_vendor.t($setup.currentTime),
    b: common_vendor.o($setup.goBack),
    c: common_vendor.p({
      title: "本周进度",
      ["show-back"]: true
    }),
    d: common_vendor.t($setup.weekRange),
    e: common_vendor.t($setup.weeklyData.summary.checkedDays),
    f: common_vendor.t($setup.weeklyData.summary.successRate),
    g: common_vendor.t($setup.weeklyData.summary.averageMood),
    h: common_vendor.p({
      type: "gradient"
    }),
    i: common_vendor.f($setup.weeklyData.weekDays, (day, index, i0) => {
      return common_vendor.e({
        a: common_vendor.t(day.weekday),
        b: common_vendor.t($setup.formatDate(day.date)),
        c: day.hasChecked
      }, day.hasChecked ? {} : {}, {
        d: day.hasChecked ? 1 : "",
        e: day.hasChecked
      }, day.hasChecked ? {
        f: common_vendor.t($setup.getMoodEmoji(day.moodLevel)),
        g: common_vendor.t($setup.getMoodLabel(day.moodLevel))
      } : {}, {
        h: index,
        i: day.hasChecked ? 1 : "",
        j: day.isToday ? 1 : "",
        k: !day.hasChecked ? 1 : "",
        l: common_vendor.o(($event) => $setup.showDayDetail(day), index)
      });
    }),
    j: common_vendor.f([5, 4, 3, 2, 1], (level, k0, i0) => {
      return {
        a: common_vendor.t(level),
        b: level
      };
    }),
    k: common_vendor.f(5, (i, k0, i0) => {
      return {
        a: i
      };
    }),
    l: common_vendor.f($setup.weeklyData.weekDays, (day, index, i0) => {
      return {
        a: $setup.getMoodBarHeight(day.moodLevel),
        b: day.hasChecked ? 1 : "",
        c: common_vendor.t($setup.getShortWeekday(day.weekday)),
        d: index
      };
    }),
    m: $setup.weeklyData.summary.successRate >= 80
  }, $setup.weeklyData.summary.successRate >= 80 ? {
    n: common_vendor.t($setup.weeklyData.summary.successRate),
    o: common_vendor.p({
      type: "success"
    })
  } : {}, {
    p: $setup.showDetailModal
  }, $setup.showDetailModal ? common_vendor.e({
    q: common_vendor.t((_a = $setup.selectedDay) == null ? void 0 : _a.weekday),
    r: common_vendor.o((...args) => $setup.closeDetailModal && $setup.closeDetailModal(...args)),
    s: common_vendor.t((_b = $setup.selectedDay) == null ? void 0 : _b.date),
    t: (_c = $setup.selectedDay) == null ? void 0 : _c.hasChecked
  }, ((_d = $setup.selectedDay) == null ? void 0 : _d.hasChecked) ? {
    v: common_vendor.t($setup.getMoodEmoji($setup.selectedDay.moodLevel)),
    w: common_vendor.t($setup.getMoodLabel($setup.selectedDay.moodLevel))
  } : {
    x: common_vendor.t(((_e = $setup.selectedDay) == null ? void 0 : _e.isToday) ? "今天还可以去打卡哦！" : "这天没有打卡记录")
  }, {
    y: ((_f = $setup.selectedDay) == null ? void 0 : _f.isToday) && !((_g = $setup.selectedDay) == null ? void 0 : _g.hasChecked)
  }, ((_h = $setup.selectedDay) == null ? void 0 : _h.isToday) && !((_i = $setup.selectedDay) == null ? void 0 : _i.hasChecked) ? {
    z: common_vendor.o($setup.goToCheckin),
    A: common_vendor.p({
      type: "primary",
      size: "medium"
    })
  } : {}, {
    B: common_vendor.o(() => {
    }),
    C: common_vendor.o((...args) => $setup.closeDetailModal && $setup.closeDetailModal(...args))
  }) : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-b4878256"]]);
wx.createPage(MiniProgramPage);
