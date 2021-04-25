<template>
  <div>
    <el-card
      class="box-card"
      v-for="(webList, group) in webListObject"
      shadow="hover"
      :key="group"
    >
      <template #header>
        <div class="card-header">
          <span class="group-info"
            ><i class="el-icon-s-tools"></i>&nbsp;&nbsp;<span class="group-text">{{ group }}</span></span
          >
          <el-button class="button" type="text" @click="addWeb"
            ><i class="el-icon-circle-plus"></i
          ></el-button>
        </div>
      </template>
      <web-group :webList="webList" :webGroups="webGroups" />
    </el-card>
    <web-form
      :modeShow="modeShow"
      :webGroups="webGroups"
      @showModel="showModel"
    />
  </div>
</template>

<script>
import WebGroup from "@/components/Web/WebGroup.vue";
import WebForm from "@/components/Web/WebForm.vue";
import { WebList, WebGroups } from "@/api/file.js";

export default {
  name: "WebList",
  components: {
    WebGroup,
    WebForm
  },
  data() {
    return {
      webListObject: {},
      webGroups: [],
      modeShow: false
    };
  },
  created() {
    this.init();
  },
  methods: {
    init() {
      this.getWebList();
      this.getWebGroups();
    },
    getWebList() {
      WebList()
        .then((res) => {
          res = res.data;
          console.log(res);
          if (res.code) {
            this.$message.error(
              res.message || "获取站点失败，请联系系统管理员！"
            );
          } else {
            this.webListObject = res.data.rows;
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    getWebGroups() {
      WebGroups()
        .then((res) => {
          res = res.data;
          console.log(res);
          if (res.code) {
            this.$message.error(
              res.message || "获取站点失败，请联系系统管理员！"
            );
          } else {
            this.webGroups = res.data || [];
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    addWeb() {
      this.modeShow = true
    },
    showModel(data){
      this.modeShow = data.show
      window.location.reload()
    }
  },
};
</script>

<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  max-width: 1040px;
  margin: 2rem auto;
}
.el-card__header {
  padding: 4px 20px;
}
.el-icon-s-tools {
  color: #409eff;
}
.group-info {
  color: #409eff;
}
.group-text{
  padding: auto 1rem;
  font-size: 0.9rem;
}
</style>
