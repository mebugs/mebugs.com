package blogService

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/cacheModel"
	"strings"
)

// getRunUrls 获得需要处理的URL列表 0：Index 1:
func getRunUrls(traceID string, runBy, page int) (backIds [][]string, total int) {
	// 获取ID列表
	idList := cacheModel.GetPostSortCache(traceID)
	if idList == nil {
		return
	}
	total = len(idList.PostNews)
	switch runBy {
	case 0: // 0 Index  1 new 2 good 3 view 4 hot 5 all=new 6 category 7 tag 8 topic
		backIds = [][]string{
			idList.Category,
			getUrlsByPage(idList.Tag, page, 20),
			getUrlsByPage(idList.PostNews, page, 18),
			getUrlsByPage(idList.PostViews, page, 6),
			getUrlsByPage(idList.PostGoods, page, 6),
			getUrlsByPage(idList.PostHots, page, 6),
		}
	case 1: // 1 new
		backIds = [][]string{getUrlsByPage(idList.PostNews, page, 18)}
	case 2: // 2 good
		backIds = [][]string{getUrlsByPage(idList.PostGoods, page, 18)}
	case 3: // 3 view
		backIds = [][]string{getUrlsByPage(idList.PostViews, page, 18)}
	case 4: // 4 hot
		backIds = [][]string{getUrlsByPage(idList.PostHots, page, 18)}
	case 6: // 6 category
		backIds = [][]string{idList.Category}
	case 7: // 7 tag
		backIds = [][]string{getUrlsByPage(idList.Tag, page, 24)}
		total = len(idList.Tag)
	default: // 5 || Other
		backIds = [][]string{idList.PostNews}
	}
	return
}

// getGroupRunUrls 获取相关分组下的文章数据
func getGroupRunUrls(traceID, url string, runBy, page int) (backIds [][]string, total int, group any) {
	switch runBy {
	case 6: //  6 category 7 tag 8 topic
		resMap := cacheModel.GetCategoryCache(traceID)
		if res, ok := resMap[url]; ok {
			total = len(res.Posts)
			backIds = [][]string{getUrlsByPage(res.Posts, page, 18)}
			res.Posts = nil
			group = res
		}
	case 7: //  6 category 7 tag 8 topic
		resMap := cacheModel.GetTagCache(traceID)
		if res, ok := resMap[url]; ok {
			total = len(res.Posts)
			backIds = [][]string{getUrlsByPage(res.Posts, page, 18)}
			res.Posts = nil
			group = res
		}
	}
	return
}

// getUrlsByPage 根据分页获得ID
func getUrlsByPage(urls []string, page, size int) []string {
	if page < 1 {
		page = 1
	}
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

// getPostByUrlSort 获取分类顺序数据 new  view  good  hot
func getPostByUrlSort(traceID string, new, view, good, hot []string) (res [][]*cacheModel.PostCache) {
	cache := cacheModel.GetPostCache(traceID)
	if cache == nil {
		return
	}
	res = [][]*cacheModel.PostCache{getPostByUrls(new, cache), getPostByUrls(view, cache), getPostByUrls(good, cache), getPostByUrls(hot, cache)}
	return
}

func getPostByUrlSortSingle(traceID string, urls []string) (res []*cacheModel.PostCache) {
	cache := cacheModel.GetPostCache(traceID)
	if cache == nil {
		return
	}
	return getPostByUrls(urls, cache)
}

func getPostByUrls(urls []string, cache map[string]*cacheModel.PostCache) (res []*cacheModel.PostCache) {
	res = make([]*cacheModel.PostCache, len(urls))
	for i, url := range urls {
		res[i] = cache[url]
	}
	return
}

func getPostByUrl(traceID, url string) (res *cacheModel.PostCache, cache map[string]*cacheModel.PostCache) {
	cache = cacheModel.GetPostCache(traceID)
	if cache == nil {
		return
	}
	res = cache[url]
	return
}

func getAttr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func getImgBase64(traceID, url, icon string) string {
	// 处理图片URL
	if !strings.HasPrefix(icon, "http") {
		if !strings.HasPrefix(icon, "/") {
			icon = url + "/" + icon
		} else {
			icon = url + icon
		}
	}
	// 1. 发起 HTTP GET 请求获取图片
	resp, err := http.Get(icon)
	if err != nil {
		log.ErrorTF(traceID, "getImgBase64 Get fail . Err is %v", err)
		return ""
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 2. 读取图片二进制数据
	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.ErrorTF(traceID, "getImgBase64 Read fail . Err is %v", err)
		return ""
	}
	// 3. 转换为 Base64 字符串
	base64Str := base64.StdEncoding.EncodeToString(imageBytes)
	mimeType := http.DetectContentType(imageBytes)
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Str)

	return dataURL
}
