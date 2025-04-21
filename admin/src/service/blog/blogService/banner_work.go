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
func checkBannerDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "url_uni") {
			// 唯一索引错误
			return baseModel.Fail(constant.BannerUniXxxNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// BannerPageQuery 分页查询对象封装
func BannerPageQuery(req *blogModel.BannerPageReq) (query *actuator.Query) {
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
	if req.Type != "" {
		query.Eq("type", req.Type)
	}
	query.Eq("status", constant.StatusOpen)
	query.Desc("update_at")
	query.LimitByPage(req.Current, req.PageSize)
	return
}
