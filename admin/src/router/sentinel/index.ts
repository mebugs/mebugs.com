/**
 * 路由哨兵
 * 监听路由变化
 * 传递路由通信
 * 判定路由权限
 */
import { setRouteEmitter } from '@/utils/routeListener'
import type { Router } from 'vue-router'

// 路由哨兵发布路由变化消息
function routerSentinelPush(router: Router) {
  router.beforeEach(async (to) => {
    // emit route change
    setRouteEmitter(to)
  })
}

// 路由哨兵检查路由访问 - 登陆 - 授权
function routerSentinelCheck(router: Router) {
  router.beforeEach(async (to, from, next) => {
    console.log(from + '>>' + to)
    next()
  })
}

// 创建路由处理中间件
export default function routerSentinel(router: Router) {
  routerSentinelPush(router) // 路由建传递路有变化信息
  routerSentinelCheck(router) // 验证是否有权限
}
