package cacheModel

import (
	"encoding/xml"
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/redis"
	"time"
)

// 初始化默认页面
func makeFirstUrls() []*URL {
	urls := make([]*URL, 0)
	today := time.Now().Format("2006-01-02")
	// 首页
	urls = append(urls, &URL{
		Loc:        constant.DomainUrl,
		LastMod:    today,
		ChangeFreq: "hourly",
		Priority:   1,
	})
	// 索引页
	urls = append(urls, &URL{
		Loc:        fmt.Sprintf(constant.DomainGroupUrl, constant.DomainUrl, "view"),
		LastMod:    today,
		ChangeFreq: "hourly",
		Priority:   0.9,
	})
	urls = append(urls, &URL{
		Loc:        fmt.Sprintf(constant.DomainGroupUrl, constant.DomainUrl, "hot"),
		LastMod:    today,
		ChangeFreq: "hourly",
		Priority:   0.9,
	})
	urls = append(urls, &URL{
		Loc:        fmt.Sprintf(constant.DomainGroupUrl, constant.DomainUrl, "good"),
		LastMod:    today,
		ChangeFreq: "hourly",
		Priority:   0.9,
	})
	urls = append(urls, &URL{
		Loc:        fmt.Sprintf(constant.DomainGroupUrl, constant.DomainUrl, "new"),
		LastMod:    today,
		ChangeFreq: "hourly",
		Priority:   0.9,
	})
	// 分类和标签
	urls = append(urls, &URL{
		Loc:        fmt.Sprintf(constant.DomainCategoryUrl, constant.DomainUrl),
		LastMod:    today,
		ChangeFreq: "daily",
		Priority:   0.8,
	})
	urls = append(urls, &URL{
		Loc:        fmt.Sprintf(constant.DomainTagUrl, constant.DomainUrl),
		LastMod:    today,
		ChangeFreq: "daily",
		Priority:   0.8,
	})
	return urls
}

// setPostUrls 追加文章URL
func setPostUrls(postUrl []*URL, doDay, path string) []*URL {
	// 文章
	postUrl = append(postUrl, &URL{
		Loc:        fmt.Sprintf(constant.DomainPostUrl, constant.DomainUrl, path),
		LastMod:    doDay,
		ChangeFreq: "weekly",
		Priority:   0.7,
	})
	return postUrl
}

// setPostUrls 追加文章URL
func setCateTagUrls(url []*URL, cateAndTagHaveMap map[string]bool, doDay, pre, path string, id uint64) []*URL {
	// 存在了
	_, ok := cateAndTagHaveMap[fmt.Sprintf("%s_%d", pre, id)]
	if ok {
		return url
	}
	// 文章
	url = append(url, &URL{
		Loc:        fmt.Sprintf(constant.DomainCateTagDetailUrl, constant.DomainUrl, pre, path),
		LastMod:    doDay,
		ChangeFreq: "weekly",
		Priority:   0.8,
	})
	// 修复重复的错误
	cateAndTagHaveMap[fmt.Sprintf("%s_%d", pre, id)] = true
	return url
}

// 初始化其他页面
func makeLastUrls(pageCache []*BannerCache) []*URL {
	urls := make([]*URL, 0)
	today := time.Now().Format("2006-01-02")
	urls = append(urls, &URL{
		Loc:        fmt.Sprintf(constant.DomainPageUrl, constant.DomainUrl),
		LastMod:    today,
		ChangeFreq: "weekly",
		Priority:   0.7,
	})
	for _, page := range pageCache {
		urls = append(urls, &URL{
			Loc:        fmt.Sprintf(constant.DomainPageDetailUrl, constant.DomainUrl, page.Url),
			LastMod:    today,
			ChangeFreq: "weekly",
			Priority:   0.7,
		})
	}
	return urls
}

func makeSiteMapCache(traceID string, firstUrl, categoryUrl, tagUrl, postUrl, lastUrl []*URL) {
	urlSet := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  mergeURLS(firstUrl, categoryUrl, tagUrl, postUrl, lastUrl),
	}
	xmlData, _ := xml.MarshalIndent(urlSet, "", "")
	xmlString := fmt.Sprintf("%s%s", xml.Header, string(xmlData))
	err := redis.Set(constant.SiteMapCache, xmlString, 0)
	if err != nil {
		log.ErrorTF(traceID, "MakeSiteMapCache Fail . Err Is : %v", err)
	}
}

func mergeURLS(firstUrl, categoryUrl, tagUrl, postUrl, lastUrl []*URL) []*URL {
	totalLen := len(firstUrl) + len(categoryUrl) + len(tagUrl) + len(postUrl) + len(lastUrl)
	result := make([]*URL, totalLen)
	pos := 0
	pos += copy(result[pos:], firstUrl)
	pos += copy(result[pos:], categoryUrl)
	pos += copy(result[pos:], tagUrl)
	pos += copy(result[pos:], postUrl)
	copy(result[pos:], lastUrl)
	return result
}
