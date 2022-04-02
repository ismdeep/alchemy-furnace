import {createApp} from 'vue'
import App from './App.vue'
import {createRouter, createWebHistory} from "vue-router";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import Home from './views/Home.vue'
import About from './views/About.vue'
import Login from './views/Login.vue'

const routes = [
  {path: '/', component: Home},
  {path: '/about', component: About},
  {path: '/login', component: Login},
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

createApp(App).use(ElementPlus).use(router).mount('#app')
