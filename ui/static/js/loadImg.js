// 初始化
init(loadNormalImg)

// 常规加载
function loadNormalImg() {
  loadImg(false)
}

// Bann独立加载
function loadBannImg() {
  loadImg(true)
}

// 加载图片
function loadImg(isBann) {
  var allImg = document.querySelectorAll('#content .mimg img')
  for (var i = 0; i < allImg.length; i++) {
    let img = allImg[i]
    let oimg = img.parentNode
    let isBannImg = oimg.classList.contains('bimg')
    if (isBann == !isBannImg) {
      continue
    }
    let imgSrc = img.getAttribute('src')
    img.setAttribute('src', '/static/img/logo.png')
    let loder = new Image()
    loder.src = imgSrc
    loder.onload = () => {
      img.setAttribute('src', imgSrc)
      setTimeout(() => {
        oimg.classList.add('dimg')
      }, 1000)
    }
    loder.onerror = () => {
      setTimeout(() => {
        oimg.classList.add('dimg')
      }, 1000)
    }
  }
}
