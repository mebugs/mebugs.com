package blogModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// TopicDoReq 文章专题 通用请求，创建&编辑可复用的字段
type TopicDoReq struct {
	Title     string               `json:"title" binding:"max=32" example:"demo"`   // 名称
	Url       string               `json:"url" binding:"max=32" example:"demo"`     // 专题地址
	Summary   string               `json:"summary" binding:"max=64" example:"demo"` // 简介
	PostIdSet []*baseModel.SortReq `json:"postIdSet" example:"1,2,3"`               // 关联文章
}

// TopicAddReq 文章专题 创建请求，酌情从通用中摘出部分字段
type TopicAddReq struct {
	Source []string `json:"source"  binding:"required" example:"0"` // 主图资源列表
	TopicDoReq
}

// TopicEditReq 文章专题 编辑请求，酌情从通用中摘出部分字段
type TopicEditReq struct {
	Id     uint64   `json:"id" binding:"required" example:"1"` // 数据ID
	Source []string `json:"source" example:"0"`                // 主图资源列表
	TopicDoReq
}

// ToDbReq 文章专题 创建转数据库
func (r *TopicAddReq) ToDbReq() *blogDB.Topic {
	now := time.Now()
	return &blogDB.Topic{
		Id:       0,
		Title:    r.Title,
		Url:      r.Url,
		Summary:  r.Summary,
		Status:   constant.StatusOpen,
		CreateAt: &now,
		UpdateAt: &now,
	}
}

// ToDbReq 文章专题 更新转数据库
func (r *TopicEditReq) ToDbReq(d *blogDB.Topic) {
	now := time.Now()
	d.Title = r.Title
	d.Url = r.Url
	d.Summary = r.Summary
	d.UpdateAt = &now
}

// TopicGetRes 文章专题 详情响应
type TopicGetRes struct {
	Id         uint64         `json:"id" example:"1"`                // 数据ID
	Title      string         `json:"title" example:"demo"`          // 名称
	Url        string         `json:"url" example:"demo"`            // 专题地址
	Summary    string         `json:"summary" example:"demo"`        // 简介
	SourceShow string         `json:"sourceShow" example:"/xxx.jpg"` // 资源图片地址
	PostIds    []uint64       `json:"postIds" example:"1"`           // 文章ID列表
	Posts      []*PostListRes `json:"posts" `                        // 关联文章
}

// TopicPageReq 文章专题 分页请求，根据实际业务替换分页条件字段
type TopicPageReq struct {
	Title string `json:"title" example:"demo"` // 名称
	Url   string `json:"url" example:"demo"`   // 分类地址
	baseModel.PageReq
}

// TopicPageRes 文章专题 分页响应，酌情从详情摘出部分字段
type TopicPageRes struct {
	TopicGetRes
}

// ToTopicGetRes 文章专题 数据库转为详情响应
func ToTopicGetRes(r *blogDB.Topic, sourceShow string) *TopicGetRes {
	return &TopicGetRes{
		Id:         r.Id,
		Title:      r.Title,
		Url:        r.Url,
		Summary:    r.Summary,
		SourceShow: sourceShow,
	}
}

// ToTopicPageRes 文章专题 数据库转分页响应
func ToTopicPageRes(list []*blogDB.Topic, sourceMap map[uint64]string) []*TopicPageRes {
	res := make([]*TopicPageRes, len(list))
	for i, r := range list {
		res[i] = &TopicPageRes{
			TopicGetRes: *ToTopicGetRes(r, sourceMap[r.SourceId]),
		}
	}
	return res
}
