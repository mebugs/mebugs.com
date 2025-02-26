package blogModel

import (
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// PostSourceDoReq 文章引用资源 通用请求，创建&编辑可复用的字段
type PostSourceDoReq struct {
	PostId   uint64 `json:"postId" example:"0"`   // 文章ID
	SourceId uint64 `json:"sourceId" example:"0"` // 资源ID
}

// PostSourceAddReq 文章引用资源 创建请求，酌情从通用中摘出部分字段
type PostSourceAddReq struct {
	PostSourceDoReq
}

// PostSourceEditReq 文章引用资源 编辑请求，酌情从通用中摘出部分字段
type PostSourceEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	PostSourceDoReq
}

// ToDbReq 文章引用资源 创建转数据库
func (r *PostSourceAddReq) ToDbReq() *blogDB.PostSource {
	return &blogDB.PostSource{
		Id:       0,
		PostId:   r.PostId,
		SourceId: r.SourceId,
	}
}

// ToDbReq 文章引用资源 更新转数据库
func (r *PostSourceEditReq) ToDbReq(d *blogDB.PostSource) {
	d.PostId = r.PostId
	d.SourceId = r.SourceId
}

// PostSourceGetRes 文章引用资源 详情响应
type PostSourceGetRes struct {
	Id       uint64 `json:"id" example:"1"`       // 数据ID
	PostId   uint64 `json:"postId" example:"0"`   // 文章ID
	SourceId uint64 `json:"sourceId" example:"0"` // 资源ID
}

// PostSourcePageReq 文章引用资源 分页请求，根据实际业务替换分页条件字段
type PostSourcePageReq struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
	baseModel.PageReq
}

// PostSourcePageRes 文章引用资源 分页响应，酌情从详情摘出部分字段
type PostSourcePageRes struct {
	PostSourceGetRes
}

// ToPostSourceGetRes 文章引用资源 数据库转为详情响应
func ToPostSourceGetRes(r *blogDB.PostSource) *PostSourceGetRes {
	return &PostSourceGetRes{
		Id:       r.Id,
		PostId:   r.PostId,
		SourceId: r.SourceId,
	}
}

// ToPostSourcePageRes 文章引用资源 数据库转分页响应
func ToPostSourcePageRes(list []*blogDB.PostSource) []*PostSourcePageRes {
	res := make([]*PostSourcePageRes, len(list))
	for i, r := range list {
		res[i] = &PostSourcePageRes{
			PostSourceGetRes: *ToPostSourceGetRes(r),
		}
	}
	return res
}
