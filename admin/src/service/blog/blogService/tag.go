package blogService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddTag 创建标签
func AddTag(traceID string, req *blogModel.TagAddReq) *baseModel.ResBody {
	// 先处理资源
	if len(req.Source) < 1 {
		return baseModel.Fail(constant.TagAddNoScNG)
	}
	dbSourceReq, errNum, err := thirdAddSource(traceID, &blogModel.SourceAddReq{
		FileType: "1", // 1 ICON
		SourceDoReq: blogModel.SourceDoReq{
			Name:         "Tag_" + req.Title,
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
	err = blogDB.TagTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddTag Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkTagDBErr(err)
	}
	// 分类无需独立的资源关联
	return baseModel.Success(constant.TagAddSS, true)
}

// PageTag 查询标签分页
func PageTag(traceID string, req *blogModel.TagPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.TagTable.Page(tagPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TagGetNG)
	}
	// 循环List获取资源地址
	sourceIds := make([]uint64, len(list))
	for i, r := range list {
		sourceIds[i] = r.SourceId
	}
	sources, err := blogDB.SourceTable.GetByIds(sourceIds)
	if err != nil {
		log.ErrorTF(traceID, "PageTag GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TagGetNG)
	}

	// 读取资源地址
	sourceMap := make(map[uint64]string, len(sources))
	for _, r := range sources {
		sourceMap[r.Id] = fmt.Sprintf(constant.SourceFileUrl, r.FilePath, fmt.Sprintf("%d", r.Id), r.BackEnd, r.Version)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToTagPageRes(list, sourceMap), total))
}

// GetTag 标签详情
func GetTag(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.TagTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TagGetNG)
	}
	source, err := blogDB.SourceTable.GetOneById(res.SourceId)
	if err != nil {
		log.ErrorTF(traceID, "GetTag GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TagGetNG)
	}
	return baseModel.SuccessUnPop(blogModel.ToTagGetRes(&res, fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version)))
}

// EditTag 编辑标签
func EditTag(traceID string, req *blogModel.TagEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.TagTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TagGetNG)
	}
	// 如果存在图片资源，更新图片
	if len(req.Source) > 0 {
		errNum, err := thirdEditSource(traceID, &blogModel.SourceEditReq{
			Id:       dbReq.SourceId,
			FileType: "1", // 1 ICON
			SourceDoReq: blogModel.SourceDoReq{
				Name:         "Tag_" + dbReq.Title,
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
	err = blogDB.TagTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditTag %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkTagDBErr(err)
	}
	return baseModel.Success(constant.TagEditSS, true)
}

// DelTag 标签移除
func DelTag(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := blogDB.TagTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TagGetNG)
	}
	// 先删除文章的关联
	err = blogDB.PostTagTable.DeleteByMap(map[string]any{"tag_id": dbReq.Id})
	if err != nil {
		log.ErrorTF(traceID, "DeletePostTag %d Fail . Err Is : %v", req.Id, err)
		return baseModel.Fail(constant.TagDelPostNG)
	}
	// 物理删除
	err = blogDB.TagTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelTag %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.TagDelNG)
	}
	return baseModel.Success(constant.TagDelSS, true)
}

// ListTag 标签列表
func ListTag(traceID string) *baseModel.ResBody {
	// 查询全部标签
	resList, err := blogDB.TagTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "ListTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.TagGetNG)
	}
	resp := make([]*baseModel.SelectNumRes, 0)
	for _, item := range resList {
		resp = append(resp, &baseModel.SelectNumRes{
			Label: item.Title,
			Value: item.Id,
		})
	}
	return baseModel.SuccessUnPop(resp)
}
