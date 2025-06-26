"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  name: "CommunityDetail",
  setup() {
    const currentTime = common_vendor.ref("9:41");
    const loading = common_vendor.ref(false);
    const loadingComments = common_vendor.ref(false);
    const hasMoreComments = common_vendor.ref(true);
    const commentText = common_vendor.ref("");
    const sortType = common_vendor.ref("time");
    const post = common_vendor.ref({
      id: 1,
      title: "🎉 100天里程碑达成！",
      content: "想跟大家分享一些这段路程的心得和体验。首先，前50天是最困难的，特别是前2周，身体和心理都会有很大的反应。我的建议是：\n\n1. 建立良好的作息时间\n2. 多运动，转移注意力\n3. 寻找新的兴趣爱好\n4. 及时寻求社区支持\n\n现在100天了，感觉整个人的精神状态都有了很大改善，希望能继续坚持下去，也希望能帮助到更多的朋友！",
      category: 4,
      categoryName: "成功故事",
      userNickname: "坚持者_阳光",
      isAnonymous: false,
      viewCount: 156,
      likeCount: 128,
      commentCount: 45,
      isLiked: false,
      createdAt: new Date(Date.now() - 18e4).toISOString()
    });
    const comments = common_vendor.ref([
      {
        id: 1,
        content: "太厉害了！100天真的不容易，请问你是怎么度过最困难的前期的？",
        userNickname: "新手_加油",
        likeCount: 12,
        isLiked: false,
        createdAt: new Date(Date.now() - 6e4).toISOString(),
        replies: [
          {
            id: 11,
            content: "我也想知道这个问题的答案",
            userNickname: "同路人",
            createdAt: new Date(Date.now() - 3e4).toISOString()
          }
        ]
      },
      {
        id: 2,
        content: "恭喜恭喜！你的分享很有用，收藏了！",
        userNickname: "学习者_努力",
        likeCount: 8,
        isLiked: true,
        createdAt: new Date(Date.now() - 12e4).toISOString(),
        replies: []
      }
    ]);
    const updateTime = () => {
      const now = /* @__PURE__ */ new Date();
      currentTime.value = `${now.getHours()}:${now.getMinutes().toString().padStart(2, "0")}`;
    };
    const goBack = () => {
      common_vendor.index.navigateBack();
    };
    const showMoreOptions = () => {
      common_vendor.index.showActionSheet({
        itemList: ["举报", "分享", "收藏"],
        success: (res) => {
          console.log("选中了第" + (res.tapIndex + 1) + "个按钮");
        }
      });
    };
    const toggleLike = () => {
      post.value.isLiked = !post.value.isLiked;
      post.value.likeCount += post.value.isLiked ? 1 : -1;
      common_vendor.index.showToast({
        title: post.value.isLiked ? "点赞成功" : "取消点赞",
        icon: "success"
      });
    };
    const toggleCommentLike = (comment) => {
      comment.isLiked = !comment.isLiked;
      comment.likeCount += comment.isLiked ? 1 : -1;
    };
    const focusCommentInput = () => {
      common_vendor.index.showToast({
        title: "评论功能开发中",
        icon: "none"
      });
    };
    const sharePost = () => {
      common_vendor.index.showToast({
        title: "分享功能开发中",
        icon: "none"
      });
    };
    const replyToComment = (comment) => {
      commentText.value = `@${comment.userNickname} `;
    };
    const submitComment = () => {
      if (!commentText.value.trim()) return;
      common_vendor.index.showToast({
        title: "评论发布成功",
        icon: "success"
      });
      commentText.value = "";
    };
    const setSortType = (type) => {
      sortType.value = type;
    };
    const loadMoreComments = () => {
      loadingComments.value = true;
      setTimeout(() => {
        loadingComments.value = false;
        hasMoreComments.value = false;
      }, 1e3);
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
      post.value.viewCount++;
    });
    return {
      currentTime,
      loading,
      loadingComments,
      hasMoreComments,
      commentText,
      sortType,
      post,
      comments,
      goBack,
      showMoreOptions,
      toggleLike,
      toggleCommentLike,
      focusCommentInput,
      sharePost,
      replyToComment,
      submitComment,
      setSortType,
      loadMoreComments,
      getAvatarClass,
      getAvatarText,
      getCategoryClass,
      formatTime
    };
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.t($setup.currentTime),
    b: common_vendor.o((...args) => $setup.goBack && $setup.goBack(...args)),
    c: common_vendor.o((...args) => $setup.showMoreOptions && $setup.showMoreOptions(...args)),
    d: common_vendor.t($setup.getAvatarText($setup.post.userNickname)),
    e: common_vendor.n($setup.getAvatarClass($setup.post.userNickname)),
    f: common_vendor.t($setup.post.userNickname),
    g: common_vendor.t($setup.formatTime($setup.post.createdAt)),
    h: common_vendor.t($setup.post.categoryName),
    i: common_vendor.n($setup.getCategoryClass($setup.post.category)),
    j: common_vendor.t($setup.post.viewCount),
    k: common_vendor.t($setup.post.title),
    l: common_vendor.t($setup.post.content),
    m: $setup.post.isLiked ? 1 : "",
    n: common_vendor.t($setup.post.likeCount),
    o: common_vendor.o((...args) => $setup.toggleLike && $setup.toggleLike(...args)),
    p: common_vendor.t($setup.post.commentCount),
    q: common_vendor.o((...args) => $setup.focusCommentInput && $setup.focusCommentInput(...args)),
    r: common_vendor.o((...args) => $setup.sharePost && $setup.sharePost(...args)),
    s: common_vendor.t($setup.comments.length),
    t: $setup.sortType === "time" ? 1 : "",
    v: common_vendor.o(($event) => $setup.setSortType("time")),
    w: $setup.sortType === "hot" ? 1 : "",
    x: common_vendor.o(($event) => $setup.setSortType("hot")),
    y: common_vendor.f($setup.comments, (comment, k0, i0) => {
      return common_vendor.e({
        a: common_vendor.t($setup.getAvatarText(comment.userNickname)),
        b: common_vendor.n($setup.getAvatarClass(comment.userNickname)),
        c: common_vendor.t(comment.userNickname),
        d: common_vendor.t($setup.formatTime(comment.createdAt)),
        e: comment.isLiked ? 1 : "",
        f: comment.likeCount > 0
      }, comment.likeCount > 0 ? {
        g: common_vendor.t(comment.likeCount)
      } : {}, {
        h: common_vendor.o(($event) => $setup.toggleCommentLike(comment), comment.id),
        i: common_vendor.o(($event) => $setup.replyToComment(comment), comment.id),
        j: common_vendor.t(comment.content),
        k: comment.replies && comment.replies.length > 0
      }, comment.replies && comment.replies.length > 0 ? {
        l: common_vendor.f(comment.replies, (reply, k1, i1) => {
          return {
            a: common_vendor.t($setup.getAvatarText(reply.userNickname)),
            b: common_vendor.n($setup.getAvatarClass(reply.userNickname)),
            c: common_vendor.t(reply.userNickname),
            d: common_vendor.t($setup.formatTime(reply.createdAt)),
            e: common_vendor.t(reply.content),
            f: reply.id
          };
        })
      } : {}, {
        m: comment.id
      });
    }),
    z: $setup.hasMoreComments
  }, $setup.hasMoreComments ? {
    A: common_vendor.t($setup.loadingComments ? "加载中..." : "查看更多评论"),
    B: common_vendor.o((...args) => $setup.loadMoreComments && $setup.loadMoreComments(...args))
  } : {}, {
    C: common_vendor.o((...args) => $setup.submitComment && $setup.submitComment(...args)),
    D: $setup.commentText,
    E: common_vendor.o(($event) => $setup.commentText = $event.detail.value),
    F: $setup.commentText.trim() ? 1 : "",
    G: common_vendor.o((...args) => $setup.submitComment && $setup.submitComment(...args)),
    H: $setup.loading
  }, $setup.loading ? {} : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-e186c394"]]);
wx.createPage(MiniProgramPage);
