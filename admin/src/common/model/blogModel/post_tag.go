package blogModel

import (
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// PostTagDoReq 文章引用标签 通用请求，创建&编辑可复用的字段
type PostTagDoReq struct {
	PostId uint64 `json:"postId" example:"0"` // 文章ID
	TagId  uint64 `json:"tagId" example:"0"`  // 标签ID
}

// PostTagAddReq 文章引用标签 创建请求，酌情从通用中摘出部分字段
type PostTagAddReq struct {
	PostTagDoReq
}

// PostTagEditReq 文章引用标签 编辑请求，酌情从通用中摘出部分字段
type PostTagEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	PostTagDoReq
}

// ToDbReq 文章引用标签 创建转数据库
func (r *PostTagAddReq) ToDbReq() *blogDB.PostTag {
	return &blogDB.PostTag{
		Id:     0,
		PostId: r.PostId,
		TagId:  r.TagId,
	}
}

// ToDbReq 文章引用标签 更新转数据库
func (r *PostTagEditReq) ToDbReq(d *blogDB.PostTag) {
	d.PostId = r.PostId
	d.TagId = r.TagId
}

// PostTagGetRes 文章引用标签 详情响应
type PostTagGetRes struct {
	Id     uint64 `json:"id" example:"1"`     // 数据ID
	PostId uint64 `json:"postId" example:"0"` // 文章ID
	TagId  uint64 `json:"tagId" example:"0"`  // 标签ID
}

// PostTagPageReq 文章引用标签 分页请求，根据实际业务替换分页条件字段
type PostTagPageReq struct {
	Title string `json:"title" example:"demo"` // 名称
	Url   string `json:"url" example:"demo"`   // 分类地址
	baseModel.PageReq
}

// PostTagPageRes 文章引用标签 分页响应，酌情从详情摘出部分字段
type PostTagPageRes struct {
	PostTagGetRes
}

// ToPostTagGetRes 文章引用标签 数据库转为详情响应
func ToPostTagGetRes(r *blogDB.PostTag) *PostTagGetRes {
	return &PostTagGetRes{
		Id:     r.Id,
		PostId: r.PostId,
		TagId:  r.TagId,
	}
}

// ToPostTagPageRes 文章引用标签 数据库转分页响应
func ToPostTagPageRes(list []*blogDB.PostTag) []*PostTagPageRes {
	res := make([]*PostTagPageRes, len(list))
	for i, r := range list {
		res[i] = &PostTagPageRes{
			PostTagGetRes: *ToPostTagGetRes(r),
		}
	}
	return res
}
