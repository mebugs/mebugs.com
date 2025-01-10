package blogModel

import (
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// PostTopicDoReq 文章引用主题 通用请求，创建&编辑可复用的字段
type PostTopicDoReq struct {
	PostId  uint64 `json:"postId" example:"0"`  // 文章ID
	TopicId uint64 `json:"topicId" example:"0"` // 专题ID
	Sort    uint16 `json:"sort" example:"0"`    // 排序序号
}

// PostTopicAddReq 文章引用主题 创建请求，酌情从通用中摘出部分字段
type PostTopicAddReq struct {
	PostTopicDoReq
}

// PostTopicEditReq 文章引用主题 编辑请求，酌情从通用中摘出部分字段
type PostTopicEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	PostTopicDoReq
}

// ToDbReq 文章引用主题 创建转数据库
func (r *PostTopicAddReq) ToDbReq() *blogDB.PostTopic {
	return &blogDB.PostTopic{
		Id:      0,
		PostId:  r.PostId,
		TopicId: r.TopicId,
		Sort:    r.Sort,
	}
}

// ToDbReq 文章引用主题 更新转数据库
func (r *PostTopicEditReq) ToDbReq(d *blogDB.PostTopic) {
	d.PostId = r.PostId
	d.TopicId = r.TopicId
	d.Sort = r.Sort
}

// PostTopicGetRes 文章引用主题 详情响应
type PostTopicGetRes struct {
	Id      uint64 `json:"id" example:"1"`      // 数据ID
	PostId  uint64 `json:"postId" example:"0"`  // 文章ID
	TopicId uint64 `json:"topicId" example:"0"` // 专题ID
	Sort    uint16 `json:"sort" example:"0"`    // 排序序号
}

// PostTopicPageReq 文章引用主题 分页请求，根据实际业务替换分页条件字段
type PostTopicPageReq struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
	baseModel.PageReq
}

// PostTopicPageRes 文章引用主题 分页响应，酌情从详情摘出部分字段
type PostTopicPageRes struct {
	PostTopicGetRes
}

// ToPostTopicGetRes 文章引用主题 数据库转为详情响应
func ToPostTopicGetRes(r *blogDB.PostTopic) *PostTopicGetRes {
	return &PostTopicGetRes{
		Id:      r.Id,
		PostId:  r.PostId,
		TopicId: r.TopicId,
		Sort:    r.Sort,
	}
}

// ToPostTopicPageRes 文章引用主题 数据库转分页响应
func ToPostTopicPageRes(list []*blogDB.PostTopic) []*PostTopicPageRes {
	res := make([]*PostTopicPageRes, len(list))
	for i, r := range list {
		res[i] = &PostTopicPageRes{
			PostTopicGetRes: *ToPostTopicGetRes(r),
		}
	}
	return res
}
