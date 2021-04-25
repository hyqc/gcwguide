import { createApp } from "vue";
import ElementPlus from "element-plus";
import "element-plus/lib/theme-chalk/index.css";
import App from "./App.vue";
import axios from "axios";
import "dayjs/locale/zh-cn";
import locale from "element-plus/lib/locale/lang/zh-cn";

const app = createApp(App);
app.config.globalProperties.$axios = axios;
app.use(ElementPlus, { locale });
app.mount("#app");
