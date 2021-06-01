var loc = window.localStorage;
var show = loc.getItem("show") ? loc.getItem("show") : "light";
var backTimer;
document.body.classList.add(show);
window.addEventListener("scroll", function () {
	let top = document.body.scrollTop || document.documentElement.scrollTop;
	if(top > 150) {
		document.body.classList.add("headerbg");
	} else {
		document.body.classList.remove("headerbg");
	}
})
// set show light or dark
function setShow() {
	if(show == "light") {
		show = "dark";
		document.body.classList.remove("light");
	} else {
		show = "light";
		document.body.classList.remove("dark");
	}
	loc.setItem("show",show);
	document.body.classList.add(show);
}
// back top 
function backTop() {
	clearInterval(backTimer);
	let startRun = document.body.scrollTop || document.documentElement.scrollTop;
	backTimer = setInterval(function () {
		startRun = Math.floor(startRun - (startRun * 0.2));
		window.scrollTo(0,startRun);
		if(Math.floor(startRun) === 0) {
			clearInterval(backTimer);
		}
	},50);
}