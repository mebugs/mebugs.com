package router

import (
	"server/src/sevices/handler"
	"server/src/sevices/handler/platformHandler"
)

// 平台路由
func PlatformRouter() {
	authRouter()
}

// 授权路由
func authRouter() {
	authPreFix := "/platform"
	Routers.Register(authPreFix+"/auth/login", handler.Func(platformHandler.AuthLogin)) // 登陆
}
