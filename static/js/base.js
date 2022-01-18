var loc = window.localStorage;
var show = loc.getItem("show") ? loc.getItem("show") : "light";
var backTimer;
var bodyNode = document.body;
bodyNode.classList.add(show);
setTimeout(function() {bodyNode.classList.add("trans");}, 100);
ready(hideLoader)
var cH = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;
var cW = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
var sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
var mHeight = false;
var mEnd = false;
var mbm = false;
var mbpm = false;
var pomenu =  document.getElementById("pomenu");
if(cW > 1200) {
  mHeight = document.getElementById("menue") ? document.getElementById("menue").getBoundingClientRect().top - 90 + sTop : false;
  mEnd = mHeight ? getMendHeight() : false;
}
window.addEventListener("scroll", function () {
	sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
	if(sTop > 50) {
		bodyNode.classList.add("headerbg");
	} else {
		// 回到顶部关闭目录
		if(cW < 1200 && pomenu) {
		  doPmenu(false);
		}
		bodyNode.classList.remove("headerbg");
	}
	if(mHeight && cW > 1200) {
		if(sTop > mHeight)  {
			bodyNode.classList.add("menus");
			if(sTop > mEnd) {
				bodyNode.classList.add("menuse");
			} else {
				bodyNode.classList.remove("menuse");
			}
		} else {
			bodyNode.classList.remove("menus");
		}
	}
  // 滑动关闭目录
  if(cW < 1200 && pomenu) {
    doPmenu(false);
  }
})
// update 20 since five sc
var upF = 0;
var upTimer;
// get menu end stop Height 
function getMendHeight() {
	upTimer = setInterval(function () {
		mEnd = doGetMendHeight();
		upF ++;
		if(upF > 20) {
			clearInterval(upTimer);
		}
	},50);
	return doGetMendHeight();
}
function doGetMendHeight() {
	var endH = document.getElementById("mother").getBoundingClientRect().top + sTop;
  return endH - cH;
}
// set show light or dark
function setShow() {
	if(show == "light") {
		show = "dark";
		bodyNode.classList.remove("light");
	} else {
		show = "light";
		bodyNode.classList.remove("dark");
	}
	loc.setItem("show",show);
	bodyNode.classList.add(show);
}
function setMbm() {
  mbm = !mbm;
  if(mbm) {
    bodyNode.classList.add("opmenu");
    if(pomenu) {
       doPmenu(false);
    }
  }else{
    bodyNode.classList.remove("opmenu");
  }
}

function doPmenu(down) {
  if(down) {
    mbpm = !mbpm;
    if(mbpm) {
      bodyNode.classList.add("oppmenu");
      if(mbm) {
        mbm = !mbm;
        bodyNode.classList.remove("opmenu");
      }
    }else{
      bodyNode.classList.remove("oppmenu");
    }
  }else{
    mbpm = false;
    bodyNode.classList.remove("oppmenu");
  }
}
// back top 
function backTop() {
	clearInterval(backTimer);
	let startRun = bodyNode.scrollTop || document.documentElement.scrollTop;
	let backOnce = Math.floor(startRun * 0.02);
	backTimer = setInterval(function () {
		startRun = Math.floor(startRun - backOnce);
		window.scrollTo(0,startRun);
		if(Math.floor(startRun) <= 0) {
			window.scrollTo(0,0);
			clearInterval(backTimer);
		}
	},10);
}

function hideLoader() {
  setTimeout(function() {bodyNode.classList.add("hideLoader");}, 200);
  var loader = document.getElementById("loader");
  setTimeout(function() {loader.style.display = 'none';}, 700);
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