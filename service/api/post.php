<?php
function UpsertPost($conn,$body) {
  $id = $body -> id;
  $title = $body -> title;
  $url = $body -> url;
  $banner = $body -> banner;
  $remark = $body -> remark;
  $cid = $body -> cid;
  $status = $body -> status;
  if ($id == "0") {
    $sql = "INSERT INTO `post_main`(`title`, `url`, `banner`, `remark`, `cid`, `status`) VALUES ('$title', '$url', '$banner', '$remark', '$cid', '$status')";
    $query = mysqli_query($conn,$sql);
    if($query) {
      $id = mysqli_insert_id($conn);
    } else {
      return [false,'文章主表添加失败'];
    }
  } else {
    $sql = "UPDATE `post_main` SET `title` = '$title', `url` = '$url', `banner` = '$banner', `remark` = '$remark', `cid` = '$cid', `status` = '$status' WHERE `id` = '$id'";
    $query = mysqli_query($conn,$sql);
    if(!$query) {
      return [false,'文章主表更新失败'];
    }
  }
  $md = $body -> md;
  $html = $body -> html;
  $menu = $body -> menu;
  $md = str_replace('\'','\'\'',$md);
  $html = str_replace('\'','\'\'',$html);
  $menu = str_replace('\'','\'\'',$menu);
  $sqlInfo = "REPLACE INTO `post_info`(`pid`,`md`, `html`, `menu`) VALUES ('$id','$md', '$html', '$menu')";
  $queryInfo = mysqli_query($conn,$sqlInfo);
  if(!$query) {
     return [false,'文章信息更新失败'];
  }
  // 更新文章标签
  $tids = $body ->tids;
  mysqli_query($conn,"DELETE FROM `post_tag` WHERE pid=".$id);
  // 关闭自动提交
  mysqli_autocommit($conn,FALSE);
  foreach($tids as $tid) {
    mysqli_query($conn,"INSERT INTO `post_tag`(`pid`, `tid`) VALUES ('$id','$tid')");
  }
  if(!mysqli_commit($conn)) {
    return [false,'文章标签更新失败'];
  }
  mysqli_autocommit($conn,TRUE);
  return [false,$id];
  // 更新TAG交给定时任务
}

function GetPost($conn,$id) {
  $sql = "SELECT * FROM `post_main` m LEFT JOIN `post_info` i ON m.id = i.pid WHERE m.id =".$id;
  $query = mysqli_query($conn,$sql);
  $post = [];
  if($query) {
    $post = mysqli_fetch_array($query,MYSQL_ASSOC);
  } else {
    return [false,'获取文章失败'];
  }
  // tagList
  $tag = "SELECT tid FROM `post_tag` t WHERE t.pid=".$id;
  $tquery = mysqli_query($conn,$tag);
  $tids = [];
  if($tquery) {
    while($tid = mysqli_fetch_row($query)) {
      array_push($tids,$tid[0]);
    }
  } else {
    return [false,'获取标签失败'];
  }
  $post['tids'] = $tids;
  return [true,$post];
}
?>