var loc, side, out;
// 菜单高度
var mdMiHs = {};
var mdMis, st, ed, outDoc, outH, outTop, isMob;

// 初始化
export function InitDom() {
  InitDomLoad(domInit());
}

// 初始化
export function InitPostDom() {
  InitDomLoad(domInit());
  InitDomLoad(initPost());
}

// 恢复部分场景
export function InitBack() {
  if (document.body.classList.contains("bzd")) {
    document.body.classList.remove("bzd");
  }
  window.removeEventListener("scroll", refreshMdMiHs);
}

// 指定加载
export function LoadPostsImg() {
  loadImg(".pos");
}

// Bann独立加载
export function LoadBannImg() {
  loadImg(".bn");
}

// 移动定时器
var runTimer = null;
export function ToWhere(where) {
  clearInterval(runTimer);
  let runNow = document.body.scrollTop || document.documentElement.scrollTop;
  let runDo = false;
  if (where > runNow) {
    // 下移
    runDo = true;
  }
  runTimer = setInterval(function () {
    // 4%的步长
    let needRun = Math.abs(runNow - where);
    let needRunSince = Math.floor(needRun * 0.04) + 4;
    let runTo = runDo ? Math.floor(runNow + needRunSince) : Math.floor(runNow - needRunSince);
    if (runDo) {
      // 下移超出
      if (runTo > where) {
        window.scrollTo(0, where);
        clearInterval(runTimer);
        return;
      }
    } else {
      // 上移小于
      if (runTo < where) {
        window.scrollTo(0, where);
        clearInterval(runTimer);
        return;
      }
    }
    // 去往指定位置
    runNow = runTo;
    window.scrollTo(0, runTo);
  }, 10);
}

function domInit() {
  side = true;
  out = true;
  loc = window.localStorage;
  setTimeout(() => {
    document.body.classList.add("bzd");
  }, 200);
  loadNormalImg();
}

// 常规加载
function loadNormalImg() {
  loadImg(".m");
}

// 加载图片 TODO
function loadImg(from) {
  var allImg = document.querySelectorAll(from + " img");
  var imgs = [];
  var index = 0;
  for (var i = 0; i < allImg.length; i++) {
    let img = allImg[i];
    let oimg = img.parentNode;
    let isBannImg = oimg.classList.contains("bgb");
    if (from == ".m" && isBannImg) {
      continue;
    }
    let imgSrc = img.getAttribute("src");
    img.setAttribute("src", "/static/img/empty.png");
    img.setAttribute("ori-src", imgSrc);
    imgs.push(img);
  }
  lazyLoadImg(imgs, index);
}

function lazyLoadImg(imgs, index) {
  if (index >= imgs.length) {
    return;
  }
  var img = imgs[index];
  let imgSrc = img.getAttribute("ori-src");
  let loder = new Image();
  loder.src = imgSrc;
  index++;
  loder.onload = () => {
    img.setAttribute("src", imgSrc);
    setTimeout(() => {
      img.parentNode.classList.add("bgd");
      lazyLoadImg(imgs, index);
    }, 100);
  };
  loder.onerror = () => {
    img.setAttribute("src", imgSrc);
    setTimeout(() => {
      img.parentNode.classList.add("bgd");
      lazyLoadImg(imgs, index);
    }, 100);
  };
}

// 初始化文章页
function initPost() {
  mdMis = document.querySelectorAll(".brl li span");
  outDoc = document.querySelector(".outline_box");
  outH = outDoc.clientHeight;
  isMob = window.innerWidth <= 1024;
  // 初始化高度
  initMdMiHs();
  // 初始化滚动事件
  window.addEventListener("scroll", refreshMdMiHs);
}
// 初始化高度
function initMdMiHs() {
  st = document.getElementById("st").offsetTop;
  ed = document.getElementById("ed").offsetTop;
  // 初始化
  mdMiHs = {};
  mdMis.forEach((mdMi) => {
    let mdMiId = mdMi.getAttribute("data-target-id");
    if (mdMiId) {
      let mdMiIdH = document.getElementById(mdMiId).offsetTop + st;
      mdMiHs[mdMiId] = mdMiIdH;
      mdMi.onclick = function () {
        event.preventDefault();
        if (isMob) {
          ToWhere(mdMiHs[mdMiId] - 60);
        } else {
          ToWhere(mdMiHs[mdMiId]);
        }
        // TODO
        // setOut();
      };
    }
  });
}
// 刷新高度并选择高度
function refreshMdMiHs() {
  st = document.getElementById("st").offsetTop;
  ed = document.getElementById("ed").offsetTop;
  let miKeys = Object.keys(mdMiHs);
  miKeys.forEach((miKey) => {
    mdMiHs[miKey] = document.getElementById(miKey).offsetTop + st;
  });
  let runNow = document.body.scrollTop || document.documentElement.scrollTop;
  refreshCheck(runNow);
}

// 刷新选择
function refreshCheck(runNow) {
  mdMis.forEach((mdMi) => {
    mdMi.classList.remove("outcheck");
  });
  // 超出边界无需配色
  if (runNow >= ed || runNow <= st) {
    return;
  }
  // 匹配选择
  let checkKey = 0;
  let checkHi = 0;
  var miKeys = Object.keys(mdMiHs);
  let pre = isMob ? 70 : 10;
  miKeys.forEach((miKey) => {
    if (runNow >= mdMiHs[miKey] - pre && runNow > checkHi) {
      checkHi = mdMiHs[miKey] - pre;
      checkKey = miKey;
    }
  });
  if (checkKey != 0) {
    mdMis.forEach((mdMi) => {
      let mdMiId = mdMi.getAttribute("data-target-id");
      if (mdMiId && mdMiId == checkKey) {
        mdMi.classList.add("outcheck");
        outTop = mdMi.offsetTop;
        if (outTop > outH - 36) {
          outDoc.scrollTop = outTop - outH + 36;
        } else {
          outDoc.scrollTop = 0;
        }
      }
    });
  }
}

// 初始化函数
export function InitDomLoad(fn) {
  if (document.addEventListener) {
    document.addEventListener("DOMContentLoaded", function () {
      document.removeEventListener("DOMContentLoaded", arguments.callee);
      fn();
    });
  } else if (document.attachEvent) {
    document.attachEvent("onreadystatechange", function () {
      if (document.readystate == "complete") {
        document.dispatchEvent("onreadystatechange", arguments.callee);
        fn();
      }
    });
  }
}
