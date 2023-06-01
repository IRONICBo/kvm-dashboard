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
    },
    {
        path: '/config',
        name: 'Config',
        component: () => import('@/views/kvm_control/Config.vue'),
        redirect: '/host/manage',
        children: [
            {
                path: '/domain/create',
                name: 'DomainCreate',
                component: () => import('@/views/kvm_control/DomainCreate.vue')
            },
            {
                path: '/domain/manage',
                name: 'DomainManage',
                component: () => import('@/views/kvm_control/DomainManage.vue')
            },
            {
                path: '/host/manage',
                name: 'HostManage',
                component: () => import('@/views/kvm_control/HostManage.vue')
            },
        ]
    }
]

// create router
const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

router.beforeEach((to, from, next) => {
    console.log('beforeEach triggered');
    console.log('From:', from.path);
    console.log('To:', to.path);
    next();
  });
  
  router.afterEach((to, from) => {
    console.log('afterEach triggered');
    console.log('From:', from.path);
    console.log('To:', to.path);
  });
  

// expose router
export default router