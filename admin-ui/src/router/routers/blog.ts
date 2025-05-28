import { MAIN_LAYOUT } from './base'
import type { RouteRecordRaw } from 'vue-router'

// 博客路由
const BLOG: RouteRecordRaw = {
  path: '/blog', // 博客管理
  name: 'Blog', // 充当权限Alias
  component: MAIN_LAYOUT, // 布局
  meta: { locale: 'menu.Blog', icon: 'icon-book', requiresAuth: true }, // 国际化、图标、鉴权标识
  children: [
    {
      path: 'source', // 资源
      name: 'BlogSource',
      component: () => import('@/views/blog/source/index.vue'),
      meta: { locale: 'menu.BlogSource', requiresAuth: true }
    },
    {
      path: 'banner', // Banner
      name: 'BlogBanner',
      component: () => import('@/views/blog/banner/index.vue'),
      meta: { locale: 'menu.BlogBanner', requiresAuth: true }
    },
    {
      path: 'category', // 分类
      name: 'BlogCategory',
      component: () => import('@/views/blog/category/index.vue'),
      meta: { locale: 'menu.BlogCategory', requiresAuth: true }
    },
    {
      path: 'tag', // 标签
      name: 'BlogTag',
      component: () => import('@/views/blog/tag/index.vue'),
      meta: { locale: 'menu.BlogTag', requiresAuth: true }
    },
    {
      path: 'post', // 文章
      name: 'BlogPost',
      component: () => import('@/views/blog/post/index.vue'),
      meta: { locale: 'menu.BlogPost', requiresAuth: true }
    },
    {
      path: 'comment', // 评论
      name: 'BlogComment',
      component: () => import('@/views/blog/comment/index.vue'),
      meta: { locale: 'menu.BlogComment', requiresAuth: true }
    },
    {
      path: 'links', // 评论
      name: 'BlogLinks',
      component: () => import('@/views/blog/links/index.vue'),
      meta: { locale: 'menu.BlogLinks', requiresAuth: true }
    }
  ]
}

export default BLOG
