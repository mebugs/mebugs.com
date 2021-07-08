var runTimer;
var runHeight;
var runAbout;
var runNow;
var runBig;
function runTo(id) {
	clearInterval(runTimer);
	runNow = bodyNode.scrollTop || document.documentElement.scrollTop;
	let runNode = document.getElementById(id);
	runHeight = runNode.getBoundingClientRect().top - 70 + runNow;
	runAbout = runHeight - runNow;
	runBig = runAbout > 0 ? true : false;
	runTimer = setInterval(function () {
		runNow = Math.floor(runNow + (runAbout * 0.05));
		window.scrollTo(0,runNow);
		if((runBig && Math.floor(runNow) >= Math.floor(runHeight)) || (!runBig && Math.floor(runNow) <= Math.floor(runHeight))) {
			window.scrollTo(0,runHeight);
			clearInterval(runTimer);
		}
		console.log("FLAG"+runBig)
		console.log("TO"+runHeight)
		console.log("NOW"+runNow)
	},10);
}