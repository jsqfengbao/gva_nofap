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
const apis_learning = require("./learning.js");
const data_learningMock = require("../data/learning-mock.js");
class LearningService {
  /**
   * 获取学习统计数据
   */
  getStats() {
    return __async(this, null, function* () {
      try {
        const res = yield apis_learning.learningApi.getStats();
        if (res.code === 0) {
          return {
            success: true,
            data: res.data
          };
        }
        throw new Error("API返回错误");
      } catch (error) {
        console.warn("获取学习统计失败，使用模拟数据:", error);
        return {
          success: true,
          data: data_learningMock.mockLearningStats,
          isMock: true
        };
      }
    });
  }
  /**
   * 获取学习内容列表
   * 注意：后端API暂未实现，直接使用模拟数据
   */
  getContents() {
    return __async(this, arguments, function* (params = {}) {
      console.warn("学习内容列表API暂未实现，使用模拟数据");
      const { page = 1, pageSize = 10, contentType } = params;
      const mockContents = data_learningMock.getContentsByType(contentType);
      const paginatedData = data_learningMock.getPaginatedContents(mockContents, page, pageSize);
      return {
        success: true,
        data: paginatedData,
        isMock: true
      };
    });
  }
  /**
   * 搜索学习内容
   * 注意：后端API暂未实现，直接使用模拟数据
   */
  searchContents() {
    return __async(this, arguments, function* (params = {}) {
      console.warn("搜索学习内容API暂未实现，使用模拟数据");
      const { keyword, page = 1, pageSize = 10 } = params;
      const searchResults = data_learningMock.searchMockContents(keyword || "");
      const paginatedData = data_learningMock.getPaginatedContents(searchResults, page, pageSize);
      return {
        success: true,
        data: paginatedData,
        isMock: true
      };
    });
  }
  /**
   * 开始学习记录
   * 注意：后端API暂未实现，返回模拟成功
   */
  startLearning(contentId) {
    return __async(this, null, function* () {
      console.warn("开始学习记录API暂未实现，返回模拟成功");
      return {
        success: true,
        data: {
          message: "学习记录已开始",
          contentId,
          startTime: (/* @__PURE__ */ new Date()).toISOString()
        },
        isMock: true
      };
    });
  }
  /**
   * 完成学习
   */
  completeLearning(_0) {
    return __async(this, arguments, function* (contentId, data = {}) {
      try {
        const res = yield apis_learning.learningApi.completeLearning(contentId, data);
        return {
          success: res.code === 0,
          data: res.data
        };
      } catch (error) {
        console.warn("完成学习记录失败:", error);
        return {
          success: true,
          data: { message: "学习记录失败，但不影响使用" },
          isMock: true
        };
      }
    });
  }
  /**
   * 点赞内容
   */
  likeContent(contentId) {
    return __async(this, null, function* () {
      try {
        const res = yield apis_learning.learningApi.likeContent(contentId);
        return {
          success: res.code === 0,
          data: res.data
        };
      } catch (error) {
        console.warn("点赞失败:", error);
        return {
          success: false,
          error: "点赞失败，请稍后重试"
        };
      }
    });
  }
  /**
   * 收藏内容
   */
  collectContent(contentId) {
    return __async(this, null, function* () {
      try {
        const res = yield apis_learning.learningApi.collectContent(contentId);
        return {
          success: res.code === 0,
          data: res.data
        };
      } catch (error) {
        console.warn("收藏失败:", error);
        return {
          success: false,
          error: "收藏失败，请稍后重试"
        };
      }
    });
  }
  /**
   * 获取内容详情
   */
  getContentDetail(id) {
    return __async(this, null, function* () {
      try {
        const res = yield apis_learning.learningApi.getDetail(id);
        if (res.code === 0) {
          return {
            success: true,
            data: res.data
          };
        }
        throw new Error("API返回错误");
      } catch (error) {
        console.warn("获取内容详情失败，使用模拟数据:", error);
        const allContents = data_learningMock.getAllMockContents();
        const content = allContents.find((item) => item.id === parseInt(id));
        if (content) {
          return {
            success: true,
            data: content,
            isMock: true
          };
        } else {
          return {
            success: false,
            error: "内容不存在"
          };
        }
      }
    });
  }
  /**
   * 转换学习统计数据格式
   */
  transformStatsForDisplay(stats) {
    return {
      weeklyHours: (stats.totalLearningTime / 60).toFixed(1),
      weeklyGoal: 5,
      weeklyProgress: Math.min(stats.totalLearningTime / 300 * 100, 100),
      articlesRead: stats.completedContents || 0,
      videosWatched: stats.likedContents || 0,
      audiosListened: stats.collectedContents || 0
    };
  }
  /**
   * 根据内容类型获取对应的模拟数据
   */
  getMockDataByType(contentType) {
    switch (contentType) {
      case 1:
        return data_learningMock.mockArticles;
      case 2:
        return data_learningMock.mockVideos;
      case 3:
        return data_learningMock.mockAudios;
      default:
        return data_learningMock.getAllMockContents();
    }
  }
}
const learningService = new LearningService();
exports.learningService = learningService;
