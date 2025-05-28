package blogService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddLinks 创建友情链接
func AddLinks(traceID string, req *blogModel.LinksAddReq) *baseModel.ResBody {
	// 先处理资源
	if len(req.Source) < 1 {
		return baseModel.Fail(constant.LinksAddNoScNG)
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
	// 创建对象初始化
	dbReq := req.ToDbReq()
	dbReq.SourceId = dbSourceReq.Id
	err = blogDB.LinksTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddLinks Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkLinksDBErr(err)
	}
	// 分类无需独立的资源关联
	return baseModel.Success(constant.LinksAddSS, true)
}

// PageLinks 查询友情链接分页
func PageLinks(traceID string, req *blogModel.LinksPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.LinksTable.Page(LinksPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageLinks Fail . Err Is : %v", err)
		return baseModel.Fail(constant.LinksGetNG)
	}
	// 循环List获取资源地址
	sourceIds := make([]uint64, len(list))
	for i, r := range list {
		sourceIds[i] = r.SourceId
	}
	sources, err := blogDB.SourceTable.GetByIds(sourceIds)
	if err != nil {
		log.ErrorTF(traceID, "PageLinks GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.LinksGetNG)
	}

	// 读取资源地址
	sourceMap := make(map[uint64]string, len(sources))
	for _, r := range sources {
		sourceMap[r.Id] = fmt.Sprintf(constant.SourceFileUrl, r.FilePath, fmt.Sprintf("%d", r.Id), r.BackEnd, r.Version)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToLinksPageRes(list, sourceMap), total))
}

// GetLinks 友情链接详情
func GetLinks(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.LinksTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetLinks Fail . Err Is : %v", err)
		return baseModel.Fail(constant.LinksGetNG)
	}
	source, err := blogDB.SourceTable.GetOneById(res.SourceId)
	if err != nil {
		log.ErrorTF(traceID, "GetLinks GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.LinksGetNG)
	}
	return baseModel.SuccessUnPop(blogModel.ToLinksGetRes(&res, fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version)))
}

// EditLinks 编辑友情链接
func EditLinks(traceID string, req *blogModel.LinksEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.LinksTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetLinks Fail . Err Is : %v", err)
		return baseModel.Fail(constant.LinksGetNG)
	}
	// 如果存在图片资源，更新图片
	if len(req.Source) > 0 {
		errNum, err := thirdEditSource(traceID, &blogModel.SourceEditReq{
			Id:       dbReq.SourceId,
			FileType: "1", // 1 ICON
			SourceDoReq: blogModel.SourceDoReq{
				Name:         "Links_" + dbReq.Title,
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
	err = blogDB.LinksTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditLinks %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkLinksDBErr(err)
	}
	return baseModel.Success(constant.LinksEditSS, true)
}

// DelLinks 友情链接移除
func DelLinks(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := blogDB.LinksTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetLinks Fail . Err Is : %v", err)
		return baseModel.Fail(constant.LinksGetNG)
	}
	// 物理删除
	err = blogDB.LinksTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelLinks %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.LinksDelNG)
	}
	return baseModel.Success(constant.LinksDelSS, true)
}
