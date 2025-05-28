package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"siteol.com/smart/src/common/mysql/blogDB"
	"siteol.com/smart/src/service"
)

func CenterIndex(c *gin.Context) {
	// 读取统计
	totalS := make([]int64, 8)
	// 文章总量 草稿总量
	normalQuery := actuator.InitQuery()
	totalS[0], _ = blogDB.PostTable.CountByQuery(normalQuery)
	query0 := actuator.InitQuery()
	query0.Eq("status", "0")
	totalS[1], _ = blogDB.PostTable.CountByQuery(query0)
	// 分类总量 TAG总量
	totalS[2], _ = blogDB.CategoryTable.CountByQuery(normalQuery)
	totalS[3], _ = blogDB.TagTable.CountByQuery(normalQuery)
	// 评论总量 待回复总量
	totalS[4], _ = blogDB.PostCommentTable.CountByQuery(normalQuery)
	query1 := actuator.InitQuery()
	query1.Eq("status", "1")
	totalS[5], _ = blogDB.PostCommentTable.CountByQuery(query1)
	// 友链总量 待审核友链总量
	totalS[6], _ = blogDB.LinksTable.CountByQuery(normalQuery)
	totalS[7], _ = blogDB.LinksTable.CountByQuery(query1)
	service.JsonRes(c, baseModel.SuccessUnPop(totalS))
}
