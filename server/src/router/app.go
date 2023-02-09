package router

import (
	"net/http"
)

// 路由列表
var Routers = newRouter()

// 初始化路由对象
func newRouter() *Router {
	return &Router{
		routers:       make(map[string]http.Handler, 0),
		middlewareFun: make([]MDFunc, 0),
	}
}

// 初始化注册路由和中间件件
func init() {
	// 注册中间件
	Routers.Use(Auth) // 鉴权中间件
	// 注册路由 - 平台路由
	PlatformRouter()
}
