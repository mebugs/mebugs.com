<?php
// 生成文件
function makeFileTask($cdnUrl,$today,$conns) {
  // 初始化连接
  include($_SERVER['DOCUMENT_ROOT'].'/service/comm/connect.php');
  // 初始化统计
  initStatics($conn,$today);
  $baseUrl = "http://www.mebugs.com";
  // 创建siteMap数组
  $siteMaps = [];
  // 查询smallImgs
  $smliSql = "SELECT * FROM `urls` WHERE `status` = 1 AND `type` = 'smaller' ORDER BY `sorts`";
  $smliImgs = mysqli_fetch_all(mysqli_query($conn,$smliSql),MYSQLI_ASSOC);
  // 查询页面底部推送数据
  $newPostSql = "SELECT * FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY id DESC LIMIT 10";
  $newPosts = mysqli_fetch_all(mysqli_query($conn,$newPostSql),MYSQLI_ASSOC);
  $hotPostSql = "SELECT * FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY monthView DESC LIMIT 10";
  $hotPosts = mysqli_fetch_all(mysqli_query($conn,$hotPostSql),MYSQLI_ASSOC);
  $radPostSql = "SELECT * FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY randno DESC LIMIT 10";
  $radPosts = mysqli_fetch_all(mysqli_query($conn,$radPostSql),MYSQLI_ASSOC);
  $doPosts = ["new" => $newPosts,"hot" => $hotPosts,"rad" => $radPosts];
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
  $postUrls = makePosts($smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn,$doPosts);
  $siteMaps = array_merge_recursive($siteMaps,$postUrls);
  // 生成文章列表页
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makePosts.php');
  $listUrls = makePostLists($fteams,$ftags,$cdnUrl,$baseUrl,$conn);
  $siteMaps = array_merge_recursive($siteMaps,$listUrls);
  // 生成分类页
  // 生成分类列表页
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makeCategory.php');
  $cateUrls = makeCategoryAndList($fteams,$ftags,$cdnUrl,$baseUrl,$conn);
  $siteMaps = array_merge_recursive($siteMaps,$cateUrls);
  // 生成标签页
  // 生成标签列表页
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makeTag.php');
  $tagUrls = makeTagAndList($fteams,$ftags,$cdnUrl,$baseUrl,$conn);
  $siteMaps = array_merge_recursive($siteMaps,$tagUrls);
  // 生成其他页面
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makePages.php');
  $friendUrls = makeFriends($fteams,$ftags,$cdnUrl,$baseUrl,$conn);
  $siteMaps = array_merge_recursive($siteMaps,$friendUrls);
  $friendSendUrls = makeSendFriends($fteams,$ftags,$cdnUrl,$baseUrl,$conn);
  $siteMaps = array_merge_recursive($siteMaps,$friendSendUrls);
  $mapUrls = makeMap($fteams,$ftags,$cdnUrl,$baseUrl,$conn);
  $siteMaps = array_merge_recursive($siteMaps,$mapUrls);
  // 生成SiteMap
  include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makeSiteMap.php');
  makeSiteMap($siteMaps);
  // 测试显示
  echo json_encode($siteMaps);
  mysqli_close($conn);
}

// 初始化统计
function initStatics($conn,$today) {
  // 统计分组开放数量
  mysqli_query($conn,"UPDATE `category` tm SET tm.num = (
  	SELECT ptm.nums FROM (
  		SELECT t.id,IFNULL(pt.nums,0) AS nums FROM `category` t
  		LEFT JOIN (SELECT cid,count(id) AS nums FROM `post_main` WHERE status = 1 GROUP BY cid) AS pt ON t.id = pt.cid
  	) ptm 
  	WHERE ptm.id = tm.id
  )");
  // 统计标签开放数量
  mysqli_query($conn,"UPDATE `tag` tm SET tm.num = (
    SELECT ptm.nums FROM (
      SELECT t.id,IFNULL(ptc.nums,0) AS nums FROM `tag` t
      LEFT JOIN (SELECT tid,count(pid) AS nums FROM (SELECT pt.* FROM `post_tag` pt LEFT JOIN `post_main` m ON m.id = pt.pid WHERE m.`status` = 1) AS ptp GROUP BY tid) AS ptc ON t.id = ptc.tid
	) ptm 
    WHERE ptm.id = tm.id
  )");
  // 随机生成排序值
  mysqli_query($conn,"update `post_main` set randno = floor(rand()*1000)");
  // 逢5对近期飙升压缩至10%
  $dayNo = substr($today,9,1);
  if($dayNo == "5") {
    mysqli_query($conn,"update `post_main` set monthView = floor(monthView/10)");
  }
}

// 测试代码
// include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/doConfig.php');
// $cdnUrl = getConfig("cdnUrl","../comm/system.php","string");
// $conns = [
//  'qqUrl' => getConfig("qqUrl","../comm/system.php","string"),
//  'githubUrl' => getConfig("githubUrl","../comm/system.php","string"),
//  'giteeUrl' => getConfig("giteeUrl","../comm/system.php","string"),
//  'emailUrl' => getConfig("emailUrl","../comm/system.php","string")
// ];
// // 当前时间
// ini_set('date.timezone', 'Asia/Shanghai');
// $time_str = date('Y-m-d H:i:s', time());
// $today = substr($time_str,0,10);
// echo $time_str;
// //setConfig("taskLastRun",$time_str,"../comm/system.php","string");// 更新执行时间
// makeFileTask($cdnUrl,$today,$conns);// 处理全局文件生成
// exit;
?>