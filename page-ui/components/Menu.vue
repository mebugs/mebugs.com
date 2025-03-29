<template>
  <div class="ll">
    <a class="llg" href="/">
      <img src="/static/img/logo.png" />
      <h1>MEBUGS</h1>
    </a>
    <div class="lmn">
      <NuxtLink :class="{ hv: route.path == '/' }" to="/">首页</NuxtLink>
      <NuxtLink :class="{ hv: route.path.includes('/posts') }" to="/posts/new">发现</NuxtLink>
      <NuxtLink :class="{ hv: route.path.includes('/topic') }" to="/topic">专题</NuxtLink>
      <NuxtLink :class="{ hv: route.path.includes('/category') }" to="/category">分类</NuxtLink>
      <NuxtLink :class="{ hv: route.path.includes('/tag') }" to="/tag">标签</NuxtLink>
      <NuxtLink :class="{ hv: route.path == '/post/about' }" to="/post/about">关于</NuxtLink>
      <NuxtLink :class="{ hv: route.path.includes('/page') }" to="/page">更多</NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
const route = useRoute();
</script>
