var bodyNode = document.body;
var sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
var loc = window.localStorage;
var the = loc.getItem("the") ? loc.getItem("the") : "light";
bodyNode.classList.add(the);
setTimeout(function() {bodyNode.classList.add("trans");}, 100);
window.addEventListener("scroll", function () {
	sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
	if(sTop > 280) {
		bodyNode.classList.add("sh");
	} else {
		bodyNode.classList.remove("sh");
	}
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