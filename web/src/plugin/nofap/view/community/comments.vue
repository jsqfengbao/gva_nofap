<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="帖子ID">
          <el-input v-model="searchInfo.postId" placeholder="帖子ID" clearable />
        </el-form-item>

        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="评论状态" clearable>
            <el-option label="全部" :value="null" />
            <el-option label="正常" :value="1" />
            <el-option label="待审核" :value="0" />
            <el-option label="已屏蔽" :value="-1" />
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="ID" prop="ID" width="70" />
        <el-table-column align="left" label="帖子ID" prop="postId" width="80" />
        <el-table-column align="left" label="作者" prop="userNickname" width="120" />
        <el-table-column align="left" label="评论内容" prop="content" min-width="200">
          <template #default="scope">
            <span>{{ scope.row.content?.substring(0, 60) }}...</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="点赞" prop="likeCount" width="70" />
        <el-table-column align="left" label="状态" prop="status" width="90">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" prop="CreatedAt" width="160">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="200"
        >
          <template #default="scope">
            <template v-if="scope.row.status !== 1">
              <el-button
                type="success"
                link
                icon="check"
                @click="changeStatus(scope.row, 1)"
              >
                通过
              </el-button>
            </template>
            <template v-if="scope.row.status !== -1">
              <el-button
                type="warning"
                link
                icon="hide"
                @click="changeStatus(scope.row, -1)"
              >
                屏蔽
              </el-button>
            </template>
            <template v-if="scope.row.status === -1">
              <el-button
                type="success"
                link
                icon="unlock"
                @click="changeStatus(scope.row, 1)"
              >
                恢复
              </el-button>
            </template>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
            >
              删除
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
  </div>
</template>

<script setup>
import {
  getCommunityCommentList,
  updateCommunityCommentStatus,
  deleteCommunityComment,
} from '@/plugin/nofap/api/nofap.js'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'NofapComments'
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  postId: '',
  status: null
})

const elSearchFormRef = ref()

// 状态工具函数
const getStatusType = (status) => {
  switch (status) {
    case 1: return 'success'
    case 0: return 'warning'
    case -1: return 'danger'
    default: return 'info'
  }
}

const getStatusText = (status) => {
  switch (status) {
    case 1: return '正常'
    case 0: return '待审核'
    case -1: return '已屏蔽'
    default: return '未知'
  }
}

// 重置
const onReset = () => {
  searchInfo.value = { postId: '', status: null }
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const params = {
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  }
  // 如果postId为空字符串，不传参
  if (!params.postId) delete params.postId

  const res = await getCommunityCommentList(params)
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 多选数据
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 修改状态
const changeStatus = (row, newStatus) => {
  const statusText = getStatusText(newStatus)
  ElMessageBox.confirm(`确定要将此评论设为${statusText}吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await updateCommunityCommentStatus(row.ID, { status: newStatus })
    if (res.code === 0) {
      ElMessage.success(`操作成功`)
      getTableData()
    }
  })
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除该评论吗?此操作不可恢复!', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteCommunityComment(row.ID)
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}
</script>
