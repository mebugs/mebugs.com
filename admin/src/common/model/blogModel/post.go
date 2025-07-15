package blogModel

import (
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// PostListRes 文章关联用
type PostListRes struct {
	Id         uint64 `json:"id" example:"1"`         // 数据ID
	Title      string `json:"title" example:"demo"`   // 标题
	CategoryId uint64 `json:"CategoryId" example:"0"` // 分组，文章可用
	Url        string `json:"url" example:"demo"`     // 文章地址
}

// ToPostListRes 文章关联列表
func ToPostListRes(list []*blogDB.Post) []*PostListRes {
	res := make([]*PostListRes, len(list))
	for i, r := range list {
		res[i] = &PostListRes{
			Id:         r.Id,
			Title:      r.Title,
			CategoryId: r.CategoryId,
			Url:        r.Url,
		}
	}
	return res
}

// PostDoReq 文章或页面 通用请求，创建&编辑可复用的字段
type PostDoReq struct {
	Title      string     `json:"title" binding:"max=128" example:"demo"`                  // 标题
	Url        string     `json:"url" binding:"max=32" example:"demo"`                     // 文章地址
	Summary    string     `json:"summary" binding:"max=512" example:"demo"`                // 摘要
	CategoryId uint64     `json:"categoryId" example:"0"`                                  // 分组，文章可用
	TagIds     []uint64   `json:"tagIds"`                                                  // 文章分类
	SourceIds  []uint64   `json:"sourceIds"`                                               // 文章关联资源
	PushAt     string     `json:"pushAt" example:"2222-22-22"`                             // 发布时间，标准时间格式（时间固定为08:00）
	Status     string     `json:"status" binding:"required,oneof='0' '1' '2'" example:"0"` // 状态，枚举：0_草稿 1_发布 2_锁定
	PostType   string     `json:"postType" example:"0"`                                    // 文章类型，枚举：0_页面 1_文章
	CreateAt   *time.Time `json:"createAt"`                                                // 创建时间
	UpdateAt   *time.Time `json:"updateAt"`                                                // 更新时间
	PostMoreDoReq
}

// PostAddReq 文章或页面 创建请求，酌情从通用中摘出部分字段
type PostAddReq struct {
	PostDoReq
	Source []string `json:"source"  binding:"required" example:"0"` // 主图资源列表（主图）
}

// PostEditReq 文章或页面 编辑请求，酌情从通用中摘出部分字段
type PostEditReq struct {
	Id     uint64   `json:"id" binding:"required" example:"1"` // 数据ID
	Source []string `json:"source" example:"0"`                // 主图资源列表（主图）
	PostDoReq
}

// ToDbReq 文章或页面 创建转数据库
func (r *PostAddReq) ToDbReq() *blogDB.Post {
	now := time.Now()
	// PUSH 的定时任务每天扫描1次，固定8点开始发布
	pushAt := &now
	// 如果指定了发布时间
	if r.PushAt != "" {
		pushTime, err := time.Parse("2006-01-02", r.PushAt)
		if err == nil {
			pushAt = &pushTime
		}
	}
	return &blogDB.Post{
		Id:         0,
		Title:      r.Title,
		Url:        r.Url,
		Summary:    r.Summary,
		CategoryId: r.CategoryId,
		PostType:   r.PostType,
		PushAt:     pushAt,
		Status:     r.Status,
		CreateAt:   &now,
		UpdateAt:   &now,
	}
}

// ToDbReq 文章或页面 更新转数据库
func (r *PostEditReq) ToDbReq(d *blogDB.Post) {
	now := time.Now()
	// PUSH 的定时任务每天扫描1次，固定8点开始发布
	pushAt := &now
	// 如果指定了发布时间
	if r.PushAt != "" {
		pushTime, err := time.Parse("2006-01-02", r.PushAt)
		if err == nil {
			pushAt = &pushTime
		}
	}
	d.Title = r.Title
	d.Url = r.Url
	d.Summary = r.Summary
	d.CategoryId = r.CategoryId
	d.PushAt = pushAt
	d.Status = r.Status
	d.UpdateAt = &now
}

// PostGetRes 文章或页面 详情响应
type PostGetRes struct {
	Id         uint64 `json:"id" example:"1"`         // 数据ID
	Title      string `json:"title" example:"demo"`   // 标题
	Url        string `json:"url" example:"demo"`     // 文章地址
	Summary    string `json:"summary" example:"demo"` // 摘要
	SourceId   uint64 `json:"sourceId" example:"0"`   // 主图资源ID
	SourceShow string `json:"sourceShow" example:"0"` // 图片信息
	PostType   string `json:"postType" example:"0"`   // 文章类型，枚举：0_页面 1_文章
	CategoryId uint64 `json:"categoryId" example:"0"` // 分组，文章可用
	Status     string `json:"status" example:"0"`     // 状态，枚举：0_正常 1_锁定 2_封存
}

// PostPageReq 文章或页面 分页请求，根据实际业务替换分页条件字段
type PostPageReq struct {
	Title      string `json:"title" example:"demo"`   // 标题
	Url        string `json:"url" example:"demo"`     // 文章地址
	CategoryId uint64 `json:"categoryId" example:"0"` // 分组，文章可用
	Status     string `json:"status" example:"0"`     // 状态，枚举：0_正常 1_锁定 2_封存
	PostType   string `json:"postType" example:"1"`   // 文章类型，枚举：0_页面 1_文章
	baseModel.PageReq
}

// PostAllRes 文章详情页完整响应
type PostAllRes struct {
	PostGetRes
	Toc       string   `json:"toc" example:"demo"`          // 文章导航
	Md        string   `json:"md" example:"demo"`           // MD源文件
	Html      string   `json:"html" example:"demo"`         // HTML源文件
	TagIds    []uint64 `json:"tagIds"`                      // 文章分类
	SourceIds []uint64 `json:"sourceIds"`                   // 文章关联资源
	PushAt    string   `json:"pushAt" example:"2222-22-22"` // 发布时间，标准时间格式（时间固定为08:00）
}

// PostPageRes 文章或页面 分页响应，酌情从详情摘出部分字段
type PostPageRes struct {
	PostGetRes
}

// ToPostGetRes 文章或页面 数据库转为详情响应
func ToPostGetRes(r *blogDB.Post, sourceShow string) *PostGetRes {
	return &PostGetRes{
		Id:         r.Id,
		Title:      r.Title,
		Url:        r.Url,
		Summary:    r.Summary,
		SourceId:   r.SourceId,
		SourceShow: sourceShow,
		PostType:   r.PostType,
		CategoryId: r.CategoryId,
		Status:     r.Status,
	}
}

// ToPostPageRes 文章或页面 数据库转分页响应
func ToPostPageRes(list []*blogDB.Post, sourceMap map[uint64]string) []*PostPageRes {
	res := make([]*PostPageRes, len(list))
	for i, r := range list {
		res[i] = &PostPageRes{
			PostGetRes: *ToPostGetRes(r, sourceMap[r.SourceId]),
		}
	}
	return res
}

// ToPostAllRes 查询完整对象
func ToPostAllRes(p *blogDB.Post, p2 *blogDB.PostMore, tagIds []uint64, url string) *PostAllRes {
	return &PostAllRes{
		PostGetRes: *ToPostGetRes(p, url),
		Toc:        p2.Toc,
		Md:         p2.Md,
		Html:       p2.Html,
		TagIds:     tagIds,
		SourceIds:  nil,
		PushAt:     p.PushAt.Format("2006-01-02"),
	}
}
