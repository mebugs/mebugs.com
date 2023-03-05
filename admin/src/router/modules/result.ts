import Layout from '@/layouts/index.vue';

export default [
  {
    path: '/list',
    name: 'list',
    component: Layout,
    redirect: '/list/base',
    meta: { title: '列表', icon: 'check-circle' },
    children: [
      {
        path: 'base',
        name: 'ListBase',
        component: () => import('@/pages/list/base/index.vue'),
        meta: { title: '默认列表' },
      },
      {
        path: 'card',
        name: 'Listcard',
        component: () => import('@/pages/list/card/index.vue'),
        meta: { title: '卡片列表' },
      },
      {
        path: 'filter',
        name: 'Listfilter',
        component: () => import('@/pages/list/filter/index.vue'),
        meta: { title: '过滤列表' },
      },
    ],
  },
  {
    path: '/form',
    name: 'form',
    component: Layout,
    redirect: '/form/base',
    meta: { title: '表单', icon: 'check-circle' },
    children: [
      {
        path: 'base',
        name: 'FormBase',
        component: () => import('@/pages/form/base/index.vue'),
        meta: { title: '默认表单' },
      },
      {
        path: 'step',
        name: 'FormStep',
        component: () => import('@/pages/form/step/index.vue'),
        meta: { title: '分布表单' },
      },
    ],
  },
  {
    path: '/detail',
    name: 'detail',
    component: Layout,
    redirect: '/detail/base',
    meta: { title: '详情', icon: 'check-circle' },
    children: [
      {
        path: 'base',
        name: 'DeBase',
        component: () => import('@/pages/detail/base/index.vue'),
        meta: { title: '基础信息' },
      },
      {
        path: 'adv',
        name: 'Deadv',
        component: () => import('@/pages/detail/advanced/index.vue'),
        meta: { title: 'ADV信息' },
      },
      {
        path: 'dep',
        name: 'Dedep',
        component: () => import('@/pages/detail/deploy/index.vue'),
        meta: { title: 'DEP信息' },
      },
      {
        path: 'sec',
        name: 'DeSec',
        component: () => import('@/pages/detail/secondary/index.vue'),
        meta: { title: 'SEC信息' },
      },
    ],
  },
  {
    path: '/result',
    name: 'result',
    component: Layout,
    redirect: '/result/success',
    meta: { title: '结果页', icon: 'check-circle' },
    children: [
      {
        path: 'success',
        name: 'ResultSuccess',
        component: () => import('@/pages/result/success/index.vue'),
        meta: { title: '成功页' },
      },
      {
        path: 'fail',
        name: 'ResultFail',
        component: () => import('@/pages/result/fail/index.vue'),
        meta: { title: '失败页' },
      },
      {
        path: 'network-error',
        name: 'ResultNetworkError',
        component: () => import('@/pages/result/network-error/index.vue'),
        meta: { title: '网络异常' },
      },
      {
        path: '403',
        name: 'Result403',
        component: () => import('@/pages/result/403/index.vue'),
        meta: { title: '无权限' },
      },
      {
        path: '404',
        name: 'Result404',
        component: () => import('@/pages/result/404/index.vue'),
        meta: { title: '访问页面不存在页' },
      },
      {
        path: '500',
        name: 'Result500',
        component: () => import('@/pages/result/500/index.vue'),
        meta: { title: '服务器出错页' },
      },
      {
        path: 'browser-incompatible',
        name: 'ResultBrowserIncompatible',
        component: () => import('@/pages/result/browser-incompatible/index.vue'),
        meta: { title: '浏览器不兼容页' },
      },
      {
        path: 'maintenance',
        name: 'ResultMaintenance',
        component: () => import('@/pages/result/maintenance/index.vue'),
        meta: { title: '系统维护页' },
      },
    ],
  },
];
