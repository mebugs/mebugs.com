package blogService

import (
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
)

// getTopicPosts 获取主题关联的文章
func getTopicPosts(traceID string, topicId uint64) (postIds []uint64, post []*blogModel.PostListRes, err error) {
	postIds, err = blogDB.PostTopic{}.GetPostIds(topicId)
	if err != nil {
		log.ErrorTF(traceID, "GetPostIdsByTopicId By %d Fail . Err Is : %v", topicId, err)
		return
	}
	list := make([]*blogDB.Post, 0)
	// 关联查询文章字段
	if len(postIds) > 0 {
		listR, errR := blogDB.PostTable.GetByIds(postIds)
		if errR != nil {
			err = errR
			log.WarnTF(traceID, "GetPostsByTopicId By %v Fail . Err Is : %v", postIds, err)
			return
		}
		listRMap := make(map[uint64]*blogDB.Post, 0)
		for _, li := range listR {
			listRMap[li.Id] = li
		}
		// 排序
		for _, id := range postIds {
			li, ok := listRMap[id]
			if ok {
				list = append(list, li)
			}
		}
	}
	post = blogModel.ToPostListRes(list)
	return
}

// syncTopicPosts 刷新主题和文档关系 editFlag:true表是更细
func syncTopicPosts(traceID string, topicId uint64, postSet []*baseModel.SortReq, editFlag bool) (err error) {
	if editFlag {
		// 移除当前权限的路由
		err = blogDB.PostTopicTable.Executor().DeleteByTopicId(topicId)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByTopicId By %d Fail . Err Is : %v", topicId, err)
			return
		}
	}
	// 重新插入路由关系
	if len(postSet) > 0 {
		topicPosts := make([]blogDB.PostTopic, len(postSet))
		for i, item := range postSet {
			topicPosts[i] = blogDB.PostTopic{
				PostId:  item.ID,
				TopicId: topicId,
				Sort:    item.Sort,
			}
		}
		err = blogDB.PostTopicTable.InsertBatch(&topicPosts)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchPostTopic By topicId %d Fail . Err Is : %v", topicId, err)

		}
	}
	return
}
