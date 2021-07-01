var loc = window.localStorage;
var show = loc.getItem("show") ? loc.getItem("show") : "dark";
var backTimer;
var bodyNode = document.body;
bodyNode.classList.add(show);
window.addEventListener("scroll", function () {
	let top = bodyNode.scrollTop || document.documentElement.scrollTop;
	if(top > 150) {
		bodyNode.classList.add("headerbg");
	} else {
		bodyNode.classList.remove("headerbg");
	}
})
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
	backTimer = setInterval(function () {
		startRun = Math.floor(startRun - (startRun * 0.2));
		window.scrollTo(0,startRun);
		if(Math.floor(startRun) === 0) {
			clearInterval(backTimer);
		}
	},50);
}