var swiper;
var fumer;
window.addEventListener("DOMContentLoaded", function () {
	// Page Init (ext IE)
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
	fumer = new Swiper('.fm', {
		loop:true,
	    autoplay: {
	        disableOnInteraction: false,
	        stopOnLastSlide: false,
	        delay: 5000
	    },
		speed: 900,
		navigation: {nextEl: '.swiper-button-next',prevEl:'.swiper-button-prev'}
	});
});