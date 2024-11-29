<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-form-item field="name" :label="$t('source.name')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.name" :max-length="32" allow-clear show-word-limit :placeholder="$t('source.name.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="fileType" :label="$t('source.fileType')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.fileType"
            :options="pop.dictList.sourceType"
            allow-clear
            allow-search
            :placeholder="$t('button.all')"
            @change="resetImg" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="img" :label="$t('source.img')" :rules="[{ required: true, message: $t('rule.required') }]">
          <label class="upImg" for="upMainImgSourceMod">
            <img :src="formData.img" v-if="formData.img" />
            <p v-else>
              <icon-upload />
              {{ $t('button.upload') }}
            </p>
          </label>
          <input
            id="upMainImgSourceMod"
            accept="image/gif, image/jpeg, image/png, image/jpg"
            type="file"
            style="display: none"
            @change="chooesMain" />
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
import { sourceInit, sourceGet, sourceEdit } from '@/api/blog/source'
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
// const sourceAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await sourceAdd.value?.validate();
      await sourceEdit(values)
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
const resetImg = () => {
  initImgQuick('')
  if (formData.value.fileType == '0') {
    initImgQuick('main')
  }
  if (formData.value.fileType == '1') {
    initImgQuick('icon')
  }
  formData.value.fileStrArray = []
  formData.value.img = ''
  const fileInput = document.getElementById('upMainImgSourceMod') as HTMLInputElement
  if (fileInput) {
    fileInput.value = ''
  }
}
const chooesMain = async (e: Event) => {
  await chooesImg(e)
    .then(() => {
      formData.value.fileStrArray = imgObj.value.baseUrls
      formData.value.img = imgObj.value.baseUrls[0]
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
