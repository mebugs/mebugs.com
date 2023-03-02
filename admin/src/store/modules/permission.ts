// 前端 roles 控制菜单权限 通过登录后的角色对菜单就行过滤处理
// 如果需要前端 roles 控制菜单权限 请使用此文件代码替换 permission.ts 的内容

import { defineStore } from 'pinia';
import { RouteRecordRaw } from 'vue-router';
import router, { allRoutes } from '@/router';
import { store } from '@/store';

function filterPermissionsRouters(routes: Array<RouteRecordRaw>, roles: Array<unknown>) {
  const res = [];
  const removeRoutes = [];
  routes.forEach((route) => {
    const children = [];
    route.children?.forEach((childRouter) => {
      const roleCode = childRouter.meta?.roleCode || childRouter.name;
      if (roles.indexOf(roleCode) !== -1) {
        children.push(childRouter);
      } else {
        removeRoutes.push(childRouter);
      }
    });
    if (children.length > 0) {
      route.children = children;
      res.push(route);
    }
  });
  return { accessedRouters: res, removeRoutes };
}

export const usePermissionStore = defineStore('permission', {
  state: () => ({
    routers: [],
    removeRoutes: [],
    initFlag: false, // 初始化路由
  }),
  actions: {
    async initRoutes(roles: Array<unknown>) {
      let accessedRouters = [];

      let removeRoutes = [];
      // special token
      if (roles && roles.length > 0) {
        if (roles.includes('all')) {
          accessedRouters = allRoutes;
        } else {
          const res = filterPermissionsRouters(allRoutes, roles);
          accessedRouters = res.accessedRouters;
          removeRoutes = res.removeRoutes;
        }
      }

      this.routers = accessedRouters;
      this.removeRoutes = removeRoutes;

      removeRoutes.forEach((item: RouteRecordRaw) => {
        if (router.hasRoute(item.name)) {
          router.removeRoute(item.name);
        }
      });
      // 完成初始化
      this.initFlag = true;
      return this.routers;
    },
    async restore() {
      this.removeRoutes.forEach((item: RouteRecordRaw) => {
        router.addRoute(item);
      });
      // 恢复初始化
      this.initFlag = false;
    },
  },
});

export function getPermissionStore() {
  return usePermissionStore(store);
}
