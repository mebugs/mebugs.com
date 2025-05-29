<template>
  <div class="la">
    <NuxtRouteAnnouncer />
    <NuxtLoadingIndicator :color="'#3491fa'" :height="2" />
    <Menu></Menu>
    <div class="lr"><NuxtPage /></div>
    <Comm></Comm>
  </div>
</template>

<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();
useHead({
  link: [{ rel: "stylesheet", href: "/static/css/base.css" }],
  // script: [{ src: "/static/js/base.js", tagPosition: "bodyClose" }],
});
</script>
