package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddTopic 	godoc
// @id			AddTopic 文章专题新建
// @Summary		文章专题新建
// @Description	新建文章专题
// @Router		/blog/topic/add [post]
// @Tags		文章专题
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.TopicAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddTopic(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.TopicAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.TopicAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddTopic(traceID, req))
	}
}

// PageTopic	godoc
// @id			PageTopic 文章专题分页
// @Summary		文章专题分页
// @Description	分页查询文章专题
// @Router		/blog/topic/page [post]
// @Tags		文章专题
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.TopicPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.TopicPageRes}}	"响应成功"
func PageTopic(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.TopicPageReq{})
	if err == nil {
		req := reqObj.(*blogModel.TopicPageReq)
		// 执行查询
		service.JsonRes(c, blogService.PageTopic(traceID, req))
	}
}

// GetTopic	godoc
// @id			GetTopic 文章专题详情
// @Summary		文章专题详情
// @Description	查询文章专题详情
// @Router		/blog/topic/get [post]
// @Tags		文章专题
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.TopicGetRes}	"响应成功"
func GetTopic(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetTopic(traceID, req))
	}
}

// EditTopic 	godoc
// @id			EditTopic 文章专题编辑
// @Summary		文章专题编辑
// @Description	基于数据ID编辑文章专题
// @Router		/blog/topic/edit [post]
// @Tags		文章专题
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.TopicEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditTopic(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.TopicEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.TopicEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditTopic(traceID, req))
	}
}

// DelTopic	godoc
// @id			DelTopic 文章专题移除
// @Summary		文章专题移除
// @Description	文章专题移除处理
// @Router		/blog/topic/del [post]
// @Tags		文章专题
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelTopic(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		service.JsonRes(c, blogService.DelTopic(traceID, req))
	}
}
