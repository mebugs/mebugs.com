<?php
// 生成文件
function makeFileTask($cdnUrl,$today,$conns) {
  // 初始化连接
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/connect.php');
  $baseUrl = "http://www.mebugs.com";
  // 创建siteMap数组
  $siteMaps = [];
  // 查询smallImgs
  $smliSql = "SELECT * FROM `urls` WHERE `status` = 1 AND `type` = 'smaller' ORDER BY `sorts`";
  $smliImgs = mysqli_fetch_all(mysqli_query($conn,$smliSql),MYSQLI_ASSOC);
  // 生成首页 lv1
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makeIndex.php');
  // 生成文章详情页
  $indexUrl = makeIndex($smliImgs,$cdnUrl,$baseUrl,$conns,$conn);
  // 生成指定Page页（0-999）预留页
  
  // 生成文章列表页
  
  // 生成分类页
  
  // 生成分类列表页
  
  // 生成标签页
  
  // 生成标签列表页
  
  // 生成SiteMap
  
  mysqli_close($conn);
}
?>