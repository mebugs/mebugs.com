<template>
  <div class="tab-bar-container">
    <div class="tab-bar-btn tab-bar-left" :class="{ tabBarLShow: tagBL }" @click="leftTab()">
      <icon-left />
    </div>
    <div class="tab-bar-contant" :style="{ left: tagRN + 'px' }">
      <tab-item v-for="(tag, index) in tagList" :key="tag.fullPath" :index="index" :item-data="tag" @set-tag-run="setTagRun" />
    </div>
    <div class="tab-bar-btn tab-bar-right" :class="{ tabBarRShow: tagBR }" @click="rightTab()">
      <icon-right />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, watch, onUnmounted, onMounted, nextTick } from 'vue';
import type { RouteLocationNormalized } from 'vue-router';
import { listenerRouteChange, removeRouteListener } from '@/utils/route-listener';
import { useAppStore, useTabBarStore } from '@/store';
// import { nextTick } from 'process';
import tabItem from './tab-item.vue';

const appStore = useAppStore();
const tabBarStore = useTabBarStore();

const affixRef = ref();
// 标签列表
const tagList = computed(() => {
  return tabBarStore.getTabList;
});

watch(
  () => appStore.navbar,
  () => {
    affixRef.value.updatePosition();
  }
);

// 标签总长度
const tagL = ref(0);
// 标签区域长度
const tagEL = ref(0);
// 左侧区按钮
const tagBL = ref(false);
const tagBR = ref(false);
// 位移值
const tagRN = ref(0);
// 左右按钮判定
function leftRight() {
  // 左按钮判定
  tagBL.value = tagRN.value < 0;
  // 判定右按钮
  tagBR.value = tagL.value > tagEL.value;
  if (tagBR.value) {
    tagBR.value = tagL.value > tagEL.value - tagRN.value;
  }
}
// 标签各项计算
function setTagRunL(path: string) {
  // 滚动区
  const tabCt = document.querySelector('.tab-bar-contant');
  // 选中签现在的位置
  let checkN = 0;
  let changeUnit = 0;
  if (tabCt) {
    tagEL.value = tabCt.clientWidth;
    const tabBi = document.querySelectorAll('.tab-bar-item');
    tagL.value = 0;
    if (tabBi && tabBi.length > 0) {
      for (let i = 0; i < tagList.value.length; i += 1) {
        const unitLength = tabBi[i].clientWidth + 8;
        tagL.value += unitLength;
        if (path !== '' && tagList.value[i].fullPath === path) {
          checkN = (tabBi[i] as HTMLElement).offsetLeft - 4;
          changeUnit = unitLength;
        }
      }
    }
    // 基于checkN判定是否要位移
    // 位置值尾部超出
    const tagElth = tagEL.value;
    let waitRun = 0;
    if (tagL.value > tagElth) {
      // 尾部值
      const endLine = checkN + changeUnit;
      if (endLine > tagElth) {
        waitRun = tagElth - endLine;
        // 头部值-偏移值<0 说明出现在左边
      } else if (checkN + tagRN.value < 0) {
        // 位置头部超出
        waitRun = -checkN;
      }
      // 启动滚动
      tagRN.value = waitRun;
    } else {
      tagRN.value = 0;
    }
    leftRight();
  }
}

// 窗口变化
const tagRe = ref(true);
function tagResize() {
  if (tagRe.value) {
    tagRe.value = false;
    setTagRunL('');
    tagRe.value = true;
  }
}
// 左切
function leftTab() {
  if (tagRN.value < 0) {
    const waitRun = tagRN.value + 150;
    tagRN.value = waitRun > 0 ? 0 : waitRun;
  } else {
    tagRN.value = 0;
  }
  leftRight();
}
// 右切
function rightTab() {
  const maxRang = tagEL.value - tagL.value - 20;
  if (tagRN.value > maxRang) {
    const waitRun = tagRN.value - 150;
    tagRN.value = waitRun < maxRang ? maxRang : waitRun;
  } else {
    tagRN.value = maxRang;
  }
  leftRight();
}
// 开放给子组件
const setTagRun = () => {
  setTagRunL('');
};
// 路由切换
listenerRouteChange((route: RouteLocationNormalized) => {
  if (!route.meta.noAffix && !tagList.value.some((tag) => tag.fullPath === route.fullPath)) {
    tabBarStore.updateTabList(route);
  }
  // 目标路由可能产生移位
  nextTick(() => {
    setTagRunL(route.fullPath);
  });
}, true);
// 标签初始化
onMounted(() => {
  window.addEventListener('resize', tagResize);
  setTagRun();
});

// 注销监听
onUnmounted(() => {
  window.removeEventListener('resize', tagResize);
  removeRouteListener();
});
</script>

<style scoped lang="less">
.tab-bar-container {
  position: relative;
  background-color: var(--color-bg-2);
  height: 34px;
  width: 100%;
  overflow: hidden;
  * {
    transition: all 0.3s linear;
  }
  .tab-bar-btn {
    width: 20px;
    height: 34px;
    position: absolute;
    background-color: var(--color-bg-2);
    box-shadow: 0 0 10px var(--color-neutral-5);
    color: var(--color-text-2);
    top: 0;
    z-index: 1;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .tab-bar-btn:hover {
    color: rgb(var(--link-6));
    box-shadow: 0 0 10px rgb(var(--link-6));
  }
  .tab-bar-left {
    left: -30px;
  }
  .tabBarLShow {
    left: 0;
  }
  .tab-bar-right {
    right: -30px;
  }
  .tabBarRShow {
    right: 0;
  }
  .tab-bar-contant {
    margin: 0 20px;
    position: relative;
    white-space: nowrap;
    left: 0;
  }
}
</style>
