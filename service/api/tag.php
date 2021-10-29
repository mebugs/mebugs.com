<?php
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

function AddTag($conn,$name,$url,$remark) {
  $sql = "INSERT INTO `tag`(`name`, `url`, `remark`) VALUES ('$name', '$url', '$remark')";
  $query = mysqli_query($conn,$sql);
  if($query) {
    $id = mysqli_insert_id($conn);
    return [true,$id];
  }
  return [false,'数据新增失败'];
}

function ModTag($conn,$id,$name,$url,$remark) {
  $sql = "UPDATE `tag` SET `name` = '$name', `url` = '$url', `remark` = '$remark'WHERE `id` = ".$id;
  $query = mysqli_query($conn,$sql);
  if($query) {
    return [true,$id];
  }
  return [false,'数据更新失败'];
}

function GetTagPage($conn,$q) {
  $ctSql = "SELECT count(id) FROM `tag` m WHERE 1=1";
  $qrSql = "SELECT m.* FROM `tag` m WHERE 1=1";
  $where = "";
  if ($q['name'] != "") {
    $where = $where . " AND m.name like '%" . $q['name'] . "%'";
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