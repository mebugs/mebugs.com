var mindex = 0;
var urlArray = ["main.html","post/postList.html"];
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
    $.ajax({
      url: "/service/login.php",
      data: {username:username,password:password},
      dataType: "json",
      success: function(res) {
        if(res.code == 200) {
          // login sucess add token and add child token
          $("body").removeClass("showLogin");
          setTimeout(function(){ $(".leare").hide();},500);
          loc.setItem("MTOKEN",res.data);
          token = res.data;
          $("#content")[0].contentWindow.vue.utoken = res.data;
        } else {
          //登录失败
        }
      },
      error: function(err) {
        //系统异常
      }
    });
  }
  
}

