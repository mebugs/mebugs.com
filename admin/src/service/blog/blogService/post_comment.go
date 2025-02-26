package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddPostComment 创建文章评论
func AddPostComment(traceID string, req *blogModel.PostCommentAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := blogDB.PostCommentTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddPostComment Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkPostCommentDBErr(err)
	}
	return baseModel.Success(constant.PostCommentAddSS, true)
}

// PagePostComment 查询文章评论分页
func PagePostComment(traceID string, req *blogModel.PostCommentPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.PostCommentTable.Page(postCommentPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PagePostComment Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostCommentGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToPostCommentPageRes(list), total))
}

// GetPostComment 文章评论详情
func GetPostComment(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.PostCommentTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPostComment Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostCommentGetNG)
	}
	return baseModel.SuccessUnPop(blogModel.ToPostCommentGetRes(&res))
}

// EditPostComment 编辑文章评论
func EditPostComment(traceID string, req *blogModel.PostCommentEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.PostCommentTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPostComment Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostCommentGetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = blogDB.PostCommentTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditPostComment %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkPostCommentDBErr(err)
	}
	return baseModel.Success(constant.PostCommentEditSS, true)
}

// DelPostComment 文章评论移除
func DelPostComment(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := blogDB.PostCommentTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPostComment Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostCommentGetNG)
	}
	//	// 文章评论禁止刪除
	//	if dbReq.Mark == constant.StatusLock {
	//		log.ErrorTF(traceID, "DelPostComment %d Fail . Can not Edit", dbReq.Id)
	//		return baseModel.Fail(constant.PostCommentMarkNG)
	//	}
	// 物理删除
	err = blogDB.PostCommentTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelPostComment %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.PostCommentDelNG)
	}
	return baseModel.Success(constant.PostCommentDelSS, true)
}
