import { MessagePlugin } from 'tdesign-vue-next';
import NProgress from 'nprogress'; // progress bar
import 'nprogress/nprogress.css'; // progress bar style

import { RouteRecordRaw } from 'vue-router';
import { getPermissionStore, getUserStore } from '@/store';
import router from '@/router';
import { PAGE_NOT_FOUND_ROUTE } from '@/utils/route/constant';

NProgress.configure({ showSpinner: false });

// 白名单定制在外部
const whiteListRouters = ['/login'];

router.beforeEach(async (to, from, next) => {
  NProgress.start();
  // 白名单页面不进行授权
  if (whiteListRouters.indexOf(to.path) !== -1) {
    next();
  } else {
    // 进行授权判定
    // 审视是否已登录
    const userStore = getUserStore();
    const { token } = userStore;
    if (token) {
      try {
        // 已登录，先看权限集数据
        const permissionStore = getPermissionStore();
        const { initFlag } = permissionStore;
        // 已完成路由初始化
        if (!initFlag) {
          // 读取用户权限
          const userInfo = await userStore.getUserInfo();
          // 开始初始化路由
          await permissionStore.initRoutes(userInfo.roles);
        }
        // 判定路由
        if (router.hasRoute(to.name)) {
          next();
        } else {
          next(`/`);
        }
      } catch (error) {
        MessagePlugin.error(error);
        next({
          path: '/login',
          query: { redirect: encodeURIComponent(to.fullPath) },
        });
        NProgress.done();
      }
    } else {
      // 未登录，去往登录入口
      next({
        path: '/login',
        query: { redirect: encodeURIComponent(to.fullPath) },
      });
    }
  }
  // NProgress.done();
});

router.afterEach((to) => {
  if (to.path === '/login') {
    const userStore = getUserStore();
    // const permissionStore = getPermissionStore();

    userStore.logout();
    // permissionStore.restoreRoutes();
  }
  NProgress.done();
});
