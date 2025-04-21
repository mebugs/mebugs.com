package blogModel

import (
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// IndexData 首页数据对象
type IndexData struct {
	Banners []*cacheModel.BannerCache `json:"banners"` // BANNER
	TagData
	CategoryData
	PostsHot  []*cacheModel.PostCache `json:"postsHot"`
	PostsView []*cacheModel.PostCache `json:"postsView"`
	PostsGood []*cacheModel.PostCache `json:"postsGood"`
	PostsNew  []*cacheModel.PostCache `json:"postsNew"`
}

// PageData 页面数据对象
type PageData struct {
	Banners []*cacheModel.BannerCache `json:"banners"` // BANNER
	TagData
	CategoryData
	PostsHot  []*cacheModel.PostCache `json:"postsHot"`
	PostsView []*cacheModel.PostCache `json:"postsView"`
	PostsGood []*cacheModel.PostCache `json:"postsGood"`
	PostsNew  []*cacheModel.PostCache `json:"postsNew"`
}

// CategoryData 分类数据对象
type CategoryData struct {
	Category []*cacheModel.CategoryCache `json:"category"`
}

// TagData 标签数据对象
type TagData struct {
	Tag []*cacheModel.TagCache `json:"tag"`
}

// TagsData 标签列表页响应
type TagsData struct {
	TagData
	Total int `json:"total"`
}

// PostsReq 查询文章的对象
type PostsReq struct {
	By    int    `json:"by"`    // 查询源：1 new 2 good 3 view 4 hot 5 all=new 6 category 7 tag 8 topic
	Value string `json:"value"` // 查询源：category tag topic 的值
	Page  int    `json:"page"`  // 查询页，每页固定15个
}

// PostsData 文章列表页响应
type PostsData struct {
	Null  bool                    `json:"null"`
	Main  any                     `json:"main"`
	Posts []*cacheModel.PostCache `json:"posts"`
	Total int                     `json:"total"`
}

// PostReq 查询文章的对象
type PostReq struct {
	Url string `json:"url"` // 文章URL
}

// PostData 文章响应
type PostData struct {
	Null     bool                      `json:"null"`
	Post     *cacheModel.PostCache     `json:"post"`
	PostMain *cacheModel.PostMainCache `json:"postMain"`
	PostMore []*cacheModel.PostCache   `json:"postMore"`
}

// PostCommReq 评论添加
type PostCommReq struct {
	User   PostUser `json:"user"`                            // PostUser 评论用户
	PostId uint64   `json:"postId"`                          // 文章ID
	Level  uint8    `json:"level"`                           // 评论等级，枚举：0_评论 1_回复
	Rid    uint64   `json:"rid"`                             // 回复的评论者ID（来自被回复的数据）
	Info   string   `json:"info" binding:"required,max=512"` // 评论内容（暂不支持HTML）
}

// PostUser 评论用户
type PostUser struct {
	ClientId string `json:"clientId" binding:"required,max=64"` // 客户端ID（前端生成，清除会丢失）
	SourceId uint64 `json:"sourceId"`                           // 头像资源ID（仅小程序提供）
	Name     string `json:"name" binding:"required,max=32"`     // 用户名
	Email    string `json:"email" binding:"max=128"`            // 邮箱
	Summary  string `json:"summary" binding:"max=128"`          // 简介
	Url      string `json:"url" binding:"max=128"`              // 三方网站（需要审核）
}

// CommentsReq 评论查询
type CommentsReq struct {
	PostId   uint64 `json:"postId"`   // 文章ID
	ClientId string `json:"clientId"` // 客户端ID（前端生成，清除会丢失）
}

// CommentsRes 评论响应
type CommentsRes struct {
	Total    int         `json:"total"`    // 响应评论
	Comments []*Comments `json:"comments"` // 评论树
}

// Comments 评论对象
type Comments struct {
	User     *CommentsUser `json:"user"`
	Id       uint64        `json:"id"`       // 当前评论ID
	Rid      uint64        `json:"rid"`      // 回复的评论者ID（来自被回复的数据）
	Info     string        `json:"info"`     // 评论内容（暂不支持HTML）
	Date     string        `json:"date"`     // 评论日期
	Comments []*Comments   `json:"comments"` // 评论树
}

// CommentsUser 评论用户
type CommentsUser struct {
	Id       uint64 `json:"id"`       // UID
	SourceId uint64 `json:"sourceId"` // 头像资源ID（仅小程序提供）
	Name     string `json:"name" `    // 用户名
	Summary  string `json:"summary" ` // 简介
	Url      string `json:"url"`      // 三方网站（需要审核）
}

var NormalUser = &CommentsUser{
	Id:       0,
	SourceId: 0,
	Name:     "查无此人",
	Summary:  "这个用户失联了...",
	Url:      "#",
}

func ToCommentsUser(user *blogDB.User) *CommentsUser {
	return &CommentsUser{
		Id:       user.Id,
		SourceId: user.SourceId,
		Name:     user.Name,
		Summary:  user.Summary,
		Url:      user.ThirdUrl,
	}
}

func ToComments(user *CommentsUser, com *blogDB.PostComment) *Comments {
	return &Comments{
		User:     user,
		Id:       com.Id,
		Rid:      com.Rid,
		Info:     com.Info,
		Date:     com.CreateAt.Format("2006-01-02"),
		Comments: make([]*Comments, 0),
	}
}
