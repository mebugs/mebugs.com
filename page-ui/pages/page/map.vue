<template>
  <div class="m">
    <div class="r">
      <div class="rb">
        <div class="bmar b1">
          <h1 class="bz bpt">📖{{ resData.title }}</h1>
          <div class="bg bpm">
            <img :src="resData.sourceShow" />
          </div>
          <div class="pdesc bz">
            <span>🎯</span>
            <p>{{ resData.summary }}</p>
          </div>
        </div>
      </div>
      <!--广告位 -->
      <div class="addeare"><Adsbygoogle ad-slot="9654230756" /></div>
      <div class="rb">
        <div class="b b1" v-ripples>
          <input class="mapSearch" v-model="seachTitle" type="text" @keyup.enter="handleSearch" placeholder="输入标题，回车搜索，支持模糊查询" />
        </div>
      </div>
      <div class="rb rc rmap" v-if="posts && posts[0]">
        <NuxtLink v-for="post in posts" class="b b3" :to="'/post/' + post.url">
          <div class="bi">
            <h1 class="bz">📖{{ post.title }}</h1>
          </div>
        </NuxtLink>
      </div>
    </div>
    <!--广告位 -->
    <div class="addeare"><Adsbygoogle ad-slot="1524459291" /></div>
    <Foot></Foot>
  </div>
</template>

<script setup lang="ts">
import { InitDom, InitBack } from "~/utils/base.js";
const runtimeConfig = useRuntimeConfig();
const needRun = useState("pageMap", () => true);
const resData = useState("resData", () => <any>{});
const seachTitle = ref("");
const posts = useState("posts", () => <any>[]);
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/page", { method: "POST", body: { url: "/page/map" } });
  resData.value = (res.data.value as any).data;
  if (resData.value && resData.value.litePosts && resData.value.litePosts[0]) {
    posts.value = resData.value.litePosts;
  }
  needRun.value = false;
  initHeader();
}
function initHeader() {
  useHead({
    title: `${resData.value?.title ? resData.value?.title : "哎呀呀！没找到！"}${runtimeConfig.public.siteName}`,
    meta: [{ hid: "description", name: "description", content: `${resData.value?.summary ? resData.value?.summary : "哎呀呀！没找到！"} - ${runtimeConfig.public.description}` }],
  });
}
const handleSearch = () => {
  if (seachTitle.value == "") {
    posts.value = resData.value.litePosts;
  } else {
    posts.value = resData.value.litePosts.filter((item: any) => {
      return item.title.includes(seachTitle.value);
    });
  }
};
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
  const res = await $fetch("/api/page/page", { server: false, method: "POST", body: { url: "/page/map" } });
  resData.value = (res as any).data;
  if (resData.value && resData.value.litePosts && resData.value.litePosts[0]) {
    posts.value = resData.value.litePosts;
  }
}
onUnmounted(() => {
  needRun.value = true;
  InitBack();
});
function initPage() {
  InitDom();
}
</script>
