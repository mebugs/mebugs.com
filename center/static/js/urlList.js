var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    search: {name:null,size:8,page:1,pages:0,api:'PageUrls'},
    item: {id:0,api:'UpsertUrls'},
    list: [],
    title: '',
    banner: [],
    ibann: null,
    fileType: null
	},
  watch: {
  },
	mounted() {
    this.init();
	},
	methods: {
		init() {
      var this_ = this;
      $(function(){
      	if(getToken(true)){
          this_.getPage();
      	}
      })
    },
    setToken() {
      this.utoken = true;
      this.init();
    },
    add() {
      this.item = {id:0,api:'UpsertUrls'}
      this.banner = []
      this.ibann = null
      this.title = '新增链接'
      this.opPop()
    },
    edit(item) {
      this.banner = []
      this.ibann = item.img
      this.item = item
      this.item.api = 'UpsertUrls'
      this.title = '编辑链接'
      this.opPop()
    },
    opPop() {
      $(".pope").show();
      setTimeout(function() {
        $("body").addClass("popup");
      },50)
    },
    clsPop() {
      $("body").removeClass("popup");
      setTimeout(function() {
        $(".pope").hide();
      },500)
    },
    save() {
      PopUp('正在提交...',2,1 );
      if(this.banner[0] && this.banner[1].indexOf('data:image') > -1) {
        if(this.item.type == "smaller") { // 启用压缩
          this.resizeBanner()
        } else {
          this.saveImg()
        }
      } else {
        this.saveUrls();
      }
    },
    chooseIbann(e) { // 选择图片
      let file = e.target.files[0]
      //var file = document.getElementsById('ibann')[0].files[0];
    	if (window.FileReader && file) //读取文件
    	{
    		let fileName = new Date().getTime()+Math.ceil(Math.random()*(999-99)+100) + file.name.substring(file.name.lastIndexOf('.'));
    		var this_ = this
    		var reader = new FileReader();
    		reader.readAsDataURL(file);
    		//监听文件读取结束后事件
    		reader.onloadend = function(e)
    		{
          // 原图
          this_.banner = []
          this_.banner[0] = this_.item.img ? this_.item.img.substring(this_.item.img.lastIndexOf("\/")+1,this_.item.img.length) : fileName;
          this_.banner[1] = e.target.result;
          this_.ibann = e.target.result;
          this_.fileType = file.type;
    		};
    	}      
    },
    resizeBanner()//Banne图片400canvas压缩
    {
    	var this_ = this;
    	// 创建一个 Image 对象
      var image = new Image();
      // 绑定 load 事件处理器，加载完成后执行
      image.onload = function() {
        // 获取 canvas DOM 对象
        var canvas = document.createElement("canvas")
        //生成首页展示图
        if (image.width > 400) {
          // 宽度等比例缩放 *=
          image.height *= 400 / image.width
          image.width = 400
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
    		var blob = canvas.toDataURL(this_.fileType)
    		this_.banner[1] = blob;
        this_.saveImg()
      }
      image.src = this_.banner[1]
    },
    saveImg() {
      PopUp('正在上传图片...',2,1 );
      ImgWork({api:"Base64Urls",imgs:this.banner}).then(res => {
        if(res.data) {
          this.ibann = res.data+'?v='+Math.ceil(Math.random()*100);
          this.banner[1] = res.data
          this.item.img = res.data
          this.saveUrls();
        }
      })
    },
    saveUrls() {
      UrlsWork(this.item).then(res => {
        PopUp('更新成功',0,1);
        this.clsPop()
        this.getPage()
      }).catch(function(err) {
        PopUp('更新失败',1,1);
      });
    },
    toPage(n) {
      if(this.search.page == n) {
        return
      }
      this.search.page = n
      this.getPage()
    },
    query() {
      this.getPage()
    },
    toPage(n) {
      if(this.search.page == n) {
        return
      }
      this.search.page = n
      this.getPage()
    },
    getPage() {
      PopUp('正在查询...',2,1 );
      UrlsWork(this.search).then(res => {
        PopUp('查询成功',0,1);
        this.list =  res.data.list;
        this.search.pages = res.data.pages;
      }).catch(function(err) {
        PopUp('查询失败',1,1);
      });
    }
	}
})