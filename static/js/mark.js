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
	let mRunSince = Math.floor(runAbout * 0.02);
	runTimer = setInterval(function () {
		runNow = Math.floor(runNow + mRunSince);
		window.scrollTo(0,runNow);
		if((runBig && Math.floor(runNow) >= Math.floor(runHeight)) || (!runBig && Math.floor(runNow) <= Math.floor(runHeight))) {
			window.scrollTo(0,runHeight);
			clearInterval(runTimer);
		}
	},10);
}