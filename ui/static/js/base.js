var body, html, loc, side, out
// 初始化
init(InitDom)
// 初始化
function InitDom() {
  side = true
  out = true
  html = document.documentElement
  body = document.body
  loc = window.localStorage
  setTimeout(function () {
    body.classList.add('trans')
    setTimeout(function () {
      body.classList.add('pin')
    }, 1000)
  }, 50)
}

// 移动端侧边菜单
function setSide() {
  if (side) {
    body.classList.add('side')
  } else {
    body.classList.remove('side')
  }
  side = !side
}

// 移动定时器
var runTimer = null
function toWhere(where) {
  clearInterval(runTimer)
  let runNow = body.scrollTop || document.documentElement.scrollTop
  let runDo = false
  if (where > runNow) {
    // 下移
    runDo = true
  }
  runTimer = setInterval(function () {
    // 4%的步长
    let needRun = Math.abs(runNow - where)
    let needRunSince = Math.floor(needRun * 0.04) + 4
    let runTo = runDo
      ? Math.floor(runNow + needRunSince)
      : Math.floor(runNow - needRunSince)
    if (runDo) {
      // 下移超出
      if (runTo > where) {
        window.scrollTo(0, where)
        clearInterval(runTimer)
        return
      }
    } else {
      // 上移小于
      if (runTo < where) {
        window.scrollTo(0, where)
        clearInterval(runTimer)
        return
      }
    }
    // 去往指定位置
    runNow = runTo
    window.scrollTo(0, runTo)
  }, 10)
}

// 初始化函数
function init(fn) {
  if (document.addEventListener) {
    document.addEventListener('DOMContentLoaded', function () {
      document.removeEventListener('DOMContentLoaded', arguments.callee)
      fn()
    })
  } else if (document.attachEvent) {
    document.attachEvent('onreadystatechange', function () {
      if (document.readystate == 'complete') {
        document.dispatchEvent('onreadystatechange', arguments.callee)
        fn()
      }
    })
  }
}
