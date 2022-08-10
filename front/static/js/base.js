// body元素
var bodyNode = document.body;
// 当前执行高度
var runNow = bodyNode.scrollTop || document.documentElement.scrollTop;
// 屏幕宽高
var cH = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;
var cW = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
// 浏览器存储
var loc = window.localStorage;
// 暗黑模式按钮
var the = loc.getItem("the") ? loc.getItem("the") : "light";
// 主元素占位(顶部和底部)
var ht = document.getElementById("ht");
var hth = 900;
var hb = document.getElementById("hb");
var hbh = 1900;
// 浮动盒子
var gebox = document.getElementById("gebox");
var geboxh = 0;
// 浮动盒子上方
var get = document.getElementById("get");
// 记忆风格模式
bodyNode.classList.add(the);
// 添加动画模式
setTimeout(function() {bodyNode.classList.add("trans");}, 100);
// 元素初始化
ready(initIn)
// 滚动行为
window.addEventListener("scroll", function () {runScroll()});

// 触发滚动变换
function runScroll() {
  runNow = bodyNode.scrollTop || document.documentElement.scrollTop;
  // 导航动效
  if (runNow > hth) {
    bodyNode.classList.add("hs");
  }else{
    if (runNow > 10) {
      bodyNode.classList.add("hh");
    }else{
      bodyNode.classList.remove("hh");
    }
    bodyNode.classList.remove("hs");
  }
  if(cW > 1200) {
    // 侧边动效
    if(gebox) {
      var nowRunBt = runNow+cH
      if(nowRunBt > geboxh){
        if(nowRunBt > hbh){
          bodyNode.classList.remove("gef");
          bodyNode.classList.add("geb");
        }else{
          bodyNode.classList.remove("geb");
          bodyNode.classList.add("gef");
        }
      }else{
        bodyNode.classList.remove("gef");
      }
    }
  }
}

// 变色
function theCheck(t) {
  bodyNode.classList.remove(the);
  loc.setItem("the",t);
  the = t;
  bodyNode.classList.add(the);
}

// 初始化
var inter = 0;
var interRun;
function initIn() {
  interRun = setInterval(function(){
    // 初始化高度值
    initH();
    inter++;
    if(inter > 4) {
      clearInterval(interRun);
    }
  },300)
}

// 初始化高度值
function initH() {
  runNow = bodyNode.scrollTop || document.documentElement.scrollTop;
  if(!ht || !hb) {
    ht = document.getElementById("ht");
    hb = document.getElementById("hb");
  }
  if(ht&&hb) {
    // offsetTop(上级元素)
    // runNode.getBoundingClientRect().top + runNow
    hth = ht.offsetTop;
    hbh = hb.offsetTop - 130;
  }
  // 如果有浮动侧边
  if(!gebox || !get){
    gebox = document.getElementById("gebox");
    get = document.getElementById("get");
  }
  if(gebox && get){
    var geth = get.getBoundingClientRect().bottom + runNow;
    geboxh = geth + gebox.offsetHeight;
    // 触发滚动变换
    runScroll();
  }
}

// 初始化钩子函数
function ready(fn){
  if(document.addEventListener){
    document.addEventListener('DOMContentLoaded',function(){
      document.removeEventListener('DOMContentLoaded',arguments.callee);
      fn();
    });
  }else if(document.attachEvent){
    document.attachEvent('onreadystatechange',function(){
      if(document.readystate=='complete'){
        document.dispatchEvent('onreadystatechange',arguments.callee);
        fn();
      }
    })
  }
}

