<template>
  <div class="f">
    <div class="fr flk"><a href="https://beian.miit.gov.cn/" target="_blank">苏ICP备20039109号</a></div>
    <div class="fr flk">
      <a v-for="item in resData" :href="item.url" target="_blank">{{ item.title }}</a>
      <NuxtLink class="bt" to="/page/link">#更多</NuxtLink>
    </div>
    <div class="fr flk"><NuxtLink to="/post/about">关于</NuxtLink>丨<NuxtLink to="/page/msg">留言</NuxtLink>丨<NuxtLink to="/page/link">友链</NuxtLink>丨<NuxtLink to="/page/map">地图</NuxtLink></div>
    <div class="fr flk fcn">
      <a href="http://wpa.qq.com/msgrd?v=3&amp;uin=7431346&amp;site=qq&amp;menu=yes" target="_blank"> <i>&#xF016;</i>QQ</a>
      <a href="http://wpa.qq.com/msgrd?v=3&amp;uin=7431346&amp;site=qq&amp;menu=yes" target="_blank"> <i>&#xF017;</i>微信</a>
      <a href="mailto:iam@qiantaoyuan.com" target="_blank"> <i>&#xF018;</i>邮箱</a>
      <a href="https://github.com/mebugs" target="_blank"> <i>&#xF019;</i>仓库</a>
      <NuxtLink to="/page/msg"> <i>&#xF020;</i>留言</NuxtLink>
    </div>
    <div class="fr flk kc"><NuxtLink class="kc" href="/">&#xF015; MEBUGS</NuxtLink></div>
  </div>
</template>

<script setup lang="ts">
const needRun = useState("pageLink", () => true);
const resData = useState("links", () => <any>[]);
const runtimeConfig = useRuntimeConfig();
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/link", { method: "POST", body: {} });
  resData.value = (res.data.value as any).data;

  needRun.value = false;
}
// 页面挂载后的初始化
onMounted(async () => {
  // 如果SSR没有
  if (needRun.value) {
    await initByClient();
  }
});
async function initByClient() {
  const res = await $fetch("/api/page/link", { server: false, method: "POST", body: {} });
  resData.value = (res as any).data;
}
onUnmounted(() => {});
</script>
