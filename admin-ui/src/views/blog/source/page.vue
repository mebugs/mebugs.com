<template>
  <a-form size="large" :model="query" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="8">
        <a-form-item field="title" :label="$t('source.name')">
          <a-input v-model="query.name" :max-length="16" allow-clear show-word-limit :placeholder="$t('source.name.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="fileType" :label="$t('source.fileType')">
          <a-select v-model="query.fileType" :options="dictList.sourceType" allow-clear allow-search :placeholder="$t('button.all')" />
        </a-form-item>
      </a-col>
      <a-col :span="8"> </a-col>
      <a-col :span="12">
        <a-space>
          <a-tooltip :content="$t('button.add')" :mini="true">
            <a-button
              v-permission="''"
              size="large"
              type="primary"
              status="danger"
              @click="pop.open('add', 0, $t('source.add'), $t('source.add.sub'), {}, search)">
              <template #icon>
                <icon-plus />
              </template>
            </a-button>
          </a-tooltip>

          <a-tooltip v-if="!doC" :content="$t('button.clean')" :mini="true">
            <a-button v-permission="''" size="large" type="primary" status="success" @click="openDelete">
              <template #icon> <icon-brush /> </template>
            </a-button>
          </a-tooltip>
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
      <a-col :span="6" v-for="record in list" :key="record.id">
        <div class="cardlist">
          <div class="upImg" :class="{ upIcon: doC }">
            <img :src="record.img" />
          </div>
          <p v-if="doC">{{ record.name + '.' + record.backEnd }}</p>
          <p v-else>{{ record.name + '.' + record.backEnd }} / {{ dictMap.sourceType[record.fileType] }}</p>
          <a-space>
            <a-tooltip v-if="!doC" :content="$t('button.get')" :mini="true">
              <a-button v-permission="''" type="text" @click="pop.open('get', record.id, $t('source.get'), record.title, {}, search)">
                <template #icon> <icon-eye /> </template>
              </a-button>
            </a-tooltip>
            <a-tooltip v-if="!doC" :content="$t('button.edit')" :mini="true">
              <a-button v-permission="''" type="text" @click="pop.open('edit', record.id, $t('source.edit'), record.title, {}, search)">
                <template #icon> <icon-edit /> </template>
              </a-button>
            </a-tooltip>
            <a-tooltip v-if="doC" :content="$t('button.check')" :mini="true">
              <a-button v-permission="''" type="text" @click="doCheck(record)">
                <template #icon> <icon-check /> </template>
              </a-button>
            </a-tooltip>
          </a-space>
        </div>
      </a-col>
    </a-row>
  </a-spin>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :width="400" :title="$t('title.clean')" @before-ok="deleting">
    <div>{{ $t('source.del.tips') }}</div>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, watch } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import usePage from '@/utils/hooks/page'
import { dictRead } from '@/api/plat/dict'
import { sourcePage, sourceDel } from '@/api/blog/source'
// 入参读取
const props = defineProps({
  pop: {
    type: Object,
    required: true,
    default: () => {
      return {} as Pop
    }
  },
  doC: {
    type: Boolean,
    default: () => {
      return false
    }
  },
  doCheck: {
    type: Function,
    default: (o: any) => {}
  }
})
// 加载中变量
const { load, setLoad } = useLoad(false)
// 当前语言
const { currentLocale } = useLocale()
// 分页
const { page, setQuery, search, changePage, resetPage } = usePage(8)
// 初始化查询对象
const initQuery = () => {
  return { name: '', fileType: '' }
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
    const res = await sourcePage({ ...query.value, ...page })
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
// 初始化字典对象
const dictList = ref({ sourceType: [] })
const dictMap = ref({ sourceType: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['sourceType'] }).then((r) => {
    dictList.value = r.data.list
    dictMap.value = r.data.map
    props.pop.dictList = dictList
    props.pop.dictMap = dictMap
  })
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
// 删除对象
const delItem = reactive({
  delConfirm: false
})
// 打开删除
function openDelete() {
  delItem.delConfirm = true
}
// 确认删除
async function deleting() {
  try {
    await sourceDel()
  } catch (err) {
    return false
  } finally {
    // Nothing
    pageQuery()
  }
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
