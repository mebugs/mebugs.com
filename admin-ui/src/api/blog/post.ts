import axios from 'axios'
import { ref } from 'vue'

export type post = {
  id: number // 数据ID
  title: string // 标题
  url: string // 静态地址
  summary: string // 摘要
  sourceShow: string
  source: Array<string> // 选择数组（资源选择）
  categoryId: Number | null // 分组
  topic: Array<string> // 专题
  page: string // 页面
  tagIds: Array<Number> // 关联标签
  status: string // 状态，枚举：0_发布 1_草稿
  pushAt: string // 计划发布日期
  md: string | undefined // markdown文本
  html: string | undefined // 编译的HTML
  toc: string | undefined // 编译的HTML
  sourceIds: Array<Number>
  postType: string
}

export function postInit(postType: string) {
  return ref<post>({
    id: 0,
    title: '',
    url: '',
    summary: '',
    sourceShow: '',
    source: [],
    categoryId: postType == '1' ? null : 0,
    topic: [],
    page: '',
    tagIds: [],
    status: '',
    pushAt: '',
    md: '',
    html: '',
    toc: '',
    sourceIds: [],
    postType: postType
  })
}

export function postAdd(req: post) {
  return axios.post('/blog/post/add', req)
}

export function postPage(req: any) {
  return axios.post('/blog/post/page', req)
}

export function postGet(req: number) {
  return axios.post('/blog/post/get', { id: req })
}

export function postEdit(req: post) {
  return axios.post('/blog/post/edit', req)
}

export function postDel(req: any) {
  return axios.post('/blog/post/del', { id: req })
}

export function postReset(req: any) {
  return axios.post('/blog/post/reset', { id: req })
}
