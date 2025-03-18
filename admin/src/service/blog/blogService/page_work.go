package blogService

import (
	"siteol.com/smart/src/common/model/cacheModel"
)

// getRunUrls 获得需要处理的URL列表 0：Index
func getRunUrls(traceID string, runBy, page int) (backIds [][]string) {
	// 获取ID列表
	idList := cacheModel.GetPostSortCache(traceID)
	if idList == nil {
		return
	}
	switch runBy {
	case 0: // Index  0 category 1 tag 2 topic 3 new 4 view 5 good 6 hot
		backIds = [][]string{
			idList.Category,
			getUrlsByPage(idList.Tag, page, 20),
			getUrlsByPage(idList.Topic, page, 4),
			getUrlsByPage(idList.PostNews, page, 15),
			getUrlsByPage(idList.PostViews, page, 6),
			getUrlsByPage(idList.PostGoods, page, 6),
			getUrlsByPage(idList.PostHots, page, 6),
		}

	}
	return
}

// getUrlsByPage 根据分页获得ID
func getUrlsByPage(urls []string, page, size int) []string {
	length := len(urls)
	start := (page - 1) * size
	end := page * size
	if length < start {
		return []string{}
	} else {
		if length <= end {
			return urls[start:]
		} else {
			return urls[start:end]
		}
	}
}

// getCategoryByUrlSort 获取分类顺序数据
func getCategoryByUrlSort(traceID string, urls []string) (res []*cacheModel.CategoryCache) {
	cache := cacheModel.GetCategoryCache(traceID)
	if cache == nil {
		return
	}
	res = make([]*cacheModel.CategoryCache, len(urls))
	for i, url := range urls {
		res[i] = cache[url]
	}
	return
}

// getTagByUrlSort 获取分类顺序数据
func getTagByUrlSort(traceID string, urls []string) (res []*cacheModel.TagCache) {
	cache := cacheModel.GetTagCache(traceID)
	if cache == nil {
		return
	}
	res = make([]*cacheModel.TagCache, len(urls))
	for i, url := range urls {
		res[i] = cache[url]
	}
	return
}

// getTopicByUrlSort 获取分类顺序数据
func getTopicByUrlSort(traceID string, urls []string) (res []*cacheModel.TopicCache) {
	cache := cacheModel.GetTopicCache(traceID)
	if cache == nil {
		return
	}
	res = make([]*cacheModel.TopicCache, len(urls))
	for i, url := range urls {
		res[i] = cache[url]
	}
	return
}

// getPostByUrlSort 获取分类顺序数据 new  view  good  hot
func getPostByUrlSort(traceID string, new, view, good, hot []string) (res [][]*cacheModel.PostCache) {
	cache := cacheModel.GetPostCache(traceID)
	if cache == nil {
		return
	}
	res = [][]*cacheModel.PostCache{getPostByIds(new, cache), getPostByIds(view, cache), getPostByIds(good, cache), getPostByIds(hot, cache)}
	return
}

func getPostByIds(urls []string, cache map[string]*cacheModel.PostCache) (res []*cacheModel.PostCache) {
	res = make([]*cacheModel.PostCache, len(urls))
	for i, url := range urls {
		res[i] = cache[url]
	}
	return
}
