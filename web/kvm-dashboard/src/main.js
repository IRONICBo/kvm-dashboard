import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import locale from 'element-plus/es/locale/lang/zh-cn' // use chinese
import axios from 'axios'
import router from './router'

const app = createApp(App)
app.config.globalProperties.$https = axios // use axios
app.use(router).use(ElementPlus, {locale}).mount('#app') // mount the router on the app

// add element-icon
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}