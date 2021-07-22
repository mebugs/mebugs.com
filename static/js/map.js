function search() {
  var pname = document.getElementById("pname").value;
  if(pname) {
    refush();
    // All Post Node
    var allNode = document.querySelectorAll("#plist a h2");
    var scNo = 0;
    if(allNode && allNode.length > 0) {
      for(var i=0; i < allNode.length; i++) {
        let h2Titile = allNode[i].innerText;
        if(h2Titile.indexOf(pname) > -1) {
          scNo ++;
        }else{
          allNode[i].parentNode.classList.add("shid");
        }
      }
      document.getElementById("scPname").innerText=pname;
      document.getElementById("scNo").innerText=scNo;
       bodyNode.classList.add("msch");
    }
    bodyNode.classList.add("msch");
  }else{
    refush()
  }
}
function refush() {
  bodyNode.classList.remove("msch");
  document.getElementById("scPname").innerText="";
  document.getElementById("scNo").innerText="";
  // move all hide
  var hideNode = document.querySelectorAll("#plist .shid");
  if(hideNode && hideNode.length > 0) {
    for(var i=0; i < hideNode.length; i++) {
      hideNode[i].classList.remove("shid");
    }
  }
  document.getElementById("pname").value = "";
}