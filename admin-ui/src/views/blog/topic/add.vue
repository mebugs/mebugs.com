<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-form-item field="sourceShow" :label="$t('topic.sourceId')" :rules="[{ required: true, message: $t('rule.required') }]">
          <label class="upImg" for="upMainImgtopicAdd">
            <img :src="formData.sourceShow" v-if="formData.sourceShow" />
            <p v-else>
              <icon-upload />
              {{ $t('button.upload') }}
            </p>
          </label>
          <input
            id="upMainImgtopicAdd"
            accept="image/gif, image/jpeg, image/png, image/jpg"
            type="file"
            style="display: none"
            @change="chooesMain" />
        </a-form-item>
      </a-col>
      <a-col :span="12"></a-col>
      <a-col :span="12">
        <a-form-item field="title" :label="$t('topic.title')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.title" :max-length="32" allow-clear show-word-limit :placeholder="$t('topic.title.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="url" :label="$t('topic.url')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.url" :max-length="32" allow-clear show-word-limit :placeholder="$t('topic.url.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-form-item field="summary" :label="$t('topic.summary')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.summary" :max-length="64" allow-clear show-word-limit :placeholder="$t('topic.summary.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-divider orientation="left">
          {{ $t('topic.posts') }}
          <span class="arco-form-item-extra" style="margin: 0 10px"> {{ $t('topic.posts.tips') }} </span>
          <a-tooltip :content="$t('topic.posts.check')" :mini="true">
            <a-button type="primary" size="small" @click="open">
              <template #icon>
                <icon-edit />
              </template>
            </a-button>
          </a-tooltip>
        </a-divider>
        <posts :list="formData.posts" :remove="true" :remove-func="removeQuick" :sort-func="sortQuick" />
      </a-col>
      <a-col :span="24">
        <a-divider />
        <div class="doBtn">
          <a-space>
            <a-button size="large" type="primary" html-type="submit" :loading="load">
              <template #icon>
                <icon-check />
              </template>
              {{ $t('button.submit') }}
            </a-button>
            <a-button size="large" @click="pop.close()">
              <template #icon>
                <icon-close />
              </template>
              {{ $t('button.cancel') }}</a-button
            >
          </a-space>
        </div>
      </a-col>
    </a-row>
  </a-form>
  <a-modal v-model:visible="check" :hide-title="true" :width="900" :on-before-ok="done" :on-before-cancell="cancel">
    <post-check v-if="check" :do-select="doSelect" :data="selectData" />
  </a-modal>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import { topicInit, topicAdd } from '@/api/blog/topic'
import { CheckList } from '@/utils/hooks/check'
import useImgs from '@/utils/hooks/imgs'
import Posts from './posts.vue'
import PostCheck from './postCheck.vue'
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
// 选择框初始化
const { check, selectData, open, cancel, done, doSelect, removeQuick, sortQuick } = CheckList(formData, 'postIds', 'posts')
// const topicAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      const sortObj: any[] = []
      formData.value.posts.forEach((item: any, i) => {
        sortObj.push({
          id: item.id,
          sort: i
        })
      })
      values.postIdSet = sortObj
      // const res = await topicAdd.value?.validate();
      await topicAdd(values)
      // Pop Close & Back
      props.pop.close()
      props.pop.callBack()
    } catch (err) {
      // DoNothing
    } finally {
      setLoad(false)
    }
  }
}
// 主图加载器
const { imgObj, initImgQuick, chooesImg } = useImgs()
initImgQuick('main')
const chooesMain = async (e: Event) => {
  await chooesImg(e)
    .then(() => {
      formData.value.source = imgObj.value.baseUrls
      formData.value.sourceShow = imgObj.value.baseUrls[0]
    })
    .catch((er) => {
      console.log(er)
    })
}
// 页面渲染
onMounted(() => {
  // Nothing
  setLoad(false)
})
</script>
