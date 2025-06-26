"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  __name: "help",
  setup(__props) {
    const loading = common_vendor.ref(false);
    const loadingText = common_vendor.ref("加载中...");
    const searchText = common_vendor.ref("");
    const showDetailModal = common_vendor.ref(false);
    const selectedFaq = common_vendor.ref(null);
    const faqCategories = common_vendor.ref([
      {
        id: 1,
        title: "账户与登录",
        icon: "👤",
        iconClass: "account-icon",
        expanded: false,
        faqs: [
          {
            id: 1,
            question: "如何注册账户？",
            answer: '本应用使用微信登录，无需单独注册。首次使用时，点击"微信登录"按钮，授权后即可自动创建账户。',
            steps: [
              '打开应用，点击"微信登录"',
              "在微信中确认授权",
              "完成用户信息设置",
              "开始使用应用功能"
            ]
          },
          {
            id: 2,
            question: "忘记了登录密码怎么办？",
            answer: "本应用使用微信登录，无需记住密码。如果无法登录，请检查微信是否正常，或重新授权登录。"
          },
          {
            id: 3,
            question: "如何更换绑定的微信账号？",
            answer: "目前暂不支持更换绑定的微信账号。如需使用其他微信账号，请联系客服处理。"
          }
        ]
      },
      {
        id: 2,
        title: "打卡与记录",
        icon: "✅",
        iconClass: "checkin-icon",
        expanded: false,
        faqs: [
          {
            id: 4,
            question: "如何进行每日打卡？",
            answer: '在首页或打卡页面点击"今日打卡"按钮，选择当前心情状态（1-5级），添加备注（可选），然后确认打卡。',
            steps: [
              "进入打卡页面",
              "选择心情等级（1-5）",
              "添加今日感想（可选）",
              "点击确认打卡"
            ]
          },
          {
            id: 5,
            question: "错过了打卡时间怎么办？",
            answer: "如果错过了当天打卡，连续天数会重置为0。建议设置打卡提醒，避免错过打卡时间。"
          },
          {
            id: 6,
            question: "可以补签之前的打卡记录吗？",
            answer: "为了保证记录的真实性，暂不支持补签功能。请坚持每日打卡，养成良好习惯。"
          }
        ]
      },
      {
        id: 3,
        title: "成就与等级",
        icon: "🏆",
        iconClass: "achievement-icon",
        expanded: false,
        faqs: [
          {
            id: 7,
            question: "如何获得成就徽章？",
            answer: "通过完成特定条件可以获得成就徽章，如连续打卡、帮助他人、学习内容等。成就会自动解锁。"
          },
          {
            id: 8,
            question: "等级是如何计算的？",
            answer: "等级基于经验值计算，通过打卡、获得成就、社区互动等方式可以获得经验值。等级越高，解锁的功能越多。"
          },
          {
            id: 9,
            question: "成就徽章有什么用？",
            answer: "成就徽章是对您坚持努力的认可，同时也能在社区中展示您的成长历程，激励自己和他人。"
          }
        ]
      },
      {
        id: 4,
        title: "社区与互助",
        icon: "👥",
        iconClass: "community-icon",
        expanded: false,
        faqs: [
          {
            id: 10,
            question: "如何在社区发帖？",
            answer: '进入社区页面，点击右下角的"+"按钮，选择帖子分类，填写标题和内容，选择是否匿名发布。'
          },
          {
            id: 11,
            question: "社区有哪些发帖规则？",
            answer: "请遵守社区规则：内容积极正面、禁止发布不当内容、尊重他人、分享真实经验。违规内容会被删除。"
          },
          {
            id: 12,
            question: "如何举报不当内容？",
            answer: '在帖子或评论右上角点击"..."菜单，选择"举报"，选择举报原因并提交。我们会及时处理。'
          }
        ]
      }
    ]);
    common_vendor.onMounted(() => {
    });
    const goBack = () => {
      common_vendor.index.navigateBack();
    };
    const onSearchInput = () => {
      console.log("搜索内容:", searchText.value);
    };
    const contactSupport = () => {
      common_vendor.index.showModal({
        title: "联系客服",
        content: "请发送邮件至 support@nofap-helper.com 或在社区发帖寻求帮助",
        showCancel: false,
        confirmText: "我知道了"
      });
    };
    const reportBug = () => {
      common_vendor.index.showModal({
        title: "反馈问题",
        content: "请详细描述遇到的问题，并发送至 bug@nofap-helper.com",
        showCancel: false,
        confirmText: "我知道了"
      });
    };
    const suggestFeature = () => {
      common_vendor.index.showModal({
        title: "功能建议",
        content: "欢迎分享您的想法！请发送建议至 feature@nofap-helper.com",
        showCancel: false,
        confirmText: "我知道了"
      });
    };
    const toggleCategory = (categoryId) => {
      const category = faqCategories.value.find((c) => c.id === categoryId);
      if (category) {
        category.expanded = !category.expanded;
      }
    };
    const showFaqDetail = (faq) => {
      selectedFaq.value = faq;
      showDetailModal.value = true;
    };
    const closeDetailModal = () => {
      showDetailModal.value = false;
      selectedFaq.value = null;
    };
    const markHelpful = () => {
      common_vendor.index.showToast({
        title: "感谢反馈",
        icon: "success"
      });
      closeDetailModal();
    };
    const markNotHelpful = () => {
      common_vendor.index.showToast({
        title: "我们会改进",
        icon: "success"
      });
      closeDetailModal();
    };
    return (_ctx, _cache) => {
      var _a, _b, _c, _d;
      return common_vendor.e({
        a: common_vendor.o(goBack),
        b: common_vendor.o([($event) => searchText.value = $event.detail.value, onSearchInput]),
        c: searchText.value,
        d: common_vendor.o(contactSupport),
        e: common_vendor.o(reportBug),
        f: common_vendor.o(suggestFeature),
        g: common_vendor.f(faqCategories.value, (category, k0, i0) => {
          return common_vendor.e({
            a: common_vendor.t(category.icon),
            b: common_vendor.n(category.iconClass),
            c: common_vendor.t(category.title),
            d: common_vendor.t(category.expanded ? "−" : "+"),
            e: category.expanded ? 1 : "",
            f: category.expanded
          }, category.expanded ? {
            g: common_vendor.f(category.faqs, (faq, k1, i1) => {
              return {
                a: common_vendor.t(faq.question),
                b: faq.id,
                c: common_vendor.o(($event) => showFaqDetail(faq), faq.id)
              };
            })
          } : {}, {
            h: category.id,
            i: common_vendor.o(($event) => toggleCategory(category.id), category.id)
          });
        }),
        h: showDetailModal.value
      }, showDetailModal.value ? common_vendor.e({
        i: common_vendor.t((_a = selectedFaq.value) == null ? void 0 : _a.question),
        j: common_vendor.o(closeDetailModal),
        k: common_vendor.t((_b = selectedFaq.value) == null ? void 0 : _b.answer),
        l: (_c = selectedFaq.value) == null ? void 0 : _c.steps
      }, ((_d = selectedFaq.value) == null ? void 0 : _d.steps) ? {
        m: common_vendor.f(selectedFaq.value.steps, (step, index, i0) => {
          return {
            a: common_vendor.t(index + 1),
            b: common_vendor.t(step),
            c: index
          };
        })
      } : {}, {
        n: common_vendor.o(markHelpful),
        o: common_vendor.o(markNotHelpful),
        p: common_vendor.o(() => {
        }),
        q: common_vendor.o(closeDetailModal)
      }) : {}, {
        r: loading.value
      }, loading.value ? {
        s: common_vendor.t(loadingText.value)
      } : {});
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-8059907d"]]);
wx.createPage(MiniProgramPage);
