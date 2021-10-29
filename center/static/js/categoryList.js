var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    item: {id:0,api:'UpsertCategory'},
    list: [],
    title: ''
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
      this.item = {id:0,api:'UpsertCategory'}
      this.title = '新增分类'
      this.opPop()
    },
    edit(item) {
      this.item = item
      this.item.api = 'UpsertCategory'
      this.title = '编辑分类'
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
      CategoryWork(this.item).then(res => {
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
    reset() {
      this.search = {title:null,cid:0,status:0,size:10,page:1,pages:0,api:'ListPost'}
    },
    getPage() {
      PopUp('正在查询...',2,1 );
      CategoryWork({api:'ListCategory'}).then(res => {
        PopUp('查询成功',0,1);
        this.list =  res.data;
      }).catch(function(err) {
        PopUp('查询失败',1,1);
      });
    }
	}
})