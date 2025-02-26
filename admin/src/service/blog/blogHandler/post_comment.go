package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddPostComment 	godoc
// @id			AddPostComment 文章评论新建
// @Summary		文章评论新建
// @Description	新建文章评论
// @Router		/blog/postComment/add [post]
// @Tags		文章评论
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostCommentAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddPostComment(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostCommentAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostCommentAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddPostComment(traceID, req))
	}
}

// PagePostComment	godoc
// @id			PagePostComment 文章评论分页
// @Summary		文章评论分页
// @Description	分页查询文章评论
// @Router		/blog/postComment/page [post]
// @Tags		文章评论
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostCommentPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.PostCommentPageRes}}	"响应成功"
func PagePostComment(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostCommentPageReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostCommentPageReq)
		// 执行查询
		service.JsonRes(c, blogService.PagePostComment(traceID, req))
	}
}

// GetPostComment	godoc
// @id			GetPostComment 文章评论详情
// @Summary		文章评论详情
// @Description	查询文章评论详情
// @Router		/blog/postComment/get [post]
// @Tags		文章评论
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.PostCommentGetRes}	"响应成功"
func GetPostComment(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetPostComment(traceID, req))
	}
}

// EditPostComment 	godoc
// @id			EditPostComment 文章评论编辑
// @Summary		文章评论编辑
// @Description	基于数据ID编辑文章评论
// @Router		/blog/postComment/edit [post]
// @Tags		文章评论
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.PostCommentEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditPostComment(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostCommentEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostCommentEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditPostComment(traceID, req))
	}
}

// DelPostComment	godoc
// @id			DelPostComment 文章评论移除
// @Summary		文章评论移除
// @Description	文章评论移除处理
// @Router		/blog/postComment/del [post]
// @Tags		文章评论
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelPostComment(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)

		service.JsonRes(c, blogService.DelPostComment(traceID, req))
	}
}
