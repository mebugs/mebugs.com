// 图片名
var fileName;
var imgData;

// 图片选择触发
function chooseImg(nodeId) {
  document.getElementById(nodeId).click();
}

// 读取截取并压缩图片
function loadImg(node,showId) {
  let file = node.files[0]
  //var file = document.getElementsById('ibann')[0].files[0];
  if (window.FileReader && file) //读取文件
  {
    fileName = new Date().getTime()+Math.ceil(Math.random()*(999-99)+100) + file.name.substring(file.name.lastIndexOf('.'));
  	var reader = new FileReader();
  	reader.readAsDataURL(file);
  	//监听文件读取结束后事件
  	reader.onloadend = function(e)
  	{
      resizeImg(e.target.result,file.type,showId);//压缩图片
  	};
  }      
}

// 压缩图片
function resizeImg(url,type,showId) {
  // 创建一个 Image 对象
  var image = new Image();
  // 绑定 load 事件处理器，加载完成后执行
  image.onload = function() {
    // 获取 canvas DOM 对象
    var canvas = document.createElement("canvas")
    // 压缩120px
    // 如果高度超标
    if (image.height > 120 && image.height >= image.width) {
    	// 宽度等比例缩放 *
    	image.height *= 120 / image.width
    	image.width = 120
    }
    // 如果宽度超标
    if (image.width > 120 && image.width > image.height) {
    	// 宽度等比例缩放 *=
    	image.width *= 120 / image.height
    	image.height = 120
    } 
    // 获取 canvas的 2d 画布对象,
    var ctx = canvas.getContext("2d")
    // canvas清屏，并设置为上面宽高
    ctx.clearRect(0, 0, canvas.width, canvas.height)
  	// 重置canvas宽高
  	canvas.width = image.width
  	canvas.height = image.height
  	// 将图像绘制到canvas上
  	ctx.drawImage(image, 0, 0, image.width, image.height)
  	// !!! 注意，image 没有加入到 dom之中
  	var blob = canvas.toDataURL(type)
    resizeImg120(blob,type,showId)
  }
  image.src = url
}

// 截取方图
function resizeImg120(url,type,showId) {
  // 创建一个 Image 对象
  var image = new Image();
  // 绑定 load 事件处理器，加载完成后执行
  image.onload = function() {
    // 获取 canvas DOM 对象
    var canvas = document.createElement("canvas")
    // 获取 canvas的 2d 画布对象,
    var ctx = canvas.getContext("2d")
    // canvas清屏，并设置为上面宽高
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    // 重置canvas宽高
    canvas.width = 120
    canvas.height = 120
    // 将图像绘制到canvas上
    if(image.height > 120)
    {
      let ys = image.height / 2 - 60;
      ctx.drawImage(image, 0, ys, 120, 120, 0, 0, 120, 120);
    }else{
      let xs = image.width / 2 - 60;
      ctx.drawImage(image, xs, 0, 120, 120, 0, 0, 120, 120);
    }
    // !!! 注意，image 没有加入到 dom之中
    var blob = canvas.toDataURL(type)
    // 展示在Img元素上
    imgData = blob
    document.getElementById(showId).src = blob
  }
  image.src = url	   
}

// 图片上传 - 由业务调用
function uploadImg() {
  PopUp("图片上传中...",2,0)
  var img = {
    timestamp: new Date().getTime(),
    imgs: [fileName,imgData]
  }
  var imgSrc = false;
  $.ajax({
    url:"/service/imgOpen.php",
    type:"POST",
    dataType:'json',
    contentType:'application/json;charset=UTF-8',
    data:JSON.stringify(img),
    async: false,
    success:function(data, status){
      if(data.code == 200 ) {
        PopUp("图片上传成功",0,2);
        imgSrc = data.data;
      }else{
        PopUp(data.msg,1,2);
      }
    },
    error:function(req,data, err){
      PopUp("请求错误",1,2);
    },
  })
  return imgSrc
}