<template>
  <div class="m">
    <div class="r">
      <div class="rb">
        <div class="bmar b1">
          <h1 class="bz bpt">📖 {{ resData.title }}</h1>
          <div class="bg bpm">
            <img :src="resData.sourceShow" />
          </div>
          <div class="pdesc bz">
            <span>🎯</span>
            <p>{{ resData.summary }}</p>
          </div>
          <div class="b pcom">
            <button type="button" @click="openAddLink()"><i>&#xF027;</i>友链申请</button>
          </div>
        </div>
      </div>
      <div class="rb rt rtgb rtlink" v-if="resData.links && resData.links[0]">
        <NuxtLink target="_blank" v-for="link in resData.links" :to="link.url" class="b b3">
          <div class="bg"><img :src="link.sourceShow" /></div>
          <div class="bi">
            <h2 class="bz">{{ link.title }}</h2>
            <div class="bit">
              <span class="bz">{{ link.summary }}</span>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
    <Foot></Foot>
  </div>
  <LinkAdd :linkAdd="openLink" />
</template>

<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();
const needRun = useState("pageLink", () => true);
const resData = useState("resData", () => <any>{});
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/page", { method: "POST", body: { url: "/page/link" } });
  resData.value = (res.data.value as any).data;
  needRun.value = false;
  initHeader();
}
function initHeader() {
  useHead({
    title: `${resData.value?.title ? resData.value?.title : "哎呀呀！没找到！"}${runtimeConfig.public.siteName}`,
    meta: [{ hid: "description", name: "description", content: `${resData.value?.summary ? resData.value?.summary : "哎呀呀！没找到！"} - ${runtimeConfig.public.description}` }],
  });
}
// 打开友链添加
const openLink = ref({ show: false });
function openAddLink() {
  openLink.value.show = true;
}
// 页面挂载后的初始化
onMounted(async () => {
  // 如果SSR没有
  if (needRun.value) {
    await initByClient();
    initHeader();
  }
  initPage();
});
async function initByClient() {
  const res = await $fetch("/api/page/page", { server: false, method: "POST", body: { url: "/page/link" } });
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
