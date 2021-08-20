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