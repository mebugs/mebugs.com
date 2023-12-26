package mysql

import "siteol.com/smart/src/common/mysql/dbPlat"

// Init 初始化全部数据库
func Init() {
	dbPlat.InitPlatFromDb()
}
