package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// Banner Banner
type Banner struct {
	Id       uint64     `json:"id"`       // 数据ID
	Tag      string     `json:"tag"`      // 标签（仅用于展示）
	Title    string     `json:"title"`    // 名称
	Url      string     `json:"url"`      // 链接地址
	Type     string     `json:"type"`     // 类型 0 banner 1 page
	Summary  string     `json:"summary"`  // 简介
	SourceId uint64     `json:"sourceId"` // 主图资源ID
	Status   string     `json:"status"`   // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}

// BannerTable Banner泛型造器
var BannerTable actuator.Table[Banner]

// DataBase 实现指定数据库
func (t Banner) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t Banner) TableName() string {
	return "banner"
}

// GetBanners 获取Banners
func (t Banner) GetBanners() (res []*Banner, err error) {
	r := blogDb.Table(t.TableName()).Order("update_at").Find(&res)
	err = r.Error
	return
}
