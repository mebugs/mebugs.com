import axios from 'axios'
import { ref } from 'vue'

export type tag = {
  id: number // 数据ID
  title: string // 标题
  url: string // 静态地址
  summary: string // 摘要
  sourceShow: string
  source: Array<string> // 选择数组（资源选择）
}

export function tagInit() {
  return ref<tag>({
    id: 0,
    title: '',
    url: '',
    summary: '',
    sourceShow: '',
    source: []
  })
}

export function tagAdd(req: tag) {
  return axios.post('/blog/tag/add', req)
}

export function tagPage(req: any) {
  return axios.post('/blog/tag/page', req)
}

export function tagGet(req: number) {
  return axios.post('/blog/tag/get', { id: req })
}

export function tagEdit(req: tag) {
  return axios.post('/blog/tag/edit', req)
}

export function tagDel(req: any) {
  return axios.post('/blog/tag/del', { id: req })
}

export function tagList() {
  return axios.post('/blog/tag/list')
}
