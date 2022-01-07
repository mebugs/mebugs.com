<?php
function UpsertUrls($conn,$urls) {
  if ($urls['id'] == "0") {
    $sql = "INSERT INTO `urls`(`type`, `img`, `url`, `sorts`, `status`) VALUES ( '".$urls['type']."', '".$urls['img']."', '".$urls['url']."', ".$urls['sorts'].", ".$urls['status'].")";
    $query = mysqli_query($conn,$sql);
    if($query) {
      $id = mysqli_insert_id($conn);
    } else {
      return [false,'添加失败'];
    }
  } else {
    $sql = "UPDATE `urls` SET `type` = '".$urls['type']."', `img` = '".$urls['img']."', `url` = '".$urls['url']."',`sorts` = '".$urls['sorts']."',`status` = '".$urls['status']."' WHERE `id` = ".$urls['id'];
    $query = mysqli_query($conn,$sql);
    if(!$query) {
      return [false,'更新失败'];
    }
  }
  return [true,$urls['id']];
}

function GetUrlsPage($conn,$q) {
  $ctSql = "SELECT count(id) FROM `urls` m WHERE 1=1";
  $qrSql = "SELECT m.* FROM `urls` m WHERE 1=1";
  $size = $q['size'];
  $qrSql = $qrSql  . " ORDER BY `id` DESC LIMIT ".($q['page']-1)*$size.",".$size;
  $pageData = [];
  $numAry = mysqli_fetch_array(mysqli_query($conn,$ctSql));
  $total = $numAry[0];
  $pageData['pages'] = ceil($total/$size);
  $query = mysqli_query($conn,$qrSql);
  $pageData['list'] = mysqli_fetch_all($query,MYSQLI_ASSOC);
  return [true,$pageData];
}
?>