package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/blogModel"
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

// GetPosts 	godoc
// @id			GetPosts 获取文章数据
// @Summary		获取文章数据
// @Description	获取文章数据
// @Router		/page/posts [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func GetPosts(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostsReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostsReq)
		// 执行创建
		service.JsonRes(c, blogService.GetPosts(traceID, req))
	}
}
