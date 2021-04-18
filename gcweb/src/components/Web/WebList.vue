<template>
  <div class="group-card">
    <web-item />
  </div>
</template>

<script>
import WebItem from "@/components/Web/WebItem.vue";

export default {
  name: "WebList",
  components: {
    WebItem,
  }
};
</script>

<style>
.group-card{
  max-width: 1068px;
  margin: 5rem auto;
  z-index: 10;
  border: 1px solid rgb(124, 124, 126);
}
</style>
