"use strict";
const common_vendor = require("../../common/vendor.js");
const data_assessmentQuestions = require("../../data/assessment-questions.js");
const _sfc_main = {
  name: "QuestionnairePage",
  data() {
    return {
      questions: data_assessmentQuestions.assessmentQuestions,
      currentQuestionIndex: 0,
      answers: {},
      selectedAnswer: null
    };
  },
  computed: {
    currentQuestion() {
      return this.questions[this.currentQuestionIndex] || {};
    },
    totalQuestions() {
      return this.questions.length;
    },
    isLastQuestion() {
      return this.currentQuestionIndex === this.totalQuestions - 1;
    }
  },
  methods: {
    selectAnswer(value) {
      this.selectedAnswer = value;
    },
    nextQuestion() {
      if (this.selectedAnswer === null) return;
      this.answers[this.currentQuestion.id] = this.selectedAnswer;
      if (this.isLastQuestion) {
        this.completeAssessment();
      } else {
        this.currentQuestionIndex++;
        this.loadQuestionAnswer();
      }
    },
    previousQuestion() {
      if (this.currentQuestionIndex > 0) {
        this.currentQuestionIndex--;
        this.loadQuestionAnswer();
      }
    },
    loadQuestionAnswer() {
      const savedAnswer = this.answers[this.currentQuestion.id];
      this.selectedAnswer = savedAnswer !== void 0 ? savedAnswer : null;
    },
    completeAssessment() {
      this.answers[this.currentQuestion.id] = this.selectedAnswer;
      common_vendor.index.showLoading({ title: "正在分析结果..." });
      setTimeout(() => {
        common_vendor.index.hideLoading();
        common_vendor.index.navigateTo({
          url: `/pages/assessment/result?answers=${encodeURIComponent(JSON.stringify(this.answers))}`
        });
      }, 2e3);
    },
    handleBack() {
      common_vendor.index.navigateBack();
    },
    showHelp() {
      common_vendor.index.showModal({
        title: "评估说明",
        content: "本评估采用专业的心理学量表，请诚实回答每个问题。所有信息严格保密。",
        showCancel: false,
        confirmText: "我知道了"
      });
    }
  },
  onLoad() {
    this.loadQuestionAnswer();
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.o((...args) => $options.handleBack && $options.handleBack(...args)),
    b: common_vendor.o((...args) => $options.showHelp && $options.showHelp(...args)),
    c: common_vendor.t($data.currentQuestionIndex + 1),
    d: common_vendor.t($options.totalQuestions),
    e: common_vendor.t(Math.round(($data.currentQuestionIndex + 1) / $options.totalQuestions * 100)),
    f: ($data.currentQuestionIndex + 1) / $options.totalQuestions * 100 + "%",
    g: common_vendor.n($options.currentQuestion.icon),
    h: common_vendor.t($options.currentQuestion.categoryName),
    i: common_vendor.t($options.currentQuestion.question),
    j: common_vendor.f($options.currentQuestion.options, (option, index, i0) => {
      return common_vendor.e({
        a: $data.selectedAnswer === option.value
      }, $data.selectedAnswer === option.value ? {} : {}, {
        b: $data.selectedAnswer === option.value ? 1 : "",
        c: common_vendor.t(option.text),
        d: index,
        e: $data.selectedAnswer === option.value ? 1 : "",
        f: common_vendor.o(($event) => $options.selectAnswer(option.value), index)
      });
    }),
    k: $data.currentQuestionIndex === 0 ? 1 : "",
    l: common_vendor.o((...args) => $options.previousQuestion && $options.previousQuestion(...args)),
    m: $data.currentQuestionIndex === 0,
    n: common_vendor.t($options.isLastQuestion ? "完成评估" : "下一题"),
    o: $data.selectedAnswer === null ? 1 : "",
    p: common_vendor.o((...args) => $options.nextQuestion && $options.nextQuestion(...args)),
    q: $data.selectedAnswer === null
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-0e64c823"]]);
wx.createPage(MiniProgramPage);
