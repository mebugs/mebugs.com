var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    search: {size:8,page:1,pages:0,api:'PageFirends'},
    item: {},
    list: [],
    statusNames : ['未审核','已通过','已驳回'],
    statusCss:['clred','clgreen','clgray']
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
      this.item.api = 'ModFirend'
      this.opPop()
    },
    opPop() {
      $(".popd").show(); 
      setTimeout(function() {
        $("body").addClass("popup");
      },50)
    },
    clsPop() {
      $("body").removeClass("popup");
      setTimeout(function() {
        $(".popd").hide();
      },500)
    },
    save(status) {
      this.item.status = status
      PopUp('正在提交...',2,1 );
      FirendWork(this.item).then(res => {
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
      FirendWork(this.search).then(res => {
        PopUp('查询成功',0,1);
        this.list =  res.data.list;
        this.search.pages = res.data.pages;
      }).catch(function(err) {
        PopUp('查询失败',1,1);
      });
    }
	}
})