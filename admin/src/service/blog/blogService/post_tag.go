package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddPostTag 创建文章引用标签
func AddPostTag(traceID string, req *blogModel.PostTagAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := blogDB.PostTagTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddPostTag Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkPostTagDBErr(err)
	}
	return baseModel.Success(constant.PostTagAddSS, true)
}

// PagePostTag 查询文章引用标签分页
func PagePostTag(traceID string, req *blogModel.PostTagPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.PostTagTable.Page(postTagPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PagePostTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostTagGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToPostTagPageRes(list), total))
}

// GetPostTag 文章引用标签详情
func GetPostTag(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.PostTagTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPostTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostTagGetNG)
	}
	return baseModel.SuccessUnPop(blogModel.ToPostTagGetRes(&res))
}

// EditPostTag 编辑文章引用标签
func EditPostTag(traceID string, req *blogModel.PostTagEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.PostTagTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPostTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostTagGetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = blogDB.PostTagTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditPostTag %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkPostTagDBErr(err)
	}
	return baseModel.Success(constant.PostTagEditSS, true)
}

// DelPostTag 文章引用标签移除
func DelPostTag(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := blogDB.PostTagTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPostTag Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostTagGetNG)
	}
	//	// 文章引用标签禁止刪除
	//	if dbReq.Mark == constant.StatusLock {
	//		log.ErrorTF(traceID, "DelPostTag %d Fail . Can not Edit", dbReq.Id)
	//		return baseModel.Fail(constant.PostTagMarkNG)
	//	}
	// 物理删除
	err = blogDB.PostTagTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelPostTag %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.PostTagDelNG)
	}
	return baseModel.Success(constant.PostTagDelSS, true)
}
