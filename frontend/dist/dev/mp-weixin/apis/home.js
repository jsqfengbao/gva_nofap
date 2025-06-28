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
const apis_checkin = require("./checkin.js");
const apis_achievement = require("./achievement.js");
const homeApi = {
  // 获取首页所有数据
  getHomeData: () => __async(exports, null, function* () {
    try {
      const [statsRes, todayRes, gameRes] = yield Promise.all([
        apis_checkin.checkinApi.getStatistics(),
        apis_checkin.checkinApi.getTodayStatus(),
        apis_achievement.achievementApi.getGameStats()
      ]);
      return {
        code: 0,
        data: {
          userStats: statsRes.data || {},
          todayStatus: todayRes.data || {},
          gameStats: gameRes.data || {}
        }
      };
    } catch (error) {
      console.error("获取首页数据失败:", error);
      return {
        code: -1,
        message: "获取首页数据失败",
        error
      };
    }
  }),
  // 获取打卡统计数据
  getCheckinStats: () => {
    return apis_checkin.checkinApi.getStatistics();
  },
  // 获取今日打卡状态
  getTodayCheckinStatus: () => {
    return apis_checkin.checkinApi.getTodayStatus();
  },
  // 获取游戏化统计数据
  getGameStats: () => {
    return apis_achievement.achievementApi.getGameStats();
  }
};
exports.homeApi = homeApi;
