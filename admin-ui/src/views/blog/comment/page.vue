<template>
  <a-form size="large" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="12"> </a-col>
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
        <div class="cardlist commList">
          <div class="postIf">
            <h3>{{ record.postName ? record.postName : '留言板' }}</h3>
            <p>{{ record.postDesc ? record.postName : '公共留言板' }}</p>
          </div>
          <div class="commDt">
            <img :src="'/source/a/' + record.user.sourceId + '.jpg'" />
            <div class="commInf">
              <h5>
                <a-tag class="fTag" :color="record.status == '0' ? 'green' : 'orange'">
                  {{ record.status == '0' ? '已通过' : '待审核' }}
                </a-tag>
                丨 {{ record.user.name }}
              </h5>
              <div class="pcominp">
                <p>{{ record.info }}</p>
              </div>
            </div>
          </div>
          <a-space class="postBtn">
            <a-tooltip :content="$t('button.edit')" :mini="true">
              <a-button size="large" v-permission="''" type="text" @click="pop.open('edit', record.id, '审核评论', {}, {}, search)">
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
import { ref, onMounted, watch } from 'vue'
import useLocale from '@/utils/hooks/locale'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import usePage from '@/utils/hooks/page'
import { postCommentPage } from '@/api/blog/comment'
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
// 列表对象
const list = ref<any>([])
// 分页检索
async function pageQuery() {
  if (load.value) return
  setLoad(true)
  try {
    const res = await postCommentPage({ ...page })
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
function init() {
  // 初始化搜索
  pageQuery()
}
// 重置查询
function resetQuery() {
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
  }
})
</script>
