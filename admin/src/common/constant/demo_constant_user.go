package constant

/**
 * 请将下述常量迁移到 response_code.go
 * 相关编号通过 src/zoom/respCode/make_test.go 生成
 */

const (
	UserAddSS    = "Sn00" // 极简用户信息创建成功
	UserEditSS   = "Sn01" // 极简用户信息编辑成功
	UserDelSS    = "Sn02" // 极简用户信息删除成功
	UserGetNG    = "Fn00" // 极简用户信息查询失败
	UserUniXxxNG = "Fn01" // 极简用户信息地址全局唯一
	UserMarkNG   = "Fn02" // 内置极简用户信息禁止删除
	UserDelNG    = "Fn03" // 极简用户信息删除

)
