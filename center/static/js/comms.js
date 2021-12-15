var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    search: {size:8,page:1,pages:0,api:'PageComms'},
    item: {},
    list: [],
    status: ['未审核','已通过','已回复'],
    statusCss:['clred','clgreen','clblue'],
    ritem: {}
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
    edit(item) {
      this.item = item
      this.item.api = 'ModComms'
      this.opPop(1)
    },
    reback(item) { // 打开回复
      if(item.status < 1) {
        PopUp('审批后回复',1,1);
        return
      }
      this.item = item
      this.ritem = {api:'AdminReplayComms',name:'米虫',email:'iam@qiantaoyuan.cn',qq:'7431346',url:'http://www.mebugs.com',avt:'/static/img/me_avator.jpg',pid:item.pid,fid:item.id,level:2,coms:''}
      this.opPop(2)
    },
    replay() { // 发起回复
      PopUp('正在提交...',2,1 );
      CommsWork(this.ritem).then(res => {
        PopUp('更新成功',0,1);
        this.clsPop()
        this.getPage()
      }).catch(function(err) {
        PopUp('更新失败',1,1);
      });
    },
    opPop(flag) {
      if(flag == 1) {
        $(".popd").show();
      } else{
        $(".popb").show();
      }  
      setTimeout(function() {
        $("body").addClass("popup");
      },50)
    },
    clsPop() {
      $("body").removeClass("popup");
      setTimeout(function() {
        $(".popd").hide();
        $(".popb").hide();
      },500)
    },
    save(status) {
      this.item.status = status
      PopUp('正在提交...',2,1 );
      CommsWork(this.item).then(res => {
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
      CommsWork(this.search).then(res => {
        PopUp('查询成功',0,1);
        this.list =  res.data.list;
        this.search.pages = res.data.pages;
      }).catch(function(err) {
        PopUp('查询失败',1,1);
      });
    }
	}
})