<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="16">
        <a-col :span="24" style="padding: 0">
          <a-form-item field="title" :label="$t('post.title')" :rules="[{ required: true, message: $t('rule.required') }]">
            <a-input v-model="formData.title" :max-length="64" allow-clear show-word-limit :placeholder="$t('post.title.place')" />
          </a-form-item>
        </a-col>
        <a-col :span="24" style="padding: 0">
          <a-form-item field="url" :label="$t('post.url')" :rules="[{ required: true, message: $t('rule.required') }]">
            <a-input v-model="formData.url" :max-length="32" allow-clear show-word-limit :placeholder="$t('post.url.place')" />
          </a-form-item>
        </a-col>
        <a-col :span="24" style="padding: 0">
          <a-form-item field="summary" :label="$t('post.summary')" :rules="[{ required: true, message: $t('rule.required') }]">
            <a-textarea
              v-model="formData.summary"
              :max-length="256"
              allow-clear
              show-word-limit
              :placeholder="$t('post.summary.place')" />
          </a-form-item>
        </a-col>
      </a-col>
      <a-col :span="8">
        <a-form-item field="sourceShow" :label="$t('post.img')" :rules="[{ required: true, message: $t('rule.required') }]">
          <label class="upImg upPostImg" for="upMainImgpostAdd">
            <img :src="formData.sourceShow" v-if="formData.sourceShow" />
            <p v-else>
              <icon-upload />
              {{ $t('button.upload') }}
            </p>
          </label>
          <input
            id="upMainImgpostAdd"
            accept="image/gif, image/jpeg, image/png, image/jpg"
            type="file"
            style="display: none"
            @change="chooesMain" />
        </a-form-item>
      </a-col>
      <a-col :span="24" style="margin: 20px 0">
        <div id="vditorAdd"></div>
      </a-col>
      <a-col :span="12" v-if="postType == '1'">
        <a-form-item field="tagIds" :label="$t('post.tags')">
          <template #extra>
            <div>
              {{ $t('post.tags.tips') }}
              <a-tooltip :content="$t('button.add')" :mini="true">
                <a-button type="primary" size="small" @click="openTagAdd">
                  <template #icon>
                    <icon-plus />
                  </template>
                </a-button>
              </a-tooltip>
            </div>
          </template>
          <a-select v-model="formData.tagIds" :options="tags" multiple allow-clear allow-search :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12" v-if="postType == '1'">
        <a-form-item field="categoryId" :label="$t('post.category')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.categoryId"
            :options="pop.dictList.categoryList"
            allow-clear
            allow-search
            :placeholder="$t('rule.select')"
            @clear="formData.categoryId = null" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="status" :label="$t('post.status')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.status"
            :options="pop.dictList.postStatus"
            allow-clear
            allow-search
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="pushAt" :label="$t('post.pushAt')">
          <a-date-picker style="width: 100%" v-model="formData.pushAt" allow-clear :placeholder="$t('post.pushAt.place')" />
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
  <a-modal v-model:visible="tagAddFlag" :title="$t('tags.add')" :footer="false" :width="900">
    <tag-add v-if="tagAddFlag" :doAdd="doAdd" :doCanc="doCanc" :do="true" />
  </a-modal>
  <!-- 添加图片 -->
  <a-modal v-model:visible="openS" :title="$t('source.check')" :footer="false" :width="900">
    <source-index v-if="openS" :doCheck="doCheck" :doC="true" />
  </a-modal>
</template>

<script lang="ts" setup>
import '~/vditor/src/assets/less/index.less'
import { onMounted, ref } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import vditor from '@/utils/hooks/vditor'
import { postInit, postAdd } from '@/api/blog/post'
import { tagList } from '@/api/blog/tag'
import useImgs from '@/utils/hooks/imgs'
import tagAdd from '../tag/add.vue'
import sourceIndex from '../source/index.vue'
// 入参读取
const props = defineProps({
  pop: {
    type: Object,
    required: true,
    default: () => {
      return {} as Pop
    }
  },
  postType: {
    type: String,
    default: () => {
      return '1' //默认文章
    }
  }
})
// 加载中变量
const { load, setLoad } = useLoad(true)
// 表单数据初始化
const formData = postInit(props.postType)
// 标签列表
const tags = ref([])
const getTagList = async () => {
  setLoad(true)
  await tagList()
    .then((r) => {
      tags.value = r.data
    })
    .finally(() => {
      setLoad(false)
    })
}
const tagAddFlag = ref(false)
const openTagAdd = () => {
  tagAddFlag.value = true
}
const doAdd = () => {
  console.log('INADD')
  tagAddFlag.value = false
  getTagList()
}
const doCanc = () => {
  tagAddFlag.value = false
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
// 打开资源选择
const openS = ref(false)
const openSource = () => {
  openS.value = true
}
// 0 编辑 1 正文预览 2 确认发布
const { vd, md, getNew, getResponse } = vditor('vditorAdd', openSource)
// docheck
const doCheck = (record: any) => {
  openS.value = false
  vd.value.insertValue(`![SC-${record.id}-${record.name}.${record.backEnd} ](${record.img})\n`, true)
}

// const postAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      const sourceList = await getResponse()
      formData.value.sourceIds = sourceList
      formData.value.html = md.html
      formData.value.md = md.md
      formData.value.toc = md.outline
      // const res = await postAdd.value?.validate();
      await postAdd(formData.value)
      // Pop Close & Back
      props.pop.close()
      props.pop.callBack()
    } catch (err) {
      // DoNothing
    } finally {
      setLoad(false)
      console.log('F')
    }
  }
}
// 页面渲染
onMounted(() => {
  // Nothing
  getTagList()
  setLoad(false)
  getNew('')
})
</script>
