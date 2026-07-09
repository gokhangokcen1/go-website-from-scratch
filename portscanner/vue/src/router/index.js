import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../pages/HomePage.vue'
import SubnettingPage from '../pages/SubnettingPage.vue'
import PortCheckerPage from '../pages/PortCheckerPage.vue'

const routes = [
  { path: '/', component: HomePage },
  { path: '/subnetting', component: SubnettingPage },
  { path: '/portchecker', component: PortCheckerPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router