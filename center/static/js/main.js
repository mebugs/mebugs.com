var vue = new Vue({
	el: '#main',
	data: {
    utoken: null,
    loading: true,
    list: [],
    sysSet: {},
    sysView: true
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
          this_.getConfig()
      	}
      })
    },
    getConfig() {
      SystemWork({api:'GetConf'}).then(res => {
        if(res.data) {
          this.sysSet = res.data
        }else {
          PopUp(res.msg,1,1);
        }
      }).catch(function(err) {
        PopUp('数据获取失败',1,1);
      });  
    },
    toSave() {
      this.sysView = false
    },
    runTask() {
      PopUp('正在提交...',2,1 );
      RunTask().then(res => {
        PopUp('启动成功',0,1);
        this.sysView = true;
        this.getConfig();
      }).catch(function(err) {
        PopUp('启动失败',1,1);
      });
    },
    save() {
      PopUp('正在提交...',2,1 );
      this.sysSet.api = 'SetConf'
      SystemWork(this.sysSet).then(res => {
        PopUp('更新成功',0,1);
        this.sysView = true
        this.getConfig()
      }).catch(function(err) {
        PopUp('更新失败',1,1);
      });
    },
    setToken() {
      this.utoken = true;
      this.init()
    }
	}
})