var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    search: {title:null,category:0,status:0,size:10,page:1},
    category: [],
    status: [{id:0,name:"请选择文章状态"},{id:1,name:"草稿"},{id:2,name:"发布"},{id:3,name:"下线"}]
	},
  watch: {
  	utoken(val,oldVal) { // 监听TOKEN
      this.setToken(val);
  	}
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
      	}
      })
    },
    setToken() {
      this.init();
    },
    getCategory() {
      this.category = [{id:1,name:"java"},{id:2,name:"go"}]
      this.category.unshift({id:0,name:"请选择文章分类"})
    }
	}
})