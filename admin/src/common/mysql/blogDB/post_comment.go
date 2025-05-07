package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// PostComment 文章评论
type PostComment struct {
	Id       uint64     `json:"id"`       // 数据ID
	PostId   uint64     `json:"postId"`   // 文章ID
	Level    uint8      `json:"level"`    // 评论等级，枚举：0_评论 1_回复
	Uid      uint64     `json:"uid"`      // 评论者ID
	Rid      uint64     `json:"rid"`      // 回复的评论者ID（来自被回复的数据）
	Info     string     `json:"info"`     // 评论内容（暂不支持HTML）
	Status   string     `json:"status"`   // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}

// PostCommentTable 文章评论泛型造器
var PostCommentTable actuator.Table[PostComment]

// DataBase 实现指定数据库
func (t PostComment) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t PostComment) TableName() string {
	return "post_comment"
}

func (t PostComment) Comments(postId uint64) (res []*PostComment, err error) {
	r := blogDb.Table(t.TableName()).Where("post_id = ?", postId).Order("`level`,create_at").Find(&res)
	err = r.Error
	return
}

func (t PostComment) FindByRid(rid uint64) (res []*PostComment, err error) {
	r := blogDb.Table(t.TableName()).Where("rid = ?", rid).Order("create_at").Find(&res)
	err = r.Error
	return
}
