import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import App from './App.vue'
import axios from 'axios'

const app = createApp(App)
app.config.globalProperties.$axios=axios
app.use(ElementPlus)
app.mount('#app')
