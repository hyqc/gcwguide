<template>
  <div class="group-card">
    <el-row :gutter="20">
      <el-col :xs="8" :sm="6" :md="4" :lg="3" :xl="3" v-for="(webInfo, index) in webList" :key="index">
        <web-item :webInfo="webInfo" :webGroups="webGroups" :canEdit="canEdit" />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import WebItem from "@/components/Web/WebItem.vue";

export default {
  name: "WebGroup",
  components: {
    WebItem,
  },
  props: {
    webList: {
      type: Array,
      default() {
        return [];
      },
    },
    webGroups: {
      type: Array,
      default() {
        return [];
      },
    },
    canEdit: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    return {
    };
  },
};
</script>

<style scoped>
.group-card {
  display: block;
  max-width: 1068px;
  height: 100%;
  z-index: 1;
  background-color: #FAFAFA;
}
</style>
