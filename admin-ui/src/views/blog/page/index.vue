<template>
  <div class="container" :class="{ sPop: pop.pop }">
    <a-card class="sGroup sList" :bordered="false">
      <page :pop="pop" :post-type="'0'" />
    </a-card>
    <a-card class="sGroup sItem" :bordered="false" :loading="false">
      <a-page-header :title="pop.header" :subtitle="pop.subHeader" @back="pop.close()" />
      <add v-if="pop.add" :pop="pop" :post-type="'0'" />
      <edit v-if="pop.edit" :pop="pop" :post-type="'0'" />
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import { onUnmounted } from 'vue'
import usePop from '@/utils/hooks/pop'
import Page from '../post/page.vue'
import Add from '../post/add.vue'
import Edit from '../post/edit.vue'
const pop = usePop()
onUnmounted(() => {
  pop.close()
})
</script>
<script lang="ts">
export default {
  name: 'BlogPage'
}
</script>
