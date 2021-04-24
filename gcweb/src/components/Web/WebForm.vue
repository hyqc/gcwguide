<template>
  <el-dialog title="编辑站点" v-model="dialogFormVisible" :width="600">
    <el-form :model="formData" :rules="rules" ref="formData">
      <el-form-item label="上传图片" :label-width="formLabelWidth" prop="pic">
        <el-upload
          :action="upload.url"
          list-type="picture-card"
          :on-preview="uploadPreview"
          :on-remove="uploadRemove"
          :on-success="uploadSuccess"
          :accept="accept"
          :limit="1"
        >
          <i class="el-icon-plus"></i>
        </el-upload>
        <el-dialog v-model="dialogVisible">
          <img width="100%" :src="dialogImageUrl" alt="" />
        </el-dialog>
      </el-form-item>
      <el-form-item
        label="站点图片地址"
        :label-width="formLabelWidth"
        prop="pic"
      >
        <el-input v-model="formData.pic" autocomplete="on"></el-input>
      </el-form-item>
      <el-form-item label="站点名称" :label-width="formLabelWidth" prop="name">
        <el-input v-model="formData.name" autocomplete="on"></el-input>
      </el-form-item>
      <el-form-item label="站点链接" :label-width="formLabelWidth" prop="host">
        <el-input v-model="formData.host" autocomplete="on"></el-input>
      </el-form-item>
      <el-form-item label="站点分组" :label-width="formLabelWidth" prop="group">
        <el-select v-model="formData.group" placeholder="请选择站点分组">
          <el-option
            v-for="(item, key) in groupsList"
            :label="item"
            :value="item"
            :key="key"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="站点描述" :label-width="formLabelWidth" prop="desc">
        <el-input v-model="formData.desc" autocomplete="on"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-form-item class="dialog-footer">
        <el-button type="primary" @click="submit('formData')">保存</el-button>
        <el-button @click="resetForm('formData')">重置</el-button>
      </el-form-item>
    </template>
  </el-dialog>
</template>

<script>
import { WebAdd, WebGroups, WebLogoUploadUrl } from "@/api/file.js";

export default {
  name: "WebForm",
  props: {
    modeShow: {
      type: Boolean,
      default: false,
    },
    webItem: {
      type: Object,
      default() {
        return {
          id: "",
          group: "",
          pic: "",
          name: "",
          host: "",
          desc: "",
        };
      },
    },
  },
  data() {
    return {
      dialogFormVisible: false,
      formLabelWidth: "120px",
      formData: null,
      upload: {
        url: WebLogoUploadUrl,
        accept: "image/*",
        dialogImageUrl: "",
      },
      rules: {
        name: [
          { required: true, message: "站点名称不能为空", trigger: "blur" },
        ],
        host: [
          { required: true, message: "站点地址不能为空", trigger: "blur" },
          { type: "url", message: "不是有效的地址", trigger: "blur" },
        ],
      },
      groupsList: [],
    };
  },
  created() {
    this.init();
  },
  methods: {
    init() {
      this.dialogFormVisible = this.modeShow;
      this.formData = JSON.parse(JSON.stringify(this.webItem));
      this.getWebsitesGroups();
    },
    submit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const _this = this;
          WebAdd(this.formData)
            .then((res) => {
              res = res.data;
              if (res.code) {
                this.$message.error(res.message || "保存失败");
              } else {
                this.$message.success({
                  message: res.message || "保存成功",
                  onClose() {
                    _this.dialogFormVisible = false;
                  },
                });
              }
            })
            .catch((err) => {
              console.log(err);
              this.$message.error("保存失败");
            });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    uploadRemove(file, fileList) {
      console.log(file, fileList);
      this.formData.pic = "";
    },
    uploadPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    uploadSuccess(res) {
      if (res.code) {
        this.$message.error(res.message || "上传失败，请稍后再试！");
      } else {
        this.formData.pic = res.data.url;
        this.dialogImageUrl = this.formData.pic;
      }
    },
    getWebsitesGroups() {
      WebGroups().then((res) => {
        res = res.data;
        this.groupsList = res.data || [];
      });
    },
  },
  watch: {
    modeShow(n) {
      this.dialogFormVisible = n;
    },
  },
};
</script>

<style scoped>
.avatar-uploader .el-upload {
  border: 1px dashed #524e4e;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 18px;
  color: #8c939d;
  width: 30px;
  height: 30px;
  text-align: center;
}
.avatar {
  width: 30px;
  height: 30px;
  display: block;
}
</style>
