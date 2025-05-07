<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="16">
        <a-row :gutter="20">
          <a-col :span="24">
            <a-form-item field="postName" label="文章信息">
              <template #extra>
                <div>{{ formData.postDesc ? formData.postName : '公共留言板' }}</div>
              </template>
              <span class="formSpan">{{ formData.postName ? formData.postName : '留言板' }}</span>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="user.name" label="用户昵称">
              <a-input v-model="formData.user.name" allow-clear show-word-limit :placeholder="$t('role.remark.place')" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="user.name" label="用户邮箱">
              <a-input v-model="formData.user.email" allow-clear show-word-limit :placeholder="$t('role.remark.place')" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="user.name" label="用户地址">
              <template #extra>
                <div>当前地址：{{ formData.user.thirdUrl ? formData.user.thirdUrl : '空' }}</div>
              </template>
              <a-input v-model="formData.user.waitUrl" allow-clear show-word-limit :placeholder="$t('role.remark.place')" />
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item field="postName" label="签名信息">
              <a-input v-model="formData.user.summary" allow-clear show-word-limit :placeholder="$t('role.remark.place')" />
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item field="postName" label="评论内容">
              <a-textarea v-model="formData.info" allow-clear show-word-limit :placeholder="$t('role.remark.place')" />
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item field="postName" label="回复内容">
              <a-textarea v-model="formData.reInfo" allow-clear show-word-limit :placeholder="$t('role.remark.place')" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="8">
        <a-row :gutter="20">
          <a-col :span="24">
            <p>相关回复</p>
            <div class="commli" v-if="formData.recList && formData.recList[0]">
              <div class="commlii" v-for="item in formData.recList">
                <h3>{{ item.user.name }}</h3>
                <p>{{ item.info }}</p>
              </div>
            </div>
            <div v-else class="commli">暂无相关回复</div>
          </a-col>
        </a-row>
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
import { postCommentInit, postCommentGet, postCommentEdit } from '@/api/blog/comment'
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
const formData = postCommentInit()
async function get() {
  setLoad(true)
  try {
    const res = await postCommentGet(props.pop.itemId)
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
      await postCommentEdit(values)
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
