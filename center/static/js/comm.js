var loc = window.localStorage;
// req from where
function getToken(flag) {
  var token = loc.getItem("MTOKEN");
  if(flag) {
    // child page
    if(parent && window.top != window.self) {
      if(token) {
        vue.utoken = true;
        return true;
      } else {
        return parent.getToken(false);
      }
    }else{
      window.close();
      return false;
    }
  }
  if(token) {
    return true;
  }else{
    showLogin();
    return false;
  }
}

function getUrlParam(name) {
  var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
  var r = window.location.search.substr(1).match(reg);  //匹配目标参数
  if (r != null) return unescape(r[2]); return null; //返回参数值
}