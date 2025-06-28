/**
 * 学习内容模拟数据
 */

// 精品课程
export const featuredCourse = {
  id: 1,
  title: '21天自控力训练营',
  summary: '科学方法帮你建立持久的自控习惯',
  thumbnailUrl: 'https://images.unsplash.com/photo-1499336315816-097655dcfbda?w=160&h=160&fit=crop',
  category: 2,
  difficulty: 2,
  contentType: 1,
  duration: 30,
  viewCount: 5200,
  likeCount: 450,
  collectCount: 320,
  commentCount: 89,
  author: '心理学专家',
  tags: '自控力,习惯养成,心理训练',
  tagList: ['自控力', '习惯养成', '心理训练']
}

// 学习统计模拟数据
export const mockLearningStats = {
  totalContents: 25,
  completedContents: 12,
  likedContents: 8,
  collectedContents: 5,
  totalLearningTime: 210, // 分钟
  avgLearningTime: 18,
  completionRate: 48,
  continuousLearning: 7,
  lastLearningTime: new Date().toISOString()
}

// 文章模拟数据
export const mockArticles = [
  {
    id: 1,
    title: '如何建立健康的生活习惯',
    summary: '从小习惯开始，逐步建立健康的生活方式，包括规律作息、健康饮食、适量运动等方面的具体建议...',
    content: '详细的文章内容...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=160&h=120&fit=crop',
    contentType: 1,
    category: 2,
    difficulty: 1,
    duration: 8,
    author: '健康专家',
    viewCount: 1200,
    likeCount: 234,
    collectCount: 89,
    commentCount: 45,
    status: 1,
    tags: '健康习惯,生活方式,作息规律',
    tagList: ['健康习惯', '生活方式', '作息规律'],
    createdAt: '2024-06-20T10:00:00Z',
    isLiked: false,
    isCollected: false,
    userRating: 0
  },
  {
    id: 2,
    title: '理解大脑的奖励机制',
    summary: '科学解释成瘾的神经学原理和应对方法，帮助你了解大脑如何工作，以及如何利用这些知识改变行为...',
    content: '详细的文章内容...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1559757148-5c350d0d3c56?w=160&h=120&fit=crop',
    contentType: 1,
    category: 1,
    difficulty: 2,
    duration: 12,
    author: '神经科学家',
    viewCount: 856,
    likeCount: 189,
    collectCount: 67,
    commentCount: 23,
    status: 1,
    tags: '神经科学,大脑机制,行为改变',
    tagList: ['神经科学', '大脑机制', '行为改变'],
    createdAt: '2024-06-19T14:30:00Z',
    isLiked: true,
    isCollected: false,
    userRating: 5
  },
  {
    id: 3,
    title: '正念冥想入门指南',
    summary: '学会用冥想管理情绪和冲动，建立内心平静。包含基础冥想技巧、呼吸方法和日常练习建议...',
    content: '详细的文章内容...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1545389336-cf090694435e?w=160&h=120&fit=crop',
    contentType: 1,
    category: 3,
    difficulty: 1,
    duration: 15,
    author: '冥想导师',
    viewCount: 743,
    likeCount: 156,
    collectCount: 98,
    commentCount: 34,
    status: 1,
    tags: '正念,冥想,情绪管理',
    tagList: ['正念', '冥想', '情绪管理'],
    createdAt: '2024-06-18T09:15:00Z',
    isLiked: false,
    isCollected: true,
    userRating: 4
  },
  {
    id: 8,
    title: '压力管理的科学方法',
    summary: '基于科学研究的压力应对策略，包括认知重构、时间管理、放松技巧等实用方法...',
    content: '详细的文章内容...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=160&h=120&fit=crop',
    contentType: 1,
    category: 2,
    difficulty: 2,
    duration: 10,
    author: '心理咨询师',
    viewCount: 1050,
    likeCount: 203,
    collectCount: 76,
    commentCount: 28,
    status: 1,
    tags: '压力管理,心理健康,应对策略',
    tagList: ['压力管理', '心理健康', '应对策略'],
    createdAt: '2024-06-17T16:20:00Z',
    isLiked: false,
    isCollected: false,
    userRating: 0
  }
]

// 视频模拟数据
export const mockVideos = [
  {
    id: 4,
    title: '正念冥想实践指导',
    summary: '跟随专业导师学习正念冥想的基础技巧，通过视频演示掌握正确的冥想姿势和呼吸方法',
    content: '视频内容描述...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1545389336-cf090694435e?w=300&h=200&fit=crop',
    contentType: 2,
    category: 3,
    difficulty: 1,
    duration: 15,
    author: '冥想导师',
    viewCount: 3200,
    likeCount: 127,
    collectCount: 89,
    commentCount: 23,
    status: 1,
    tags: '正念,冥想,实践指导',
    tagList: ['正念', '冥想', '实践指导'],
    mediaUrl: 'https://example.com/video/meditation.mp4',
    createdAt: '2024-06-17T16:45:00Z',
    isLiked: true,
    isCollected: false,
    userRating: 5
  },
  {
    id: 5,
    title: '呼吸练习技巧讲解',
    summary: '通过正确的呼吸方法缓解压力和焦虑，学习腹式呼吸、4-7-8呼吸法等实用技巧',
    content: '视频内容描述...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300&h=200&fit=crop',
    contentType: 2,
    category: 2,
    difficulty: 1,
    duration: 10,
    author: '呼吸治疗师',
    viewCount: 1800,
    likeCount: 89,
    collectCount: 45,
    commentCount: 15,
    status: 1,
    tags: '呼吸练习,压力缓解,焦虑管理',
    tagList: ['呼吸练习', '压力缓解', '焦虑管理'],
    mediaUrl: 'https://example.com/video/breathing.mp4',
    createdAt: '2024-06-16T11:20:00Z',
    isLiked: false,
    isCollected: true,
    userRating: 4
  },
  {
    id: 9,
    title: '运动与心理健康',
    summary: '了解运动对心理健康的积极影响，学习如何通过适量运动改善情绪和减少压力',
    content: '视频内容描述...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=300&h=200&fit=crop',
    contentType: 2,
    category: 2,
    difficulty: 1,
    duration: 20,
    author: '运动心理学家',
    viewCount: 2100,
    likeCount: 156,
    collectCount: 78,
    commentCount: 32,
    status: 1,
    tags: '运动,心理健康,情绪管理',
    tagList: ['运动', '心理健康', '情绪管理'],
    mediaUrl: 'https://example.com/video/exercise.mp4',
    createdAt: '2024-06-15T14:30:00Z',
    isLiked: false,
    isCollected: false,
    userRating: 0
  }
]

// 音频模拟数据
export const mockAudios = [
  {
    id: 6,
    title: '睡前放松冥想',
    summary: '帮助改善睡眠质量的引导冥想，通过温和的声音指导让你放松身心，进入深度睡眠',
    content: '音频内容描述...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1542662565-7e4b16f20bfb?w=160&h=120&fit=crop',
    contentType: 3,
    category: 3,
    difficulty: 1,
    duration: 20,
    author: '睡眠专家',
    viewCount: 890,
    likeCount: 67,
    collectCount: 123,
    commentCount: 18,
    status: 1,
    tags: '睡眠,放松,冥想',
    tagList: ['睡眠', '放松', '冥想'],
    mediaUrl: 'https://example.com/audio/sleep-meditation.mp3',
    createdAt: '2024-06-15T20:00:00Z',
    isLiked: false,
    isCollected: true,
    userRating: 5,
    isDownloaded: true
  },
  {
    id: 7,
    title: '专注力训练音频',
    summary: '提升注意力和专注能力的训练课程，通过特殊的音频频率和引导语帮助你提高专注力',
    content: '音频内容描述...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1499209974431-9dddcece7f88?w=160&h=120&fit=crop',
    contentType: 3,
    category: 2,
    difficulty: 2,
    duration: 25,
    author: '专注力教练',
    viewCount: 612,
    likeCount: 45,
    collectCount: 67,
    commentCount: 12,
    status: 1,
    tags: '专注力,注意力训练,效率提升',
    tagList: ['专注力', '注意力训练', '效率提升'],
    mediaUrl: 'https://example.com/audio/focus-training.mp3',
    createdAt: '2024-06-14T08:30:00Z',
    isLiked: true,
    isCollected: false,
    userRating: 4,
    isDownloaded: false
  },
  {
    id: 10,
    title: '情绪调节引导音频',
    summary: '学习如何识别和调节负面情绪，通过专业的心理技巧帮助你保持情绪平衡',
    content: '音频内容描述...',
    thumbnailUrl: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=160&h=120&fit=crop',
    contentType: 3,
    category: 1,
    difficulty: 2,
    duration: 18,
    author: '情绪管理师',
    viewCount: 756,
    likeCount: 89,
    collectCount: 54,
    commentCount: 21,
    status: 1,
    tags: '情绪调节,心理健康,自我管理',
    tagList: ['情绪调节', '心理健康', '自我管理'],
    mediaUrl: 'https://example.com/audio/emotion-regulation.mp3',
    createdAt: '2024-06-13T15:45:00Z',
    isLiked: false,
    isCollected: false,
    userRating: 0,
    isDownloaded: false
  }
]

// 分类统计模拟数据
export const mockCategoryStats = [
  {
    category: 1,
    categoryName: '心理健康',
    totalContents: 8,
    completedContents: 3,
    completionRate: 38
  },
  {
    category: 2,
    categoryName: '生活方式',
    totalContents: 12,
    completedContents: 7,
    completionRate: 58
  },
  {
    category: 3,
    categoryName: '正念冥想',
    totalContents: 5,
    completedContents: 2,
    completionRate: 40
  }
]

// 推荐内容模拟数据
export const mockRecommendations = [
  {
    recommendType: 'continue',
    reason: '继续上次未完成的学习',
    contents: [mockArticles[0]]
  },
  {
    recommendType: 'similar',
    reason: '基于你的兴趣推荐',
    contents: [mockVideos[0], mockAudios[0]]
  },
  {
    recommendType: 'trending',
    reason: '热门内容推荐',
    contents: [mockArticles[1], mockVideos[1]]
  }
]

// 获取所有内容（合并文章、视频、音频）
export const getAllMockContents = () => {
  return [...mockArticles, ...mockVideos, ...mockAudios].sort((a, b) => 
    new Date(b.createdAt) - new Date(a.createdAt)
  )
}

// 根据内容类型过滤
export const getContentsByType = (contentType) => {
  const allContents = getAllMockContents()
  if (!contentType) return allContents
  return allContents.filter(content => content.contentType === contentType)
}

// 搜索内容
export const searchMockContents = (keyword) => {
  const allContents = getAllMockContents()
  if (!keyword) return []
  
  const lowerKeyword = keyword.toLowerCase()
  return allContents.filter(content => 
    content.title.toLowerCase().includes(lowerKeyword) ||
    content.summary.toLowerCase().includes(lowerKeyword) ||
    content.tags.toLowerCase().includes(lowerKeyword)
  )
}

// 分页获取内容
export const getPaginatedContents = (contents, page = 1, pageSize = 10) => {
  const startIndex = (page - 1) * pageSize
  const endIndex = startIndex + pageSize
  
  return {
    list: contents.slice(startIndex, endIndex),
    total: contents.length,
    page,
    pageSize,
    hasMore: endIndex < contents.length
  }
} 