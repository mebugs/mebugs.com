<template>
  <div class="m">
    <div class="r">
      <div class="bcn">
        <div class="totp">
          <div class="bg">
            <i class="bz">&#xF020;</i>
          </div>
          <div class="bi">
            <h2 class="bz">更多页面</h2>
          </div>
        </div>
      </div>
    </div>
    <div class="r">
      <div class="rb rc" v-if="resData && resData[0]">
        <NuxtLink v-for="page in resData" class="b b3" :to="page.url">
          <div class="bg">
            <img :src="page.sourceShow" />
          </div>
          <div class="bi">
            <h1 class="bz">📖 {{ page.title }}</h1>
          </div>
        </NuxtLink>
      </div>
    </div>
    <Foot></Foot>
  </div>
</template>

<script setup lang="ts">
import { InitDom, InitBack } from "~/utils/base.js";
const runtimeConfig = useRuntimeConfig();
const needRun = useState("pageData", () => true);
const resData = useState("resData", () => <any>{});
useHead({
  title: `更多页面${runtimeConfig.public.siteName}`,
  meta: [{ hid: "description", name: "description", content: `${runtimeConfig.public.description}` }],
});
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/pages", {});
  resData.value = (res.data.value as any).data;
  needRun.value = false;
}
// 页面挂载后的初始化
onMounted(async () => {
  // 如果SSR没有
  if (needRun.value) {
    await initByClient();
  }
  initPage();
});
async function initByClient() {
  const res = await $fetch("/api/page/pages", { server: false, method: "GET" });
  resData.value = (res as any).data;
}
onUnmounted(() => {
  needRun.value = true;
  InitBack();
});
function initPage() {
  InitDom();
}
</script>
