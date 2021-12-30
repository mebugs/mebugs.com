var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    loading: true,
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
          IndexWork().then(res => {
            if(res.data) {
              setTimeout(function() {
                this_.list = res.data;
                this_.loading = false;
              },500);
            }else {
              PopUp(res.msg,1,1);
            }
          }).catch(function(err) {
            PopUp('数据获取失败',1,1);
          });  
      	}
      })
    },
    setToken() {
      this.utoken = true;
      this.init()
    }
	}
})