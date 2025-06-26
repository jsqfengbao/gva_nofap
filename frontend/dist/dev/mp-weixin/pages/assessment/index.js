"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  setup(__props) {
    const historyList = common_vendor.ref([]);
    common_vendor.onMounted(() => {
      loadHistoryList();
    });
    const startAssessment = () => {
      common_vendor.index.navigateTo({
        url: "/pages/assessment/questionnaire"
      });
    };
    const viewResult = (item) => {
      common_vendor.index.navigateTo({
        url: `/pages/assessment/result?id=${item.id}`
      });
    };
    const loadHistoryList = () => {
      historyList.value = [];
    };
    const formatDate = (dateStr) => {
      const date = new Date(dateStr);
      return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, "0")}-${String(date.getDate()).padStart(2, "0")}`;
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.o(startAssessment),
        b: historyList.value.length === 0
      }, historyList.value.length === 0 ? {} : {
        c: common_vendor.f(historyList.value, (item, k0, i0) => {
          return {
            a: common_vendor.t(formatDate(item.createTime)),
            b: common_vendor.t(item.score),
            c: common_vendor.t(item.level),
            d: common_vendor.o(($event) => viewResult(item), item.id),
            e: item.id
          };
        })
      });
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-9dc5ab2e"]]);
wx.createPage(MiniProgramPage);
