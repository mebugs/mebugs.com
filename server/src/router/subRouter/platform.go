package subRouter

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/router/middleware"
	"siteOl.com/stone/server/src/sevices/plat/platHandler"
)

// PlatFormRouter 平台业务路由
func PlatFormRouter(router *gin.Engine) {
	PlatFormRouter := router.Group("/plat", middleware.AuthMiddleWare)
	{
		// 账号相关
		PlatFormRouter.POST("/accounts/add", platHandler.AddAccount)

	}
}
