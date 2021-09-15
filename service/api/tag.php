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
  
}
?>