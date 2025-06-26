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
  name: "CheckinHistory",
  setup() {
    const loading = common_vendor.ref(false);
    const selectedMonth = common_vendor.ref("");
    const currentPage = common_vendor.ref(1);
    const pageSize = 20;
    const historyData = common_vendor.reactive({
      list: [],
      total: 0,
      page: 1,
      pageSize: 20
    });
    const availableMonths = common_vendor.ref([]);
    const moodOptions = [
      { level: 1, emoji: "😰", label: "很糟糕" },
      { level: 2, emoji: "😟", label: "不太好" },
      { level: 3, emoji: "😐", label: "一般" },
      { level: 4, emoji: "😊", label: "不错" },
      { level: 5, emoji: "😄", label: "很棒" }
    ];
    const hasMore = common_vendor.computed(() => {
      return historyData.list.length < historyData.total;
    });
    const currentMonthCount = common_vendor.computed(() => {
      if (selectedMonth.value) {
        return historyData.list.length;
      }
      const currentMonth = (/* @__PURE__ */ new Date()).toISOString().slice(0, 7);
      return historyData.list.filter(
        (item) => item.checkinDate.startsWith(currentMonth)
      ).length;
    });
    const averageMood = common_vendor.computed(() => {
      if (historyData.list.length === 0) return "0.0";
      const sum = historyData.list.reduce((acc, item) => acc + item.moodLevel, 0);
      return (sum / historyData.list.length).toFixed(1);
    });
    const getMoodEmoji = (level) => {
      const mood = moodOptions.find((m) => m.level === level);
      return mood ? mood.emoji : "😐";
    };
    const getMoodLabel = (level) => {
      const mood = moodOptions.find((m) => m.level === level);
      return mood ? mood.label : "一般";
    };
    const formatDate = (dateString) => {
      const date = new Date(dateString);
      const today = /* @__PURE__ */ new Date();
      const diffTime = today - date;
      const diffDays = Math.ceil(diffTime / (1e3 * 60 * 60 * 24));
      if (diffDays === 0) return "今天";
      if (diffDays === 1) return "昨天";
      if (diffDays === 2) return "前天";
      return `${date.getMonth() + 1}月${date.getDate()}日`;
    };
    const getWeekday = (dateString) => {
      const date = new Date(dateString);
      const weekdays = ["周日", "周一", "周二", "周三", "周四", "周五", "周六"];
      return weekdays[date.getDay()];
    };
    const formatTime = (dateString) => {
      const date = new Date(dateString);
      return date.toLocaleTimeString("zh-CN", {
        hour: "2-digit",
        minute: "2-digit"
      });
    };
    const formatMonth = (monthString) => {
      const [year, month] = monthString.split("-");
      return `${year}年${month}月`;
    };
    const generateAvailableMonths = () => {
      const months = [];
      const now = /* @__PURE__ */ new Date();
      for (let i = 0; i < 6; i++) {
        const date = new Date(now.getFullYear(), now.getMonth() - i, 1);
        const monthString = `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, "0")}`;
        months.push(monthString);
      }
      availableMonths.value = months;
    };
    const selectMonth = (month) => {
      selectedMonth.value = month;
      currentPage.value = 1;
      historyData.list = [];
      loadHistoryData();
    };
    const loadHistoryData = () => __async(this, null, function* () {
      if (loading.value) return;
      loading.value = true;
      try {
        const token = common_vendor.index.getStorageSync("token");
        if (!token) {
          common_vendor.index.redirectTo({ url: "/pages/auth/login" });
          return;
        }
        const params = {
          page: currentPage.value,
          pageSize
        };
        if (selectedMonth.value) {
          params.month = selectedMonth.value;
        }
        const queryString = Object.keys(params).map((key) => `${key}=${encodeURIComponent(params[key])}`).join("&");
        const res = yield common_vendor.index.request({
          url: `http://localhost:8888/api/v1/miniprogram/checkin/history?${queryString}`,
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (res.data.code === 0) {
          const data = res.data.data;
          if (currentPage.value === 1) {
            historyData.list = data.list || [];
          } else {
            historyData.list.push(...data.list || []);
          }
          historyData.total = data.total || 0;
          historyData.page = data.page || 1;
          historyData.pageSize = data.pageSize || pageSize;
        } else {
          common_vendor.index.showToast({
            title: res.data.msg || "获取数据失败",
            icon: "none"
          });
        }
      } catch (error) {
        console.error("获取打卡历史失败:", error);
        common_vendor.index.showToast({
          title: "网络错误，请重试",
          icon: "none"
        });
      } finally {
        loading.value = false;
      }
    });
    const loadMore = () => {
      if (!hasMore.value || loading.value) return;
      currentPage.value++;
      loadHistoryData();
    };
    common_vendor.onMounted(() => {
      generateAvailableMonths();
      loadHistoryData();
    });
    return {
      // 数据
      loading,
      selectedMonth,
      historyData,
      availableMonths,
      // 计算属性
      hasMore,
      currentMonthCount,
      averageMood,
      // 方法
      getMoodEmoji,
      getMoodLabel,
      formatDate,
      getWeekday,
      formatTime,
      formatMonth,
      selectMonth,
      loadMore
    };
  }
};
if (!Array) {
  const _component_NfNavBar = common_vendor.resolveComponent("NfNavBar");
  const _component_NfCard = common_vendor.resolveComponent("NfCard");
  (_component_NfNavBar + _component_NfCard)();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.p({
      title: "打卡历史",
      showBack: true
    }),
    b: common_vendor.o(($event) => $setup.selectMonth("")),
    c: $setup.selectedMonth === "" ? 1 : "",
    d: common_vendor.f($setup.availableMonths, (month, k0, i0) => {
      return {
        a: common_vendor.t($setup.formatMonth(month)),
        b: month,
        c: common_vendor.o(($event) => $setup.selectMonth(month), month),
        d: $setup.selectedMonth === month ? 1 : ""
      };
    }),
    e: common_vendor.t($setup.historyData.total),
    f: common_vendor.t($setup.currentMonthCount),
    g: common_vendor.t($setup.selectedMonth ? "当月" : "本月"),
    h: common_vendor.t($setup.averageMood),
    i: common_vendor.p({
      type: "basic"
    }),
    j: $setup.loading
  }, $setup.loading ? {} : $setup.historyData.list.length === 0 ? {
    l: common_vendor.t($setup.selectedMonth ? "该月份暂无打卡记录" : "暂无打卡记录")
  } : {
    m: common_vendor.f($setup.historyData.list, (checkin, k0, i0) => {
      return common_vendor.e({
        a: common_vendor.t($setup.formatDate(checkin.checkinDate)),
        b: common_vendor.t($setup.getWeekday(checkin.checkinDate)),
        c: common_vendor.t($setup.getMoodEmoji(checkin.moodLevel)),
        d: common_vendor.t($setup.getMoodLabel(checkin.moodLevel)),
        e: common_vendor.f(5, (star, k1, i1) => {
          return {
            a: star,
            b: star <= checkin.moodLevel ? 1 : ""
          };
        }),
        f: checkin.notes
      }, checkin.notes ? {
        g: common_vendor.t(checkin.notes)
      } : {}, {
        h: common_vendor.t($setup.formatTime(checkin.checkinDate)),
        i: common_vendor.t(checkin.rewards),
        j: checkin.id
      });
    })
  }, {
    k: $setup.historyData.list.length === 0,
    n: $setup.hasMore && !$setup.loading
  }, $setup.hasMore && !$setup.loading ? {
    o: common_vendor.o((...args) => $setup.loadMore && $setup.loadMore(...args))
  } : {}, {
    p: !$setup.hasMore && $setup.historyData.list.length > 0
  }, !$setup.hasMore && $setup.historyData.list.length > 0 ? {} : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-f3aa1d8e"]]);
wx.createPage(MiniProgramPage);
