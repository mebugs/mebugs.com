package blogService

import (
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// syncPostSources 刷新文档和资源关系 editFlag:true表是更细
func syncPostSources(traceID string, postId uint64, sourceIds []uint64, editFlag bool) (err error) {
	if editFlag {
		// 移除当前权限的路由
		err = blogDB.PostSourceTable.Executor().DeleteByPostId(postId)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByPostId By %d Fail . Err Is : %v", postId, err)
			return
		}
	}
	// 重新插入路由关系
	if len(sourceIds) > 0 {
		postSources := make([]blogDB.PostSource, len(sourceIds))
		for i, item := range sourceIds {
			postSources[i] = blogDB.PostSource{
				Id:       0,
				PostId:   postId,
				SourceId: item,
			}
		}
		err = blogDB.PostSourceTable.InsertBatch(&postSources)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchPostSources By topicId %d Fail . Err Is : %v", postId, err)

		}
	}
	return
}
