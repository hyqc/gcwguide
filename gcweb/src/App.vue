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
  background: url(/background.jpg) no-repeat fixed top;
}
.container {
  width: 100%;
  min-height: 100%;
  position: relative;
  display: block;
}</style
>>
