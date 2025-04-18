package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// User 极简用户信息
type User struct {
	Id       uint64     `json:"id"`       // 数据ID
	ClientId string     `json:"clientId"` // 客户端ID（前端生成，清除会丢失）
	Name     string     `json:"name"`     // 用户名
	Email    string     `json:"email"`    // 邮箱
	Code     string     `json:"code"`     // 用户编码（允许访问用户）默认=ID
	Summary  string     `json:"summary"`  // 简介
	ThirdUrl string     `json:"thirdUrl"` // 三方网站（需要审核）
	WaitUrl  string     `json:"waitUrl"`  // 待审核三方URL
	SourceId uint64     `json:"sourceId"` // 头像资源ID（仅小程序提供）
	OpenId   string     `json:"openId"`   // 小程序账号绑定登录（允许授权转移）
	Status   string     `json:"status"`   // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}

// UserTable 极简用户信息泛型造器
var UserTable actuator.Table[User]

// DataBase 实现指定数据库
func (t User) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t User) TableName() string {
	return "user"
}
