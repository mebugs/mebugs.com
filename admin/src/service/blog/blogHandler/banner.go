package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddBanner 	godoc
// @id			AddBanner Banner新建
// @Summary		Banner新建
// @Description	新建Banner
// @Router		/blog/banner/add [post]
// @Tags		Banner
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.BannerAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddBanner(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.BannerAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.BannerAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddBanner(traceID, req))
	}
}

// PageBanner	godoc
// @id			PageBanner Banner分页
// @Summary		Banner分页
// @Description	分页查询Banner
// @Router		/blog/banner/page [post]
// @Tags		Banner
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.BannerPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.BannerPageRes}}	"响应成功"
func PageBanner(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.BannerPageReq{})
	if err == nil {
		req := reqObj.(*blogModel.BannerPageReq)
		// 执行查询
		service.JsonRes(c, blogService.PageBanner(traceID, req))
	}
}

// GetBanner	godoc
// @id			GetBanner Banner详情
// @Summary		Banner详情
// @Description	查询Banner详情
// @Router		/blog/banner/get [post]
// @Tags		Banner
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.BannerGetRes}	"响应成功"
func GetBanner(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetBanner(traceID, req))
	}
}

// EditBanner 	godoc
// @id			EditBanner Banner编辑
// @Summary		Banner编辑
// @Description	基于数据ID编辑Banner
// @Router		/blog/banner/edit [post]
// @Tags		Banner
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.BannerEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditBanner(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.BannerEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.BannerEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditBanner(traceID, req))
	}
}

// DelBanner	godoc
// @id			DelBanner Banner移除
// @Summary		Banner移除
// @Description	Banner移除处理
// @Router		/blog/banner/del [post]
// @Tags		Banner
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelBanner(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		service.JsonRes(c, blogService.DelBanner(traceID, req))
	}
}
