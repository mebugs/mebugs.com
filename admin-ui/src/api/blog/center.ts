import axios from 'axios'

export function centerIndex() {
  return axios.post('/blog/center/index', {})
}
