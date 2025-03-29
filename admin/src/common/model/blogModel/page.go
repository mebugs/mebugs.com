package blogModel

import "siteol.com/smart/src/common/model/cacheModel"

// IndexData 首页数据对象
type IndexData struct {
	Banners []*cacheModel.BannerCache `json:"banners"` // BANNER
	TagData
	CategoryData
	TopicsData
	PostsHot  []*cacheModel.PostCache `json:"postsHot"`
	PostsView []*cacheModel.PostCache `json:"postsView"`
	PostsGood []*cacheModel.PostCache `json:"postsGood"`
	PostsNew  []*cacheModel.PostCache `json:"postsNew"`
}

// TopicsData 专栏数据对象
type TopicsData struct {
	Topic []*cacheModel.TopicCache `json:"topic"`
}

// CategoryData 分类数据对象
type CategoryData struct {
	Category []*cacheModel.CategoryCache `json:"category"`
}

// TagData 标签数据对象
type TagData struct {
	Tag []*cacheModel.TagCache `json:"tag"`
}

// TagsData 标签列表页响应
type TagsData struct {
	TagData
	Total int `json:"total"`
}

// PostsReq 查询文章的对象
type PostsReq struct {
	By    int    `json:"by"`    // 查询源：1 new 2 good 3 view 4 hot 5 all=new 6 category 7 tag 8 topic
	Value string `json:"value"` // 查询源：category tag topic 的值
	Page  int    `json:"page"`  // 查询页，每页固定15个
}

// PostsData 文章列表页响应
type PostsData struct {
	Null  bool                    `json:"null"`
	Main  any                     `json:"main"`
	Posts []*cacheModel.PostCache `json:"posts"`
	Total int                     `json:"total"`
}

// PostReq 查询文章的对象
type PostReq struct {
	Url string `json:"url"` // 文章URL
}

// PostData 文章响应
type PostData struct {
	Null     bool                      `json:"null"`
	Post     *cacheModel.PostCache     `json:"post"`
	PostMain *cacheModel.PostMainCache `json:"postMain"`
	PostMore []*cacheModel.PostCache   `json:"postMore"`
}
