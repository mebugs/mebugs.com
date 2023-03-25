package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/validate"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

//  登陆授权

// AuthLogin /open/auth/login 开放账密登陆
func AuthLogin(c *gin.Context) {
	// TraceID 日志追踪
	traceID := c.GetString(constant.TraceID)
	// 参数读取
	req := &platModel.AuthLogin{}
	// 校验并且 解析请求数据
	err, reqObj := validate.Readable(c, req)
	if err != nil {
		return
	}
	req = reqObj.(*platModel.AuthLogin)
	// 执行登陆
	res := platService.Login(traceID, req)
	c.Set(constant.RespBody, res)
}
