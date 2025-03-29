var body, html, loc, side, out;
// 初始化
function InitDom() {
  initDomLoad(domInit());
}

function domInit() {
  side = true;
  out = true;
  html = document.documentElement;
  body = document.body;
  loc = window.localStorage;
  setTimeout(() => {
    body.classList.add("bzd");
  }, 200);
  loadNormalImg();
}

// 恢复部分场景
function InitBack() {
  if (body.classList.contains("bzd")) {
    body.classList.remove("bzd");
  }
}

// 移动定时器
var runTimer = null;
function toWhere(where) {
  clearInterval(runTimer);
  let runNow = body.scrollTop || document.documentElement.scrollTop;
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

// 常规加载
function loadNormalImg() {
  loadImg(".m");
}

// 指定加载
function LoadPostsImg() {
  loadImg(".pos");
}

// Bann独立加载
function LoadBannImg() {
  loadImg(".bn");
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

// 初始化函数
function initDomLoad(fn) {
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
