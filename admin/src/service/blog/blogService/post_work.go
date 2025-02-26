package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkPostDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "xxx_uni") {
			// 唯一索引错误
			return baseModel.Fail(constant.PostUniXxxNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 分页查询对象封装
func postPageQuery(req *blogModel.PostPageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	if req.Title != "" {
		query.Like("title", req.Title)
	}
	if req.Url != "" {
		query.Like("url", req.Url)
	}
	if req.CategoryId != 0 {
		query.Eq("category_id", req.CategoryId)
	}
	if req.PostType != "" {
		query.Eq("post_type", req.PostType)
	}
	if req.Status != "" {
		query.Eq("status", req.Status)
	}
	// query.Eq("status", constant.StatusOpen)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}
