// 当前加载的位置
var bannIndex = 0
// 全部浮动子节点
var bannChild = null
// 定时器
var bannInterval = null
// Bann元素
var bannObj = document.getElementById("bann")

// -- 操作
function toBannLeft() {
  event.stopPropagation()
  resetBannInterval()
  runBannLeft()
}

// ++ 操作
function toBannRight() {
  event.stopPropagation()
  resetBannInterval()
  runBannRight()
}

// 重置定时器
function resetBannInterval(init) {
  clearInterval(bannInterval)
  bannInterval = setInterval(function() {
    runBannRight()
  }, 4000)
}

// 执行左翼
function runBannLeft() {
  bannIndex = fixIndex(bannIndex - 1)
  bannChild[bannIndex].classList.remove("bpre")
  bannChild[fixIndex(bannIndex + 1)].classList.remove("bnow")
  bannChild[fixIndex(bannIndex + 2)].classList.remove("bnext")
  bannChild[fixIndex(bannIndex + 3)].classList.remove("bwait")
  runBann()
}

// 执行右移
function runBannRight() {
  bannIndex = fixIndex(bannIndex + 1)
  bannChild[bannIndex].classList.remove("bnext")
  bannChild[fixIndex(bannIndex + 1)].classList.remove("bwait")
  bannChild[fixIndex(bannIndex + 2)].classList.remove("bpre")
  bannChild[fixIndex(bannIndex + 3)].classList.remove("bnow")
  runBann()
}

// 新Index定位
function runBann() {
  bannChild[bannIndex].classList.add("bnow")
  bannChild[fixIndex(bannIndex + 1)].classList.add("bnext")
  bannChild[fixIndex(bannIndex + 2)].classList.add("bwait")
  bannChild[fixIndex(bannIndex + 3)].classList.add("bpre")
}

function initBann() {
  bannChild = document.querySelectorAll('.bne>.bni')
  runBann()
  resetBannInterval()
}

function fixIndex(index) {
  if(index > 3) {
    index = index - 4
  } else if(index < 0) {
    index = index + 4
  }
  return index
}

// 元素初始化
ready(initBann)