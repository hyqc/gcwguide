<template>
  <div class="container">
    <web-list />
    <div class="footer">
      <p>
        <a href="https://www.jianean.com/" target="blank">由简言提供</a>
      </p>
      <p>
        <a href="https://www.jianean.com/" target="blank"
          >https://www.jianean.com</a
        >
      </p>
    </div>
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
  background-color: #e9e9eb;
}
.container {
  width: 100%;
  min-height: 100%;
  position: relative;
  display: block;
}
.footer {
  text-align: center;
}
.footer a {
  text-decoration: none;
}
</style>
