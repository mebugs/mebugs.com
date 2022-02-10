<?php

function makeFriends($fteams,$ftags,$cdnUrl,$baseUrl,$conn) {
  // 目标文件
  $file_sc = $_SERVER['DOCUMENT_ROOT']."/page/link.html";
  // siteMap生成对象
  $backUrls = [];
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  $title = "友链 - 做一个有理想的米虫，全栈程序员，乐观主义者，坚信一切都是最好的安排！ - www.mebugs.com";
  $description = "程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！";
  $pageType = "link";//友链
  $pageSubType = "link";
  include(__DIR__.'/makeHead.php'); 

?>
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_12 pst"> 
     <div class="pt pct rtl"> 
      <h1>友情链接丨米虫的邻居们丨来呀，交个朋友！</h1> 
     </div> 
     <div class="pdest rtr">
       这里是与米虫同住地球村的小伙伴们，友情互链！排名不分先后，暂未对站点进行详细分类！同时欢迎新的小伙伴&amp;大佬们与米虫交个朋友~.~
      <a class="bgl" href="/page/makeFriend.html">+申请友链+</a> 
     </div> 
    </div> 
   </div> 
  </div> 
  <?php
    $sql = "SELECT * FROM `friend` WHERE status = 1 ORDER BY findex,id";
    $query = mysqli_query($conn,$sql);
  ?>
  <div class="r">
   <div class="w tb_15 teame linke"> 
    <?php while($friend = mysqli_fetch_assoc($query)){ ?>
    <a href="<?php echo $friend['furl']; ?>" class="c c_4">
     <div class="teami"> 
      <img src="<?php echo $friend['ficon']; ?>" /> 
      <div class="tmin"> 
       <h1><?php echo $friend['fname']; ?></h1> 
       <p><?php echo $friend['fdesc']; ?></p> 
      </div>
     </div> 
    </a> 
    <?php } ?>
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
  array_push($backUrls,["loc" => $baseUrl."/page/link.html", "priority" => 0.8]);
  array_push($backUrls,["loc" => $baseUrl."/page/makeFriend.html", "priority" => 0.8]);
  return $backUrls;
}

?>