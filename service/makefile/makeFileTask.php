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
  // 查询足部数据
  $fteamSql = "SELECT * FROM `category` ORDER BY num DESC";
  $fteams = mysqli_fetch_all(mysqli_query($conn,$fteamSql),MYSQLI_ASSOC);
  $ftagSql = "SELECT * FROM `tag` ORDER BY num DESC LIMIT 30";
  $ftags = mysqli_fetch_all(mysqli_query($conn,$ftagSql),MYSQLI_ASSOC);
  // 生成首页 lv1
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makeIndex.php');
  $indexUrls = makeIndex($smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn);
  $siteMaps = array_merge_recursive($siteMaps,$indexUrls);
  // 生成文章详情页
  // 生成指定Page页（0-999）预留页
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makePost.php');
  $postUrls = makePosts($smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn);
  // 生成文章列表页
  
  // 生成分类页
  
  // 生成分类列表页
  
  // 生成标签页
  
  // 生成标签列表页
  
  // 生成SiteMap
  
  mysqli_close($conn);
}

// 测试代码
include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/doConfig.php');
$cdnUrl = getConfig("cdnUrl","./comm/system.php","string");
    $conns = [
      'qqUrl' => getConfig("qqUrl","./comm/system.php","string"),
      'githubUrl' => getConfig("githubUrl","./comm/system.php","string"),
      'giteeUrl' => getConfig("giteeUrl","./comm/system.php","string"),
      'emailUrl' => getConfig("emailUrl","./comm/system.php","string")
    ];
// 当前时间
      ini_set('date.timezone', 'Asia/Shanghai');
      $time_str = date('Y-m-d H:i:s', time());
      $today = substr($time_str,0,10);
makeFileTask($cdnUrl,$today,$conns);// 处理全局文件生成
exit;
?>