package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// PostMore 文章关联
type PostMore struct {
	Id   uint64 `json:"id"`   // 数据ID
	Toc  string `json:"toc"`  // 文章导航
	Md   string `json:"md"`   // MD源文件
	Html string `json:"html"` // HTML源文件
}

// PostMoreTable 文章关联泛型造器
var PostMoreTable actuator.Table[PostMore]

// DataBase 实现指定数据库
func (t PostMore) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t PostMore) TableName() string {
	return "post_more"
}
