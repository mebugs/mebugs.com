import axios from 'axios'
import { ref } from 'vue'

export type source = {
  id: number // 数据ID
  name: string // 标题
  img: string // 图片地址
  fileType: string // 文件类型
  fileStrArray: Array<string> // 图片资源
}

export function sourceInit() {
  return ref<source>({
    id: 0,
    name: '',
    img: '',
    fileType: '',
    fileStrArray: []
  })
}

export function sourceAdd(req: source) {
  return axios.post('/blog/source/add', req)
}

export function sourcePage(req: any) {
  return axios.post('/blog/source/page', req)
}

export function sourceGet(req: number) {
  return axios.post('/blog/source/get', { id: req })
}

export function sourceEdit(req: source) {
  return axios.post('/blog/source/edit', req)
}

export function sourceDel() {
  return axios.post('/blog/source/del')
}
