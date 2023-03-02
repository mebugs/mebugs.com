// 搜索标题
function scTit() {
  var scArray = $("#sc").serializeArray();
  var scObj = arrayToJsonObj(scArray);
  if(scObj.tit) {
    doFilter(scObj.tit)
  }else{
    renovate()
  }
  return false;
}
// 启动过滤
function doFilter(tit) {
  var allList = $("#plist a");
  var scNo = 0;
  if(allList && allList.length > 0) {
    for(var i=0;i<allList.length;i++) {
      var item = $(allList[i]);
      if(item.text().indexOf(tit) > -1) {
        item.show(400);
        scNo++;
      }else{
        item.hide(400);
      }
    }
  }
  $("#titVal").text(tit);
  $("#titNo").text(scNo);
  $("#sci").show(400);
}

// 刷新
function renovate() {
  $("#sci").hide(400);
  $("#titVal").innerText = "***";
  $("#titNo").innerText = 0;
  var allList = $("#plist a");
  if(allList && allList.length > 0) {
    for(var i=0;i<allList.length;i++) {
      $(allList[i]).show(400);
    }
  }
}