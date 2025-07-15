<template>
  <div class="la">
    <NuxtRouteAnnouncer />
    <NuxtLoadingIndicator :color="'#3491fa'" :height="2" />
    <Menu></Menu>
    <div class="lr"><NuxtPage /></div>
    <Comm></Comm>
  </div>
</template>

<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();
useHead({
  link: [{ rel: "stylesheet", href: "/static/css/base.css" }],
  // script: [{ src: "/static/js/base.js", tagPosition: "bodyClose" }],
});
// GoogleADS
onMounted(() => {
  const script = document.createElement("script");
  script.src = "https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-2699978133125218";
  script.async = true;
  script.crossOrigin = "anonymous";
  document.head.appendChild(script);
  // @ts-ignore - Google Ads global variable not defined in types
  (adsbygoogle = window.adsbygoogle || []).push({});
  // 初始化百度统计
  // @ts-ignore - baidu
  window._hmt = window._hmt || [];
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?64b27967e9d2fa5d1d4c6f945051603e";
  var s = document.getElementsByTagName("script")[0];
  // @ts-ignore - baidu
  document.head.appendChild(hm);
});
</script>
