<?php
function UpsertCategory($conn,$body) {
  $id = $body -> id;
  $name = $body -> name;
  $url = $body -> url;
  $icon = $body -> icon;
  $remark = $body -> remark;
  if ($id == "0") {
    $sql = "INSERT INTO `category`(`name`, `url`, `icon`, `remark`) VALUES ('$name', '$url', '$icon', '$remark')";
    $query = mysqli_query($conn,$sql);
    if($query) {
      $id = mysqli_insert_id($conn);
    } else {
      return [false,'分类添加失败'];
    }
  } else {
    $sql = "UPDATE `category` SET `name` = '$name', `url` = '$url', `icon` = '$icon', `remark` = '$remark' WHERE `id` = '$id'";
    $query = mysqli_query($conn,$sql);
    if(!$query) {
      return [false,'分类更新失败'];
    }
  }
  return [true,$id];
}

function ListCategory($conn) {
  $sql = "SELECT * FROM `category`";
  $query = mysqli_query($conn,$sql);
  $data = mysqli_fetch_all($query,MYSQLI_ASSOC);
  return [true,$data];
}
?>