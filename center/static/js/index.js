var mindex = 0;
var urlArray = ["main.html","post/postList.html","category/categoryList.html"];
function toMenu(index) {
  if(index == mindex) {
    return;
  }
  $(".mck").removeClass("mck");
  $(".lmenu p").eq(index).addClass("mck");
  mindex = index;
  // 触发IFRAME变动
  $("#content").attr("src",urlArray[index]);
}

function showLogin() {
  $(".leare").show();
  setTimeout(function(){$("body").addClass("showLogin");},50);
}

function login() {
  var username = $("#username").val();
  var password = $("#password").val();
  if(username && password) {
    PopUp('正在登录...',2,1 );
    $.ajax({
      url: "/service/login.php",
      data: {username:username,password:password},
      dataType: "json",
      success: function(res) {
        if(res.code == 200) {
          PopUp('登录成功',0,1);
          // login sucess add token and add child token
          $("body").removeClass("showLogin");
          setTimeout(function(){ $(".leare").hide();},500);
          loc.setItem("MTOKEN",res.data);
          $("#content")[0].contentWindow.vue.setToken();
        } else {
          //登录失败
          PopUp(res.msg,1,1);
        }
      },
      error: function(err) {
        //系统异常
        PopUp('系统异常',0,1);
      }
    });
  }
  
}

