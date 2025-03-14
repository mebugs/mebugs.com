var swiper;
init(initSwiper);
// Page Init (ext IE)
function initSwiper() {
  // banner
  swiper = new Swiper(".swiper-container", {
    loop: true,
    autoplay: {
      disableOnInteraction: false,
      stopOnLastSlide: false,
      delay: 3500000,
    },
    grabCursor: true,
    speed: 400,
    pagination: { el: ".swiper-pagination", clickable: true },
    on: {
      init: function (swiper) {
        loadBannImg();
      },
    },
  });
}
