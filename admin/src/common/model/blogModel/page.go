package blogModel

import "siteol.com/smart/src/common/model/cacheModel"

// IndexData 首页数据对象
type IndexData struct {
	Banners   []*cacheModel.BannerCache   `json:"banners"` // BANNER
	Category  []*cacheModel.CategoryCache `json:"category"`
	Tag       []*cacheModel.TagCache      `json:"tag"`
	Topic     []*cacheModel.TopicCache    `json:"topic"`
	PostsHot  []*cacheModel.PostCache     `json:"postsHot"`
	PostsView []*cacheModel.PostCache     `json:"postsView"`
	PostsGood []*cacheModel.PostCache     `json:"postsGood"`
	PostsNew  []*cacheModel.PostCache     `json:"postsNew"`
}

// PostsReq 查询文章的对象
type PostsReq struct {
	By    string `json:"by"`    // 查询源：new good view hot category tag topic all=new
	Value string `json:"value"` // 查询源：category tag topic
	Page  int    `json:"page"`  // 查询页，每页固定15个
}
