    <!-- page common controller -->
    <div class="c c_12">
     <div class="peare">
      <?php 
      if($pages <= 7) {
        for($ro = 1;$ro <= $pages;$ro++) {
          $riclass='';
          if($ro == $page) {
            $riclass='class="pnow"';
          }
          echo '<a href="'.$pathDir.$pathBe.'_'.$ro.'.html" '.$riclass.'>'.$ro.'</a>';
        }
      }
      if($pages > 7 && $page < 5) {
        for($ro = 1;$ro <= 5;$ro++) {
          $riclass='';
          if($ro == $page) {
            $riclass='class="pnow"';
          }
          echo '<a href="'.$pathDir.$pathBe.'_'.$ro.'.html" '.$riclass.'>'.$ro.'</a>';
        }
        echo '<a href="'.$pathDir.$pathBe.'_6.html">...</a>';
        echo '<a href="'.$pathDir.$pathBe.'_'.$pages.'.html">'.$pages.'</a>';
      }
      if($pages > 7 && $page >= 5 && $page <= $pages - 4) {
        echo '<a href="'.$pathDir.$pathBe.'_1.html">1</a>';
        echo '<a href="'.$pathDir.$pathBe.'_'.($page-2).'.html">...</a>';
        echo '<a href="'.$pathDir.$pathBe.'_'.($page-1).'.html">'.($page-1).'</a>'; 
        echo '<a href="'.$pathDir.$pathBe.'_'.$page.'.html" class="pnow">'.$page.'</a>'; 
        echo '<a href="'.$pathDir.$pathBe.'_'.($page+1).'.html">'.($page+1).'</a>';
        echo '<a href="'.$pathDir.$pathBe.'_'.($page+2).'.html">...</a>'; 
        echo '<a href="'.$pathDir.$pathBe.'_'.$pages.'.html">'.$pages.'</a>';
      }
      if($pages > 7 && $page >= 5 && $page > $pages - 4) {
        echo '<a href="'.$pathDir.$pathBe.'_1.html">1</a>';
        echo '<a href="'.$pathDir.$pathBe.'_'.($page-5).'.html">...</a>';
        for($ro = ($pages - 4);$ro <= $pages;$ro++) {
          $riclass='';
          if($ro == $page) {
            $riclass='class="pnow"';
          }
          echo '<a href="'.$pathDir.$pathBe.'_'.$ro.'.html" '.$riclass.'>'.$ro.'</a>';
        }
      }
      ?>
     </div>
    </div>