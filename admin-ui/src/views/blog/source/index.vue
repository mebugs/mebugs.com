<template>
  <div class="container" :class="{ sPop: pop.pop }">
    <a-card class="sGroup sList" :bordered="false">
      <page :pop="pop" :doCheck="doCheck" :doC="doC" />
    </a-card>
    <a-card class="sGroup sItem" :bordered="false" :loading="false">
      <a-page-header :title="pop.header" :subtitle="pop.subHeader" @back="pop.close()" />
      <add v-if="pop.add" :pop="pop" :doC="doC" />
      <get v-if="pop.get" :pop="pop" />
      <edit v-if="pop.edit" :pop="pop" />
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import { onUnmounted } from 'vue'
import usePop from '@/utils/hooks/pop'
import Page from './page.vue'
import Add from './add.vue'
import Get from './get.vue'
import Edit from './edit.vue'
defineProps({
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
const pop = usePop()
onUnmounted(() => {
  pop.close()
})
</script>
<script lang="ts">
export default {
  name: 'BlogSource'
}
</script>
