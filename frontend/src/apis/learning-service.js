/**
 * 学习内容服务层
 * 提供学习相关功能的统一接口，包含API调用和模拟数据回退机制
 */
import learningApi from './learning.js'
import { 
  mockLearningStats,
  mockArticles,
  mockVideos,
  mockAudios,
  getAllMockContents,
  getContentsByType,
  searchMockContents,
  getPaginatedContents
} from '@/data/learning-mock'

class LearningService {
  /**
   * 获取学习统计数据
   */
  async getStats() {
    try {
      const res = await learningApi.getStats()
      if (res.code === 0) {
        return {
          success: true,
          data: res.data
        }
      }
      throw new Error('API返回错误')
    } catch (error) {
      console.warn('获取学习统计失败，使用模拟数据:', error)
      return {
        success: true,
        data: mockLearningStats,
        isMock: true
      }
    }
  }

  /**
   * 获取学习内容列表
   * 注意：后端API暂未实现，直接使用模拟数据
   */
  async getContents(params = {}) {
    // 暂时直接使用模拟数据，因为后端API未实现
    console.warn('学习内容列表API暂未实现，使用模拟数据')
    
    const { page = 1, pageSize = 10, contentType } = params
    const mockContents = getContentsByType(contentType)
    const paginatedData = getPaginatedContents(mockContents, page, pageSize)
    
    return {
      success: true,
      data: paginatedData,
      isMock: true
    }
  }

  /**
   * 搜索学习内容
   * 注意：后端API暂未实现，直接使用模拟数据
   */
  async searchContents(params = {}) {
    // 暂时直接使用模拟数据，因为后端API未实现
    console.warn('搜索学习内容API暂未实现，使用模拟数据')
    
    const { keyword, page = 1, pageSize = 10 } = params
    const searchResults = searchMockContents(keyword || '')
    const paginatedData = getPaginatedContents(searchResults, page, pageSize)
    
    return {
      success: true,
      data: paginatedData,
      isMock: true
    }
  }

  /**
   * 开始学习记录
   * 注意：后端API暂未实现，返回模拟成功
   */
  async startLearning(contentId) {
    // 暂时直接返回成功，因为后端API未实现
    console.warn('开始学习记录API暂未实现，返回模拟成功')
    
    return {
      success: true,
      data: { 
        message: '学习记录已开始',
        contentId: contentId,
        startTime: new Date().toISOString()
      },
      isMock: true
    }
  }

  /**
   * 完成学习
   */
  async completeLearning(contentId, data = {}) {
    try {
      const res = await learningApi.completeLearning(contentId, data)
      return {
        success: res.code === 0,
        data: res.data
      }
    } catch (error) {
      console.warn('完成学习记录失败:', error)
      return {
        success: true,
        data: { message: '学习记录失败，但不影响使用' },
        isMock: true
      }
    }
  }

  /**
   * 点赞内容
   */
  async likeContent(contentId) {
    try {
      const res = await learningApi.likeContent(contentId)
      return {
        success: res.code === 0,
        data: res.data
      }
    } catch (error) {
      console.warn('点赞失败:', error)
      return {
        success: false,
        error: '点赞失败，请稍后重试'
      }
    }
  }

  /**
   * 收藏内容
   */
  async collectContent(contentId) {
    try {
      const res = await learningApi.collectContent(contentId)
      return {
        success: res.code === 0,
        data: res.data
      }
    } catch (error) {
      console.warn('收藏失败:', error)
      return {
        success: false,
        error: '收藏失败，请稍后重试'
      }
    }
  }

  /**
   * 获取内容详情
   */
  async getContentDetail(id) {
    try {
      const res = await learningApi.getDetail(id)
      if (res.code === 0) {
        return {
          success: true,
          data: res.data
        }
      }
      throw new Error('API返回错误')
    } catch (error) {
      console.warn('获取内容详情失败，使用模拟数据:', error)
      
      // 从模拟数据中查找
      const allContents = getAllMockContents()
      const content = allContents.find(item => item.id === parseInt(id))
      
      if (content) {
        return {
          success: true,
          data: content,
          isMock: true
        }
      } else {
        return {
          success: false,
          error: '内容不存在'
        }
      }
    }
  }

  /**
   * 转换学习统计数据格式
   */
  transformStatsForDisplay(stats) {
    return {
      weeklyHours: (stats.totalLearningTime / 60).toFixed(1),
      weeklyGoal: 5,
      weeklyProgress: Math.min((stats.totalLearningTime / 300) * 100, 100),
      articlesRead: stats.completedContents || 0,
      videosWatched: stats.likedContents || 0,
      audiosListened: stats.collectedContents || 0
    }
  }

  /**
   * 根据内容类型获取对应的模拟数据
   */
  getMockDataByType(contentType) {
    switch (contentType) {
      case 1:
        return mockArticles
      case 2:
        return mockVideos
      case 3:
        return mockAudios
      default:
        return getAllMockContents()
    }
  }
}

// 导出单例实例
export const learningService = new LearningService()
export default learningService 