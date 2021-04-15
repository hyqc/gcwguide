<template>
  <div>
    <el-card class="box-card" shadow="never">
      <template #header>
        <div class="card-header">
          <span>默认</span>
          <el-button class="button" type="text">编辑</el-button>
        </div>
      </template>
      <div v-for="o in 4" :key="o" class="text item">
        <el-row :gutter="48">
          <el-col :span="6">
            <el-card shadow="hover" class="web-item">
              总是显示3
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover" class="web-item">
              总是显示4
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover" class="web-item">
              鼠标悬浮时显示5
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover" class="web-item">
              从不显示6
            </el-card>
          </el-col>
        </el-row>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
  props: {
    msg: String,
  },
};
</script>

<style scoped>
.web-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
  text-align: center;
}
.el-row {
  margin: 2rem;
}
</style>
