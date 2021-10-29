
var topWin= (function (p,c){
	while(p!=c){
		c = p		
		p = p.parent
	}
	return c
})(window.parent,window);
var popHave = false;
// 原生弹窗控件
function PopUp(text,type,time) {
  var iconHtml = '<span>✔</span>';
  if (type == 1) {
    iconHtml = '<span>✘</span>';
  }
  if (type == 2) {
    iconHtml = '<span>•</span>';
    time = 0;
  }
  if(topWin.document.getElementById("popeare")) {
    popHave = true;
  }
  if(popHave) {
    topWin.document.getElementById("popmsg").innerHTML = '<p>'+iconHtml+text+'</p>';
  }else{
    // 弹层元素
    var popDiv = topWin.document.createElement("div");
    popDiv.classList.add("popeare");
    popDiv.setAttribute("id","popeare");
    popDiv.innerHTML = '<div class="popmsg" id="popmsg"><p>'+iconHtml+text+'</p></div>';
    topWin.document.body.appendChild(popDiv);
    // 提示框样式
    var cssEle = topWin.document.createElement("style");
    cssEle.innerHTML = '.popeare,.popeare *{transition: all .3s cubic-bezier(0, 0.26, 0.53, 1.78);}.popeare{left:0;top:0;width:100vw;height:100vh;position:fixed;z-index:200;background:rgba(0,0,0,0);display:none;opacity:0}.popmsg{width:300px;background:rgba(0,0,0,0.5);height:100px;position:fixed;top:calc(50vh - 50px);transform: scale(0.8);left:calc(50vw - 150px);border-radius:10px;padding:10px}.popmsg p{line-height:40px;color:#eee;text-align:center;font-size:15px;font-weight:900}.popmsg p span{display:block;width:40px;height:40px;color:#ddd;margin:0 auto;border-radius:20px;font-size:26px;line-height:30px;border:3px solid #ddd;margin-top:5px}.showPop .popeare{opacity:1}.showPop .popmsg{transform: scale(1);}';
    topWin.document.body.appendChild(cssEle);
    popHave = true;
  }
  // 先移除弹出效果
  topWin.document.body.classList.remove("showPop");
  setTimeout(function(){
    // 显示弹层
    topWin.document.getElementById("popeare").style.display = "block";
    setTimeout(function(){
      // 触发延迟样式
      topWin.document.body.classList.add("showPop");
      if(time != 0) { // 0表示手工触发关闭
        setTimeout(function(){
          // 定时退出动画
          topWin.document.body.classList.remove("showPop");
          setTimeout(function(){
            // 隐藏元素
            topWin.document.getElementById("popeare").style.display = "none";
          },400);
        },time*1000);
      }
    },10);
  },10);
}