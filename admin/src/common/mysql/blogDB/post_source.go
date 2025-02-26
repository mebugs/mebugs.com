package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// PostSource 文章引用资源
type PostSource struct {
	Id       uint64 `json:"id"`       // 数据ID
	PostId   uint64 `json:"postId"`   // 文章ID
	SourceId uint64 `json:"sourceId"` // 资源ID
}

// PostSourceTable 文章引用资源泛型造器
var PostSourceTable actuator.Table[PostSource]

// DataBase 实现指定数据库
func (t PostSource) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t PostSource) TableName() string {
	return "post_source"
}

// DeleteByPostId 根据文章ID移除标签关联
func (t PostSource) DeleteByPostId(postId uint64) (err error) {
	r := blogDb.Where("post_id = ?", postId).Delete(&t)
	err = r.Error
	return
}
