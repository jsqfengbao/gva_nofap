"use strict";
Object.defineProperty(exports, Symbol.toStringTag, { value: "Module" });
const common_vendor = require("./common/vendor.js");
const components_ui_index = require("./components/ui/index.js");
if (!Math) {
  "./pages/index/index.js";
  "./pages/welcome/welcome.js";
  "./pages/assessment/index.js";
  "./pages/assessment/assessment.js";
  "./pages/assessment/questionnaire.js";
  "./pages/assessment/result.js";
  "./pages/checkin/index.js";
  "./pages/checkin/history.js";
  "./pages/checkin/weekly.js";
  "./pages/checkin/calendar.js";
  "./pages/achievement/index.js";
  "./pages/progress/index.js";
  "./pages/community/index.js";
  "./pages/community/post.js";
  "./pages/community/detail.js";
  "./pages/emergency/index.js";
  "./pages/emergency/articles.js";
  "./pages/learning/index.js";
  "./pages/learning/detail.js";
  "./pages/profile/index.js";
  "./pages/profile/setup.js";
  "./pages/profile/auth.js";
  "./pages/profile/notification.js";
  "./pages/profile/privacy.js";
  "./pages/profile/help.js";
  "./pages/about/about.js";
  "./pages/legal/service-agreement.js";
  "./pages/legal/privacy-policy.js";
}
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "App",
  setup(__props) {
    common_vendor.onLaunch(() => {
      console.log("App Launch");
      const isFirstLaunch = !common_vendor.index.getStorageSync("hasLaunched");
      if (isFirstLaunch) {
        common_vendor.index.setStorageSync("hasLaunched", true);
        common_vendor.index.redirectTo({
          url: "/pages/welcome/welcome"
        });
      }
    });
    common_vendor.onShow(() => {
      console.log("App Show");
    });
    common_vendor.onHide(() => {
      console.log("App Hide");
    });
    return (_ctx, _cache) => {
      return {};
    };
  }
});
function createApp() {
  const app = common_vendor.createSSRApp(_sfc_main);
  const pinia = common_vendor.createPinia();
  app.use(pinia);
  app.use(components_ui_index.UIComponents);
  return {
    app
  };
}
createApp().app.mount("#app");
exports.createApp = createApp;
