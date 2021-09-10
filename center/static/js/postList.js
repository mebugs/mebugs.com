var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    search: {title:null,cid:0,status:0,size:10,page:1,pages:0},
    category: [],
    status: [{id:0,name:"请选择文章状态"},{id:1,name:"草稿"},{id:2,name:"发布"},{id:3,name:"下线"}],
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
      this.category = [{id:1,name:"java"},{id:2,name:"go"}]
      this.category.unshift({id:0,name:"请选择文章分类"})
    },
    add() {
      window.location.href = "postEdit.html?id=0"
    },
    edit(id) {
      
    },
    search() {
      this.getPage()
    },
    reset() {
      this.search = {title:null,cid:0,status:0,size:10,page:1,pages:0}
    },
    getPage() {
      this.search.pages = 10
      this.list = [
        {id:1,title:"测试标题1测试标题1测试标题1",banner:"/test/test1.jpg",remark:"测试标题1测试标题1测试标题1",cid:1,status:1,category:"Java",statusName:"草稿"},
        {id:2,title:"测试标题22啥打法是否该222",banner:"/test/test3.jpg",remark:"测试标题2321rreq3w4东风风光32sd试标题1",cid:1,status:2,category:"Go",statusName:"发布"},
        {id:3,title:"测试标题开发看电视了坚实的22222",banner:"/test/test9.jpg",remark:"测试标题2321rreq3w大幅度发鬼地方432sd试标题1",cid:2,status:3,category:"Go",statusName:"下线"},
        {id:4,title:"测试标题开发看电视了坚实的22222",banner:"/test/test9.jpg",remark:"测试标题2321rreq3w大幅度发鬼地方432sd试标题1",cid:2,status:3,category:"Go",statusName:"下线"},
        {id:5,title:"测试标题开发看电视了坚实的22222",banner:"/test/test9.jpg",remark:"测试标题2321rreq3w大幅度发鬼地方432sd试标题1",cid:2,status:3,category:"Go",statusName:"下线"},
        {id:6,title:"测试标题开发看电视了坚实的22222",banner:"/test/test9.jpg",remark:"测试标题2321rreq3w大幅度发鬼地方432sd试标题1",cid:2,status:3,category:"Go",statusName:"下线"},
        {id:7,title:"测试标题开发看电视了坚实的22222",banner:"/test/test9.jpg",remark:"测试标题2321rreq3w大幅度发鬼地方432sd试标题1",cid:2,status:3,category:"Go",statusName:"下线"},
        {id:8,title:"测试标题开发看电视了坚实的22222",banner:"/test/test9.jpg",remark:"测试标题2321rreq3w大幅度发鬼地方432sd试标题1",cid:2,status:3,category:"Go",statusName:"下线"},
        {id:9,title:"测试标题开发看电视了坚实的22222",banner:"/test/test9.jpg",remark:"测试标题2321rreq3w大幅度发鬼地方432sd试标题1",cid:2,status:3,category:"Go",statusName:"下线"},
        {id:10,title:"测试标题开发看电视了坚实的22222",banner:"/test/test9.jpg",remark:"测试标题2321rreq3w大幅度发鬼地方432sd试标题1",cid:2,status:3,category:"Go",statusName:"下线"},
      ]
    }
	}
})