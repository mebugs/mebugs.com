var swiper
init(initSwiper)
// Page Init (ext IE)
function initSwiper() {
  // banner
  swiper = new Swiper('.bann', {
    loop: true,
    autoplay: {
      disableOnInteraction: false,
      stopOnLastSlide: false,
      delay: 3500
    },
    speed: 600,
    pagination: { el: '.swiper-pagination', clickable: true },
    on: {
      init: function (swiper) {
        loadBannImg()
      }
    }
  })
}
