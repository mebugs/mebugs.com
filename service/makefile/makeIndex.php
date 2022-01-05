<?php 
//生成INDEX页面代码 CDN切换 数据库对象
function makeIndex($cdnUrl,$baseUrl.,$file_sc = $_SERVER['DOCUMENT_ROOT']."/index.html";//资源文件$conn) {
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
  // 查询数据
  
?>