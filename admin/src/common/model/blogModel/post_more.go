package blogModel

import (
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// PostMoreDoReq 文章关联 通用请求，创建&编辑可复用的字段
type PostMoreDoReq struct {
	Toc  string `json:"toc" example:"demo"`  // 文章导航
	Md   string `json:"md" example:"demo"`   // MD源文件
	Html string `json:"html" example:"demo"` // HTML源文件
}

// PostMoreAddReq 文章关联 创建请求，酌情从通用中摘出部分字段
type PostMoreAddReq struct {
	PostMoreDoReq
}

// PostMoreEditReq 文章关联 编辑请求，酌情从通用中摘出部分字段
type PostMoreEditReq struct {
	PostMoreDoReq
}

// ToDbReq 文章关联 创建转数据库
func (r *PostMoreAddReq) ToDbReq() *blogDB.PostMore {
	return &blogDB.PostMore{
		Id:   0,
		Md:   r.Md,
		Html: r.Html,
	}
}

// ToDbReq 文章关联 更新转数据库
func (r *PostMoreEditReq) ToDbReq(d *blogDB.PostMore) {
	d.Md = r.Md
	d.Html = r.Html
}

// PostMoreGetRes 文章关联 详情响应
type PostMoreGetRes struct {
	Id   uint64 `json:"id" example:"1"`      // 数据ID
	Md   string `json:"md" example:"demo"`   // MD源文件
	Html string `json:"html" example:"demo"` // HTML源文件
}

// PostMorePageReq 文章关联 分页请求，根据实际业务替换分页条件字段
type PostMorePageReq struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
	baseModel.PageReq
}

// PostMorePageRes 文章关联 分页响应，酌情从详情摘出部分字段
type PostMorePageRes struct {
	PostMoreGetRes
}

// ToPostMoreGetRes 文章关联 数据库转为详情响应
func ToPostMoreGetRes(r *blogDB.PostMore) *PostMoreGetRes {
	return &PostMoreGetRes{
		Id:   r.Id,
		Md:   r.Md,
		Html: r.Html,
	}
}

// ToPostMorePageRes 文章关联 数据库转分页响应
func ToPostMorePageRes(list []*blogDB.PostMore) []*PostMorePageRes {
	res := make([]*PostMorePageRes, len(list))
	for i, r := range list {
		res[i] = &PostMorePageRes{
			PostMoreGetRes: *ToPostMoreGetRes(r),
		}
	}
	return res
}
