var vue = new Vue({
	el: '#main',
	data: {
    utoken: null
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
      $(function(){
      	if(getToken(true)){
      	  
      	}
      })
    },
    setToken() {
      this.init()
    }
	}
})