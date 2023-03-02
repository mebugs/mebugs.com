package router

import (
	"io/ioutil"

	"siteOl.com/stone/server/src/router/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	gin.ForceConsoleColor()

	router := gin.Default()
	// 公共的Panic中间件
	router.Use(middleware.Recover)

	return router
}
