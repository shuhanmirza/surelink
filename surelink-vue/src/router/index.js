import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/',
        name: 'home',
        component: () => import('../components/HomeVue.vue')
    },
    {
        path: '/about',
        name: 'about',
        component: () => import('../components/AboutUs.vue')
    },
    {
        path: '/404',
        name: "notFound",
        component: () => import('../components/NotFound.vue')
    },
    {
        path: '/:link_uuid',
        name: 'redirection',
        component: () => import('../components/RedirectionConfirmation')
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
