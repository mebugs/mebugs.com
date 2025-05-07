package blogModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// PostUserReq 极简用户信息 通用请求，创建&编辑可复用的字段
type PostUserReq struct {
	Id       uint64 `json:"id"`                                        // 用户ID
	Name     string `json:"name" binding:"max=32" example:"demo"`      // 用户名
	Email    string `json:"email" example:"0"`                         // 邮箱
	Summary  string `json:"summary" binding:"max=64" example:"demo"`   // 简介
	ThirdUrl string `json:"thirdUrl" binding:"max=255" example:"demo"` // 三方网站（需要审核）
	WaitUrl  string `json:"waitUrl" binding:"max=255" example:"demo"`  // 待审核三方URL
	SourceId uint64 `json:"sourceId" example:"0"`                      // 头像资源ID（仅小程序提供）
}

// PostCommentDoReq 文章评论 通用请求，创建&编辑可复用的字段
type PostCommentDoReq struct {
	Uid    uint64       `json:"uid" example:"0"`                         // 评论者ID
	User   *PostUserReq `json:"user" binding:"required" example:"demo"`  // 评论新对象更新User
	Info   string       `json:"info" binding:"max=512" example:"demo"`   // 评论内容（暂不支持HTML）
	ReInfo string       `json:"reInfo" binding:"max=512" example:"demo"` // 回复内容如果有
	// Status string       `json:"status" binding:"required,oneof='0' '2'" example:"0"` // 状态，枚举：0_正常 2_封存
}

// PostCommentEditReq 文章评论 编辑请求，酌情从通用中摘出部分字段
type PostCommentEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	PostCommentDoReq
}

// ToDbReq 文章评论 更新转数据库
func (r *PostCommentEditReq) ToDbReq(d *blogDB.PostComment) {
	now := time.Now()
	d.Uid = r.Uid
	d.Info = r.Info
	d.Status = constant.StatusOpen
	d.UpdateAt = &now
}

// ToReCommDbReq 文章评论 更新转数据库
func ToReCommDbReq(r *blogDB.PostComment, rInfo string) *blogDB.PostComment {
	return &blogDB.PostComment{
		PostId:   r.PostId,
		Level:    1, // 回复必然为1
		Uid:      1,
		Rid:      r.Id,
		Info:     rInfo,
		Status:   constant.StatusOpen,
		CreateAt: r.UpdateAt,
		UpdateAt: r.UpdateAt,
	}
}

// PostCommentGetRes 文章评论 详情响应
type PostCommentGetRes struct {
	Id       uint64               `json:"id" example:"1"`          // 数据ID
	PostId   uint64               `json:"postId" example:"0"`      // 文章ID
	PostName string               `json:"postName" example:"demo"` // 文章名称
	PostDesc string               `json:"postDesc" example:"demo"` // 文章名称
	Level    uint8                `json:"level" example:"0"`       // 评论等级，枚举：0_评论 1_回复
	Uid      uint64               `json:"uid" example:"0"`         // 评论者ID
	User     *PostUserReq         `json:"user" example:"demo"`     // 评论新对象更新User
	Info     string               `json:"info" example:"demo"`     // 评论内容（暂不支持HTML）
	Status   string               `json:"status" example:"0"`      // 状态，枚举：0_正常 1_锁定 2_封存
	RecList  []*PostCommentGetRes `json:"recList"`                 // 关联相关的回复
}

// PostCommentPageReq 文章评论 分页请求，根据实际业务替换分页条件字段
type PostCommentPageReq struct {
	baseModel.PageReq
}

// PostCommentPageRes 文章评论 分页响应，酌情从详情摘出部分字段
type PostCommentPageRes struct {
	PostCommentGetRes
}

// ToPostCommentGetRes 文章评论 数据库转为详情响应
func ToPostCommentGetRes(r *blogDB.PostComment, user *blogDB.User, post *blogDB.Post) *PostCommentGetRes {
	return &PostCommentGetRes{
		Id:       r.Id,
		PostId:   r.PostId,
		PostName: post.Title,
		PostDesc: post.Summary,
		Level:    r.Level,
		Uid:      r.Uid,
		Info:     r.Info,
		Status:   r.Status,
		User: &PostUserReq{
			Id:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			Summary:  user.Summary,
			ThirdUrl: user.ThirdUrl,
			WaitUrl:  user.WaitUrl,
			SourceId: user.SourceId,
		},
	}
}
