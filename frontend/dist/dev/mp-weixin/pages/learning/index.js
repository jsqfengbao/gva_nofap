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
const utils_format = require("../../utils/format.js");
const NfTabBar = () => "../../components/ui/navigation/NfTabBar.js";
const _sfc_main = {
  name: "LearningIndex",
  components: {
    NfTabBar
  },
  data() {
    return {
      currentTime: utils_format.getCurrentTime(),
      activeTabIndex: 0,
      categoryTabs: [
        { name: "推荐", type: "recommend" },
        { name: "文章", type: "article" },
        { name: "视频", type: "video" },
        { name: "音频", type: "audio" }
      ],
      // 精品课程
      featuredCourse: {
        id: 1,
        title: "21天自控力训练营",
        summary: "科学方法帮你建立持久的自控习惯",
        thumbnail: "https://images.unsplash.com/photo-1499336315816-097655dcfbda?w=160&h=160&fit=crop",
        category: 2,
        difficulty: 2
      },
      // 学习统计
      learningStats: {
        weeklyHours: 3.5,
        weeklyGoal: 5,
        weeklyProgress: 70,
        articlesRead: 12,
        videosWatched: 5,
        audiosListened: 8
      },
      // 内容数据
      articles: [],
      videos: [],
      audios: [],
      // 分页
      currentPage: 1,
      pageSize: 10,
      hasMore: true,
      loadingMore: false,
      // 搜索
      showSearchModal: false,
      searchKeyword: "",
      searchResults: []
    };
  },
  computed: {
    currentArticles() {
      return this.activeTabIndex === 0 || this.activeTabIndex === 1 ? this.articles.slice(0, this.activeTabIndex === 0 ? 2 : this.articles.length) : [];
    },
    currentVideos() {
      return this.activeTabIndex === 0 || this.activeTabIndex === 2 ? this.videos.slice(0, this.activeTabIndex === 0 ? 1 : this.videos.length) : [];
    },
    currentAudios() {
      return this.activeTabIndex === 0 || this.activeTabIndex === 3 ? this.audios.slice(0, this.activeTabIndex === 0 ? 1 : this.audios.length) : [];
    }
  },
  onLoad() {
    this.initPage();
  },
  onShow() {
    this.updateTime();
  },
  methods: {
    // 初始化页面
    initPage() {
      return __async(this, null, function* () {
        this.updateTime();
        this.loadMockData();
        try {
          yield this.loadLearningStats();
        } catch (error) {
          console.log("使用模拟数据");
        }
      });
    },
    // 加载模拟数据
    loadMockData() {
      this.articles = [
        {
          id: 1,
          title: "如何建立健康的生活习惯",
          summary: "从小习惯开始，逐步建立健康的生活方式...",
          thumbnailUrl: "https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=160&h=120&fit=crop",
          duration: 8,
          likeCount: 234,
          viewCount: 1200,
          contentType: 1,
          category: 2,
          createdAt: "2024-06-20T10:00:00Z"
        },
        {
          id: 2,
          title: "理解大脑的奖励机制",
          summary: "科学解释成瘾的神经学原理和应对方法...",
          thumbnailUrl: "https://images.unsplash.com/photo-1559757148-5c350d0d3c56?w=160&h=120&fit=crop",
          duration: 12,
          likeCount: 189,
          viewCount: 856,
          contentType: 1,
          category: 1,
          createdAt: "2024-06-19T14:30:00Z"
        },
        {
          id: 3,
          title: "正念冥想入门指南",
          summary: "学会用冥想管理情绪和冲动，建立内心平静...",
          thumbnailUrl: "https://images.unsplash.com/photo-1545389336-cf090694435e?w=160&h=120&fit=crop",
          duration: 15,
          likeCount: 156,
          viewCount: 743,
          contentType: 1,
          category: 3,
          createdAt: "2024-06-18T09:15:00Z"
        }
      ];
      this.videos = [
        {
          id: 4,
          title: "正念冥想入门指南",
          summary: "学会用冥想管理情绪和冲动",
          thumbnailUrl: "https://images.unsplash.com/photo-1545389336-cf090694435e?w=300&h=200&fit=crop",
          duration: 15,
          likeCount: 127,
          viewCount: 3200,
          commentCount: 23,
          contentType: 2,
          category: 3,
          createdAt: "2024-06-17T16:45:00Z",
          videoUrl: "https://example.com/video/meditation.mp4"
        },
        {
          id: 5,
          title: "呼吸练习技巧讲解",
          summary: "通过正确的呼吸方法缓解压力和焦虑",
          thumbnailUrl: "https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300&h=200&fit=crop",
          duration: 10,
          likeCount: 89,
          viewCount: 1800,
          commentCount: 15,
          contentType: 2,
          category: 2,
          createdAt: "2024-06-16T11:20:00Z",
          videoUrl: "https://example.com/video/breathing.mp4"
        }
      ];
      this.audios = [
        {
          id: 6,
          title: "睡前放松冥想",
          summary: "帮助改善睡眠质量的引导冥想",
          thumbnailUrl: "https://images.unsplash.com/photo-1542662565-7e4b16f20bfb?w=160&h=120&fit=crop",
          duration: 20,
          likeCount: 67,
          viewCount: 890,
          contentType: 3,
          category: 3,
          createdAt: "2024-06-15T20:00:00Z",
          audioUrl: "https://example.com/audio/sleep-meditation.mp3",
          isDownloaded: true
        },
        {
          id: 7,
          title: "专注力训练音频",
          summary: "提升注意力和专注能力的训练课程",
          thumbnailUrl: "https://images.unsplash.com/photo-1499209974431-9dddcece7f88?w=160&h=120&fit=crop",
          duration: 25,
          likeCount: 45,
          viewCount: 612,
          contentType: 3,
          category: 2,
          createdAt: "2024-06-14T08:30:00Z",
          audioUrl: "https://example.com/audio/focus-training.mp3",
          isDownloaded: false
        }
      ];
    },
    // 加载学习统计
    loadLearningStats() {
      return __async(this, null, function* () {
        try {
          const res = yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/stats",
            method: "GET",
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.learningStats = {
              weeklyHours: (res.data.data.totalLearningTime / 60).toFixed(1),
              weeklyGoal: 5,
              weeklyProgress: Math.min(res.data.data.totalLearningTime / 300 * 100, 100),
              articlesRead: res.data.data.completedContents || 0,
              videosWatched: res.data.data.likedContents || 0,
              audiosListened: res.data.data.collectedContents || 0
            };
          }
        } catch (error) {
          console.error("加载学习统计失败:", error);
        }
      });
    },
    // 加载内容列表
    loadContent(loadMore = false) {
      return __async(this, null, function* () {
        if (this.loadingMore) return;
        this.loadingMore = true;
        try {
          const contentType = this.getContentType();
          const res = yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/contents",
            method: "GET",
            data: {
              page: loadMore ? this.currentPage + 1 : 1,
              pageSize: this.pageSize,
              contentType: contentType || void 0,
              sortBy: "view_count",
              order: "desc"
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            const newContent = res.data.data.list || [];
            if (loadMore) {
              this.currentPage++;
              this.appendContent(newContent);
            } else {
              this.currentPage = 1;
              this.setContent(newContent);
            }
            this.hasMore = newContent.length === this.pageSize;
          }
        } catch (error) {
          console.error("加载内容失败:", error);
          common_vendor.index.showToast({
            title: "加载失败",
            icon: "error"
          });
        } finally {
          this.loadingMore = false;
        }
      });
    },
    // 获取当前内容类型
    getContentType() {
      const typeMap = {
        0: null,
        // 推荐
        1: 1,
        // 文章
        2: 2,
        // 视频
        3: 3
        // 音频
      };
      return typeMap[this.activeTabIndex];
    },
    // 设置内容数据
    setContent(content) {
      content.forEach((item) => {
        if (item.contentType === 1) {
          if (!this.articles.find((a) => a.id === item.id)) {
            this.articles.push(item);
          }
        } else if (item.contentType === 2) {
          if (!this.videos.find((v) => v.id === item.id)) {
            this.videos.push(item);
          }
        } else if (item.contentType === 3) {
          if (!this.audios.find((a) => a.id === item.id)) {
            this.audios.push(item);
          }
        }
      });
    },
    // 追加内容数据
    appendContent(content) {
      this.setContent(content);
    },
    // 切换标签
    switchTab(index) {
      return __async(this, null, function* () {
        if (this.activeTabIndex === index) return;
        this.activeTabIndex = index;
        this.currentPage = 1;
        this.hasMore = true;
        yield this.loadContent();
      });
    },
    // 判断是否显示区域
    shouldShowSection(type) {
      if (this.activeTabIndex === 0) return true;
      const typeMap = {
        "articles": 1,
        "videos": 2,
        "audios": 3
      };
      return this.activeTabIndex === typeMap[type];
    },
    // 获取区域标题
    getSectionTitle(type) {
      const titleMap = {
        "articles": this.activeTabIndex === 0 ? "热门文章" : "文章列表",
        "videos": this.activeTabIndex === 0 ? "推荐视频" : "视频列表",
        "audios": this.activeTabIndex === 0 ? "音频内容" : "音频列表"
      };
      return titleMap[type];
    },
    // 查看更多
    viewMore(type) {
      const typeMap = {
        "articles": 1,
        "videos": 2,
        "audios": 3
      };
      this.switchTab(typeMap[type]);
    },
    // 查看内容详情
    viewContent(content) {
      common_vendor.index.navigateTo({
        url: `/pages/learning/detail?id=${content.id}`
      });
    },
    // 开始学习
    startLearning(content) {
      this.viewContent(content);
    },
    // 播放音频
    playAudio(audio) {
      return __async(this, null, function* () {
        try {
          yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/start",
            method: "POST",
            data: {
              contentId: audio.id
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          common_vendor.index.navigateTo({
            url: `/pages/learning/audio?id=${audio.id}`
          });
        } catch (error) {
          console.error("播放音频失败:", error);
          common_vendor.index.showToast({
            title: "播放失败",
            icon: "error"
          });
        }
      });
    },
    // 显示搜索
    showSearch() {
      this.showSearchModal = true;
    },
    // 隐藏搜索
    hideSearch() {
      this.showSearchModal = false;
      this.searchKeyword = "";
      this.searchResults = [];
    },
    // 搜索输入
    onSearchInput() {
      return __async(this, null, function* () {
        if (!this.searchKeyword.trim()) {
          this.searchResults = [];
          return;
        }
        try {
          const res = yield common_vendor.index.request({
            url: "/api/v1/miniprogram/learning/contents",
            method: "GET",
            data: {
              page: 1,
              pageSize: 10,
              keyword: this.searchKeyword
            },
            header: {
              "Authorization": `Bearer ${common_vendor.index.getStorageSync("token")}`
            }
          });
          if (res.data.code === 0) {
            this.searchResults = res.data.data.list || [];
          }
        } catch (error) {
          console.error("搜索失败:", error);
        }
      });
    },
    // 加载更多
    loadMore() {
      if (this.hasMore && !this.loadingMore) {
        this.loadContent(true);
      }
    },
    // 更新时间
    updateTime() {
      this.currentTime = utils_format.getCurrentTime();
    },
    // 格式化数字
    formatNumber: utils_format.formatNumber,
    // 格式化时长
    formatDuration(minutes) {
      const hours = Math.floor(minutes / 60);
      const mins = minutes % 60;
      if (hours > 0) {
        return `${hours}:${mins.toString().padStart(2, "0")}`;
      }
      return `${mins}:00`;
    }
  }
};
if (!Array) {
  const _component_nf_tab_bar = common_vendor.resolveComponent("nf-tab-bar");
  _component_nf_tab_bar();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.t($data.currentTime),
    b: common_vendor.o((...args) => $options.showSearch && $options.showSearch(...args)),
    c: common_vendor.f($data.categoryTabs, (tab, index, i0) => {
      return {
        a: common_vendor.t(tab.name),
        b: index,
        c: $data.activeTabIndex === index ? 1 : "",
        d: common_vendor.o(($event) => $options.switchTab(index), index)
      };
    }),
    d: $data.activeTabIndex === 0
  }, $data.activeTabIndex === 0 ? {
    e: $data.featuredCourse.thumbnail,
    f: common_vendor.t($data.featuredCourse.title),
    g: common_vendor.t($data.featuredCourse.summary),
    h: common_vendor.o(($event) => $options.startLearning($data.featuredCourse))
  } : {}, {
    i: $data.activeTabIndex === 0
  }, $data.activeTabIndex === 0 ? {
    j: common_vendor.t($data.learningStats.weeklyHours),
    k: $data.learningStats.weeklyProgress + "%",
    l: common_vendor.t($data.learningStats.weeklyGoal),
    m: common_vendor.t($data.learningStats.articlesRead),
    n: common_vendor.t($data.learningStats.videosWatched),
    o: common_vendor.t($data.learningStats.audiosListened)
  } : {}, {
    p: $options.shouldShowSection("articles")
  }, $options.shouldShowSection("articles") ? {
    q: common_vendor.t($options.getSectionTitle("articles")),
    r: common_vendor.o(($event) => $options.viewMore("articles")),
    s: common_vendor.f($options.currentArticles, (article, k0, i0) => {
      return {
        a: article.thumbnailUrl,
        b: common_vendor.t(article.title),
        c: common_vendor.t(article.summary),
        d: common_vendor.t(article.duration),
        e: common_vendor.t(article.likeCount),
        f: common_vendor.t($options.formatNumber(article.viewCount)),
        g: article.id,
        h: common_vendor.o(($event) => $options.viewContent(article), article.id)
      };
    })
  } : {}, {
    t: $options.shouldShowSection("videos")
  }, $options.shouldShowSection("videos") ? {
    v: common_vendor.t($options.getSectionTitle("videos")),
    w: common_vendor.o(($event) => $options.viewMore("videos")),
    x: common_vendor.f($options.currentVideos, (video, k0, i0) => {
      return {
        a: video.thumbnailUrl,
        b: common_vendor.t($options.formatDuration(video.duration)),
        c: common_vendor.t(video.title),
        d: common_vendor.t(video.summary),
        e: common_vendor.t($options.formatNumber(video.viewCount)),
        f: common_vendor.t(video.likeCount),
        g: common_vendor.t(video.commentCount),
        h: video.id,
        i: common_vendor.o(($event) => $options.viewContent(video), video.id)
      };
    })
  } : {}, {
    y: $options.shouldShowSection("audios")
  }, $options.shouldShowSection("audios") ? {
    z: common_vendor.t($options.getSectionTitle("audios")),
    A: common_vendor.o(($event) => $options.viewMore("audios")),
    B: common_vendor.f($options.currentAudios, (audio, k0, i0) => {
      return common_vendor.e({
        a: common_vendor.t(audio.title),
        b: common_vendor.t(audio.summary),
        c: common_vendor.t(audio.duration),
        d: audio.isDownloaded
      }, audio.isDownloaded ? {} : {}, {
        e: common_vendor.o(($event) => $options.playAudio(audio), audio.id),
        f: audio.id,
        g: common_vendor.o(($event) => $options.viewContent(audio), audio.id)
      });
    })
  } : {}, {
    C: $data.hasMore
  }, $data.hasMore ? {
    D: common_vendor.t($data.loadingMore ? "加载中..." : "上拉加载更多")
  } : {}, {
    E: common_vendor.o((...args) => $options.loadMore && $options.loadMore(...args)),
    F: common_vendor.p({
      current: "learning"
    }),
    G: $data.showSearchModal
  }, $data.showSearchModal ? common_vendor.e({
    H: common_vendor.o([($event) => $data.searchKeyword = $event.detail.value, (...args) => $options.onSearchInput && $options.onSearchInput(...args)]),
    I: $data.searchKeyword,
    J: common_vendor.o((...args) => $options.hideSearch && $options.hideSearch(...args)),
    K: $data.searchResults.length > 0
  }, $data.searchResults.length > 0 ? {
    L: common_vendor.f($data.searchResults, (result, k0, i0) => {
      return {
        a: common_vendor.t(result.title),
        b: common_vendor.t(result.summary),
        c: result.id,
        d: common_vendor.o(($event) => $options.viewContent(result), result.id)
      };
    })
  } : $data.searchKeyword ? {} : {}, {
    M: $data.searchKeyword,
    N: common_vendor.o(() => {
    }),
    O: common_vendor.o((...args) => $options.hideSearch && $options.hideSearch(...args))
  }) : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render]]);
wx.createPage(MiniProgramPage);
