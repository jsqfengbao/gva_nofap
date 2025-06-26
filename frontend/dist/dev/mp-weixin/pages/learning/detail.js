"use strict";
var __defProp = Object.defineProperty;
var __defProps = Object.defineProperties;
var __getOwnPropDescs = Object.getOwnPropertyDescriptors;
var __getOwnPropSymbols = Object.getOwnPropertySymbols;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __propIsEnum = Object.prototype.propertyIsEnumerable;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp.call(b, prop))
      __defNormalProp(a, prop, b[prop]);
  if (__getOwnPropSymbols)
    for (var prop of __getOwnPropSymbols(b)) {
      if (__propIsEnum.call(b, prop))
        __defNormalProp(a, prop, b[prop]);
    }
  return a;
};
var __spreadProps = (a, b) => __defProps(a, __getOwnPropDescs(b));
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
  name: "LearningDetail",
  data() {
    return {
      currentTime: "9:41",
      contentId: null,
      content: {},
      learningRecord: null,
      relatedContents: [],
      // 音频播放状态
      isAudioPlaying: false,
      audioCurrentTime: 0,
      audioDuration: 0,
      audioProgress: 0,
      audioContext: null,
      // 评分
      showRatingModal: false,
      currentRating: 0
    };
  },
  onLoad(options) {
    if (options.id) {
      this.contentId = parseInt(options.id);
      this.initPage();
    }
  },
  onShow() {
    this.updateTime();
  },
  onUnload() {
    if (this.audioContext) {
      this.audioContext.stop();
    }
  },
  methods: {
    initPage() {
      return __async(this, null, function* () {
        common_vendor.index.showLoading({ title: "加载中..." });
        try {
          yield Promise.all([
            this.loadContent(),
            this.loadLearningRecord(),
            this.loadRelatedContent()
          ]);
          yield this.startLearningRecord();
        } catch (error) {
          console.error("页面初始化失败:", error);
          common_vendor.index.showToast({
            title: "加载失败",
            icon: "error"
          });
        } finally {
          common_vendor.index.hideLoading();
        }
      });
    },
    // 加载内容详情
    loadContent() {
      return __async(this, null, function* () {
        try {
          const res = yield common_vendor.index.request({
            url: `/api/v1/miniprogram/learning/contents/${this.contentId}`,
            method: "GET",
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.content = res.data.data;
          }
        } catch (error) {
          console.error("加载内容失败:", error);
          throw error;
        }
      });
    },
    // 加载学习记录
    loadLearningRecord() {
      return __async(this, null, function* () {
        try {
          const res = yield common_vendor.index.request({
            url: `/api/v1/miniprogram/learning/record/${this.contentId}`,
            method: "GET",
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.learningRecord = res.data.data;
          }
        } catch (error) {
          console.error("加载学习记录失败:", error);
        }
      });
    },
    // 加载相关内容
    loadRelatedContent() {
      return __async(this, null, function* () {
        try {
          const res = yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/recommendations",
            method: "GET",
            data: {
              contentId: this.contentId,
              type: "similar",
              limit: 3
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.relatedContents = res.data.data.contents || [];
          }
        } catch (error) {
          console.error("加载相关内容失败:", error);
        }
      });
    },
    // 开始学习记录
    startLearningRecord() {
      return __async(this, null, function* () {
        try {
          yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/start",
            method: "POST",
            data: {
              contentId: this.contentId
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
        } catch (error) {
          console.error("开始学习记录失败:", error);
        }
      });
    },
    // 返回上一页
    goBack() {
      common_vendor.index.navigateBack();
    },
    // 切换收藏状态
    toggleCollect() {
      return __async(this, null, function* () {
        try {
          const res = yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/collect",
            method: "POST",
            data: {
              contentId: this.contentId
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.content.isCollected = !this.content.isCollected;
            common_vendor.index.showToast({
              title: this.content.isCollected ? "收藏成功" : "取消收藏",
              icon: "success"
            });
          }
        } catch (error) {
          console.error("收藏操作失败:", error);
          common_vendor.index.showToast({
            title: "操作失败",
            icon: "error"
          });
        }
      });
    },
    // 切换点赞状态
    toggleLike() {
      return __async(this, null, function* () {
        try {
          const res = yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/like",
            method: "POST",
            data: {
              contentId: this.contentId
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.content.isLiked = !this.content.isLiked;
            this.content.likeCount += this.content.isLiked ? 1 : -1;
            common_vendor.index.showToast({
              title: this.content.isLiked ? "点赞成功" : "取消点赞",
              icon: "success"
            });
          }
        } catch (error) {
          console.error("点赞操作失败:", error);
          common_vendor.index.showToast({
            title: "操作失败",
            icon: "error"
          });
        }
      });
    },
    // 分享内容
    shareContent() {
      common_vendor.index.share({
        provider: "weixin",
        type: 0,
        title: this.content.title,
        summary: this.content.summary,
        imageUrl: this.content.thumbnailUrl,
        success: () => {
          common_vendor.index.showToast({
            title: "分享成功",
            icon: "success"
          });
        }
      });
    },
    // 标记为已完成
    markAsCompleted() {
      return __async(this, null, function* () {
        try {
          const res = yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/complete",
            method: "POST",
            data: {
              contentId: this.contentId,
              progress: 100
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.learningRecord = __spreadProps(__spreadValues({}, this.learningRecord), {
              isCompleted: true,
              progress: 100,
              completedAt: (/* @__PURE__ */ new Date()).toISOString()
            });
            common_vendor.index.showToast({
              title: "恭喜完成学习！",
              icon: "success"
            });
          }
        } catch (error) {
          console.error("完成学习失败:", error);
          common_vendor.index.showToast({
            title: "操作失败",
            icon: "error"
          });
        }
      });
    },
    // 重新学习
    restartLearning() {
      return __async(this, null, function* () {
        try {
          yield this.startLearningRecord();
          this.learningRecord.isCompleted = false;
          this.learningRecord.progress = 0;
          common_vendor.index.showToast({
            title: "开始重新学习",
            icon: "success"
          });
        } catch (error) {
          console.error("重新学习失败:", error);
        }
      });
    },
    // 音频播放控制
    toggleAudioPlay() {
      if (this.isAudioPlaying) {
        this.pauseAudio();
      } else {
        this.playAudio();
      }
    },
    playAudio() {
      if (!this.audioContext) {
        this.audioContext = common_vendor.index.createInnerAudioContext();
        this.audioContext.src = this.content.audioUrl;
        this.audioContext.onPlay(() => {
          this.isAudioPlaying = true;
        });
        this.audioContext.onPause(() => {
          this.isAudioPlaying = false;
        });
        this.audioContext.onTimeUpdate(() => {
          this.audioCurrentTime = this.audioContext.currentTime;
          this.audioDuration = this.audioContext.duration;
          this.audioProgress = this.audioCurrentTime / this.audioDuration * 100;
        });
        this.audioContext.onEnded(() => {
          this.isAudioPlaying = false;
          this.markAsCompleted();
        });
      }
      this.audioContext.play();
    },
    pauseAudio() {
      if (this.audioContext) {
        this.audioContext.pause();
      }
    },
    // 视频播放事件
    onVideoPlay() {
      console.log("视频开始播放");
    },
    onVideoPause() {
      console.log("视频暂停");
    },
    onVideoTimeUpdate(e) {
      const { currentTime, duration } = e.detail;
      if (duration > 0) {
        const progress = currentTime / duration * 100;
        this.updateLearningProgress(progress);
      }
    },
    // 更新学习进度
    updateLearningProgress(progress) {
      return __async(this, null, function* () {
        try {
          yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/progress",
            method: "POST",
            data: {
              contentId: this.contentId,
              progress
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (this.learningRecord) {
            this.learningRecord.progress = progress;
          }
        } catch (error) {
          console.error("更新学习进度失败:", error);
        }
      });
    },
    // 显示评分弹窗
    rateContent() {
      this.showRatingModal = true;
      this.currentRating = 0;
    },
    // 隐藏评分弹窗
    hideRating() {
      this.showRatingModal = false;
    },
    // 设置评分
    setRating(rating) {
      this.currentRating = rating;
    },
    // 提交评分
    submitRating() {
      return __async(this, null, function* () {
        if (this.currentRating === 0) {
          common_vendor.index.showToast({
            title: "请选择评分",
            icon: "none"
          });
          return;
        }
        try {
          const res = yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/rate",
            method: "POST",
            data: {
              contentId: this.contentId,
              rating: this.currentRating
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.hideRating();
            common_vendor.index.showToast({
              title: "评分成功",
              icon: "success"
            });
          }
        } catch (error) {
          console.error("评分失败:", error);
          common_vendor.index.showToast({
            title: "评分失败",
            icon: "error"
          });
        }
      });
    },
    // 查看相关内容
    viewRelated(item) {
      common_vendor.index.redirectTo({
        url: `/pages/learning/detail?id=${item.id}`
      });
    },
    // 获取分类名称
    getCategoryName(category) {
      const categoryMap = {
        1: "科普知识",
        2: "康复指导",
        3: "心理健康",
        4: "经验分享"
      };
      return categoryMap[category] || "其他";
    },
    // 格式化数字
    formatNumber(num) {
      if (num >= 1e3) {
        return (num / 1e3).toFixed(1) + "k";
      }
      return num.toString();
    },
    // 格式化日期
    formatDate(dateStr) {
      const date = new Date(dateStr);
      return `${date.getMonth() + 1}月${date.getDate()}日`;
    },
    // 格式化日期时间
    formatDateTime(dateStr) {
      const date = new Date(dateStr);
      return `${date.getMonth() + 1}月${date.getDate()}日 ${date.getHours()}:${date.getMinutes().toString().padStart(2, "0")}`;
    },
    // 格式化时间
    formatTime(seconds) {
      const mins = Math.floor(seconds / 60);
      const secs = Math.floor(seconds % 60);
      return `${mins}:${secs.toString().padStart(2, "0")}`;
    },
    // 更新时间
    updateTime() {
      const now = /* @__PURE__ */ new Date();
      this.currentTime = `${now.getHours()}:${now.getMinutes().toString().padStart(2, "0")}`;
    }
  }
};
if (!Array) {
  const _component_nf_button = common_vendor.resolveComponent("nf-button");
  _component_nf_button();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.t($data.currentTime),
    b: common_vendor.o((...args) => $options.goBack && $options.goBack(...args)),
    c: common_vendor.t($data.content.isCollected ? "★" : "☆"),
    d: common_vendor.o((...args) => $options.toggleCollect && $options.toggleCollect(...args)),
    e: common_vendor.o((...args) => $options.shareContent && $options.shareContent(...args)),
    f: common_vendor.t($options.getCategoryName($data.content.category)),
    g: common_vendor.n("category-" + $data.content.category),
    h: common_vendor.t($data.content.title),
    i: common_vendor.t($data.content.duration),
    j: common_vendor.t($options.formatNumber($data.content.viewCount)),
    k: common_vendor.t($data.content.likeCount),
    l: common_vendor.t($options.formatDate($data.content.createdAt)),
    m: common_vendor.t($data.content.summary),
    n: $data.content.contentType === 2
  }, $data.content.contentType === 2 ? {
    o: $data.content.videoUrl,
    p: $data.content.thumbnailUrl,
    q: common_vendor.o((...args) => $options.onVideoPlay && $options.onVideoPlay(...args)),
    r: common_vendor.o((...args) => $options.onVideoPause && $options.onVideoPause(...args)),
    s: common_vendor.o((...args) => $options.onVideoTimeUpdate && $options.onVideoTimeUpdate(...args))
  } : {}, {
    t: $data.content.contentType === 3
  }, $data.content.contentType === 3 ? {
    v: $data.content.thumbnailUrl || "/static/images/audio-placeholder.png",
    w: common_vendor.t($data.isAudioPlaying ? "⏸️" : "▶️"),
    x: common_vendor.o((...args) => $options.toggleAudioPlay && $options.toggleAudioPlay(...args)),
    y: common_vendor.t($data.content.title),
    z: common_vendor.t($options.formatTime($data.audioCurrentTime)),
    A: $data.audioProgress + "%",
    B: common_vendor.o((...args) => _ctx.seekAudio && _ctx.seekAudio(...args)),
    C: common_vendor.t($options.formatTime($data.audioDuration))
  } : {}, {
    D: $data.content.contentType === 1
  }, $data.content.contentType === 1 ? {
    E: $data.content.content
  } : {}, {
    F: $data.content.contentType !== 1 && $data.content.textContent
  }, $data.content.contentType !== 1 && $data.content.textContent ? {
    G: $data.content.textContent
  } : {}, {
    H: $data.learningRecord
  }, $data.learningRecord ? common_vendor.e({
    I: common_vendor.t(Math.round($data.learningRecord.progress)),
    J: $data.learningRecord.progress + "%",
    K: common_vendor.t($options.formatDateTime($data.learningRecord.startTime)),
    L: $data.learningRecord.completedAt
  }, $data.learningRecord.completedAt ? {
    M: common_vendor.t($options.formatDateTime($data.learningRecord.completedAt))
  } : {}, {
    N: common_vendor.t(Math.round($data.learningRecord.learningTime / 60))
  }) : {}, {
    O: common_vendor.t($data.content.isLiked ? "❤️" : "🤍"),
    P: $data.content.isLiked ? 1 : "",
    Q: common_vendor.t($data.content.isLiked ? "已点赞" : "点赞"),
    R: common_vendor.o((...args) => $options.toggleLike && $options.toggleLike(...args)),
    S: common_vendor.t($data.content.isCollected ? "★" : "☆"),
    T: $data.content.isCollected ? 1 : "",
    U: common_vendor.t($data.content.isCollected ? "已收藏" : "收藏"),
    V: common_vendor.o((...args) => $options.toggleCollect && $options.toggleCollect(...args)),
    W: common_vendor.o((...args) => $options.shareContent && $options.shareContent(...args)),
    X: common_vendor.o((...args) => $options.rateContent && $options.rateContent(...args)),
    Y: $data.relatedContents.length > 0
  }, $data.relatedContents.length > 0 ? {
    Z: common_vendor.f($data.relatedContents, (item, k0, i0) => {
      return {
        a: item.thumbnailUrl,
        b: common_vendor.t(item.title),
        c: common_vendor.t($options.getCategoryName(item.category)),
        d: common_vendor.t(item.duration),
        e: item.id,
        f: common_vendor.o(($event) => $options.viewRelated(item), item.id)
      };
    })
  } : {}, {
    aa: !$data.learningRecord || !$data.learningRecord.isCompleted
  }, !$data.learningRecord || !$data.learningRecord.isCompleted ? {
    ab: common_vendor.o($options.markAsCompleted),
    ac: common_vendor.p({
      type: "primary"
    })
  } : {
    ad: common_vendor.o($options.restartLearning),
    ae: common_vendor.p({
      type: "secondary"
    })
  }, {
    af: $data.showRatingModal
  }, $data.showRatingModal ? {
    ag: common_vendor.f(5, (star, k0, i0) => {
      return {
        a: star,
        b: star <= $data.currentRating ? 1 : "",
        c: common_vendor.o(($event) => $options.setRating(star), star)
      };
    }),
    ah: common_vendor.o($options.hideRating),
    ai: common_vendor.p({
      type: "secondary"
    }),
    aj: common_vendor.o($options.submitRating),
    ak: common_vendor.p({
      type: "primary"
    }),
    al: common_vendor.o(() => {
    }),
    am: common_vendor.o((...args) => $options.hideRating && $options.hideRating(...args))
  } : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render]]);
wx.createPage(MiniProgramPage);
