<template>
  <div class="m">
    <div class="r">
      <div class="bcn">
        <div class="totp">
          <div class="bg">
            <i class="bz">&#xF014;</i>
          </div>
          <div class="bi">
            <h2 class="bz">专题专栏</h2>
          </div>
        </div>
      </div>
    </div>
    <div class="r">
      <div class="rb rc" v-if="resData.topic && resData.topic[0]">
        <NuxtLink v-for="top in resData.topic" class="b b3" :to="'/topic/' + top.url">
          <div class="bg">
            <img :src="top.sourceShow" />
            <div class="bgt">
              <span class="bz">{{ top.num }}</span>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
    <Foot></Foot>
  </div>
</template>

<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();
const needRun = useState("topicsRun", () => true);
const resData = useState("resData", () => <any>{});
useHead({
  title: `专题专栏${runtimeConfig.public.siteName}`,
  meta: [{ hid: "description", name: "description", content: `专题专栏 - ${runtimeConfig.public.description}` }],
});
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/topics", {});
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
  const res = await $fetch("/api/page/topics", { server: false, method: "GET" });
  resData.value = (res as any).data;
}
onUnmounted(() => {
  needRun.value = true;
  // @ts-ignore
  InitBack();
});
function initPage() {
  // @ts-ignore
  InitDom();
}
</script>
