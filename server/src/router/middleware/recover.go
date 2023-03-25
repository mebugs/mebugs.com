package middleware

import (
	"runtime/debug"
	"siteOl.com/stone/server/src/utils/log"

	"github.com/gin-gonic/gin"
)

// Recover 公共Panic
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.ErrorF("panic:%s, stack:%s", err, string(debug.Stack()))
			// returnJSON(c, "Recover", "Recover")
			c.Abort()
		}
	}()
	c.Next()
}
