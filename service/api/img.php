<?php
// 默认上传到Post
function UploadBaseArray($imgs,$path="post/") {
  $fileName = $imgs[0];
  $upPath = $_SERVER['DOCUMENT_ROOT'].'/static/upload/'.$path;
  $imgx = array();
  // 单图与多图
  if(count($imgs) > 2) {
    array_push($imgx,$imgs[1],$imgs[2],$imgs[3]);
  } else {
    array_push($imgx,$imgs[1]);
  }
  for($i=0;$i<count($imgx);$i++)
  {
    if (strstr($imgx[$i],","))
    {
      $value = explode(',',$imgx[$i]);
      $imgx[$i] = $value[1];
    }
    $upPathName = $upPath . $fileName;
    if($i == 1)
    {
      $upPathName = $upPath . '400_' . $fileName;
    }
    if($i == 2)
    {
      $upPathName = $upPath . '100_' . $fileName;
    }
    file_put_contents($upPathName, base64_decode($imgx[$i]));
  }
  return [true,'/static/upload/'. $path . $fileName];
}
?>