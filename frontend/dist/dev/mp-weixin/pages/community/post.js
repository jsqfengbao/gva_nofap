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
const _sfc_main = {
  name: "CommunityPost",
  setup() {
    const currentTime = common_vendor.ref("9:41");
    const selectedCategory = common_vendor.ref(null);
    const postTitle = common_vendor.ref("");
    const postContent = common_vendor.ref("");
    const isAnonymous = common_vendor.ref(false);
    const isPublishing = common_vendor.ref(false);
    const categories = common_vendor.ref([
      { id: 1, name: "经验分享", icon: "💡" },
      { id: 2, name: "求助求鼓励", icon: "🤝" },
      { id: 3, name: "日常打卡", icon: "📅" },
      { id: 4, name: "成功故事", icon: "🎉" }
    ]);
    const canPublish = common_vendor.computed(() => {
      return selectedCategory.value && postTitle.value.trim().length > 0 && postContent.value.trim().length > 0 && !isPublishing.value;
    });
    const updateTime = () => {
      const now = /* @__PURE__ */ new Date();
      currentTime.value = `${now.getHours()}:${now.getMinutes().toString().padStart(2, "0")}`;
    };
    const selectCategory = (categoryId) => {
      selectedCategory.value = categoryId;
    };
    const toggleAnonymous = () => {
      isAnonymous.value = !isAnonymous.value;
    };
    const onTitleInput = (e) => {
      postTitle.value = e.detail.value;
    };
    const onContentInput = (e) => {
      postContent.value = e.detail.value;
    };
    const goBack = () => {
      common_vendor.index.navigateBack();
    };
    const publishPost = () => __async(this, null, function* () {
      if (!canPublish.value) return;
      isPublishing.value = true;
      try {
        yield new Promise((resolve) => setTimeout(resolve, 2e3));
        common_vendor.index.showToast({ title: "发布成功", icon: "success" });
        setTimeout(() => common_vendor.index.navigateBack(), 1500);
      } catch (error) {
        common_vendor.index.showToast({ title: "发布失败", icon: "none" });
      } finally {
        isPublishing.value = false;
      }
    });
    common_vendor.onMounted(() => {
      updateTime();
      setInterval(updateTime, 6e4);
    });
    return {
      currentTime,
      selectedCategory,
      postTitle,
      postContent,
      isAnonymous,
      isPublishing,
      categories,
      canPublish,
      selectCategory,
      toggleAnonymous,
      onTitleInput,
      onContentInput,
      goBack,
      publishPost
    };
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.t($setup.currentTime),
    b: common_vendor.o((...args) => $setup.goBack && $setup.goBack(...args)),
    c: !$setup.canPublish ? 1 : "",
    d: common_vendor.o((...args) => $setup.publishPost && $setup.publishPost(...args)),
    e: common_vendor.f($setup.categories, (category, k0, i0) => {
      return {
        a: common_vendor.t(category.icon),
        b: common_vendor.t(category.name),
        c: category.id,
        d: $setup.selectedCategory === category.id ? 1 : "",
        e: common_vendor.o(($event) => $setup.selectCategory(category.id), category.id)
      };
    }),
    f: common_vendor.o([($event) => $setup.postTitle = $event.detail.value, (...args) => $setup.onTitleInput && $setup.onTitleInput(...args)]),
    g: $setup.postTitle,
    h: common_vendor.t($setup.postTitle.length),
    i: common_vendor.o([($event) => $setup.postContent = $event.detail.value, (...args) => $setup.onContentInput && $setup.onContentInput(...args)]),
    j: $setup.postContent,
    k: common_vendor.t($setup.postContent.length),
    l: $setup.isAnonymous ? 1 : "",
    m: common_vendor.o((...args) => $setup.toggleAnonymous && $setup.toggleAnonymous(...args)),
    n: $setup.isPublishing
  }, $setup.isPublishing ? {} : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-af14bcb1"]]);
wx.createPage(MiniProgramPage);
