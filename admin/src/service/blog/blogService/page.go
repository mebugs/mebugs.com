package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/model/cacheModel"
)

// GetIndex 获取首页数据
func GetIndex(traceID string) *baseModel.ResBody {
	indexData := &blogModel.IndexData{
		Banners: cacheModel.GetBannerCache(traceID),
	}
	// 读取ID对象，并处理结果
	indexIds := getRunUrls(traceID, 0, 1)
	// 处理首页数据
	indexData.Category = getCategoryByUrlSort(traceID, indexIds[0])
	indexData.Tag = getTagByUrlSort(traceID, indexIds[1])
	indexData.Topic = getTopicByUrlSort(traceID, indexIds[2])
	posts := getPostByUrlSort(traceID, indexIds[3], indexIds[4], indexIds[5], indexIds[6])
	indexData.PostsNew = posts[0]
	indexData.PostsView = posts[1]
	indexData.PostsGood = posts[2]
	indexData.PostsHot = posts[3]
	return baseModel.Success(constant.Success, indexData)
}

// GetPosts 获取文章数据
func GetPosts(traceID string, req *blogModel.PostsReq) *baseModel.ResBody {
	return baseModel.Success(constant.Success, nil)
}
