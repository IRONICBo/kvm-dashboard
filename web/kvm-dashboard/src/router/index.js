import {createRouter, createWebHistory} from 'vue-router'

// router options
const routes = [
    {
        path: '/',
        name: 'Home',
        hidden: true,
        component: () => import('@/views/Home.vue')
    },
    // {
    //     path: '/login',
    //     name: 'Login',
    //     component: Login
    // }
]

// create router
const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

// expose router
export default router