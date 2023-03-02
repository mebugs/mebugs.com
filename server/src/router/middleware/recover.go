package middleware

import (
	"net/http"
	"runtime/debug"

	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/utils/logUtil"

	"github.com/gin-gonic/gin"
)

// Recover 公共Panic
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logUtil.ErrorF("panic:%s, stack:%s", err, string(debug.Stack()))
			c.JSON(http.StatusOK, resp.JsonFail("Panic"))
			c.Abort()
		}
	}()
	c.Next()
}
