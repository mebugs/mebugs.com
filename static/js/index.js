var swiper;
window.addEventListener("DOMContentLoaded", function () {
	// Page Init (ext IE)
	// banner
	swiper = new Swiper('.bann', {
	    autoplay: {
	        disableOnInteraction: false,
	        stopOnLastSlide: false,
	        delay: 3000
	    },
		watchSlidesProgress: true,
	    loop:true,
		speed: 500,
		grabCursor: true,
	    pagination: {el: '.swiper-pagination'}
	});
});