<?php
function AddComms($conn,$body,$status,$admin) {
  $name = $body -> name;$name = str_replace('\'','\'\'',$name);
  $email = $body -> email;$email = str_replace('\'','\'\'',$email);
  $qq = $body -> qq;$qq = str_replace('\'','\'\'',$qq);
  $url = $body -> url;$url = str_replace('\'','\'\'',$url);
  $avt = $body -> avt;$avt = str_replace('\'','\'\'',$avt);
  $pid = $body -> pid;
  $fid = $body -> fid;
  $level = $body -> level;
  $coms = $body -> coms;$coms = str_replace('\'','\'\'',$coms);
  $sql = "INSERT INTO `comms`(`name`, `email`, `qq`, `url`, `avt`, `pid`, `fid`, `level`, `coms`, `send_time`, `status`, `admin`) VALUES ('$name', '$email', '$qq', '$url', '$avt', '$pid', '$fid', '$level', '$coms', now(), '$status', '$admin')";
  $query = mysqli_query($conn,$sql);
  if($query) {
    $id = mysqli_insert_id($conn);
    return [true,$id];
  }
  return [false,'评论提交失败'];
}

// 获取评论管理员分页（级联文章）
function GetCommsPageManage($conn,$q) {
  $ctSql = "SELECT count(id) FROM `comms` s WHERE s.`admin` = 0";
  $qrSql = "SELECT s.*,m.title,m.url AS purl,p.`name` AS fname,p.`coms` AS fcoms, ap.`coms` AS acoms,IFNULL(ap.`id`,0) AS aid FROM `comms` s LEFT JOIN `post_main` m ON m.id = s.pid LEFT JOIN `comms` p ON s.fid = p.id LEFT JOIN (SELECT a.* FROM `comms` a WHERE a.`admin` = 1) ap ON s.id = ap.fid WHERE s.`admin` = 0";
  $where = "";
  $ctSql = $ctSql . $where;
  $size = $q['size'];
  $qrSql = $qrSql . $where . " ORDER BY s.`status`,s.`id` DESC LIMIT ".($q['page']-1)*$size.",".$size;
  $pageData = [];
  $numAry = mysqli_fetch_array(mysqli_query($conn,$ctSql));
  $total = $numAry[0];
  $pageData['pages'] = ceil($total/$size);
  $query = mysqli_query($conn,$qrSql);
  $pageData['list'] = mysqli_fetch_all($query,MYSQLI_ASSOC);
  return [true,$pageData];
}

// 更新审核评论
function ModCommsPageManage($conn,$update) {
  $sql = "";
  if($update['status'] == 3) {
    $sql = $sql."UPDATE `comms` SET `status` = '".$update['status']."' WHERE `id` = ".$update['id'];
  } else {
    $sql = $sql."UPDATE `comms` SET `name` = '".$update['name']."', `email` = '".$update['email']."',`qq` = '".$update['qq']."',`url` = '".$update['url']."',`coms` = '".$update['coms']."',`status` = '".$update['status']."' WHERE `id` = ".$update['id'];
  }
  $query = mysqli_query($conn,$sql);
  if($query) {
    return [true,$update['id']];
  }
  return [false,'数据更新失败'];
}

function UpdateCommsStatus($conn,$id,$status) {
  $sql = "UPDATE `comms` SET `status` = '$status' WHERE `id` = ".$id;
  $query = mysqli_query($conn,$sql);
  if($query) {
    return [true,$id];
  }
  return [false,'数据更新失败'];
}

?>