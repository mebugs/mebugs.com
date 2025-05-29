<template>
  <div class="errTips" :class="{ tipsHas: comm.err }"><i>&#xF029;</i>{{ comm.errTips }}</div>
  <div class="succTips" :class="{ tipsHas: comm.ok }"><i>&#xF030;</i>提交成功，请等待审核通过后全面展示！</div>
  <div class="commz" id="commz" @click="comm.choose ? closeComv() : closeComm()"></div>
  <div class="commp" id="commp">
    <h1 v-if="comm.rid == 0">发表评论</h1>
    <h1 v-else="comm.rid == 0">提交回复</h1>
    <button class="commx" type="button" @click="closeComm()"><i>&#xF021;</i></button>
    <div class="commr">
      <div class="commavt">
        <img :src="'/source/a/' + comm.user.sourceId + '.jpg'" />
        <button type="button" @click="openComv()">更换</button>
      </div>
      <div class="comminf">
        <div class="comminfr">
          <span class="ineed">昵称</span>
          <input type="text" v-model="comm.user.name" placeholder="必填，请输入昵称" />
        </div>
        <div class="comminfr">
          <span>邮箱</span>
          <input type="text" v-model="comm.user.email" placeholder="选填，请输入邮箱" />
        </div>
        <div class="comminfr">
          <span>链接</span>
          <input type="text" v-model="comm.user.url" placeholder="选填，请输入网址" />
        </div>
        <div class="comminfr">
          <span>签名</span>
          <input type="text" v-model="comm.user.summary" placeholder="选填，请输入签名" />
        </div>
        <div class="comminfr commdtl">
          <span class="ineed">评论</span>
          <textarea v-model="comm.info" placeholder="必填,请输入您的精彩观点！" />
        </div>
        <p>温馨提示：系统将通过浏览器临时记忆您曾经填写的个人信息且支持修改，评论提交后仅自己可见，内容需要经过审核后方可全面展示。</p>
      </div>
    </div>
    <button class="commsd" type="button" :class="{ sending: comm.send }" @click="sendComm()">提交</button>
  </div>
  <div class="commp comvp" id="comvp">
    <h1>选择头像</h1>
    <button class="commx" type="button" @click="closeComv()"><i>&#xF021;</i></button>
    <div class="commr comvr">
      <img v-for="item in avts" :src="'/source/a/' + item + '.jpg'" :class="{ avtc: item == comm.user.sourceId }" @click="chooseComv(item)" />
    </div>
  </div>
</template>

<script setup lang="ts">
const avts = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
// 入参读取
const props = defineProps({
  post: {
    type: Object,
    default: () => {
      return {
        postId: 0,
        level: 0,
        rid: 0,
        show: false,
      };
    },
  },
});
watch(props.post, (newV) => {
  let commz = document.getElementById("commz");
  let commp = document.getElementById("commp");
  if (newV.show) {
    commz?.classList.add("commidx");
    setTimeout(() => {
      commz?.classList.add("commtrans");
      setTimeout(() => {
        commz?.classList.add("commzs");
        commp?.classList.add("commolk");
      }, 100);
    }, 50);
    initUser();
  }
});
const comm = ref({
  user: { clientId: "", name: "", url: "", summary: "", sourceId: 0, email: "" },
  postId: 0,
  level: 0,
  rid: 0,
  info: "",
  choose: false,
  ok: false,
  err: false,
  errTips: "",
  send: false,
});
function initUser() {
  // 获取客户端ID
  let clientInfo = window.localStorage.getItem("MEBUGS_COMM_USER");
  if (clientInfo) {
    comm.value.user = JSON.parse(clientInfo);
  } else {
    comm.value.user.clientId = generateRandomString();
    comm.value.user.sourceId = Math.floor(Math.random() * 12) + 1;
    window.localStorage.setItem("MEBUGS_COMM_CLIENTID", comm.value.user.clientId);
    window.localStorage.setItem("MEBUGS_COMM_USER", JSON.stringify(comm.value.user));
  }
  comm.value.postId = props.post.postId;
  comm.value.level = props.post.level;
  comm.value.rid = props.post.rid;
  comm.value.info = "";
}
// 提交评论
async function sendComm() {
  if (comm.value.send) {
    return;
  }
  comm.value.send = true;
  comm.value.errTips = "";
  if (comm.value.user.name == "") {
    comm.value.errTips = "请输入昵称";
  }
  if (comm.value.info == "") {
    comm.value.errTips = "请输入评论";
  }
  if (comm.value.errTips != "") {
    comm.value.err = true;
    setTimeout(() => {
      comm.value.err = false;
      comm.value.send = false;
    }, 2000);
    return;
  }
  window.localStorage.setItem("MEBUGS_COMM_USER", JSON.stringify(comm.value.user));
  // 提交评论
  const res = await $fetch("/api/page/comm/add", { server: false, method: "POST", body: comm.value });
  const data = res as any;
  if (data) {
    if (data.code != "S000") {
      comm.value.errTips = data.msg;
    } else {
      comm.value.ok = true;
    }
  } else {
    comm.value.errTips = "系统异常";
  }
  if (comm.value.errTips != "") {
    comm.value.err = true;
    setTimeout(() => {
      comm.value.err = false;
      comm.value.send = false;
    }, 2000);
    return;
  }
  if (comm.value.ok) {
    setTimeout(() => {
      comm.value.ok = false;
      comm.value.send = false;
    }, 2000);
    // 关闭评论框
    comm.value.info = "";
    closeComm();
  }
}
function closeComm() {
  window.localStorage.setItem("MEBUGS_COMM_USER", JSON.stringify(comm.value.user));
  props.post.show = false;
  let commz = document.getElementById("commz");
  let commp = document.getElementById("commp");
  commz?.classList.remove("commzs");
  commp?.classList.remove("commolk");
  setTimeout(() => {
    commz?.classList.remove("commtrans");
    commz?.classList.remove("commidx");
  }, 310);
}
function openComv() {
  comm.value.choose = true;
  let comvp = document.getElementById("comvp");
  comvp?.classList.add("commpa");
}
function closeComv() {
  comm.value.choose = false;
  let comvp = document.getElementById("comvp");
  comvp?.classList.remove("commpa");
}
function chooseComv(i: number) {
  comm.value.user.sourceId = i;
  window.localStorage.setItem("MEBUGS_COMM_USER", JSON.stringify(comm.value.user));
  closeComv();
}
function generateRandomString() {
  const characters = "abcdefghijklmnopqrstuvwxyz0123456789";
  let result = "";
  const charactersLength = characters.length;
  for (let i = 0; i < 64; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}
// 页面挂载后的初始化
onMounted(async () => {
  if (process.client) {
    initUser();
  }
});
</script>
