var bodyNode = document.body;
var runNow = bodyNode.scrollTop || document.documentElement.scrollTop;
var loc = window.localStorage;
var the = loc.getItem("the") ? loc.getItem("the") : "light";
var hi = document.getElementById("hi");
var hih = 900;
bodyNode.classList.add(the);
setTimeout(function() {bodyNode.classList.add("trans");}, 100);
ready(initIn)
window.addEventListener("scroll", function () {
  runNow = bodyNode.scrollTop || document.documentElement.scrollTop;
  if(!hi) {
    hi = document.getElementById("hi");
  }
  // offsetTop(上级元素)   // runNode.getBoundingClientRect().top + runNow
  hih = hi.offsetTop;
  if (runNow > hih) {
    bodyNode.classList.add("hs");
  }else{
    if (runNow > 10) {
      bodyNode.classList.add("hh");
    }else{
      bodyNode.classList.remove("hh");
    }
    bodyNode.classList.remove("hs");
  }
	// sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
	// if(sTop > 280) {
	// 	bodyNode.classList.add("sh");
	// } else {
	// 	bodyNode.classList.remove("sh");
	// }
});
// 导航栏
function doMenu() {
  if(bodyNode.classList.contains("om")){
    bodyNode.classList.remove("om");
  }else{
    bodyNode.classList.add("om");
  }
}
// 变色
function theCheck(t) {
  bodyNode.classList.remove(the);
  loc.setItem("the",t);
  the = t;
  bodyNode.classList.add(the);
}
function initIn() {
  // setTimeout(function() {bodyNode.classList.add("init");}, 200);
}

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