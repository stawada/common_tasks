import './assets/reset.css' //CSSリセット
import './assets/main.css'
import axios from 'axios'　//axios

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import VueCookies from "vue-cookies"

const app = createApp(App)
const base_url = 'http://localhost:8080/attendance/'

app.config.globalProperties.$axios = axios;　//グローバル変数宣言
app.config.globalProperties.BASE_URL = base_url;
app.config.globalProperties.$cookies = VueCookies;

app.use(createPinia())
app.use(router)
app.use(VueCookies)

app.mount('#app')

