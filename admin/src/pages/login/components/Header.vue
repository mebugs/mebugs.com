<template>
  <header class="login-header">
    <logo-full-icon class="logo" />
    <div class="operations-container">
      <t-button theme="default" shape="square" variant="text" @click="navToGitHub">
        <t-icon name="logo-github" class="icon" />
      </t-button>
      <t-button theme="default" shape="square" variant="text" @click="navToHelper">
        <t-icon name="help-circle" class="icon" />
      </t-button>
      <t-button theme="default" shape="square" variant="text" @click="changeMode">
        <icon v-if="settingStore.displayMode == 'dark'" size="26px" name="darkMode" url="/static/svg/mode.js" />
        <icon v-else size="26px" name="lightMode" url="/static/svg/mode.js" />
      </t-button>
    </div>
  </header>
</template>

<script setup lang="ts">
import { Icon } from 'tdesign-icons-vue-next';
import LogoFullIcon from '@/assets/assets-logo-full.svg?component';
import { useSettingStore } from '@/store';

const settingStore = useSettingStore();
// 切换视觉模式
const changeMode = () => {
  const modeNew = settingStore.displayMode === 'light' ? 'dark' : 'light';
  settingStore.updateConfig({
    mode: modeNew,
  });
};

const navToGitHub = () => {
  window.open('https://github.com/tencent/tdesign-vue-next-starter');
};

const navToHelper = () => {
  window.open('http://tdesign.tencent.com/starter/docs/get-started');
};
</script>

<style lang="less" scoped>
.login-header {
  padding: 12px 28px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  backdrop-filter: blur(5px);
  color: var(--td-text-color-primary);

  .logo {
    width: 188px;
  }

  .operations-container {
    display: flex;
    align-items: center;
    .t-button {
      margin-left: 16px;
    }

    .icon {
      height: 20px;
      width: 20px;
      padding: 6px;
      box-sizing: content-box;

      &:hover {
        cursor: pointer;
      }
    }
  }
}
</style>
