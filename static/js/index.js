var swiper;
// Page Init (ext IE)
window.addEventListener("DOMContentLoaded", function () {
	// banner
	swiper = new Swiper('.bann', {
	    loop:true,
		autoplay: {
		    disableOnInteraction: false,
		    stopOnLastSlide: false,
		    delay: 3000
		},
		speed: 600,
	    pagination: {el: '.swiper-pagination',clickable:true}
	});
});