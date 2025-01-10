import axios from 'axios'
import { ref } from 'vue'

export type topic = {
  id: number // 数据ID
  title: string // 标题
  url: string // 静态地址
  summary: string // 摘要
  sourceShow: string
  source: Array<string> // 选择数组（资源选择）
  postIdSet: Array<any>
  posts: Array<any>
}

export function topicInit() {
  return ref<topic>({
    id: 0,
    title: '',
    url: '',
    summary: '',
    sourceShow: '',
    source: [],
    postIdSet: [],
    posts: []
  })
}

export function topicAdd(req: topic) {
  return axios.post('/blog/topic/add', req)
}

export function topicPage(req: any) {
  return axios.post('/blog/topic/page', req)
}

export function topicGet(req: number) {
  return axios.post('/blog/topic/get', { id: req })
}

export function topicEdit(req: topic) {
  return axios.post('/blog/topic/edit', req)
}

export function topicDel(req: any) {
  return axios.post('/blog/topic/del', { id: req })
}
