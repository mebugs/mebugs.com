package router

import (
	"server/src/sevices/handler"
	"server/src/sevices/handler/platHandler"
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
