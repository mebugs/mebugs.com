package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/data/validate"
	"siteOl.com/stone/server/src/utils/log"
)

/**
 *
 * 账号控制器
 *
 *
 * @author 米虫@mebugs.com
 * @since 2023-03-08
 */

type Test struct {
	A string `json:"a" binding:"required"`
	B string `json:"b" binding:"required"`
}

// AddAccount 创建账号
func AddAccount(c *gin.Context) {
	// TraceID 日志追踪
	traceID := c.GetString(constant.TraceID)
	// TODO 获取登录用户对象
	//loginUser := plat.GetLoginUser(c)
	//if loginUser == nil {
	//	c.Set(constant.RespBody, resp.AuthErr)
	//	return
	//}

	// 参数读取
	req := &Test{}
	// 校验并且 解析请求数据
	err, reqObj := validate.Readable(c, req)
	if err != nil {
		return
	}
	req = reqObj.(*Test)
	log.InfoTF(traceID, "%v", req)
	//// 查询深度 不传0 不查询子集 1查询本级+子集 2查询本级+2层子集 3查询本级+3层子集
	//// 最大查3层
	//if bundleReq.Depth > 3 {
	//	bundleReq.Depth = 3
	//}
	//
	//res, err := bundleService.GetBundle(bundleReq, userCard, traceID)
	//if err != nil {
	//	// 权益包不存在
	//	if err.Error() == model.ErrorNotFound {
	//		c.Set(model.RespBody, model.ResponseBundleNotFound)
	//		return
	//	}
	//	// 数据模版不存在
	//	if err.Error() == model.ErrPageModDataNotFound {
	//		c.Set(model.RespBody, model.ResponseErrPageModDataNotFound)
	//		return
	//	}
	//	c.Set(model.RespBody, model.ResponseSystemError)
	//	return
	//}
	c.Set(constant.RespBody, resp.OK)
	return
}
