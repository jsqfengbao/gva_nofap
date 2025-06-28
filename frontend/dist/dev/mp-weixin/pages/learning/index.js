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
const apis_learningService = require("../../apis/learning-service.js");
const data_learningMock = require("../../data/learning-mock.js");
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
      featuredCourse: data_learningMock.featuredCourse,
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
        yield this.loadLearningStats();
      });
    },
    // 加载模拟数据
    loadMockData() {
      this.articles = [...data_learningMock.mockArticles];
      this.videos = [...data_learningMock.mockVideos];
      this.audios = [...data_learningMock.mockAudios];
    },
    // 加载学习统计
    loadLearningStats() {
      return __async(this, null, function* () {
        const result = yield apis_learningService.learningService.getStats();
        if (result.success) {
          this.learningStats = apis_learningService.learningService.transformStatsForDisplay(result.data);
          if (result.isMock) {
            console.log("使用模拟学习统计数据");
          }
        } else {
          console.error("加载学习统计失败");
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
          const page = loadMore ? this.currentPage + 1 : 1;
          const result = yield apis_learningService.learningService.getContents({
            page,
            pageSize: this.pageSize,
            contentType: contentType || void 0,
            sortBy: "view_count",
            order: "desc"
          });
          if (result.success) {
            const newContent = result.data.list || [];
            if (loadMore) {
              this.currentPage++;
              this.appendContent(newContent);
            } else {
              this.currentPage = 1;
              this.setContent(newContent);
            }
            this.hasMore = result.data.hasMore !== void 0 ? result.data.hasMore : newContent.length === this.pageSize;
            if (result.isMock) {
              console.log("使用模拟内容数据");
            }
          }
        } catch (error) {
          console.error("加载内容失败:", error);
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
        const result = yield apis_learningService.learningService.startLearning(audio.id);
        if (result.isMock) {
          console.log("学习记录使用模拟模式");
        }
        common_vendor.index.navigateTo({
          url: `/pages/learning/audio?id=${audio.id}`
        });
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
        const result = yield apis_learningService.learningService.searchContents({
          page: 1,
          pageSize: 10,
          keyword: this.searchKeyword
        });
        if (result.success) {
          this.searchResults = result.data.list || [];
          if (result.isMock) {
            console.log("搜索使用模拟数据");
          }
        } else {
          console.error("搜索失败");
          this.searchResults = [];
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
    e: $data.featuredCourse.thumbnailUrl,
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
        g: common_vendor.t(video.commentCount || 0),
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
