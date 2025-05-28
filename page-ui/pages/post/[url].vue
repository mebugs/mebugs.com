<template>
  <div class="m post" v-if="resData.post">
    <div class="r">
      <div class="rbx">
        <div class="bx b1">
          <h1 class="bz bpt">📖 {{ resData.post.title }}</h1>
          <div class="bi bix">
            <div class="bip bz">
              <em>发布：</em><span><i>&#xF008;</i> {{ resData.post.pushAt }}</span>
            </div>
            <div class="bip bz">
              <em>热度：</em><span><i>&#xF010;</i> {{ resData.post.views }}</span>
            </div>
            <div class="bip bz">
              <em>趋势：</em><span><i>&#xF007;</i> {{ resData.post.hots }}</span>
            </div>
            <div class="bip bz">
              <em>权重：</em><span><i>&#xF011;</i> {{ resData.post.goods }}</span>
            </div>
            <div class="bip bz">
              <em>分类：</em>
              <NuxtLink :to="'/category/' + resData.post.category"><i>&#xF012;</i> {{ resData.post.categoryName }}</NuxtLink>
            </div>
            <div class="bip bz" v-if="resData.post.tagUrls && resData.post.tagUrls[0]">
              <em>标签：</em>
              <NuxtLink v-for="(ta, ti) in resData.post.tagUrls" :to="'/tag/' + ta"><i>&#xF013;</i> {{ resData.post.tagNames[ti] }}</NuxtLink>
            </div>
          </div>
        </div>
        <div class="bx bl">
          <div class="bg bpm">
            <img :src="resData.post.sourcePath + '.' + resData.post.sourceBack" />
          </div>
          <div class="pdesc bz">
            <span>🎯</span>
            <p>{{ resData.post.summary }}</p>
          </div>
          <em id="st"></em>
          <div class="vditor-reset bz" v-html="resData.postMain.html"></div>
          <em id="ed"></em>
          <div class="b pcom">
            <button type="button" @click="putComms()"><i>&#xF009;</i>评论</button>
          </div>
          <div class="pcoms">
            <div class="pcomt bz" v-if="comms.total">
              当前文章有<b>{{ comms.total }}</b
              >条讨论，加入讨论吧!
            </div>
            <div class="pcomt bz" v-else>当前文章暂无讨论，留下脚印吧！</div>
            <div class="pcomls" v-if="comms.total && comms.comments && comms.comments[0]">
              <div class="pcomll" v-for="item in comms.comments">
                <div class="pcomavt"><img :src="'/source/a/' + item.user.sourceId + '.jpg'" /></div>
                <div class="pcominf">
                  <div class="pcomina">
                    <h2>
                      <a :href="item.user.url ? item.user.url : '#'" target="_blank">{{ item.user.name }}</a>
                      <em>{{ item.date }}</em>
                    </h2>
                    <p>{{ item.user.summary }}</p>
                    <button type="button" @click="putCommsRes(item.id)"><i>&#xF009;</i></button>
                  </div>
                  <div class="pcominp">
                    <p>{{ item.info }}</p>
                  </div>
                </div>
                <div class="pcomll" v-for="it in item.comments">
                  <div class="pcomavt"><img :src="'/source/a/' + it.user.sourceId + '.jpg'" /></div>
                  <div class="pcominf">
                    <div class="pcomina">
                      <h2>
                        <a :href="it.user.url ? it.user.url : '#'" target="_blank">{{ it.user.name }}</a>
                        <em>{{ it.date }}</em>
                      </h2>
                      <p>{{ it.user.summary }}</p>
                      <button type="button" @click="putCommsRes(it.id)"><i>&#xF009;</i></button>
                    </div>
                    <div class="pcominp">
                      <p>{{ it.info }}</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="bx br brl">
          <div class="brdh bz" v-html="resData.postMain.toc"></div>
        </div>
      </div>
    </div>
    <div class="r rpm">
      <div class="rp">
        <h1 class="bz"><i>&#xF023;</i> 相关阅读</h1>
      </div>
      <div class="rb" v-if="resData.postMore && resData.postMore[0]">
        <NuxtLink class="b b3" v-for="post in resData.postMore" :to="'/post/' + post.url">
          <div class="bg">
            <img :src="post.sourcePath + '_2.' + post.sourceBack" />
          </div>
          <div class="bi">
            <h1 class="bz">{{ post.title }}</h1>
          </div>
        </NuxtLink>
      </div>
    </div>
    <Foot></Foot>
  </div>
  <Comns :post="openComms" />
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
    title: `${resData.value?.post?.title ? resData.value?.post?.title : "哎呀呀！没找到！"}${runtimeConfig.public.siteName}`,
    meta: [{ hid: "description", name: "description", content: `${resData.value?.post?.summary ? resData.value?.post?.summary : "哎呀呀！没找到！"} - ${runtimeConfig.public.description}` }],
    link: [{ href: "/static/css/ant-design.css", rel: "stylesheet" }],
    script: [{ src: "/static/js/post.js", tagPosition: "bodyClose" }],
  });
}
// 评论数据
const comms = ref({ total: 0, comments: <any>[] });
async function getCommnents() {
  let clientId = window.localStorage.getItem("MEBUGS_COMM_CLIENTID");
  if (!clientId) {
    clientId = "";
  }
  const commRes = await $fetch("/api/page/comments", { server: false, method: "POST", body: { postId: resData.value.post.id, clientId: clientId } });
  comms.value = (commRes as any).data;
}
// 提交评论
const openComms = ref({ postId: 0, level: 0, rid: 0, show: false });
function putComms() {
  openComms.value.postId = resData.value.post.id;
  openComms.value.rid = 0;
  openComms.value.level = 0;
  openComms.value.show = true;
}
function putCommsRes(rid: number) {
  openComms.value.postId = resData.value.post.id;
  openComms.value.rid = rid;
  openComms.value.level = 1;
  openComms.value.show = true;
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
var intv: NodeJS.Timeout;
var waitPostLoad: NodeJS.Timeout;
// 执行深度处理
function initPageGoods() {
  var nums: number = 1;
  intv = setInterval(() => {
    if (nums <= 6) {
      $fetch("/api/page/post/good", { server: false, method: "POST", body: { url: url } });
      nums++;
    } else {
      clearInterval(intv);
    }
  }, 60000);
}
onUnmounted(() => {
  needRun.value = true;
  clearInterval(intv);
  clearInterval(waitPostLoad);
  // @ts-ignore
  InitBack();
});
function initPage() {
  // @ts-ignore
  InitDom();
  initPageGoods();
  getCommnents();
  initPost();
}
function initPost() {
  waitPostLoad = setInterval(() => {
    // @ts-ignore
    if (typeof mdMiHs !== "undefined") {
      // @ts-ignore
      InitPostDom();
      clearInterval(waitPostLoad);
    }
  }, 100);
}
</script>
