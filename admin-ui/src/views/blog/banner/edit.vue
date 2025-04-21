<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-form-item field="type" :label="$t('banner.type')">
          <span class="formSpan">{{ pop.dictMap.bannerType[formData.type] }}</span>
        </a-form-item>
      </a-col>
      <a-col :span="12"></a-col>
      <a-col :span="12">
        <a-form-item field="sourceShow" :label="$t('banner.sourceId')" :rules="[{ required: true, message: $t('rule.required') }]">
          <label class="upImg" for="upMainImgbannerMod">
            <img :src="formData.sourceShow" v-if="formData.sourceShow" />
            <p v-else>
              <icon-upload />
              {{ $t('button.upload') }}
            </p>
          </label>
          <input
            id="upMainImgbannerMod"
            accept="image/gif, image/jpeg, image/png, image/jpg"
            type="file"
            style="display: none"
            @change="chooesMain" />
        </a-form-item>
      </a-col>
      <a-col :span="12"></a-col>
      <a-col :span="12">
        <a-form-item field="title" :label="$t('banner.title')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.title" :max-length="32" allow-clear show-word-limit :placeholder="$t('banner.title.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="tag" :label="$t('banner.tag')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.tag" :max-length="12" allow-clear show-word-limit :placeholder="$t('banner.tag.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-form-item field="url" :label="$t('banner.url')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.url" :max-length="128" allow-clear show-word-limit :placeholder="$t('banner.url.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-form-item field="summary" :label="$t('banner.summary')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.summary" :max-length="64" allow-clear show-word-limit :placeholder="$t('banner.summary.place')" />
        </a-form-item>
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
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import { bannerInit, bannerGet, bannerEdit } from '@/api/blog/banner'
import useImgs from '@/utils/hooks/imgs'
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
const formData = bannerInit()
async function get() {
  setLoad(true)
  try {
    const res = await bannerGet(props.pop.itemId)
    formData.value = res.data
    return
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
// const bannerAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await bannerAdd.value?.validate();
      await bannerEdit(values)
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
  get()
})
</script>
