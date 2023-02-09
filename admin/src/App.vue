<template>
  <router-view :class="[mode]" />
</template>
<script setup lang="ts">
import { computed } from 'vue';
import STYLE_CONFIG from '@/config/style';
import { useSettingStore } from '@/store';

const store = useSettingStore();

const mode = computed(() => {
  return store.displayMode;
});

// 初始化配置
const initStyleConfig = () => {
  const styleConfig = STYLE_CONFIG;
  for (const key in styleConfig) {
    styleConfig[key] = store[key];
  }
  return styleConfig;
};
store.updateConfig(initStyleConfig());
</script>
