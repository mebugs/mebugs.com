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
	indexIds := getRunIds(traceID, 0, 1)
	// 处理首页数据
	indexData.Category = getCategoryByIdSort(traceID, indexIds[0])
	indexData.Tag = getTagByIdSort(traceID, indexIds[1])
	indexData.Topic = getTopicByIdSort(traceID, indexIds[2])
	posts := getPostByIdSort(traceID, indexIds[3], indexIds[4], indexIds[5], indexIds[6])
	indexData.PostsNew = posts[0]
	indexData.PostsView = posts[1]
	indexData.PostsGood = posts[2]
	indexData.PostsHot = posts[3]
	return baseModel.Success(constant.Success, indexData)
}
