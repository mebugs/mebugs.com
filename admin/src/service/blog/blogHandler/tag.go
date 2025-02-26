package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddTag 	godoc
// @id			AddTag 标签新建
// @Summary		标签新建
// @Description	新建标签
// @Router		/blog/tag/add [post]
// @Tags		标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.TagAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.TagAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.TagAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddTag(traceID, req))
	}
}

// PageTag	godoc
// @id			PageTag 标签分页
// @Summary		标签分页
// @Description	分页查询标签
// @Router		/blog/tag/page [post]
// @Tags		标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.TagPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.TagPageRes}}	"响应成功"
func PageTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.TagPageReq{})
	if err == nil {
		req := reqObj.(*blogModel.TagPageReq)
		// 执行查询
		service.JsonRes(c, blogService.PageTag(traceID, req))
	}
}

// GetTag	godoc
// @id			GetTag 标签详情
// @Summary		标签详情
// @Description	查询标签详情
// @Router		/blog/tag/get [post]
// @Tags		标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.TagGetRes}	"响应成功"
func GetTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetTag(traceID, req))
	}
}

// EditTag 	godoc
// @id			EditTag 标签编辑
// @Summary		标签编辑
// @Description	基于数据ID编辑标签
// @Router		/blog/tag/edit [post]
// @Tags		标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.TagEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.TagEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.TagEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditTag(traceID, req))
	}
}

// DelTag	godoc
// @id			DelTag 标签移除
// @Summary		标签移除
// @Description	标签移除处理
// @Router		/blog/tag/del [post]
// @Tags		标签
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelTag(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)

		service.JsonRes(c, blogService.DelTag(traceID, req))
	}
}

// ListTag	godoc
// @id			ListTag 标签列表
// @Summary		标签列表
// @Description	标签列表处理
// @Router		/blog/tag/list [post]
// @Tags		文章分类
// @Accept		json
// @Produce		json
// @Security	Token
// @Success		200		{object}	baseModel.ResBody{data=[]baseModel.SelectRes}	"响应成功"
func ListTag(c *gin.Context) {
	traceID := c.GetString(constant.ContextTraceID)
	service.JsonRes(c, blogService.ListTag(traceID))
}
