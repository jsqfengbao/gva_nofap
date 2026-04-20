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
        <el-form-item label="昵称/昵称">
          <el-input v-model="searchInfo.nickname" placeholder="搜索用户昵称" clearable />
        </el-form-item>

        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="用户状态" clearable>
            <el-option label="全部" :value="null" />
            <el-option label="正常" :value="1" />
            <el-option label="禁用" :value="0" />
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

        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="头像" prop="avatarUrl" width="80">
          <template #default="scope">
            <el-avatar :size="40" :src="scope.row.avatarUrl" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="昵称" prop="nickname" width="150" />
        <el-table-column align="left" label="OpenID" prop="openId" min-width="180" />
        <el-table-column align="left" label="连续打卡" prop="currentStreak" width="90">
          <template #default="scope">
            <el-tag type="primary">{{ scope.row.currentStreak }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="最长打卡" prop="longestStreak" width="90" />
        <el-table-column align="left" label="状态" prop="status" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="注册时间" prop="CreatedAt" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="160"
        >
          <template #default="scope">
            <el-button
              v-if="scope.row.status === 1"
              type="danger"
              link
              icon="lock"
              class="table-button"
              @click="toggleStatus(scope.row)"
            >
              禁用
            </el-button>
            <el-button
              v-else
              type="success"
              link
              icon="unlock"
              class="table-button"
              @click="toggleStatus(scope.row)"
            >
              启用
            </el-button>
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
  getNofapUserList,
  updateNofapUserStatus,
} from '@/plugin/nofap/api/nofap.js'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'NofapUsers'
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  nickname: '',
  status: null
})

const elSearchFormRef = ref()

// 重置
const onReset = () => {
  searchInfo.value = { nickname: '', status: null }
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
  const res = await getNofapUserList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
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

// 切换用户状态
const toggleStatus = (row) => {
  const newStatus = row.status === 1 ? 0 : 1
  const actionText = newStatus === 1 ? '启用' : '禁用'
  ElMessageBox.confirm(`确定要${actionText}该用户吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await updateNofapUserStatus(row.ID, { status: newStatus })
    if (res.code === 0) {
      ElMessage.success(`${actionText}成功`)
      getTableData()
    }
  })
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除该用户吗?此操作不可恢复!', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    // TODO: 需要后端实现删除接口
    ElMessage.warning('删除功能需要后端API支持')
  })
}
</script>
