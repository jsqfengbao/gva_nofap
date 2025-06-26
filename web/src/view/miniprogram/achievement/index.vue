<template>
  <div class="achievement-management">
    <!-- 搜索区域 -->
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="成就名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="搜索成就名称" clearable />
        </el-form-item>
        <el-form-item label="成就类型" prop="type">
          <el-select v-model="searchInfo.type" placeholder="请选择成就类型" clearable>
            <el-option label="打卡类" value="checkin" />
            <el-option label="等级类" value="level" />
            <el-option label="社区类" value="community" />
            <el-option label="学习类" value="learning" />
            <el-option label="特殊类" value="special" />
          </el-select>
        </el-form-item>
        <el-form-item label="稀有度" prop="rarity">
          <el-select v-model="searchInfo.rarity" placeholder="请选择稀有度" clearable>
            <el-option label="普通" value="common" />
            <el-option label="稀有" value="rare" />
            <el-option label="史诗" value="epic" />
            <el-option label="传说" value="legendary" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="isActive">
          <el-select v-model="searchInfo.isActive" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="true" />
            <el-option label="禁用" :value="false" />
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
                <el-icon><Trophy /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.totalAchievements }}</div>
                <div class="stat-label">总成就数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon active-icon">
                <el-icon><Check /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.activeAchievements }}</div>
                <div class="stat-label">启用中</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon unlocked-icon">
                <el-icon><Star /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.totalUnlocked }}</div>
                <div class="stat-label">总解锁次数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon legendary-icon">
                <el-icon><Crown /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.legendaryCount }}</div>
                <div class="stat-label">传说成就</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <!-- 数据表格 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addAchievement">新增成就</el-button>
        <el-button type="success" icon="check" @click="batchEnable" :disabled="!multipleSelection.length">批量启用</el-button>
        <el-button type="warning" icon="close" @click="batchDisable" :disabled="!multipleSelection.length">批量禁用</el-button>
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
        <el-table-column align="left" label="图标" width="80">
          <template #default="scope">
            <div class="achievement-icon" :class="`rarity-${scope.row.rarity}`">
              {{ scope.row.icon }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="成就名称" prop="name" width="200" />
        <el-table-column align="left" label="描述" prop="description" width="250" show-overflow-tooltip />
        <el-table-column align="left" label="类型" width="100">
          <template #default="scope">
            <el-tag :type="getTypeTag(scope.row.type)">
              {{ getTypeText(scope.row.type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="稀有度" width="100">
          <template #default="scope">
            <el-tag :type="getRarityTag(scope.row.rarity)">
              {{ getRarityText(scope.row.rarity) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="条件" width="150">
          <template #default="scope">
            <span>{{ getConditionText(scope.row) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="奖励经验" prop="rewardExp" width="100" />
        <el-table-column align="left" label="解锁次数" prop="unlockCount" width="100" />
        <el-table-column align="left" label="状态" width="80">
          <template #default="scope">
            <el-switch 
              v-model="scope.row.isActive" 
              @change="toggleAchievementStatus(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" width="180">
          <template #default="scope">
            <span>{{ formatTime(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" width="150">
          <template #default="scope">
            <el-button type="primary" link size="small" @click="viewAchievement(scope.row)">查看</el-button>
            <el-button type="warning" link size="small" @click="editAchievement(scope.row)">编辑</el-button>
            <el-button type="danger" link size="small" @click="deleteAchievement(scope.row)">删除</el-button>
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

    <!-- 成就详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="成就详情" width="700px">
      <div v-if="currentAchievement" class="achievement-detail">
        <div class="achievement-header">
          <div class="achievement-icon-large" :class="`rarity-${currentAchievement.rarity}`">
            {{ currentAchievement.icon }}
          </div>
          <div class="achievement-info">
            <h3>{{ currentAchievement.name }}</h3>
            <p>{{ currentAchievement.description }}</p>
            <div class="achievement-meta">
              <el-tag :type="getTypeTag(currentAchievement.type)">{{ getTypeText(currentAchievement.type) }}</el-tag>
              <el-tag :type="getRarityTag(currentAchievement.rarity)">{{ getRarityText(currentAchievement.rarity) }}</el-tag>
            </div>
          </div>
        </div>

        <el-descriptions :column="2" border>
          <el-descriptions-item label="解锁条件">{{ getConditionText(currentAchievement) }}</el-descriptions-item>
          <el-descriptions-item label="奖励经验">{{ currentAchievement.rewardExp }}</el-descriptions-item>
          <el-descriptions-item label="解锁次数">{{ currentAchievement.unlockCount }}</el-descriptions-item>
          <el-descriptions-item label="状态">{{ currentAchievement.isActive ? '启用' : '禁用' }}</el-descriptions-item>
          <el-descriptions-item label="创建时间" :span="2">{{ formatTime(currentAchievement.CreatedAt) }}</el-descriptions-item>
        </el-descriptions>

        <div class="achievement-stats">
          <h4>解锁统计</h4>
          <div class="stats-chart">
            <!-- 这里可以添加图表组件 -->
            <p>解锁趋势图表</p>
          </div>
        </div>
      </div>
    </el-dialog>

    <!-- 新增/编辑成就对话框 -->
    <el-dialog v-model="editDialogVisible" :title="isEdit ? '编辑成就' : '新增成就'" width="800px">
      <el-form ref="editFormRef" :model="editForm" :rules="editRules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="成就名称" prop="name">
              <el-input v-model="editForm.name" placeholder="请输入成就名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="成就图标" prop="icon">
              <el-input v-model="editForm.icon" placeholder="请输入图标emoji" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="成就描述" prop="description">
          <el-input v-model="editForm.description" type="textarea" :rows="2" placeholder="请输入成就描述" />
        </el-form-item>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="成就类型" prop="type">
              <el-select v-model="editForm.type" placeholder="请选择成就类型">
                <el-option label="打卡类" value="checkin" />
                <el-option label="等级类" value="level" />
                <el-option label="社区类" value="community" />
                <el-option label="学习类" value="learning" />
                <el-option label="特殊类" value="special" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="稀有度" prop="rarity">
              <el-select v-model="editForm.rarity" placeholder="请选择稀有度">
                <el-option label="普通" value="common" />
                <el-option label="稀有" value="rare" />
                <el-option label="史诗" value="epic" />
                <el-option label="传说" value="legendary" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="条件类型" prop="conditionType">
              <el-select v-model="editForm.conditionType" placeholder="请选择条件类型">
                <el-option label="连续打卡" value="consecutive_checkins" />
                <el-option label="总打卡数" value="total_checkins" />
                <el-option label="达到等级" value="reach_level" />
                <el-option label="社区发帖" value="community_posts" />
                <el-option label="学习时长" value="learning_time" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="条件值" prop="conditionValue">
              <el-input-number v-model="editForm.conditionValue" :min="1" placeholder="条件数值" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="奖励经验" prop="rewardExp">
              <el-input-number v-model="editForm.rewardExp" :min="0" placeholder="奖励经验值" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态" prop="isActive">
              <el-switch v-model="editForm.isActive" active-text="启用" inactive-text="禁用" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveAchievement">保存</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Trophy, Check, Star, Crown } from '@element-plus/icons-vue'
import { getAchievementList, getAchievementDetail, createAchievement, updateAchievement, deleteAchievement, getAchievementStatistics } from '@/api/miniprogram'
import { formatTimeToStr } from '@/utils/date'

defineOptions({
  name: 'AchievementManagement'
})

// 响应式数据
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = reactive({
  name: '',
  type: '',
  rarity: '',
  isActive: ''
})

const statistics = reactive({
  totalAchievements: 0,
  activeAchievements: 0,
  totalUnlocked: 0,
  legendaryCount: 0
})

const detailDialogVisible = ref(false)
const editDialogVisible = ref(false)
const currentAchievement = ref(null)
const multipleSelection = ref([])
const isEdit = ref(false)

// 编辑表单
const editForm = reactive({
  ID: '',
  name: '',
  icon: '',
  description: '',
  type: 'checkin',
  rarity: 'common',
  conditionType: 'consecutive_checkins',
  conditionValue: 1,
  rewardExp: 10,
  isActive: true
})

const editRules = {
  name: [
    { required: true, message: '请输入成就名称', trigger: 'blur' }
  ],
  icon: [
    { required: true, message: '请输入成就图标', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入成就描述', trigger: 'blur' }
  ],
  type: [
    { required: true, message: '请选择成就类型', trigger: 'change' }
  ],
  rarity: [
    { required: true, message: '请选择稀有度', trigger: 'change' }
  ],
  conditionType: [
    { required: true, message: '请选择条件类型', trigger: 'change' }
  ],
  conditionValue: [
    { required: true, message: '请输入条件值', trigger: 'blur' }
  ]
}

// 获取表格数据
const getTableData = async() => {
  const table = await getAchievementList({
    page: page.value,
    pageSize: pageSize.value,
    name: searchInfo.name,
    type: searchInfo.type,
    rarity: searchInfo.rarity,
    isActive: searchInfo.isActive
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
  const stats = await getAchievementStatistics()
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
  searchInfo.name = ''
  searchInfo.type = ''
  searchInfo.rarity = ''
  searchInfo.isActive = ''
  getTableData()
}

// 选择变化
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 查看成就详情
const viewAchievement = async(row) => {
  const res = await getAchievementDetail(row.ID)
  if (res.code === 0) {
    currentAchievement.value = res.data
    detailDialogVisible.value = true
  }
}

// 新增成就
const addAchievement = () => {
  isEdit.value = false
  resetEditForm()
  editDialogVisible.value = true
}

// 编辑成就
const editAchievement = (row) => {
  isEdit.value = true
  Object.assign(editForm, row)
  editDialogVisible.value = true
}

// 保存成就
const saveAchievement = async() => {
  const res = isEdit.value 
    ? await updateAchievement(editForm)
    : await createAchievement(editForm)
    
  if (res.code === 0) {
    ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
    editDialogVisible.value = false
    getTableData()
    getStatisticsData()
  } else {
    ElMessage.error(res.msg || '操作失败')
  }
}

// 切换成就状态
const toggleAchievementStatus = async(row) => {
  const res = await updateAchievement({
    ID: row.ID,
    isActive: row.isActive
  })
  if (res.code === 0) {
    ElMessage.success('状态更新成功')
    getStatisticsData()
  } else {
    ElMessage.error(res.msg || '状态更新失败')
    row.isActive = !row.isActive // 回滚状态
  }
}

// 删除成就
const deleteAchievement = (row) => {
  ElMessageBox.confirm(
    `确定要删除成就 "${row.name}" 吗？此操作不可恢复！`,
    '确认删除',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async() => {
    const res = await deleteAchievement(row.ID)
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
const batchEnable = () => {
  ElMessage.info('批量启用功能开发中...')
}

const batchDisable = () => {
  ElMessage.info('批量禁用功能开发中...')
}

const batchDelete = () => {
  ElMessage.info('批量删除功能开发中...')
}

// 重置编辑表单
const resetEditForm = () => {
  Object.assign(editForm, {
    ID: '',
    name: '',
    icon: '',
    description: '',
    type: 'checkin',
    rarity: 'common',
    conditionType: 'consecutive_checkins',
    conditionValue: 1,
    rewardExp: 10,
    isActive: true
  })
}

// 工具函数
const formatTime = (time) => {
  if (!time) return '-'
  return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss')
}

const getTypeText = (type) => {
  const typeMap = {
    checkin: '打卡类',
    level: '等级类',
    community: '社区类',
    learning: '学习类',
    special: '特殊类'
  }
  return typeMap[type] || '未知'
}

const getTypeTag = (type) => {
  const tagMap = {
    checkin: 'primary',
    level: 'success',
    community: 'warning',
    learning: 'info',
    special: 'danger'
  }
  return tagMap[type] || 'info'
}

const getRarityText = (rarity) => {
  const rarityMap = {
    common: '普通',
    rare: '稀有',
    epic: '史诗',
    legendary: '传说'
  }
  return rarityMap[rarity] || '未知'
}

const getRarityTag = (rarity) => {
  const tagMap = {
    common: 'info',
    rare: 'primary',
    epic: 'warning',
    legendary: 'danger'
  }
  return tagMap[rarity] || 'info'
}

const getConditionText = (achievement) => {
  const conditionMap = {
    consecutive_checkins: '连续打卡',
    total_checkins: '总打卡数',
    reach_level: '达到等级',
    community_posts: '社区发帖',
    learning_time: '学习时长'
  }
  const conditionText = conditionMap[achievement.conditionType] || '未知条件'
  return `${conditionText} ${achievement.conditionValue}`
}

// 生命周期
onMounted(() => {
  getTableData()
  getStatisticsData()
})
</script>

<style lang="scss" scoped>
.achievement-management {
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
          
          &.active-icon {
            background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
          }
          
          &.unlocked-icon {
            background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
          }
          
          &.legendary-icon {
            background: linear-gradient(135deg, #ffd700 0%, #ffb347 100%);
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
  
  .achievement-icon {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 20px;
    margin: 0 auto;
    
    &.rarity-common {
      background: #909399;
      color: white;
    }
    
    &.rarity-rare {
      background: #409eff;
      color: white;
    }
    
    &.rarity-epic {
      background: #a855f7;
      color: white;
    }
    
    &.rarity-legendary {
      background: linear-gradient(135deg, #ffd700 0%, #ffb347 100%);
      color: white;
    }
  }
  
  .achievement-detail {
    .achievement-header {
      display: flex;
      align-items: center;
      margin-bottom: 24px;
      
      .achievement-icon-large {
        width: 80px;
        height: 80px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 40px;
        margin-right: 20px;
        
        &.rarity-common {
          background: #909399;
          color: white;
        }
        
        &.rarity-rare {
          background: #409eff;
          color: white;
        }
        
        &.rarity-epic {
          background: #a855f7;
          color: white;
        }
        
        &.rarity-legendary {
          background: linear-gradient(135deg, #ffd700 0%, #ffb347 100%);
          color: white;
        }
      }
      
      .achievement-info {
        flex: 1;
        
        h3 {
          margin: 0 0 8px 0;
          color: #303133;
        }
        
        p {
          margin: 0 0 12px 0;
          color: #606266;
          line-height: 1.4;
        }
        
        .achievement-meta {
          .el-tag {
            margin-right: 8px;
          }
        }
      }
    }
    
    .achievement-stats {
      margin-top: 24px;
      
      h4 {
        margin: 0 0 16px 0;
        color: #303133;
      }
      
      .stats-chart {
        height: 200px;
        display: flex;
        align-items: center;
        justify-content: center;
        background: #f5f7fa;
        border-radius: 8px;
        color: #909399;
      }
    }
  }
}
</style> 