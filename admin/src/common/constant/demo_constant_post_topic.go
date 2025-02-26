package constant

/**
 * 请将下述常量迁移到 response_code.go
 * 相关编号通过 src/zoom/respCode/make_test.go 生成
 */

const (
	PostTopicAddSS    = "Sn00" // 文章引用主题创建成功
	PostTopicEditSS   = "Sn01" // 文章引用主题编辑成功
	PostTopicDelSS    = "Sn02" // 文章引用主题删除成功
	PostTopicGetNG    = "Fn00" // 文章引用主题查询失败
	PostTopicUniXxxNG = "Fn01" // 文章引用主题地址全局唯一
	PostTopicMarkNG   = "Fn02" // 内置文章引用主题禁止删除
	PostTopicDelNG    = "Fn03" // 文章引用主题删除

)
