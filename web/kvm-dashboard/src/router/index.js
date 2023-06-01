import {createRouter, createWebHistory} from 'vue-router'

// router options
const routes = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/home',
        name: 'Home',
        component: () => import('@/views/Home.vue')
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/Login.vue')
    }
]

// create router
const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

// expose router
export default router