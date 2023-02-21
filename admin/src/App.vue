<template>
  <t-config-provider :global-config="globalConfig">
    <router-view :class="[mode]" />
  </t-config-provider>
</template>
<script setup lang="ts">
import { computed } from 'vue';
import { Icon } from 'tdesign-icons-vue-next';
import STYLE_CONFIG from '@/config/style';
import { useSettingStore } from '@/store';
// 创建全区配置
const globalConfig = {
  icon: Icon,
};
// 用户设置配置读取
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
