package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// PagePostComment 查询文章评论分页
func PagePostComment(traceID string, req *blogModel.PostCommentPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.PostCommentTable.Page(postCommentPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PagePostComment Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostCommentGetNG)
	}
	// 关联文章和回复
	postMap := map[uint64]*blogDB.Post{}
	postMap[0] = &blogDB.Post{}
	userMap := map[uint64]*blogDB.User{}
	commList := make([]*blogModel.PostCommentPageRes, len(list))
	for i, com := range list {
		// 先查文章
		post, ok := postMap[com.PostId]
		if !ok {
			post1, err := blogDB.PostTable.GetOneById(com.PostId)
			if err != nil {
				post = &blogDB.Post{}
			} else {
				post = &post1
			}
			postMap[com.PostId] = post
		}
		// 再查用户
		user, ok := userMap[com.Uid]
		if !ok {
			user1, err := blogDB.UserTable.GetOneById(com.Uid)
			if err != nil {
				user = &blogDB.User{}
			} else {
				user = &user1
			}
			userMap[com.Uid] = user
		}
		commList[i] = &blogModel.PostCommentPageRes{
			PostCommentGetRes: *blogModel.ToPostCommentGetRes(com, user, post),
		}
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(commList, total))
}

// GetPostComment 文章评论详情
func GetPostComment(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.PostCommentTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPostComment Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostCommentGetNG)
	}
	// 关联文章和回复
	postMap := map[uint64]*blogDB.Post{}
	postMap[0] = &blogDB.Post{}
	userMap := map[uint64]*blogDB.User{}
	// 先查文章
	post, err := blogDB.PostTable.GetOneById(res.PostId)
	if err != nil {
		post = blogDB.Post{}
	}
	postMap[res.PostId] = &post
	// 再查用户
	user, err := blogDB.UserTable.GetOneById(res.Uid)
	if err != nil {
		user = blogDB.User{}
	}
	userMap[res.Uid] = &user
	list, err := blogDB.PostCommentTable.Executor().FindByRid(res.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPostComment ByRid Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PostCommentGetNG)
	}

	resBody := blogModel.ToPostCommentGetRes(&res, &user, &post)
	commList := make([]*blogModel.PostCommentGetRes, len(list))
	for i, com := range list {
		// 先查文章
		post, ok := postMap[com.PostId]
		if !ok {
			post1, err := blogDB.PostTable.GetOneById(com.PostId)
			if err != nil {
				post = &blogDB.Post{}
			} else {
				post = &post1
			}
			postMap[com.PostId] = post
		}
		// 再查用户
		user, ok := userMap[com.Uid]
		if !ok {
			user1, err := blogDB.UserTable.GetOneById(com.Uid)
			if err != nil {
				user = &blogDB.User{}
			} else {
				user = &user1
			}
			userMap[com.Uid] = user
		}
		commList[i] = blogModel.ToPostCommentGetRes(com, user, post)
	}
	resBody.RecList = commList
	return baseModel.SuccessUnPop(resBody)
}

// EditPostComment 编辑文章评论
func EditPostComment(traceID string, req *blogModel.PostCommentEditReq) *baseModel.ResBody {
	// 更新用户
	user, err := blogDB.UserTable.GetOneById(req.User.Id)
	if err != nil {
		log.ErrorTF(traceID, "EditPostComment GetUser Fail . Err Is : %v", err)
		return baseModel.SysErr
	}
	now := time.Now()
	user.UpdateAt = &now
	user.Name = req.User.Name
	user.Summary = req.User.Summary
	user.ThirdUrl = req.User.WaitUrl // URL点击通过或拒绝，拒绝变为#
	err = blogDB.UserTable.UpdateOne(&user)
	if err != nil {
		log.ErrorTF(traceID, "EditPostComment UpDate User Failed . Err Is %v", err)
		return baseModel.FailWithMsg("评论用户初始化失败")
	}
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
		return baseModel.SysErr
	}
	if req.ReInfo != "" {
		rOne := blogModel.ToReCommDbReq(&dbReq, req.ReInfo)
		err = blogDB.PostCommentTable.InsertOne(rOne)
		if err != nil {
			log.ErrorTF(traceID, "EditPostComment ReComm Fail . Err Is : %v", err)
			return baseModel.SysErr
		}
	}
	return baseModel.Success(constant.PostCommentEditSS, true)
}
