package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/redis"
)

// GetIndex 获取首页数据
func GetIndex(traceID string) *baseModel.ResBody {
	indexData := &blogModel.IndexData{
		Banners: cacheModel.GetBannerCache(traceID),
	}
	// 读取ID对象，并处理结果
	indexIds, _ := getRunUrls(traceID, 0, 1)
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
	var urls [][]string
	var total int
	var main any
	// 0 Index  1 new 2 good 3 view 4 hot 5 all=new 6 category 7 tag 8 topic
	if req.By <= 5 {
		urls, total = getRunUrls(traceID, req.By, req.Page)
	} else {
		// 6 category 7 tag 8 topic
		urls, total, main = getGroupRunUrls(traceID, req.Value, req.By, req.Page)
	}
	if total == 0 {
		return baseModel.Success(constant.Success, &blogModel.PostsData{Null: true})
	}
	posts := getPostByUrlSortSingle(traceID, urls[0])
	return baseModel.Success(constant.Success, &blogModel.PostsData{
		Main:  main,
		Posts: posts,
		Total: total,
	})
}

// GetTopics 获取专栏数据
func GetTopics(traceID string) *baseModel.ResBody {
	urls, _ := getRunUrls(traceID, 8, 1)
	topicsMap := cacheModel.GetTopicCache(traceID)
	topics := make([]*cacheModel.TopicCache, len(urls[0]))
	for i, url := range urls[0] {
		topics[i] = topicsMap[url]
	}
	return baseModel.Success(constant.Success, &blogModel.TopicsData{Topic: topics})
}

// GetCategoryList 获取专栏数据
func GetCategoryList(traceID string) *baseModel.ResBody {
	urls, _ := getRunUrls(traceID, 6, 1)
	categoryMap := cacheModel.GetCategoryCache(traceID)
	category := make([]*cacheModel.CategoryCache, len(urls[0]))
	for i, url := range urls[0] {
		category[i] = categoryMap[url]
	}
	return baseModel.Success(constant.Success, &blogModel.CategoryData{Category: category})
}

// GetTags 获取标签数据
func GetTags(traceID string, req *blogModel.PostsReq) *baseModel.ResBody {
	urls, total := getRunUrls(traceID, 7, req.Page)
	tagMap := cacheModel.GetTagCache(traceID)
	tag := make([]*cacheModel.TagCache, len(urls[0]))
	for i, url := range urls[0] {
		tag[i] = tagMap[url]
	}
	return baseModel.Success(constant.Success, &blogModel.TagsData{
		TagData: blogModel.TagData{Tag: tag},
		Total:   total,
	})
}

// GetPostDetail 获取标签数据
func GetPostDetail(traceID string, req *blogModel.PostReq) *baseModel.ResBody {
	postData := &blogModel.PostData{}
	// 先尝试获得缓存
	postBase, allMap := getPostByUrl(traceID, req.Url)
	if postBase == nil {
		postData.Null = true
		return baseModel.Success(constant.Success, postData)
	}
	postData.Post = postBase
	postMain := cacheModel.GetPostMainCache(traceID, postBase.Url)
	// 缓存不存在，开始构建缓存，如果是不存在的文章，则构建一个null=true的缓存
	if postMain == nil {
		// 刷新文章正文信息
		postMain = cacheModel.SyncPostMain(traceID, postBase)
		// 刷新失败
		if postMain == nil {
			postData.Null = true
			return baseModel.Success(constant.Success, postData)
		}
	}
	postData.PostMain = postMain
	postMore := make([]*cacheModel.PostCache, len(postMain.Like))
	for i, like := range postMain.Like {
		postMore[i] = allMap[like]
	}
	postData.PostMore = postMore
	// 处理最新统计数据
	postBase.Views++
	postBase.Hots++
	// 写回缓存
	allMap[postBase.Url] = postBase
	_ = redis.Set(constant.PagePostCache, allMap, 0)
	return baseModel.Success(constant.Success, postData)
}

// SetPostGood 追加文章深度
func SetPostGood(traceID string, req *blogModel.PostReq) *baseModel.ResBody {
	postBase, allMap := getPostByUrl(traceID, req.Url)
	if postBase == nil {
		return baseModel.Success(constant.Success, nil)
	}
	// 处理最新统计数据
	postBase.Hots++
	postBase.Goods++
	// 写回缓存
	allMap[postBase.Url] = postBase
	_ = redis.Set(constant.PagePostCache, allMap, 0)
	return baseModel.Success(constant.Success, nil)
}
