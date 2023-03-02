// body元素
var bodyNode = document.body
// 当前执行高度
var runNow = bodyNode.scrollTop || document.documentElement.scrollTop
// 屏幕宽高
var cH = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight
var cW = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth
// 浏览器存储
var loc = window.localStorage
// 主题模式(初始化)
var mod = loc.getItem("mod") ? loc.getItem("mod") : 0
var mods = ["light","dark"]
var modBts = ["&#xF116;","&#xF115;"]
var modDoc = null
// DoMenu操作
var menuFlag = false
// DoCata操作
var cataFlag = false
var cataBts = ["&#xF113;","&#xF114;"]
var cataDoc = null
// DoConn操作
var connFlag = false
// 返回顶部定时器
var backTimer = null

// 元素初始化
ready(initIn)
// 滚动行为
window.addEventListener("scroll", function() {
  runScroll()
})
// 初始化函数
function initIn() {
  modDoc = document.getElementById("mod")
  cataDoc= document.getElementById("hcata")
  doMod(false)
  // 添加动画模式
  setTimeout(function() {
    bodyNode.classList.add("trans")
  }, 100)
}

// 触发滚动变换
function runScroll() {
  runNow = bodyNode.scrollTop || document.documentElement.scrollTop
  // 导航动效
  if(runNow > 20) {
    bodyNode.classList.add("hs")
  } else {
    bodyNode.classList.remove("hs")
  }
  // 菜单复位
  if(menuFlag) {
    menuFlag = !menuFlag
    bodyNode.classList.remove("omenu")
  }
  // 目录复位
  if(cataFlag) {
    cataFlag = !cataFlag
    cataDoc.innerHTML = cataBts[cataFlag?1:0]
    bodyNode.classList.remove("ocata")
  }
  // 联系复位
  if(connFlag) {
    connFlag = !connFlag
    bodyNode.classList.remove("oconn")
  }
}

// 主题切换
function doMod(flag) {
  if(flag) {
    bodyNode.classList.remove(mods[mod])
    mod = mod==0 ? 1 : 0
  }
  bodyNode.classList.add(mods[mod])
  modDoc.innerHTML = modBts[mod]
  loc.setItem("mod",mod)
  // 菜单复位
  if(menuFlag) {
    menuFlag = !menuFlag
    bodyNode.classList.remove("omenu")
  }
}

// 菜单函数
function doMenu() {
  if(menuFlag) {
    bodyNode.classList.remove("omenu")
  } else {
    bodyNode.classList.add("omenu")
  }
  menuFlag = !menuFlag
  if(cataFlag) {
    cataFlag = !cataFlag
    bodyNode.classList.remove("ocata")
  }
}

// 目录函数
function doCata() {
  if(cataFlag) {
    bodyNode.classList.remove("ocata")
  } else {
    bodyNode.classList.add("ocata")
  }
  cataFlag = !cataFlag
  cataDoc.innerHTML = cataBts[cataFlag?1:0]
  if(menuFlag) {
    menuFlag = !menuFlag
    bodyNode.classList.remove("omenu")
  }
}

// 底部联系
function doConn() {
  if(connFlag) {
    bodyNode.classList.remove("oconn")
  } else {
    bodyNode.classList.add("oconn")
  }
  connFlag = !connFlag
}

// 返回顶部
function doBackTop() {
  clearInterval(backTimer)
  let backNow = bodyNode.scrollTop || document.documentElement.scrollTop
  backTimer = setInterval(function() {
    // 5%的步长
    let mRunSince = Math.floor(backNow * 0.02) + 4
    backNow = Math.floor(backNow - mRunSince)
    window.scrollTo(0, backNow)
    if(Math.floor(backNow) <= 0) {
      window.scrollTo(0, 0)
      clearInterval(backTimer)
    }
  }, 10)
}


// 初始化钩子函数
function ready(fn) {
  if(document.addEventListener) {
    document.addEventListener('DOMContentLoaded', function() {
      document.removeEventListener('DOMContentLoaded', arguments.callee)
      fn()
    })
  } else if(document.attachEvent) {
    document.attachEvent('onreadystatechange', function() {
      if(document.readystate == 'complete') {
        document.dispatchEvent('onreadystatechange', arguments.callee)
        fn()
      }
    })
  }
}

// 重载界面
function resize() {
  // 当前执行高度
  runNow = bodyNode.scrollTop || document.documentElement.scrollTop
  // 屏幕宽高
  cH = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight
  cW = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth
  // 初始化方法
  initIn()
}

// 自适应监听代码
window.addEventListener("resize", function() {
  resize()
})

// ArrarToJsonObj
function arrayToJsonObj(array) {
  var obj = {}
  array.forEach(i => {
    obj[i.name] = i.value
  })
  return obj
}