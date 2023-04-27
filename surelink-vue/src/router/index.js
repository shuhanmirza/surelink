import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RedirectView from "@/views/RedirectView.vue";
import ErrorView from "@/views/ErrorView.vue";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/:link([A-Za-z0-9=]{6})',
    name: 'redirect',
    component: RedirectView
  },
  {
    path: '*',
    name: 'error',
    component: ErrorView
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
