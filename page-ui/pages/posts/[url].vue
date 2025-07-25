<template>
  <div class="m">
    <div class="r">
      <div class="bcn">
        <div class="totp">
          <div class="bg">
            <i class="bz" v-html="postsMain.main.icon"></i>
          </div>
          <div class="bi">
            <h2 class="bz">{{ postsMain.main.tip }}</h2>
          </div>
        </div>
      </div>
      <div class="bcn">
        <NuxtLink class="toot" v-for="other in postsMain.others" :to="'/posts/' + other.url">
          <div class="bg">
            <i class="bz" v-html="other.icon"></i>
          </div>
          <div class="bi">
            <h2 class="bz">{{ other.tip }}</h2>
          </div>
        </NuxtLink>
      </div>
    </div>
    <!--广告位 -->
    <div class="addeare"><Adsbygoogle ad-slot="6001361550" /></div>
    <div class="r">
      <div class="rb pos" :class="{ bzr: !changeDown }" v-if="resData.posts && resData.posts[0]">
        <NuxtLink class="b b3" v-for="post in resData.posts" :to="'/post/' + post.url">
          <div class="bg">
            <img :src="post.sourcePath + '_1.' + post.sourceBack" />
            <div class="bgt">
              <span class="bz"><i>&#xF013;</i>{{ post.categoryName }}</span>
              <span class="bz" v-if="urlCode == 1"><i>&#xF008;</i>{{ post.pushAt }}</span>
              <span class="bz" v-if="urlCode == 2"><i>&#xF011;</i>{{ post.goods }}</span>
              <span class="bz" v-if="urlCode == 3"><i>&#xF010;</i>{{ post.views }}</span>
              <span class="bz" v-if="urlCode == 4"><i>&#xF007;</i>{{ post.hots }}</span>
            </div>
          </div>
          <div class="bi">
            <h1 class="bz">📖{{ post.title }}</h1>
            <div class="bit">
              <span class="bz" v-for="t in post.tagNames"><i>#</i>{{ t }}</span>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
    <div class="r" v-if="resData.posts && resData.posts[0]">
      <div class="bcn">
        <button type="button" v-ripples :class="{ nc: !changeSet.pre }" @click="go(0, true)"><i>&#xF003;</i></button>
        <button type="button" v-ripples :class="{ nc: !changeSet.pre }" @click="go(-1, false)"><i>&#xF004;</i></button>
        <p>{{ changeSet.page }}</p>
        <button type="button" v-ripples :class="{ nc: !changeSet.next }" @click="go(1, false)"><i>&#xF005;</i></button>
        <button type="button" v-ripples :class="{ nc: !changeSet.next }" @click="go(changeSet.maxPage, true)"><i>&#xF006;</i></button>
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
const route = useRoute();
const router = useRouter();
const url = route.params.url;
let urlCode: number = 1;
switch (url as string) {
  case "new":
    urlCode = 1;
    break;
  case "good":
    urlCode = 2;
    break;
  case "view":
    urlCode = 3;
    break;
  case "hot":
    urlCode = 4;
    break;
  default: // 其他情况去404
    router.replace({ path: "/404" });
}
const needRun = useState("postsRun", () => true);
const resData = useState("resData", () => <any>{});
const postsMain = useState("postsMain", () => syncPostsMain());
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
    body: { by: urlCode, page: 1 },
  });
  resData.value = (res.data.value as any).data;
  syncPageSet(Math.ceil(resData.value.total / 18));
  initHeader();
  needRun.value = false;
}

// 翻页
let query = { by: urlCode, page: 1 };
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
  syncPageSet(Math.ceil(resData.value.total / 18));
  await nextTick();
  ToWhere(0);
  setTimeout(() => {
    changeDown.value = true;
    if (lazyLoad) {
      LoadPostsImg();
    }
  }, 200);
}
function syncPostsMain() {
  const postsMainArray = [
    { url: "new", tip: "新鲜发布", desc: "围观最新发布的内容", icon: "&#xF008;" },
    { url: "new", tip: "新鲜发布", desc: "围观最新发布的内容", icon: "&#xF008;" },
    { url: "good", tip: "深度好文", desc: "深度价值指数的内容", icon: "&#xF011;" },
    { url: "view", tip: "全站热门", desc: "访问流量最高的内容", icon: "&#xF010;" },
    { url: "hot", tip: "近期上升", desc: "最近备受关注的内容", icon: "&#xF007;" },
  ];
  let mainObj = { main: <any>{}, others: <any>[] };
  for (let i = 1; i < postsMainArray.length; i++) {
    if (i == urlCode) {
      mainObj.main = postsMainArray[i];
    } else {
      mainObj.others.push(postsMainArray[i]);
    }
  }
  return mainObj;
}
function initHeader() {
  useHead({
    title: `${postsMain.value.main.tip} - 文章列表${runtimeConfig.public.siteName}`,
    meta: [{ hid: "description", name: "description", content: `${postsMain.value.main.desc} - 文章列表 - ${runtimeConfig.public.description}` }],
  });
}
// 页面挂载后的初始化
onMounted(async () => {
  // 如果SSR没有
  if (needRun.value) {
    postsMain.value = syncPostsMain();
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
