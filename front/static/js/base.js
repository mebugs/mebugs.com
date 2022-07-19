var bodyNode = document.body;
var sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
var loc = window.localStorage;
var the = loc.getItem("the") ? loc.getItem("the") : "light";
bodyNode.classList.add(the);
setTimeout(function() {bodyNode.classList.add("trans");}, 100);
ready(initIn)
window.addEventListener("scroll", function () {
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