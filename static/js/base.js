var loc = window.localStorage;
var show = loc.getItem("show") ? loc.getItem("show") : "light";
var backTimer;
var bodyNode = document.body;
bodyNode.classList.add(show);
setTimeout(function() {bodyNode.classList.add("trans");}, 100);
var mHeight = document.getElementById("menue") ? document.getElementById("menue").getBoundingClientRect().top - 90 : false;
var mEnd = mHeight ? getMendHeight() : false;
var cH = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;
var sTop = 0;
window.addEventListener("scroll", function () {
	sTop = bodyNode.scrollTop || document.documentElement.scrollTop;
	if(sTop > 50) {
		bodyNode.classList.add("headerbg");
	} else {
		bodyNode.classList.remove("headerbg");
	}
	if(mHeight) {
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
})
// update 20 since five sc
var upF = 0;
var upTimer;
// get menu end stop Height 
function getMendHeight() {
	upTimer = setInterval(function () {
		mEnd = doGetMendHeight();
		upF ++;
		console.log(mEnd);
		if(upF > 20) {
			clearInterval(upTimer);
		}
	},3000);
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
// WaitForSVGInjectInit
function initSvg(obj) {
	try{
	    SVGInject(obj);
	}catch(e){
	    setTimeout(function(){
			initSvg(obj);
		},200)
	}
}
