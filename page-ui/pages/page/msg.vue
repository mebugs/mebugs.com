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
          <div class="b pcom">
            <button type="button" v-ripples @click="putComms()"><i>&#xF009;</i>评论</button>
          </div>
          <div class="pcoms">
            <div class="pcomt bz" v-if="comms.total">
              留言板有<b>{{ comms.total }}</b
              >条讨论，加入讨论吧!
            </div>
            <div class="pcomt bz" v-else>留言板暂无讨论，留下脚印吧！</div>
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
      </div>
    </div>
    <Foot></Foot>
  </div>
  <Comns :post="openComms" />
</template>

<script setup lang="ts">
import { InitDom, InitBack } from "~/utils/base.js";
const runtimeConfig = useRuntimeConfig();
const needRun = useState("pageMsg", () => true);
const resData = useState("resData", () => <any>{});
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/page", { method: "POST", body: { url: "/page/msg" } });
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
// 评论数据
const comms = ref({ total: 0, comments: <any>[] });
async function getCommnents() {
  let clientId = window.localStorage.getItem("MEBUGS_COMM_CLIENTID");
  if (!clientId) {
    clientId = "";
  }
  const commRes = await $fetch("/api/page/comments", { server: false, method: "POST", body: { postId: 0, clientId: clientId } });
  comms.value = (commRes as any).data;
}
// 提交评论
const openComms = ref({ postId: 0, level: 0, rid: 0, show: false });
function putComms() {
  openComms.value.postId = 0;
  openComms.value.rid = 0;
  openComms.value.level = 0;
  openComms.value.show = true;
}
function putCommsRes(rid: number) {
  openComms.value.postId = 0;
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
  getCommnents();
  initPage();
});
async function initByClient() {
  const res = await $fetch("/api/page/page", { server: false, method: "POST", body: { url: "/page/msg" } });
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
