<template>
  <el-container class="container">
    <el-header class="header">Header</el-header>
    <el-main class="main">
      <web-list />
    </el-main>
    <el-footer class="footer">Footer</el-footer>
  </el-container>
</template>

<script>
import WebList from "./components/WebList.vue";

export default {
  name: "App",
  components: {
    WebList,
  },
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
}
#app {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
}
.body {
  padding-bottom: 60px;
}
.container {
  width: 100%;
  min-height: 100%;
  position: relative;
  background: white;
  display: block;
}
.header,
.footer {
  height: 60px;
  width: 100%;
  display: block;
  border-bottom: 1px solid rgb(180, 184, 176);
}
.main {
  display: block;
  overflow: auto;
}
.footer {
  height: 60px;
  position: absolute;
  bottom: 0px;
  left: 0px;
  text-align: center;
  border-top: 1px solid rgb(180, 184, 176);
}
.backtop {
  height: 100vh;
  overflow-x: hidden;
}
</style>
