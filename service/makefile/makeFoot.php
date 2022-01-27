  <footer> 
   <div class="w"> 
    <div class="c c_6"> 
     <div class="ft fl"> 
      <h2>排序分类</h2> 
      <p>
        <a href="/posts/best_1.html">好文荐读</a>
        <a href="/posts/hot_1.html">全站热门</a>
        <a href="/posts/new_1.html">最新发布</a>
        <a href="/posts/up_1.html">近期飙升</a>
        <a href="/posts/rand_1.html">随机排序</a>
      <?php foreach ($fteams as $fteam) {
        echo '<a href="/category/'.$fteam['url'].'_1.html">'.$fteam['name'].'</a>';
      } ?>
      </p>
     </div> 
    </div>
    <div class="c c_6"> 
     <div class="ft fr"> 
      <h2>热点标签</h2> 
      <p> 
      <?php foreach ($ftags as $tag) {
        echo '<a href="/tag/'.$tag['url'].'_1.html">'.$tag['name'].'</a>';
      } ?>
      </p>
     </div> 
    </div> 
    <div class="c c_12"> 
     <div class="fu"> 
      <p>米虫<em>[独立博客]</em> 丨 <a href="/">www.mebugs.com</a> 丨 <a href="https://beian.miit.gov.cn/" target="_blank">苏ICP备20039109号</a>丨 做一个有理想的米虫</p> 
      <p><a href="/page/about.html">关于</a> 丨 <a href="/page/msg.html">留言</a> 丨 <a href="/page/link.html">友链</a> 丨 <a href="/page/map.html">地图</a></p> 
     </div> 
    </div> 
   </div> 
  </footer> 
  <div class="back_top" onclick="backTop()"> </div> 
  <?php if($pageType == "post" && $post != null && $post["openComms"] == 0) { // 列表页POST对象预置为null?> 
  <div class="rzz" id="cavt">
    <div class="rea cave box">
      <div class="reat">
        <h1>选择个人头像</h1>
      </div>
      <div class="avts">
        <img src="/static/upload/avtor/1.jpg" onclick="setAvt(this,1)" />
        <img src="/static/upload/avtor/2.jpg" onclick="setAvt(this,2)" />
        <img src="/static/upload/avtor/3.jpg" onclick="setAvt(this,3)" />
        <img src="/static/upload/avtor/4.jpg" onclick="setAvt(this,4)" />
        <img src="/static/upload/avtor/5.jpg" onclick="setAvt(this,5)" />
        <img src="/static/upload/avtor/6.jpg" onclick="setAvt(this,6)" />
        <img src="/static/upload/avtor/7.jpg" onclick="setAvt(this,7)" />
        <img src="/static/upload/avtor/8.jpg" onclick="setAvt(this,8)" />
        <img src="/static/upload/avtor/9.jpg" onclick="setAvt(this,9)" />
        <img src="/static/upload/avtor/10.jpg" onclick="setAvt(this,10)" />
        <img src="/static/upload/avtor/11.jpg" onclick="setAvt(this,11)" />
        <img src="/static/upload/avtor/12.jpg" onclick="setAvt(this,12)" />
      </div>
      <div class="sendc"> 
        <a class="rcls" href="javascript:clsAvt()">关闭</a>
      </div> 
    </div>
  </div>
  <div class="rzz" id="rzz">
    <div class="rea box">
      <div class="reat">
        <h1 id="rwork"></h1>
      </div>
      <div class="comme">
       <div class="namee">  
        <div class="nxinfo"> 
         <div class="nipt"> 
          <p>昵称</p> 
          <input id="rname" type="text" placeholder="*必填,请输入您的昵称" /> 
         </div> 
         <div class="nipt"> 
          <p>邮箱</p> 
          <input id="remail" type="text" placeholder="推荐选填,请输入您的电子邮箱" /> 
         </div> 
         <div class="nipt"> 
          <p>QQ</p> 
          <input id="rqq" type="text" placeholder="选填,请输入您的联系QQ" /> 
         </div> 
         <div class="nipt"> 
          <p>网址</p> 
          <input id="rurl" type="text" placeholder="选填,请输入您的个人主页地址" /> 
         </div> 
        </div> 
       </div> 
       <div class="ctxe"> 
        <textarea id="rcoms" name="rcoms" placeholder="*必填,请输入您的精彩观点"></textarea> 
       </div> 
       <div class="sendc"> 
        <a class="rcls" href="javascript:clsComms()">取消</a> 
        <a id="repMethod" class="bgl" href="#">回复观点</a> 
       </div> 
      </div> 
    </div>
  </div>
  <div class="rzz" id="chelp">
    <div class="rea box chelp">
      <div class="reat">
        <h1>评论提示</h1>
      </div>
      <div class="helpe">
        <ul>
          <li>头像：系统为您提供了12个头像自由选择，初次打开随机为你选择一个</li>
          <li>邮箱：可选提交邮箱，该信息不会外泄，或将上线管理员回复邮件通知</li>
          <li>网址：可选提交网址，评论区该地址将以外链的形式展示在您的昵称上</li>
          <li>记忆：浏览器将记忆您已选择或填写过得信息，下次评论无需重复输入</li>
          <li>审核：提供一个和谐友善的评论环境，本站所有评论需要经过人工审核</li>
        </ul>
      </div>
      <div class="sendc">
       <a class="bgl" href="javascript:clsHelpComms()">了解</a> 
      </div>
    </div>
  </div>
  <?php } ?>
  </div> 
  <script type="text/javascript">
    <?php 
    if($pageType == "post" && $post != null){
      echo "var pid = ".$post["id"].";";
    }else{
      echo "var pid = 0;";
    }
    ?>
  </script>
  <script src="/static/js/base.js?v=1.0"></script> 
  <?php if($pageType == "index") { ?>
  <script src="/static/lib/svg/svg-inject.min.js?v=1.0"></script>
  <?php } ?>
  <?php if($pageType == "index") { ?>
  <script src="/static/lib/swiper/swiper-bundle.min.js?v=1.0"></script> 
  <script src="/static/js/index.js?v=1.0"></script>  
  <?php } ?>
  <?php if($pageType == "post") { ?>
   <script src="/static/lib/jquery/jquery-3.6.0.min.js?v=1.0"></script>
   <script src="/static/js/popup.js?v=1.0"></script>  
   <script src="/static/js/post.js?v=1.0"></script> 
   <script src="/static/js/mark.js?v=1.0"></script> 
  <?php } ?>
 </body>
</html>