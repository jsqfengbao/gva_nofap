<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-title">总用户数</div>
          <div class="stat-value">{{ statistics.totalUsers }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-title">活跃用户</div>
          <div class="stat-value">{{ statistics.activeUsers }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-title">总帖子数</div>
          <div class="stat-value">{{ statistics.totalPosts }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-title">总打卡天数</div>
          <div class="stat-value">{{ statistics.totalCheckins }}</div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px;">
      <el-col :span="24">
        <el-card>
          <h3>NOPEMON 戒色小程序管理后台</h3>
          <p>欢迎使用 NOPEMON 戒色小程序管理系统，请通过左侧菜单进入具体功能模块：</p>
          <ul>
            <li>👥 <strong>用户管理</strong> - 查看和管理小程序注册用户</li>
            <li>📚 <strong>学习内容</strong> - 管理戒色文章和学习资料</li>
            <li>🆘 <strong>紧急资源</strong> - 管理"好色了"点击后的紧急求助资源</li>
            <li>💬 <strong>社区管理</strong> - 审核管理用户发帖和评论</li>
            <li>🏆 <strong>成就管理</strong> - 管理系统成就</li>
          </ul>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { getNofapStatistics } from '@/plugin/nofap/api/nofap.js'
import { ref, onMounted } from 'vue'

defineOptions({
  name: 'NofapIndex'
})

const statistics = ref({
  totalUsers: 0,
  activeUsers: 0,
  totalPosts: 0,
  totalCheckins: 0
})

const loadStatistics = async () => {
  const res = await getNofapStatistics()
  if (res.code === 0) {
    statistics.value = res.data
  }
}

onMounted(() => {
  loadStatistics()
})
</script>

<style scoped>
.stat-card {
  text-align: center;
  padding: 20px;
}
.stat-title {
  font-size: 14px;
  color: #909399;
  margin-bottom: 8px;
}
.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #409eff;
}
</style>
