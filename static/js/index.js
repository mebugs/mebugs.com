var swiper;
var fumer;
var width = document.body.clientWidth;
// Page Init (ext IE)
window.addEventListener("DOMContentLoaded", function () {
	// banner
	swiper = new Swiper('.bann', { loop:true, autoplay: { disableOnInteraction: false, stopOnLastSlide: false, delay: 3000 }, speed: 600, pagination: {el: '.swiper-pagination',clickable:true} });
	// indexFm make for mobile , every slide use two node
	if(width < 800) { var fmList = document.querySelectorAll('.fm .swiper-slide a'); if(fmList && fmList.length > 0) { var html = '<div class="swiper-wrapper" >'; for(var i=0;i<fmList.length;i++) { if(i%2 ==0) { if(i!=0) { html += '</div>'; } html += '<div class="swiper-slide">'; } html += fmList[i].outerHTML; } html += '</div></div><div class="swiper-button-next"></div><div class="swiper-button-prev"></div>'; var indexFm = document.getElementById('indexfm'); indexFm.innerHTML = html; } }
	// indexFm
	fumer = new Swiper('.fm', { loop:true, autoplay: { disableOnInteraction: false, stopOnLastSlide: false, delay: 6000 }, speed: 900, navigation: {nextEl: '.swiper-button-next',prevEl:'.swiper-button-prev'} });
});