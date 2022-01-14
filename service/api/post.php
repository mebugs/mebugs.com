<?php
function UpsertPost($conn,$body) {
  $id = $body -> id;
  $title = $body -> title;
  $url = $body -> url;
  $banner = $body -> banner;
  $remark = $body -> remark;
  $cid = $body -> cid;
  $status = $body -> status;
  $openComms = $body -> openComms;
  if ($id == "0") {
    $sql = "INSERT INTO `post_main`(`title`, `url`, `banner`, `remark`, `cid`, `status`, `openComms`) VALUES ('$title', '$url', '$banner', '$remark', '$cid', '$status', '$openComms')";
    $query = mysqli_query($conn,$sql);
    if($query) {
      $id = mysqli_insert_id($conn);
    } else {
      return [false,'文章主表添加失败'];
    }
  } else {
    $sql = "UPDATE `post_main` SET `title` = '$title', `url` = '$url', `banner` = '$banner', `remark` = '$remark', `cid` = '$cid', `status` = '$status', `openComms` = '$openComms' WHERE `id` = '$id'";
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
  $tids = $body -> tids;
  mysqli_query($conn,"DELETE FROM `post_tag` WHERE pid=".$id);
  if ($tids != null && count($tids) >0) {
    // 关闭自动提交
    mysqli_autocommit($conn,FALSE);
    foreach($tids as $tid) {
      mysqli_query($conn,"INSERT INTO `post_tag`(`pid`, `tid`) VALUES ('$id','$tid')");
    }
    if(!mysqli_commit($conn)) {
      return [false,'文章标签更新失败'];
    }
    mysqli_autocommit($conn,TRUE);
  }
  return [true,$id];
  // 更新TAG交给定时任务
}

function GetPost($conn,$id) {
  $sql = "SELECT * FROM `post_main` m LEFT JOIN `post_info` i ON m.id = i.pid WHERE m.id =".$id;// m.id > 100 AND
  $query = mysqli_query($conn,$sql);
  $post = [];
  if($query) {
    $post = mysqli_fetch_array($query,MYSQLI_ASSOC);
  } else {
    return [false,'获取文章失败'];
  }
  // tagList
  $tag = "SELECT tid FROM `post_tag` t WHERE t.pid=".$id;
  $tquery = mysqli_query($conn,$tag);
  $tids = [];
  if($tquery) {
    while($tid = mysqli_fetch_row($tquery)) {
      array_push($tids,$tid[0]);
    }
  } else {
    return [false,'获取标签失败'];
  }
  $post['tids'] = $tids;
  return [true,$post];
}

function ListPost($conn,$q) {
  $ctSql = "SELECT count(id) FROM `post_main` m ";# WHERE m.id > 100
  $qrSql = "SELECT m.*,c.name AS category FROM `post_main` m LEFT JOIN `category` c ON c.id = m.cid WHERE 1=1";// m.id > 100
  $where = "";
  if ($q['cid'] != 0) {
    $where = $where . " AND m.cid = " . $q['cid'];
  }
  if ($q['status'] != 0) {
    $where = $where . " AND m.status = " . $q['status'];
  }
  if ($q['title'] != "") {
    $where = $where . " AND m.title like '%" . $q['title'] . "%'";
  }
  $ctSql = $ctSql . $where;
  $size = $q['size'];
  $qrSql = $qrSql . $where . " ORDER BY `id` DESC LIMIT ".($q['page']-1)*$size.",".$size;
  $pageData = [];
  $numAry = mysqli_fetch_array(mysqli_query($conn,$ctSql));
  $total = $numAry[0];
  $pageData['pages'] = ceil($total/$size);
  $query = mysqli_query($conn,$qrSql);
  $pageData['list'] = mysqli_fetch_all($query,MYSQLI_ASSOC);
  return [true,$pageData];
}
?>