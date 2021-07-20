var avtNos = 10;
var localUserStr = window.localStorage.getItem("localUser");
var localUser;
if(localUserStr) {
  localUser = JSON.parse(localUserStr);
}else{
  localUser = { name: "", email: "", qq: "", url: "", avt: ""};
}
if(localUser.avt == "") {
  localUser.avt = "/static/upload/avtor/"+randomAvt()+".jpg";
  document.getElementById("avt").src = localUser.avt;
}
function chooseAvt() {
	
}
function setAvt(no) {
	localUser.avt = "/static/upload/avtor/"+no+".jpg";
	document.getElementById("avt").src = localUser.avt;
}
function randomAvt() {
  return 1;//Math.floor(Math.random()*avtNos+1);
}
// PostID Level 123 ParentID
function sendComms(pid, level, fid) {
  // validate
  
  // cache
  var userCache = JSON.stringify(localUser);
  window.localStorage.setItem("localUser",userCache);
  // send comms
}