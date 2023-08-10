import './assets/reset.css' //CSSリセット
import './assets/main.css'
import axios from 'axios'　//axios

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.config.globalProperties.$axios = axios;　//グローバル変数宣言

app.use(createPinia())
app.use(router)

app.mount('#app')
