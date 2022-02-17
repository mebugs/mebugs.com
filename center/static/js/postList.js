var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    search: {title:null,cid:0,status:-1,size:8,page:1,pages:0,api:'ListPost'},
    category: [],
    status: [{id:-1,name:"全部"},{id:0,name:"草稿"},{id:1,name:"发布"},{id:2,name:"隐藏"}],
    statusNames: ['草稿', '发布', '隐藏'],
    list: []
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
          // 获取分类
          this_.getCategory();
          this_.getPage();
      	}
      })
    },
    setToken() {
      this.utoken = true;
      this.init();
    },
    getCategory() {
      CategoryWork({api:"ListCategory"}).then(res => {
        this.category = res.data;
        this.category.unshift({id:0,name:"请选择文章分类"})
      }).catch(function(err) {
        PopUp('分类查询失败',1,1);
      });  
    },
    add() {
      window.location.href = "postEdit.html?id=0"
    },
    edit(id) {
      window.location.href = "postEdit.html?id="+id
    },
    toPage(n) {
      if(this.search.page == n) {
        return
      }
      this.search.page = n
      this.getPage()
    },
    query() {
      this.search.page = 1
      this.getPage()
    },
    reset() {
      this.search = {title:null,cid:0,status:-1,size:8,page:1,pages:0,api:'ListPost'}
    },
    getPage() {
      PopUp('正在查询...',2,1 );
      PostWork(this.search).then(res => {
        PopUp('查询成功',0,1);
        this.list =  res.data.list;
        this.search.pages = res.data.pages;
      }).catch(function(err) {
        PopUp('查询失败',1,1);
      });
    }
	}
})