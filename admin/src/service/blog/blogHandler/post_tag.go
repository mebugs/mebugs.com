package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddPostTag 	godoc
// @id			AddPostTag 文章引用标签新建
// @Summary		文章引用标签新建
// @Description	新建文章引用标签
// @Router		/blog/postTag/add [post]
// @Tags		文章引用标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostTagAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddPostTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostTagAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostTagAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddPostTag(traceID, req))
	}
}

// PagePostTag	godoc
// @id			PagePostTag 文章引用标签分页
// @Summary		文章引用标签分页
// @Description	分页查询文章引用标签
// @Router		/blog/postTag/page [post]
// @Tags		文章引用标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostTagPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.PostTagPageRes}}	"响应成功"
func PagePostTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostTagPageReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostTagPageReq)
		// 执行查询
		service.JsonRes(c, blogService.PagePostTag(traceID, req))
	}
}

// GetPostTag	godoc
// @id			GetPostTag 文章引用标签详情
// @Summary		文章引用标签详情
// @Description	查询文章引用标签详情
// @Router		/blog/postTag/get [post]
// @Tags		文章引用标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.PostTagGetRes}	"响应成功"
func GetPostTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetPostTag(traceID, req))
	}
}

// EditPostTag 	godoc
// @id			EditPostTag 文章引用标签编辑
// @Summary		文章引用标签编辑
// @Description	基于数据ID编辑文章引用标签
// @Router		/blog/postTag/edit [post]
// @Tags		文章引用标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostTagEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditPostTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostTagEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostTagEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditPostTag(traceID, req))
	}
}

// DelPostTag	godoc
// @id			DelPostTag 文章引用标签移除
// @Summary		文章引用标签移除
// @Description	文章引用标签移除处理
// @Router		/blog/postTag/del [post]
// @Tags		文章引用标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelPostTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)

		service.JsonRes(c, blogService.DelPostTag(traceID, req))
	}
}
