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
const _sfc_main = {
  name: "CalendarPage",
  components: {
    NfNavBar,
    NfCard
  },
  setup() {
    const calendarData = common_vendor.ref({
      year: (/* @__PURE__ */ new Date()).getFullYear(),
      month: (/* @__PURE__ */ new Date()).getMonth() + 1,
      days: [],
      summary: {
        checkedDays: 0,
        totalDays: 0,
        successRate: 0,
        averageMood: 0,
        bestStreak: 0
      }
    });
    const loading = common_vendor.ref(false);
    const weekdays = ["日", "一", "二", "三", "四", "五", "六"];
    const currentTime = common_vendor.computed(() => {
      return (/* @__PURE__ */ new Date()).toLocaleTimeString("zh-CN", {
        hour: "2-digit",
        minute: "2-digit"
      });
    });
    const calendarDays = common_vendor.computed(() => {
      const year = calendarData.value.year;
      const month = calendarData.value.month;
      const firstDay = new Date(year, month - 1, 1);
      const lastDay = new Date(year, month, 0);
      const startWeekday = firstDay.getDay();
      const prevMonth = month === 1 ? 12 : month - 1;
      const prevYear = month === 1 ? year - 1 : year;
      const prevMonthLastDay = new Date(prevYear, prevMonth, 0).getDate();
      const nextMonth2 = month === 12 ? 1 : month + 1;
      const nextYear = month === 12 ? year + 1 : year;
      const days = [];
      for (let i = startWeekday - 1; i >= 0; i--) {
        const day = prevMonthLastDay - i;
        days.push({
          day,
          date: `${prevYear}-${String(prevMonth).padStart(2, "0")}-${String(day).padStart(2, "0")}`,
          hasChecked: false,
          moodLevel: 0,
          isToday: false,
          isThisMonth: false
        });
      }
      const today = (/* @__PURE__ */ new Date()).toISOString().split("T")[0];
      for (let day = 1; day <= lastDay.getDate(); day++) {
        const dateStr = `${year}-${String(month).padStart(2, "0")}-${String(day).padStart(2, "0")}`;
        const dayData = calendarData.value.days.find((d) => d.date === dateStr) || {};
        days.push({
          day,
          date: dateStr,
          hasChecked: dayData.hasChecked || false,
          moodLevel: dayData.moodLevel || 0,
          isToday: dateStr === today,
          isThisMonth: true
        });
      }
      const remainingDays = 42 - days.length;
      for (let day = 1; day <= remainingDays; day++) {
        days.push({
          day,
          date: `${nextYear}-${String(nextMonth2).padStart(2, "0")}-${String(day).padStart(2, "0")}`,
          hasChecked: false,
          moodLevel: 0,
          isToday: false,
          isThisMonth: false
        });
      }
      return days;
    });
    const showDayDetail = (day) => {
      if (!day.isThisMonth) return;
      console.log("Day detail:", day);
    };
    const previousMonth = () => {
      if (calendarData.value.month === 1) {
        calendarData.value.year--;
        calendarData.value.month = 12;
      } else {
        calendarData.value.month--;
      }
      loadCalendarData();
    };
    const nextMonth = () => {
      if (calendarData.value.month === 12) {
        calendarData.value.year++;
        calendarData.value.month = 1;
      } else {
        calendarData.value.month++;
      }
      loadCalendarData();
    };
    const goBack = () => {
      common_vendor.index.navigateBack();
    };
    const loadCalendarData = () => __async(this, null, function* () {
      if (loading.value) return;
      loading.value = true;
      try {
        const token = common_vendor.index.getStorageSync("token");
        if (!token) {
          common_vendor.index.redirectTo({ url: "/pages/auth/login" });
          return;
        }
        const res = yield common_vendor.index.request({
          url: `http://localhost:8888/api/v1/miniprogram/checkin/calendar?year=${calendarData.value.year}&month=${calendarData.value.month}`,
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (res.data.code === 0) {
          Object.assign(calendarData.value, res.data.data);
        } else {
          common_vendor.index.showToast({
            title: res.data.msg || "获取数据失败",
            icon: "none"
          });
        }
      } catch (error) {
        console.error("获取日历数据失败:", error);
        common_vendor.index.showToast({
          title: "网络错误，请重试",
          icon: "none"
        });
      } finally {
        loading.value = false;
      }
    });
    common_vendor.onMounted(() => {
      loadCalendarData();
    });
    return {
      // 数据
      calendarData,
      loading,
      weekdays,
      // 计算属性
      currentTime,
      calendarDays,
      // 方法
      showDayDetail,
      previousMonth,
      nextMonth,
      goBack
    };
  }
};
if (!Array) {
  const _component_NfNavBar = common_vendor.resolveComponent("NfNavBar");
  const _component_NfCard = common_vendor.resolveComponent("NfCard");
  (_component_NfNavBar + _component_NfCard)();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.t($setup.currentTime),
    b: common_vendor.o($setup.goBack),
    c: common_vendor.p({
      title: "月度日历",
      ["show-back"]: true
    }),
    d: common_vendor.o((...args) => $setup.previousMonth && $setup.previousMonth(...args)),
    e: common_vendor.t($setup.calendarData.year),
    f: common_vendor.t($setup.calendarData.month),
    g: common_vendor.o((...args) => $setup.nextMonth && $setup.nextMonth(...args)),
    h: common_vendor.t($setup.calendarData.summary.checkedDays),
    i: common_vendor.t($setup.calendarData.summary.successRate),
    j: common_vendor.t($setup.calendarData.summary.averageMood),
    k: common_vendor.t($setup.calendarData.summary.bestStreak),
    l: common_vendor.p({
      type: "gradient"
    }),
    m: common_vendor.f($setup.weekdays, (weekday, k0, i0) => {
      return {
        a: common_vendor.t(weekday),
        b: weekday
      };
    }),
    n: common_vendor.f($setup.calendarDays, (day, index, i0) => {
      return common_vendor.e({
        a: common_vendor.t(day.day),
        b: day.hasChecked && day.isThisMonth
      }, day.hasChecked && day.isThisMonth ? {
        c: common_vendor.n("mood-" + day.moodLevel)
      } : day.isThisMonth ? {} : {}, {
        d: day.isThisMonth,
        e: index,
        f: day.isToday ? 1 : "",
        g: day.hasChecked ? 1 : "",
        h: !day.isThisMonth ? 1 : "",
        i: common_vendor.o(($event) => $setup.showDayDetail(day), index)
      });
    })
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-5c084a8a"]]);
wx.createPage(MiniProgramPage);
