<template>
  <!--表格，子页无需吸顶 -->
  <a-table
    :scrollbar="false"
    :sticky-header="false"
    :row-key="'id'"
    :columns="posts"
    :draggable="drag"
    :data="list"
    :pagination="false"
    size="medium"
    style="margin-top: 10px"
    @change="sortChange">
    <template #operations="{ record }" v-if="remove">
      <a-space>
        <a-tooltip :content="$t('button.delete')" :mini="true">
          <a-button type="text" size="small" @click="remove(record.id)">
            <template #icon> <icon-delete /> </template>
          </a-button>
        </a-tooltip>
      </a-space>
    </template>
  </a-table>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
// 入参读取
const props = defineProps({
  list: {
    type: Object,
    required: true,
    default: () => {
      return []
    }
  },
  dictMap: {
    type: Object,
    required: true,
    default: () => {
      return {}
    }
  },
  remove: {
    type: Boolean,
    default: () => {
      return false
    }
  },
  removeFunc: {
    type: Function,
    default: (id: number) => {
      return id
    }
  },
  sortFunc: {
    type: Function,
    default: (ids: any) => {
      return ids
    }
  }
})
// 排序
function sortChange(data: any) {
  props.sortFunc(data)
}
// 移除数据
function remove(id: number) {
  props.removeFunc(id)
}
// 响应拖拽
const drag = computed(() => {
  if (props.remove) {
    return { type: 'handle', width: 40 }
  }
  return {}
})
// 路由列表对象
const posts = computed(() => {
  let c: any[] = [
    { title: t('post.title'), dataIndex: 'title' },
    { title: t('post.url'), dataIndex: 'url' }
  ]
  if (props.remove) {
    c.push({ title: t('base.oper'), slotName: 'operations', width: 80 })
  }
  return c
})
</script>
