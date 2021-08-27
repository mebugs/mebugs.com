var loc = window.localStorage;
var token;
// req from where
function getToken(flag) {
  if(flag) {
    // child page
    if(parent && window.top != window.self) {
      if(parent.getToken(false))
      {
        vue.utoken=parent.token;
        return true;
      }else{
        return false;
      }
    }else{
      alert("非法访问！");
      window.close();
      return false;
    }
  }else{
    token = loc.getItem("MTOKEN");
  }
  if(token) {
    return true;
  }else{
    showLogin();
    return false;
  }
}