package constant

/**
 * 请将下述常量迁移到 response_code.go
 * 相关编号通过 src/zoom/respCode/make_test.go 生成
 */

const (
	PostMoreAddSS    = "Sn00" // 文章关联创建成功
	PostMoreEditSS   = "Sn01" // 文章关联编辑成功
	PostMoreDelSS    = "Sn02" // 文章关联删除成功
	PostMoreGetNG    = "Fn00" // 文章关联查询失败
	PostMoreUniXxxNG = "Fn01" // 文章关联地址全局唯一
	PostMoreMarkNG   = "Fn02" // 内置文章关联禁止删除
	PostMoreDelNG    = "Fn03" // 文章关联删除

)
