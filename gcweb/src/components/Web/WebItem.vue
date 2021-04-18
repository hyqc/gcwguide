<template>
  <el-card shadow="hover" @click="openWeb" class="web-item">
    <div></div>
    <el-tooltip class="item" effect="dark" content="Top Left 提示文字" placement="top-start">
      <span>言简言</span>
    </el-tooltip>
    <div></div>
  </el-card>
</template>

<script>
export default {
  name: "WebItem",
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
    return {};
  },
  methods: {
    openWeb() {
      window.open(this.webInfo.host);
    },
  },
};
</script>

<style>
.web-item {
  transition: all 0.2s ease-in-out 0s;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-end;
  width: 100%;
  height: 100%;
  padding: 6px;
  border-radius: 4px;
  height: 60px;
  width: 120px;
  cursor: pointer;
  color: rgb(25, 61, 90);
  opacity: 1;
}
</style>
