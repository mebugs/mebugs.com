package platformDb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/config"
	"siteOl.com/stone/server/src/utils/logUtil"
)

var platformDb *gorm.DB

// InitPlatFromDb 初始化平台数据库
func InitPlatFromDb() {
	// 采用默认配置打开数据可
	db, err := gorm.Open(mysql.Open(config.RunConfig.MySQL.Platform), &gorm.Config{})
	if err != nil {
		logUtil.FatalF("Open PlatFromDb Fail . Err Is : %s", err)
		return
	}
	platformDb = db
	logUtil.InfoF("Init PlatFromDb Success . ")
}
