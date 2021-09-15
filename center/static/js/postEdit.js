$(function(){
	
})
var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    category: [],
    status: [{id:0,name:"请选择文章状态"},{id:1,name:"草稿"},{id:2,name:"发布"},{id:3,name:"下线"}],
    hdTitle: '',
    tagSc: null,
    tagList: [],
    tagShow: [],
    tagShowFlag: false,
    tagCheck: [],
    newTag: {api:'AddTag',id:0,name:null,remark:null,url:null},
    fileName: null,
    post: {},
    ibann: '',
    ibanns: [null,null,null,null],
		md: null,
		html: '实时预览区',
		index: 0,
		toc: [],
		renderer: null,
		option: {},
		mtree: [],
		hideView: false,
		showMenu: false
	},
	watch: {
		md(val,oldVal) { // 监听MD输入框
		this.makeHtml()
		},
		toc(val,oldVal) { // 监听菜单变化
		this.makeMenus()
		}
	},
	mounted() {
		this.init()
	},
	methods: {
		init() { // 初始化
      var this_ = this;
      $(function(){
      	if(getToken(true)){
          // 获取分类
          this_.getCategory();
          // 获取标签
          this_.getAllTag();
          // 获取文章信息
          this_.getPostInfo();
          // 初始化markdown
          this_.initMarkDown();
      	}
      })
		},
    getPostInfo() {
      var id = getUrlParam("id");
      if(id != 0) {
        this.hdTitle = "编辑文章"
        this.md = "原文MARKDOWN文本"
        this.post = {id:id}
      } else {
        this.hdTitle = "新增文章"
        this.post = {id:0,cid:0,status:0}
      }
    },
    save() {
      // 数据检查和组装
      if(this.post.title && this.post.remark && this.md && this.post.url && this.post.cid && this.post.status) {
        if(this.tagCheck.length < 1) {
          PopUp("至少需要选择一个标签",1,1);
          return;
        }
        if(!this.fileName) {
          PopUp("请上传题图",1,1);
          return;
        }
        if(this.ibanns[0] && this.ibanns[1].indexOf('data:image') > -1) {
          PopUp('正在上传题图...',2,1 );
          ImgWork({api:"Base64",imgs:this.ibanns}).then(res => {
            if(res.data) {
              this.ibann = res.data;
              this.savePost();
            }
          })
        } else {
          this.savePost();
        }
      } else {
        PopUp("请输入或选择相关内容",1,1);
      } 
    },
    savePost() { // POST 赋值检查
      PopUp('res.msg',0,1);
    },
    setToken() {
      this.utoken = true;
      this.init();
    },
    getCategory() {
      this.category = [{id:1,name:"java"},{id:2,name:"go"}]
      this.category.unshift({id:0,name:"请选择文章分类"})
    },
    getAllTag() {
      TagWork({api:'AllTag'}).then(res => {
        this.tagList = res.data;
        this.tagShow = this.tagList;
      });
    },
    getScTag(e) {
      var val = e.target.value;
      if(val) {
        // 模糊大小写匹配
        this.tagShow = this.tagList.filter(function(item) { return item.name.toLowerCase().indexOf(val.toLowerCase()) != -1 ;});
      } else {
        this.tagShow = this.tagList
      }
      this.newTag = {api:'AddTag',id:0,name:val,remark:null,url:null};
      this.tagShowFlag = true;
    },
    pushTag(i) {
      let have = this.tagCheck.filter(function(item) { return item.id == i.id ;});
      if(have && have.length >0) {
        PopUp("该标签已存在",1,1);
      }else{
        this.tagCheck.push(i);
        this.tagShowFlag = false;
        this.tagSc = null;
        this.newTag = {api:'AddTag',id:0,name:null,remark:null,url:null};
      }
    },
    removeTag(id) {
      this.tagCheck = this.tagCheck.filter(function(item) { return item.id != id ;});
    },
    closeTag() {
      this.tagShowFlag = false;
    },
    addTag() {
      if(this.newTag.name && this.newTag.remark && this.newTag.url) {
        PopUp('正在提交...',2,1 );
        TagWork(this.newTag).then(res => {
          PopUp(res.msg,0,1);
          this.newTag.id = res.data;
          this.tagList.push(this.newTag);
          this.tagCheck.push(this.newTag);
          this.tagShowFlag = false;
          this.tagSc = null;
          this.newTag = {api:'AddTag',id:0,name:null,remark:null,url:null};
        });
      }else{
        PopUp("请输入全部内容后提交",1,1);
      } 
    },
    chooseIbann(e) { // 上传封面图片
      let file = e.target.files[0]
      //var file = document.getElementsById('ibann')[0].files[0];
			if (window.FileReader && file) //读取文件
			{
				let fileName = file.name;
				this.banners = [false,false,false];
				var this_ = this
				var reader = new FileReader();
				reader.readAsDataURL(file);
				//监听文件读取结束后事件
				reader.onloadend = function(e)
				{
          // 原图
          this_.ibanns = []
          this_.fileName = this_.post.banner ? this_.post.banner : fileName;
          this_.ibanns[0] = this_.fileName;
					this_.ibann = e.target.result;
          this_.ibanns[1] = this_.ibann;
					this_.resizeBanner(e.target.result,file.type);//压缩图片
				};
			}      
    },
    resizeBanner(url,type)//Banne图片400canvas压缩
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
				var blob = canvas.toDataURL(type)
				this_.ibanns[2] = blob;
        this_.resizeBanner100(blob,type)
		  }
		  image.src = url
		},
    resizeBanner100(url,type)//Banne图片100canvas压缩
    {
    	var this_ = this;
    	// 创建一个 Image 对象
      var image = new Image();
      // 绑定 load 事件处理器，加载完成后执行
      image.onload = function() {
        // 获取 canvas DOM 对象
        var canvas = document.createElement("canvas")
        //生成首页展示图
        if (image.width > 100) {
          // 宽度等比例缩放 *=
          image.height *= 100 / image.width
          image.width = 100
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
    		this_.ibanns[3] = blob;
      }
      image.src = url
    },
    initMarkDown() {
      this.renderer = new marked.Renderer();
      var this_ = this;
      // 重写标题
      this_.renderer.heading = function(text, level, raw) {
      	this_.index+=1;
      	this_.toc.push({ 'id':this_.index, 'level':level, 'text':text, 'child': [] });
      	return '<h'+level+' id="menu_'+this_.index+'">'+text+'</h'+level+'>';
      };
      //重写a标签，在新标签打开
      this_.renderer.link = function(href,title,text){
      	return '<a class="mlink" href="'+href+'" title="'+text+'" target="_blank">'+text+'</a>';
      }
      this_.option = {
      	renderer: this_.renderer, // 渲染器
      	breaks: true, // 开启回车换行能力
      	highlight: function(code) { return hljs.highlightAuto(code).value; }, // 代码高亮配置
      	langPrefix: "hljs language-", // 配合代码高亮补齐样式
      	xhtml: true // 自动闭合单标签
      	//headerPrefix: "menu_" // 为标题创建前缀 (通过渲染器改写标题)
      }
    },
		makeHtml() { // 生成HTML
		this.index = 0;
		this.toc = [];
		var this_ = this;
		this_.html = marked(this_.md,this_.option);
		},
		makeMenus() { // 生成菜单对象
		this.mtree = []
		this.toc.forEach(item => {
			if(item.level == 1){
			this.mtree.push(item)
			}
			if(item.level == 2) {
			let ml = this.mtree.length
			this.mtree[ml-1].child.push(item)
			}
			if(item.level == 3) {
			let ml = this.mtree.length
			let cl = this.mtree[ml-1].child.length
			this.mtree[ml-1].child[cl-1].child.push(item)
			}
		})
		console.log(this.mtree)
		},
		doView() { // 显示隐藏预览
		this.hideView = !this.hideView
    console.log($("#html").prop("outerHTML"));
		},
		doMenu() { // 显示隐藏目录
		this.showMenu = !this.showMenu
		},
		img(){
		
		}
	}
})