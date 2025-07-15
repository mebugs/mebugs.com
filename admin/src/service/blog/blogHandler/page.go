package blogHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/model/cacheModel"
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

// GetPages	godoc
// @id			GetPages 获取页面数据
// @Summary		获取页面数据
// @Description	获取页面数据
// @Router		/page/pages [get]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func GetPages(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	service.JsonRes(c, blogService.GetPage(traceID))
}

// GetPage	godoc
// @id			GetPage 获取页面数据
// @Summary		获取页面数据
// @Description	获取页面数据
// @Router		/page/page [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func GetPage(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostReq)
		// 执行创建
		service.JsonRes(c, blogService.GetPageDetail(traceID, req))
	}
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

// GetCategoryList 	godoc
// @id			GetCategoryList 获取分类数据
// @Summary		获取分类数据
// @Description	获取分类数据
// @Router		/page/categoryList [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func GetCategoryList(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	service.JsonRes(c, blogService.GetCategoryList(traceID))
}

// GetTags 	godoc
// @id			GetTags 获取标签数据
// @Summary		获取标签数据
// @Description	获取标签数据
// @Router		/page/tags [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func GetTags(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostsReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostsReq)
		// 执行创建
		service.JsonRes(c, blogService.GetTags(traceID, req))
	}
}

// GetPostDetail 	godoc
// @id			GetPostDetail 获取文章详情
// @Summary		获取文章详情
// @Description	获取文章详情
// @Router		/page/post [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func GetPostDetail(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostReq)
		// 执行创建
		service.JsonRes(c, blogService.GetPostDetail(traceID, req))
	}
}

// SetPostGood 	godoc
// @id			SetPostGood 文章深度追加
// @Summary		文章深度追加
// @Description	文章深度追加
// @Router		/page/post/good [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func SetPostGood(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostReq)
		// 执行创建
		service.JsonRes(c, blogService.SetPostGood(traceID, req))
	}
}

// AddComm 	godoc
// @id			AddComm 添加文章评论
// @Summary		文章深度追加
// @Description	文章深度追加
// @Router		/page/comm/add [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func AddComm(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.PostCommReq{})
	if err == nil {
		req := reqObj.(*blogModel.PostCommReq)
		// 执行创建
		service.JsonRes(c, blogService.AddComm(traceID, req))
	}
}

// Comments 	godoc
// @id			AddComm 查询评论
// @Summary		查询评论
// @Description	查询评论
// @Router		/page/comments [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func Comments(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.CommentsReq{})
	if err == nil {
		req := reqObj.(*blogModel.CommentsReq)
		// 执行创建
		service.JsonRes(c, blogService.Comments(traceID, req))
	}
}

// PageLink 	godoc
// @id			PageLink 获取友链
// @Summary		获取友链
// @Description	获取友链
// @Router		/page/link [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func PageLink(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	service.JsonRes(c, blogService.PageLink(traceID))
}

// LinkScan 	godoc
// @id			LinkScan 链接扫描
// @Summary		链接扫描
// @Description	链接扫描
// @Router		/page/link/scan [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func LinkScan(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.LinksScanReq{})
	if err == nil {
		req := reqObj.(*blogModel.LinksScanReq)
		// 执行创建
		service.JsonRes(c, blogService.LinkScan(traceID, req))
	}
}

// LinkAdd 	godoc
// @id			LinkAdd 链接提交
// @Summary		链接提交
// @Description	链接提交
// @Router		/page/link/add [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func LinkAdd(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.LinksAddClientReq{})
	if err == nil {
		req := reqObj.(*blogModel.LinksAddClientReq)
		// 执行创建
		service.JsonRes(c, blogService.LinkAdd(traceID, req))
	}
}

// SiteMap 	godoc
// @id			SiteMap 索引
// @Summary		索引
// @Description	索引
// @Router		/page/sitemap.xml [post]
// @Tags		Page
// @Accept		json
// @Produce		json
// @Security	Token
func SiteMap(c *gin.Context) {
	traceID := c.GetString(constant.ContextTraceID)
	xmlString := cacheModel.GetSiteMapCache(traceID)
	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}
