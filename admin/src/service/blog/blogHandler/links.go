package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddLinks 	godoc
// @id			AddLinks 友情链接新建
// @Summary		友情链接新建
// @Description	新建友情链接
// @Router		/blog/links/add [post]
// @Tags		友情链接
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.LinksAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddLinks(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.LinksAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.LinksAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddLinks(traceID, req))
	}
}

// PageLinks	godoc
// @id			PageLinks 友情链接分页
// @Summary		友情链接分页
// @Description	分页查询友情链接
// @Router		/blog/links/page [post]
// @Tags		友情链接
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.LinksPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.LinksPageRes}}	"响应成功"
func PageLinks(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.LinksPageReq{})
	if err == nil {
		req := reqObj.(*blogModel.LinksPageReq)
		// 执行查询
		service.JsonRes(c, blogService.PageLinks(traceID, req))
	}
}

// GetLinks	godoc
// @id			GetLinks 友情链接详情
// @Summary		友情链接详情
// @Description	查询友情链接详情
// @Router		/blog/links/get [post]
// @Tags		友情链接
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.LinksGetRes}	"响应成功"
func GetLinks(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetLinks(traceID, req))
	}
}

// EditLinks 	godoc
// @id			EditLinks 友情链接编辑
// @Summary		友情链接编辑
// @Description	基于数据ID编辑友情链接
// @Router		/blog/links/edit [post]
// @Tags		友情链接
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.LinksEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditLinks(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.LinksEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.LinksEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditLinks(traceID, req))
	}
}

// DelLinks	godoc
// @id			DelLinks 友情链接移除
// @Summary		友情链接移除
// @Description	友情链接移除处理
// @Router		/blog/links/del [post]
// @Tags		友情链接
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelLinks(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		service.JsonRes(c, blogService.DelLinks(traceID, req))
	}
}
