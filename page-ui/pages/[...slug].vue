<template>
  <div class="m">
    <Null></Null>
    <!--广告位 -->
    <div class="addeare"><Adsbygoogle ad-slot="1524459291" /></div>

    <Foot></Foot>
  </div>
</template>

<script setup lang="ts">
import Null from "~/components/Null.vue";

const runtimeConfig = useRuntimeConfig();
useHead({
  title: `哎呀呀！没找到！${runtimeConfig.public.siteName}`,
  meta: [{ hid: "description", name: "description", content: `哎呀呀！没找到！ - ${runtimeConfig.public.description}` }],
});
</script>
