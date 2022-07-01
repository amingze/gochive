// 不需要鉴权的业务路由
import { RouteRecordRaw } from 'vue-router';

const commonRoutes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'home',
        meta: {
            title: '',
            icon: '',
        },
        component: () => import('@/views/home/index.vue'),
    },{
        path: '/login',
        name: 'login',
        meta: {
            title: 'User Login',
            icon: '',
        },
        component: () => import('@/views/sign/Login.vue'),
    },
];

export default commonRoutes;
