<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item field="name" :label="$t('source.name')">
            <span class="formSpan">{{ formData.name }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="fileType" :label="$t('source.fileType')" :rules="[{ required: true, message: $t('rule.required') }]">
            <span class="formSpan">{{ pop.dictMap.sourceType[formData.fileType] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="img" :label="$t('source.img')" :rules="[{ required: true, message: $t('rule.required') }]">
            <label class="upImg" for="upMainImg">
              <img :src="formData.img" />
            </label>
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
import { sourceInit, sourceGet } from '@/api/blog/source'
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
const formData = sourceInit()
async function get() {
  setLoad(true)
  try {
    const res = await sourceGet(props.pop.itemId)
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
