import type { RouteRecordRaw } from 'vue-router'

// 基础路由（免授权路由）
const base: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: 'index'
  },
  {
    path: '/index',
    name: 'index',
    component: () => import('@/views/base/index.vue')
  }
]

export default base
