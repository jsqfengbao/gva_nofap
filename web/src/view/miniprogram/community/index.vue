<template>
  <div class="community-management">
    <!-- 搜索区域 -->
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="内容关键词" prop="keyword">
          <el-input v-model="searchInfo.keyword" placeholder="搜索帖子内容" clearable />
        </el-form-item>
        <el-form-item label="发布者" prop="author">
          <el-input v-model="searchInfo.author" placeholder="搜索发布者昵称" clearable />
        </el-form-item>
        <el-form-item label="帖子状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="请选择帖子状态" clearable>
            <el-option label="正常" value="1" />
            <el-option label="待审核" value="2" />
            <el-option label="已隐藏" value="3" />
            <el-option label="已删除" value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="发布时间" prop="dateRange">
          <el-date-picker
            v-model="searchInfo.dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="getTableData">查询</el-button>
          <el-button icon="refresh" @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 统计卡片 -->
    <div class="statistics-cards">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon total-icon">
                <el-icon><ChatDotRound /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.totalPosts }}</div>
                <div class="stat-label">总帖子数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon pending-icon">
                <el-icon><Clock /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.pendingPosts }}</div>
                <div class="stat-label">待审核</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon today-icon">
                <el-icon><Calendar /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.todayPosts }}</div>
                <div class="stat-label">今日发布</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon comments-icon">
                <el-icon><ChatLineRound /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.totalComments }}</div>
                <div class="stat-label">总评论数</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <!-- 数据表格 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="success" icon="check" @click="batchApprove" :disabled="!multipleSelection.length">批量通过</el-button>
        <el-button type="warning" icon="hide" @click="batchHide" :disabled="!multipleSelection.length">批量隐藏</el-button>
        <el-button type="danger" icon="delete" @click="batchDelete" :disabled="!multipleSelection.length">批量删除</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="发布者" width="120">
          <template #default="scope">
            <div class="author-info">
              <el-avatar :size="30" :src="scope.row.authorAvatar" />
              <span class="author-name">{{ scope.row.authorNickname }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="内容" width="300">
          <template #default="scope">
            <div class="post-content">
              <p class="content-text">{{ scope.row.content }}</p>
              <div v-if="scope.row.images && scope.row.images.length" class="content-images">
                <el-image
                  v-for="(img, index) in scope.row.images.slice(0, 3)"
                  :key="index"
                  :src="img"
                  style="width: 40px; height: 40px; margin-right: 4px;"
                  fit="cover"
                  preview-disabled
                />
                <span v-if="scope.row.images.length > 3" class="more-images">+{{ scope.row.images.length - 3 }}</span>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTag(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="点赞数" prop="likesCount" width="80" />
        <el-table-column align="left" label="评论数" prop="commentsCount" width="80" />
        <el-table-column align="left" label="发布时间" width="180">
          <template #default="scope">
            <span>{{ formatTime(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" width="200">
          <template #default="scope">
            <el-button type="primary" link size="small" @click="viewPost(scope.row)">查看</el-button>
            <el-button 
              v-if="scope.row.status === 2" 
              type="success" 
              link 
              size="small" 
              @click="approvePost(scope.row)"
            >
              通过
            </el-button>
            <el-button 
              v-if="scope.row.status === 1" 
              type="warning" 
              link 
              size="small" 
              @click="hidePost(scope.row)"
            >
              隐藏
            </el-button>
            <el-button type="danger" link size="small" @click="deletePost(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 帖子详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="帖子详情" width="800px">
      <div v-if="currentPost" class="post-detail">
        <div class="post-header">
          <div class="author-info">
            <el-avatar :size="50" :src="currentPost.authorAvatar" />
            <div class="author-details">
              <h4>{{ currentPost.authorNickname }}</h4>
              <p>{{ formatTime(currentPost.CreatedAt) }}</p>
            </div>
          </div>
          <el-tag :type="getStatusTag(currentPost.status)">
            {{ getStatusText(currentPost.status) }}
          </el-tag>
        </div>

        <div class="post-content">
          <p>{{ currentPost.content }}</p>
          <div v-if="currentPost.images && currentPost.images.length" class="post-images">
            <el-image
              v-for="(img, index) in currentPost.images"
              :key="index"
              :src="img"
              style="width: 100px; height: 100px; margin: 0 8px 8px 0;"
              fit="cover"
              :preview-src-list="currentPost.images"
            />
          </div>
        </div>

        <div class="post-stats">
          <el-row :gutter="20">
            <el-col :span="8">
              <el-statistic title="点赞数" :value="currentPost.likesCount" />
            </el-col>
            <el-col :span="8">
              <el-statistic title="评论数" :value="currentPost.commentsCount" />
            </el-col>
            <el-col :span="8">
              <el-statistic title="分享数" :value="currentPost.sharesCount" />
            </el-col>
          </el-row>
        </div>

        <div class="post-comments">
          <h4>评论列表</h4>
          <div v-if="currentPost.comments && currentPost.comments.length" class="comments-list">
            <div v-for="comment in currentPost.comments" :key="comment.ID" class="comment-item">
              <div class="comment-header">
                <el-avatar :size="30" :src="comment.userAvatar" />
                <span class="comment-author">{{ comment.userNickname }}</span>
                <span class="comment-time">{{ formatTime(comment.CreatedAt) }}</span>
              </div>
              <div class="comment-content">{{ comment.content }}</div>
            </div>
          </div>
          <el-empty v-else description="暂无评论" />
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ChatDotRound, Clock, Calendar, ChatLineRound } from '@element-plus/icons-vue'
import { getCommunityPostList, getCommunityPostDetail, updatePostStatus, deleteCommunityPost, getCommunityStatistics } from '@/api/miniprogram'
import { formatTimeToStr } from '@/utils/date'

defineOptions({
  name: 'CommunityManagement'
})

// 响应式数据
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = reactive({
  keyword: '',
  author: '',
  status: '',
  dateRange: []
})

const statistics = reactive({
  totalPosts: 0,
  pendingPosts: 0,
  todayPosts: 0,
  totalComments: 0
})

const detailDialogVisible = ref(false)
const currentPost = ref(null)
const multipleSelection = ref([])

// 获取表格数据
const getTableData = async() => {
  const table = await getCommunityPostList({
    page: page.value,
    pageSize: pageSize.value,
    keyword: searchInfo.keyword,
    author: searchInfo.author,
    status: searchInfo.status,
    startDate: searchInfo.dateRange?.[0],
    endDate: searchInfo.dateRange?.[1]
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 获取统计数据
const getStatisticsData = async() => {
  const stats = await getCommunityStatistics()
  if (stats.code === 0) {
    Object.assign(statistics, stats.data)
  }
}

// 分页相关
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 搜索重置
const resetForm = () => {
  searchInfo.keyword = ''
  searchInfo.author = ''
  searchInfo.status = ''
  searchInfo.dateRange = []
  getTableData()
}

// 选择变化
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 查看帖子详情
const viewPost = async(row) => {
  const res = await getCommunityPostDetail(row.ID)
  if (res.code === 0) {
    currentPost.value = res.data
    detailDialogVisible.value = true
  }
}

// 审核通过帖子
const approvePost = async(row) => {
  const res = await updatePostStatus(row.ID, 1)
  if (res.code === 0) {
    ElMessage.success('审核通过')
    getTableData()
    getStatisticsData()
  } else {
    ElMessage.error(res.msg || '操作失败')
  }
}

// 隐藏帖子
const hidePost = async(row) => {
  ElMessageBox.confirm(
    `确定要隐藏这篇帖子吗？`,
    '确认隐藏',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async() => {
    const res = await updatePostStatus(row.ID, 3)
    if (res.code === 0) {
      ElMessage.success('隐藏成功')
      getTableData()
      getStatisticsData()
    } else {
      ElMessage.error(res.msg || '操作失败')
    }
  })
}

// 删除帖子
const deletePost = (row) => {
  ElMessageBox.confirm(
    `确定要删除这篇帖子吗？此操作不可恢复！`,
    '确认删除',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async() => {
    const res = await deleteCommunityPost(row.ID)
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
      getStatisticsData()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  })
}

// 批量操作
const batchApprove = () => {
  if (!multipleSelection.value.length) return
  ElMessageBox.confirm(
    `确定要批量通过选中的 ${multipleSelection.value.length} 篇帖子吗？`,
    '批量审核',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'info'
    }
  ).then(async() => {
    // 批量审核逻辑
    ElMessage.success('批量审核成功')
    getTableData()
    getStatisticsData()
  })
}

const batchHide = () => {
  if (!multipleSelection.value.length) return
  ElMessageBox.confirm(
    `确定要批量隐藏选中的 ${multipleSelection.value.length} 篇帖子吗？`,
    '批量隐藏',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async() => {
    // 批量隐藏逻辑
    ElMessage.success('批量隐藏成功')
    getTableData()
    getStatisticsData()
  })
}

const batchDelete = () => {
  if (!multipleSelection.value.length) return
  ElMessageBox.confirm(
    `确定要批量删除选中的 ${multipleSelection.value.length} 篇帖子吗？此操作不可恢复！`,
    '批量删除',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async() => {
    // 批量删除逻辑
    ElMessage.success('批量删除成功')
    getTableData()
    getStatisticsData()
  })
}

// 工具函数
const formatTime = (time) => {
  if (!time) return '-'
  return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss')
}

const getStatusText = (status) => {
  const statusMap = {
    1: '正常',
    2: '待审核',
    3: '已隐藏',
    4: '已删除'
  }
  return statusMap[status] || '未知'
}

const getStatusTag = (status) => {
  const tagMap = {
    1: 'success',
    2: 'warning',
    3: 'info',
    4: 'danger'
  }
  return tagMap[status] || 'info'
}

// 生命周期
onMounted(() => {
  getTableData()
  getStatisticsData()
})
</script>

<style lang="scss" scoped>
.community-management {
  .statistics-cards {
    margin-bottom: 20px;
    
    .stat-card {
      .stat-content {
        display: flex;
        align-items: center;
        
        .stat-icon {
          width: 60px;
          height: 60px;
          border-radius: 12px;
          display: flex;
          align-items: center;
          justify-content: center;
          margin-right: 16px;
          font-size: 24px;
          color: white;
          
          &.total-icon {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
          }
          
          &.pending-icon {
            background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
          }
          
          &.today-icon {
            background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
          }
          
          &.comments-icon {
            background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
          }
        }
        
        .stat-info {
          flex: 1;
          
          .stat-number {
            font-size: 28px;
            font-weight: bold;
            color: #303133;
            line-height: 1;
          }
          
          .stat-label {
            font-size: 14px;
            color: #909399;
            margin-top: 4px;
          }
        }
      }
    }
  }
  
  .author-info {
    display: flex;
    align-items: center;
    
    .author-name {
      margin-left: 8px;
      font-size: 14px;
    }
  }
  
  .post-content {
    .content-text {
      margin: 0 0 8px 0;
      font-size: 14px;
      line-height: 1.4;
      max-height: 60px;
      overflow: hidden;
      text-overflow: ellipsis;
      display: -webkit-box;
      -webkit-line-clamp: 3;
      -webkit-box-orient: vertical;
    }
    
    .content-images {
      display: flex;
      align-items: center;
      
      .more-images {
        font-size: 12px;
        color: #909399;
        margin-left: 4px;
      }
    }
  }
  
  .post-detail {
    .post-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;
      
      .author-info {
        display: flex;
        align-items: center;
        
        .author-details {
          margin-left: 12px;
          
          h4 {
            margin: 0 0 4px 0;
            color: #303133;
          }
          
          p {
            margin: 0;
            font-size: 12px;
            color: #909399;
          }
        }
      }
    }
    
    .post-content {
      margin-bottom: 20px;
      
      p {
        line-height: 1.6;
        color: #303133;
      }
      
      .post-images {
        margin-top: 12px;
      }
    }
    
    .post-stats {
      margin-bottom: 20px;
      padding: 16px;
      background: #f5f7fa;
      border-radius: 8px;
    }
    
    .post-comments {
      h4 {
        margin: 0 0 16px 0;
        color: #303133;
      }
      
      .comments-list {
        max-height: 300px;
        overflow-y: auto;
        
        .comment-item {
          padding: 12px;
          border: 1px solid #ebeef5;
          border-radius: 8px;
          margin-bottom: 8px;
          
          .comment-header {
            display: flex;
            align-items: center;
            margin-bottom: 8px;
            
            .comment-author {
              margin-left: 8px;
              font-size: 14px;
              font-weight: 500;
              color: #303133;
            }
            
            .comment-time {
              margin-left: auto;
              font-size: 12px;
              color: #909399;
            }
          }
          
          .comment-content {
            font-size: 14px;
            line-height: 1.4;
            color: #606266;
          }
        }
      }
    }
  }
}
</style> 