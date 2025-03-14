package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// GetIndex 	godoc
// @id			GetIndex 获取首页数据
// @Summary		获取首页数据
// @Description	获取首页数据
// @Router		/page/index [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func GetIndex(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	service.JsonRes(c, blogService.GetIndex(traceID))
}
