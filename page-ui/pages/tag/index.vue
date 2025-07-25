<template>
  <div class="m">
    <div class="r">
      <div class="bcn">
        <div class="totp">
          <div class="bg">
            <i class="bz">&#xF013;</i>
          </div>
          <div class="bi">
            <h2 class="bz">核心标签</h2>
          </div>
        </div>
      </div>
    </div>
    <!--广告位 -->
    <div class="addeare"><Adsbygoogle ad-slot="6001361550" /></div>
    <div class="r">
      <div class="rb rt rttg pos" :class="{ bzr: !changeDown }" v-if="resData.tag && resData.tag[0]">
        <NuxtLink v-for="ta in resData.tag" :to="'/tag/' + ta.url" class="b b4">
          <div class="bg"><img :src="ta.sourceShow" /></div>
          <div class="bi">
            <span class="bno bz">{{ ta.num }}</span>
            <h2 class="bz">{{ ta.title }}</h2>
            <div class="bit">
              <span class="bz">{{ ta.summary }}</span>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
    <div class="r" v-if="resData.tag && resData.tag[0]">
      <div class="bcn">
        <button type="button" :class="{ nc: !changeSet.pre }" @click="go(0, true)"><i>&#xF003;</i></button>
        <button type="button" :class="{ nc: !changeSet.pre }" @click="go(-1, false)"><i>&#xF004;</i></button>
        <p>{{ changeSet.page }}</p>
        <button type="button" :class="{ nc: !changeSet.next }" @click="go(1, false)"><i>&#xF005;</i></button>
        <button type="button" :class="{ nc: !changeSet.next }" @click="go(changeSet.maxPage, true)"><i>&#xF006;</i></button>
      </div>
    </div>
    <!--广告位 -->
    <div class="addeare"><Adsbygoogle ad-slot="1524459291" /></div>
    <Foot></Foot>
  </div>
</template>

<script setup lang="ts">
import { InitDom, InitBack, LoadPostsImg, ToWhere } from "~/utils/base.js";
const runtimeConfig = useRuntimeConfig();
const needRun = useState("tagRun", () => true);
const resData = useState("resData", () => <any>{});
const changeSet = useState("pageSet", () => <any>{ pre: false, next: true, page: 1, maxPage: 0 });
useHead({
  title: `核心标签${runtimeConfig.public.siteName}`,
  meta: [{ hid: "description", name: "description", content: `核心标签 - ${runtimeConfig.public.description}` }],
});
// 相应数据
let changeDown = ref(true);
function syncPageSet(maxPage: number) {
  changeSet.value.maxPage = maxPage;
  changeSet.value.pre = changeSet.value.page > 1;
  changeSet.value.next = changeSet.value.page < changeSet.value.maxPage;
}
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/tags", {
    method: "POST",
    body: { by: 7, page: 1 },
  });
  resData.value = (res.data.value as any).data;
  syncPageSet(Math.ceil(resData.value.total / 24));
  initHeader();
  needRun.value = false;
}

// 翻页
let query = { by: 7, page: 1 };
function go(i: number, to: boolean) {
  if (changeSet.value.page == 1 && i < 1) {
    return;
  }
  if (changeSet.value.page == changeSet.value.maxPage && i > 0) {
    return;
  }
  if (to) {
    if (i == 0) {
      i = 1;
    }
    changeSet.value.page = i;
  } else {
    changeSet.value.page = changeSet.value.page + i;
  }
  query.page = changeSet.value.page;
  getTags(true);
}
// 查询更多文章
async function getTags(lazyLoad: boolean) {
  changeDown.value = false;
  const res = await $fetch("/api/page/tags", {
    server: false,
    method: "POST",
    body: query,
  });
  resData.value = (res as any)?.data;
  syncPageSet(Math.ceil(resData.value.total / 24));
  await nextTick();
  ToWhere(0);
  setTimeout(() => {
    changeDown.value = true;
    if (lazyLoad) {
      // @ts-ignore
      LoadPostsImg();
    }
  }, 200);
}
function initHeader() {
  useHead({
    title: `核心标签${runtimeConfig.public.siteName}`,
    meta: [{ hid: "description", name: "description", content: `核心标签 - ${runtimeConfig.public.description}` }],
  });
}
// 页面挂载后的初始化
onMounted(async () => {
  // 如果SSR没有
  if (needRun.value) {
    await getTags(false);
    initHeader();
  }
  initPage();
});
onUnmounted(() => {
  needRun.value = true;
  InitBack();
});
function initPage() {
  InitDom();
}
</script>
