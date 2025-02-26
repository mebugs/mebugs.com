package blogModel

import (
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// PostCommentDoReq 文章评论 通用请求，创建&编辑可复用的字段
type PostCommentDoReq struct {
	PostId   uint64     `json:"postId" example:"0"`                                      // 文章ID
	Level    uint8      `json:"level" binding:"oneof='0' '1'" example:"0"`               // 评论等级，枚举：0_评论 1_回复
	Uid      uint64     `json:"uid" example:"0"`                                         // 评论者ID
	Rid      uint64     `json:"rid" example:"0"`                                         // 回复的评论者ID（来自被回复的数据）
	Info     string     `json:"info" binding:"max=512" example:"demo"`                   // 评论内容（暂不支持HTML）
	Status   string     `json:"status" binding:"required,oneof='0' '1' '2'" example:"0"` // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"`                                                // 创建时间
	UpdateAt *time.Time `json:"updateAt"`                                                // 更新时间
}

// PostCommentAddReq 文章评论 创建请求，酌情从通用中摘出部分字段
type PostCommentAddReq struct {
	PostCommentDoReq
}

// PostCommentEditReq 文章评论 编辑请求，酌情从通用中摘出部分字段
type PostCommentEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	PostCommentDoReq
}

// ToDbReq 文章评论 创建转数据库
func (r *PostCommentAddReq) ToDbReq() *blogDB.PostComment {
	now := time.Now()
	return &blogDB.PostComment{
		Id:       0,
		PostId:   r.PostId,
		Level:    r.Level,
		Uid:      r.Uid,
		Rid:      r.Rid,
		Info:     r.Info,
		Status:   r.Status,
		CreateAt: &now,
		UpdateAt: &now,
	}
}

// ToDbReq 文章评论 更新转数据库
func (r *PostCommentEditReq) ToDbReq(d *blogDB.PostComment) {
	now := time.Now()
	d.PostId = r.PostId
	d.Level = r.Level
	d.Uid = r.Uid
	d.Rid = r.Rid
	d.Info = r.Info
	d.Status = r.Status
	d.CreateAt = &now
	d.UpdateAt = &now
}

// PostCommentGetRes 文章评论 详情响应
type PostCommentGetRes struct {
	Id       uint64     `json:"id" example:"1"`      // 数据ID
	PostId   uint64     `json:"postId" example:"0"`  // 文章ID
	Level    uint8      `json:"level" example:"0"`   // 评论等级，枚举：0_评论 1_回复
	Uid      uint64     `json:"uid" example:"0"`     // 评论者ID
	Rid      uint64     `json:"rid" example:"0"`     // 回复的评论者ID（来自被回复的数据）
	Info     string     `json:"info" example:"demo"` // 评论内容（暂不支持HTML）
	Status   string     `json:"status" example:"0"`  // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"`            // 创建时间
	UpdateAt *time.Time `json:"updateAt"`            // 更新时间
}

// PostCommentPageReq 文章评论 分页请求，根据实际业务替换分页条件字段
type PostCommentPageReq struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
	baseModel.PageReq
}

// PostCommentPageRes 文章评论 分页响应，酌情从详情摘出部分字段
type PostCommentPageRes struct {
	PostCommentGetRes
}

// ToPostCommentGetRes 文章评论 数据库转为详情响应
func ToPostCommentGetRes(r *blogDB.PostComment) *PostCommentGetRes {
	return &PostCommentGetRes{
		Id:       r.Id,
		PostId:   r.PostId,
		Level:    r.Level,
		Uid:      r.Uid,
		Rid:      r.Rid,
		Info:     r.Info,
		Status:   r.Status,
		CreateAt: r.CreateAt,
		UpdateAt: r.UpdateAt,
	}
}

// ToPostCommentPageRes 文章评论 数据库转分页响应
func ToPostCommentPageRes(list []*blogDB.PostComment) []*PostCommentPageRes {
	res := make([]*PostCommentPageRes, len(list))
	for i, r := range list {
		res[i] = &PostCommentPageRes{
			PostCommentGetRes: *ToPostCommentGetRes(r),
		}
	}
	return res
}
