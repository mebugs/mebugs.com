package blogService

import (
	"siteol.com/smart/src/common/model/cacheModel"
)

// getRunIds 获得需要处理的ID列表 0：Index
func getRunIds(traceID string, runBy, page int) (backIds [][]uint64) {
	// 获取ID列表
	idList := cacheModel.GetPostSortCache(traceID)
	if idList == nil {
		return
	}
	switch runBy {
	case 0: // Index  0 category 1 tag 2 topic 3 new 4 view 5 good 6 hot
		backIds = [][]uint64{
			idList.Category,
			getIdsByPage(idList.Tag, page, 20),
			getIdsByPage(idList.Topic, page, 4),
			getIdsByPage(idList.PostNews, page, 15),
			getIdsByPage(idList.PostViews, page, 6),
			getIdsByPage(idList.PostGoods, page, 6),
			getIdsByPage(idList.PostHots, page, 6),
		}

	}
	return
}

// getIdsByPage 根据分页获得ID
func getIdsByPage(ids []uint64, page, size int) []uint64 {
	length := len(ids)
	start := (page - 1) * size
	end := page * size
	if length < start {
		return []uint64{}
	} else {
		if length <= end {
			return ids[start:]
		} else {
			return ids[start:end]
		}
	}
}

// getCategoryByIdSort 获取分类顺序数据
func getCategoryByIdSort(traceID string, ids []uint64) (res []*cacheModel.CategoryCache) {
	cache := cacheModel.GetCategoryCache(traceID)
	if cache == nil {
		return
	}
	res = make([]*cacheModel.CategoryCache, len(ids))
	for i, id := range ids {
		res[i] = cache[id]
	}
	return
}

// getTagByIdSort 获取分类顺序数据
func getTagByIdSort(traceID string, ids []uint64) (res []*cacheModel.TagCache) {
	cache := cacheModel.GetTagCache(traceID)
	if cache == nil {
		return
	}
	res = make([]*cacheModel.TagCache, len(ids))
	for i, id := range ids {
		res[i] = cache[id]
	}
	return
}

// getTopicByIdSort 获取分类顺序数据
func getTopicByIdSort(traceID string, ids []uint64) (res []*cacheModel.TopicCache) {
	cache := cacheModel.GetTopicCache(traceID)
	if cache == nil {
		return
	}
	res = make([]*cacheModel.TopicCache, len(ids))
	for i, id := range ids {
		res[i] = cache[id]
	}
	return
}

// getPostByIdSort 获取分类顺序数据 new  view  good  hot
func getPostByIdSort(traceID string, new, view, good, hot []uint64) (res [][]*cacheModel.PostCache) {
	cache := cacheModel.GetPostCache(traceID)
	if cache == nil {
		return
	}
	res = [][]*cacheModel.PostCache{getPostByIds(new, cache), getPostByIds(view, cache), getPostByIds(good, cache), getPostByIds(hot, cache)}
	return
}

func getPostByIds(ids []uint64, cache map[uint64]*cacheModel.PostCache) (res []*cacheModel.PostCache) {
	res = make([]*cacheModel.PostCache, len(ids))
	for i, id := range ids {
		res[i] = cache[id]
	}
	return
}
