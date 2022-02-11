<?php 

function makeSiteMap($siteMaps) {
  // 启动缓存输出 采用模版形式无需缓冲区
  ob_start();
  // 目标文件
  $file_sc = $_SERVER['DOCUMENT_ROOT']."/sitemap.xml";
?>
<?xml version="1.0" encoding="utf-8"?>
<urlset>
    <?php foreach ($siteMaps as $site) { ?>
    <url>
        <loc><?php echo $site['loc']; ?></loc>
        <changefreq>daily</changefreq>
        <priority><?php echo $site['priority']; ?></priority>
    </url>
    <?php } ?>
</urlset>
<?php
  //页面绘制完成
	//存储新的资源数据
	$content = ob_get_contents();
	//写入文件
	file_put_contents($file_sc,$content);
	//擦除缓冲区并关闭
	ob_end_clean();
}
?>