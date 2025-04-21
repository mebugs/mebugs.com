import axios from 'axios'
import { ref } from 'vue'

export type banner = {
  id: number // 数据ID
  tag: string // 标签
  title: string // 标题
  url: string // 静态地址
  type: string
  summary: string // 摘要
  sourceShow: string
  source: Array<string> // 选择数组（资源选择）
}

export function bannerInit() {
  return ref<banner>({
    id: 0,
    title: '',
    tag: '',
    type: '',
    url: '',
    summary: '',
    sourceShow: '',
    source: []
  })
}

export function bannerAdd(req: banner) {
  return axios.post('/blog/banner/add', req)
}

export function bannerPage(req: any) {
  return axios.post('/blog/banner/page', req)
}

export function bannerGet(req: number) {
  return axios.post('/blog/banner/get', { id: req })
}

export function bannerEdit(req: banner) {
  return axios.post('/blog/banner/edit', req)
}

export function bannerDel(req: any) {
  return axios.post('/blog/banner/del', { id: req })
}
