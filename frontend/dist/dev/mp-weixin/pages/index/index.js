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
const apis_home = require("../../apis/home.js");
const _sfc_main = {
  __name: "index",
  setup(__props) {
    const currentTime = common_vendor.ref("");
    const userLevel = common_vendor.ref(1);
    const streakDays = common_vendor.ref(0);
    const hasCheckedToday = common_vendor.ref(false);
    const currentExp = common_vendor.ref(0);
    const nextLevelExp = common_vendor.ref(100);
    const todayAchievements = common_vendor.ref([]);
    const recentAchievements = common_vendor.ref([]);
    const userStats = common_vendor.reactive({
      level: 1,
      experience: 0,
      currentStreak: 0,
      longestStreak: 0,
      totalDays: 0,
      successRate: 0
    });
    const gameStats = common_vendor.reactive({
      currentLevel: 1,
      currentExp: 0,
      nextLevelExp: 100,
      levelProgress: 0,
      totalAchievements: 0,
      unlockedAchievements: 0,
      achievementRate: 0,
      recentAchievements: []
    });
    const greeting = common_vendor.computed(() => {
      const hour = (/* @__PURE__ */ new Date()).getHours();
      if (hour < 12) return "早上好";
      if (hour < 18) return "下午好";
      return "晚上好";
    });
    const levelTitle = common_vendor.computed(() => {
      const level = userLevel.value;
      if (level >= 50) return "传奇大师";
      if (level >= 40) return "宗师";
      if (level >= 30) return "专家";
      if (level >= 25) return "资深导师";
      if (level >= 20) return "高级导师";
      if (level >= 15) return "导师";
      if (level >= 10) return "进阶行者";
      if (level >= 5) return "初级学徒";
      if (level >= 1) return "新手";
      return "未知";
    });
    const levelProgress = common_vendor.computed(() => {
      if (nextLevelExp.value <= 0) return 0;
      return Math.round(currentExp.value / nextLevelExp.value * 100);
    });
    const daysToMilestone = common_vendor.computed(() => {
      const milestones = [7, 14, 21, 30, 60, 90, 180, 365];
      const current = streakDays.value;
      for (let milestone of milestones) {
        if (current < milestone) {
          return milestone - current;
        }
      }
      return 30;
    });
    const milestoneProgress = common_vendor.computed(() => {
      const milestones = [7, 14, 21, 30, 60, 90, 180, 365];
      const current = streakDays.value;
      let previousMilestone = 0;
      for (let milestone of milestones) {
        if (current < milestone) {
          const progress = (current - previousMilestone) / (milestone - previousMilestone) * 100;
          return Math.max(0, Math.min(100, progress));
        }
        previousMilestone = milestone;
      }
      return 100;
    });
    const dailyTip = common_vendor.ref("当感到冲动时，尝试做10个深呼吸或者快速做20个俯卧撑。身体运动能有效转移注意力并释放内啡肽。");
    const updateTime = () => {
      const now = /* @__PURE__ */ new Date();
      const hours = now.getHours().toString().padStart(2, "0");
      const minutes = now.getMinutes().toString().padStart(2, "0");
      currentTime.value = `${hours}:${minutes}`;
    };
    const loadUserData = () => __async(this, null, function* () {
      try {
        const token = common_vendor.index.getStorageSync("token");
        if (!token) {
          console.log("用户未登录");
          return;
        }
        const result = yield apis_home.homeApi.getHomeData();
        if (result.code === 0) {
          const { userStats: stats, todayStatus, gameStats: game } = result.data;
          if (stats) {
            Object.assign(userStats, stats);
            streakDays.value = stats.currentStreak || 0;
            userLevel.value = stats.level || 1;
          }
          if (todayStatus) {
            hasCheckedToday.value = todayStatus.hasChecked || false;
          }
          if (game) {
            Object.assign(gameStats, game);
            userLevel.value = game.currentLevel || 1;
            currentExp.value = game.currentExp || 0;
            nextLevelExp.value = game.nextLevelExp || 100;
            recentAchievements.value = game.recentAchievements || [];
          }
        } else {
          console.error("获取首页数据失败:", result.message);
          common_vendor.index.showToast({
            title: "数据加载失败",
            icon: "none"
          });
        }
      } catch (error) {
        console.error("加载用户数据失败:", error);
        common_vendor.index.showToast({
          title: "网络错误，请重试",
          icon: "none"
        });
      }
    });
    const formatTime = (dateStr) => {
      if (!dateStr) return "";
      const date = new Date(dateStr);
      const now = /* @__PURE__ */ new Date();
      const diffTime = now - date;
      const diffDays = Math.floor(diffTime / (1e3 * 60 * 60 * 24));
      if (diffDays === 0) return "今天获得";
      if (diffDays === 1) return "昨天获得";
      if (diffDays < 7) return `${diffDays}天前获得`;
      return `${Math.floor(diffDays / 7)}周前获得`;
    };
    const handleCheckin = () => {
      if (hasCheckedToday.value) {
        goToCheckin();
      } else {
        goToCheckin();
      }
    };
    const goToCheckin = () => {
      common_vendor.index.switchTab({ url: "/pages/checkin/index" });
    };
    const goToAchievement = () => {
      common_vendor.index.navigateTo({ url: "/pages/achievement/index" });
    };
    const goToCommunity = () => {
      common_vendor.index.switchTab({ url: "/pages/community/index" });
    };
    const goToEmergency = () => {
      common_vendor.index.navigateTo({ url: "/pages/emergency/index" });
    };
    const goToProfile = () => {
      common_vendor.index.switchTab({ url: "/pages/profile/index" });
    };
    const goToNotifications = () => {
      common_vendor.index.showToast({
        title: "通知功能开发中",
        icon: "none"
      });
    };
    common_vendor.onMounted(() => {
      updateTime();
      setInterval(updateTime, 6e4);
      loadUserData();
    });
    common_vendor.onShow(() => {
      loadUserData();
    });
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.t(currentTime.value),
        b: common_vendor.t(greeting.value),
        c: common_vendor.o(goToNotifications),
        d: common_vendor.t(userLevel.value),
        e: common_vendor.o(goToProfile),
        f: common_vendor.t(streakDays.value),
        g: common_vendor.t(daysToMilestone.value),
        h: milestoneProgress.value + "%",
        i: common_vendor.t(hasCheckedToday.value ? "今日已打卡 ✓" : "今日打卡 ✓"),
        j: hasCheckedToday.value ? 1 : "",
        k: common_vendor.o(handleCheckin),
        l: common_vendor.o(goToCheckin),
        m: common_vendor.t(userLevel.value),
        n: common_vendor.t(levelTitle.value),
        o: common_vendor.t(currentExp.value),
        p: common_vendor.t(nextLevelExp.value),
        q: levelProgress.value + "%",
        r: common_vendor.o(goToAchievement),
        s: common_vendor.o(goToEmergency),
        t: common_vendor.o(goToCommunity),
        v: todayAchievements.value.length > 0
      }, todayAchievements.value.length > 0 ? {
        w: common_vendor.f(todayAchievements.value, (achievement, k0, i0) => {
          return {
            a: common_vendor.t(achievement.name),
            b: common_vendor.t(achievement.rewards),
            c: common_vendor.t(achievement.rewards),
            d: achievement.id
          };
        })
      } : {}, {
        x: common_vendor.t(dailyTip.value),
        y: recentAchievements.value.length > 0
      }, recentAchievements.value.length > 0 ? {
        z: common_vendor.f(recentAchievements.value, (achievement, k0, i0) => {
          return {
            a: common_vendor.t(achievement.name),
            b: common_vendor.t(formatTime(achievement.unlockedAt)),
            c: achievement.id
          };
        })
      } : {});
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-83a5a03c"]]);
wx.createPage(MiniProgramPage);
