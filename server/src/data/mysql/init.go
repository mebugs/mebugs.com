package mysql

import "siteOl.com/stone/server/src/data/mysql/platDb"

// Init 初始化全部数据库
func Init() {
	platDb.InitPlatFromDb() // 平台基础数据库
}
