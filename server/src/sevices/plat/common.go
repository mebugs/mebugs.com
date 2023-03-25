package plat

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
)

/**
 *
 * 平台公共方法
 *
 *
 * @author 米虫@mebugs.com
 * @since 2023-03-08
 */

// GetLoginUser 从上下文获取登录用户授权机构体
func GetLoginUser(c *gin.Context) *platDb.AuthUser {
	obj, ok := c.Get(constant.AuthUser)
	if ok {
		authUser := &platDb.AuthUser{}
		authUser = obj.(*platDb.AuthUser)
		return authUser
	}
	return nil
}
