package routers

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/router/middleware"
	"siteol.com/smart/src/service/blog/blogHandler"
)

// BlogRouter 博客路由
func BlogRouter(router *gin.Engine) {
	blogRouter := router.Group("/blog", middleware.CommMiddleWare) // 授权中间件
	{
		// 首页统计
		indexRouter := blogRouter.Group("/center")
		{
			indexRouter.POST("/index", blogHandler.CenterIndex)
		}
		// 资源配置表相关
		sourceRouter := blogRouter.Group("/source")
		{
			sourceRouter.POST("/add", blogHandler.AddSource)
			sourceRouter.POST("/page", blogHandler.PageSource)
			sourceRouter.POST("/get", blogHandler.GetSource)
			sourceRouter.POST("/edit", blogHandler.EditSource)
			sourceRouter.POST("/del", blogHandler.DelSource)
		}

		// Banner
		bannerRouter := blogRouter.Group("/banner")
		{
			bannerRouter.POST("/add", blogHandler.AddBanner)
			bannerRouter.POST("/page", blogHandler.PageBanner)
			bannerRouter.POST("/get", blogHandler.GetBanner)
			bannerRouter.POST("/edit", blogHandler.EditBanner)
			bannerRouter.POST("/del", blogHandler.DelBanner)
		}

		// 文章分类相关
		categoryRouter := blogRouter.Group("/category")
		{
			categoryRouter.POST("/add", blogHandler.AddCategory)
			categoryRouter.POST("/page", blogHandler.PageCategory)
			categoryRouter.POST("/get", blogHandler.GetCategory)
			categoryRouter.POST("/edit", blogHandler.EditCategory)
			categoryRouter.POST("/del", blogHandler.DelCategory)
			categoryRouter.POST("/merge", blogHandler.MergeCategory)
			categoryRouter.POST("/list", blogHandler.ListCategory)
		}

		// 标签相关
		tagRouter := blogRouter.Group("/tag")
		{
			tagRouter.POST("/add", blogHandler.AddTag)
			tagRouter.POST("/page", blogHandler.PageTag)
			tagRouter.POST("/get", blogHandler.GetTag)
			tagRouter.POST("/edit", blogHandler.EditTag)
			tagRouter.POST("/del", blogHandler.DelTag)
			tagRouter.POST("/list", blogHandler.ListTag)
		}

		// 文章专题相关
		topicRouter := blogRouter.Group("/topic")
		{
			topicRouter.POST("/add", blogHandler.AddTopic)
			topicRouter.POST("/page", blogHandler.PageTopic)
			topicRouter.POST("/get", blogHandler.GetTopic)
			topicRouter.POST("/edit", blogHandler.EditTopic)
			topicRouter.POST("/del", blogHandler.DelTopic)
		}

		// 文章或页面相关
		postRouter := blogRouter.Group("/post")
		{
			postRouter.POST("/add", blogHandler.AddPost)
			postRouter.POST("/page", blogHandler.PagePost)
			postRouter.POST("/get", blogHandler.GetPost)
			postRouter.POST("/edit", blogHandler.EditPost)
		}

		// 文章评论相关
		postCommentRouter := blogRouter.Group("/postComment")
		{
			postCommentRouter.POST("/page", blogHandler.PagePostComment)
			postCommentRouter.POST("/get", blogHandler.GetPostComment)
			postCommentRouter.POST("/edit", blogHandler.EditPostComment)
		}

		// 友情链接相关
		linksRouter := blogRouter.Group("/links")
		{
			linksRouter.POST("/add", blogHandler.AddLinks)
			linksRouter.POST("/page", blogHandler.PageLinks)
			linksRouter.POST("/get", blogHandler.GetLinks)
			linksRouter.POST("/edit", blogHandler.EditLinks)
			linksRouter.POST("/del", blogHandler.DelLinks)
		}
	}

	pageRouter := router.Group("/page", middleware.CommMiddleWare) // 授权中间件
	{
		pageRouter.GET("/index", blogHandler.GetIndex)               // 获取首页数据
		pageRouter.GET("/pages", blogHandler.GetPages)               // 获取页面
		pageRouter.POST("/page", blogHandler.GetPage)                // 获取页面
		pageRouter.POST("/posts", blogHandler.GetPosts)              // 获取文章列表数据
		pageRouter.GET("/categoryList", blogHandler.GetCategoryList) // 获取分类列表数据
		pageRouter.POST("/tags", blogHandler.GetTags)                // 获取文章列表数据
		pageRouter.POST("/post", blogHandler.GetPostDetail)          // 获取文章详情
		pageRouter.POST("/post/good", blogHandler.SetPostGood)       // Push文章Goods
		pageRouter.POST("/comments", blogHandler.Comments)           // 获取评论
		pageRouter.POST("/comm/add", blogHandler.AddComm)            // 提交评论
		pageRouter.POST("/link", blogHandler.PageLink)               // 获取友链
		pageRouter.POST("/link/scan", blogHandler.LinkScan)          // 链接扫描
		pageRouter.POST("/link/add", blogHandler.LinkAdd)            // 链接添加

	}
	openRouter := router.Group("/open")
	{
		openRouter.GET("/sitemap.xml", blogHandler.SiteMap) // SITEMAP
	}
}
