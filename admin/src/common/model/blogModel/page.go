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
