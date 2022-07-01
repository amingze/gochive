// 需要鉴权的业务路由
import { RouteRecordRaw } from 'vue-router';

const asyncRoutes: Array<RouteRecordRaw> = [
  {
    path: '/process',
    name: 'process',
    meta: {
      title: 'Template configuration process',
      icon: '',
    },
    component: () => import('@/views/example/MarkdownPage.vue'),
  },
    {
        path: '/memo',
        name: 'memo',
        meta: {
            title: 'Memo Home',
            icon: '',
        },
        component: () => import('@/views/memo/MemoIndex.vue'),
    },
];

export default asyncRoutes;
