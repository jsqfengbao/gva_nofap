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
if (!Array) {
  const _component_NfNavBar = common_vendor.resolveComponent("NfNavBar");
  _component_NfNavBar();
}
const _sfc_main = {
  __name: "index",
  setup(__props) {
    const loading = common_vendor.ref(false);
    const activeCategory = common_vendor.ref(1);
    const showModal = common_vendor.ref(false);
    const selectedAchievement = common_vendor.ref(null);
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
    const achievements = common_vendor.reactive({
      totalAchievements: 0,
      unlockedAchievements: 0,
      unlockRate: 0,
      categories: {
        1: [],
        // 打卡类
        2: [],
        // 等级类
        3: [],
        // 社区类
        4: [],
        // 学习类
        5: []
        // 特殊类
      },
      recentUnlocked: []
    });
    const categoryNames = {
      1: "打卡类",
      2: "等级类",
      3: "社区类",
      4: "学习类",
      5: "特殊类"
    };
    const rarityNames = {
      1: "普通",
      2: "稀有",
      3: "史诗",
      4: "传说"
    };
    const gamificationService = common_vendor.computed(() => {
      return {
        GetLevelTitle: (level) => {
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
        }
      };
    });
    const switchCategory = (category) => {
      activeCategory.value = category;
    };
    const getRarityClass = (rarity) => {
      const classes = {
        1: "common",
        2: "rare",
        3: "epic",
        4: "legendary"
      };
      return classes[rarity] || "common";
    };
    const getRarityText = (rarity) => {
      return rarityNames[rarity] || "普通";
    };
    const getCategoryText = (category) => {
      return categoryNames[category] || "未知";
    };
    const formatDate = (dateStr) => {
      if (!dateStr) return "";
      const date = new Date(dateStr);
      return `${date.getMonth() + 1}月${date.getDate()}日`;
    };
    const formatDateTime = (dateStr) => {
      if (!dateStr) return "";
      const date = new Date(dateStr);
      return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, "0")}-${date.getDate().toString().padStart(2, "0")} ${date.getHours().toString().padStart(2, "0")}:${date.getMinutes().toString().padStart(2, "0")}`;
    };
    const showAchievementDetail = (achievement) => {
      selectedAchievement.value = achievement;
      showModal.value = true;
    };
    const closeModal = () => {
      showModal.value = false;
      selectedAchievement.value = null;
    };
    const loadGameStats = () => __async(this, null, function* () {
      try {
        const token = common_vendor.index.getStorageSync("token");
        if (!token) {
          common_vendor.index.redirectTo({ url: "/pages/auth/login" });
          return;
        }
        const res = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/achievement/game-stats",
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (res.data.code === 0) {
          Object.assign(gameStats, res.data.data);
        }
      } catch (error) {
        console.error("获取游戏统计失败:", error);
      }
    });
    const loadAchievements = () => __async(this, null, function* () {
      try {
        const token = common_vendor.index.getStorageSync("token");
        if (!token) {
          common_vendor.index.redirectTo({ url: "/pages/auth/login" });
          return;
        }
        const res = yield common_vendor.index.request({
          url: "http://localhost:8888/api/v1/miniprogram/achievement/list",
          method: "GET",
          header: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });
        if (res.data.code === 0) {
          Object.assign(achievements, res.data.data);
        }
      } catch (error) {
        console.error("获取成就列表失败:", error);
      }
    });
    const loadData = () => __async(this, null, function* () {
      loading.value = true;
      try {
        yield Promise.all([
          loadGameStats(),
          loadAchievements()
        ]);
      } finally {
        loading.value = false;
      }
    });
    common_vendor.onMounted(() => {
      loadData();
    });
    return (_ctx, _cache) => {
      var _a, _b, _c, _d, _e, _f, _g, _h, _i, _j, _k;
      return common_vendor.e({
        a: common_vendor.p({
          title: "成就系统",
          ["show-back"]: true
        }),
        b: common_vendor.t(gamificationService.value.GetLevelTitle(gameStats.currentLevel)),
        c: common_vendor.t(gameStats.currentLevel),
        d: gameStats.levelProgress + "%",
        e: common_vendor.t(gameStats.currentExp),
        f: common_vendor.t(gameStats.nextLevelExp),
        g: common_vendor.t(gameStats.unlockedAchievements),
        h: common_vendor.t(gameStats.totalAchievements),
        i: common_vendor.t(Math.round(gameStats.achievementRate)),
        j: gameStats.recentAchievements && gameStats.recentAchievements.length > 0
      }, gameStats.recentAchievements && gameStats.recentAchievements.length > 0 ? {
        k: common_vendor.f(gameStats.recentAchievements, (achievement, k0, i0) => {
          return {
            a: common_vendor.t(getRarityText(achievement.rarity)),
            b: common_vendor.n(getRarityClass(achievement.rarity)),
            c: common_vendor.t(achievement.name),
            d: common_vendor.t(formatDate(achievement.unlockedAt)),
            e: achievement.id
          };
        })
      } : {}, {
        l: common_vendor.f(categoryNames, (categoryName, category, i0) => {
          return {
            a: common_vendor.t(categoryName),
            b: activeCategory.value === category ? 1 : "",
            c: category,
            d: common_vendor.o(($event) => switchCategory(category), category)
          };
        }),
        m: achievements.categories[activeCategory.value]
      }, achievements.categories[activeCategory.value] ? {
        n: common_vendor.f(achievements.categories[activeCategory.value], (achievement, k0, i0) => {
          return common_vendor.e({
            a: achievement.isUnlocked
          }, achievement.isUnlocked ? {} : {}, {
            b: common_vendor.n(getRarityClass(achievement.rarity)),
            c: common_vendor.t(achievement.name),
            d: common_vendor.t(achievement.description),
            e: common_vendor.t(achievement.rewards),
            f: common_vendor.t(getRarityText(achievement.rarity)),
            g: achievement.isUnlocked
          }, achievement.isUnlocked ? {} : {}, {
            h: achievement.isUnlocked ? 1 : "",
            i: achievement.id,
            j: common_vendor.o(($event) => showAchievementDetail(achievement), achievement.id)
          });
        })
      } : {}, {
        o: (_a = selectedAchievement.value) == null ? void 0 : _a.isUnlocked
      }, ((_b = selectedAchievement.value) == null ? void 0 : _b.isUnlocked) ? {} : {}, {
        p: common_vendor.t((_c = selectedAchievement.value) == null ? void 0 : _c.name),
        q: common_vendor.t(getRarityText((_d = selectedAchievement.value) == null ? void 0 : _d.rarity)),
        r: common_vendor.n(getRarityClass((_e = selectedAchievement.value) == null ? void 0 : _e.rarity)),
        s: common_vendor.t((_f = selectedAchievement.value) == null ? void 0 : _f.description),
        t: common_vendor.t(getCategoryText((_g = selectedAchievement.value) == null ? void 0 : _g.category)),
        v: common_vendor.t((_h = selectedAchievement.value) == null ? void 0 : _h.rewards),
        w: (_i = selectedAchievement.value) == null ? void 0 : _i.isUnlocked
      }, ((_j = selectedAchievement.value) == null ? void 0 : _j.isUnlocked) ? {
        x: common_vendor.t(formatDateTime((_k = selectedAchievement.value) == null ? void 0 : _k.unlockedAt))
      } : {}, {
        y: common_vendor.o(closeModal),
        z: common_vendor.o(() => {
        }),
        A: showModal.value,
        B: common_vendor.o(closeModal)
      });
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-6ea57778"]]);
wx.createPage(MiniProgramPage);
