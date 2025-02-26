package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// PostTag 文章引用标签
type PostTag struct {
	Id     uint64 `json:"id"`     // 数据ID
	PostId uint64 `json:"postId"` // 文章ID
	TagId  uint64 `json:"tagId"`  // 标签ID
}

// PostTagTable 文章引用标签泛型造器
var PostTagTable actuator.Table[PostTag]

// DataBase 实现指定数据库
func (t PostTag) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t PostTag) TableName() string {
	return "post_tag"
}

// DeleteByPostId 根据文章ID移除标签关联
func (t PostTag) DeleteByPostId(postId uint64) (err error) {
	r := blogDb.Where("post_id = ?", postId).Delete(&t)
	err = r.Error
	return
}

// GetTagIds 获取文章对应的标签ID
func (t PostTag) GetTagIds(postId uint64) (res []uint64, err error) {
	r := blogDb.Table(t.TableName()).Distinct("tag_id").Where("post_id = ?", postId).Find(&res)
	err = r.Error
	return
}
