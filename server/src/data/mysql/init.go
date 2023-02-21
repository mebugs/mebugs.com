package mysql

import "siteOl.com/stone/server/src/data/mysql/platformDb"

// InitMySQL 初始化全部数据库
func InitMySQL() {
	platformDb.InitPlatFromDb()
}
