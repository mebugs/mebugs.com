<template>
  <div class="errTips" :class="{ tipsHas: link.err }"><i>&#xF029;</i>{{ link.errTips }}</div>
  <div class="succTips" :class="{ tipsHas: link.ok }"><i>&#xF030;</i>{{ link.succTips }}</div>
  <div class="commz" id="commz" @click="closeLink()"></div>
  <div class="commp commlink" id="commlk">
    <h1>提交友链</h1>
    <button class="commx" type="button" @click="closeLink()"><i>&#xF021;</i></button>
    <label class="commavt" for="logoAdd">
      <img :src="link.sourceShow" v-if="link.sourceShow" />
    </label>
    <p>
      你可点击上方直接上传LOGO，也可以输入链接后自动同步贵站的数据。<br />
      链接为完整地址，不要涉及重定向，防盗链可能导致LOGO获取失败。<br />
      链接示例：https://www.mebugs.com
    </p>
    <input id="logoAdd" accept="image/gif, image/jpeg, image/png, image/jpg" type="file" style="display: none" @change="chooesImg" />
    <div class="comminf">
      <div class="comminfr syncinfr">
        <span class="ineed">链接</span>
        <input type="text" v-model="link.url" placeholder="必填，请输入网站地址" />
        <button type="button" :class="{ sending: link.send }" @click="syncLink()">同步</button>
      </div>
      <div class="comminfr">
        <span class="ineed">名称</span>
        <input type="text" v-model="link.title" placeholder="必填，请输入名称，建议8字以内" />
      </div>
      <div class="comminfr commdtl">
        <span>简介</span>
        <textarea v-model="link.summary" placeholder="选填,请输入站点简介，建议64字以内" />
      </div>
    </div>
    <button class="commsd" type="button" :class="{ sending: link.send }" @click="sendLink()">提交</button>
  </div>
</template>

<script setup lang="ts">
// 入参读取
const props = defineProps({
  linkAdd: {
    type: Object,
    default: () => {
      return {
        show: false,
      };
    },
  },
});
watch(props.linkAdd, () => {
  let commz = document.getElementById("commz");
  let commlk = document.getElementById("commlk");
  if (props.linkAdd.show) {
    commz?.classList.add("commidx");
    setTimeout(() => {
      commz?.classList.add("commtrans");
      setTimeout(() => {
        commz?.classList.add("commzs");
        commlk?.classList.add("commolk");
      }, 100);
    }, 50);
    initUser();
  }
});
const user = ref({ clientId: "", name: "", url: "", summary: "", sourceId: 0, email: "" });
const link = ref<any>({});
link.value = initLink();
function initLink() {
  return {
    clientId: "",
    source: [],
    sourceShow: "/static/img/logo.png",
    url: "",
    title: "",
    summary: "",
    ok: false,
    err: false,
    errTips: "",
    succTips: "提交成功，请等待审核通过后展示！",
    send: false,
  };
}
function initUser() {
  // 获取客户端ID
  let clientInfo = window.localStorage.getItem("MEBUGS_COMM_USER");
  if (clientInfo) {
    user.value = JSON.parse(clientInfo);
  } else {
    user.value.clientId = generateRandomString();
    user.value.sourceId = Math.floor(Math.random() * 12) + 1;
    window.localStorage.setItem("MEBUGS_COMM_CLIENTID", user.value.clientId);
    window.localStorage.setItem("MEBUGS_COMM_USER", JSON.stringify(user.value));
  }
  link.value = initLink();
  link.value.clientId = user.value.clientId;
}
// 提交同步
async function syncLink() {
  if (link.value.send) {
    return;
  }
  link.value.send = true;
  link.value.errTips = "";
  if (link.value.url == "") {
    link.value.errTips = "请输入站点链接";
  }
  if (link.value.errTips != "") {
    link.value.err = true;
    setTimeout(() => {
      link.value.err = false;
      link.value.send = false;
    }, 2000);
    return;
  }
  // 提交链接
  const res = await $fetch("/api/page/link/scan", { server: false, method: "POST", body: { url: link.value.url } });
  const data = res as any;
  if (data) {
    if (data.code != "S000") {
      link.value.errTips = data.msg;
    } else {
      link.value.ok = true;
      link.value.title = data.data.title;
      link.value.summary = data.data.summary;
      link.value.source = [data.data.sourceShow];
      link.value.sourceShow = [data.data.sourceShow];
      link.value.succTips = "同步完成";
    }
  } else {
    link.value.errTips = "系统异常";
  }
  if (link.value.errTips != "") {
    link.value.err = true;
    setTimeout(() => {
      link.value.err = false;
      link.value.send = false;
    }, 2000);
    return;
  }
  if (link.value.ok) {
    setTimeout(() => {
      link.value.ok = false;
      link.value.send = false;
    }, 2000);
  }
}
// 提交评论
async function sendLink() {
  if (link.value.send) {
    return;
  }
  link.value.send = true;
  link.value.errTips = "";
  if (link.value.url == "") {
    link.value.errTips = "请输入站点链接";
  }
  if (link.value.title == "") {
    link.value.errTips = "请输入站点标题";
  }
  if (link.value.source.url == 0) {
    link.value.errTips = "请提交站点Logo";
  }
  if (Array.from(link.value.title).length > 128) {
    link.value.errTips = "站点链接最大128位";
  }
  if (Array.from(link.value.title).length > 32) {
    link.value.errTips = "站点标题最大32位";
  }
  if (Array.from(link.value.summary).length > 256) {
    link.value.errTips = "站点描述最大256位";
  }
  if (link.value.errTips != "") {
    link.value.err = true;
    setTimeout(() => {
      link.value.err = false;
      link.value.send = false;
    }, 2000);
    return;
  }
  // 提交链接
  const res = await $fetch("/api/page/link/add", { server: false, method: "POST", body: link.value });
  const data = res as any;
  if (data) {
    if (data.code != "S000") {
      link.value.errTips = data.msg;
    } else {
      link.value.succTips = "提交成功，请等待审核通过后展示！";
      link.value.ok = true;
    }
  } else {
    link.value.errTips = "系统异常";
  }
  if (link.value.errTips != "") {
    link.value.err = true;
    setTimeout(() => {
      link.value.err = false;
      link.value.send = false;
    }, 2000);
    return;
  }
  if (link.value.ok) {
    setTimeout(() => {
      link.value.ok = false;
      link.value.send = false;
    }, 2000);
    closeLink();
  }
}
const chooesImg = (e: Event) => {
  return new Promise(async (resolve, reject) => {
    if (e.target) {
      const files = (e.target as HTMLInputElement).files;
      if (files && files[0]) {
        let file = files[0];
        if (window.FileReader && file) {
          if (window.FileReader && file) {
            const url = await readFile(file).catch((e) => {
              link.value.err = true;
              link.value.errTips = "读取图片失败";
              setTimeout(() => {
                link.value.err = false;
                link.value.send = false;
              }, 2000);
              return;
            });
            link.value.source = [];
            await resizeImg(url as string, file.type, 200).catch((e) => {
              link.value.err = true;
              link.value.errTips = "图片缩放失败";
              setTimeout(() => {
                link.value.err = false;
                link.value.send = false;
              }, 2000);
              return;
            });
            resolve(0);
            return;
          }
        }
      }
    }
    return;
  });
};
const readFile = (file: File) => {
  return new Promise((resolve, reject) => {
    var reader = new FileReader();
    reader.readAsDataURL(file);
    //监听文件读取结束后事件
    reader.onloadend = function (ex: ProgressEvent<FileReader>) {
      // 原图
      if (ex.target) {
        resolve(ex.target.result);
        return;
      }
    };
    reader.onerror = function () {
      reject(2); // 读取失败
      return;
    };
  });
};
const resizeImg = (url: string | ArrayBuffer | null, type: string, thisSize: number) => {
  return new Promise((resolve, reject) => {
    // 创建一个 Image 对象
    var image = new Image();
    // 绑定 load 事件处理器，加载完成后执行
    image.onload = function () {
      // 获取 canvas DOM 对象
      var canvas = document.createElement("canvas");
      // 生成缩放图（注意压缩<=0表示无需压缩）
      if (thisSize > 0 && image.width > thisSize) {
        // 宽度等比例缩放 *=
        image.height *= thisSize / image.width;
        image.width = thisSize;
      }
      // 获取 canvas的 2d 画布对象,
      var ctx = canvas.getContext("2d");
      if (ctx) {
        // canvas清屏，并设置为上面宽高
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        // 重置canvas宽高
        canvas.width = image.width;
        canvas.height = image.height;
        // 将图像绘制到canvas上
        ctx.drawImage(image, 0, 0, image.width, image.height);
        // !!! 注意，image 没有加入到 dom之中
        var blob = canvas.toDataURL(type);
        link.value.source.push(blob);
        link.value.sourceShow = blob;
        resolve(0);
        return;
      }
    };
    image.src = url as string;
  });
};
function closeLink() {
  props.linkAdd.show = false;
  let commz = document.getElementById("commz");
  let commlk = document.getElementById("commlk");
  commz?.classList.remove("commzs");
  commlk?.classList.remove("commolk");
  setTimeout(() => {
    commz?.classList.remove("commtrans");
    commz?.classList.remove("commidx");
  }, 310);
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
