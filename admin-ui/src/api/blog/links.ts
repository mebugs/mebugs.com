import axios from 'axios'
import { ref } from 'vue'

export type links = {
  id: number // 数据ID
  title: string // 标题
  url: string // 静态地址
  summary: string // 摘要
  sourceShow: string
  source: Array<string> // 选择数组（资源选择）
  status: string
}

export function linksInit() {
  return ref<links>({
    id: 0,
    title: '',
    url: '',
    summary: '',
    sourceShow: '',
    source: [],
    status: ''
  })
}

export function linksAdd(req: links) {
  return axios.post('/blog/links/add', req)
}

export function linksPage(req: any) {
  return axios.post('/blog/links/page', req)
}

export function linksGet(req: number) {
  return axios.post('/blog/links/get', { id: req })
}

export function linksEdit(req: links) {
  return axios.post('/blog/links/edit', req)
}

export function linksDel(req: any) {
  return axios.post('/blog/links/del', { id: req })
}
