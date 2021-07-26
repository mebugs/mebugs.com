$(function(){
	
})
new Vue({
	el: '#content',
	data: {
		md: '',
		html: '',
		index: 0,
		toc: [],
		renderer: null,
		option: {},
		mtree: [],
		hideView: false,
		showMenu: false
	},
	watch: {
		md(val,oldVal) { // 监听MD输入框
		this.makeHtml()
		},
		toc(val,oldVal) { // 监听菜单变化
		this.makeMenus()
		}
	},
	mounted() {
		this.init()
	},
	methods: {
		init() { // 初始化
		this.renderer = new marked.Renderer();
		var this_ = this;
		// 重写标题
		this_.renderer.heading = function(text, level, raw) {
			this_.index+=1;
			this_.toc.push({ 'id':this_.index, 'level':level, 'text':text, 'child': [] });
			return '<h'+level+' id="menu_'+this_.index+'">'+text+'</h'+level+'>';
		};
		//重写a标签，在新标签打开
		this_.renderer.link = function(href,title,text){
			return '<a class="mlink" href="'+href+'" title="'+text+'" target="_blank">'+text+'</a>';
		}
		this_.option = {
			renderer: this_.renderer, // 渲染器
			breaks: true, // 开启回车换行能力
			highlight: function(code) { return hljs.highlightAuto(code).value; }, // 代码高亮配置
			langPrefix: "hljs language-", // 配合代码高亮补齐样式
			xhtml: true // 自动闭合单标签
			//headerPrefix: "menu_" // 为标题创建前缀 (通过渲染器改写标题)
		}
		},
		makeHtml() { // 生成HTML
		this.index = 0;
		this.toc = [];
		var this_ = this;
		this_.html = marked(this_.md,this_.option);
		},
		makeMenus() { // 生成菜单对象
		this.mtree = []
		this.toc.forEach(item => {
			if(item.level == 1){
			this.mtree.push(item)
			}
			if(item.level == 2) {
			let ml = this.mtree.length
			this.mtree[ml-1].child.push(item)
			}
			if(item.level == 3) {
			let ml = this.mtree.length
			let cl = this.mtree[ml-1].child.length
			this.mtree[ml-1].child[cl-1].child.push(item)
			}
		})
		console.log(this.mtree)
		},
		doView() { // 显示隐藏预览
		this.hideView = !this.hideView
    console.log($("#html").prop("outerHTML"));
		},
		doMenu() { // 显示隐藏目录
		this.showMenu = !this.showMenu
		},
		img(){
		
		}
	}
})