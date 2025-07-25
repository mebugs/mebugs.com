<template>
  <div id="lmu" class="ll">
    <NuxtLink v-ripples class="llg" href="/">
      <img src="/static/img/logo.png" />
      <h1>MEBUGS</h1>
    </NuxtLink>
    <div class="lmn">
      <NuxtLink v-ripples :class="{ hv: route.path == '/' }" to="/">首页</NuxtLink>
      <NuxtLink v-ripples :class="{ hv: route.path.includes('/posts') }" to="/posts/new">发现</NuxtLink>
      <NuxtLink v-ripples :class="{ hv: route.path.includes('/category') }" to="/category">分类</NuxtLink>
      <NuxtLink v-ripples :class="{ hv: route.path.includes('/tag') }" to="/tag">标签</NuxtLink>
      <NuxtLink v-ripples :class="{ hv: route.path == '/post/about' }" to="/post/about">关于</NuxtLink>
      <NuxtLink v-ripples :class="{ hv: route.path.includes('/page') }" to="/page">更多</NuxtLink>
    </div>
  </div>
  <div class="tp">
    <div class="tbt"><i v-ripples class="bz" @click="openMenu">&#xF031;</i></div>
    <div class="tcen">
      <NuxtLink v-ripples href="/"><img src="/static/img/logo.png" /> <span>MEBUGS</span></NuxtLink>
    </div>
    <div class="tbt"><i v-if="route.path.includes('/post/') && width < 1060" v-ripples @click="openDh" class="bz">&#xF001;</i></div>
  </div>
  <div id="muzz" class="muzz" @click="openMenu"></div>
  <div id="dhzz" class="dhzz" @click="openDh"></div>
</template>

<script setup lang="ts">
const route = useRoute();
const opm = ref(false);
const opd = ref(false);
const width = ref(1061);
// 页面挂载后的初始化
onMounted(async () => {
  width.value = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
});

// 深度监听路由对象变化
watch(
  () => route.path, // 或监听完整路由对象 route
  (newVal, oldVal) => {
    if (process.client) {
      let muzz = document.getElementById("muzz");
      if (muzz?.classList.contains("commzs")) {
        opm.value = true;
        openMenu();
      }
    }
  },
  { deep: true, immediate: true }
);
const openDh = () => {
  let muzz = document.getElementById("dhzz");
  let mull = document.getElementById("bdh");
  if (opd.value) {
    muzz?.classList.remove("commzs");
    mull?.classList.remove("opdh");
    setTimeout(() => {
      muzz?.classList.remove("commtrans");
      muzz?.classList.remove("commidx");
    }, 310);
    opd.value = false;
  } else {
    muzz?.classList.add("commidx");
    setTimeout(() => {
      muzz?.classList.add("commtrans");
      setTimeout(() => {
        muzz?.classList.add("commzs");
        mull?.classList.add("opdh");
      }, 100);
    }, 50);
    opd.value = true;
  }
};
const openMenu = () => {
  let muzz = document.getElementById("muzz");
  let mull = document.getElementById("lmu");
  if (opm.value) {
    muzz?.classList.remove("commzs");
    mull?.classList.remove("opmu");
    setTimeout(() => {
      muzz?.classList.remove("commtrans");
      muzz?.classList.remove("commidx");
    }, 310);
    opm.value = false;
  } else {
    muzz?.classList.add("commidx");
    setTimeout(() => {
      muzz?.classList.add("commtrans");
      setTimeout(() => {
        muzz?.classList.add("commzs");
        mull?.classList.add("opmu");
      }, 100);
    }, 50);
    opm.value = true;
  }
};
</script>
