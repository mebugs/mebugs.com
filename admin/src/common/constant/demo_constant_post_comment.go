package constant

/**
 * 请将下述常量迁移到 response_code.go
 * 相关编号通过 src/zoom/respCode/make_test.go 生成
 */

const (
	PostCommentAddSS    = "Sn00" // 文章评论创建成功
	PostCommentEditSS   = "Sn01" // 文章评论编辑成功
	PostCommentDelSS    = "Sn02" // 文章评论删除成功
	PostCommentGetNG    = "Fn00" // 文章评论查询失败
	PostCommentUniXxxNG = "Fn01" // 文章评论地址全局唯一
	PostCommentMarkNG   = "Fn02" // 内置文章评论禁止删除
	PostCommentDelNG    = "Fn03" // 文章评论删除

)
