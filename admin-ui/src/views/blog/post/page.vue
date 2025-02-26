<template>
  <a-form size="large" :model="query" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="8">
        <a-form-item field="title" :label="$t('post.title')">
          <a-input v-model="query.title" allow-clear show-word-limit :placeholder="$t('post.title.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="status" :label="$t('post.status')">
          <a-select v-model="query.status" :options="dictList.postStatus" allow-clear allow-search :placeholder="$t('button.all')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item v-if="postType == '1'" field="category" :label="$t('post.category')">
          <a-select
            v-model="query.categoryId"
            :options="dictList.categoryList"
            allow-clear
            allow-search
            @clear="toNull"
            :placeholder="$t('button.all')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-space>
          <a-tooltip :content="$t('button.add')" :mini="true">
            <a-button
              v-permission="''"
              size="large"
              type="primary"
              status="danger"
              @click="
                pop.open(
                  'add',
                  0,
                  postType == '1' ? $t('post.add') : $t('page.add'),
                  postType == '1' ? $t('post.add.sub') : $t('page.add.sub'),
                  {},
                  search
                )
              ">
              <template #icon>
                <icon-plus />
              </template>
            </a-button>
          </a-tooltip>
          <!-- <a-divider direction="vertical" /> -->
        </a-space>
      </a-col>
      <a-col :span="12" class="doBtn">
        <a-space>
          <a-tooltip :content="$t('button.search')" :mini="true">
            <a-button v-permission="''" size="large" type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
            </a-button>
          </a-tooltip>
          <a-tooltip :content="$t('button.reset')" :mini="true">
            <a-button v-permission="''" size="large" @click="resetQuery">
              <template #icon>
                <icon-refresh />
              </template>
            </a-button>
          </a-tooltip>
        </a-space>
      </a-col>
    </a-row>
  </a-form>
  <a-divider />
  <a-spin :loading="load">
    <a-row :gutter="20">
      <a-col :span="24" class="doBtn" style="margin-bottom: 20px">
        <a-pagination :total="page.total" :page-size="page.pageSize" show-total @change="changePage" />
      </a-col>
      <a-col :span="12" v-for="record in list" :key="record.id">
        <div class="cardlist postList">
          <div class="upIcon postImg">
            <a-tag class="fTag" :color="flagTag[record.status]">{{ dictMap.postStatus[record.status] }}</a-tag>
            <p v-if="postType == '1'">{{ dictMap.categoryList[record.categoryId] }}</p>
            <img :src="record.sourceShow" />
          </div>
          <div class="postInfo">
            <h3>
              {{ record.title }}
            </h3>
            <p>
              {{ record.summary }}
            </p>
          </div>
          <a-space class="postBtn">
            <a-tooltip :content="$t('button.edit')" :mini="true">
              <a-button
                size="large"
                v-permission="''"
                type="text"
                @click="pop.open('edit', record.id, postType == '1' ? $t('post.edit') : $t('page.edit'), record.title, {}, search)">
                <template #icon> <icon-edit /> </template>
              </a-button>
            </a-tooltip>
          </a-space>
        </div>
      </a-col>
    </a-row>
  </a-spin>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, watch } from 'vue'
import useLocale from '@/utils/hooks/locale'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import usePage from '@/utils/hooks/page'
import { dictRead } from '@/api/plat/dict'
import { postPage, postDel } from '@/api/blog/post'
import { categoryList } from '@/api/blog/category'
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
const { load, setLoad } = useLoad(false)
// 当前语言
const { currentLocale } = useLocale()
// 分页
const { page, setQuery, search, changePage, resetPage } = usePage(10)
// 初始化查询对象
const initQuery = () => {
  return { title: '', url: '', status: '', categoryId: null, postType: props.postType }
}
const toNull = () => {
  query.value.categoryId = null
}
// 查询对象
const query = ref(initQuery())
// 列表对象
const list = ref<any>([])
// 分页检索
async function pageQuery() {
  if (load.value) return
  setLoad(true)
  try {
    const res = await postPage({ ...query.value, ...page })
    list.value = res.data.list
    page.total = res.data.total
  } catch (e) {
    // 清空数据
    list.value = []
    page.current = 1
    page.total = 0
  } finally {
    setLoad(false)
  }
}
// 初始化分页
setQuery(pageQuery)
// 状态标签
const flagTag: any = { '0': 'orange', '1': 'green', '2': 'red' }
// 初始化字典对象
const dictList = ref({ postStatus: [], categoryList: [] })
const dictMap = ref({ postStatus: {} as any, categoryList: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['postStatus'] }).then((r) => {
    dictList.value.postStatus = r.data.list.postStatus
    dictMap.value.postStatus = r.data.map.postStatus

    props.pop.dictMap = dictMap
  })
  await categoryList().then((dr) => {
    dictList.value.categoryList = dr.data.list
    dictMap.value.categoryList = dr.data.map
  })
  props.pop.dictList = dictList
}
function init() {
  // 初始化后端字典对象
  dictInit()
  // 初始化搜索
  pageQuery()
}
// 重置查询
function resetQuery() {
  query.value = initQuery()
  list.value = []
  // 重置分页
  resetPage()
  init()
}
// 页面渲染
onMounted(() => {
  init()
})
// 语言监听
watch(currentLocale, (n, o) => {
  if (n !== o) {
    dictInit()
  }
})
</script>
