<template>
  <div class="m post" v-if="resData.post">
    <div class="r">
      <div class="rbx">
        <div class="bx b1">
          <h1 class="bz">📖 {{ resData.post.title }}</h1>
        </div>
      </div>
    </div>
    <div class="r">
      <div class="rbx">
        <div class="bx bl">
          <div class="bg bpm">
            <img :src="resData.post.sourcePath + '.' + resData.post.sourceBack" />
          </div>
        </div>
        <div class="b bx br">
          <div class="bi bix">
            <div class="bip bz">
              <em>发布日期：</em><span><i>&#xF008;</i> {{ resData.post.pushAt }}</span>
            </div>
            <div class="bip bz">
              <em>访问流量：</em><span><i>&#xF010;</i> {{ resData.post.views }}</span>
            </div>
            <div class="bip bz">
              <em>热门指数：</em><span><i>&#xF007;</i> {{ resData.post.hots }}</span>
            </div>
            <div class="bip bz">
              <em>推荐权重：</em><span><i>&#xF011;</i> {{ resData.post.goods }}</span>
            </div>
            <div class="bip bz">
              <em>所述分类：</em>
              <NuxtLink :to="'/category/' + resData.post.category"><i>&#xF012;</i> {{ resData.post.categoryName }}</NuxtLink>
            </div>
            <div class="bip bz" v-if="resData.post.tagUrls && resData.post.tagUrls[0]">
              <em>相关标签：</em>
              <NuxtLink v-for="(ta, ti) in resData.post.tagUrls" :to="'/tag/' + ta"><i>&#xF013;</i> {{ resData.post.tagNames[ti] }}</NuxtLink>
            </div>
            <div class="bip bz" v-if="resData.post.topicUrls && resData.post.topicUrls[0]">
              <em>收录专栏：</em>
              <NuxtLink v-for="(ta, ti) in resData.post.topicUrls" :to="'/tag/' + ta"><i>&#xF014;</i> {{ resData.post.topicNames[ti] }}</NuxtLink>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="r">
      <div class="rbx">
        <div class="bx bl">
          <div class="pdesc bz">
            <span>🎯</span>
            <p>{{ resData.post.summary }}</p>
          </div>
          <em id="st"></em>
          <div class="vditor-reset bz" v-html="resData.postMain.html"></div>
          <em id="ed"></em>
        </div>
        <div class="bx br brdh bz">
          <div class="brl" v-html="resData.postMain.toc"></div>
        </div>
      </div>
    </div>
    <Foot></Foot>
  </div>
</template>

<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();
const route = useRoute();
const url = route.params.url;
const needRun = useState("postRun", () => true);
const resData = useState("resData", () => <any>{});
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/post", {
    method: "POST",
    body: { url: url },
  });
  resData.value = (res.data.value as any).data;
  initHeader();
  needRun.value = false;
}

function initHeader() {
  useHead({
    title: `${resData.value.null ? "哎呀呀！没找到！" : resData.value?.post?.title}${runtimeConfig.public.siteName}`,
    meta: [{ hid: "description", name: "description", content: `${resData.value.null ? "哎呀呀！没找到！" : resData.value?.post?.summary} - ${runtimeConfig.public.description}` }],
    link: [{ href: "/static/css/ant-design.css", rel: "stylesheet" }],
    script: [{ src: "/static/js/post.js", tagPosition: "bodyClose" }],
  });
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
  const res = await $fetch("/api/page/post", { server: false, method: "POST", body: { url: url } });
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
  // @ts-ignore
  InitPostDom();
}
</script>
