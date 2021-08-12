var loc = window.localStorage;
var show = loc.getItem("show") ? loc.getItem("show") : "light";
var backTimer;
var bodyNode = document.body;
bodyNode.classList.add(show);
setTimeout(function() {bodyNode.classList.add("trans");}, 100);
var cH = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;
var cW = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
var sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
var mHeight = false;
var mRHeight = false;
var mEnd = false;
var mbm = false;
var mbpm = false;
var pomenu =  document.getElementById("pomenu");
if(cW > 1200) {
  mHeight = document.getElementById("menue") ? document.getElementById("menue").getBoundingClientRect().top - 90 + sTop : false;
  mRHeight = mHeight ? document.getElementById("rhot").getBoundingClientRect().top + sTop : false;
  mEnd = mHeight ? getMendHeight() : false;
}
window.addEventListener("scroll", function () {
	sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
	if(sTop > 50) {
		bodyNode.classList.add("headerbg");
	} else {
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
    if(mRHeight) {
      if(sTop > mRHeight)  {
        bodyNode.classList.add("rfix");
      } else {
        bodyNode.classList.remove("rfix");
      }
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
		if(upF > 200 || sTop > mRHeight) {
			clearInterval(upTimer);
		}
	},50);
	return doGetMendHeight();
}
function doGetMendHeight() {
	var endH = document.getElementById("mother").getBoundingClientRect().top + sTop;
  mRHeight = document.getElementById("rhot").getBoundingClientRect().top + sTop;
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
      pomenu.innerText="关闭目录";
      if(mbm) {
        mbm = !mbm;
        bodyNode.classList.remove("opmenu");
      }
    }else{
      bodyNode.classList.remove("oppmenu");
      pomenu.innerText="打开目录";
    }
  }else{
    mbpm = false;
    bodyNode.classList.remove("oppmenu");
    pomenu.innerText="打开目录";
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
