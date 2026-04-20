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
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="搜索成就名称" clearable />
        </el-form-item>

        <el-form-item label="分类">
          <el-select v-model="searchInfo.category" placeholder="选择分类" clearable>
            <el-option label="全部" :value="null" />
            <el-option label="坚持打卡" value="streak" />
            <el-option label="里程碑" value="milestone" />
            <el-option label="特殊" value="special" />
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
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">
          新增
        </el-button>
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
        >
          删除
        </el-button>
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

        <el-table-column align="left" label="ID" prop="ID" width="70" />
        <el-table-column align="left" label="图标" prop="icon" width="80">
          <template #default="scope">
            <i :class="scope.row.icon" style="font-size: 24px; color: #409eff;"></i>
          </template>
        </el-table-column>
        <el-table-column align="left" label="名称" prop="name" width="150" />
        <el-table-column align="left" label="描述" prop="description" min-width="180" />
        <el-table-column align="left" label="分类" prop="category" width="100">
          <template #default="scope">
            <el-tag>{{ scope.row.category }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="解锁条件" prop="requirement" width="120">
          <template #default="scope">
            {{ scope.row.requirement }}天
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" prop="sortOrder" width="80" />
        <el-table-column align="left" label="状态" prop="status" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="180"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateInfoFunc(scope.row)"
            >
              编辑
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

    <el-drawer
      v-model="dialogFormVisible"
      destroy-on-close
      size="800"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加成就' : '编辑成就' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog"> 确 定 </el-button>
            <el-button @click="closeDialog"> 取 消 </el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="top"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="名称:" prop="name" required>
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入成就名称"
          />
        </el-form-item>
        <el-form-item label="描述:" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="成就描述"
          />
        </el-form-item>
        <el-form-item label="图标Class:" prop="icon">
          <el-input
            v-model="formData.icon"
            placeholder="el-icon-xxx"
          />
          <div style="font-size: 12px; color: #909399; margin-top: 5px;">使用 Element Plus icon class</div>
        </el-form-item>
        <el-form-item label="分类:" prop="category">
          <el-select v-model="formData.category" placeholder="选择分类" style="width: 100%">
            <el-option label="坚持打卡" value="streak" />
            <el-option label="里程碑" value="milestone" />
            <el-option label="特殊成就" value="special" />
          </el-select>
        </el-form-item>
        <el-form-item label="解锁天数:" prop="requirement" required>
          <el-input-number v-model="formData.requirement" :min="1" :max="36500" />
          <div style="font-size: 12px; color: #909399; margin-top: 5px;">需要连续打卡多少天解锁</div>
        </el-form-item>
        <el-form-item label="排序:" prop="sortOrder">
          <el-input-number v-model="formData.sortOrder" :min="0" :max="9999" />
          <div style="font-size: 12px; color: #909399; margin-top: 5px;">数字越小越靠前</div>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getAchievementList,
  createAchievement,
  updateAchievement,
  deleteAchievement,
} from '@/plugin/nofap/api/nofap.js'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'NofapAchievements'
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  name: '',
  category: null
})

const elSearchFormRef = ref()

// 重置
const onReset = () => {
  searchInfo.value = { name: '', category: null }
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
  const res = await getAchievementList({
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

// 验证规则
const rule = reactive({
  name: [{ required: true, message: '请输入成就名称', trigger: 'blur' }],
  requirement: [{ required: true, message: '请输入解锁天数', trigger: 'blur' }]
})

// 多选数据
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 弹窗相关
const type = ref('')
const dialogFormVisible = ref(false)
const elFormRef = ref()
const formData = ref({
  name: '',
  description: '',
  icon: '',
  category: 'streak',
  requirement: 7,
  sortOrder: 100,
  status: 1
})

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    description: '',
    icon: '',
    category: 'streak',
    requirement: 7,
    sortOrder: 100,
    status: 1
  }
}

// 更新行
const updateInfoFunc = async (row) => {
  type.value = 'update'
  formData.value = { ...row }
  dialogFormVisible.value = true
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除这个成就吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteAchievement(row.ID)
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    if (type.value === 'create') {
      res = await createAchievement(formData.value)
    } else {
      res = await updateAchievement(formData.value)
    }
    if (res.code === 0) {
      ElMessage.success(`${type.value === 'create' ? '创建' : '更新'}成功`)
      closeDialog()
      getTableData()
    }
  })
}
</script>
