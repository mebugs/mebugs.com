// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  devServer: {
    host: "0.0.0.0", // 开放所有网络接口访问
    port: 3000, // 可自定义端口
  },
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
      siteName: " - 米虫博客 - www.mebugs.com",
      description: "米虫博客，程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！",
      backServer: "http://localhost:8000",
      apiBase: "/api",
    },
  },
  plugins: [
    { src: "~/plugins/ripples.js", ssr: false }, // ssr: false 表示仅客户端生效
    // { src: "~/plugins/adsense.client.ts", mode: "client" },
    { src: "~/plugins/baidu.client.ts", ssr: false },
  ],
  // google ads
  modules: ["@nuxtjs/google-adsense"],
  googleAdsense: {
    id: "ca-pub-2699978133125218",
    adFormat: "auto",
    onPageLoad: false,
    pageLevelAds: false,
    overlayBottom: true,
    test: process.env.NODE_ENV != "production",
  },
  experimental: {
    payloadExtraction: false, // 禁用预加载提示（间接关闭data-hid）
    resourceHints: false, // 显式禁用资源提示
  },
  build: {
    transpile: ["swiper"],
  },
});
