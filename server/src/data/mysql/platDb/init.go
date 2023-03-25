package platDb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/config"
	"siteOl.com/stone/server/src/utils/log"
	"time"
)

var platformDb *gorm.DB

// InitPlatFromDb 初始化平台数据库
func InitPlatFromDb() {
	// 采用默认配置打开数据可
	db, err := gorm.Open(mysql.Open(config.RunConfig.MySQL.Platform), &gorm.Config{})
	if err != nil {
		log.FatalF("Open PlatDb Fail . Err Is : %s", err)
		return
	}
	platformDb = db
	log.InfoF("Init PlatDb Success . ")
}

// Common 平台通用信息体
type Common struct {
	Status   uint8      // 状态 1正常 2锁定 3封存
	CreateAt *time.Time // 创建时间
	UpdateAt *time.Time // 更新时间
}
