package cacheModel

import (
	"encoding/json"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/redis"
)

// GetBannerCache 获取Banner配置
func GetBannerCache(traceID string) (res []*BannerCache) {
	// 读取缓存
	str, err := redis.Get(constant.PageBannerCache)
	if err != nil {
		log.WarnTF(traceID, "GetBannerCache Fail . Err Is : %v", err)
		return

	}
	res = make([]*BannerCache, 0)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetBannerCache Fail . Err Is : %v", err)
	}
	return
}

// GetCategoryCache 获取Category缓存
func GetCategoryCache(traceID string) (res map[uint64]*CategoryCache) {
	// 读取缓存
	str, err := redis.Get(constant.PageCategoryCache)
	if err != nil {
		log.WarnTF(traceID, "GetCategoryCache Fail . Err Is : %v", err)
		return

	}
	res = make(map[uint64]*CategoryCache)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetCategoryCache Fail . Err Is : %v", err)
	}
	return
}

// GetTagCache 获取Tag配置
func GetTagCache(traceID string) (res map[uint64]*TagCache) {
	// 读取缓存
	str, err := redis.Get(constant.PageTagCache)
	if err != nil {
		log.WarnTF(traceID, "GetTagCache Fail . Err Is : %v", err)
		return

	}
	res = make(map[uint64]*TagCache)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetTagCache Fail . Err Is : %v", err)
	}
	return
}

// GetTopicCache 获取Topic缓存
func GetTopicCache(traceID string) (res map[uint64]*TopicCache) {
	// 读取缓存
	str, err := redis.Get(constant.PageTopicCache)
	if err != nil {
		log.WarnTF(traceID, "GetTopicCache Fail . Err Is : %v", err)
		return

	}
	res = make(map[uint64]*TopicCache)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetTopicCache Fail . Err Is : %v", err)
	}
	return
}

// GetPostCache 获取Post缓存
func GetPostCache(traceID string) (res map[uint64]*PostCache) {
	// 读取缓存
	str, err := redis.Get(constant.PagePostCache)
	if err != nil {
		log.WarnTF(traceID, "GetPostCache Fail . Err Is : %v", err)
		return

	}
	res = make(map[uint64]*PostCache)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetPostCache Fail . Err Is : %v", err)
	}
	return
}

// GetPostSortCache 获取PostSort缓存
func GetPostSortCache(traceID string) (res *BlogListCache) {
	// 读取缓存
	str, err := redis.Get(constant.PagePostSortCache)
	if err != nil {
		log.WarnTF(traceID, "GetPostSortCache Fail . Err Is : %v", err)
		return

	}
	res = &BlogListCache{}
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetPostSortCache Fail . Err Is : %v", err)
	}
	return
}
