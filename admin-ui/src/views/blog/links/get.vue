<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item field="sourceShow" :label="$t('links.sourceId')">
            <label class="upIcon" for="upMainImglinks">
              <img :src="formData.sourceShow" />
            </label>
          </a-form-item>
        </a-col>
        <a-col :span="12"></a-col>
        <a-col :span="12">
          <a-form-item field="title" :label="$t('links.title')">
            <span class="formSpan">{{ formData.title }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="url" :label="$t('links.url')">
            <span class="formSpan">{{ formData.url }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="24">
          <a-form-item field="summary" :label="$t('links.summary')">
            <span class="formSpan">{{ formData.summary }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="status" :label="$t('links.status')">
            <span class="formSpan">
              <a-tag class="fTag" :color="formData.status == '0' ? 'green' : formData.status == '2' ? 'red' : 'orange'">
                {{ formData.status == '0' ? '已通过' : formData.status == '2' ? '已拒绝' : '待审核' }}
              </a-tag>
            </span>
          </a-form-item>
        </a-col>
      </a-row>
    </a-form>
  </a-spin>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import { linksInit, linksGet } from '@/api/blog/links'
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
const formData = linksInit()
async function get() {
  setLoad(true)
  try {
    const res = await linksGet(props.pop.itemId)
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
