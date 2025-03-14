// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  app: {
    head: {
      charset: "utf-8",
      viewport: "width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no",
    },
  },
  nitro: {
    devProxy: {
      "/api": { target: "http://localhost:8000", changeOrigin: true },
    },
  },
  runtimeConfig: {
    public: {
      siteName: "米虫博客",
      backServer: "http://localhost:8000",
      apiBase: "/api",
    },
  },
});
