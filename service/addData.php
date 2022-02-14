<?php
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/openCheck.php');
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/connect.php');
$pid = $body -> pid;
$api = $body -> api;
if($api != "" && $pid != "" && $pid > 999) {
  $sql = "UPDATE `post_main` SET `view` = `view`+1, `monthView` = `monthView`+1 WHERE `id` = ".$pid;
  if($api == "g") {
    $sql = "UPDATE `post_main` SET `depth` = `depth`+1 WHERE `id` = ".$pid;
  }
  mysqli_query($conn,$sql);
}
mysqli_close($conn);
exit;
?>