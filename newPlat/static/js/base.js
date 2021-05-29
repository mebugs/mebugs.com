var loc = window.localStorage;
var show = loc.getItem("show") ? loc.getItem("show") : "light";
document.body.classList.add(show);
window.addEventListener("scroll", function () {
	let top = document.body.scrollTop || document.documentElement.scrollTop;
	if(top > 80) {
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