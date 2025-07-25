// plugins/baidu.client.ts
export default defineNuxtPlugin(() => {
  if (process.client) {
    // 加载百度统计脚本
    const script = document.createElement("script");
    script.src = "https://hm.baidu.com/hm.js?64b27967e9d2fa5d1d4c6f945051603e";
    script.async = true;
    document.head.appendChild(script);

    // 监听路由变化
    const router = useRouter();
    router.afterEach((to) => {
      console.log("Try baidu");
      // @ts-ignore - baidu
      window._hmt = window._hmt || [];
      // @ts-ignore - baidu
      window._hmt.push(["_trackPageview", to.fullPath]);
    });
  }
});
