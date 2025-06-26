<template>
  <div class="user-management">
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="getTableData">
        <el-form-item label="用户昵称" prop="nickname">
          <el-input v-model="searchInfo.nickname" placeholder="搜索用户昵称" clearable />
        </el-form-item>
        <el-form-item label="用户状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="请选择用户状态" clearable>
            <el-option label="正常" value="1" />
            <el-option label="禁用" value="2" />
            <el-option label="封禁" value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="注册时间" prop="dateRange">
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
              <div class="stat-icon user-icon">
                <el-icon><User /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.totalUsers }}</div>
                <div class="stat-label">总用户数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon active-icon">
                <el-icon><TrendCharts /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.activeUsers }}</div>
                <div class="stat-label">活跃用户</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon new-icon">
                <el-icon><UserFilled /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.newUsers }}</div>
                <div class="stat-label">新增用户</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon retention-icon">
                <el-icon><DataAnalysis /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.retentionRate }}%</div>
                <div class="stat-label">留存率</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="exportUsers">导出用户</el-button>
        <el-button type="warning" icon="refresh" @click="getTableData">刷新</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="头像" width="80">
          <template #default="scope">
            <el-avatar :size="40" :src="scope.row.avatarUrl" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="昵称" prop="nickname" width="150" />
        <el-table-column align="left" label="微信OpenID" prop="openid" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="等级" width="80">
          <template #default="scope">
            <el-tag type="success">{{ scope.row.level }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="经验值" prop="experience" width="100" />
        <el-table-column align="left" label="连续打卡" width="100">
          <template #default="scope">
            <span class="streak-text">{{ scope.row.currentStreak }}天</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="隐私级别" width="100">
          <template #default="scope">
            <el-tag :type="getPrivacyTagType(scope.row.privacyLevel)">
              {{ getPrivacyText(scope.row.privacyLevel) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" width="80">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="最后登录" width="180">
          <template #default="scope">
            <span>{{ formatTime(scope.row.lastLoginAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="注册时间" width="180">
          <template #default="scope">
            <span>{{ formatTime(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" width="200">
          <template #default="scope">
            <el-button type="primary" link icon="view" size="small" @click="viewUser(scope.row)">查看</el-button>
            <el-button type="warning" link icon="edit" size="small" @click="editUser(scope.row)">编辑</el-button>
            <el-button 
              :type="scope.row.status === 1 ? 'danger' : 'success'" 
              link 
              :icon="scope.row.status === 1 ? 'close' : 'check'" 
              size="small" 
              @click="toggleUserStatus(scope.row)"
            >
              {{ scope.row.status === 1 ? '禁用' : '启用' }}
            </el-button>
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

    <!-- 用户详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="用户详情" width="800px">
      <div v-if="currentUser" class="user-detail">
        <el-row :gutter="20">
          <el-col :span="8">
            <div class="user-avatar">
              <el-avatar :size="120" :src="currentUser.avatarUrl" />
              <div class="user-basic">
                <h3>{{ currentUser.nickname }}</h3>
                <p>等级 {{ currentUser.level }} - {{ currentUser.levelTitle }}</p>
              </div>
            </div>
          </el-col>
          <el-col :span="16">
            <el-descriptions :column="2" border>
              <el-descriptions-item label="用户ID">{{ currentUser.ID }}</el-descriptions-item>
              <el-descriptions-item label="微信OpenID">{{ currentUser.openid }}</el-descriptions-item>
              <el-descriptions-item label="性别">{{ getGenderText(currentUser.gender) }}</el-descriptions-item>
              <el-descriptions-item label="地区">{{ currentUser.city }}, {{ currentUser.province }}</el-descriptions-item>
              <el-descriptions-item label="经验值">{{ currentUser.experience }}</el-descriptions-item>
              <el-descriptions-item label="当前连击">{{ currentUser.currentStreak }}天</el-descriptions-item>
              <el-descriptions-item label="最长连击">{{ currentUser.longestStreak }}天</el-descriptions-item>
              <el-descriptions-item label="总打卡数">{{ currentUser.totalCheckins }}</el-descriptions-item>
              <el-descriptions-item label="隐私级别">{{ getPrivacyText(currentUser.privacyLevel) }}</el-descriptions-item>
              <el-descriptions-item label="账户状态">{{ getStatusText(currentUser.status) }}</el-descriptions-item>
              <el-descriptions-item label="注册时间">{{ formatTime(currentUser.CreatedAt) }}</el-descriptions-item>
              <el-descriptions-item label="最后登录">{{ formatTime(currentUser.lastLoginAt) }}</el-descriptions-item>
            </el-descriptions>
          </el-col>
        </el-row>

        <!-- 用户统计数据 -->
        <div class="user-stats">
          <h4>用户统计</h4>
          <el-row :gutter="20">
            <el-col :span="6">
              <el-statistic title="获得成就" :value="currentUser.achievementCount" suffix="个" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="社区发帖" :value="currentUser.postCount" suffix="篇" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="学习时长" :value="currentUser.learningMinutes" suffix="分钟" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="帮助他人" :value="currentUser.helpCount" suffix="次" />
            </el-col>
          </el-row>
        </div>
      </div>
    </el-dialog>

    <!-- 编辑用户对话框 -->
    <el-dialog v-model="editDialogVisible" title="编辑用户" width="600px">
      <el-form ref="editFormRef" :model="editForm" :rules="editRules" label-width="100px">
        <el-form-item label="用户昵称" prop="nickname">
          <el-input v-model="editForm.nickname" placeholder="请输入用户昵称" />
        </el-form-item>
        <el-form-item label="隐私级别" prop="privacyLevel">
          <el-select v-model="editForm.privacyLevel" placeholder="请选择隐私级别">
            <el-option label="公开" :value="1" />
            <el-option label="好友可见" :value="2" />
            <el-option label="私密" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="账户状态" prop="status">
          <el-select v-model="editForm.status" placeholder="请选择账户状态">
            <el-option label="正常" :value="1" />
            <el-option label="禁用" :value="2" />
            <el-option label="封禁" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="editForm.remark" type="textarea" :rows="3" placeholder="管理员备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveUser">保存</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { User, TrendCharts, UserFilled, DataAnalysis } from '@element-plus/icons-vue'
import { getUserList, getUserDetail, updateUserStatus, getUserStatistics } from '@/api/miniprogram'
import { formatTimeToStr } from '@/utils/date'

defineOptions({
  name: 'UserManagement'
})

// 响应式数据
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = reactive({
  nickname: '',
  status: '',
  dateRange: []
})

const statistics = reactive({
  totalUsers: 0,
  activeUsers: 0,
  newUsers: 0,
  retentionRate: 0
})

const detailDialogVisible = ref(false)
const editDialogVisible = ref(false)
const currentUser = ref(null)
const multipleSelection = ref([])

// 编辑表单
const editForm = reactive({
  ID: '',
  nickname: '',
  privacyLevel: 1,
  status: 1,
  remark: ''
})

const editRules = {
  nickname: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' }
  ],
  privacyLevel: [
    { required: true, message: '请选择隐私级别', trigger: 'change' }
  ],
  status: [
    { required: true, message: '请选择账户状态', trigger: 'change' }
  ]
}

// 获取表格数据
const getTableData = async() => {
  const table = await getUserList({
    page: page.value,
    pageSize: pageSize.value,
    nickname: searchInfo.nickname,
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
const getStatistics = async() => {
  const stats = await getUserStatistics()
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
  searchInfo.nickname = ''
  searchInfo.status = ''
  searchInfo.dateRange = []
  getTableData()
}

// 选择变化
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 查看用户详情
const viewUser = async(row) => {
  const res = await getUserDetail(row.ID)
  if (res.code === 0) {
    currentUser.value = res.data
    detailDialogVisible.value = true
  }
}

// 编辑用户
const editUser = (row) => {
  Object.assign(editForm, {
    ID: row.ID,
    nickname: row.nickname,
    privacyLevel: row.privacyLevel,
    status: row.status,
    remark: row.remark || ''
  })
  editDialogVisible.value = true
}

// 保存用户
const saveUser = async() => {
  const res = await updateUserStatus(editForm)
  if (res.code === 0) {
    ElMessage.success('用户信息更新成功')
    editDialogVisible.value = false
    getTableData()
  } else {
    ElMessage.error(res.msg || '更新失败')
  }
}

// 切换用户状态
const toggleUserStatus = (row) => {
  const action = row.status === 1 ? '禁用' : '启用'
  ElMessageBox.confirm(
    `确定要${action}用户 ${row.nickname} 吗？`,
    '确认操作',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async() => {
    const newStatus = row.status === 1 ? 2 : 1
    const res = await updateUserStatus({
      ID: row.ID,
      status: newStatus
    })
    if (res.code === 0) {
      ElMessage.success(`${action}成功`)
      getTableData()
    } else {
      ElMessage.error(res.msg || `${action}失败`)
    }
  })
}

// 导出用户
const exportUsers = () => {
  ElMessage.info('导出功能开发中...')
}

// 工具函数
const formatTime = (time) => {
  if (!time) return '-'
  return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss')
}

const getStatusText = (status) => {
  const statusMap = {
    1: '正常',
    2: '禁用',
    3: '封禁'
  }
  return statusMap[status] || '未知'
}

const getStatusTagType = (status) => {
  const typeMap = {
    1: 'success',
    2: 'warning',
    3: 'danger'
  }
  return typeMap[status] || 'info'
}

const getPrivacyText = (level) => {
  const levelMap = {
    1: '公开',
    2: '好友可见',
    3: '私密'
  }
  return levelMap[level] || '未知'
}

const getPrivacyTagType = (level) => {
  const typeMap = {
    1: 'success',
    2: 'warning',
    3: 'info'
  }
  return typeMap[level] || 'info'
}

const getGenderText = (gender) => {
  const genderMap = {
    1: '男',
    2: '女',
    0: '未知'
  }
  return genderMap[gender] || '未知'
}

// 生命周期
onMounted(() => {
  getTableData()
  getStatistics()
})
</script>

<style lang="scss" scoped>
.user-management {
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
          
          &.user-icon {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
          }
          
          &.active-icon {
            background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
          }
          
          &.new-icon {
            background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
          }
          
          &.retention-icon {
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
  
  .streak-text {
    color: #67c23a;
    font-weight: 600;
  }
  
  .user-detail {
    .user-avatar {
      text-align: center;
      
      .user-basic {
        margin-top: 16px;
        
        h3 {
          margin: 0 0 8px 0;
          color: #303133;
        }
        
        p {
          margin: 0;
          color: #909399;
          font-size: 14px;
        }
      }
    }
    
    .user-stats {
      margin-top: 24px;
      
      h4 {
        margin: 0 0 16px 0;
        color: #303133;
      }
    }
  }
}
</style> 