var avtNos = 12;
var localUserStr = window.localStorage.getItem("localUser");
var localUser;
var crTip = document.getElementById("rwork");
var crMet = document.getElementById("repMethod");
var rzz = document.getElementById("rzz");
var cavt = document.getElementById("cavt");
var myavt = document.getElementById("avt");
if(localUserStr) {
  localUser = JSON.parse(localUserStr);
}else{
  localUser = { name: "", email: "", qq: "", url: "", avt: ""};
}
if(localUser.avt == "") {
  localUser.avt = "/static/upload/avtor/"+randomAvt()+".jpg";
}
myavt.src = localUser.avt;
initUserInput();
function chooseAvt() {
	cavt.style.display = "block";
	setTimeout(function() {
	  bodyNode.classList.add("cavts");
	}, 100);
}
function setAvt(node,no) {
  var allAvt = document.querySelectorAll(".avts img");
  for(var i=0; i < allAvt.length; i++) {
    allAvt[i].classList.remove("checkd");
  }
  node.classList.add("checkd");
	localUser.avt = "/static/upload/avtor/"+no+".jpg";
	myavt.src = localUser.avt;
}
function clsAvt() {
  bodyNode.classList.remove("cavts");
  setTimeout(function() {
    cavt.style.display = "none";
  }, 500);
}
function randomAvt() {
  return Math.floor(Math.random()*avtNos+1);
}
// CommsReplay
function sendCommsR(pid, level, fid) {
  crTip.innerText = "评论楼层";
  crMet.href = "javascript:sendComms("+pid+","+level+","+fid+")";
  initUserInput();
  opnComms();
}
// CommsReplayReplay
function sendCommsRr(pid, level, fid) {
  crTip.innerText = "回复评论";
  crMet.href = "javascript:sendComms("+pid+","+level+","+fid+")";
  initUserInput();
  opnComms();
}
// show replay
function opnComms() {
  rzz.style.display = "block";
  setTimeout(function() {
    bodyNode.classList.add("replay");
  }, 100);
}
// close Replay
function clsComms() {
  bodyNode.classList.remove("replay");
  setTimeout(function() {
    rzz.style.display = "none";
  }, 500);
}
// PostID Level 123 ParentID
function sendComms(pid, level, fid) {
  // validate
  var coms = "";
  if(level == 1) {
    localUser.name = document.getElementById("name").value;
    localUser.email = document.getElementById("email").value;
    localUser.qq = document.getElementById("qq").value;
    localUser.url = document.getElementById("url").value;
    coms = document.getElementById("coms").value;
  } else {
    localUser.name = document.getElementById("rname").value;
    localUser.email = document.getElementById("remail").value;
    localUser.qq = document.getElementById("rqq").value;
    localUser.url = document.getElementById("rurl").value;
    coms = document.getElementById("rcoms").value;
  }
  var commObj = localUser;
  commObj.coms = coms;
  if (commObj.coms && commObj.name) {
    // cache
    var userCache = JSON.stringify(localUser);
    window.localStorage.setItem("localUser",userCache);
    // send comms
  } else {
    PopUp("昵称与评论内容为必填项",1,1)
  }
}
// initinput Base Info
function initUserInput() {
  document.getElementById("name").value = localUser.name;
  document.getElementById("rname").value = localUser.name;
  document.getElementById("email").value = localUser.email;
  document.getElementById("remail").value = localUser.email;
  document.getElementById("qq").value = localUser.qq;
  document.getElementById("rqq").value = localUser.qq;
  document.getElementById("url").value = localUser.url;
  document.getElementById("rurl").value = localUser.url;
}