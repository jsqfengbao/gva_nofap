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
  __name: "index",
  setup(__props) {
    const currentTime = common_vendor.ref("");
    const currentStreak = common_vendor.ref(0);
    const longestStreak = common_vendor.ref(0);
    const weeklyData = common_vendor.ref([
      { label: "周一", hasChecked: true, height: 75 },
      { label: "周二", hasChecked: true, height: 62 },
      { label: "周三", hasChecked: true, height: 87 },
      { label: "周四", hasChecked: true, height: 50 },
      { label: "周五", hasChecked: true, height: 100 },
      { label: "周六", hasChecked: false, height: 25 },
      { label: "周日", hasChecked: false, height: 25 }
    ]);
    const progressAchievements = common_vendor.ref([
      {
        id: 1,
        name: "30天挑战",
        icon: "🏅",
        category: "streak",
        current: 23,
        target: 30,
        progress: 77
      },
      {
        id: 2,
        name: "社区贡献者",
        icon: "⭐",
        category: "community",
        current: 6,
        target: 10,
        progress: 60
      },
      {
        id: 3,
        name: "学习达人",
        icon: "📚",
        category: "learning",
        current: 8,
        target: 20,
        progress: 40
      }
    ]);
    const monthlyStats = common_vendor.reactive({
      successRate: 85,
      learningDays: 12,
      experience: 156
    });
    const userStats = common_vendor.reactive({
      currentStreak: 0,
      longestStreak: 0,
      totalDays: 0,
      successRate: 0
    });
    const weeklyCompletionRate = common_vendor.computed(() => {
      const checkedDays = weeklyData.value.filter((day) => day.hasChecked).length;
      return Math.round(checkedDays / weeklyData.value.length * 100);
    });
    const updateTime = () => {
      const now = /* @__PURE__ */ new Date();
      const hours = now.getHours().toString().padStart(2, "0");
      const minutes = now.getMinutes().toString().padStart(2, "0");
      currentTime.value = `${hours}:${minutes}`;
    };
    const loadProgressData = () => __async(this, null, function* () {
      try {
        const token = common_vendor.index.getStorageSync("token");
        if (!token) {
          console.log("用户未登录");
          return;
        }
        const statsRes = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/checkin/statistics",
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (statsRes.data.code === 0) {
          const data = statsRes.data.data;
          Object.assign(userStats, data);
          currentStreak.value = data.currentStreak || 0;
          longestStreak.value = data.longestStreak || 0;
        }
        const weeklyRes = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/checkin/weekly-progress",
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (weeklyRes.data.code === 0) {
          const weeklyProgressData = weeklyRes.data.data;
          updateWeeklyChart(weeklyProgressData);
        }
        const achievementsRes = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/achievement/progress",
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (achievementsRes.data.code === 0) {
          const achievementsData = achievementsRes.data.data;
          updateAchievementProgress(achievementsData);
        }
      } catch (error) {
        console.error("加载进度数据失败:", error);
      }
    });
    const updateWeeklyChart = (data) => {
      if (!data || !data.weeklyCheckins) return;
      const days = ["周一", "周二", "周三", "周四", "周五", "周六", "周日"];
      weeklyData.value = data.weeklyCheckins.map((checkin, index) => ({
        label: days[index],
        hasChecked: checkin.hasChecked || false,
        height: checkin.hasChecked ? Math.max(50, checkin.moodLevel * 20) : 25
      }));
    };
    const updateAchievementProgress = (achievements) => {
      if (!achievements || !achievements.length) return;
      progressAchievements.value = achievements.map((achievement) => ({
        id: achievement.id,
        name: achievement.name,
        icon: getAchievementIcon(achievement.category),
        category: achievement.category,
        current: achievement.progress || 0,
        target: achievement.targetValue || 100,
        progress: Math.round((achievement.progress || 0) / (achievement.targetValue || 100) * 100)
      }));
    };
    const getAchievementIcon = (category) => {
      const icons = {
        "checkin": "🏅",
        "streak": "🔥",
        "community": "⭐",
        "learning": "📚",
        "level": "🎮"
      };
      return icons[category] || "🏆";
    };
    const exportData = () => {
      common_vendor.index.showToast({
        title: "数据导出功能开发中",
        icon: "none"
      });
    };
    common_vendor.onMounted(() => {
      updateTime();
      setInterval(updateTime, 6e4);
      loadProgressData();
    });
    common_vendor.onShow(() => {
      loadProgressData();
    });
    return (_ctx, _cache) => {
      return {
        a: common_vendor.t(currentTime.value),
        b: common_vendor.o(exportData),
        c: common_vendor.t(currentStreak.value),
        d: common_vendor.t(longestStreak.value),
        e: common_vendor.f(weeklyData.value, (day, index, i0) => {
          return {
            a: day.hasChecked ? 1 : "",
            b: !day.hasChecked ? 1 : "",
            c: day.height + "%",
            d: common_vendor.t(day.label),
            e: index
          };
        }),
        f: common_vendor.t(weeklyCompletionRate.value),
        g: common_vendor.f(progressAchievements.value, (achievement, k0, i0) => {
          return {
            a: common_vendor.t(achievement.icon),
            b: common_vendor.n(achievement.category),
            c: common_vendor.t(achievement.name),
            d: common_vendor.n(achievement.category),
            e: achievement.progress + "%",
            f: common_vendor.t(achievement.current),
            g: common_vendor.t(achievement.target),
            h: achievement.id
          };
        }),
        h: common_vendor.t(monthlyStats.successRate),
        i: common_vendor.t(monthlyStats.learningDays),
        j: common_vendor.t(monthlyStats.experience)
      };
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-1169f7c9"]]);
wx.createPage(MiniProgramPage);
