package middleware

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/utils/comm"
	"siteOl.com/stone/server/src/utils/log"
)

// OpenMiddleWare 鉴权中间件
func OpenMiddleWare(c *gin.Context) {
	// 生成请求唯一标志
	traceID := comm.TraceID()
	log.InfoTF(traceID, "OpenMiddleWare URL = %s", c.Request.URL.String())
	defer func() {
		// 退出前的追加处理
		returnJSON(c, "OpenMiddleWare", traceID)
	}()
	c.Set(constant.TraceID, traceID)
	setLang(c) // 设置语言
	// 免鉴权
	c.Next()
}

// AuthMiddleWare 鉴权中间件
func AuthMiddleWare(c *gin.Context) {
	// 生成请求唯一标志
	traceID := comm.TraceID()
	log.InfoTF(traceID, "AuthMiddleWare URL = %s", c.Request.URL.String())
	defer func() {
		// 退出前的追加处理
		returnJSON(c, "AuthMiddleWare", traceID)
	}()
	c.Set(constant.TraceID, traceID)
	setLang(c) // 设置语言
	// 读取鉴权信息
	// TODO
	c.Next()
}
