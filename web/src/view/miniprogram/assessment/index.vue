<template>
  <div class="assessment-management">
    <!-- 搜索区域 -->
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户昵称" prop="nickname">
          <el-input v-model="searchInfo.nickname" placeholder="搜索用户昵称" clearable />
        </el-form-item>
        <el-form-item label="评估类型" prop="assessmentType">
          <el-select v-model="searchInfo.assessmentType" placeholder="请选择评估类型" clearable>
            <el-option label="成瘾程度评估" value="addiction" />
            <el-option label="心理健康评估" value="mental" />
            <el-option label="恢复进度评估" value="recovery" />
          </el-select>
        </el-form-item>
        <el-form-item label="风险等级" prop="riskLevel">
          <el-select v-model="searchInfo.riskLevel" placeholder="请选择风险等级" clearable>
            <el-option label="低风险" value="low" />
            <el-option label="中风险" value="medium" />
            <el-option label="高风险" value="high" />
          </el-select>
        </el-form-item>
        <el-form-item label="评估时间" prop="dateRange">
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
                <el-icon><Document /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.totalAssessments }}</div>
                <div class="stat-label">总评估数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon low-risk-icon">
                <el-icon><SuccessFilled /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.lowRisk }}</div>
                <div class="stat-label">低风险</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon medium-risk-icon">
                <el-icon><WarningFilled /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.mediumRisk }}</div>
                <div class="stat-label">中风险</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card">
            <div class="stat-content">
              <div class="stat-icon high-risk-icon">
                <el-icon><CircleCloseFilled /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ statistics.highRisk }}</div>
                <div class="stat-label">高风险</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <!-- 数据表格 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="download" @click="exportAssessments">导出评估数据</el-button>
        <el-button type="success" icon="pie-chart" @click="showStatistics">统计分析</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        :data="tableData"
        row-key="ID"
      >
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="用户昵称" prop="userNickname" width="150" />
        <el-table-column align="left" label="评估类型" width="120">
          <template #default="scope">
            <el-tag :type="getAssessmentTypeTag(scope.row.assessmentType)">
              {{ getAssessmentTypeText(scope.row.assessmentType) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="总分" prop="totalScore" width="80" />
        <el-table-column align="left" label="风险等级" width="100">
          <template #default="scope">
            <el-tag :type="getRiskLevelTag(scope.row.riskLevel)">
              {{ getRiskLevelText(scope.row.riskLevel) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="建议" prop="recommendations" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="评估时间" width="180">
          <template #default="scope">
            <span>{{ formatTime(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" width="150">
          <template #default="scope">
            <el-button type="primary" link size="small" @click="viewAssessment(scope.row)">查看详情</el-button>
            <el-button type="danger" link size="small" @click="deleteAssessment(scope.row)">删除</el-button>
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

    <!-- 评估详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="评估详情" width="900px">
      <div v-if="currentAssessment" class="assessment-detail">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-descriptions title="基本信息" :column="1" border>
              <el-descriptions-item label="用户昵称">{{ currentAssessment.userNickname }}</el-descriptions-item>
              <el-descriptions-item label="评估类型">{{ getAssessmentTypeText(currentAssessment.assessmentType) }}</el-descriptions-item>
              <el-descriptions-item label="总分">{{ currentAssessment.totalScore }}</el-descriptions-item>
              <el-descriptions-item label="风险等级">
                <el-tag :type="getRiskLevelTag(currentAssessment.riskLevel)">
                  {{ getRiskLevelText(currentAssessment.riskLevel) }}
                </el-tag>
              </el-descriptions-item>
              <el-descriptions-item label="评估时间">{{ formatTime(currentAssessment.CreatedAt) }}</el-descriptions-item>
            </el-descriptions>
          </el-col>
          <el-col :span="12">
            <div class="score-chart">
              <h4>分数分布</h4>
              <!-- 这里可以添加图表组件 -->
              <div class="score-breakdown">
                <div v-for="(score, index) in currentAssessment.scores" :key="index" class="score-item">
                  <span class="score-label">{{ score.dimension }}:</span>
                  <span class="score-value">{{ score.score }}</span>
                  <el-progress 
                    :percentage="(score.score / score.maxScore) * 100" 
                    :color="getScoreColor(score.score / score.maxScore)"
                  />
                </div>
              </div>
            </div>
          </el-col>
        </el-row>

        <div class="recommendations-section">
          <h4>建议与措施</h4>
          <el-card>
            <p>{{ currentAssessment.recommendations }}</p>
          </el-card>
        </div>

        <div class="answers-section">
          <h4>详细答案</h4>
          <el-table :data="currentAssessment.answers" style="width: 100%">
            <el-table-column label="题目" prop="question" width="400" />
            <el-table-column label="答案" prop="answer" width="200" />
            <el-table-column label="得分" prop="score" width="100" />
          </el-table>
        </div>
      </div>
    </el-dialog>

    <!-- 统计分析对话框 -->
    <el-dialog v-model="statisticsDialogVisible" title="统计分析" width="1000px">
      <div class="statistics-content">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-card title="风险等级分布">
              <!-- 饼图组件位置 -->
              <div class="chart-placeholder">
                <p>风险等级分布图表</p>
              </div>
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card title="评估趋势">
              <!-- 折线图组件位置 -->
              <div class="chart-placeholder">
                <p>评估趋势图表</p>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Document, SuccessFilled, WarningFilled, CircleCloseFilled } from '@element-plus/icons-vue'
import { getAssessmentList, getAssessmentDetail, deleteAssessmentRecord, getAssessmentStatistics } from '@/api/miniprogram'
import { formatTimeToStr } from '@/utils/date'

defineOptions({
  name: 'AssessmentManagement'
})

// 响应式数据
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = reactive({
  nickname: '',
  assessmentType: '',
  riskLevel: '',
  dateRange: []
})

const statistics = reactive({
  totalAssessments: 0,
  lowRisk: 0,
  mediumRisk: 0,
  highRisk: 0
})

const detailDialogVisible = ref(false)
const statisticsDialogVisible = ref(false)
const currentAssessment = ref(null)

// 获取表格数据
const getTableData = async() => {
  const table = await getAssessmentList({
    page: page.value,
    pageSize: pageSize.value,
    nickname: searchInfo.nickname,
    assessmentType: searchInfo.assessmentType,
    riskLevel: searchInfo.riskLevel,
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
  const stats = await getAssessmentStatistics()
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
  searchInfo.assessmentType = ''
  searchInfo.riskLevel = ''
  searchInfo.dateRange = []
  getTableData()
}

// 查看评估详情
const viewAssessment = async(row) => {
  const res = await getAssessmentDetail(row.ID)
  if (res.code === 0) {
    currentAssessment.value = res.data
    detailDialogVisible.value = true
  }
}

// 删除评估记录
const deleteAssessment = (row) => {
  ElMessageBox.confirm(
    `确定要删除用户 ${row.userNickname} 的评估记录吗？`,
    '确认删除',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async() => {
    const res = await deleteAssessmentRecord(row.ID)
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
      getStatisticsData()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  })
}

// 导出评估数据
const exportAssessments = () => {
  ElMessage.info('导出功能开发中...')
}

// 显示统计分析
const showStatistics = () => {
  statisticsDialogVisible.value = true
}

// 工具函数
const formatTime = (time) => {
  if (!time) return '-'
  return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss')
}

const getAssessmentTypeText = (type) => {
  const typeMap = {
    addiction: '成瘾程度评估',
    mental: '心理健康评估',
    recovery: '恢复进度评估'
  }
  return typeMap[type] || '未知'
}

const getAssessmentTypeTag = (type) => {
  const tagMap = {
    addiction: 'danger',
    mental: 'warning',
    recovery: 'success'
  }
  return tagMap[type] || 'info'
}

const getRiskLevelText = (level) => {
  const levelMap = {
    low: '低风险',
    medium: '中风险',
    high: '高风险'
  }
  return levelMap[level] || '未知'
}

const getRiskLevelTag = (level) => {
  const tagMap = {
    low: 'success',
    medium: 'warning',
    high: 'danger'
  }
  return tagMap[level] || 'info'
}

const getScoreColor = (percentage) => {
  if (percentage >= 0.8) return '#f56c6c'
  if (percentage >= 0.6) return '#e6a23c'
  return '#67c23a'
}

// 生命周期
onMounted(() => {
  getTableData()
  getStatisticsData()
})
</script>

<style lang="scss" scoped>
.assessment-management {
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
          
          &.low-risk-icon {
            background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
          }
          
          &.medium-risk-icon {
            background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
          }
          
          &.high-risk-icon {
            background: linear-gradient(135deg, #ff6b6b 0%, #ee5a24 100%);
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
  
  .assessment-detail {
    .score-chart {
      h4 {
        margin: 0 0 16px 0;
        color: #303133;
      }
      
      .score-breakdown {
        .score-item {
          margin-bottom: 12px;
          
          .score-label {
            display: inline-block;
            width: 100px;
            font-size: 14px;
            color: #606266;
          }
          
          .score-value {
            display: inline-block;
            width: 40px;
            font-weight: bold;
            color: #303133;
          }
        }
      }
    }
    
    .recommendations-section, .answers-section {
      margin-top: 24px;
      
      h4 {
        margin: 0 0 16px 0;
        color: #303133;
      }
    }
  }
  
  .statistics-content {
    .chart-placeholder {
      height: 300px;
      display: flex;
      align-items: center;
      justify-content: center;
      background: #f5f7fa;
      border-radius: 8px;
      color: #909399;
    }
  }
}
</style> 