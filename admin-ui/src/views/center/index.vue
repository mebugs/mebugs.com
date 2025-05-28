<template>
  <div class="container cindx">
    <a-card class="ccard" v-for="(i, idx) in names" :title="i.title" hoverable :loading="load">
      <template #extra>
        <a-link @click="to(i.link)">前往</a-link>
      </template>
      <span style="font-size: 20px">{{ counts[idx] }}</span>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { centerIndex } from '@/api/blog/center'
import useLoad from '@/utils/hooks/load'
import { useRouter } from 'vue-router'
// 路由对象
const router = useRouter()
const to = (url: string) => {
  router.push(url)
}
const names = [
  { title: '文章总量', link: '/blog/post' },
  { title: '文章草稿', link: '/blog/post' },
  { title: '分类总量', link: '/blog/category' },
  { title: '标签总量', link: '/blog/tag' },
  { title: '评论总量', link: '/blog/comment' },
  { title: '待审评论', link: '/blog/comment' },
  { title: '友链总量', link: '/blog/links' },
  { title: '待审友链', link: '/blog/links' }
]
// 加载中变量
const { load, setLoad } = useLoad(true)
const counts = ref(<any>[0, 0, 0, 0, 0, 0, 0, 0])
// 分页检索
async function init() {
  setLoad(true)
  try {
    const res = await centerIndex()
    counts.value = res.data
  } catch (e) {
  } finally {
    setLoad(false)
  }
}
init()
</script>
<script lang="ts">
export default {
  name: 'CenterIndex'
}
</script>

<style scoped lang="less"></style>
