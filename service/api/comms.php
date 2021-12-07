<?php
function AddComms($conn,$body) {
  $name = $body -> name;$name = str_replace('\'','\'\'',$name);
  $email = $body -> email;$email = str_replace('\'','\'\'',$email);
  $qq = $body -> qq;$qq = str_replace('\'','\'\'',$qq);
  $url = $body -> url;$url = str_replace('\'','\'\'',$url);
  $avt = $body -> avt;$avt = str_replace('\'','\'\'',$avt);
  $pid = $body -> pid;
  $fid = $body -> fid;
  $level = $body -> level;
  $coms = $body -> coms;$coms = str_replace('\'','\'\'',$coms);
  $sql = "INSERT INTO `comms`(`name`, `email`, `qq`, `url`, `avt`, `pid`, `fid`, `level`, `coms`, `send_time`) VALUES ('$name', '$email', '$qq', '$url', '$avt', '$pid', '$fid', '$level', '$coms', now())";
  $query = mysqli_query($conn,$sql);
  if($query) {
    $id = mysqli_insert_id($conn);
    return [true,$id];
  }
  return [false,'评论提交失败'];
}

// 获取评论管理员分页（级联文章）
function GetCommsPageManage($conn,$q) {
  $ctSql = "SELECT count(id) FROM `comms` s";
  $qrSql = "SELECT s.*,m.title,m.url AS purl,p.`name` AS fname,p.`coms` AS fcoms FROM `comms` s LEFT JOIN `post_main` m ON m.id = s.pid LEFT JOIN `comms` p ON p.fid = s.id";
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


function GetAllTag($conn) {
  return GetTagList($conn);
}

function GetTagList($conn,$page=0,$size=0) {
  $sql = "SELECT * FROM `tag` ORDER BY num DESC";
  // 分页查询请求
  if($page!=0){
    $sql = $sql." LIMIT ".($page-1)*$size.",".$size;
  }
  $query = mysqli_query($conn,$sql);
  if($query)
  {
  	return [true,mysqli_fetch_all($query,MYSQLI_ASSOC)];
  }
  return [false,'数据查询失败'];
}


function ModTag($conn,$id,$name,$url,$remark) {
  $sql = "UPDATE `tag` SET `name` = '$name', `url` = '$url', `remark` = '$remark'WHERE `id` = ".$id;
  $query = mysqli_query($conn,$sql);
  if($query) {
    return [true,$id];
  }
  return [false,'数据更新失败'];
}


?>