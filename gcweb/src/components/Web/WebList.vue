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
            ><i class="el-icon-s-tools"></i>&nbsp;&nbsp;<span
              class="group-text"
              >{{ group }}</span
            ></span
          >
          <span>
            <el-button class="button" type="text" @click="login"
              ><i :class="loginIcon" :title="canEdit ? '退出登录' : '登录'"></i
            ></el-button>
            <el-button class="button" type="text" @click="addWeb" title="添加站点"
              ><i class="el-icon-circle-plus"></i
            ></el-button>
          </span>
        </div>
      </template>
      <web-group :webList="webList" :webGroups="webGroups" :canEdit="canEdit" />
    </el-card>
    <web-form
      :modeShow="modeShow"
      :webGroups="webGroups"
      @showModel="showModel"
    />

    <login-form :modeShow="modeLoginForm" @showModel="showLoginModel" />
  </div>
</template>

<script>
import WebGroup from "@/components/Web/WebGroup.vue";
import WebForm from "@/components/Web/WebForm.vue";
import LoginForm from "@/components/Auth/LoginForm.vue";
import { WebList, WebGroups } from "@/api/website.js";
import { CookieGetToken, CookieRemoveToken } from "@/api/cookie.js";

export default {
  name: "WebList",
  components: {
    WebGroup,
    WebForm,
    LoginForm,
  },
  data() {
    return {
      webListObject: {},
      webGroups: [],
      modeShow: false,
      modeLoginForm: false,
      canEdit: false,
      token: "",
      loginIcon: ""
    };
  },
  created() {
    this.init();
  },
  methods: {
    init() {
      this.token = CookieGetToken();
      this.isCanEdit()
      this.setLoginIcon()
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
      if (this.token) {
        this.modeShow = true;
      } else {
        this.modeLoginForm = true;
      }
    },
    login() {
      if (this.canEdit) {
        const _this = this;
        this.$confirm("确定要退出登录吗？").then(() => {
          CookieRemoveToken();
          CookieGetToken();
          _this.isCanEdit();
          _this.setLoginIcon()
        });
      } else {
        this.modeLoginForm = true;
      }
    },
    showModel(data) {
      this.modeShow = data.show;
      window.location.reload();
    },
    showLoginModel(data) {
      this.modeLoginForm = data.show;
      this.token = CookieGetToken();
      this.setLoginIcon()
      this.isCanEdit()
    },
    isCanEdit() {
      return (this.canEdit = !!CookieGetToken());
    },
    setLoginIcon() {
      return this.loginIcon = CookieGetToken() ? "el-icon-switch-button" : "el-icon-user-solid";
    },
  },
  watch: {
    token() {
      this.setLoginIcon()
      this.isCanEdit()
    },
    canEdit(n) {
      console.log(n)
    },
    loginIcon(n){
      console.log(n)
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
.group-text {
  padding: auto 1rem;
  font-size: 0.9rem;
}
</style>
