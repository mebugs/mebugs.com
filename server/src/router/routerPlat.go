package router

import (
	"server/src/handler"
	"server/src/handler/platHandler"
)

// 平台路由
func PlatRouter() {
	authRouter()
}

// 授权路由
func authRouter() {
	authPreFix := "/auth"
	Routers.Register(authPreFix+"/login", handler.Func(platHandler.AuthLogin)) // 登陆
}
