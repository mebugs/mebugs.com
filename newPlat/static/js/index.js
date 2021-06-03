var swiper;
window.addEventListener("DOMContentLoaded", function () {
	// Page Init (ext IE)
	// banner
	swiper = new Swiper('.bann', {
		effect : 'cube',
	    autoplay: {
	        disableOnInteraction: false,
	        stopOnLastSlide: false,
	        delay: 2000
	    },
		watchSlidesProgress: true,
	    loop:true,
		speed: 1000,
		grabCursor: true,
	    pagination: {el: '.swiper-pagination'}
	});
});