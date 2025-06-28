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
function getHomeData() {
  return __async(this, null, function* () {
    var _a, _b, _c;
    try {
      const [statsRes, todayRes, gameRes] = yield Promise.all([
        apis_checkin.checkinApi.getStatistics(),
        apis_checkin.checkinApi.getTodayStatus(),
        apis_achievement.achievementApi.getGameStats()
      ]);
      return {
        code: 0,
        data: {
          userStats: ((_a = statsRes.data) == null ? void 0 : _a.data) || {},
          todayStatus: ((_b = todayRes.data) == null ? void 0 : _b.data) || {},
          gameStats: ((_c = gameRes.data) == null ? void 0 : _c.data) || {}
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
  });
}
function getCheckinStats() {
  return apis_checkin.checkinApi.getStatistics();
}
function getTodayCheckinStatus() {
  return apis_checkin.checkinApi.getTodayStatus();
}
function getGameStats() {
  return apis_achievement.achievementApi.getGameStats();
}
const homeApi = {
  getHomeData,
  getCheckinStats,
  getTodayCheckinStatus,
  getGameStats
};
exports.homeApi = homeApi;
