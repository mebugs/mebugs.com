package blogService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddPost 创建文章或页面
func AddPost(traceID string, req *blogModel.PostAddReq) *baseModel.ResBody {
	// 先处理资源
	if len(req.Source) < 1 {
		return baseModel.Fail(constant.PostAddNoScNG)
	}
	// 主图资源
	dbSourceReq, errNum, err := thirdAddSource(traceID, &blogModel.SourceAddReq{
		FileType: "0", // 0 MAIN
		SourceDoReq: blogModel.SourceDoReq{
			Name:         "Post_" + req.Title,
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
	dbReq.SourceId = dbSourceReq.Id // 关联主图
	// 先建文章
	err = blogDB.PostTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddPost Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkPostDBErr(err)
	}
	// 关联标签
	// 创建文章和标签的关系
	err = syncPostTags(traceID, dbReq.Id, req.TagIds, false)
	if err != nil {
		log.ErrorTF(traceID, "AddPost syncPostTags %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.PostTagNG)
	}
	// 处理文章内容
	err = blogDB.PostMoreTable.InsertOne(blogDB.PostMore{
		Id:   dbReq.Id,
		Toc:  req.Toc,
		Md:   req.Md,
		Html: req.Html,
	})
	if err != nil {
		log.ErrorTF(traceID, "AddPost AddPostMore %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.PostMoreNG)
	}
	// 记录资源绑定
	err = syncPostSources(traceID, dbReq.Id, req.SourceIds, false)
	if err != nil {
		log.ErrorTF(traceID, "AddPost syncPostSources %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.PostSourceNG)
	}
	return baseModel.Success(constant.PostAddSS, true)
}

// PagePost 查询文章或页面分页
func PagePost(traceID string, req *blogModel.PostPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.PostTable.Page(postPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PagePost Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostGetNG)
	}
	// 循环List获取资源地址
	sourceIds := make([]uint64, len(list))
	for i, r := range list {
		sourceIds[i] = r.SourceId
	}
	sources, err := blogDB.SourceTable.GetByIds(sourceIds)
	if err != nil {
		log.ErrorTF(traceID, "PageCategory GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostGetNG)
	}
	// 读取资源地址
	sourceMap := make(map[uint64]string, len(sources))
	for _, r := range sources {
		sourceMap[r.Id] = fmt.Sprintf(constant.SourceFileUrl, r.FilePath, fmt.Sprintf("%d", r.Id), r.BackEnd, r.Version)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToPostPageRes(list, sourceMap), total))
}

// GetPost 文章或页面详情
func GetPost(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.PostTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPost Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostGetNG)
	}
	// 获取文章拓展信息
	moreRes, err := blogDB.PostMoreTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPost More Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostGetNG)
	}
	// 获取标签列表
	tagIds, err := getPostTagIds(traceID, req.Id)
	if err != nil {
		return baseModel.Fail(constant.PostGetNG)
	}
	// 获取文章主图
	source, err := blogDB.SourceTable.GetOneById(res.SourceId)
	if err != nil {
		log.ErrorTF(traceID, "GetPost Source Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostGetNG)
	}
	sourceUrl := fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version)
	return baseModel.SuccessUnPop(blogModel.ToPostAllRes(&res, &moreRes, tagIds, sourceUrl))
}

// EditPost 编辑文章或页面
func EditPost(traceID string, req *blogModel.PostEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.PostTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPost Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostGetNG)
	}
	// 如果存在图片资源，更新图片
	if len(req.Source) > 0 {
		errNum, err := thirdEditSource(traceID, &blogModel.SourceEditReq{
			Id:       dbReq.SourceId,
			FileType: "0", // 0 MAIN
			SourceDoReq: blogModel.SourceDoReq{
				Name:         "Post_" + req.Title,
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
	// 对象更新
	req.ToDbReq(&dbReq)
	err = blogDB.PostTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditPost %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkPostDBErr(err)
	}
	// 关联标签
	// 创建文章和标签的关系
	err = syncPostTags(traceID, dbReq.Id, req.TagIds, true)
	if err != nil {
		log.ErrorTF(traceID, "EditPost syncPostTags %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.PostTagNG)
	}
	// 处理文章内容
	err = blogDB.PostMoreTable.UpdateOne(blogDB.PostMore{
		Id:   dbReq.Id,
		Toc:  req.Toc,
		Md:   req.Md,
		Html: req.Html,
	})
	if err != nil {
		log.ErrorTF(traceID, "EditPost EditPostMore %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.PostMoreNG)
	}
	// 记录资源绑定
	err = syncPostSources(traceID, dbReq.Id, req.SourceIds, true)
	if err != nil {
		log.ErrorTF(traceID, "EditPost syncPostSources %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.PostSourceNG)
	}
	return baseModel.Success(constant.PostEditSS, true)
}
