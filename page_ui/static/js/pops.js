var popHave = false
var popRun = false
// 原生弹窗控件
function PopUp(text, type, time) {
  var icon = '✔'
  if(type == 1) {
    icon = '✘'
  }
  if(type == 2) {
    icon = '•'
  }
  if(document.getElementById("popD")) {
    popHave = true
  }
  var popHtml = '<div class="popE"><span>' + icon + '</span><p>' + text + '</p></div>'
  if(popHave) {
    document.getElementById("popD").innerHTML = popHtml
  } else {
    // 弹层元素
    var popDiv = document.createElement("div")
    popDiv.classList.add("popD")
    popDiv.setAttribute("id", "popD")
    popDiv.innerHTML = popHtml
    bodyNode.appendChild(popDiv)
    popHave = true
  }
  // 先移除弹出效果
  if(type == 2) {
    bodyNode.classList.add("popLd")
  } else {
    bodyNode.classList.remove("popLd")
  }
  setTimeout(function() {
    // 显示弹层
    document.getElementById("popD").style.display = "block"
    setTimeout(function() {
      if(!popRun) {
        bodyNode.classList.add("popOp")
      }
      popRun = true
      if(time != 0) { // 0表示手工触发关闭
        setTimeout(function() {
          // 定时退出动画
          popRun = false
          bodyNode.classList.remove("popOp")
          setTimeout(function() {
            document.getElementById("popD").style.display = "none"
          }, 500)
        }, time * 1000)
      }
    }, 100)
  }, 100)
}