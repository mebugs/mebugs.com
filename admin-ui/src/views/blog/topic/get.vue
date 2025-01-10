<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item field="sourceShow" :label="$t('topic.sourceId')">
            <label class="upImg" for="upMainImgtopic">
              <img :src="formData.sourceShow" />
            </label>
          </a-form-item>
        </a-col>
        <a-col :span="12"></a-col>
        <a-col :span="12">
          <a-form-item field="title" :label="$t('topic.title')">
            <span class="formSpan">{{ formData.title }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="url" :label="$t('topic.url')">
            <span class="formSpan">{{ formData.url }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="24">
          <a-form-item field="summary" :label="$t('topic.summary')">
            <span class="formSpan">{{ formData.summary }}</span>
          </a-form-item>
        </a-col>
      </a-row>
      <a-col :span="24">
        <a-divider orientation="left">{{ $t('topic.posts') }} </a-divider>
        <posts :list="formData.posts" :remove="false" />
      </a-col>
    </a-form>
  </a-spin>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import { topicInit, topicGet } from '@/api/blog/topic'
import Posts from './posts.vue'
// 入参读取
const props = defineProps({
  pop: {
    type: Object,
    required: true,
    default: () => {
      return {} as Pop
    }
  }
})
// 加载中变量
const { load, setLoad } = useLoad(false)
// 表单数据初始化
const formData = topicInit()
async function get() {
  setLoad(true)
  try {
    const res = await topicGet(props.pop.itemId)
    formData.value = res.data
    return
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
// 页面渲染
onMounted(() => {
  // Nothing
  get()
})
</script>
