package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddPost 	godoc
// @id			AddPost 文章或页面新建
// @Summary		文章或页面新建
// @Description	新建文章或页面
// @Router		/blog/post/add [post]
// @Tags		文章或页面
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddPost(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddPost(traceID, req))
	}
}

// PagePost	godoc
// @id			PagePost 文章或页面分页
// @Summary		文章或页面分页
// @Description	分页查询文章或页面
// @Router		/blog/post/page [post]
// @Tags		文章或页面
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.PostPageRes}}	"响应成功"
func PagePost(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostPageReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostPageReq)
		// 执行查询
		service.JsonRes(c, blogService.PagePost(traceID, req))
	}
}

// GetPost	godoc
// @id			GetPost 文章或页面详情
// @Summary		文章或页面详情
// @Description	查询文章或页面详情
// @Router		/blog/post/get [post]
// @Tags		文章或页面
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.PostAllRes}	"响应成功"
func GetPost(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetPost(traceID, req))
	}
}

// EditPost 	godoc
// @id			EditPost 文章或页面编辑
// @Summary		文章或页面编辑
// @Description	基于数据ID编辑文章或页面
// @Router		/blog/post/edit [post]
// @Tags		文章或页面
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditPost(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditPost(traceID, req))
	}
}
