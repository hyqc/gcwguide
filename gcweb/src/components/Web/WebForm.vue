<template>
  <el-dialog title="编辑站点" v-model="dialogFormVisible" :width="600">
    <el-form :model="formData">
        <el-form-item label="站点头像" :label-width="formLabelWidth">
        <el-input v-model="formData.host" autocomplete="on"></el-input>
      </el-form-item>
      <el-form-item label="站点链接" :label-width="formLabelWidth">
        <el-input v-model="formData.host" autocomplete="on"></el-input>
      </el-form-item>
      <el-form-item label="站点名称" :label-width="formLabelWidth">
        <el-input v-model="formData.name" autocomplete="on"></el-input>
      </el-form-item>
      <el-form-item label="站点分组" :label-width="formLabelWidth">
        <el-input v-model="formData.name" autocomplete="on"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogFormVisible = false"
          >确 定</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>

<script>
export default {
    name: "WebForm",
    props: {
        modeShow: {
            type: Boolean,
            default: false
        },
        webItem: {
            type: Object,
            default(){
                return {}
            }
        }
    },
    data(){
        return {
            dialogFormVisible: false,
            formLabelWidth: '120px',
            formData: {
                name: '122',
                host: '111'
            }
        }
    },
    created(){
        this.dialogFormVisible = this.modeShow
    },
    methods: {},
    watch:{
        modeShow(n){
            this.dialogFormVisible = n
        }
    }
}
</script>
