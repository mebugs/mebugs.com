<template>
  <div class="m">
    <div class="r bn swiper-container" v-if="banners && banners[0]">
      <div class="swiper-wrapper">
        <div class="swiper-slide" v-for="bann in banners">
          <div class="b">
            <div class="bg bgb"><img :src="bann.sourceShow" /></div>
            <div class="bi bib">
              <h1 class="bz">{{ bann.title }}</h1>
              <div class="bit">
                <span class="bz"><i>#</i>{{ bann.tag }}</span>
              </div>
              <p class="bz">{{ bann.summary }}</p>
              <div class="bia"><a class="bz" :href="bann.url">点击前往</a></div>
            </div>
          </div>
        </div>
      </div>
      <!-- Add Pagination -->
      <div class="swiper-pagination"></div>
    </div>
    <div class="r">
      <div class="rb rt">
        <a href="/" class="b b4">
          <div class="bg"><i class="bz">&#xF007;</i></div>
          <div class="bi">
            <h2 class="bz">近期上升</h2>
            <div class="bit"><span class="bz">最近备受关注的内容</span></div>
          </div>
        </a>
        <a href="/" class="b b4">
          <div class="bg"><i class="bz">&#xF010;</i></div>
          <div class="bi">
            <h2 class="bz">全站热门</h2>
            <div class="bit"><span class="bz">访问流量最高的内容</span></div>
          </div>
        </a>
        <a href="/" class="b b4">
          <div class="bg"><i class="bz">&#xF011;</i></div>
          <div class="bi">
            <h2 class="bz">深度好文</h2>
            <div class="bit"><span class="bz">深度价值指数的内容</span></div>
          </div>
        </a>
        <a href="/" class="b b4">
          <div class="bg"><i class="bz">&#xF008;</i></div>
          <div class="bi">
            <h2 class="bz">新鲜发布</h2>
            <div class="bit"><span class="bz">围观最新发布的内容</span></div>
          </div>
        </a>
      </div>
    </div>
    <div class="r">
      <div class="rp">
        <h1 class="bz"><i>&#xF007;</i> 近期上升</h1>
      </div>
      <div class="rb ro" v-if="postsHot && postsHot[0]">
        <a class="b b3" v-for="post in postsHot" :href="'/post/' + post.url">
          <div class="bg">
            <img :src="post.sourcePath + '_1.' + post.sourceBack" />
            <div class="bgt">
              <span class="bz"><i>&#xF013;</i>{{ post.categoryName }}</span>
              <span class="bz"><i>&#xF007;</i>{{ post.hots }}</span>
            </div>
          </div>
          <div class="bi">
            <h1 class="bz">{{ post.title }}</h1>
            <div class="bit">
              <span class="bz" v-for="t in post.tagNames"><i>#</i>{{ t }}</span>
            </div>
          </div>
        </a>
      </div>
    </div>
    <div class="r">
      <div class="rp">
        <h1><i>&#xF012;</i> 文章分类</h1>
      </div>
      <div class="rb rt rtg" v-if="category && category[0]">
        <a v-for="cat in category" :href="'/category/' + cat.url" class="b b4">
          <div class="bg"><img :src="cat.sourceShow" /></div>
          <div class="bi">
            <span class="bno bz">{{ cat.num }}</span>
            <h2 class="bz">{{ cat.title }}</h2>
            <div class="bit">
              <span class="bz">{{ cat.summary }}</span>
            </div>
          </div>
        </a>
      </div>
    </div>
    <div class="r">
      <div class="rp">
        <h1><i>&#xF010;</i> 全站热门</h1>
      </div>
      <div class="rb ro" v-if="postsView && postsView[0]">
        <a class="b b3" v-for="post in postsView" :href="'/post/' + post.url">
          <div class="bg">
            <img :src="post.sourcePath + '_1.' + post.sourceBack" />
            <div class="bgt">
              <span class="bz"><i>&#xF013;</i>{{ post.categoryName }}</span>
              <span class="bz"><i>&#xF010;</i>{{ post.views }}</span>
            </div>
          </div>
          <div class="bi">
            <h1 class="bz">{{ post.title }}</h1>
            <div class="bit">
              <span class="bz" v-for="t in post.tagNames"><i>#</i>{{ t }}</span>
            </div>
          </div>
        </a>
      </div>
    </div>
    <div class="r">
      <div class="rp">
        <h1><i>&#xF014;</i> 专题专栏</h1>
      </div>
      <div class="rb rc" v-if="topic && topic[0]">
        <a v-for="to in topic" class="b b4" href="/">
          <div class="bg">
            <img :src="to.sourceShow" />
            <div class="bgt">
              <span class="bz">{{ to.num }}</span>
            </div>
          </div>
        </a>
      </div>
    </div>
    <div class="r">
      <div class="rp">
        <h1><i>&#xF011;</i> 深度好文</h1>
      </div>
      <div class="rb ro" v-if="postsGood && postsGood[0]">
        <a class="b b3" v-for="post in postsGood" :href="'/post/' + post.url">
          <div class="bg">
            <img :src="post.sourcePath + '_1.' + post.sourceBack" />
            <div class="bgt">
              <span class="bz"><i>&#xF013;</i>{{ post.categoryName }}</span>
              <span class="bz"><i>&#xF011;</i>{{ post.goods }}</span>
            </div>
          </div>
          <div class="bi">
            <h1 class="bz">{{ post.title }}</h1>
            <div class="bit">
              <span class="bz" v-for="t in post.tagNames"><i>#</i>{{ t }}</span>
            </div>
          </div>
        </a>
      </div>
    </div>
    <div class="r">
      <div class="rp">
        <h1><i>&#xF013;</i> 核心标签</h1>
      </div>
      <div class="rb rtt" v-if="tag && tag[0]">
        <a v-for="ta in tag" :href="'/tag/' + ta.url" class="b b5">
          <div class="bg"><img :src="ta.sourceShow" /></div>
          <div class="bi">
            <h2 class="bz">{{ ta.title }}</h2>
          </div>
        </a>
      </div>
    </div>
    <div class="r">
      <div class="rp">
        <h1><i>&#xF008;</i> 新鲜发布</h1>
      </div>
      <div class="rb ro" v-if="postsNew && postsNew[0]">
        <a class="b b3" v-for="post in postsNew" :href="'/post/' + post.url">
          <div class="bg">
            <img :src="post.sourcePath + '_1.' + post.sourceBack" />
            <div class="bgt">
              <span class="bz"><i>&#xF013;</i>{{ post.categoryName }}</span>
              <span class="bz"><i>&#xF008;</i>{{ post.pushAt }}</span>
            </div>
          </div>
          <div class="bi">
            <h1 class="bz">{{ post.title }}</h1>
            <div class="bit">
              <span class="bz" v-for="t in post.tagNames"><i>#</i>{{ t }}</span>
            </div>
          </div>
        </a>
      </div>
    </div>
    <div class="r">
      <a class="b bm" href="/post">查看更多<em></em></a>
    </div>
    <Foot></Foot>
  </div>
</template>
<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();
useHead({
  link: [{ href: "https://cdn.staticfile.net/Swiper/11.0.5/swiper-bundle.min.css", rel: "stylesheet" }],
  script: [
    { src: "https://cdn.staticfile.net/Swiper/11.0.5/swiper-bundle.min.js", tagPosition: "bodyClose" },
    { src: "/static/js/index.js", tagPosition: "bodyClose" },
  ],
});

// 读取页面数据
const { data } = await useFetch(runtimeConfig.public.backServer + "/page/index", {});
const resData = (data.value as any).data;
const { banners, category, tag, topic, postsHot, postsView, postsGood, postsNew } = resData;
// let data2 = ref({});
// async function handleFormSubmit() {
//   data2.value = await $fetch("/api/page/index");
// }
</script>
