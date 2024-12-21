package blogModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// TagDoReq 标签 通用请求，创建&编辑可复用的字段
type TagDoReq struct {
	Title   string `json:"title" binding:"max=32" example:"demo"`   // 名称
	Url     string `json:"url" binding:"max=32" example:"demo"`     // 标签地址
	Summary string `json:"summary" binding:"max=64" example:"demo"` // 简介
}

// TagAddReq 标签 创建请求，酌情从通用中摘出部分字段
type TagAddReq struct {
	Source []string `json:"source"  binding:"required" example:"0"` // 主图资源列表
	TagDoReq
}

// TagEditReq 标签 编辑请求，酌情从通用中摘出部分字段
type TagEditReq struct {
	Id     uint64   `json:"id" binding:"required" example:"1"` // 数据ID
	Source []string `json:"source" example:"0"`                // 主图资源列表
	TagDoReq
}

// ToDbReq 标签 创建转数据库
func (r *TagAddReq) ToDbReq() *blogDB.Tag {
	now := time.Now()
	return &blogDB.Tag{
		Id:       0,
		Title:    r.Title,
		Url:      r.Url,
		Summary:  r.Summary,
		Status:   constant.StatusOpen,
		CreateAt: &now,
		UpdateAt: &now,
	}
}

// ToDbReq 标签 更新转数据库
func (r *TagEditReq) ToDbReq(d *blogDB.Tag) {
	now := time.Now()
	d.Title = r.Title
	d.Url = r.Url
	d.Summary = r.Summary
	d.UpdateAt = &now
}

// TagGetRes 标签 详情响应
type TagGetRes struct {
	Id         uint64 `json:"id" example:"1"`                // 数据ID
	Title      string `json:"title" example:"demo"`          // 名称
	Url        string `json:"url" example:"demo"`            // 标签地址
	Summary    string `json:"summary" example:"demo"`        // 简介
	SourceShow string `json:"sourceShow" example:"/xxx.jpg"` // 资源图片地址
}

// TagPageReq 标签 分页请求，根据实际业务替换分页条件字段
type TagPageReq struct {
	Title string `json:"title" example:"demo"` // 名称
	Url   string `json:"url" example:"demo"`   // 分类地址
	baseModel.PageReq
}

// TagPageRes 标签 分页响应，酌情从详情摘出部分字段
type TagPageRes struct {
	TagGetRes
}

// ToTagGetRes 标签 数据库转为详情响应
func ToTagGetRes(r *blogDB.Tag, sourceShow string) *TagGetRes {
	return &TagGetRes{
		Id:         r.Id,
		Title:      r.Title,
		Url:        r.Url,
		Summary:    r.Summary,
		SourceShow: sourceShow,
	}
}

// ToTagPageRes 标签 数据库转分页响应
func ToTagPageRes(list []*blogDB.Tag, sourceMap map[uint64]string) []*TagPageRes {
	res := make([]*TagPageRes, len(list))
	for i, r := range list {
		res[i] = &TagPageRes{
			TagGetRes: *ToTagGetRes(r, sourceMap[r.SourceId]),
		}
	}
	return res
}
