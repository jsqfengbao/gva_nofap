<template>
  <view class="assessment-index">
    <!-- 自定义导航栏 -->
    <view class="custom-navbar">
      <view class="navbar-content">
        <view class="navbar-title">评估测试</view>
      </view>
    </view>

    <!-- 页面内容 -->
    <view class="content">
      <!-- 页面标题 -->
      <view class="page-header">
        <text class="title">色隐指数评估</text>
        <text class="subtitle">了解自己的现状，制定合适的戒色计划</text>
      </view>

      <!-- 评估卡片 -->
      <view class="assessment-card">
        <view class="card-header">
          <text class="card-title">专业评估测试</text>
          <text class="card-desc">通过科学的问卷评估，了解您当前的色隐程度</text>
        </view>
        
        <view class="card-content">
          <view class="assessment-info">
            <view class="info-item">
              <text class="info-label">测试时间</text>
              <text class="info-value">约 5-10 分钟</text>
            </view>
            <view class="info-item">
              <text class="info-label">题目数量</text>
              <text class="info-value">20 道题</text>
            </view>
            <view class="info-item">
              <text class="info-label">评估维度</text>
              <text class="info-value">心理、行为、认知</text>
            </view>
          </view>
        </view>

        <view class="card-actions">
          <button class="btn-primary" @click="startAssessment">
            开始评估
          </button>
        </view>
      </view>

      <!-- 历史记录 -->
      <view class="history-section">
        <view class="section-header">
          <text class="section-title">历史记录</text>
        </view>
        
        <view class="history-list">
          <view v-if="historyList.length === 0" class="empty-state">
            <text class="empty-text">暂无评估记录</text>
          </view>
          
          <view v-else class="history-item" v-for="item in historyList" :key="item.id">
            <view class="history-content">
              <text class="history-date">{{ formatDate(item.createTime) }}</text>
              <text class="history-score">得分：{{ item.score }}</text>
              <text class="history-level">{{ item.level }}</text>
            </view>
            <view class="history-action">
              <button class="btn-secondary" @click="viewResult(item)">
                查看详情
              </button>
            </view>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// 历史记录数据
const historyList = ref([])

// 页面加载
onMounted(() => {
  loadHistoryList()
})

// 开始评估
const startAssessment = () => {
  uni.navigateTo({
    url: '/pages/assessment/questionnaire'
  })
}

// 查看结果详情
const viewResult = (item: any) => {
  uni.navigateTo({
    url: `/pages/assessment/result?id=${item.id}`
  })
}

// 加载历史记录
const loadHistoryList = () => {
  // TODO: 从服务器加载历史记录
  // 暂时使用模拟数据
  historyList.value = []
}

// 格式化日期
const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
}
</script>

<style lang="scss" scoped>
.assessment-index {
  min-height: 100vh;
  background-color: var(--background);
}

.custom-navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 44px;
  background-color: var(--surface);
  border-bottom: 1px solid var(--border);
  z-index: 100;
  padding-top: env(safe-area-inset-top);

  .navbar-content {
    height: 44px;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0 16px;
  }

  .navbar-title {
    font-size: 18px;
    font-weight: 600;
    color: var(--text-primary);
  }
}

.content {
  padding-top: calc(44px + env(safe-area-inset-top));
  padding: 16px;
  padding-top: calc(44px + env(safe-area-inset-top) + 16px);
}

.page-header {
  text-align: center;
  margin-bottom: 24px;

  .title {
    display: block;
    font-size: 24px;
    font-weight: 700;
    color: var(--text-primary);
    margin-bottom: 8px;
  }

  .subtitle {
    display: block;
    font-size: 14px;
    color: var(--text-secondary);
  }
}

.assessment-card {
  background-color: var(--surface);
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 24px;
  box-shadow: var(--shadow-sm);

  .card-header {
    margin-bottom: 16px;

    .card-title {
      display: block;
      font-size: 18px;
      font-weight: 600;
      color: var(--text-primary);
      margin-bottom: 4px;
    }

    .card-desc {
      display: block;
      font-size: 14px;
      color: var(--text-secondary);
    }
  }

  .card-content {
    margin-bottom: 20px;
  }

  .assessment-info {
    .info-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 8px 0;
      border-bottom: 1px solid var(--border-light);

      &:last-child {
        border-bottom: none;
      }

      .info-label {
        font-size: 14px;
        color: var(--text-secondary);
      }

      .info-value {
        font-size: 14px;
        color: var(--text-primary);
        font-weight: 500;
      }
    }
  }

  .card-actions {
    .btn-primary {
      width: 100%;
      height: 44px;
      background-color: var(--primary);
      color: var(--text-white);
      border: none;
      border-radius: 8px;
      font-size: 16px;
      font-weight: 600;
    }
  }
}

.history-section {
  .section-header {
    margin-bottom: 16px;

    .section-title {
      font-size: 18px;
      font-weight: 600;
      color: var(--text-primary);
    }
  }

  .history-list {
    .empty-state {
      text-align: center;
      padding: 40px 20px;
      background-color: var(--surface);
      border-radius: 12px;

      .empty-text {
        font-size: 14px;
        color: var(--text-muted);
      }
    }

    .history-item {
      background-color: var(--surface);
      border-radius: 8px;
      padding: 16px;
      margin-bottom: 8px;
      display: flex;
      justify-content: space-between;
      align-items: center;

      .history-content {
        flex: 1;

        .history-date {
          display: block;
          font-size: 14px;
          color: var(--text-secondary);
          margin-bottom: 4px;
        }

        .history-score {
          display: block;
          font-size: 16px;
          font-weight: 600;
          color: var(--text-primary);
        }

        .history-level {
          display: block;
          font-size: 12px;
          color: var(--text-muted);
        }
      }

      .history-action {
        .btn-secondary {
          padding: 6px 12px;
          background-color: var(--border-light);
          color: var(--text-primary);
          border: none;
          border-radius: 6px;
          font-size: 12px;
        }
      }
    }
  }
}
</style> 