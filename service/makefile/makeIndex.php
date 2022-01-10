<?php 
//生成INDEX页面代码 CDN切换 数据库对象
function makeIndex($smliImgs,$cdnUrl,$baseUrl,$conns,$conn) {
	// 启动缓存输出 采用模版形式无需缓冲区
	ob_start();
  // 目标文件
  $file_sc = $_SERVER['DOCUMENT_ROOT']."/index.html";
  // siteMap生成对象
  $backUrl = ["loc" => $baseUrl, "priority" => 1];
  $title = "米虫 - 做一个有理想的米虫，全栈程序员，乐观主义者，坚信一切都是最好的安排！ - www.mebugs.com";
  $description = "程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！";
  $pageType = "index";//页面类型首页
  include(__DIR__.'/makeHead.php'); 
?>
  <!-- index top -->
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_8 c_bn"> 
     <div class="swiper-container bann rtl"> 
      <div class="swiper-wrapper"> 
      <?php
        // 首页Banner配置
        $bnerSql = "SELECT * FROM `urls` WHERE `status` = 1 AND `type` = 'banner' ORDER BY `sorts`";
        $bners = mysqli_query($conn,$bnerSql);
        while($bner = mysqli_fetch_assoc($bners)){
          echo '<div class="swiper-slide"<a href=">'.$bner['url'].'"><img src="'.$cdnUrl.$bner['img'].'" /></a></div>';
        }
      ?> 
      </div> 
      <!-- Add Pagination --> 
      <div class="swiper-pagination"></div> 
     </div> 
    </div> 
    <div class="c c_4"> 
     <div class="mine rtr"> 
      <a href="/page/about.html"> 
       <div class="meavt"> 
        <img class="avtor" src="/static/img/me_avator.jpg" /> 
        <div class="ainfo"> 
         <h1>米虫</h1> 
         <p>做一个有理想的米虫，全栈程序员，乐观主义者，坚信一切都是最好的安排！</p> 
        </div> 
       </div> </a> 
      <div class="mecon"> 
       <a href="<?php echo $conns['qqUrl']; ?>" class="conn cqq"></a> 
       <a href="<?php echo $conns['emailUrl']; ?>" class="conn cml"></a> 
       <a href="<?php echo $conns['giteeUrl']; ?>" class="conn cge"></a> 
       <a href="<?php echo $conns['githubUrl']; ?>" class="conn cgb"></a> 
      </div> 
      <div class="mequc"> 
      <?php foreach ($smliImgs as $smliImg) {
        echo '<div class="mequci"><a href="'.$smliImg['url'].'"><img class="scale" src="'.$cdnUrl.$smliImg['img'].'" /></a></div> ';
      } ?>
      </div> 
     </div> 
    </div> 
   </div> 
  </div> 
  <!-- index main -->
  <div class="r"> 
   <div class="w tb_15"> 
    <!-- index main left -->
    <div class="c c_8"> 
     <!-- index goods -->
     <div class="row"> 
      <div class="pt"> 
       <h1>好文荐读</h1> 
       <a class="mr" href="/posts/best_1.html">More&gt;</a> 
      </div> 
<?php 
// 查询深度好文数据
$goodSql = "SELECT * FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY depth DESC LIMIT 10";
$goods = mysqli_query($conn,$goodSql);
$gi = 1;
?>
      <div class="box pm pa_6"> 
      <?php while($good = mysqli_fetch_assoc($goods)){ ?>	
      <a <?php if($gi < 3) { ?> class="fulla" <?php } ?> href="/post/<?php echo $good['url']; ?>.html"> 
        <?php if($gi >= 3) { ?> <div class="pimel"> <?php } ?>
        <img class="scale" src="<?php 
          if($gi < 3) {
            echo $cdnUrl."/static/upload/banner/400_".$good['banner']; 
          } else {
            echo $cdnUrl."/static/upload/banner/100_".$good['banner']; 
          } ?>" /> 
        <?php if($gi >= 3) { ?> </div> <?php } ?>
        <div class="<?php if($gi < 3) { echo "pbline"; } else { echo "pinfr"; } ?>"> 
         <h2><?php echo $good['title']; ?></h2> 
         <p class="pdesc"><?php echo $good['remark']; ?></p> 
        </div> 
      </a> 
      <?php  $gi++; } ?>
      </div> 
     </div> 
     <!-- index news -->
     <div class="row"> 
      <div class="pt"> 
       <h1>新鲜发布</h1> 
       <a class="mr" href="/posts/new_1.html">More&gt;</a> 
      </div> 
<?php 
// 查询最新数据
$newSql = "SELECT * FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY id DESC LIMIT 8";
$news = mysqli_query($conn,$newSql);
$ni = 1;
?>
      <div class="box pm pa_6"> 
      <?php while($new = mysqli_fetch_assoc($news)){ ?>
      <a <?php if($ni < 3) { ?> class="fulla" <?php } ?> href="/post/<?php echo $new['url']; ?>.html"> 
        <?php if($ni >= 3) { ?> <div class="pimel"> <?php } ?>
        <img class="scale" src="<?php 
          if($ni < 3) {
            echo $cdnUrl."/static/upload/banner/400_".$new['banner']; 
          } else {
            echo $cdnUrl."/static/upload/banner/100_".$new['banner']; 
          } ?>" /> 
        <?php if($ni >= 3) { ?> </div> <?php } ?>
        <div class="<?php if($ni < 3) { echo "pbline"; } else { echo "pinfr"; } ?>"> 
         <h2><?php echo $new['title']; ?></h2> 
         <p class="pdesc"><?php echo $new['remark']; ?></p> 
        </div> 
      </a> 
      <?php  $ni++; } ?>
      </div> 
     </div> 
     <!-- index ups -->
     <div class="row"> 
      <div class="pt"> 
       <!-- 30Day --> 
       <h1>近期热门</h1> 
       <a class="mr" href="/posts/up_1.html">More&gt;</a> 
      </div> 
<?php 
// 查询最新数据
$upSql = "SELECT * FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY monthView DESC LIMIT 10";
$ups = mysqli_query($conn,$upSql);
$ui = 1;
?>
      <div class="box pa_6 pm"> 
      <?php while($up = mysqli_fetch_assoc($ups)){ ?>
      <a <?php if($ui < 3) { ?> class="fulla" <?php } ?> href="/post/<?php echo $up['url']; ?>.html"> 
        <?php if($ui >= 3) { ?> <div class="pimel"> <?php } ?>
        <img class="scale" src="<?php 
          if($ui < 3) {
            echo $cdnUrl."/static/upload/banner/400_".$up['banner']; 
          } else {
            echo $cdnUrl."/static/upload/banner/100_".$up['banner']; 
          } ?>" /> 
        <?php if($ui >= 3) { ?> </div> <?php } ?>
        <div class="<?php if($ui < 3) { echo "pbline"; } else { echo "pinfr"; } ?>"> 
         <h2><?php echo $up['title']; ?></h2> 
         <p class="pdesc"><?php echo $up['remark']; ?></p> 
        </div> 
      </a> 
      <?php  $ui++; } ?>
      </div> 
     </div> 
    </div> 
    <!-- index main right -->
    <div class="c c_4"> 
     <!-- index main hot -->
     <div class="row"> 
      <div class="pt"> 
       <h1>全站热门</h1> 
       <a class="mr" href="/posts/hot_1.html">More&gt;</a> 
      </div> 
<?php 
// 查询热门数据
$hotSql = "SELECT * FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY view DESC LIMIT 10";
$hots = mysqli_query($conn,$hotSql);
$hi = 1;
?>
      <div class="box lm"> 
      <?php while($hot = mysqli_fetch_assoc($hots)){ ?>
      <a href="/post/<?php echo $hot['url']; ?>.html">  
        <div class="pimel ltop"><?php echo $hi; ?></div> 
        <div class="pinfr"> 
         <h2><?php echo $hot['title']; ?></h2> 
        </div> 
      </a>
      <?php  $hi++; } ?>
      </div> 
     </div> 
     <!-- index team -->
     <div class="row"> 
      <div class="pt"> 
       <h1>文章分类</h1> 
       <a class="mr" href="/category">More&gt;</a> 
      </div> 
<?php 
// 查询分类
$teamSql = "SELECT * FROM `category` ORDER BY sorts";
$teams = mysqli_query($conn,$teamSql);
?>
      <div class="box itm"> 
      <?php while($team = mysqli_fetch_assoc($teams)){ ?>
      <a href="/category/<?php echo $team['url']; ?>_1.html"> 
        <img src="<?php echo $cdnUrl.$team['icon']; ?>" onload="initSvg(this)" /> 
        <h2><?php echo $team['name']; ?></h2>
      </a>
      <?php } ?>
      </div> 
     </div> 
     <div class="row"> 
      <div class="pt"> 
       <h1>热门标签</h1> 
       <a class="mr" href="/tags">More&gt;</a> 
      </div> 
<?php 
// 查询分类
$tagSql = "SELECT * FROM `tag` ORDER BY num DESC LIMIT 60";
$tags = mysqli_query($conn,$tagSql);
?>
      <div class="box iag"> 
        <div class="ags">
        <?php while($tag = mysqli_fetch_assoc($tags)){ ?>
        <a href="/tag/<?php echo $tag['url']; ?>_1.html"><?php echo $tag['name']; ?></a> 
        <?php } ?>
       </div> 
      </div> 
     </div> 
    </div> 
    <div class="c c_12"> 
     <div class="row"> 
      <div class="pt"> 
       <h1>随便看看</h1> 
       <a class="mr" href="/posts/rand_1.html">More&gt;</a> 
      </div> 
<?php 
// 查询随机数据
$randSql = "SELECT * FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY randno DESC LIMIT 10";
$rands = mysqli_query($conn,$randSql);
$ri = 1;
?>
      <div class="box pa_4 pm">
      <?php while($rand = mysqli_fetch_assoc($rands)){ ?>
      <a <?php if($ri < 4) { ?> class="fulla" <?php } ?> href="/post/<?php echo $rand['url']; ?>.html"> 
        <?php if($ri >= 4) { ?> <div class="pimel"> <?php } ?>
        <img class="scale" src="<?php 
          if($ri < 4) {
            echo $cdnUrl."/static/upload/banner/400_".$rand['banner']; 
          } else {
            echo $cdnUrl."/static/upload/banner/100_".$rand['banner']; 
          } ?>" /> 
        <?php if($ri >= 4) { ?> </div> <?php } ?>
        <div class="<?php if($ri < 4) { echo "pbline"; } else { echo "pinfr"; } ?>"> 
         <h2><?php echo $rand['title']; ?></h2> 
         <p class="pdesc"><?php echo $rand['remark']; ?></p> 
        </div> 
      </a> 
      <?php  $ri++; } ?>
      </div> 
     </div> 
     <div class="row"> 
      <div class="pt pct"> 
       <h1>友情链接</h1> 
      </div> 
<?php 
// 查询首页友链
$friendSql = "SELECT * FROM `friend` WHERE `status` = 1 AND `findex` = 1";
$friends = mysqli_query($conn,$friendSql);
?>
      <div class="box lk"> 
      <?php while($friend = mysqli_fetch_assoc($friends)){ ?>
      <a href="<?php echo $friend['furl']; ?>" target="_blank"><?php echo $friend['fname']; ?></a> 
      <?php } ?>
      <div class="lmr"> 
        <a href="/page/link.html">更多友链</a> 
        <a href="/page/makeFriend.html">友链申请</a> 
      </div> 
      </div> 
     </div> 
    </div> 
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
  return $backUrl;
} 
?>