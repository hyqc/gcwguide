<template>
  <div class="container">
    <web-list />
  </div>
</template>

<script>
import WebList from "@/components/Web/WebList.vue";

export default {
  name: "App",
  components: {
    WebList,
  },
};
</script>

<style>
body {
  /* background: url(/background.jpg) no-repeat fixed top; */
  background-color: #E9E9EB;
}
.container {
  width: 100%;
  min-height: 100%;
  position: relative;
  display: block;
}
</style>
