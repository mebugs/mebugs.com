package blogService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"siteol.com/smart/src/common/redis"
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
		return baseModel.Fail(constant.PageCommUserUpsertNG)
	}
	// 开始添加评论（当日同UID在相同文章下最多提交5条评论）
	userPostCache := cacheModel.GetPostCommUserCache(traceID, user.Id)
	if num, ok := userPostCache[req.PostId]; ok {
		if num >= 5 {
			return baseModel.Fail(constant.PageCommCommentPutLimit)
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
		return baseModel.Fail(constant.PageCommCommentPutNG)
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
