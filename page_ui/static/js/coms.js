var coms = null
var comsTi = null
var comsTis = ["", "评论文章", "评论楼层", "评论消息"]
var comh = null
var user = {
  name: "",
  email: "",
  qq: "",
  url: "",
  avt: -1
}
var atNo = -1
var atSrc = ""
var allAvt = null
var nameE = null
var detailE = null
ready(initCm)

function initCm() {
  coms = document.getElementById("coms")
  comsTi = document.getElementById("comsTi")
  comh = document.getElementById("comh")
  nameE = document.getElementById("nameE")
  detailE = document.getElementById("detail")
  allAvt = document.querySelectorAll(".avts img")
  initForm()
}

function initForm() {
  let userLocStr = window.localStorage.getItem("mebugs_user")
  if(userLocStr) {
    user = JSON.parse(userLocStr)
  }
  // 头像不存在随机初始化
  if(user.avt == -1) {
    user.avt = randomAt()
  }
  atNo = user.avt
  atSrc = "/static/upload/avtor/" + atNo + ".jpg"
  // 初始化赋值
  document.getElementById("name").value = user.name
  document.getElementById("email").value = user.email
  document.getElementById("qq").value = user.qq
  document.getElementById("url").value = user.url
  document.getElementById("at").src = atSrc
}
// 目前12个头像
function randomAt() {
  return Math.floor(Math.random() * 12)
}
// 打开评论
// cid 0 评论文章 n 回复消息
// level 1 评论文章 2 回复评论 3 回复回复
function openCm(cid, level) {
  coms.style.display = "block"
  comsTi.innerText = comsTis[level]
  setTimeout(function() {
    bodyNode.classList.add("ocoms")
  }, 100)
}

function closeCm() {
  nameE.classList.remove("errN")
  detailE.classList.remove("errN")
  bodyNode.classList.remove("ocoms")
  setTimeout(function() {
    coms.style.display = "none"
  }, 500)
}

// 打开头像选择框
function openAt() {
  comh.style.display = "block"
  setAt(atNo)
  setTimeout(function() {
    bodyNode.classList.add("ocomh")
  }, 100)
}

function closeAt() {
  bodyNode.classList.remove("ocomh")
  setTimeout(function() {
    comh.style.display = "none"
  }, 500)
}

function setAt(index) {
  for(var i = 0;i < allAvt.length;i++) {
    allAvt[i].classList.remove("checkd")
    if(index == i) {
      allAvt[i].classList.add("checkd")
    }
  }
  atNo = index
}

function saveAt() {
  user.avt = atNo
  atSrc = "/static/upload/avtor/" + atNo + ".jpg"
  document.getElementById("at").src = atSrc
  closeAt()
}

// 提交评论
var sendFlag = false

function sendCm() {
  if(sendFlag) {
    PopUp("请勿频繁提交评论", 1, 2)
    return
  }
  let name = document.getElementById("name").value.trim()
  if(!name || name == "") {
    nameE.classList.add("errN")
  } else {
    nameE.classList.remove("errN")
  }
  let detail = document.getElementById("detail").value.trim()
  if(!detail || detail == "") {
    detailE.classList.add("errN")
  } else {
    detailE.classList.remove("errN")
  }
  let email = document.getElementById("email").value.trim()
  let qq = document.getElementById("qq").value.trim()
  let url = document.getElementById("url").value.trim()
  PopUp("提交中", 2, 0)
  sendFlag = true
  user.name = name
  user.email = email
  user.qq = qq
  user.url = url
  setTimeout(function() {
    PopUp("成功", 0, 2)
    saveUser()
    closeCm()
  }, 3000)
  setTimeout(function() {
    sendFlag = false
  }, 30000)
}

function saveUser() {
  window.localStorage.setItem("mebugs_user", JSON.stringify(user))
}