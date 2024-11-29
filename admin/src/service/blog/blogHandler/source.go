package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddSource 	godoc
// @id			AddSource 提交资源
// @Summary		提交资源
// @Description	提交资源
// @Router		/blog/source/add [post]
// @Tags		资源配置表
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.SourceAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddSource(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.SourceAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.SourceAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddSource(traceID, req))
	}
}

// PageSource	godoc
// @id			PageSource 资源配置表分页
// @Summary		资源配置表分页
// @Description	分页查询资源配置表
// @Router		/blog/source/page [post]
// @Tags		资源配置表
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.SourcePageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.SourcePageRes}}	"响应成功"
func PageSource(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.SourcePageReq{})
	if err == nil {
		req := reqObj.(*blogModel.SourcePageReq)
		// 执行查询
		service.JsonRes(c, blogService.PageSource(traceID, req))
	}
}

// GetSource	godoc
// @id			GetSource 资源配置表详情
// @Summary		资源配置表详情
// @Description	查询资源配置表详情
// @Router		/blog/source/get [post]
// @Tags		资源配置表
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.SourceGetRes}	"响应成功"
func GetSource(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetSource(traceID, req))
	}
}

// EditSource 	godoc
// @id			EditSource 资源配置表编辑
// @Summary		资源配置表编辑
// @Description	基于数据ID编辑资源配置表
// @Router		/blog/source/edit [post]
// @Tags		资源配置表
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.SourceEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditSource(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.SourceEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.SourceEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditSource(traceID, req))
	}
}

// DelSource	godoc
// @id			DelSource 资源配置表清理
// @Summary		资源配置表清理
// @Description	资源配置表清理处理
// @Router		/blog/source/del [post]
// @Tags		资源配置表
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelSource(c *gin.Context) {
	traceID := c.GetString(constant.ContextTraceID)
	service.JsonRes(c, blogService.DelSource(traceID))
}
