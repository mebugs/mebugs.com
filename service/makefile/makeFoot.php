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
  <div class="back_top" onclick="backTop()"> 
  </div> 
  <script src="/static/js/base.js?v=1.0"></script> 
  <?php if($pageType == "index") { ?>
  <script src="/static/lib/svg/svg-inject.min.js?v=1.0"></script>
  <?php } ?>
  <?php if($pageType == "index") { ?>
  <script src="/static/lib/swiper/swiper-bundle.min.js?v=1.0"></script> 
  <script src="/static/js/index.js?v=1.0"></script>  
  <?php } ?>
 </body>
</html>