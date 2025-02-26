package constant

/**
 * 请将下述常量迁移到 response_code.go
 * 相关编号通过 src/zoom/respCode/make_test.go 生成
 */

const (
	PostAddSS     = "Sn00" // 文章或页面创建成功
	PostEditSS    = "Sn01" // 文章或页面编辑成功
	PostDelSS     = "Sn02" // 文章或页面删除成功
	PostGetNG     = "Fn00" // 文章或页面查询失败
	PostAddNoScNG = "Fn00" // 文章或页面必须有主图
	PostTagNG     = "FNN"  // 文章关联标签更新失败
	PostSourceNG  = "FNN"  // 文章关联资源更新失败
	PostMoreNG    = "FNN"  // 文章关联详情更新失败
	PostUniXxxNG  = "Fn01" // 文章或页面地址全局唯一
	PostMarkNG    = "Fn02" // 内置文章或页面禁止删除
	PostDelNG     = "Fn03" // 文章或页面删除

)
