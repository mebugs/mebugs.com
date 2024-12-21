package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// Category 文章分类
type Category struct {
	Id       uint64     `json:"id"`       // 数据ID
	Title    string     `json:"title"`    // 名称
	Url      string     `json:"url"`      // 分类地址
	Summary  string     `json:"summary"`  // 简介
	SourceId uint64     `json:"sourceId"` // 主图资源ID
	Status   string     `json:"status"`   // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}

// CategoryTable 文章分类泛型造器
var CategoryTable actuator.Table[Category]

// DataBase 实现指定数据库
func (t Category) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t Category) TableName() string {
	return "category"
}
