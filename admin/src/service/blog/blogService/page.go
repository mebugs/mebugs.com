package blogService

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"gorm.io/gorm"
	"net/http"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"siteol.com/smart/src/common/redis"
	"strings"
	"time"
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
	posts := getPostByUrlSort(traceID, indexIds[2], indexIds[3], indexIds[4], indexIds[5])
	indexData.PostsNew = posts[0]
	indexData.PostsView = posts[1]
	indexData.PostsGood = posts[2]
	indexData.PostsHot = posts[3]
	return baseModel.Success(constant.Success, indexData)
}

// GetPage 获取页面数据
func GetPage(traceID string) *baseModel.ResBody {
	return baseModel.Success(constant.Success, cacheModel.GetPagesCache(traceID))
}

// GetPageDetail 获取页面数据
func GetPageDetail(traceID string, req *blogModel.PostReq) *baseModel.ResBody {
	allPage := cacheModel.GetPagesCache(traceID)
	res := &cacheModel.BannerCache{}
	for _, item := range allPage {
		if item.Url == req.Url {
			res = item
			// 读取友情链接
			if item.Url == "/page/link" {
				res.Links = cacheModel.GetLinksCache(traceID)
			}
			// 读取全部文章
			if item.Url == "/page/map" {
				indexIds, _ := getRunUrls(traceID, 5, 1)
				posts := getPostByUrlSortSingle(traceID, indexIds[0])
				res.LitePosts = cacheModel.ToLitePost(posts)
			}
		}
	}
	return baseModel.Success(constant.Success, res)
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
	postBase.Views++
	postBase.Hots++
	postBase.Goods++
	// 写回缓存
	allMap[postBase.Url] = postBase
	_ = redis.Set(constant.PagePostCache, allMap, 0)
	return baseModel.Success(constant.Success, nil)
}

// AddComm 提交文章评论
func AddComm(traceID string, req *blogModel.PostCommReq) *baseModel.ResBody {
	// 查询用户
	now := time.Now()
	user, err := blogDB.UserTable.GetOneByObject(&blogDB.User{ClientId: req.User.ClientId})
	if err != nil {
		user = blogDB.User{
			Status:   constant.StatusOpen,
			CreateAt: &now,
		}
	}
	user.UpdateAt = &now
	user.Name = req.User.Name
	user.Email = req.User.Email
	user.Summary = req.User.Summary
	user.SourceId = req.User.SourceId
	user.ClientId = req.User.ClientId
	if user.ThirdUrl != req.User.Url {
		user.WaitUrl = req.User.Url
	}
	if user.Id == 0 {
		err = blogDB.UserTable.InsertOne(&user)
	} else {
		err = blogDB.UserTable.UpdateOne(&user)
	}
	if err != nil {
		log.ErrorTF(traceID, "AddComm Add/UpDate User Failed . Err Is %v", err)
		return baseModel.FailWithMsg("评论用户初始化失败")
	}
	// 开始添加评论（当日同UID在相同文章下最多提交5条评论）
	userPostCache := cacheModel.GetPostCommUserCache(traceID, user.Id)
	if num, ok := userPostCache[req.PostId]; ok {
		if num >= 5 {
			return baseModel.FailWithMsg("用户超出文章的当日评论限制")
		}
		userPostCache[req.PostId] = num + 1
	} else {
		userPostCache[req.PostId] = 1
	}
	err = blogDB.PostCommentTable.InsertOne(&blogDB.PostComment{
		Id:       0,
		PostId:   req.PostId,
		Level:    req.Level,
		Uid:      user.Id,
		Rid:      req.Rid,
		Info:     req.Info,
		Status:   constant.StatusLock, // 待审核
		CreateAt: &now,
		UpdateAt: &now,
	})
	if err != nil {
		log.ErrorTF(traceID, "AddComm PostComment Failed . Err Is %v", err)
		return baseModel.FailWithMsg("评论提交失败，系统异常")
	}
	tom := now.AddDate(0, 0, 1)
	expTime := time.Date(tom.Year(), tom.Month(), tom.Day(), 0, 0, 0, 0, time.Local)
	expD := expTime.Sub(now)
	_ = redis.SetByTimeDuration(fmt.Sprintf(constant.PagePostCommUserCache, user.Id), userPostCache, expD)
	return baseModel.Success(constant.Success, nil)
}

// Comments 评论查询
func Comments(traceID string, req *blogModel.CommentsReq) *baseModel.ResBody {
	res := &blogModel.CommentsRes{}
	// 查询全部评论，以时间倒序
	comments, err := blogDB.PostCommentTable.Executor().Comments(req.PostId)
	if err != nil {
		return baseModel.Success(constant.Success, res)
	}
	var uid uint64
	userMap := make(map[uint64]*blogModel.CommentsUser)
	if req.ClientId != "" {
		user, err := blogDB.UserTable.GetOneByObject(&blogDB.User{ClientId: req.ClientId})
		if err == nil {
			uid = user.Id
			userMap[uid] = blogModel.ToCommentsUser(&user)
		}
	}
	// 记录需要展示的ID信息
	total := 0
	ids := make([]uint64, 0)
	commentsMap := make(map[uint64]*blogModel.Comments)
	for _, comment := range comments {
		// 关闭的评论
		if comment.Status == constant.StatusClose {
			continue
		}
		// 非受访用户的待评审评论
		if comment.Status == constant.StatusLock && comment.Uid != uid {
			continue
		}
		// 计数
		total++
		user, ok := userMap[comment.Uid]
		if !ok {
			dbUser, err := blogDB.UserTable.GetOneById(comment.Uid)
			if err != nil {
				user = blogModel.NormalUser
			} else {
				user = blogModel.ToCommentsUser(&dbUser)
			}
			userMap[comment.Uid] = user
		}
		// 一级评论
		if comment.Level == 0 {
			ids = append(ids, comment.Id)
			commentsMap[comment.Id] = blogModel.ToComments(user, comment)
		} else {
			// 二级评论
			rComm := commentsMap[comment.Rid] // 必然存在
			if rComm == nil {
				continue
			}
			rComm.Comments = append(rComm.Comments, blogModel.ToComments(user, comment))
			commentsMap[comment.Id] = rComm
		}
	}
	// 组装
	resComments := make([]*blogModel.Comments, len(ids))
	for i, id := range ids {
		resComments[i] = commentsMap[id]
	}
	res.Total = total
	res.Comments = resComments
	return baseModel.Success(constant.Success, res)
}

// PageLink 获取友链数据
func PageLink(traceID string) *baseModel.ResBody {
	links := cacheModel.GetLinksCache(traceID)
	liteLinks := make([]*cacheModel.LinksCacheLite, len(links))
	for i, link := range links {
		liteLinks[i] = link.LinksCacheLite
	}
	return baseModel.Success(constant.Success, liteLinks)
}

// LinkScan 链接扫描
func LinkScan(traceID string, req *blogModel.LinksScanReq) *baseModel.ResBody {
	// 读取站点信息
	// 获取网页内容
	resp, err := http.Get(req.Url)
	if err != nil {
		log.ErrorTF(traceID, "LinkScan Query %s Failed. Err Is %v ", req.Url, err)
		return baseModel.FailWithMsg("请求失败，请检查链接地址！")
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		log.ErrorTF(traceID, "LinkScan Query %s Res Failed. Code Is %d ", req.Url, resp.StatusCode)
		return baseModel.FailWithMsg(fmt.Sprintf("请求失败，响应码: %d！", resp.StatusCode))
	}
	// 解析HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.ErrorTF(traceID, "LinkScan Query %s Format Failed. Err Is %v ", req.Url, err)
		return baseModel.FailWithMsg("目标链接响应数据不符合HTML规范！")
	}
	var title, description, icon string
	// 遍历HTML节点
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil {
					title = n.FirstChild.Data
				}
			case "meta":
				if getAttr(n, "name") == "description" {
					description = getAttr(n, "content")
				}
			case "link":
				rel := getAttr(n, "rel")
				if strings.Contains(rel, "icon") && icon == "" {
					icon = getAttr(n, "href")
				}
			}
		}

		// 递归遍历子节点
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return baseModel.Success(constant.Success, &blogModel.LinksGetRes{
		Title:      title,
		Url:        req.Url,
		Summary:    description,
		SourceShow: getImgBase64(traceID, req.Url, icon),
	})
}

// LinkAdd 链接提交
func LinkAdd(traceID string, req *blogModel.LinksAddClientReq) *baseModel.ResBody {
	// 一天最多提交三次链接
	linkNo := cacheModel.GetClientLinkCache(traceID, req.ClientId)
	if linkNo >= 3 {
		return baseModel.FailWithMsg("用户超出友链提交的当日限制")
	} else {
		linkNo++
	}
	// 查询链接存在性
	one, err := blogDB.LinksTable.GetOneByObject(&blogDB.Links{Url: req.Url})
	if err != nil && !errors.As(err, &gorm.ErrRecordNotFound) {
		return baseModel.FailWithMsg("友链检查失败，系统异常")
	}
	if one.Url == req.Url {
		return baseModel.FailWithMsg("友链地址已存在")
	}
	dbSourceReq, errNum, err := thirdAddSource(traceID, &blogModel.SourceAddReq{
		FileType: "1", // 1 ICON
		SourceDoReq: blogModel.SourceDoReq{
			Name:         "Links_" + req.Title,
			FileStrArray: req.Source,
		},
	})
	if err != nil {
		switch errNum {
		case 1:
			// 解析数据库错误
			return checkSourceDBErr(err)
		case 2:
			return baseModel.Fail(constant.SourceFileUpNg)
		}
	}
	now := time.Now()
	link := &blogDB.Links{
		Title:    req.Title,
		Url:      req.Url,
		Summary:  req.Summary,
		SourceId: dbSourceReq.Id,
		Status:   "1",
		CreateAt: &now,
		UpdateAt: &now,
	}
	err = blogDB.LinksTable.InsertOne(link)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, constant.DBDuplicateErr) {
			if strings.Contains(errStr, "url_uni") {
				// 唯一索引错误
				return baseModel.FailWithMsg("友链地址已存在")
			}
		}
		log.ErrorTF(traceID, "LinkAdd Failed . Err Is %v", err)
		return baseModel.FailWithMsg("友链提交失败，系统异常")
	}
	tom := now.AddDate(0, 0, 1)
	expTime := time.Date(tom.Year(), tom.Month(), tom.Day(), 0, 0, 0, 0, time.Local)
	expD := expTime.Sub(now)
	_ = redis.SetByTimeDuration(fmt.Sprintf(fmt.Sprintf(constant.PageClientLinkCache, req.ClientId)), linkNo, expD)
	return baseModel.Success(constant.Success, nil)
}
