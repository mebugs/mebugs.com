package router

import (
	"io/ioutil"
	"siteOl.com/stone/server/src/router/subRouter"

	"siteOl.com/stone/server/src/router/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	gin.ForceConsoleColor() // 颜色日志

	router := gin.Default()
	// 公共的Panic中间件
	router.Use(middleware.Recover)
	// 开放路由
	subRouter.OpenRouter(router)
	// 平台路由
	subRouter.PlatFormRouter(router)
	return router
}
