package constant

/**
 * 请将下述常量迁移到 response_code.go
 * 相关编号通过 src/zoom/respCode/make_test.go 生成
 */

const (
	PostSourceAddSS    = "Sn00" // 文章引用资源创建成功
	PostSourceEditSS   = "Sn01" // 文章引用资源编辑成功
	PostSourceDelSS    = "Sn02" // 文章引用资源删除成功
	PostSourceGetNG    = "Fn00" // 文章引用资源查询失败
	PostSourceUniXxxNG = "Fn01" // 文章引用资源地址全局唯一
	PostSourceMarkNG   = "Fn02" // 内置文章引用资源禁止删除
	PostSourceDelNG    = "Fn03" // 文章引用资源删除

)
