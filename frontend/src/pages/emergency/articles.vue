<template>
  <view class="articles-page">
    <!-- Header -->
    <view class="header">
      <view class="nav-button" @click="goBack">
        <text class="nav-icon">←</text>
      </view>
      <text class="header-title">励志文章</text>
      <view class="nav-placeholder"></view>
    </view>

    <!-- Search Bar -->
    <view class="search-section">
      <view class="search-bar">
        <text class="search-icon">🔍</text>
        <input 
          class="search-input" 
          placeholder="搜索文章内容..." 
          v-model="searchQuery"
          @input="searchArticles"
        />
      </view>
      
      <!-- Filter Tags -->
      <scroll-view class="filter-tags" scroll-x>
        <view 
          class="tag" 
          :class="{ active: selectedCategory === null }"
          @click="filterByCategory(null)"
        >
          <text class="tag-text">全部</text>
        </view>
        <view 
          v-for="category in categories" 
          :key="category.id"
          class="tag" 
          :class="{ active: selectedCategory === category.id }"
          @click="filterByCategory(category.id)"
        >
          <text class="tag-text">{{ category.name }}</text>
        </view>
      </scroll-view>
    </view>

    <!-- Articles List -->
    <scroll-view class="articles-list" scroll-y @scrolltolower="loadMore">
      <view 
        v-for="article in filteredArticles" 
        :key="article.id"
        class="article-card"
        @click="readArticle(article)"
      >
        <view class="article-header">
          <view class="category-badge">
            <text class="badge-text">{{ getCategoryName(article.category) }}</text>
          </view>
          <text class="reading-time">{{ article.readingTime }}分钟阅读</text>
        </view>
        
        <text class="article-title">{{ article.title }}</text>
        <text class="article-summary">{{ article.summary }}</text>
        
        <view class="article-footer">
          <view class="article-stats">
            <text class="stat-item">👁️ {{ article.views }}</text>
            <text class="stat-item">❤️ {{ article.likes }}</text>
            <text class="stat-item">💬 {{ article.comments }}</text>
          </view>
          <text class="article-date">{{ formatDate(article.createdAt) }}</text>
        </view>
      </view>
      
      <!-- Loading -->
      <view v-if="loading" class="loading">
        <text class="loading-text">加载中...</text>
      </view>
      
      <!-- No Results -->
      <view v-if="!loading && filteredArticles.length === 0" class="no-results">
        <text class="no-results-icon">📖</text>
        <text class="no-results-text">暂无相关文章</text>
        <text class="no-results-subtitle">试试其他关键词或分类</text>
      </view>
    </scroll-view>

    <!-- Article Modal -->
    <view v-if="showArticleModal" class="modal-overlay" @click="closeArticleModal">
      <view class="article-modal" @click.stop>
        <view class="modal-header">
          <view class="modal-header-info">
            <text class="modal-title">{{ selectedArticle.title }}</text>
            <text class="modal-category">{{ getCategoryName(selectedArticle.category) }}</text>
          </view>
          <text class="close-button" @click="closeArticleModal">✕</text>
        </view>
        
        <scroll-view class="modal-content" scroll-y>
          <text class="article-content">{{ selectedArticle.content }}</text>
        </scroll-view>
        
        <view class="modal-footer">
          <view class="action-buttons">
            <view class="action-button like-button" @click="likeArticle">
              <text class="button-icon">❤️</text>
              <text class="button-text">{{ selectedArticle.likes }}</text>
            </view>
            <view class="action-button share-button" @click="shareArticle">
              <text class="button-icon">📤</text>
              <text class="button-text">分享</text>
            </view>
            <view class="action-button collect-button" @click="collectArticle">
              <text class="button-icon">⭐</text>
              <text class="button-text">收藏</text>
            </view>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { ref, computed, onMounted } from 'vue'

export default {
  name: 'EmergencyArticles',
  setup() {
    const searchQuery = ref('')
    const selectedCategory = ref(null)
    const loading = ref(false)
    const showArticleModal = ref(false)
    const selectedArticle = ref({})

    const categories = ref([
      { id: 1, name: '戒色心得' },
      { id: 2, name: '心理调节' },
      { id: 3, name: '健康生活' },
      { id: 4, name: '成功故事' },
      { id: 5, name: '科学解读' }
    ])

    const articles = ref([
      {
        id: 1,
        title: "你比你想象的更坚强",
        summary: "每个人都有克服困难的能力，关键是要相信自己，坚持下去。",
        content: "人生路上，我们总是会遇到各种各样的挑战和困难。有时候，这些困难看起来是那么的巨大，似乎无法克服。但是，我想告诉你，你比你想象的更坚强。\n\n每一次的挫折，都是一次成长的机会。每一次的跌倒，都是为了更好地站起来。在戒色的路上也是如此，虽然会有冲动，会有迷茫，但这些都是正常的。重要的是，我们要学会从失败中学习，从困难中成长。\n\n记住，没有人是完美的，没有人能够一帆风顺。但是，只要我们有坚定的信念，有不放弃的精神，我们就能够克服任何困难。\n\n相信自己，你一定可以做到的！",
        category: 4,
        readingTime: 3,
        views: 1234,
        likes: 89,
        comments: 23,
        createdAt: '2024-01-15'
      },
      {
        id: 2,
        title: "每个选择都是新的开始",
        summary: "不要为过去的错误而自责，重要的是从现在开始做出正确的选择。",
        content: "生活就是由无数个选择组成的。每一天，每一刻，我们都在做着各种各样的选择。有些选择让我们感到自豪，有些选择让我们后悔不已。\n\n在戒色的道路上，我们可能会犯错，可能会复发。这时候，很多人会陷入自责和绝望中。但我想告诉你，过去的选择并不能决定你的未来。每一个新的选择，都是一个新的开始。\n\n重要的不是你跌倒了多少次，而是你爬起来了多少次。每一次的重新开始，都意味着你还有机会，还有希望。\n\n所以，不要被过去束缚，不要被错误定义。从现在开始，做出正确的选择，朝着目标前进。每一个选择，都可能改变你的人生。",
        category: 1,
        readingTime: 4,
        views: 987,
        likes: 67,
        comments: 15,
        createdAt: '2024-01-10'
      },
      {
        id: 3,
        title: "科学认识冲动的本质",
        summary: "了解大脑的工作机制，用科学的方法来应对冲动。",
        content: "冲动是一种自然的生理和心理现象。当我们面临诱惑时，大脑中的奖励系统会被激活，释放多巴胺，让我们产生强烈的欲望。\n\n从神经科学的角度来看，冲动的产生主要涉及大脑的边缘系统，特别是伏隔核和杏仁核。这些区域负责处理奖励和情绪反应。当我们长期接触某种刺激时，大脑会形成神经通路，使得冲动变得更加强烈。\n\n但是，我们的大脑也有前额叶皮质，这是负责理性思考和自我控制的区域。通过训练和练习，我们可以增强前额叶皮质的功能，提高自控能力。\n\n冥想、运动、充足的睡眠都有助于增强大脑的自控能力。当我们了解了冲动的科学本质，就能更好地应对它，而不是被它所控制。",
        category: 5,
        readingTime: 5,
        views: 756,
        likes: 45,
        comments: 8,
        createdAt: '2024-01-08'
      },
      {
        id: 4,
        title: "建立健康的生活方式",
        summary: "培养良好的生活习惯，从身体和心理两个方面提升自己。",
        content: "健康的生活方式是戒色成功的重要基础。当我们的身体和心理都处于良好状态时，我们就有更强的意志力来抵抗诱惑。\n\n**身体健康方面：**\n1. 规律作息：保证充足的睡眠，早睡早起\n2. 适度运动：每天至少30分钟的有氧运动\n3. 均衡饮食：多吃蔬菜水果，少吃加工食品\n4. 充足饮水：每天至少8杯水\n\n**心理健康方面：**\n1. 冥想练习：每天10-20分钟的正念冥想\n2. 社交活动：与朋友家人保持良好关系\n3. 兴趣爱好：培养健康的兴趣爱好\n4. 学习成长：不断学习新知识，提升自己\n\n记住，改变需要时间，不要急于求成。一步一步地建立健康的生活方式，你会发现自己变得越来越强大。",
        category: 3,
        readingTime: 6,
        views: 1100,
        likes: 78,
        comments: 19,
        createdAt: '2024-01-05'
      },
      {
        id: 5,
        title: "如何应对情绪低落",
        summary: "当情绪低落时，学会正确的方法来调节心情，避免冲动行为。",
        content: "情绪低落是每个人都会遇到的情况，在戒色的过程中更是如此。当我们感到沮丧、焦虑或者孤独时，很容易产生冲动行为。\n\n**识别情绪低落的信号：**\n- 感到疲倦和无精打采\n- 对平时喜欢的事情失去兴趣\n- 睡眠质量下降或过度睡眠\n- 食欲不振或暴饮暴食\n- 难以集中注意力\n\n**应对策略：**\n1. **接纳情绪**：不要抗拒或否定负面情绪，接纳它们的存在\n2. **深呼吸练习**：使用4-7-8呼吸法来平静心情\n3. **运动释放**：适度的运动能够释放内啡肽，改善心情\n4. **寻求支持**：与朋友、家人或专业人士交流\n5. **写日记**：记录自己的感受，帮助理清思绪\n6. **正念冥想**：专注于当下，不被负面思绪困扰\n\n记住，情绪低落是暂时的，它会过去。关键是要学会健康的应对方式，而不是逃避到不健康的行为中。",
        category: 2,
        readingTime: 5,
        views: 892,
        likes: 56,
        comments: 12,
        createdAt: '2024-01-03'
      }
    ])

    const filteredArticles = computed(() => {
      let result = articles.value

      // Filter by category
      if (selectedCategory.value !== null) {
        result = result.filter(article => article.category === selectedCategory.value)
      }

      // Filter by search query
      if (searchQuery.value.trim()) {
        const query = searchQuery.value.toLowerCase()
        result = result.filter(article => 
          article.title.toLowerCase().includes(query) ||
          article.summary.toLowerCase().includes(query) ||
          article.content.toLowerCase().includes(query)
        )
      }

      return result
    })

    const getCategoryName = (categoryId) => {
      const category = categories.value.find(cat => cat.id === categoryId)
      return category ? category.name : '未知分类'
    }

    const formatDate = (dateString) => {
      const date = new Date(dateString)
      const now = new Date()
      const diffTime = Math.abs(now - date)
      const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
      
      if (diffDays === 1) {
        return '今天'
      } else if (diffDays === 2) {
        return '昨天'
      } else if (diffDays <= 7) {
        return `${diffDays}天前`
      } else {
        return date.toLocaleDateString('zh-CN')
      }
    }

    const goBack = () => {
      uni.navigateBack()
    }

    const searchArticles = () => {
      // Search functionality is handled by computed property
    }

    const filterByCategory = (categoryId) => {
      selectedCategory.value = categoryId
    }

    const readArticle = (article) => {
      selectedArticle.value = article
      showArticleModal.value = true
      
      // 增加阅读量
      article.views++
    }

    const closeArticleModal = () => {
      showArticleModal.value = false
    }

    const likeArticle = () => {
      selectedArticle.value.likes++
      uni.showToast({
        title: '已点赞',
        icon: 'success',
        duration: 1000
      })
    }

    const shareArticle = () => {
      uni.showActionSheet({
        itemList: ['分享到微信', '分享到朋友圈', '复制链接'],
        success: (res) => {
          uni.showToast({
            title: '分享成功',
            icon: 'success'
          })
        }
      })
    }

    const collectArticle = () => {
      uni.showToast({
        title: '已收藏',
        icon: 'success',
        duration: 1000
      })
    }

    const loadMore = () => {
      if (loading.value) return
      
      loading.value = true
      
      // 模拟加载更多数据
      setTimeout(() => {
        loading.value = false
      }, 1000)
    }

    onMounted(() => {
      // 初始化时可以加载文章数据
    })

    return {
      searchQuery,
      selectedCategory,
      loading,
      showArticleModal,
      selectedArticle,
      categories,
      filteredArticles,
      getCategoryName,
      formatDate,
      goBack,
      searchArticles,
      filterByCategory,
      readArticle,
      closeArticleModal,
      likeArticle,
      shareArticle,
      collectArticle,
      loadMore
    }
  }
}
</script>

<style scoped>
.articles-page {
  min-height: 100vh;
  background: #F8FAFC;
}

/* Header */
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  background: white;
  border-bottom: 1px solid #E5E7EB;
}

.nav-button {
  width: 40px;
  height: 40px;
  background: #F3F4F6;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.nav-placeholder {
  width: 40px;
  height: 40px;
}

.nav-icon {
  font-size: 16px;
  color: #374151;
}

.header-title {
  font-size: 18px;
  font-weight: bold;
  color: #1F2937;
}

/* Search Section */
.search-section {
  padding: 16px 20px;
  background: white;
  border-bottom: 1px solid #E5E7EB;
}

.search-bar {
  display: flex;
  align-items: center;
  background: #F3F4F6;
  border-radius: 12px;
  padding: 12px 16px;
  margin-bottom: 16px;
}

.search-icon {
  font-size: 16px;
  color: #6B7280;
  margin-right: 8px;
}

.search-input {
  flex: 1;
  font-size: 16px;
  color: #1F2937;
  background: transparent;
  border: none;
  outline: none;
}

.filter-tags {
  white-space: nowrap;
}

.tag {
  display: inline-block;
  padding: 8px 16px;
  background: #F3F4F6;
  border-radius: 20px;
  margin-right: 12px;
  transition: all 0.2s;
}

.tag.active {
  background: #3B82F6;
}

.tag-text {
  font-size: 14px;
  color: #6B7280;
}

.tag.active .tag-text {
  color: white;
}

/* Articles List */
.articles-list {
  flex: 1;
  padding: 0 20px 80px;
}

.article-card {
  background: white;
  border-radius: 16px;
  padding: 20px;
  margin-bottom: 16px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.article-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.category-badge {
  background: #EFF6FF;
  border-radius: 12px;
  padding: 4px 12px;
}

.badge-text {
  font-size: 12px;
  color: #3B82F6;
  font-weight: 500;
}

.reading-time {
  font-size: 12px;
  color: #6B7280;
}

.article-title {
  font-size: 16px;
  font-weight: 600;
  color: #1F2937;
  margin-bottom: 8px;
  display: block;
  line-height: 1.4;
}

.article-summary {
  font-size: 14px;
  color: #6B7280;
  line-height: 1.5;
  margin-bottom: 16px;
  display: block;
}

.article-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.article-stats {
  display: flex;
  gap: 16px;
}

.stat-item {
  font-size: 12px;
  color: #6B7280;
}

.article-date {
  font-size: 12px;
  color: #9CA3AF;
}

/* Loading */
.loading {
  padding: 20px;
  text-align: center;
}

.loading-text {
  color: #6B7280;
  font-size: 14px;
}

/* No Results */
.no-results {
  padding: 60px 20px;
  text-align: center;
}

.no-results-icon {
  font-size: 48px;
  color: #D1D5DB;
  margin-bottom: 16px;
  display: block;
}

.no-results-text {
  font-size: 16px;
  color: #6B7280;
  margin-bottom: 8px;
  display: block;
}

.no-results-subtitle {
  font-size: 14px;
  color: #9CA3AF;
  display: block;
}

/* Modal */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 20px;
}

.article-modal {
  background: white;
  border-radius: 16px;
  width: 100%;
  max-width: 500px;
  max-height: 80vh;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 20px;
  border-bottom: 1px solid #E5E7EB;
}

.modal-header-info {
  flex: 1;
}

.modal-title {
  font-size: 18px;
  font-weight: bold;
  color: #1F2937;
  margin-bottom: 4px;
  display: block;
  line-height: 1.4;
}

.modal-category {
  font-size: 12px;
  color: #3B82F6;
  display: block;
}

.close-button {
  font-size: 20px;
  color: #6B7280;
  padding: 4px;
  margin-left: 16px;
}

.modal-content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}

.article-content {
  font-size: 15px;
  line-height: 1.7;
  color: #374151;
  white-space: pre-line;
}

.modal-footer {
  padding: 20px;
  border-top: 1px solid #E5E7EB;
}

.action-buttons {
  display: flex;
  gap: 12px;
}

.action-button {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  padding: 12px;
  border-radius: 8px;
  background: #F3F4F6;
}

.like-button:active {
  background: #FEE2E2;
}

.share-button:active {
  background: #DBEAFE;
}

.collect-button:active {
  background: #FEF3C7;
}

.button-icon {
  font-size: 16px;
}

.button-text {
  font-size: 14px;
  color: #374151;
  font-weight: 500;
}
</style> 