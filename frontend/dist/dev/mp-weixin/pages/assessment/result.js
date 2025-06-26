"use strict";
const common_vendor = require("../../common/vendor.js");
const data_assessmentQuestions = require("../../data/assessment-questions.js");
const _sfc_main = {
  name: "AssessmentResultPage",
  data() {
    return {
      assessmentResult: null,
      levelInfo: null
    };
  },
  methods: {
    calculateResult(answers) {
      const totalAnswers = Object.values(answers);
      const totalScore = totalAnswers.reduce((sum, value) => sum + (value || 0), 0) * 4;
      const level = data_assessmentQuestions.assessmentConfig.levels.find(
        (l) => totalScore >= l.range[0] && totalScore <= l.range[1]
      ) || data_assessmentQuestions.assessmentConfig.levels[0];
      return {
        totalScore: Math.min(200, totalScore),
        categoryScores: {
          frequency: Math.random() * 100,
          control: Math.random() * 100,
          impact: Math.random() * 100,
          psychology: Math.random() * 100,
          social: Math.random() * 100,
          health: Math.random() * 100,
          cognitive: Math.random() * 100
        },
        level,
        completedAt: (/* @__PURE__ */ new Date()).toISOString()
      };
    },
    getCategoryName(category) {
      const names = {
        frequency: "行为频率",
        control: "自控能力",
        impact: "生活影响",
        psychology: "心理状态",
        social: "社交关系",
        health: "身体健康",
        cognitive: "认知功能"
      };
      return names[category] || category;
    },
    getProgressColor(score) {
      if (score < 30) return "#10B981";
      if (score < 60) return "#F59E0B";
      if (score < 80) return "#EF4444";
      return "#DC2626";
    },
    goBack() {
      common_vendor.index.navigateBack();
    },
    startRecoveryPlan() {
      common_vendor.index.showToast({
        title: "正在为您制定计划...",
        icon: "none"
      });
      setTimeout(() => {
        common_vendor.index.switchTab({
          url: "/pages/index/index"
        });
      }, 1500);
    },
    saveResult() {
      try {
        const existingResults = common_vendor.index.getStorageSync("assessmentHistory") || [];
        existingResults.push(this.assessmentResult);
        common_vendor.index.setStorageSync("assessmentHistory", existingResults);
        common_vendor.index.showToast({
          title: "评估结果已保存",
          icon: "success"
        });
      } catch (error) {
        common_vendor.index.showToast({
          title: "保存失败，请重试",
          icon: "error"
        });
      }
    },
    retakeAssessment() {
      common_vendor.index.showModal({
        title: "重新评估",
        content: "确定要重新进行评估吗？当前结果将会被覆盖。",
        success: (res) => {
          if (res.confirm) {
            common_vendor.index.redirectTo({
              url: "/pages/assessment/questionnaire"
            });
          }
        }
      });
    }
  },
  onLoad(options) {
    if (options.answers) {
      try {
        const answers = JSON.parse(decodeURIComponent(options.answers));
        this.assessmentResult = this.calculateResult(answers);
        this.levelInfo = this.assessmentResult.level;
        this.levelInfo.recommendations = [
          "建立规律的作息时间",
          "培养健康的兴趣爱好",
          "寻求专业指导帮助",
          "使用戒色助手的各项功能"
        ];
      } catch (error) {
        console.error("解析评估答案失败:", error);
        common_vendor.index.showToast({
          title: "数据加载失败",
          icon: "error"
        });
        setTimeout(() => {
          common_vendor.index.navigateBack();
        }, 2e3);
      }
    }
  }
};
if (!Array) {
  const _component_NfNavBar = common_vendor.resolveComponent("NfNavBar");
  const _component_NfCard = common_vendor.resolveComponent("NfCard");
  const _component_NfButton = common_vendor.resolveComponent("NfButton");
  (_component_NfNavBar + _component_NfCard + _component_NfButton)();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.o($options.goBack),
    b: common_vendor.p({
      title: "评估结果",
      ["show-back"]: true
    }),
    c: $data.levelInfo.color,
    d: $data.levelInfo.color + "20",
    e: common_vendor.t($data.assessmentResult.totalScore),
    f: common_vendor.t($data.levelInfo.name),
    g: $data.levelInfo.color,
    h: common_vendor.p({
      type: "gradient"
    }),
    i: common_vendor.t($data.levelInfo.description),
    j: common_vendor.p({
      title: "结果解读",
      icon: "fa-lightbulb"
    }),
    k: common_vendor.f($data.assessmentResult.categoryScores, (score, category, i0) => {
      return {
        a: common_vendor.t($options.getCategoryName(category)),
        b: common_vendor.t(Math.round(score)),
        c: score + "%",
        d: $options.getProgressColor(score),
        e: category
      };
    }),
    l: common_vendor.p({
      title: "详细分析",
      icon: "fa-chart-bar"
    }),
    m: common_vendor.f($data.levelInfo.recommendations, (recommendation, index, i0) => {
      return {
        a: common_vendor.t(recommendation),
        b: index
      };
    }),
    n: common_vendor.p({
      title: "个性化建议",
      icon: "fa-heart"
    }),
    o: common_vendor.o($options.startRecoveryPlan),
    p: common_vendor.p({
      type: "primary",
      size: "large",
      label: "开始康复计划",
      ["full-width"]: true,
      ["icon-left"]: "fa-play"
    }),
    q: common_vendor.o($options.saveResult),
    r: common_vendor.p({
      type: "secondary",
      size: "medium",
      label: "保存评估结果",
      ["full-width"]: true,
      ["icon-left"]: "fa-save"
    }),
    s: common_vendor.o($options.retakeAssessment),
    t: common_vendor.p({
      type: "secondary",
      size: "medium",
      label: "重新评估",
      ["full-width"]: true,
      ["icon-left"]: "fa-redo"
    })
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-79bc4ac3"]]);
wx.createPage(MiniProgramPage);
