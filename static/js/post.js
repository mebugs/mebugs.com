var avtNos = 12;
var localUserStr = window.localStorage.getItem("localUser");
var localUser;
var crTip = document.getElementById("rwork");
var crMet = document.getElementById("repMethod");
var rzz = document.getElementById("rzz");
var cavt = document.getElementById("cavt");
var myavt = document.getElementById("avt");
var chelp = document.getElementById("chelp");
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
  clsAvt();
}
function clsAvt() {
  bodyNode.classList.remove("cavts");
  setTimeout(function() {
    cavt.style.display = "none";
  }, 400);
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
  }, 400);
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
    PopUp("提交中...",2,1)
    // cache
    var userCache = JSON.stringify(localUser);
    window.localStorage.setItem("localUser",userCache);
    // send comms
    commObj.pid = pid
    commObj.fid = fid
    commObj.level = level
    commObj.timestamp=new Date().getTime()
    $.ajax({
      url:"/service/commsOpen.php",
      type:"POST",
      dataType:'json',
      contentType:'application/json;charset=UTF-8',
      data:JSON.stringify(commObj),
      success:function(data, status){
        if(data.code == 200 ) {
          PopUp(data.msg,0,1);
          var coms = "";
          if(level == 1) {
            document.getElementById("coms").value = "";
          } else {
            document.getElementById("rcoms").value = "";
            clsComms();
          }
        }else{
          PopUp(data.msg,1,1);
        }
      },
      error:function(req,data, err){
        PopUp("请求错误",1,1);
      },
    })
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
// order排序查询
var page = 1; // 默认第二页
var pages = 1; // 总页数存在查看更多会返回新的pages
var order = "desc";// 默认倒序
function toCommsOrder(flag) {
  // 切换排序需要清空显示区域的元素
  $("#commea").empty();
  $(".cko").removeClass("cko");
  if(flag) {
    $(".ckOrder").children("span").eq(0).addClass("cko");
  }else{
    $(".ckOrder").children("span").eq(1).addClass("cko");
  }
  $("#comodoo").hide();
  $("#comonul").hide();
  page = 1;
  order = flag?"desc":"asc";
  // 查询分页数据
  getCommsList();
}
function moreComs() {
  page++;
  $("#comodoo").hide();
  $("#comonul").hide();
  // 查询分页数据
  getCommsList();
}
function getCommsList() {
  var pageQuery = {
    page : page,
    order: order,
    pid : pid,
    timestamp : new Date().getTime()
  }
  $("#comolod").show();
  $.ajax({
    url:"/service/commsQueryOpen.php",
    type:"POST",
    dataType:'json',
    contentType:'application/json;charset=UTF-8',
    data:JSON.stringify(pageQuery),
    success:function(data, status){
      $("#comolod").hide();
      if(data.code == 200 ) {
        pages =  data.data.pages;
        if(pages < page) {
          PopUp("没有更多数据啦",1,1);
        }else if(pages == page) {
          $("#comonul").show();
        }else{
          $("#comodoo").show();
        }
        var list = data.data.list;
        if(list != null && list.length > 0) {
          var html = "";
          list.forEach(i=>{
            var liHtml = '<li> <div class="comli"> <img src="'+i.avt+'"/><div class="comp"><h1><b><a target="_blank" href="'+i.url+'">'+i.name+'</a></b>丨'+i.send_time+' 发表观点丨<a href="javascript:sendCommsR('+pid+',2,'+i.id+')">评论TA</a></h1><p>'+i.coms+'</p></div></div>';
            if(i.child && i.child.length >0) {
              var child = i.child;
              liHtml += '<ul>';
              child.forEach(c=>{
                var cLiHtml = '<li> <div class="comli"> <div class="comp"> <h1> <b><a target="_blank" href="'+c.url+'">'+c.name+'</a></b>丨'+c.send_time;
                if(c.fid == i.id) {
                  cLiHtml += ' 评论本楼丨';
                }else{
                  cLiHtml += ' 回复<b><a target="_blank" href="'+c.pUrl+'">'+c.pName+'</a></b>丨';
                }
                cLiHtml += '<a href="javascript:sendCommsR('+pid+',3,'+c.id+')">回复TA</a></h1> </h1><p>'+c.coms+'</p> </div> </div> </li>';
                liHtml += cLiHtml;
              });
              liHtml += '</ul>';
            }
            liHtml += '</li>';
            html += liHtml;
          });
          $("#commea").append(html);
        }
      }else{
        PopUp(data.msg,1,1);
      }
    },
    error:function(req,data, err){
      $("#comolod").hide(400);
      PopUp("请求错误",1,1);
    },
  })
  
}

function showAll() {
  $(".commbox").removeClass("zdbox");
}

ready(makeZd);
// 如果评论区长度高于400 默认启用折叠
function makeZd() {
  var commbox = $(".commbox");
  if(commbox) {
    if(commbox.height() > 400) {
      commbox.addClass("zdbox");
    }
  }
}

// 打开评论提示
function helpComms() {
  chelp.style.display = "block";
  setTimeout(function() {
    bodyNode.classList.add("shelp");
  }, 100);
}

function clsHelpComms() {
  bodyNode.classList.remove("shelp");
  setTimeout(function() {
    chelp.style.display = "none";
  }, 400);
}

if(pid > 999) {
  addRead("c");
  setInterval(function() {
    addRead("g");
  },200000);
}

function addRead(a) {
  $.ajax({
    url:"/service/addData.php",
    type:"POST",
    dataType:'json',
    contentType:'application/json;charset=UTF-8',
    data:JSON.stringify({pid:pid,api:a}),
    success:function(data, status){},
    error:function(req,data, err){},
  })
}