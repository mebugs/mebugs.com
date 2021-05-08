var swiper;
(function($) {
	swiper = new Swiper('.bn', {
		effect : 'coverflow',
		coverflowEffect: {
		    rotate: 30,
		    stretch: 20,
		    depth: 60,
		    modifier: 2,
		    slideShadows : true
		},
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
})(jQuery)