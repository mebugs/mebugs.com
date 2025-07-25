// plugins/error-handler.ts
export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.config.errorHandler = (error, instance, info) => {
    console.error("Golobal Vue Err:", error);
  };
  nuxtApp.hook("app:error", (err) => {
    console.error("App Err:", err); // 捕获初始化阶段的崩溃
  });
});
