<template>
  <div class="web-item">
    <div
      class="web-item-avatar"
      @click="openWeb"
      @mouseover="showBackgroundEvent(true)"
      @mouseout="showBackgroundEvent(false)"
    >
      <img src="/favicon.ico" class="web-item-avatar-img" />
    </div>
    <div class="web-item-title" @click="openWeb">
      网易
    </div>

    <el-dropdown
      @command="webItemToolEvent"
      trigger="click"
      class="web-item-tool"
    >
      <i class="el-icon-more"></i>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item icon="el-icon-edit-outline" @click="editWebItemEvent">编辑</el-dropdown-item>
          <el-dropdown-item icon="el-icon-delete">删除</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
    <div
      class="web-item-background"
      :class="showBackgroundClass"
    ></div>
    <WebForm :mode-show="modeShow" />
  </div>
</template>

<script>
import WebForm from "@/components/Web/WebForm.vue";

export default {
  name: "WebItem",
  components: {
    WebForm,
  },
  props: {
    webInfo: {
      type: Object,
      default() {
        return {
          id: "",
          group: "default",
          name: "简言",
          pic: "https://www.jianean.com/favicon.ico",
          host: "https://www.jianean.com",
          desc: "简言",
        };
      },
    },
  },
  data() {
    return {
      showBlack: false,
      modeShow: false
    };
  },
  computed: {
    showBackgroundClass() {
      return this.showBlack
        ? "web-item-background-show"
        : "web-item-background-hide";
    },
  },
  methods: {
    openWeb() {
      window.open(this.webInfo.host);
    },
    showBackgroundEvent(show) {
      this.showBlack = !!show;
    },
    webItemToolEvent(){

    },
    editWebItemEvent() {
      this.modeShow = true
    },
  },
};
</script>

<style>
.web-item {
  position: relative;
  display: block;
  border-radius: 4px;
  height: 100px;
  width: 100px;
  cursor: pointer;
}
.web-item-background {
  position: absolute;
  top: 0;
  left: 0;
  display: block;
  border-radius: 4px;
  height: 100px;
  width: 100px;
  cursor: pointer;
  background-color: gray;
  z-index: 1;
}
.web-item-background:hover {
  opacity: 0.6;
}
.web-item-background-show {
  opacity: 0.6;
}
.web-item-background-hide {
  opacity: 0;
}
.web-item-avatar {
  position: absolute;
  margin: 22px;
  padding: auto;
  display: block;
  border-radius: 4px;
  height: 50px;
  width: 50px;
  cursor: pointer;
  background-color: white;
  z-index: 999;
}
.web-item-avatar-img {
  margin: 7px;
  height: 36px;
  width: 36px;
}
.web-item-title {
  position: absolute;
  bottom: 8px;
  right: 4px;
  text-align: center;
  width: 98px;
  color: white;
  font-size: 12px;
  z-index: 999;
}
.web-item-tool {
  position: absolute;
  top: 6px;
  right: 10px;
  width: 10px;
  color: white;
  display: none;
  z-index: 999;
  padding: 2px;
}
.web-item:hover .web-item-tool {
  display: block;
}
</style>
