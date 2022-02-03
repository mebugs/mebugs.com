<?php
// 查询tag分页和详情
function makeTagAndList($fteams,$ftags,$cdnUrl,$baseUrl,$conn){
  $backUrls = [];
  // 查询全部标签
  $countSql = "SELECT count(id) FROM `tag` m";
  $count = mysqli_fetch_array(mysqli_query($conn,$countSql))[0];
  $size = 25;// 测试5
  $pages = ceil($count/$size);
  // 查询标签的分页
  for($n = 1;$n <= $pages;$n ++) {
    // 查询当页数据
    $querySql = "SELECT * FROM `tag` t ORDER BY t.num DESC LIMIT ".($n-1)*$size.",".$size;
    $tags = mysqli_fetch_all(mysqli_query($conn,$querySql),MYSQLI_ASSOC);
    makeTag($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$tags,$n,$pages);
    array_push($backUrls,["loc" => $baseUrl."/tag/list_".$n.".html", "priority" => 0.8]);
    // 生成相关Tag的列表页
    
  }
  return $backUrls;
}

// 生成tag的分页数据
function makeTag($fteams,$ftags,$cdnUrl,$baseUrl,$conn,$tags,$page,$pages) {
  $file_sc = $_SERVER['DOCUMENT_ROOT']."/tag/list_".$page.".html";
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  $title = "文章标签 - 第".$page."页"." - 米虫 - 做一个有理想的米虫 - www.mebugs.com";
  $description = "文章标签 - 程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！";
  $pageType = "tag";//标签
  $pageSubType = "tag";//标签
  include(__DIR__.'/makeHead.php');
?>
  <!-- tag top -->
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_12"> 
     <div class="pt pct rtl"> 
      <h1>文章标签</h1> 
     </div> 
    </div> 
   </div> 
  </div> 
  <div class="r"> 
   <div class="w tb_15 teame tage"> 
    <?php foreach ($tags as $tag) {?>
    <a href="/category/<?php echo $tag["url"]; ?>_1.html" class="c c_20">
     <div class="tmin">
      <span><?php echo $tag["num"]; ?></span> 
      <h1><?php echo $tag["name"]; ?></h1> 
      <p><?php echo $tag["remark"]; ?></p> 
     </div>
    </a> 
    <?php }?>
    <?php
     $pathDir="/tag";
     $pathBe="/list";
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

// 生成tag的全部数据

?>