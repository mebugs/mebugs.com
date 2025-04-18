package blogHandler

//
//import (
//	"github.com/gin-gonic/gin"
//	"siteol.com/smart/src/common/model/baseModel"
//	"siteol.com/smart/src/common/model/blogModel"
//	"siteol.com/smart/src/service"
//	"siteol.com/smart/src/service/blog/blogService"
//)
//
//// AddUser 	godoc
//// @id			AddUser 极简用户信息新建
//// @Summary		极简用户信息新建
//// @Description	新建极简用户信息
//// @Router		/blog/user/add [post]
//// @Tags		极简用户信息
//// @Accept		json
//// @Produce		json
//// @Security	Token
//// @Param		req	body		blogModel.UserAddReq	true	"请求"
//// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
//func AddUser(c *gin.Context) {
//	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.UserAddReq{})
//	if err == nil {
//		req := reqObj.(*blogModel.UserAddReq)
//		// 执行创建
//		service.JsonRes(c, blogService.AddUser(traceID, req))
//	}
//}
//
//// PageUser	godoc
//// @id			PageUser 极简用户信息分页
//// @Summary		极简用户信息分页
//// @Description	分页查询极简用户信息
//// @Router		/blog/user/page [post]
//// @Tags		极简用户信息
//// @Accept		json
//// @Produce		json
//// @Security	Token
//// @Param		req	body		blogModel.UserPageReq	true	"请求"
//// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.UserPageRes}}	"响应成功"
//func PageUser(c *gin.Context) {
//	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.UserPageReq{})
//	if err == nil {
//		req := reqObj.(*blogModel.UserPageReq)
//		// 执行查询
//		service.JsonRes(c, blogService.PageUser(traceID, req))
//	}
//}
//
//// GetUser	godoc
//// @id			GetUser 极简用户信息详情
//// @Summary		极简用户信息详情
//// @Description	查询极简用户信息详情
//// @Router		/blog/user/get [post]
//// @Tags		极简用户信息
//// @Accept		json
//// @Produce		json
//// @Security	Token
//// @Param		req	body		baseModel.IdReq	true	"请求"
//// @Success		200	{object}	baseModel.ResBody{data=blogModel.UserGetRes}	"响应成功"
//func GetUser(c *gin.Context) {
//	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
//	if err == nil {
//		req := reqObj.(*baseModel.IdReq)
//		// 执行查询
//		service.JsonRes(c, blogService.GetUser(traceID, req))
//	}
//}
//
//// EditUser 	godoc
//// @id			EditUser 极简用户信息编辑
//// @Summary		极简用户信息编辑
//// @Description	基于数据ID编辑极简用户信息
//// @Router		/blog/user/edit [post]
//// @Tags		极简用户信息
//// @Accept		json
//// @Produce		json
//// @Security	Token
//// @Param		req	body		blogModel.UserEditReq	true	"请求"
//// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
//func EditUser(c *gin.Context) {
//	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.UserEditReq{})
//	if err == nil {
//		req := reqObj.(*blogModel.UserEditReq)
//		// 执行编辑
//		service.JsonRes(c, blogService.EditUser(traceID, req))
//	}
//}
//
//// DelUser	godoc
//// @id			DelUser 极简用户信息移除
//// @Summary		极简用户信息移除
//// @Description	极简用户信息移除处理
//// @Router		/blog/user/del [post]
//// @Tags		极简用户信息
//// @Accept		json
//// @Produce		json
//// @Security	Token
//// @Param		req		body		baseModel.IdReq	true			"请求"
//// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
//func DelUser(c *gin.Context) {
//	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
//	if err == nil {
//		req := reqObj.(*baseModel.IdReq)
//
//		service.JsonRes(c, blogService.DelUser(traceID, req))
//	}
//}
