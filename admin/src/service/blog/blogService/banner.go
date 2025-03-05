package blogService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddBanner 创建Banner
func AddBanner(traceID string, req *blogModel.BannerAddReq) *baseModel.ResBody {
	// 先处理资源
	if len(req.Source) < 1 {
		return baseModel.Fail(constant.BannerAddNoScNG)
	}
	dbSourceReq, errNum, err := thirdAddSource(traceID, &blogModel.SourceAddReq{
		FileType: "0", // 1 MAIN
		SourceDoReq: blogModel.SourceDoReq{
			Name:         "Banner_" + req.Title,
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
	err = blogDB.BannerTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddBanner Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkBannerDBErr(err)
	}
	// 分类无需独立的资源关联
	return baseModel.Success(constant.BannerAddSS, true)
}

// PageBanner 查询Banner分页
func PageBanner(traceID string, req *blogModel.BannerPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.BannerTable.Page(BannerPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageBanner Fail . Err Is : %v", err)
		return baseModel.Fail(constant.BannerGetNG)
	}
	// 循环List获取资源地址
	sourceIds := make([]uint64, len(list))
	for i, r := range list {
		sourceIds[i] = r.SourceId
	}
	sources, err := blogDB.SourceTable.GetByIds(sourceIds)
	if err != nil {
		log.ErrorTF(traceID, "PageBanner GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.BannerGetNG)
	}

	// 读取资源地址
	sourceMap := make(map[uint64]string, len(sources))
	for _, r := range sources {
		sourceMap[r.Id] = fmt.Sprintf(constant.SourceFileUrl, r.FilePath, fmt.Sprintf("%d", r.Id), r.BackEnd, r.Version)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToBannerPageRes(list, sourceMap), total))
}

// GetBanner Banner详情
func GetBanner(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.BannerTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetBanner Fail . Err Is : %v", err)
		return baseModel.Fail(constant.BannerGetNG)
	}
	source, err := blogDB.SourceTable.GetOneById(res.SourceId)
	if err != nil {
		log.ErrorTF(traceID, "GetBanner GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.BannerGetNG)
	}
	return baseModel.SuccessUnPop(blogModel.ToBannerGetRes(&res, fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, source.Version)))
}

// EditBanner 编辑Banner
func EditBanner(traceID string, req *blogModel.BannerEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.BannerTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetBanner Fail . Err Is : %v", err)
		return baseModel.Fail(constant.BannerGetNG)
	}
	// 如果存在图片资源，更新图片
	if len(req.Source) > 0 {
		errNum, err := thirdEditSource(traceID, &blogModel.SourceEditReq{
			Id:       dbReq.SourceId,
			FileType: "0", // 1 MAIN
			SourceDoReq: blogModel.SourceDoReq{
				Name:         "Banner_" + dbReq.Title,
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
	err = blogDB.BannerTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditBanner %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkBannerDBErr(err)
	}
	return baseModel.Success(constant.BannerEditSS, true)
}

// DelBanner Banner移除
func DelBanner(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := blogDB.BannerTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetBanner Fail . Err Is : %v", err)
		return baseModel.Fail(constant.BannerGetNG)
	}
	// 物理删除
	err = blogDB.BannerTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelBanner %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.BannerDelNG)
	}
	return baseModel.Success(constant.BannerDelSS, true)
}
