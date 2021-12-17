var friend = {
  iurl:"",
  furl:"",
  fname:"",
  fdesc:"",
  ficon:"" 
}

function getInfo(){
  getFriend();
  if(friend.iurl == "") {
    PopUp("请输入站点地址",1,1);
    return;
  }
  
  friend.api="CheckFirend";
  friend.timestamp=new Date().getTime();
  PopUp("读取中...",2,1)
  $.ajax({
    url:"/service/firendOpen.php",
    type:"POST",
    dataType:'json',
    contentType:'application/json;charset=UTF-8',
    data:JSON.stringify(friend),
    success:function(data, status){
      if(data.code == 200 ) {
        PopUp(data.msg,0,1);
        if(data.data) {
          var page = data.data;
          friend.fname = page.title?page.title:"";
          friend.fdesc = page.description?page.description:"";
          friend.ficon = page.icon?page.icon:friend.furl+"/favicon.ico";
          if(page.icon){
            if(page.icon.startsWith("/")) {
              friend.ficon = friend.furl+page.icon;
            } else if(!data.data.icon.startsWith("http")){
              friend.ficon = friend.furl+'/'+page.icon;
            } else {
              friend.ficon = page.icon;
            }
          }
          loadInfo();
        }
      }else{
        PopUp(data.msg,1,1);
      }
    },
    error:function(req,data, err){
      PopUp("请求错误",1,1);
    },
  })
}


// 同步输入框数据
function getFriend() {
  friend.iurl = document.getElementById("furl").value;
  if(friend.iurl.endsWith("/")) {
    friend.iurl = friend.iurl.substr(0,friend.iurl.length-1);
    document.getElementById("furl").value = friend.iurl;
  }
  friend.furl = document.getElementById("ffix").value + friend.iurl;
  friend.fname = document.getElementById("fname").value;
  friend.fdesc = document.getElementById("fdesc").value;
}

// 根据对象加载内容
function loadInfo(){
  document.getElementById("fname").value = friend.fname;
  document.getElementById("fdesc").value = friend.fdesc;
  document.getElementById("ficon").src = friend.ficon;
}