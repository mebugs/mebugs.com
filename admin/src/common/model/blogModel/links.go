package blogModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// LinksScanReq 友情链接扫描
type LinksScanReq struct {
	Url string `json:"url" binding:"required,max=128" example:"demo"` // 分类地址
}

// LinksDoReq 友情链接 通用请求，创建&编辑可复用的字段
type LinksDoReq struct {
	Title   string `json:"title" binding:"max=32" example:"demo"`    // 名称
	Url     string `json:"url" binding:"max=128" example:"demo"`     // 分类地址
	Summary string `json:"summary" binding:"max=256" example:"demo"` // 简介
}

// LinksAddClientReq 友情链接 客户端创建请求
type LinksAddClientReq struct {
	ClientId string `json:"clientId" binding:"required,max=64"` // 客户端ID（前端生成，清除会丢失）
	LinksAddReq
}

// LinksAddReq 友情链接 创建请求，酌情从通用中摘出部分字段
type LinksAddReq struct {
	Source []string `json:"source"  binding:"required" example:"0"` // 主图资源列表
	LinksDoReq
}

// LinksEditReq 友情链接 编辑请求，酌情从通用中摘出部分字段
type LinksEditReq struct {
	Id     uint64   `json:"id" binding:"required" example:"1"`     // 数据ID
	Source []string `json:"source" example:"0"`                    // 主图资源列表
	Status string   `json:"status" binding:"required" example:"0"` // 链接状态
	LinksDoReq
}

// ToDbReq 友情链接 创建转数据库
func (r *LinksAddReq) ToDbReq() *blogDB.Links {
	now := time.Now()
	return &blogDB.Links{
		Id:       0,
		Title:    r.Title,
		Url:      r.Url,
		Summary:  r.Summary,
		Status:   constant.StatusOpen, // 创建为成功态
		CreateAt: &now,
		UpdateAt: &now,
	}
}

// ToDbReq 友情链接 更新转数据库
func (r *LinksEditReq) ToDbReq(d *blogDB.Links) {
	now := time.Now()
	d.Title = r.Title
	d.Url = r.Url
	d.Summary = r.Summary
	d.Status = r.Status
	d.UpdateAt = &now
}

// LinksGetRes 友情链接 详情响应
type LinksGetRes struct {
	Id         uint64 `json:"id" example:"1"`                // 数据ID
	Title      string `json:"title" example:"demo"`          // 名称
	Url        string `json:"url" example:"demo"`            // 分类地址
	Summary    string `json:"summary" example:"demo"`        // 简介
	SourceShow string `json:"sourceShow" example:"/xxx.jpg"` // 资源图片地址
	Status     string `json:"status" example:"0"`            // 状态，枚举：0_正常 1_锁定 2_封存
}

// LinksPageReq 友情链接 分页请求，根据实际业务替换分页条件字段
type LinksPageReq struct {
	Title string `json:"title" example:"demo"` // 名称
	Url   string `json:"url" example:"demo"`   // 分类地址
	baseModel.PageReq
}

// LinksPageRes 友情链接 分页响应，酌情从详情摘出部分字段
type LinksPageRes struct {
	LinksGetRes
}

// ToLinksGetRes 友情链接 数据库转为详情响应
func ToLinksGetRes(r *blogDB.Links, sourceShow string) *LinksGetRes {
	return &LinksGetRes{
		Id:         r.Id,
		Title:      r.Title,
		Url:        r.Url,
		Summary:    r.Summary,
		SourceShow: sourceShow,
		Status:     r.Status,
	}
}

// ToLinksPageRes 友情链接 数据库转分页响应
func ToLinksPageRes(list []*blogDB.Links, sourceMap map[uint64]string) []*LinksPageRes {
	res := make([]*LinksPageRes, len(list))
	for i, r := range list {
		res[i] = &LinksPageRes{
			LinksGetRes: *ToLinksGetRes(r, sourceMap[r.SourceId]),
		}
	}
	return res
}
