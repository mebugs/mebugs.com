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
  return $backUrls;
}

function makeSendFriends($fteams,$ftags,$cdnUrl,$baseUrl,$conn){
  // 目标文件
  $file_sc = $_SERVER['DOCUMENT_ROOT']."/page/makeFriend.html";
  // siteMap生成对象
  $backUrls = [];
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  $title = "申请友链 - 做一个有理想的米虫，全栈程序员，乐观主义者，坚信一切都是最好的安排！ - www.mebugs.com";
  $description = "程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！";
  $pageType = "link";//友链
  $pageSubType = "sendLink";
  include(__DIR__.'/makeHead.php'); 
?>  
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_12 pst"> 
     <div class="pt pct rtl"> 
      <h1>友链申请</h1> 
     </div> 
     <div class="pdest srm rtr">
      欢迎新伙伴来到本页面，请您依照下方表单提示提交友链信息。<br>
      当您在本页面提交申请后，1-2个工作日内完成处理，请优先添加本站为友情链接。<br>
      申请首页友链需要贵站在首页位置添加本站为友情链接。<br>
      原则上对申请站点无太多要求，原创类、独立博客类、资源类、教程类站点优先。<br>
      原则上不关注PR或搜索引擎收录，贵站需要至少保持月更，如您长期无更新（三个月）可能会撤销链接。<br>
      本站可能不定期访问友站，如您的站点已关闭或停止运营，本站撤链将不另行通知。<br>
      本站拒绝以下类型站点申请：采集站、聚合站、涉嫌诈骗或非法类信息站、伪原创，拒绝理由将回复在内容下。
     </div> 
    </div> 
   </div> 
  </div> 
  <div class="r"> 
   <div class="w tb_15"> 
    <div class="c c_12 mkf">
      <div class="box fform">
        <div class="fline">
          <select name="ffix" id="ffix">
            <option value ="http://" selected="selected">http://</option>
            <option value ="https://">https://</option>
          </select>
          <input name="furl" id="furl" placeholder="*必填,请在左侧选择访问协议">
        </div>
        <div class="fline">
          <p><a href="javascript:getInfo()">快速采集</a>输入地址后点击【快速采集】自动获取站点信息，可修改。</p>
        </div>
        <div class="fline">
          <span>站点名称</span><input name="fname" id="fname" type="text" placeholder="*必填,最高16字(超出提交会被截取)" maxlength="16"/>
        </div>
        <div class="fline">
          <span class="desc">站点描述</span>
          <textarea name="fdesc" id="fdesc" type="text" placeholder="*必填,最高128字(超出提交会被截取)" maxlength="128"></textarea>
        </div>
        <div class="fline">
          <span>友链类型</span>
          <select class="rsel" name="findex" id="findex">
            <option value ="0" selected="selected">仅申请友链</option>
            <option value ="1">同时申请首页友链</option>
          </select>
        </div>
        <div class="fline">
          <span class="ico">站点图标</span>
          <div class="fimg" onclick="chooseImg('fuimg')">
            <img id="ficon" src="/static/img/icon.png" onerror="javascript:this.src='/static/img/icon.png'" /> 
            <p>自行上传</p> 
          </div>
          <a class="rsend" href="javascript:sendInfo()">提交申请</a>
          <input name="fuimg" id="fuimg"  accept="image/gif, image/jpeg, image/png, image/jpg" type="file" style="display: none;" onchange="loadImg(this,'ficon')">
        </div>
      </div>
    </div>
    <div class="w sfr">
      <div class="pt pct rtl c_12">
       <h1>最新记录</h1> 
      </div>
      <!-- 处理记录 -->
      <?php
        $sql = "SELECT * FROM `friend` ORDER BY id DESC LIMIT 8";
        $query = mysqli_query($conn,$sql);
      ?>
      
      <?php while($friend = mysqli_fetch_assoc($query)){ ?>
      <div class="c c_3">
        <div class="box sline">
          <h2><?php echo $friend['fname']; ?></h2>
          <p>
            <span>
              <?php if($friend['status']==0) { echo "待审核";}else if($friend['status']==1){echo "通过";}else{echo "拒绝";} ?>
            </span>
            <?php echo $friend['reson']; ?>
          </p>
        </div>
      </div>
      <?php } ?>
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
  array_push($backUrls,["loc" => $baseUrl."/page/makeFriend.html", "priority" => 0.8]);
  return $backUrls;
}


function makeMap($fteams,$ftags,$cdnUrl,$baseUrl,$conn){
  // 目标文件
  $file_sc = $_SERVER['DOCUMENT_ROOT']."/page/map.html";
  // siteMap生成对象
  $backUrls = [];
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  $title = "文章地图 - 做一个有理想的米虫，全栈程序员，乐观主义者，坚信一切都是最好的安排！ - www.mebugs.com";
  $description = "程序员米虫的个人修炼手册，做一个有理想的米虫，乐观主义者，坚信一切都是最好的安排！技术学习笔记、经验教训小结、工作心得体会、自我充电提升！希望您也能有所收获！";
  $pageType = "map";//友链
  $pageSubType = "map";
  include(__DIR__.'/makeHead.php'); 
?>  
  <div class="r rt"> 
   <div class="w tb_15"> 
    <div class="c c_12 pst"> 
     <div class="pt pct rtl"> 
      <h1>全站文章地图</h1> 
     </div> 
     <div class="pdest rtr">
      文章地图为您提供全站文章的超链接，本页面提供了简易的便捷模糊搜索能力（仅针对标题）
     </div> 
    </div> 
   </div> 
  </div> 
  <div class="r"> 
   <div class="w tb_15"> 
   <div class="c c_12">
    <div class="row">
      <div class="box">
        <div class="qsc">
          <input id="pname" type="text" placeholder="请输入您想搜索的文章名" />
          <a href="javascript:search()">快速搜索</a> 
        </div>
      </div>
    </div>
    <div class="row">
      <div class="box">
       <p class="scback">为您搜索到 [<b id="scPname"></b>] 相关文章 [<b id="scNo"></b>] 篇。<a href="javascript:refush()">重置</a></p>   
       <div class="pm pa_4 plist" id="plist">
        <?php
          $newSql = "SELECT url,title FROM `post_main` WHERE id > 999 AND `status` = 1 ORDER BY id DESC";
          $news = mysqli_query($conn,$newSql);
        ?>
        <?php while($new = mysqli_fetch_assoc($news)){ ?>
        <a href="/post/<?php echo $new['url']; ?>.html"><h2><?php echo $new['title']; ?></h2></a>
        <?php } ?>
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
  array_push($backUrls,["loc" => $baseUrl."/page/map.html", "priority" => 0.8]);
  return $backUrls;
}
?>