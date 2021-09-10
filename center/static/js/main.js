var vue = new Vue({
	el: '#main',
	data: {
    utoken: null
	},
  watch: {
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
      this.utoken = true;
      this.init()
    }
	}
})