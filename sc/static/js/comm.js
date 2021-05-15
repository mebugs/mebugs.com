(function($) {
	$(document).ready(function(){
		setTimeout(function(){
			$(".wait").removeClass("wait");
		},1000);
	});
	$("#back_top").on("click",function(){
		$('body,html').animate({ scrollTop: 0 }, 1000);
	});
})(jQuery)