package cacheModel

import (
	"encoding/xml"
	"fmt"
	"math"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/blogDB"
	"siteol.com/smart/src/common/redis"
	"siteol.com/smart/src/common/utils"
	"sort"
	"time"
)

// BlogListCache 博客列表ID缓存
type BlogListCache struct {
	PostNews  []string `json:"postNews"`
	PostViews []string `json:"postViews"`
	PostGoods []string `json:"postGoods"`
	PostHots  []string `json:"postHots"`
	Category  []string `json:"category"`
	Tag       []string `json:"tag"`
}

// BannerCache 轮播缓存
type BannerCache struct {
	Title      string           `json:"title"`               // 名称
	Tag        string           `json:"tag"`                 // 标签
	Url        string           `json:"url"`                 // 分类地址
	Summary    string           `json:"summary"`             // 简介
	SourceShow string           `json:"sourceShow"`          // 资源图片地址
	Links      []*LinksCache    `json:"links,omitempty"`     // 友情链接
	LitePosts  []*PostLiteCache `json:"litePosts,omitempty"` // 全部文章
}

func ToLitePost(posts []*PostCache) (litePosts []*PostLiteCache) {
	litePosts = make([]*PostLiteCache, len(posts))
	for i, post := range posts {
		litePosts[i] = &PostLiteCache{
			Title: post.Title,
			Url:   post.Url,
		}
	}
	return
}

// PostLiteCache 轻量的文章
type PostLiteCache struct {
	Title string `json:"title"` // 标题
	Url   string `json:"url"`   // 文章地址
}

// PostCache 文章缓存（列表缓存，全文）
type PostCache struct {
	Id           uint64   `json:"id"`           // 数据ID
	Title        string   `json:"title"`        // 标题
	Url          string   `json:"url"`          // 文章地址
	Summary      string   `json:"summary"`      // 摘要
	SourcePath   string   `json:"sourcePath"`   // 图片路径
	SourceBack   string   `json:"sourceBack"`   // 图片后缀
	Category     string   `json:"category"`     // 分组URL
	CategoryName string   `json:"categoryName"` // 分组
	TagUrls      []string `json:"tagUrls"`      // 标签URL
	TagNames     []string `json:"tagNames"`     // 标签列表
	PushAt       string   `json:"pushAt"`       // 发布时间
	Views        uint64   `json:"views" `       // 总浏览量 ++
	Goods        uint64   `json:"goods"`        // 总支持量 60秒+1，页面不刷新最多+6次
	Hots         uint64   `json:"hots"`         // 近期热度 ++ Goods触发时++，每天对数据进行0.99取整进1处理
}

// PostMainCache 文章更新信息（触发时缓存12H）
type PostMainCache struct {
	Toc  string   `json:"toc"`  // 文章导航
	Html string   `json:"html"` // HTML源文件
	Like []string `json:"like"` // 相关文章
}

// LinksCacheLite 轻量版友链
type LinksCacheLite struct {
	Title string `json:"title" example:"demo"` // 名称
	Url   string `json:"url" example:"demo"`   // 分类地址
}

// LinksCache 友情链接 详情响应
type LinksCache struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
	*LinksCacheLite
	Summary    string `json:"summary" example:"demo"`        // 简介
	SourceShow string `json:"sourceShow" example:"/xxx.jpg"` // 资源图片地址
}

// CategoryCache 分类缓存
type CategoryCache struct {
	Id         uint64   `json:"id"`         // 数据ID
	Title      string   `json:"title"`      // 名称
	Url        string   `json:"url"`        // 分类地址
	Summary    string   `json:"summary"`    // 简介
	SourceShow string   `json:"sourceShow"` // 资源图片地址
	Num        uint64   `json:"num"`        // 数据量
	Posts      []string `json:"posts"`      // 对应的文章
}

// TagCache 标签缓存
type TagCache struct {
	Id         uint64   `json:"id"`         // 数据ID
	Title      string   `json:"title"`      // 名称
	Url        string   `json:"url"`        // 分类地址
	Summary    string   `json:"summary"`    // 简介
	SourceShow string   `json:"sourceShow"` // 资源图片地址
	Num        uint64   `json:"num"`        // 数据量
	Posts      []string `json:"posts"`      // 对应的文章
}

// BlogSort 排序，三大组排序（时间单独一个序）
type BlogSort struct {
	Url string
	Num uint64
}

// PostSortCache 顺序缓存
type PostSortCache struct {
	News  []uint64 // 新旧
	Views []uint64 // 查看
	Goods []uint64 // 好文
	Hots  []uint64 // 上升
}

// BlogSortArray 自定义排序
type BlogSortArray []*BlogSort

func (p BlogSortArray) Len() int {
	return len(p)
}

// Less 倒序处理
func (p BlogSortArray) Less(i, j int) bool {
	return p[i].Num > p[j].Num
}

func (p BlogSortArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// SyncBlogs 每小时刷新一次缓存
func SyncBlogs(traceID string) {
	for {
		log.InfoTF(traceID, "Start SyncBlogPostAllCache")
		_ = SyncBlogPostAllCache(traceID)
		log.InfoTF(traceID, "End SyncBlogPostAllCache")
		time.Sleep(time.Hour)
	}
}

// URLSet SITEMAP对象
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []*URL   `xml:"url"`
}
type URL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty"`
	Priority   float64 `xml:"priority,omitempty"`
}

// SyncBlogPostAllCache 同步全部文章/分类/标签/专题的缓存
func SyncBlogPostAllCache(traceID string) (err error) {
	// 先处理缓存入库
	postMapInnerDb(traceID)
	firstUrl := makeFirstUrls()
	cateAndTagHaveMap := make(map[string]bool)
	categoryUrl := make([]*URL, 0)
	tagUrl := make([]*URL, 0)
	postUrl := make([]*URL, 0)
	// 刷新Banner缓存
	lastUrl := SyncBanners(traceID)
	// 刷新友链
	SyncLinks(traceID)

	// 获取已发布的文章列表
	posts, err := blogDB.PostTable.Executor().GetPosts()
	if err != nil {
		log.ErrorTF(traceID, "SyncBlogPostAllCache Fail . Err Is : %v", err)
		return
	}
	postsMap := make(map[uint64]*PostCache)
	postNewArray := make([]string, len(posts))
	postViewSortArray := make(BlogSortArray, len(posts))
	postGoodSortArray := make(BlogSortArray, len(posts))
	postHotSortArray := make(BlogSortArray, len(posts))
	categorySortArray := make(BlogSortArray, 0)
	tagSortArray := make(BlogSortArray, 0)
	topicSortArray := make(BlogSortArray, 0)
	// 查分类
	categoryMap, err := getCategory(traceID)
	if err != nil {
		return
	}
	// 查标签
	tagMap, err := getTag(traceID)
	if err != nil {
		return
	}
	// 遍历文章
	for i, post := range posts {
		// 获取主图信息
		source, err := blogDB.SourceTable.GetOneById(post.SourceId)
		if err != nil {
			log.ErrorTF(traceID, "SyncBlogPostAllCache GetSource %d Fail . Err Is : %v", post.SourceId, err)
			continue
		}
		// 列表页使用_2图
		postsMap[post.Id] = &PostCache{
			Id:           post.Id,
			Title:        post.Title,
			Url:          post.Url,
			Summary:      post.Summary,
			SourcePath:   fmt.Sprintf(constant.SourceFilePath, source.FilePath, fmt.Sprintf("%d", source.Id)),
			SourceBack:   source.BackEnd,
			Category:     categoryMap[post.CategoryId].Url,
			CategoryName: categoryMap[post.CategoryId].Title,
			PushAt:       post.PushAt.Format("2006-01-02"),
			Views:        post.Views,
			Goods:        post.Goods,
			Hots:         post.Hots,
		}
		// SiteMap日志
		doDay := post.PushAt.Format("2006-01-02")
		// 分组追加
		categoryUrl = setCateTagUrls(categoryUrl, cateAndTagHaveMap, doDay, "category", categoryMap[post.CategoryId].Url, post.CategoryId)
		// 为分类填充数据
		categoryMap[post.CategoryId].Num++
		categoryMap[post.CategoryId].Posts = append(categoryMap[post.CategoryId].Posts, post.Url)
		// 查询归属标签
		tagIds, err := blogDB.PostTagTable.Executor().GetTagIds(post.Id)
		if err != nil {
			log.ErrorTF(traceID, "SyncBlogPostAllCache GetTag %d Fail . Err Is : %v", post.Id, err)
			continue
		}
		tagUrls := make([]string, len(tagIds))
		tagNames := make([]string, len(tagIds))
		for i, tagId := range tagIds {
			tagUrls[i] = tagMap[tagId].Url
			tagNames[i] = tagMap[tagId].Title
			tagMap[tagId].Num++
			tagMap[tagId].Posts = append(tagMap[tagId].Posts, post.Url)
			// 标签追加
			tagUrl = setCateTagUrls(tagUrl, cateAndTagHaveMap, doDay, "tag", tagMap[tagId].Url, tagId)
		}
		postsMap[post.Id].TagUrls = tagUrls
		postsMap[post.Id].TagNames = tagNames
		// 填充时间顺序
		postNewArray[i] = post.Url
		// about 放在最后
		if post.Url == "about" {
			postViewSortArray[i] = &BlogSort{Url: post.Url, Num: 0}
			postGoodSortArray[i] = &BlogSort{Url: post.Url, Num: 0}
			postHotSortArray[i] = &BlogSort{Url: post.Url, Num: 0}
		} else {
			postViewSortArray[i] = &BlogSort{Url: post.Url, Num: post.Views}
			postGoodSortArray[i] = &BlogSort{Url: post.Url, Num: post.Goods}
			postHotSortArray[i] = &BlogSort{Url: post.Url, Num: post.Hots}
		}
		// 开始文章添加
		postUrl = setPostUrls(postUrl, doDay, post.Url)
	}
	// 循环完成后，填充分类、标签、主题的排序对象
	for _, caC := range categoryMap {
		categorySortArray = append(categorySortArray, &BlogSort{Url: caC.Url, Num: caC.Num})
	}
	for _, taC := range tagMap {
		tagSortArray = append(tagSortArray, &BlogSort{Url: taC.Url, Num: taC.Num})
	}
	// 对数据进行Sort
	sort.Sort(postViewSortArray)
	sort.Sort(postGoodSortArray)
	sort.Sort(postHotSortArray)
	sort.Sort(categorySortArray)
	sort.Sort(tagSortArray)
	sort.Sort(topicSortArray)
	// 组装ID缓存对象
	pageSortCache := &BlogListCache{
		PostNews:  postNewArray,
		PostViews: sortToArray(postViewSortArray),
		PostGoods: sortToArray(postGoodSortArray),
		PostHots:  sortToArray(postHotSortArray),
		Category:  sortToArray(categorySortArray),
		Tag:       sortToArray(tagSortArray),
	}
	// 保存缓存
	err = redis.Set(constant.PagePostCache, makePostUrlMap(postsMap), 0)
	if err != nil {
		log.InfoTF(traceID, "SyncBlogPostAllCache CachePagePosts Fail . Err Is : %v", err)
	}
	err = redis.Set(constant.PageCategoryCache, makeCategoryUrlMap(categoryMap), 0)
	if err != nil {
		log.InfoTF(traceID, "SyncBlogPostAllCache CachePageCategories Fail . Err Is : %v", err)
	}
	err = redis.Set(constant.PageTagCache, makeTagUrlMap(tagMap), 0)
	if err != nil {
		log.InfoTF(traceID, "SyncBlogPostAllCache PageTagCache Fail . Err Is : %v", err)
	}
	err = redis.Set(constant.PagePostSortCache, pageSortCache, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncBlogPostAllCache CachePageSorts Fail . Err Is : %v", err)
	}
	// 生成siteMAP缓存
	makeSiteMapCache(traceID, firstUrl, categoryUrl, tagUrl, postUrl, lastUrl)
	return
}

// makePostUrlMap 生产文章URLMap
func makePostUrlMap(sMap map[uint64]*PostCache) map[string]*PostCache {
	res := make(map[string]*PostCache)
	for _, s := range sMap {
		res[s.Url] = s
	}
	return res
}

// makeCategoryUrlMap 生产分类URLMap
func makeCategoryUrlMap(sMap map[uint64]*CategoryCache) map[string]*CategoryCache {
	res := make(map[string]*CategoryCache)
	for _, s := range sMap {
		res[s.Url] = s
	}
	return res
}

// makeTagUrlMap 生产标签URLMap
func makeTagUrlMap(sMap map[uint64]*TagCache) map[string]*TagCache {
	res := make(map[string]*TagCache)
	for _, s := range sMap {
		res[s.Url] = s
	}
	return res
}

// Sort转ID
func sortToArray(sortArray BlogSortArray) []string {
	array := make([]string, len(sortArray))
	for i, sor := range sortArray {
		array[i] = sor.Url
	}
	return array
}

// 读取标签数据
func getTag(traceID string) (tagMap map[uint64]*TagCache, err error) {
	tagMap = make(map[uint64]*TagCache)
	// 最后需要组装
	tagList, err := blogDB.TagTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "SyncBlogPostAllCache GetTag Fail . Err Is : %v", err)
		return
	}
	for _, tag := range tagList {
		source, err := blogDB.SourceTable.GetOneById(tag.SourceId)
		if err != nil {
			log.ErrorTF(traceID, "SyncBlogPostAllCache GetTagSource %d Fail . Err Is : %v", tag.SourceId, err)
			continue
		}
		tagMap[tag.Id] = &TagCache{
			Id:         tag.Id,
			Title:      tag.Title,
			Url:        tag.Url,
			Summary:    tag.Summary,
			SourceShow: fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version),
			Num:        0,
		}
	}
	return
}

// 读取分类数据
func getCategory(traceID string) (categoryMap map[uint64]*CategoryCache, err error) {
	categoryMap = make(map[uint64]*CategoryCache)
	// 最后需要组装
	categoryList, err := blogDB.CategoryTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "SyncBlogPostAllCache GetCategory Fail . Err Is : %v", err)
		return
	}
	for _, category := range categoryList {
		source, err := blogDB.SourceTable.GetOneById(category.SourceId)
		if err != nil {
			log.ErrorTF(traceID, "SyncBlogPostAllCache GetCategorySource %d Fail . Err Is : %v", category.SourceId, err)
			continue
		}
		categoryMap[category.Id] = &CategoryCache{
			Id:         category.Id,
			Title:      category.Title,
			Url:        category.Url,
			Summary:    category.Summary,
			SourceShow: fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version),
			Num:        0,
		}
	}
	return
}

// SyncBanners 刷新Banner
func SyncBanners(traceID string) []*URL {
	banners, err := blogDB.BannerTable.Executor().GetBanners()
	if err != nil {
		log.ErrorTF(traceID, "SyncBanners GetBanners Fail . Err Is : %v", err)
		return nil
	}
	bannerCache := make([]*BannerCache, 0)
	pageCache := make([]*BannerCache, 0)
	for _, banner := range banners {
		cache := &BannerCache{
			Title:   banner.Title,
			Tag:     banner.Tag,
			Url:     banner.Url,
			Summary: banner.Summary,
		}
		source, err := blogDB.SourceTable.GetOneById(banner.SourceId)
		if err != nil {
			log.ErrorTF(traceID, "SyncBanners GetBannerSource %d Fail . Err Is : %v", banner.SourceId, err)
		}
		cache.SourceShow = fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version)
		if banner.Type == constant.StatusOpen {
			bannerCache = append(bannerCache, cache)
		} else {
			pageCache = append(pageCache, cache)
		}
	}
	// 写入缓存
	err = redis.Set(constant.PageBannerCache, bannerCache, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncBanners SetBannersCache Fail . Err Is : %v", err)
	}
	err = redis.Set(constant.PagePageCache, pageCache, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncBanners SetPagesCache Fail . Err Is : %v", err)
	}
	return makeLastUrls(pageCache)
}

// SyncLinks 刷新Links
func SyncLinks(traceID string) {
	links, err := blogDB.LinksTable.Executor().GetLinks()
	if err != nil {
		log.ErrorTF(traceID, "SyncLinks GetBanners Fail . Err Is : %v", err)
		return
	}
	linkCache := make([]*LinksCache, 0)
	for _, link := range links {
		cache := &LinksCache{
			Id: link.Id,
			LinksCacheLite: &LinksCacheLite{
				Title: link.Title,
				Url:   link.Url,
			},
			Summary: link.Summary,
		}
		source, err := blogDB.SourceTable.GetOneById(link.SourceId)
		if err != nil {
			log.ErrorTF(traceID, "SyncLinks GetSource %d Fail . Err Is : %v", link.SourceId, err)
		}
		cache.SourceShow = fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version)

		linkCache = append(linkCache, cache)
	}
	// 写入缓存
	err = redis.Set(constant.PageLinksCache, linkCache, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncLinks SetLinksCache Fail . Err Is : %v", err)
	}
}

// SyncPostMain 刷新文章正文16H
func SyncPostMain(traceID string, postBase *PostCache) (res *PostMainCache) {
	postMore, err := blogDB.PostMoreTable.GetOneById(postBase.Id)
	if err != nil {
		return
	}
	res = &PostMainCache{
		Toc:  postMore.Toc,
		Html: postMore.Html,
	}
	// 开始计算Like
	like := make([]string, 0)
	like = append(like, postBase.Url)
	if len(postBase.TagUrls) > 0 {
		tagPosts := make([]string, 0)
		tagCache := GetTagCache(traceID)
		for _, tag := range postBase.TagUrls {
			if tagInfo, ok := tagCache[tag]; ok {
				tagPosts = append(tagPosts, tagInfo.Posts...)
			}
		}
		// 去重
		tagPosts = utils.ArrayToSet(tagPosts)
		if len(tagPosts) > 0 {
			// 排除自己意外取9个
			like = utils.ShuffleAndSelectEx(tagPosts, like, 9)
		}
	}
	// 不满足9位，提前分组下文章
	if len(like) < 10 {
		// 由于要去掉自己，所以去掉前应该为10
		needCount := 10 - len(like)
		category := GetCategoryCache(traceID)
		if cat, ok := category[postBase.Category]; ok {
			if len(cat.Posts) > 0 {
				// 尝试补充剩余的数量
				like = utils.ShuffleAndSelectEx(cat.Posts, like, needCount)
			}
		}
	}
	// 不满足9位，获取最新文章 // 由于要去掉自己，所以去掉前应该为10
	if len(like) < 10 {
		needCount := 10 - len(like)
		postSort := GetPostSortCache(traceID)
		// 从最新文章获得前N位的数据，去掉已存在的 // 尝试补充剩余的数量
		like = utils.GetFirstAndExt(postSort.PostNews, like, needCount)
	}
	// 去掉第一位是自己
	res.Like = like[1:]
	// 写入缓存
	err = redis.SetByTimeDuration(fmt.Sprintf(constant.PagePostUrl, postBase.Url), res, 16*time.Hour)
	if err != nil {
		log.InfoTF(traceID, "SyncPostMain Fail . Err Is : %v", err)
	}
	return
}

// postMapInnerDb 更新文章处理数据
func postMapInnerDb(traceId string) {
	// 判断是否执行Hots消除
	today := time.Now().Format("20060102")
	updateDate, _ := redis.Get(constant.PagePostUpdateDate)
	needRun := today != updateDate
	if needRun {
		_ = redis.Set(constant.PagePostUpdateDate, today, 0)
	}
	postMap := GetPostCache(traceId)
	if postMap == nil {
		return
	}
	i := 0
	req := make([]*blogDB.PostUpdateData, 0)
	// 300个一轮
	for _, post := range postMap {
		if i == 300 {
			// 先执行一轮
			err := blogDB.PostTable.Executor().UpdatePostDataBatch(req)
			if err != nil {
				log.ErrorTF(traceId, "UpdatePostDataBatch Fail . Err Is : %v", err)
			}
			i = 0
			req = make([]*blogDB.PostUpdateData, 0)
		}
		hots := post.Hots
		if needRun {
			hots = uint64(math.Ceil(float64(hots) * 0.99))
		}
		req = append(req, &blogDB.PostUpdateData{
			Id:    post.Id,
			Views: post.Views,
			Goods: post.Goods,
			Hots:  hots,
		})
		i++
	}
	if len(req) > 0 {
		// 最后一轮
		err := blogDB.PostTable.Executor().UpdatePostDataBatch(req)
		if err != nil {
			log.ErrorTF(traceId, "UpdatePostDataBatch Fail . Err Is : %v", err)
		}
	}
	log.InfoTF(traceId, "UpdatePostDataBatch Done ")
}
