<?php 
//生成INDEX页面代码 CDN切换 数据库对象
function makePosts($smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn) {
  // 查询全部启用文章以id正序  < 999 写入page目录
  $backUrls = [];
  // 两个不同目录
  $root = $_SERVER['DOCUMENT_ROOT'];
  $pageBase = "/page/";
  $postBase = "/post/";
  // 查询全部开放文章
  $postSql = "SELECT * FROM `post_main` m WHERE m.`status` > 0";
  $posts = mysqli_query($conn,$postSql);
  while($post = mysqli_fetch_assoc($bners)){
    $postPageUrl = $postBase.$post['url'].".html";
    if($post['id'] < 999) {
      $postPageUrl = $pageBase.$post['url'].".html";
    }
    makePost($post,$postPageUrl,$smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn)
    array_push($backUrls,["loc" => $baseUrl.$postPageUrl, "priority" => 1])
  }
  
} 

function makePost($post,$postPageUrl,$smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn) {
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  // 目标文件
  $file_sc = $_SERVER['DOCUMENT_ROOT'].$postPageUrl;
  $title = $post['title']." - 米虫 - 做一个有理想的米虫 - www.mebugs.com";
  $description = $post['remark']." - 米虫 - 做一个有理想的米虫 - www.mebugs.com";
  $pageType = "post";//页面类型文章页
  $pageSubType = "post";
  if($post['id'] < 1000) {
     $pageSubType = $post['url']; // 固定的定制URL页 如 about msg
  }
  include(__DIR__.'/makeHead.php');
?>
  <!-- POST TOP -->
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_12 pst"> 
     <div class="pt pct rtl"> 
      <h1><?php echo $post['title']; ?></h1> 
     </div> 
     <div class="pdest rtr"><?php echo $post['remark']; ?></div> 
    </div> 
   </div> 
  </div> 
<?php 
$postInfoSql = "SELECT * FROM `post_info` WHERE pid = ".$post['id'];
$postInfo = mysqli_fetch_assoc(mysqli_query($conn,$postInfoSql));
?>
  <div class="r"> 
   <div class="w tb_15"> 
    <div class="c c_3 prel"> 
     <div class="menue" id="menue"><?php echo $cdnUrl."/static/upload/banner/".$post['banner']; ?></div> 
    </div> 
    <div class="c c_9">
     <!-- post banner info -->
     <div class="mban"> 
      <img src="<?php echo $postInfo['menu']; ?>" /> 
      <?php if($post['id'] > 999) { ?>
      <div class="pfro">
        <p>
          <span>所属分类</span>
          <b class="pta">
            <a href="/category/java_1.html">Java</a>
          </b>
        </p>
        <p>
          <span>相关标签</span>
          <b class="pta">
            <a href="/tags/springboot_1.html">SpringBoot</a>
            <a href="/tags/springboot_1.html">配置文件</a>
            <a href="/tags/springboot_1.html">优先级</a>
            <a href="/tags/springboot_1.html">环境</a>
            <a href="/tags/springboot_1.html">配置项</a>
          </b>
        </p>
      </div>
      <?php }?>
     </div>
     <!-- post content -->
     <div id="html" class="output mark"><?php echo $postInfo['html']; ?></div>
    </div>
    <div class="c c_9" id="mother"> 
     <!-- author -->
     <div class="row">
      <div class="authors">
       <div class="mea"> 
        <img class="meavt" src="/static/img/me_avator.jpg" /> 
        <div class="meainfo">
            <h1>米虫</h1>
            <p>做一个有理想的米虫，伪全栈程序猿，乐观主义者，坚信一切都是最好的安排！</p> 
            <div class="meconn">
             <a href="<?php echo $conns['qqUrl']; ?>" class="conn cqq"></a>
             <a href="<?php echo $conns['emailUrl']; ?>" class="conn cml"></a> 
             <a href="<?php echo $conns['giteeUrl']; ?>" class="conn cge"></a> 
             <a href="<?php echo $conns['githubUrl']; ?>" class="conn cgb"></a>
            </div> 
           </div> 
       </div>
       <div class="mecpr">
        <p>本站由个人原创、收集或整理，如涉及侵权请联系删除</p>
        <p>本站内容支持转发，希望贵方携带转载信息和原文链接</p>
        <p>本站具有时效性，不提供有效、可用和准确等相关保证</p>
        <p>本站不提供免费技术支持，暂不推荐您使用案例商业化</p>
       </div>
      </div>
     </div>
     <!-- comms form -->
     <div class="row">
      <div class="pt"> <h1>发表观点</h1> </div> 
      <div class="box"> 
       <div class="comme"> 
        <div class="namee"> 
         <div class="navt" onclick="chooseAvt()"><img id="avt" src="/static/upload/avtor/1.jpg" /><p>选择头像</p></div> 
         <div class="nxinfo"> 
          <div class="nipt"><p>昵称</p><input id="name" type="text" placeholder="*必填,请输入您的昵称" /></div> 
          <div class="nipt"><p>邮箱</p><input id="email" type="text" placeholder="推荐选填,请输入您的电子邮箱" /></div> 
          <div class="nipt"><p>QQ</p><input id="qq" type="text" placeholder="选填,请输入您的联系QQ" /></div> 
          <div class="nipt"><p>网址</p><input id="url" type="text" placeholder="选填,请输入您的个人主页地址" /></div> 
         </div> 
        </div> 
        <div class="ctxe"><textarea id="coms" name="coms" placeholder="*必填,请输入您的精彩观点"></textarea></div> 
        <div class="sendc"><a href="javascript:sendComms(<?php echo $post['id']; ?>,1,0)">提交观点</a></div> 
       </div>
      </div> 
     </div>
     
<?php }?>
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
  array_push($backUrls,["loc" => $baseUrl, "priority" => 1])
  return $backUrls;
} 
?>