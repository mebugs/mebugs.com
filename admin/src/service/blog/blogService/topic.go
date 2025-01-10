package blogService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddTopic 创建文章专题
func AddTopic(traceID string, req *blogModel.TopicAddReq) *baseModel.ResBody {
	// 先处理资源
	if len(req.Source) < 1 {
		return baseModel.Fail(constant.TopicAddNoScNG)
	}
	dbSourceReq, errNum, err := thirdAddSource(traceID, &blogModel.SourceAddReq{
		FileType: "0", // 0 MAIN
		SourceDoReq: blogModel.SourceDoReq{
			Name:         "Topic_" + req.Title,
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
	// 创建对象初始化
	dbReq := req.ToDbReq()
	dbReq.SourceId = dbSourceReq.Id
	err = blogDB.TopicTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddTopic Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkTopicDBErr(err)
	}
	// 创建文件和主题的关系
	err = syncTopicPosts(traceID, dbReq.Id, req.PostIdSet, false)
	if err != nil {
		// 移除当前值
		errD := blogDB.TopicTable.DeleteOne(dbReq.Id)
		if errD != nil {
			log.ErrorTF(traceID, "AddTopic syncTopicPosts %d Rollback Fail . Err Is : %v", dbReq.Id, err)
		}
		return baseModel.Fail(constant.TopicPostNG)
	}
	return baseModel.Success(constant.TopicAddSS, true)
}

// PageTopic 查询文章专题分页
func PageTopic(traceID string, req *blogModel.TopicPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.TopicTable.Page(topicPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageTopic Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TopicGetNG)
	}
	// 循环List获取资源地址
	sourceIds := make([]uint64, len(list))
	for i, r := range list {
		sourceIds[i] = r.SourceId
	}
	sources, err := blogDB.SourceTable.GetByIds(sourceIds)
	if err != nil {
		log.ErrorTF(traceID, "PageTopic GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TopicGetNG)
	}

	// 读取资源地址
	sourceMap := make(map[uint64]string, len(sources))
	for _, r := range sources {
		sourceMap[r.Id] = fmt.Sprintf(constant.SourceFileUrl, r.FilePath, fmt.Sprintf("%d", r.Id), r.BackEnd, r.Version)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToTopicPageRes(list, sourceMap), total))
}

// GetTopic 文章专题详情
func GetTopic(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.TopicTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetTopic Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TopicGetNG)
	}
	source, err := blogDB.SourceTable.GetOneById(res.SourceId)
	if err != nil {
		log.ErrorTF(traceID, "GetTopic GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TopicGetNG)
	}
	resObj := blogModel.ToTopicGetRes(&res, fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version))
	// 查询文章
	postIds, posts, _ := getTopicPosts(traceID, res.Id)
	resObj.PostIds = postIds
	resObj.Posts = posts
	return baseModel.SuccessUnPop(resObj)
}

// EditTopic 编辑文章专题
func EditTopic(traceID string, req *blogModel.TopicEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.TopicTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetTopic Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TopicGetNG)
	}
	// 如果存在图片资源，更新图片
	if len(req.Source) > 0 {
		errNum, err := thirdEditSource(traceID, &blogModel.SourceEditReq{
			Id:       dbReq.SourceId,
			FileType: "0", // 0 MAIN
			SourceDoReq: blogModel.SourceDoReq{
				Name:         "Topic_" + req.Title,
				FileStrArray: req.Source,
			},
		})
		if err != nil {
			switch errNum {
			case 1:
				return baseModel.Fail(constant.SourceGetNG)
			case 2:
				return baseModel.Fail(constant.SourceFileUpNg)
			case 3:
				return checkSourceDBErr(err)
			}
		}
	}
	// 创建文件和主题的关系
	err = syncTopicPosts(traceID, dbReq.Id, req.PostIdSet, true)
	if err != nil {
		log.ErrorTF(traceID, "EditTopic syncTopicPosts %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.TopicPostNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = blogDB.TopicTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditTopic %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkTopicDBErr(err)
	}
	return baseModel.Success(constant.TopicEditSS, true)
}

// DelTopic 文章专题移除
func DelTopic(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := blogDB.TopicTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetTopic Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TopicGetNG)
	}
	// 先移除POST关系
	err = syncTopicPosts(traceID, dbReq.Id, nil, true)
	if err != nil {
		log.ErrorTF(traceID, "DelTopic syncTopicPosts %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.TopicPostNG)
	}
	// 物理删除
	err = blogDB.TopicTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelTopic %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.TopicDelNG)
	}
	return baseModel.Success(constant.TopicDelSS, true)
}
