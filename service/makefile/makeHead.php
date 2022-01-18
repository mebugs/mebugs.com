<!DOCTYPE html>
<html>
 <head> 
  <meta charset="UTF-8" /> 
  <meta http-equiv="X-UA-Compatible" content="IE=edge" /> 
  <meta name="renderer" content="webkit" /> 
  <meta name="viewport" content="width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" /> 
  <!-- <meta name="referrer" content="no-referrer"/> --> 
  <meta name="apple-mobile-web-app-capable" content="yes" /> 
  <title><?php echo $title; ?></title> 
  <meta name="expires" content="Mon, 09 Apr 2999 09:09:09 GMT" /> 
  <meta name="description" content="<?php echo $description; ?>" /> 
  <meta name="keywords" content="米虫,程序员,博客,学习,工作,物联网,IT,经验,总结,分享,Java,Go,Linux,前端,数据库" /> 
  <link rel="icon" href="<?php echo $cdnUrl; ?>/static/img/icon.png" /> 
  <link rel="shortcut icon" href="<?php echo $cdnUrl; ?>/static/img/icon.png" /> 
  <link rel="apple-touch-icon-precomposed" href="<?php echo $cdnUrl; ?>/static/img/icon.png" /> 
  <link href="<?php echo $cdnUrl; ?>/static/css/load.css?v=1.0" rel="stylesheet" type="text/css" />
  <link href="<?php echo $cdnUrl; ?>/static/css/base.css?v=1.0" rel="stylesheet" type="text/css" />
  <?php if($pageType == "index") { ?>
  <link href="<?php echo $cdnUrl; ?>/static/css/index.css?v=1.0" rel="stylesheet" type="text/css" /> 
  <link href="<?php echo $cdnUrl; ?>/static/lib/swiper/swiper-bundle.min.css?v=1.0" rel="stylesheet" type="text/css" />
  <?php } ?>
  <?php if($pageType == "post") { ?>
  <link href="<?php echo $cdnUrl; ?>/static/css/post.css?v=1.0" rel="stylesheet" type="text/css" />
  <link href="<?php echo $cdnUrl; ?>/static/css/coms.css?v=1.0" rel="stylesheet" type="text/css" /> 
  <link href="<?php echo $cdnUrl; ?>/static/css/marked/marked.css?v=1.0" rel="stylesheet" type="text/css" /> 
  <?php } ?>
  <link href="<?php echo $cdnUrl; ?>/static/css/mobile/base.css?v=1.0" rel="stylesheet" type="text/css" media="screen and (max-device-width:1200px)" /> 
  <?php if($pageType == "index") { ?><script type="text/javascript">
    // WaitForSVGInjectInit
    function initSvg(obj) {
    	try{
    	    SVGInject(obj);
    	}catch(e){
    	    setTimeout(function(){
    			initSvg(obj);
    		},200)
    	}
    }
  </script><?php } ?>
 </head> 
 <body> 
 <div class="loader" id="loader">
   <div class="load">
     <hr/><hr/><hr/><hr/>
   </div>
 </div>
  <header> 
   <div class="w"> 
    <a class="logo" href="/"></a> 
    <div class="menu"> 
     <a <?php if($pageType == "index") { ?>class="mck"<?php } ?> href="/">首页</a> 
     <a <?php if($pageSubType == "post") { ?>class="mck"<?php } ?> href="/posts/new_1.html">文章</a> 
     <a <?php if($pageSubType == "category") { ?>class="mck"<?php } ?> href="/category">分类</a> 
     <a <?php if($pageSubType == "tag") { ?>class="mck"<?php } ?> href="/tag">标签</a> 
     <a <?php if($pageSubType == "about") { ?>class="mck"<?php } ?> href="/page/about.html">关于</a> 
     <a <?php if($pageType == "link") { ?>class="mck"<?php } ?> href="/page/link.html">友链</a> 
     <a <?php if($pageSubType == "msg") { ?>class="mck"<?php } ?> href="/page/msg.html">留言</a> 
     <a <?php if($pageType == "map") { ?>class="mck"<?php } ?> href="/page/map.html">地图</a> 
    </div> 
    <div class="temp" onclick="setShow()"></div><div class="mobim" onclick="setMbm()"><ul><li></li><li></li><li></li></ul></div> 
    <?php if($pageType == "post") { ?>
    <div class="pomenu" onclick="doPmenu(true)"><div class="mmenu" id="pomenu" ><ul><li>‹</li><li></li><li>›</li></ul></div></div>
    <?php } ?>
   </div> 
  </header>