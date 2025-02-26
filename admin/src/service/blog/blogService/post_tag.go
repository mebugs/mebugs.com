package blogService

import (
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// syncPostTags 刷新文档和标签关系 editFlag:true表是更细
func syncPostTags(traceID string, postId uint64, tagIds []uint64, editFlag bool) (err error) {
	if editFlag {
		// 移除当前权限的路由
		err = blogDB.PostTagTable.Executor().DeleteByPostId(postId)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByPostId By %d Fail . Err Is : %v", postId, err)
			return
		}
	}
	// 重新插入路由关系
	if len(tagIds) > 0 {
		postTags := make([]blogDB.PostTag, len(tagIds))
		for i, item := range tagIds {
			postTags[i] = blogDB.PostTag{
				Id:     0,
				PostId: postId,
				TagId:  item,
			}
		}
		err = blogDB.PostTagTable.InsertBatch(&postTags)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchPostTags By topicId %d Fail . Err Is : %v", postId, err)

		}
	}
	return
}

// getPostTags 获取文章标签IDS
func getPostTagIds(traceID string, postId uint64) (tagIds []uint64, err error) {
	tagIds, err = blogDB.PostTag{}.GetTagIds(postId)
	if err != nil {
		log.ErrorTF(traceID, "GetPostTagIds By %d Fail . Err Is : %v", postId, err)
	}
	return
}
