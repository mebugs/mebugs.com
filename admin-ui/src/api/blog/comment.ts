import axios from 'axios'
import { ref } from 'vue'

export type postComment = {
  id: number // 数据ID
  postId: number // 文章ID
  postName: string // 文章名称
  postDesc: string // 文章名称
  level: number // 评论等级，枚举：0_评论 1_回复
  uid: number // 评论者ID
  user: any // 用户对象
  info: string // 评论内容（暂不支持HTML）
  status: string // 状态，枚举：0_：0_正常 1_锁定 2_封存
  recList: Array<any>
  reInfo: string // 回复
}

export function postCommentInit() {
  return ref<postComment>({
    id: 0, // 数据ID
    postId: 0, // 文章ID
    postName: '', // 文章名称
    postDesc: '', // 文章名称
    level: 0, // 评论等级，枚举：0_评论 1_回复
    uid: 0, // 评论者ID
    user: {}, // 用户对象
    info: '', // 评论内容（暂不支持HTML）
    status: '1', // 状态，枚举：0_：0_正常 1_草稿 2_封存
    recList: [],
    reInfo: '' // 回复
  })
}

export function postCommentPage(req: any) {
  return axios.post('/blog/postComment/page', req)
}

export function postCommentGet(req: number) {
  return axios.post('/blog/postComment/get', { id: req })
}

export function postCommentEdit(req: postComment) {
  return axios.post('/blog/postComment/edit', req)
}
