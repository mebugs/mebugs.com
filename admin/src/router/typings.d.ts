import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    needAuth: boolean // 是否需要授权
    icon?: string // 图标
    locale?: string // 国际化文言
    hideInMenu?: boolean // 不显示在菜单
    hideChildrenInMenu?: boolean // 隐藏子节点
    order?: number // 菜单排序（暂不使用）
    noAffix?: boolean // 不显示在顶部Tab中
    ignoreCache?: boolean // 不进行路由缓存
  }
}
