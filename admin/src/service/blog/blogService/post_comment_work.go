package blogService

import (
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/actuator"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 分页查询对象封装
func postCommentPageQuery(req *blogModel.PostCommentPageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	query.Ne("uid", 1)
	query.Desc("status")
	query.Desc("create_at")
	query.LimitByPage(req.Current, req.PageSize)
	return
}
