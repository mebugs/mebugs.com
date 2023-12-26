package dbPlat

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// Router 路由表
type Router struct {
	ID           uint64 // 默认数据ID
	Name         string // 路由名称，用于界面展示，与权限关联
	Url          string // 路由地址，后端访问URL 后端不再URL中携带参数，统一Post处理内容
	Type         uint8  // 免授权路由 0 授权 1 免授权
	ReqLogPrint  uint8  // 请求日志打印 0 打印 1 不打印
	ReqLogSecure string // 请求日志脱敏字段，逗号分隔
	ResLogPrint  uint8  // 响应日志打印 0 打印 1 不打印
	ResLogSecure string // 响应日志脱敏字段，逗号分隔
	Common
}

// RouterTable 路由泛型造器
var RouterTable actuator.Table[Router]

// DataBase 实现指定数据库
func (t Router) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Router) TableName() string {
	return "router"
}
