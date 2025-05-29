<template>
  <div class="m">
    <div class="r">
      <div class="bcn">
        <div class="totp" v-if="resData.main">
          <div class="bg">
            <img :src="resData.main.sourceShow" />
          </div>
          <div class="bi">
            <h2 class="bz">{{ resData.main.title }}</h2>
            <div class="bit">
              <span class="bz">{{ resData.main.summary }}</span>
            </div>
          </div>
          <div class="bt">
            <NuxtLink class="bz" to="/category">返回分类</NuxtLink>
          </div>
        </div>
      </div>
    </div>
    <div class="r">
      <div class="rb pos" :class="{ bzr: !changeDown }" v-if="resData.posts && resData.posts[0]">
        <NuxtLink class="b b3" v-for="post in resData.posts" :to="'/post/' + post.url">
          <div class="bg">
            <img :src="post.sourcePath + '_1.' + post.sourceBack" />
            <div class="bgt">
              <span class="bz"><i>&#xF013;</i>{{ post.categoryName }}</span>
              <span class="bz"><i>&#xF008;</i>{{ post.pushAt }}</span>
            </div>
          </div>
          <div class="bi">
            <h1 class="bz">📖 {{ post.title }}</h1>
            <div class="bit">
              <span class="bz" v-for="t in post.tagNames"><i>#</i>{{ t }}</span>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
    <div class="r" v-if="resData.posts && resData.posts[0]">
      <div class="bcn">
        <button type="button" :class="{ nc: !changeSet.pre }" @click="go(1, true)"><i>&#xF003;</i></button>
        <button type="button" :class="{ nc: !changeSet.pre }" @click="go(-1, false)"><i>&#xF004;</i></button>
        <p>{{ changeSet.page }}</p>
        <button type="button" :class="{ nc: !changeSet.next }" @click="go(1, false)"><i>&#xF005;</i></button>
        <button type="button" :class="{ nc: !changeSet.next }" @click="go(changeSet.maxPage, true)"><i>&#xF006;</i></button>
      </div>
    </div>
    <Null v-if="resData.null"></Null>
    <Foot></Foot>
  </div>
</template>

<script setup lang="ts">
import { InitDom, InitBack, LoadPostsImg, ToWhere } from "~/utils/base.js";
const runtimeConfig = useRuntimeConfig();
const route = useRoute();
const url = route.params.url;
const needRun = useState("postsRun", () => true);
const resData = useState("resData", () => <any>{});
const changeSet = useState("pageSet", () => <any>{ pre: false, next: true, page: 1, maxPage: 0 });
// 相应数据
let changeDown = ref(true);
function syncPageSet(maxPage: number) {
  changeSet.value.maxPage = maxPage;
  changeSet.value.pre = changeSet.value.page > 1;
  changeSet.value.next = changeSet.value.page < changeSet.value.maxPage;
}
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/posts", {
    method: "POST",
    body: { by: 6, value: url, page: 1 },
  });
  resData.value = (res.data.value as any).data;
  syncPageSet(Math.ceil(resData.value.total / 15));
  initHeader();
  needRun.value = false;
}

// 翻页
let query = { by: 6, value: url, page: 1 };
function go(i: number, to: boolean) {
  if (to) {
    changeSet.value.page = i;
  } else {
    if (changeSet.value.page == 1 && i < 0) {
      return;
    }
    if (changeSet.value.page == changeSet.value.maxPage && i > 0) {
      return;
    }
    changeSet.value.page = changeSet.value.page + i;
  }
  query.page = changeSet.value.page;
  getPosts(true);
}
// 查询更多文章
async function getPosts(lazyLoad: boolean) {
  changeDown.value = false;
  const res = await $fetch("/api/page/posts", {
    server: false,
    method: "POST",
    body: query,
  });
  resData.value = (res as any)?.data;
  syncPageSet(Math.ceil(resData.value.total / 15));
  await nextTick();
  ToWhere(0);
  setTimeout(() => {
    changeDown.value = true;
    if (lazyLoad) {
      LoadPostsImg();
    }
  }, 200);
}
function initHeader() {
  useHead({
    title: `${resData.value?.main?.title ? resData.value?.main?.title : "哎呀呀！没找到！"} - 文章分类${runtimeConfig.public.siteName}`,
    meta: [
      { hid: "description", name: "description", content: `${resData.value?.main?.summary ? resData.value?.main?.summary : "哎呀呀！没找到！"} - 文章分类 - ${runtimeConfig.public.description}` },
    ],
  });
}
// 页面挂载后的初始化
onMounted(async () => {
  // 如果SSR没有
  if (needRun.value) {
    await getPosts(false);
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
