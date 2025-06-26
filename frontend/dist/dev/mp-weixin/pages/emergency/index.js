"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  name: "EmergencyHelp",
  setup() {
    const currentTime = common_vendor.ref("9:41");
    const isBreathing = common_vendor.ref(true);
    const onlineVolunteers = common_vendor.ref(23);
    const showHelpModal = common_vendor.ref(false);
    const showResourceModal = common_vendor.ref(false);
    const selectedResource = common_vendor.reactive({
      title: "",
      content: "",
      id: 0
    });
    const currentQuote = common_vendor.ref({
      text: "你比你想象的更坚强，这个时刻会过去的。",
      author: "戒色社区"
    });
    const quotes = [
      {
        text: "你比你想象的更坚强，这个时刻会过去的。",
        author: "戒色社区"
      },
      {
        text: "每个选择都是新的开始，每一刻都有改变的可能。",
        author: "戒色助手"
      },
      {
        text: "成长需要时间，对自己要有耐心。",
        author: "励志语录"
      }
    ];
    const emergencyResources = common_vendor.ref([
      {
        id: 1,
        title: "4-7-8深呼吸练习",
        type: 1,
        content: "这是一种简单而有效的呼吸技巧，可以帮助您快速平静下来。\n\n步骤：\n1. 吸气4秒钟\n2. 屏住呼吸7秒钟\n3. 呼气8秒钟\n4. 重复3-4次\n\n这种呼吸方式可以激活副交感神经系统，帮助身体放松。",
        duration: 180
      },
      {
        id: 2,
        title: "5分钟正念冥想",
        type: 2,
        content: "正念冥想可以帮助您专注当下，减少负面情绪的影响。\n\n指导：\n1. 找一个安静的地方坐下\n2. 闭上眼睛，专注于呼吸\n3. 当思绪飘散时，温和地把注意力拉回呼吸\n4. 观察身体的感觉，不做判断\n5. 保持5分钟\n\n记住：没有'正确'或'错误'的冥想，只要保持观察即可。",
        duration: 300
      },
      {
        id: 5,
        title: "快速运动指导",
        type: 5,
        content: "运动可以释放内啡肽，帮助改善心情。这里有一套简单的运动，无需器械：\n\n**热身（1分钟）**\n• 原地踏步 30秒\n• 手臂绕圈 30秒\n\n**主要运动（3分钟）**\n• 俯卧撑 20个（可膝盖着地）\n• 深蹲 20个\n• 平板支撑 30秒\n• 开合跳 20个\n\n**放松（1分钟）**\n• 深呼吸 30秒\n• 拉伸手臂和腿部 30秒",
        duration: 300
      }
    ]);
    const updateTime = () => {
      const now = /* @__PURE__ */ new Date();
      const hours = now.getHours().toString().padStart(2, "0");
      const minutes = now.getMinutes().toString().padStart(2, "0");
      currentTime.value = `${hours}:${minutes}`;
    };
    const rotateQuote = () => {
      const randomIndex = Math.floor(Math.random() * quotes.length);
      currentQuote.value = quotes[randomIndex];
    };
    const goBack = () => {
      common_vendor.index.navigateBack();
    };
    const callEmergency = () => {
      common_vendor.index.showModal({
        title: "紧急联系",
        content: "如果遇到生命危险，请立即拨打120或当地急救电话",
        confirmText: "知道了",
        showCancel: false
      });
    };
    const startBreathingExercise = () => {
      const resource = emergencyResources.value.find((r) => r.type === 1);
      if (resource) {
        openResourceModal(resource);
      }
    };
    const startPhysicalExercise = () => {
      const resource = emergencyResources.value.find((r) => r.type === 5);
      if (resource) {
        openResourceModal(resource);
      }
    };
    const openActivity = (type) => {
      let resource;
      switch (type) {
        case "meditation":
          resource = emergencyResources.value.find((r) => r.type === 2);
          break;
        case "music":
          common_vendor.index.showToast({
            title: "音乐功能开发中",
            icon: "none"
          });
          return;
        case "reading":
          common_vendor.index.navigateTo({
            url: "/pages/emergency/articles"
          });
          return;
        case "puzzle":
          common_vendor.index.showToast({
            title: "游戏功能开发中",
            icon: "none"
          });
          return;
      }
      if (resource) {
        openResourceModal(resource);
      }
    };
    const openResourceModal = (resource) => {
      selectedResource.title = resource.title;
      selectedResource.content = resource.content;
      selectedResource.id = resource.id;
      showResourceModal.value = true;
    };
    const closeResourceModal = () => {
      showResourceModal.value = false;
    };
    const useResource = () => {
      common_vendor.index.showToast({
        title: "开始使用",
        icon: "success"
      });
      closeResourceModal();
    };
    const rateResource = () => {
      common_vendor.index.showToast({
        title: "感谢你的反馈",
        icon: "success"
      });
    };
    const requestHelp = () => {
      showHelpModal.value = true;
    };
    const closeHelpModal = () => {
      showHelpModal.value = false;
    };
    const selectHelpType = (type) => {
      closeHelpModal();
      common_vendor.index.showModal({
        title: "求助已发送",
        content: "你的求助信息已发送给在线志愿者，他们会尽快回复你。",
        confirmText: "好的",
        showCancel: false,
        success: () => {
          common_vendor.index.showToast({
            title: "正在为您匹配志愿者",
            icon: "loading",
            duration: 2e3
          });
        }
      });
    };
    const viewStories = () => {
      common_vendor.index.navigateTo({
        url: "/pages/community/index?category=4"
        // 成功故事分类
      });
    };
    common_vendor.onMounted(() => {
      updateTime();
      setInterval(updateTime, 6e4);
      setInterval(rotateQuote, 3e4);
    });
    return {
      currentTime,
      isBreathing,
      onlineVolunteers,
      currentQuote,
      showHelpModal,
      showResourceModal,
      selectedResource,
      goBack,
      callEmergency,
      startBreathingExercise,
      startPhysicalExercise,
      openActivity,
      requestHelp,
      closeHelpModal,
      selectHelpType,
      viewStories,
      closeResourceModal,
      useResource,
      rateResource
    };
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.t($setup.currentTime),
    b: common_vendor.o((...args) => $setup.goBack && $setup.goBack(...args)),
    c: common_vendor.o((...args) => $setup.callEmergency && $setup.callEmergency(...args)),
    d: $setup.isBreathing ? 1 : "",
    e: common_vendor.o((...args) => $setup.startBreathingExercise && $setup.startBreathingExercise(...args)),
    f: common_vendor.o((...args) => $setup.startPhysicalExercise && $setup.startPhysicalExercise(...args)),
    g: common_vendor.o(($event) => $setup.openActivity("puzzle")),
    h: common_vendor.o(($event) => $setup.openActivity("meditation")),
    i: common_vendor.o(($event) => $setup.openActivity("music")),
    j: common_vendor.o(($event) => $setup.openActivity("reading")),
    k: common_vendor.t($setup.onlineVolunteers),
    l: common_vendor.o((...args) => $setup.requestHelp && $setup.requestHelp(...args)),
    m: common_vendor.o((...args) => $setup.viewStories && $setup.viewStories(...args)),
    n: common_vendor.t($setup.currentQuote.text),
    o: common_vendor.t($setup.currentQuote.author),
    p: $setup.showHelpModal
  }, $setup.showHelpModal ? {
    q: common_vendor.o((...args) => $setup.closeHelpModal && $setup.closeHelpModal(...args)),
    r: common_vendor.o(($event) => $setup.selectHelpType(1)),
    s: common_vendor.o(($event) => $setup.selectHelpType(2)),
    t: common_vendor.o(($event) => $setup.selectHelpType(3)),
    v: common_vendor.o(($event) => $setup.selectHelpType(4)),
    w: common_vendor.o(() => {
    }),
    x: common_vendor.o((...args) => $setup.closeHelpModal && $setup.closeHelpModal(...args))
  } : {}, {
    y: $setup.showResourceModal
  }, $setup.showResourceModal ? {
    z: common_vendor.t($setup.selectedResource.title),
    A: common_vendor.o((...args) => $setup.closeResourceModal && $setup.closeResourceModal(...args)),
    B: common_vendor.t($setup.selectedResource.content),
    C: common_vendor.o((...args) => $setup.useResource && $setup.useResource(...args)),
    D: common_vendor.o((...args) => $setup.rateResource && $setup.rateResource(...args)),
    E: common_vendor.o(() => {
    }),
    F: common_vendor.o((...args) => $setup.closeResourceModal && $setup.closeResourceModal(...args))
  } : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-f6b69931"]]);
wx.createPage(MiniProgramPage);
