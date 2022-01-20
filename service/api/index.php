<?php
// 获取首页数据
function GetIndexData($conn) {
  // 统计数据
  $postCount = "SELECT count(id) FROM `post_main`";
  $categoryCount = "SELECT count(id) FROM `category`";
  $tagCount = "SELECT count(id) FROM `tag`";
  $commsCount = "SELECT count(id) FROM `comms`";
  $friendCount = "SELECT count(id) FROM `friend`";
  // 草稿
  $postDOCount = "SELECT count(id) FROM `post_main` WHERE status = 0";
  // 待审核评论
  $commsDOCount = "SELECT count(id) FROM `comms` WHERE status = 0";
  // 待审核友链
  $friendDOCount = "SELECT count(id) FROM `friend` WHERE status = 0";
  $data = [
    GetCount($postCount,$conn),
    GetCount($categoryCount,$conn),
    GetCount($tagCount,$conn),
    GetCount($commsCount,$conn),
    GetCount($friendCount,$conn),
    GetCount($postDOCount,$conn),
    GetCount($commsDOCount,$conn),
    GetCount($friendDOCount,$conn)
  ];
  return [true,$data];
}

function GetCount($sql,$conn) {
  $numAry = mysqli_fetch_array(mysqli_query($conn,$sql));
  $total = $numAry[0];
  return $total;
}

?>