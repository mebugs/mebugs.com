// 菜单高度
var mdMiHs = {}
var mdMis, st, end, outDoc, outH, outTop, isMob
init(initPost)

// 初始化文章页
function initPost() {
  mdMis = document.querySelectorAll('.outline li span')
  outDoc = document.querySelector('.vditor-outline__content')
  outH = outDoc.clientHeight
  isMob = window.innerWidth <= 1024
  // 初始化高度
  initMdMiHs()
  // 初始化滚动事件
  window.addEventListener('scroll', function () {
    refreshMdMiHs()
  })
}
// 初始化高度
function initMdMiHs() {
  mdMis.forEach((mdMi) => {
    let mdMiId = mdMi.getAttribute('data-target-id')
    if (mdMiId) {
      let mdMiIdH = document.getElementById(mdMiId).offsetTop
      mdMiHs[mdMiId] = mdMiIdH
      mdMi.onclick = function () {
        event.preventDefault()
        if (isMob) {
          toWhere(mdMiHs[mdMiId] - 60)
        } else {
          toWhere(mdMiHs[mdMiId])
        }
        setOut()
      }
    }
  })
  st = document.getElementById('st').offsetTop
  ed = document.getElementById('ed').offsetTop
}
// 刷新高度并选择高度
function refreshMdMiHs() {
  let miKeys = Object.keys(mdMiHs)
  miKeys.forEach((miKey) => {
    mdMiHs[miKey] = document.getElementById(miKey).offsetTop
  })
  st = document.getElementById('st').offsetTop
  ed = document.getElementById('ed').offsetTop
  let runNow = body.scrollTop || document.documentElement.scrollTop
  refreshCheck(runNow)
}

// 刷新选择
function refreshCheck(runNow) {
  mdMis.forEach((mdMi) => {
    mdMi.classList.remove('outcheck')
  })
  // 超出边界无需配色
  if (runNow >= ed || runNow <= st) {
    return
  }
  // 匹配选择
  let checkKey = 0
  let checkHi = 0
  var miKeys = Object.keys(mdMiHs)
  let pre = isMob ? 70 : 10
  miKeys.forEach((miKey) => {
    if (runNow >= mdMiHs[miKey] - pre && runNow > checkHi) {
      checkHi = mdMiHs[miKey] - pre
      checkKey = miKey
    }
  })
  if (checkKey != 0) {
    mdMis.forEach((mdMi) => {
      let mdMiId = mdMi.getAttribute('data-target-id')
      if (mdMiId && mdMiId == checkKey) {
        mdMi.classList.add('outcheck')
        outTop = mdMi.offsetTop
        if (outTop > outH - 36) {
          outDoc.scrollTop = outTop - outH + 36
        } else {
          outDoc.scrollTop = 0
        }
      }
    })
  }
}

// 移动端侧边大纲
function setOut() {
  if (out) {
    body.classList.add('out')
  } else {
    body.classList.remove('out')
  }
  out = !out
}

// 代码复制
function copyCode(node) {
  var text = node.parentNode.nextSibling.innerText
  if (navigator.clipboard) {
    // clipboard api 复制
    navigator.clipboard.writeText(text)
  } else {
    // 上层texteare复制
    node.previousElementSibling.value = text
    node.previousElementSibling.select()
    document.execCommand('copy', true)
  }
  node.setAttribute('aria-label', '已复制')
}
