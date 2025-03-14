package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// PostTopic 文章引用主题
type PostTopic struct {
	Id      uint64 `json:"id"`      // 数据ID
	PostId  uint64 `json:"postId"`  // 文章ID
	TopicId uint64 `json:"topicId"` // 专题ID
	Sort    uint16 `json:"sort"`    // 排序序号
}

// PostTopicTable 文章引用主题泛型造器
var PostTopicTable actuator.Table[PostTopic]

// DataBase 实现指定数据库
func (t PostTopic) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t PostTopic) TableName() string {
	return "post_topic"
}

// DeleteByTopicId 根据主题ID移除文章
func (t PostTopic) DeleteByTopicId(topicId uint64) (err error) {
	r := blogDb.Where("topic_id = ?", topicId).Delete(&t)
	err = r.Error
	return
}

// GetPostIds 获取主题对应的文章ID
func (t PostTopic) GetPostIds(topicId uint64) (res []uint64, err error) {
	r := blogDb.Table(t.TableName()).Distinct("post_id").Where("topic_id = ?", topicId).Order("sort").Find(&res)
	err = r.Error
	return
}

// GetTopicIds 获取主题ID
func (t PostTopic) GetTopicIds(postId uint64) (res []uint64, err error) {
	r := blogDb.Table(t.TableName()).Distinct("topic_id").Where("post_id = ?", postId).Order("sort").Find(&res)
	err = r.Error
	return
}
