<template>
  <div class="m">
    <div class="r">
      <div class="bcn">
        <div class="totp">
          <div class="bg">
            <i class="bz">&#xF012;</i>
          </div>
          <div class="bi">
            <h2 class="bz">文章分类</h2>
          </div>
        </div>
      </div>
    </div>
    <!--广告位 -->
    <div class="addeare"><Adsbygoogle ad-slot="6001361550" /></div>
    <div class="r">
      <div class="rb rt rtgb" v-if="resData.category && resData.category[0]">
        <NuxtLink v-for="cat in resData.category" :to="'/category/' + cat.url" class="b b3">
          <div class="bg"><img :src="cat.sourceShow" /></div>
          <div class="bi">
            <span class="bno bz">{{ cat.num }}</span>
            <h2 class="bz">{{ cat.title }}</h2>
            <div class="bit">
              <span class="bz">{{ cat.summary }}</span>
            </div>
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
const needRun = useState("categorysRun", () => true);
const resData = useState("resData", () => <any>{});
useHead({
  title: `文章分类${runtimeConfig.public.siteName}`,
  meta: [{ hid: "description", name: "description", content: `文章分类 - ${runtimeConfig.public.description}` }],
});
// 读取页面数据
if (process.server) {
  const res = await useFetch(runtimeConfig.public.backServer + "/page/categoryList", {});
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
  const res = await $fetch("/api/page/categoryList", { server: false, method: "GET" });
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
