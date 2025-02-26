package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"strings"
)

// 解析数据库错误
func checkPostMoreDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "xxx_uni") {
			// 唯一索引错误
			return baseModel.Fail(constant.PostMoreUniXxxNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}
