var ca = null
var md = null
var caAi = null
var mdAi = null
var mdAiHs = {}
var initOk = false
ready(initPost)
function initPost() {
  ca = document.getElementById("cata")
  caAi = document.querySelectorAll(".cata .list li a")
  md = document.getElementById("md")
  mdAi = document.querySelectorAll(".md a:not([href])")
  refreshAi()
  initCaAi()
}

// 抵达顶部和更多
function postGo(i) {
  var goTo = 0;
  if(i==0) {
    goTo = document.getElementById("st").offsetTop-110
  }else{
    goTo = document.getElementById("ed").offsetTop-60
  }
  toWhere(goTo)
}

// 滚动行为
window.addEventListener("scroll", function () {readAi()})

function readAi() {
  let runNow = bodyNode.scrollTop || document.documentElement.scrollTop
  // 元素初始化完成
  if(initOk) {
    refreshAi()
    // 判断谁被选中
    refreshCi(runNow)
  }
}

// 初始化点击
function initCaAi() {
  caAi.forEach(cAi => {
    let cHref = cAi.getAttribute("href")
    cHref = cHref.substr(1)
    cAi.setAttribute("cid",cHref)
    cAi.onclick = function(){
      event.preventDefault()
      let cid = this.getAttribute("cid")
      toWhere(mdAiHs[cid])
    }
  })
}

// 刷新高度
function refreshAi() {
  mdAi.forEach(mAi => {
    mdAiHs[mAi.getAttribute("id")] = mAi.offsetTop//+80
  })
  // ed = document.getElementById("ed").offsetTop//+80
  initOk = true 
}

// 判断选中
function refreshCi(runNow){
  caAi.forEach(cAi => {
    cAi.classList.remove("cak")
  })
  if(runNow > ed) {
    return
  }
  let ci = 0
  // 触发选中的A
  let cId = null
  for(let key in mdAiHs) {
    if(ci==0) {
      if(runNow < mdAiHs[key]-5) {
        return
      }
    }
    ci++
    if(runNow >= mdAiHs[key]-5) {
      cId = key
    }
  }
  // 超出结束范围
  if(runNow >= document.getElementById("ed").offsetTop-100){
    cId = null
  }
  caAi.forEach(cAi => {
    let caId = cAi.getAttribute("cid")
    if(caId == cId) {
      cAi.classList.add("cak")
    }
  })
}

// 移动定时器
var runTimer = null

function toWhere(where) {
  clearInterval(runTimer)
  let runNow = bodyNode.scrollTop || document.documentElement.scrollTop
  let runDo = false
  if(where > runNow) {
    // 下移
    runDo = true
  }
  runTimer = setInterval(function () {  
    // 5%的步长
    let needRun = Math.abs(runNow - where)
    let needRunSince = Math.floor(needRun * 0.03)+4
    let runTo = runDo ? Math.floor(runNow + needRunSince):Math.floor(runNow - needRunSince)
    if(runDo) {
      // 下移超出
      if(runTo > where) {
        window.scrollTo(0,where)
        clearInterval(runTimer)
        return
      }
    }else{
      // 上移小于
      if(runTo < where) {
        window.scrollTo(0,where)
        clearInterval(runTimer)
        return
      }
    }
    // 去往指定位置
    runNow = runTo
    window.scrollTo(0,runTo)
  },10)	
}
