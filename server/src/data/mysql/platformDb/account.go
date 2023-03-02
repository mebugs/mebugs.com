package platformDb

import (
	"siteOl.com/stone/server/src/data/mysql/actuator"
	"time"
)

// Account 数据库表结构映射
type Account struct {
	ID         uint64     // 账号ID
	Account    string     // 账号
	Encryption string     // 密文密码
	SaltKey    string     // 盐值秘钥
	PwdExpTime *time.Time // 密码超期时间（修改后的90天）
	TenantId   uint64     // 租户ID
	Status     uint8      // 账号状态 1正常 2锁定 3封存
	Mark       uint8      // 变更标识 1可变更 2禁止变更（除密码）
	CreateAt   *time.Time // 创建时间
	UpdateAt   *time.Time // 更新时间
}

// FindOne 基于对象实施查询
func (a *Account) FindOne() (res *Account, err error) {
	res = new(Account)
	err = actuator.FindOneByObject(platformDb, a, res)
	return
}
