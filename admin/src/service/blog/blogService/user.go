package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// AddUser 创建极简用户信息
func AddUser(traceID string, req *blogModel.UserAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := blogDB.UserTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddUser Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkUserDBErr(err)
	}
	return baseModel.Success(constant.UserAddSS, true)
}

// PageUser 查询极简用户信息分页
func PageUser(traceID string, req *blogModel.UserPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.UserTable.Page(userPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageUser Fail . Err Is : %v", err)
		return baseModel.Fail(constant.UserGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToUserPageRes(list), total))
}

// GetUser 极简用户信息详情
func GetUser(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.UserTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetUser Fail . Err Is : %v", err)
		return baseModel.Fail(constant.UserGetNG)
	}
	return baseModel.SuccessUnPop(blogModel.ToUserGetRes(&res))
}

// EditUser 编辑极简用户信息
func EditUser(traceID string, req *blogModel.UserEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.UserTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetUser Fail . Err Is : %v", err)
		return baseModel.Fail(constant.UserGetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = blogDB.UserTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditUser %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkUserDBErr(err)
	}
	return baseModel.Success(constant.UserEditSS, true)
}

// DelUser 极简用户信息移除
func DelUser(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := blogDB.UserTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetUser Fail . Err Is : %v", err)
		return baseModel.Fail(constant.UserGetNG)
	}
	//	// 极简用户信息禁止刪除
	//	if dbReq.Mark == constant.StatusLock {
	//		log.ErrorTF(traceID, "DelUser %d Fail . Can not Edit", dbReq.Id)
	//		return baseModel.Fail(constant.UserMarkNG)
	//	}
	// 物理删除
	err = blogDB.UserTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelUser %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.UserDelNG)
	}
	return baseModel.Success(constant.UserDelSS, true)
}
