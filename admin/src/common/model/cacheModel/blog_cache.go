package cacheModel

import (
	"encoding/json"
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/redis"
	"strconv"
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

// GetPagesCache 获取Pages配置
func GetPagesCache(traceID string) (res []*BannerCache) {
	// 读取缓存
	str, err := redis.Get(constant.PagePageCache)
	if err != nil {
		log.WarnTF(traceID, "GetPagesCache Fail . Err Is : %v", err)
		return

	}
	res = make([]*BannerCache, 0)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetPagesCache Fail . Err Is : %v", err)
	}
	return
}

// GetLinksCache 获取Links
func GetLinksCache(traceID string) (res []*LinksCache) {
	// 读取缓存
	str, err := redis.Get(constant.PageLinksCache)
	if err != nil {
		log.WarnTF(traceID, "GetLinksCache Fail . Err Is : %v", err)
		return

	}
	res = make([]*LinksCache, 0)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetLinksCache Fail . Err Is : %v", err)
	}
	return
}

// GetCategoryCache 获取Category缓存
func GetCategoryCache(traceID string) (res map[string]*CategoryCache) {
	// 读取缓存
	str, err := redis.Get(constant.PageCategoryCache)
	if err != nil {
		log.WarnTF(traceID, "GetCategoryCache Fail . Err Is : %v", err)
		return

	}
	res = make(map[string]*CategoryCache)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetCategoryCache Fail . Err Is : %v", err)
	}
	return
}

// GetTagCache 获取Tag配置
func GetTagCache(traceID string) (res map[string]*TagCache) {
	// 读取缓存
	str, err := redis.Get(constant.PageTagCache)
	if err != nil {
		log.WarnTF(traceID, "GetTagCache Fail . Err Is : %v", err)
		return

	}
	res = make(map[string]*TagCache)
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetTagCache Fail . Err Is : %v", err)
	}
	return
}

// GetPostCache 获取Post缓存
func GetPostCache(traceID string) (res map[string]*PostCache) {
	// 读取缓存
	str, err := redis.Get(constant.PagePostCache)
	if err != nil {
		log.WarnTF(traceID, "GetPostCache Fail . Err Is : %v", err)
		return

	}
	res = make(map[string]*PostCache)
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

// GetPostMainCache 获取文章换算
func GetPostMainCache(traceID, url string) (res *PostMainCache) {
	// 读取缓存
	str, err := redis.Get(fmt.Sprintf(constant.PagePostCache, url))
	if err != nil {
		log.WarnTF(traceID, "GetPostCache Fail . Err Is : %v", err)
		return
	}
	res = &PostMainCache{}
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetPostCache Fail . Err Is : %v", err)
	}
	return
}

// GetPostCommUserCache 获取用户当日评论缓存
func GetPostCommUserCache(traceID string, uid uint64) (res map[uint64]uint64) {
	res = make(map[uint64]uint64)
	// 读取缓存
	str, err := redis.Get(fmt.Sprintf(constant.PagePostCommUserCache, uid))
	if err != nil {
		log.WarnTF(traceID, "GetPostCommUserCache Fail . Err Is : %v", err)
		return
	}
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal GetPostCommUserCache Fail . Err Is : %v", err)
	}
	return
}

// GetClientLinkCache 获取客户端提交的链接上线
func GetClientLinkCache(traceID, clientId string) (res int64) {
	// 读取缓存
	str, err := redis.Get(fmt.Sprintf(constant.PageClientLinkCache, clientId))
	if err != nil {
		log.WarnTF(traceID, "GetPostCommUserCache Fail . Err Is : %v", err)
		return
	}
	res, _ = strconv.ParseInt(str, 10, 64)
	return
}

// GetSiteMapCache 获取地图
func GetSiteMapCache(traceID string) (res string) {
	// 读取缓存
	res, err := redis.Get(constant.SiteMapCache)
	if err != nil {
		log.WarnTF(traceID, "GetSiteMapCache Fail . Err Is : %v", err)
		return
	}
	return
}
