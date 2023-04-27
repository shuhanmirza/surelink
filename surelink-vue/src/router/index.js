import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RedirectView from "@/views/RedirectView.vue";
import { pathToRegexp } from 'path-to-regexp';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/:link([a-zA-Z]+)',
    name: 'redirect',
    component: RedirectView
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
