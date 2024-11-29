package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddSource 创建资源配置表
func AddSource(traceID string, req *blogModel.SourceAddReq) *baseModel.ResBody {

	dbReq, errNum, err := thirdAddSource(traceID, req)
	if err != nil {
		switch errNum {
		case 1:
			// 解析数据库错误
			return checkSourceDBErr(err)
		case 2:
			return baseModel.Fail(constant.SourceFileUpNg)
		}
	}
	return baseModel.Success(constant.SourceAddSS, dbReq.Id)
}

// PageSource 查询资源配置表分页
func PageSource(traceID string, req *blogModel.SourcePageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.SourceTable.Page(sourcePageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.SourceGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToSourcePageRes(list), total))
}

// GetSource 资源配置表详情
func GetSource(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.SourceTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.SourceGetNG)
	}
	return baseModel.SuccessUnPop(blogModel.ToSourceGetRes(&res))
}

// EditSource 编辑资源配置表
func EditSource(traceID string, req *blogModel.SourceEditReq) *baseModel.ResBody {
	errNum, err := thirdEditSource(traceID, req)
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
	return baseModel.Success(constant.SourceEditSS, true)
}

// DelSource 资源配置表清理
func DelSource(traceID string) *baseModel.ResBody {
	return baseModel.Success(constant.SourceDelSS, true)
}
