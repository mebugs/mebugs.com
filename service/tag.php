<?php
header('Access-Control-Allow-Headers: MTOKEN,Content-Type');
$token = $_SERVER['HTTP_MTOKEN'];//校验登陆合法性早-变_并追加HTTP
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/auth.php');
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
$api = $body -> api;
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/connect.php');
include($_SERVER['DOCUMENT_ROOT'].'/service/api/tag.php');
$ret = [];
if($api == "AllTag") {
  $ret = GetTagList($conn);
}
if($api == "PageTag") {
  $query = [ 'name' => $body -> name,'size' => $body -> size,'page' => $body -> page ];
  $ret = GetTagPage($conn,$query);
}
if($api == "AddTag" || $api == "ModTag") {
  $id = $body -> id;
  $name = $body -> name;
  $url = $body -> url;
  $remark = $body -> remark;
  if($api == "AddTag") {
    $ret = AddTag($conn,$name,$url,$remark);
  } else if($api == "ModTag") {
    $ret = ModTag($conn,$id,$name,$url,$remark);
  } 
}
$code = $ret[0] ? 200 : 100;
$msg = $ret[0] ? '操作成功':$ret[1];
$data = $ret[0] ? $ret[1]: null;
mysqli_close($conn);
echo json_encode(['code' => $code,'msg' => $msg ,'data' => $data]);
exit;
?>