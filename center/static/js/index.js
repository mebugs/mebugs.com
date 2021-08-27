var mindex = 0;
function toMenu(index) {
  if(index == mindex) {
    return;
  }
  $(".mck").removeClass("mck");
  $(".lmenu p").eq(index).addClass("mck");
  mindex = index;
  // 触发IFRAME变动
}

function showLogin() {
  $(".leare").show();
  setTimeout(function(){$("body").addClass("showLogin");},50);
}

function login() {
  // login sucess add token and add child token
  $("body").removeClass("showLogin");
  setTimeout(function(){ $(".leare").hide();},500);
  loc.setItem("MTOKEN","A");
  token = "A";
  $("#content")[0].contentWindow.vue.utoken = "A";
}

