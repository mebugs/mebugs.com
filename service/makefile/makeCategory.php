<?php
// 制作分类总页
// 制作分类列表页
function makeCategoryAndList($fteams,$ftags,$cdnUrl,$baseUrl,$conn)
{
  // 分类目前默认只有一页
  $backUrls = [];
  $query = "SELECT * FROM `category` ORDER BY sorts";
  $teams = mysqli_fetch_all(mysqli_query($conn,$query),MYSQLI_ASSOC);
  // 制作分类总页
  makeCategory($fteams,$ftags,$teams,$cdnUrl,$baseUrl,$conn);
  array_push($backUrls,["loc" => $baseUrl."/category/list_1.html", "priority" => 0.8]);
  // 逐一查询分类分页
  foreach ($teams as $team) {
    $count = $team["num"];
    $size = 15;
    $pages = ceil($count/$size);
    // 查询最新文章
    for($n = 1;$n <= $pages;$n ++) {
      // 查询当页数据
      $querySql = "SELECT * FROM `post_main` m WHERE m.`status` = 1 AND m.`cid` = ".$team["id"]." ORDER BY id DESC LIMIT ".($n-1)*$size.",".$size;
      makeCategoryList($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$querySql,$team,$n,$pages);
      array_push($backUrls,["loc" => $baseUrl."/category/".$team["url"]."_".$n.".html", "priority" => 0.8]);
    }
  }
  return $backUrls;
}

// 制作分类分页
function makeCategoryList($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$querySql,$team,$page,$pages) {
  $file_sc = $_SERVER['DOCUMENT_ROOT']."/category/".$team["url"]."_".$page.".html";
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  $title = "[".$team["name"]."]分类 - 第".$page."页"." - 米虫 - 做一个有理想的米虫 - www.mebugs.com";
  $description = $team["name"]."分类 - 程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！";
  $pageType = "postList";//文章列表页
  $pageSubType = "category";//分类
  include(__DIR__.'/makeHead.php');
?>
  <!-- team top -->
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_12"> 
     <div class="ql qtag pt pct rtl">
       <div class="timg">
         <img onload="initSvg(this)" src="<?php echo $team["icon"]; ?>" /> 
         <h1><?php echo $team["name"]; ?></h1>
       </div>    
       <div class="tmin">
         <p><?php echo $team["remark"]; ?></p> 
       </div>
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
     $pathDir="/category";
     $pathBe="/".$team["url"];
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

// 制作分类总页
function makeCategory($fteams,$ftags,$teams,$cdnUrl,$baseUrl,$conn) {
  $file_sc = $_SERVER['DOCUMENT_ROOT']."/category/list_1.html";
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  $title = "文章分类 - 米虫 - 做一个有理想的米虫 - www.mebugs.com";
  $description = "文章分类 - 程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！";
  $pageType = "category";//文章分类页
  $pageSubType = "category";//文章分类页
  include(__DIR__.'/makeHead.php');
?>
  <!-- top -->
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_12"> 
     <div class="pt pct rtl"> 
      <h1>文章分类</h1> 
     </div> 
    </div> 
   </div> 
  </div> 
  <div class="r"> 
   <div class="w tb_15 teame">
    <?php foreach ($teams as $team) {?>
    <a href="/category/<?php echo $team["url"]; ?>_1.html" class="c c_3">
     <div class="teami"> 
      <img onload="initSvg(this)" src="<?php echo $team["icon"]; ?>" /> 
      <div class="tmin">
       <span><?php echo $team["num"]; ?></span> 
       <h1><?php echo $team["name"]; ?></h1> 
       <p><?php echo $team["remark"]; ?></p> 
      </div>
     </div>
    </a> 
    <?php }?>
    <?php
     $pathDir="/category";
     $pathBe="/list";
     $pages=1;
     $page=1;
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