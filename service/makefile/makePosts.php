<?php 
//生成INDEX页面代码 CDN切换 数据库对象
function makePostLists($fteams,$ftags,$cdnUrl,$baseUrl,$conn) {
  // 查询全部启用文章的数量用于构建分页
  $backUrls = [];
  // 查询全部开放文章
  $countSql = "SELECT count(id) FROM `post_main` m WHERE m.`status` = 1";
  $count = mysqli_fetch_array(mysqli_query($conn,$countSql))[0];
  $size = 15;
  $pages = ceil($count/$size);
  // 查询最新文章
  for($n = 1;$n <= $pages;$n ++) {
    // 查询当页数据
    $querySql = "SELECT * FROM `post_main` m WHERE m.`status` = 1 ORDER BY id DESC LIMIT ".($n-1)*$size.",".$size;
    $postsType = "new";
    $tit = "新鲜发布";
    makePostList($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$querySql,$postsType,$tit,$n,$pages);
    array_push($backUrls,["loc" => $baseUrl."/posts/new_".$n.".html", "priority" => 0.8]);
  }
  // 查询全站荐读
  for($n = 1;$n <= $pages;$n ++) {
    // 查询当页数据
    $querySql = "SELECT * FROM `post_main` m WHERE m.`status` = 1 ORDER BY depth DESC LIMIT ".($n-1)*$size.",".$size;
    $postsType = "best";
    $tit = "好文荐读";
    makePostList($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$querySql,$postsType,$tit,$n,$pages);
    array_push($backUrls,["loc" => $baseUrl."/posts/best_".$n.".html", "priority" => 0.8]);
  }
  // 查询近期飙升
  for($n = 1;$n <= $pages;$n ++) {
    // 查询当页数据
    $querySql = "SELECT * FROM `post_main` m WHERE m.`status` = 1 ORDER BY monthView DESC LIMIT ".($n-1)*$size.",".$size;
    $postsType = "up";
    $tit = "近期飙升";
    makePostList($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$querySql,$postsType,$tit,$n,$pages);
    array_push($backUrls,["loc" => $baseUrl."/posts/up_".$n.".html", "priority" => 0.8]);
  }
  // 查询全站热门
  for($n = 1;$n <= $pages;$n ++) {
    // 查询当页数据
    $querySql = "SELECT * FROM `post_main` m WHERE m.`status` = 1 ORDER BY view DESC LIMIT ".($n-1)*$size.",".$size;
    $postsType = "hot";
    $tit = "全站热门";
    makePostList($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$querySql,$postsType,$tit,$n,$pages);
    array_push($backUrls,["loc" => $baseUrl."/posts/hot_".$n.".html", "priority" => 0.8]);
  }
  // 查询今日随选
  for($n = 1;$n <= $pages;$n ++) {
    // 查询当页数据
    $querySql = "SELECT * FROM `post_main` m WHERE m.`status` = 1 ORDER BY randno DESC LIMIT ".($n-1)*$size.",".$size;
    $postsType = "rand";
    $tit = "今日随选";
    makePostList($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$querySql,$postsType,$tit,$n,$pages);
    array_push($backUrls,["loc" => $baseUrl."/posts/rand_".$n.".html", "priority" => 0.8]);
  }
  return $backUrls;
} 

function makePostList($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$querySql,$postType,$tit,$page,$pages) {
  $postsUrl = "/posts/".$postType."_".$page.".html";
  $file_sc = $_SERVER['DOCUMENT_ROOT'].$postsUrl;
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  $title = "[".$tit."] - 第".$page."页"." - 米虫 - 做一个有理想的米虫 - www.mebugs.com";
  $description = $tit." - 程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！";
  $pageType = "postList";//文章页
  $pageSubType = "post";//文章列表页
  include(__DIR__.'/makeHead.php');
?>
  <!-- posts top -->
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_12"> 
     <div class="ql pt pct rtl"> 
      <a href="/posts/new_1.html" <?php if($postType == "new") { echo 'class="cp"';} ?>>新鲜</a>
      <a href="/posts/best_1.html" <?php if($postType == "best") { echo 'class="cp"';} ?>>荐读</a> 
      <a href="/posts/up_1.html" <?php if($postType == "up") { echo 'class="cp"';} ?>>飙升</a>
      <a href="/posts/hot_1.html" <?php if($postType == "hot") { echo 'class="cp"';} ?>>热门</a>
      <a href="/posts/rand_1.html" <?php if($postType == "rand") { echo 'class="cp"';} ?>>随选</a> 
     </div> 
    </div> 
    <!--
    <div class="c c_6"> 
     <div class="ql qr rtr"> 
      <p>第 <?php echo $page; ?> 页</p> 
      <p>总 <?php echo $pages; ?> 页</p> 
     </div> 
    </div> 
    -->
   </div> 
  </div>
  <!-- post records -->
  <div class="r">
   <div class="w tb_15 pagee"> 
   <?php 
   $lists = mysqli_query($conn,$querySql);
   ?>
    <?php while($list = mysqli_fetch_assoc($lists)){ ?>	
    <a class="c c_4" href="/post/<?php echo $list['url']; ?>.html">
     <div class="pime">
      <img class="scale" src="<?php echo $cdnUrl."/static/upload/banner/400_".$list['banner']; ?> " />
     </div> 
     <div class="pinfb"> 
      <h2><?php echo $list['title']; ?></h2>
      <p class="pdesc"><?php echo $list['remark']; ?></p>  
     </div> </a> 
    <?php } ?>
    <?php
     $pathDir="/posts";
     $pathBe="/".$postType;
     include(__DIR__.'/makePageHtml.php');
    ?>
   </div>
  </div>
<?php
  include(__DIR__.'/makeFoot.php');
	//页面绘制完成
	//存储新的资源数据
	$content = ob_get_contents();
	//写入文件
	file_put_contents($file_sc,$content);
	//擦除缓冲区并关闭
	ob_end_clean();
}
?>