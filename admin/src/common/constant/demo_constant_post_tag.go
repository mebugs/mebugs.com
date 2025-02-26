package constant

/**
 * 请将下述常量迁移到 response_code.go
 * 相关编号通过 src/zoom/respCode/make_test.go 生成
 */

const (
	PostTagAddSS    = "Sn00" // 文章引用标签创建成功
	PostTagEditSS   = "Sn01" // 文章引用标签编辑成功
	PostTagDelSS    = "Sn02" // 文章引用标签删除成功
	PostTagGetNG    = "Fn00" // 文章引用标签查询失败
	PostTagUniXxxNG = "Fn01" // 文章引用标签地址全局唯一
	PostTagMarkNG   = "Fn02" // 内置文章引用标签禁止删除
	PostTagDelNG    = "Fn03" // 文章引用标签删除

)
