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
const NfTabBar = () => "../../components/ui/navigation/NfTabBar.js";
const _sfc_main = {
  name: "CommunityIndex",
  components: {
    NfTabBar
  },
  setup() {
    const currentTime = common_vendor.ref("9:41");
    const selectedCategory = common_vendor.ref(0);
    const posts = common_vendor.ref([]);
    const loading = common_vendor.ref(false);
    const hasMore = common_vendor.ref(true);
    const page = common_vendor.ref(1);
    const updateTime = () => {
      const now = /* @__PURE__ */ new Date();
      currentTime.value = `${now.getHours()}:${now.getMinutes().toString().padStart(2, "0")}`;
    };
    const selectCategory = (category) => {
      selectedCategory.value = category;
      page.value = 1;
      posts.value = [];
      loadPosts();
    };
    const loadPosts = () => __async(this, null, function* () {
      if (loading.value) return;
      loading.value = true;
      try {
        const mockPosts = [
          {
            id: 1,
            title: "🎉 100天里程碑达成！",
            content: "想跟大家分享一些这段路程的心得...",
            category: 4,
            categoryName: "成功故事",
            userNickname: "坚持者_阳光",
            isAnonymous: false,
            viewCount: 156,
            likeCount: 128,
            commentCount: 45,
            isLiked: false,
            createdAt: new Date(Date.now() - 18e4).toISOString()
          },
          {
            id: 2,
            title: "感觉很困难，需要大家的鼓励 😔",
            content: "第7天了，总是想要放弃...",
            category: 2,
            categoryName: "求助求鼓励",
            userNickname: "新手_求助",
            isAnonymous: false,
            viewCount: 89,
            likeCount: 32,
            commentCount: 18,
            isLiked: false,
            createdAt: new Date(Date.now() - 9e5).toISOString()
          },
          {
            id: 3,
            title: "第30天打卡 ✅",
            content: "今天是第30天，感觉状态不错！",
            category: 3,
            categoryName: "日常打卡",
            userNickname: "努力的小伙",
            isAnonymous: false,
            viewCount: 45,
            likeCount: 28,
            commentCount: 12,
            isLiked: true,
            createdAt: new Date(Date.now() - 18e5).toISOString()
          },
          {
            id: 4,
            title: "分享一些有效的转移注意力方法",
            content: "运动、阅读、冥想都是很好的方式...",
            category: 1,
            categoryName: "经验分享",
            userNickname: "经验分享者",
            isAnonymous: false,
            viewCount: 234,
            likeCount: 89,
            commentCount: 36,
            isLiked: false,
            createdAt: new Date(Date.now() - 36e5).toISOString()
          }
        ];
        if (page.value === 1) {
          posts.value = mockPosts;
        }
        hasMore.value = false;
      } catch (error) {
        console.error("加载帖子失败:", error);
        common_vendor.index.showToast({
          title: "网络错误",
          icon: "none"
        });
      } finally {
        loading.value = false;
      }
    });
    const loadMorePosts = () => {
      if (hasMore.value && !loading.value) {
        page.value++;
        loadPosts();
      }
    };
    const toggleLike = (post) => __async(this, null, function* () {
      post.isLiked = !post.isLiked;
      post.likeCount += post.isLiked ? 1 : -1;
      common_vendor.index.showToast({
        title: post.isLiked ? "点赞成功" : "取消点赞",
        icon: "success"
      });
    });
    const viewPostDetail = (postId) => {
      common_vendor.index.navigateTo({
        url: `/pages/community/detail?id=${postId}`
      });
    };
    const openPostModal = () => {
      common_vendor.index.navigateTo({
        url: "/pages/community/post"
      });
    };
    const openSearch = () => {
      common_vendor.index.showToast({
        title: "搜索功能开发中",
        icon: "none"
      });
    };
    const encourageUser = (post) => {
      common_vendor.index.showToast({
        title: "已给TA发送鼓励",
        icon: "success"
      });
    };
    const getAvatarClass = (nickname) => {
      if (nickname === "匿名用户") return "anonymous";
      const hash = nickname.charCodeAt(0) % 5;
      return `avatar-${hash}`;
    };
    const getAvatarText = (nickname) => {
      if (nickname === "匿名用户") return "匿";
      return nickname.substring(0, 1);
    };
    const getCategoryClass = (category) => {
      const classes = ["", "experience", "help", "checkin", "success"];
      return classes[category] || "";
    };
    const formatTime = (dateStr) => {
      const date = new Date(dateStr);
      const now = /* @__PURE__ */ new Date();
      const diff = now - date;
      if (diff < 6e4) {
        return "刚刚";
      } else if (diff < 36e5) {
        return `${Math.floor(diff / 6e4)}分钟前`;
      } else if (diff < 864e5) {
        return `${Math.floor(diff / 36e5)}小时前`;
      } else {
        return date.toLocaleDateString();
      }
    };
    common_vendor.onMounted(() => {
      updateTime();
      setInterval(updateTime, 6e4);
      loadPosts();
    });
    return {
      currentTime,
      selectedCategory,
      posts,
      loading,
      hasMore,
      selectCategory,
      loadMorePosts,
      toggleLike,
      viewPostDetail,
      openPostModal,
      openSearch,
      encourageUser,
      getAvatarClass,
      getAvatarText,
      getCategoryClass,
      formatTime
    };
  }
};
if (!Array) {
  const _component_nf_tab_bar = common_vendor.resolveComponent("nf-tab-bar");
  _component_nf_tab_bar();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.t($setup.currentTime),
    b: common_vendor.o((...args) => $setup.openSearch && $setup.openSearch(...args)),
    c: common_vendor.o((...args) => $setup.openPostModal && $setup.openPostModal(...args)),
    d: $setup.selectedCategory === 0 ? 1 : "",
    e: common_vendor.o(($event) => $setup.selectCategory(0)),
    f: $setup.selectedCategory === 1 ? 1 : "",
    g: common_vendor.o(($event) => $setup.selectCategory(1)),
    h: $setup.selectedCategory === 2 ? 1 : "",
    i: common_vendor.o(($event) => $setup.selectCategory(2)),
    j: $setup.selectedCategory === 3 ? 1 : "",
    k: common_vendor.o(($event) => $setup.selectCategory(3)),
    l: $setup.selectedCategory === 4 ? 1 : "",
    m: common_vendor.o(($event) => $setup.selectCategory(4)),
    n: common_vendor.f($setup.posts, (post, k0, i0) => {
      return common_vendor.e({
        a: common_vendor.t($setup.getAvatarText(post.userNickname)),
        b: common_vendor.n($setup.getAvatarClass(post.userNickname)),
        c: common_vendor.t(post.userNickname),
        d: common_vendor.t($setup.formatTime(post.createdAt)),
        e: common_vendor.t(post.categoryName),
        f: common_vendor.n($setup.getCategoryClass(post.category)),
        g: common_vendor.t(post.title),
        h: common_vendor.t(post.content),
        i: post.isLiked ? 1 : "",
        j: common_vendor.t(post.likeCount),
        k: common_vendor.o(($event) => $setup.toggleLike(post), post.id),
        l: common_vendor.t(post.commentCount),
        m: common_vendor.o(($event) => $setup.viewPostDetail(post.id), post.id),
        n: common_vendor.t(post.viewCount),
        o: post.category === 2
      }, post.category === 2 ? {
        p: common_vendor.o(($event) => $setup.encourageUser(post), post.id)
      } : {}, {
        q: post.id,
        r: common_vendor.o(($event) => $setup.viewPostDetail(post.id), post.id)
      });
    }),
    o: $setup.hasMore
  }, $setup.hasMore ? {
    p: common_vendor.t($setup.loading ? "加载中..." : "上拉加载更多")
  } : {}, {
    q: !$setup.hasMore && $setup.posts.length > 0
  }, !$setup.hasMore && $setup.posts.length > 0 ? {} : {}, {
    r: common_vendor.o((...args) => $setup.loadMorePosts && $setup.loadMorePosts(...args)),
    s: common_vendor.p({
      current: "community"
    })
  });
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-485a05f3"]]);
wx.createPage(MiniProgramPage);
