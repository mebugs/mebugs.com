package router

import (
	"net/http"
)

// 中间件方法结构体
type MDFunc func(handler http.Handler) http.Handler

// 免鉴权白名单
var whiteApis = []string{
	"/login", // 登录入口免鉴权
}

// 基础中间件（登陆-授权）
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO 处理授权
		// TODO 白名单判定
		// TODO 读取JWT
		// TODO 读取用户的路由权限集
		// TODO 校验是否具备权限
		next.ServeHTTP(w, r)
	})
}
