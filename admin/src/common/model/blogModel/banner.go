package blogModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// BannerDoReq Banner通用请求，创建&编辑可复用的字段
type BannerDoReq struct {
	Title   string `json:"title" binding:"max=32" example:"demo"`   // 名称
	Tag     string `json:"tag" binding:"max=12" example:"demo"`     // 标签
	Url     string `json:"url" binding:"max=128" example:"demo"`    // 分类地址
	Type    string `json:"type" binding:"required" example:"1"`     // 类型 0 Banner 1 Page
	Summary string `json:"summary" binding:"max=64" example:"demo"` // 简介
}

// BannerAddReq Banner创建请求，酌情从通用中摘出部分字段
type BannerAddReq struct {
	Source []string `json:"source"  binding:"required" example:"0"` // 主图资源列表
	BannerDoReq
}

// BannerEditReq Banner编辑请求，酌情从通用中摘出部分字段
type BannerEditReq struct {
	Id     uint64   `json:"id" binding:"required" example:"1"` // 数据ID
	Source []string `json:"source" example:"0"`                // 主图资源列表
	BannerDoReq
}

// ToDbReq Banner创建转数据库
func (r *BannerAddReq) ToDbReq() *blogDB.Banner {
	now := time.Now()
	return &blogDB.Banner{
		Id:       0,
		Title:    r.Title,
		Url:      r.Url,
		Tag:      r.Tag,
		Type:     r.Type,
		Summary:  r.Summary,
		Status:   constant.StatusOpen,
		CreateAt: &now,
		UpdateAt: &now,
	}
}

// ToDbReq Banner更新转数据库
func (r *BannerEditReq) ToDbReq(d *blogDB.Banner) {
	now := time.Now()
	d.Title = r.Title
	d.Url = r.Url
	d.Tag = r.Tag
	d.Summary = r.Summary
	d.UpdateAt = &now
}

// BannerGetRes Banner详情响应
type BannerGetRes struct {
	Id         uint64 `json:"id" example:"1"`                      // 数据ID
	Title      string `json:"title" example:"demo"`                // 名称
	Tag        string `json:"tag" example:"demo"`                  // 标签
	Type       string `json:"type" binding:"required" example:"1"` // 类型 0 Banner 1 Page
	Url        string `json:"url" example:"demo"`                  // 分类地址
	Summary    string `json:"summary" example:"demo"`              // 简介
	SourceShow string `json:"sourceShow" example:"/xxx.jpg"`       // 资源图片地址
	Status     string `json:"status" example:"0"`                  // 状态，枚举：0_正常 1_锁定 2_封存
}

// BannerPageReq Banner分页请求，根据实际业务替换分页条件字段
type BannerPageReq struct {
	Title string `json:"title" example:"demo"` // 名称
	Url   string `json:"url" example:"demo"`   // 分类地址
	Type  string `json:"type" example:"1"`     // 类型 0 Banner 1 Page
	baseModel.PageReq
}

// BannerPageRes Banner分页响应，酌情从详情摘出部分字段
type BannerPageRes struct {
	BannerGetRes
}

// ToBannerGetRes Banner数据库转为详情响应
func ToBannerGetRes(r *blogDB.Banner, sourceShow string) *BannerGetRes {
	return &BannerGetRes{
		Id:         r.Id,
		Title:      r.Title,
		Tag:        r.Tag,
		Url:        r.Url,
		Type:       r.Type,
		Summary:    r.Summary,
		SourceShow: sourceShow,
		Status:     r.Status,
	}
}

// ToBannerPageRes Banner数据库转分页响应
func ToBannerPageRes(list []*blogDB.Banner, sourceMap map[uint64]string) []*BannerPageRes {
	res := make([]*BannerPageRes, len(list))
	for i, r := range list {
		res[i] = &BannerPageRes{
			BannerGetRes: *ToBannerGetRes(r, sourceMap[r.SourceId]),
		}
	}
	return res
}
