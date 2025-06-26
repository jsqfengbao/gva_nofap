<template>
  <div class="learning-management">
    <!-- 搜索区域 -->
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="内容标题" prop="title">
          <el-input v-model="searchInfo.title" placeholder="搜索内容标题" clearable />
        </el-form-item>
        <el-form-item label="内容类型" prop="type">
          <el-select v-model="searchInfo.type" placeholder="请选择内容类型" clearable>
            <el-option label="文章" value="article" />
            <el-option label="视频" value="video" />
            <el-option label="音频" value="audio" />
          </el-select>
        </el-form-item>
        <el-form-item label="分类" prop="category">
          <el-select v-model="searchInfo.category" placeholder="请选择分类" clearable>
            <el-option label="科普知识" value="knowledge" />
            <el-option label="康复指导" value="recovery" />
            <el-option label="心理健康" value="mental" />
            <el-option label="经验分享" value="experience" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="请选择状态" clearable>
            <el-option label="已发布" value="published" />
            <el-option label="草稿" value="draft" />
            <el-option label="已下线" value="offline" />
          </el-select>
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
                <el-icon><Document /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.totalContent }}</div>
                <div class="stat-label">总内容数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon published-icon">
                <el-icon><Check /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.publishedContent }}</div>
                <div class="stat-label">已发布</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon draft-icon">
                <el-icon><Edit /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.draftContent }}</div>
                <div class="stat-label">草稿</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon views-icon">
                <el-icon><View /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.totalViews }}</div>
                <div class="stat-label">总浏览量</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <!-- 数据表格 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addContent">新增内容</el-button>
        <el-button type="success" icon="upload" @click="batchPublish" :disabled="!multipleSelection.length">批量发布</el-button>
        <el-button type="warning" icon="download" @click="batchOffline" :disabled="!multipleSelection.length">批量下线</el-button>
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
        <el-table-column align="left" label="封面" width="100">
          <template #default="scope">
            <el-image
              v-if="scope.row.coverImage"
              :src="scope.row.coverImage"
              style="width: 60px; height: 40px;"
              fit="cover"
              :preview-src-list="[scope.row.coverImage]"
            />
            <span v-else class="no-image">无封面</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="标题" prop="title" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="类型" width="100">
          <template #default="scope">
            <el-tag :type="getTypeTag(scope.row.type)">
              {{ getTypeText(scope.row.type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="分类" width="120">
          <template #default="scope">
            <el-tag :type="getCategoryTag(scope.row.category)">
              {{ getCategoryText(scope.row.category) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTag(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="浏览量" prop="viewCount" width="100" />
        <el-table-column align="left" label="点赞数" prop="likeCount" width="100" />
        <el-table-column align="left" label="收藏数" prop="favoriteCount" width="100" />
        <el-table-column align="left" label="时长" width="100">
          <template #default="scope">
            <span v-if="scope.row.duration">{{ formatDuration(scope.row.duration) }}</span>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" width="180">
          <template #default="scope">
            <span>{{ formatTime(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" width="200">
          <template #default="scope">
            <el-button type="primary" link size="small" @click="viewContent(scope.row)">查看</el-button>
            <el-button type="warning" link size="small" @click="editContent(scope.row)">编辑</el-button>
            <el-button 
              v-if="scope.row.status === 'draft'" 
              type="success" 
              link 
              size="small" 
              @click="publishContent(scope.row)"
            >
              发布
            </el-button>
            <el-button 
              v-if="scope.row.status === 'published'" 
              type="info" 
              link 
              size="small" 
              @click="offlineContent(scope.row)"
            >
              下线
            </el-button>
            <el-button type="danger" link size="small" @click="deleteContent(scope.row)">删除</el-button>
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

    <!-- 内容详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="内容详情" width="900px">
      <div v-if="currentContent" class="content-detail">
        <el-row :gutter="20">
          <el-col :span="16">
            <el-descriptions title="基本信息" :column="2" border>
              <el-descriptions-item label="标题">{{ currentContent.title }}</el-descriptions-item>
              <el-descriptions-item label="类型">{{ getTypeText(currentContent.type) }}</el-descriptions-item>
              <el-descriptions-item label="分类">{{ getCategoryText(currentContent.category) }}</el-descriptions-item>
              <el-descriptions-item label="状态">{{ getStatusText(currentContent.status) }}</el-descriptions-item>
              <el-descriptions-item label="时长">{{ formatDuration(currentContent.duration) }}</el-descriptions-item>
              <el-descriptions-item label="难度">{{ getDifficultyText(currentContent.difficulty) }}</el-descriptions-item>
              <el-descriptions-item label="浏览量">{{ currentContent.viewCount }}</el-descriptions-item>
              <el-descriptions-item label="点赞数">{{ currentContent.likeCount }}</el-descriptions-item>
              <el-descriptions-item label="收藏数">{{ currentContent.favoriteCount }}</el-descriptions-item>
              <el-descriptions-item label="评分">{{ currentContent.rating }}/5</el-descriptions-item>
              <el-descriptions-item label="创建时间" :span="2">{{ formatTime(currentContent.CreatedAt) }}</el-descriptions-item>
            </el-descriptions>
          </el-col>
          <el-col :span="8">
            <div class="content-cover">
              <h4>封面图片</h4>
              <el-image
                v-if="currentContent.coverImage"
                :src="currentContent.coverImage"
                style="width: 100%; height: 200px;"
                fit="cover"
                :preview-src-list="[currentContent.coverImage]"
              />
              <div v-else class="no-cover">暂无封面</div>
            </div>
          </el-col>
        </el-row>

        <div class="content-description">
          <h4>内容描述</h4>
          <p>{{ currentContent.description }}</p>
        </div>

        <div class="content-body">
          <h4>内容正文</h4>
          <div class="content-preview" v-html="currentContent.content"></div>
        </div>
      </div>
    </el-dialog>

    <!-- 新增/编辑内容对话框 -->
    <el-dialog v-model="editDialogVisible" :title="isEdit ? '编辑内容' : '新增内容'" width="1000px">
      <el-form ref="editFormRef" :model="editForm" :rules="editRules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="内容标题" prop="title">
              <el-input v-model="editForm.title" placeholder="请输入内容标题" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="内容类型" prop="type">
              <el-select v-model="editForm.type" placeholder="请选择内容类型">
                <el-option label="文章" value="article" />
                <el-option label="视频" value="video" />
                <el-option label="音频" value="audio" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="分类" prop="category">
              <el-select v-model="editForm.category" placeholder="请选择分类">
                <el-option label="科普知识" value="knowledge" />
                <el-option label="康复指导" value="recovery" />
                <el-option label="心理健康" value="mental" />
                <el-option label="经验分享" value="experience" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="难度等级" prop="difficulty">
              <el-select v-model="editForm.difficulty" placeholder="请选择难度等级">
                <el-option label="入门" value="beginner" />
                <el-option label="进阶" value="intermediate" />
                <el-option label="高级" value="advanced" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="时长(分钟)" prop="duration">
              <el-input-number v-model="editForm.duration" :min="1" :max="999" placeholder="内容时长" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-select v-model="editForm.status" placeholder="请选择状态">
                <el-option label="草稿" value="draft" />
                <el-option label="已发布" value="published" />
                <el-option label="已下线" value="offline" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="封面图片" prop="coverImage">
          <el-input v-model="editForm.coverImage" placeholder="请输入封面图片URL" />
        </el-form-item>

        <el-form-item label="内容描述" prop="description">
          <el-input v-model="editForm.description" type="textarea" :rows="3" placeholder="请输入内容描述" />
        </el-form-item>

        <el-form-item label="内容正文" prop="content">
          <el-input v-model="editForm.content" type="textarea" :rows="10" placeholder="请输入内容正文" />
        </el-form-item>

        <el-form-item v-if="editForm.type !== 'article'" label="媒体文件" prop="mediaUrl">
          <el-input v-model="editForm.mediaUrl" placeholder="请输入视频/音频文件URL" />
        </el-form-item>

        <el-form-item label="标签" prop="tags">
          <el-input v-model="editForm.tags" placeholder="请输入标签，用逗号分隔" />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveContent">保存</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Document, Check, Edit, View } from '@element-plus/icons-vue'
import { getLearningContentList, getLearningContentDetail, createLearningContent, updateLearningContent, deleteLearningContent, getLearningStatistics } from '@/api/miniprogram'
import { formatTimeToStr } from '@/utils/date'

defineOptions({
  name: 'LearningManagement'
})

// 响应式数据
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = reactive({
  title: '',
  type: '',
  category: '',
  status: ''
})

const statistics = reactive({
  totalContent: 0,
  publishedContent: 0,
  draftContent: 0,
  totalViews: 0
})

const detailDialogVisible = ref(false)
const editDialogVisible = ref(false)
const currentContent = ref(null)
const multipleSelection = ref([])
const isEdit = ref(false)

// 编辑表单
const editForm = reactive({
  ID: '',
  title: '',
  type: 'article',
  category: 'knowledge',
  difficulty: 'beginner',
  duration: 10,
  status: 'draft',
  coverImage: '',
  description: '',
  content: '',
  mediaUrl: '',
  tags: ''
})

const editRules = {
  title: [
    { required: true, message: '请输入内容标题', trigger: 'blur' }
  ],
  type: [
    { required: true, message: '请选择内容类型', trigger: 'change' }
  ],
  category: [
    { required: true, message: '请选择分类', trigger: 'change' }
  ],
  description: [
    { required: true, message: '请输入内容描述', trigger: 'blur' }
  ],
  content: [
    { required: true, message: '请输入内容正文', trigger: 'blur' }
  ]
}

// 获取表格数据
const getTableData = async() => {
  const table = await getLearningContentList({
    page: page.value,
    pageSize: pageSize.value,
    title: searchInfo.title,
    type: searchInfo.type,
    category: searchInfo.category,
    status: searchInfo.status
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
  const stats = await getLearningStatistics()
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
  searchInfo.title = ''
  searchInfo.type = ''
  searchInfo.category = ''
  searchInfo.status = ''
  getTableData()
}

// 选择变化
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 查看内容详情
const viewContent = async(row) => {
  const res = await getLearningContentDetail(row.ID)
  if (res.code === 0) {
    currentContent.value = res.data
    detailDialogVisible.value = true
  }
}

// 新增内容
const addContent = () => {
  isEdit.value = false
  resetEditForm()
  editDialogVisible.value = true
}

// 编辑内容
const editContent = (row) => {
  isEdit.value = true
  Object.assign(editForm, {
    ID: row.ID,
    title: row.title,
    type: row.type,
    category: row.category,
    difficulty: row.difficulty,
    duration: row.duration,
    status: row.status,
    coverImage: row.coverImage,
    description: row.description,
    content: row.content,
    mediaUrl: row.mediaUrl,
    tags: row.tags ? row.tags.join(',') : ''
  })
  editDialogVisible.value = true
}

// 保存内容
const saveContent = async() => {
  const formData = {
    ...editForm,
    tags: editForm.tags ? editForm.tags.split(',').map(tag => tag.trim()) : []
  }
  
  const res = isEdit.value 
    ? await updateLearningContent(formData)
    : await createLearningContent(formData)
    
  if (res.code === 0) {
    ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
    editDialogVisible.value = false
    getTableData()
    getStatisticsData()
  } else {
    ElMessage.error(res.msg || '操作失败')
  }
}

// 发布内容
const publishContent = async(row) => {
  const res = await updateLearningContent({
    ID: row.ID,
    status: 'published'
  })
  if (res.code === 0) {
    ElMessage.success('发布成功')
    getTableData()
    getStatisticsData()
  } else {
    ElMessage.error(res.msg || '发布失败')
  }
}

// 下线内容
const offlineContent = async(row) => {
  const res = await updateLearningContent({
    ID: row.ID,
    status: 'offline'
  })
  if (res.code === 0) {
    ElMessage.success('下线成功')
    getTableData()
    getStatisticsData()
  } else {
    ElMessage.error(res.msg || '下线失败')
  }
}

// 删除内容
const deleteContent = (row) => {
  ElMessageBox.confirm(
    `确定要删除内容 "${row.title}" 吗？此操作不可恢复！`,
    '确认删除',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async() => {
    const res = await deleteLearningContent(row.ID)
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
const batchPublish = () => {
  ElMessage.info('批量发布功能开发中...')
}

const batchOffline = () => {
  ElMessage.info('批量下线功能开发中...')
}

const batchDelete = () => {
  ElMessage.info('批量删除功能开发中...')
}

// 重置编辑表单
const resetEditForm = () => {
  Object.assign(editForm, {
    ID: '',
    title: '',
    type: 'article',
    category: 'knowledge',
    difficulty: 'beginner',
    duration: 10,
    status: 'draft',
    coverImage: '',
    description: '',
    content: '',
    mediaUrl: '',
    tags: ''
  })
}

// 工具函数
const formatTime = (time) => {
  if (!time) return '-'
  return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss')
}

const formatDuration = (duration) => {
  if (!duration) return '-'
  return `${duration}分钟`
}

const getTypeText = (type) => {
  const typeMap = {
    article: '文章',
    video: '视频',
    audio: '音频'
  }
  return typeMap[type] || '未知'
}

const getTypeTag = (type) => {
  const tagMap = {
    article: 'primary',
    video: 'success',
    audio: 'warning'
  }
  return tagMap[type] || 'info'
}

const getCategoryText = (category) => {
  const categoryMap = {
    knowledge: '科普知识',
    recovery: '康复指导',
    mental: '心理健康',
    experience: '经验分享'
  }
  return categoryMap[category] || '未知'
}

const getCategoryTag = (category) => {
  const tagMap = {
    knowledge: 'primary',
    recovery: 'success',
    mental: 'warning',
    experience: 'info'
  }
  return tagMap[category] || 'info'
}

const getStatusText = (status) => {
  const statusMap = {
    published: '已发布',
    draft: '草稿',
    offline: '已下线'
  }
  return statusMap[status] || '未知'
}

const getStatusTag = (status) => {
  const tagMap = {
    published: 'success',
    draft: 'warning',
    offline: 'info'
  }
  return tagMap[status] || 'info'
}

const getDifficultyText = (difficulty) => {
  const difficultyMap = {
    beginner: '入门',
    intermediate: '进阶',
    advanced: '高级'
  }
  return difficultyMap[difficulty] || '未知'
}

// 生命周期
onMounted(() => {
  getTableData()
  getStatisticsData()
})
</script>

<style lang="scss" scoped>
.learning-management {
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
          
          &.published-icon {
            background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
          }
          
          &.draft-icon {
            background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
          }
          
          &.views-icon {
            background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
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
  
  .no-image {
    color: #c0c4cc;
    font-size: 12px;
  }
  
  .content-detail {
    .content-cover {
      h4 {
        margin: 0 0 16px 0;
        color: #303133;
      }
      
      .no-cover {
        height: 200px;
        display: flex;
        align-items: center;
        justify-content: center;
        background: #f5f7fa;
        color: #909399;
        border-radius: 8px;
      }
    }
    
    .content-description, .content-body {
      margin-top: 24px;
      
      h4 {
        margin: 0 0 16px 0;
        color: #303133;
      }
      
      p {
        line-height: 1.6;
        color: #606266;
      }
      
      .content-preview {
        max-height: 300px;
        overflow-y: auto;
        padding: 16px;
        background: #f5f7fa;
        border-radius: 8px;
        line-height: 1.6;
        color: #606266;
      }
    }
  }
}
</style> 