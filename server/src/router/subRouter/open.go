package subRouter

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/router/middleware"
	"siteOl.com/stone/server/src/sevices/plat/platHandler"
)

// OpenRouter Open 开放路由
func OpenRouter(router *gin.Engine) {
	PlatFormRouter := router.Group("/open", middleware.OpenMiddleWare)
	{
		// 开放账密登陆
		PlatFormRouter.POST("/auth/login", platHandler.AuthLogin)

	}
}
