<?php 
//生成INDEX页面代码 CDN切换 数据库对象
function makePosts($smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn,$doPosts) {
  // 查询全部启用文章以id正序  < 999 写入page目录
  $backUrls = [];
  // 两个不同目录
  $pageBase = "/page/";
  $postBase = "/post/";
  // 查询全部开放文章
  $postSql = "SELECT * FROM `post_main` m WHERE m.`status` > 0";
  $posts = mysqli_query($conn,$postSql);
  while($post = mysqli_fetch_assoc($posts)){
    $postPageUrl = $postBase.$post['url'].".html";
    if($post['id'] < 999) {
      $postPageUrl = $pageBase.$post['url'].".html";
    }
    makePost($post,$postPageUrl,$smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn,$doPosts);
    array_push($backUrls,["loc" => $baseUrl.$postPageUrl, "priority" => 1]);
  }
  return $backUrls;
} 

function makePost($post,$postPageUrl,$smliImgs,$fteams,$ftags,$cdnUrl,$baseUrl,$conns,$conn,$doPosts) {
  // 目标文件
  $file_sc = $_SERVER['DOCUMENT_ROOT'].$postPageUrl;
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
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
     <div class="menue" id="menue"><?php echo $postInfo['menu']; ?></div> 
    </div> 
    <div class="c c_9">
     <!-- post banner info -->
     <div class="mban"> 
      <img src="<?php echo $cdnUrl."/static/upload/banner/".$post['banner']; ?>" /> 
      <?php 
      if($post['id'] > 999) {
        // 读取分类名
        $postTeamSql = "SELECT * FROM `category` WHERE id =".$post['cid'];
        $postTeam = mysqli_fetch_assoc(mysqli_query($conn,$postTeamSql));
        // 读取tage集
        $postTagsSql = "SELECT t.* FROM `post_tag` pt LEFT JOIN `tag` t ON t.id = pt.tid WHERE pt.pid = ".$post['id'];
        $postTags = mysqli_query($conn,$postTagsSql)
      ?>
      <div class="pfro">
        <p>
          <span>所属分类</span>
          <b class="pta"><?php echo '<a href="/category/'.$postTeam['url'].'_1.html">'.$postTeam['name'].'</a>';  ?> </b>
        </p>
        <p>
          <span>相关标签</span>
          <b class="pta"><?php while($postTag = mysqli_fetch_assoc($postTags)){ echo '<a href="/tags/'.$postTag['url'].'_1.html">'.$postTag['url'].'</a>'; } ?></b>
        </p>
      </div>
      <?php }?>
     </div>
     <!-- post content -->
     <div id="html" class="output mark"><?php echo $postInfo['html']; ?></div>
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
    </div>
    <div class="c c_12" id="mother"> 
     <?php if($post["openComms"] == 0) { ?>
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
     <!-- comms eare -->
     <div class="row">
      <div class="commi">
      <?php
      $commsCount = mysqli_fetch_array(mysqli_query($conn,"SELECT count(id) FROM `comms` c WHERE c.`status` > 0 AND c.pid = ".$post["id"]))[0];
      if($commsCount == 0) {
        echo '<p>当前还没有观点发布，欢迎您留下足迹！</p>';
      } else {
        echo '<p>当前累计<b>'.$commsCount.'</b>条观点！一起看看吧！</p>';
        echo '<div class="ckOrder"><span class="cko" onclick="toCommsOrder(true)">最新</span><span onclick="toCommsOrder(false)">最早</span></div>';
      }
      // 查询pageCount
      $pageCount = mysqli_fetch_array(mysqli_query($conn,"SELECT count(id) FROM `comms` c WHERE c.`status` > 0 AND c.`level` = 1 AND c.pid = ".$post["id"]))[0];
      ?>
      </div>
      <?php if($commsCount > 0) { ?>
      <div class="box">
       <div class="comml">
        <ul id="commea">
        <?php
        // 以及评论 查询倒序
        $commNews = mysqli_query($conn,"SELECT * FROM `comms` c WHERE c.`status` > 0 AND c.`level` = 1 AND c.pid = ".$post["id"]." ORDER BY id DESC"); ?>
         <?php while($commNew = mysqli_fetch_assoc($commNews)){ ?>
         <li>
          <div class="comli"> 
           <img src="<?php echo $commNew["avt"]; ?>"/> 
           <div class="comp"> 
            <h1>
             <b><a target="_blank" href="<?php echo $commNew["url"]; ?>"><?php echo $commNew["name"]; ?></a></b>
             <?php echo "丨".$commNew["send_time"]." 发表观点丨"; ?>
             <a href="javascript:sendCommsR(<?php echo $post['id']; ?>,2,<?php echo $commNew["id"]; ?>)">评论TA</a>
            </h1> 
            <p><?php echo $commNew["coms"]; ?></p> 
           </div> 
          </div>
          <?php
          // 二级评论数量
          $lvTwoCounts = mysqli_fetch_array(mysqli_query($conn,"SELECT count(id) FROM `comms` c WHERE c.`status` > 0 AND c.`level` > 1 AND c.fid = ".$commNew["id"]))[0];
          ?>
          <?php if($lvTwoCounts > 0) {?>
          <!-- 二三级评论 -->
          <ul>
            <?php
            // 查询二三级评论列表
            //$lvTTSql = "SELECT * FROM (SELECT ts.`name` AS pName,ts.url AS pUrl,tc.coms,tc.id,tc.name,tc.url,tc.avt,tc.fid,tc.send_time FROM `comms` tc LEFT JOIN `comms` ts ON ts.id = tc.fid WHERE tc.`status` > 0 AND tc.fid = ".$commNew["id"]." UNION ALL SELECT cs.`name` AS pName,cs.url AS pUrl,c.coms,c.id,c.name,c.url,c.avt,c.fid,c.send_time FROM `comms` c LEFT JOIN `comms` cs ON cs.id = c.fid WHERE c.fid IN (SELECT id FROM `comms` WHERE `status` = 1 AND fid = ".$commNew["id"].") AND c.`status` > 0) tt ORDER BY tt.id";
            $lvTTSql = "SELECT ts.`name` AS pName,ts.url AS pUrl,tc.coms,tc.id,tc.name,tc.url,tc.avt,tc.fid,tc.send_time FROM `comms` tc LEFT JOIN `comms` ts ON ts.id = tc.fid WHERE tc.`status` > 0 AND tc.gid = ".$commNew["id"]." ORDER BY tc.id";
            $lvTTs = mysqli_query($conn,$lvTTSql);
            ?>
            <?php while($lvTT = mysqli_fetch_assoc($lvTTs)){ ?>
            <li>
             <div class="comli"> 
              <div class="comp">
               <h1>
                 <b><a target="_blank" href="<?php echo $lvTT["url"]; ?>"><?php echo $lvTT["name"]; ?></a></b>
                 <?php 
                 echo "丨".$lvTT["send_time"];
                 if($lvTT["fid"] == $commNew["id"]) { 
                   echo ' 评论本楼丨';
                 }else{
                   echo ' 回复<b><a target="_blank" href="'.$lvTT["pUrl"].'">'.$lvTT["pName"].'</a></b>丨';
                 }
                 ?>
                 <a href="javascript:sendCommsR(<?php echo $post['id']; ?>,3,<?php echo $lvTT["id"]; ?>)">回复TA</a></h1> 
               </h1> 
               <p><?php echo $lvTT["coms"]; ?></p> 
              </div> 
             </div>
            </li>
            <?php }?>
          </ul>
          <?php }?>
         </li>
         <?php }?>
        </ul>
       </div>
       <div class="comore">
        <?php 
        if($pageCount > 5) { 
          echo '<a id="comodoo" href="javascript:moreComs()">更多观点</a><p id="comonul" style="display:none;">已经到底啦~.~</p>';
        }else{
          echo '<a id="comodoo" style="display:none;" href="javascript:moreComs()">更多观点</a><p id="comonul">已经到底啦~.~</p>';
        } 
        ?>
        <p id="comolod" style="display:none;">加载中...</p>
       </div>
      </div>
      <?php }?>
     </div> 
     <?php }?>
     <!-- sml imgs -->
     <div class="row">
      <div class="box"> 
       <div class="mequc indimg">
        <?php foreach ($smliImgs as $smliImg) {
          echo '<div class="mequci"><a href="'.$smliImg['url'].'"><img class="scale" src="'.$cdnUrl.$smliImg['img'].'" /></a></div> ';
        } ?>
       </div> 
      </div> 
     </div>
     <?php if($post['id'] > 999) { ?>
     <!-- like post -->
     
     <?php }?>
    </div>
    <!-- more post new hot rand -->
    <div class="c c_4">
     <!-- new -->
     <div class="row"> 
      <div class="pt"><h1>新鲜发布</h1><a class="mr" href="/posts/new_1.html">More&gt;</a></div> 
      <?php $npi = 1; ?>
      <div class="box lm"> 
       <?php foreach ($doPosts['new'] as $newP) { ?>
       <a href="/post/<?php echo $newP['url']; ?>.html">  
        <div class="pimel ltop"><?php echo $npi; ?></div> 
        <div class="pinfr"><h2><?php echo $newP['title']; ?></h2></div> 
       </a>
       <?php  $npi++; } ?>
      </div> 
     </div>  
    </div>
    <div class="c c_4">
      <!-- new -->
      <div class="row"> 
       <div class="pt"><h1>近期热门</h1><a class="mr" href="/posts/up_1.html">More&gt;</a></div> 
       <?php $upi = 1; ?>
       <div class="box lm"> 
        <?php foreach ($doPosts['hot'] as $upP) { ?>
        <a href="/post/<?php echo $upP['url']; ?>.html">  
         <div class="pimel ltop"><?php echo $upi; ?></div> 
         <div class="pinfr"><h2><?php echo $upP['title']; ?></h2></div> 
        </a>
        <?php  $upi++; } ?>
       </div> 
      </div>  
    </div>
    <div class="c c_4">
      <!-- new -->
      <div class="row"> 
       <div class="pt"><h1>随便看看</h1><a class="mr" href="/posts/rand_1.html">More&gt;</a></div> 
       <?php $rpi = 1; ?>
       <div class="box lm"> 
        <?php foreach ($doPosts['new'] as $radP) { ?>
        <a href="/post/<?php echo $radP['url']; ?>.html">  
         <div class="pimel ltop"><?php echo $rpi; ?></div> 
         <div class="pinfr"><h2><?php echo $radP['title']; ?></h2></div> 
        </a>
        <?php  $rpi++; } ?>
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
}?>