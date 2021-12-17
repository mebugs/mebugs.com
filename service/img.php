<?php
header('Access-Control-Allow-Headers: MTOKEN,Content-Type');
$token = $_SERVER['HTTP_MTOKEN'];//校验登陆合法性早-变_并追加HTTP
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/auth.php');
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
$api = $body -> api;
include($_SERVER['DOCUMENT_ROOT'].'/service/api/img.php');
$ret = [];
if($api == "Base64") {
  $imgs = $body -> imgs;
  $ret = UploadBaseArray($imgs);
}
if($api == "Base64Banner") {
  $imgs = $body -> imgs;
  $ret = UploadBaseArray($imgs,"banner/");
}
$code = $ret[0] ? 200 : 100;
$msg = $ret[0] ? '操作成功':$ret[1];
$data = $ret[0] ? $ret[1]: null;
echo json_encode(['code' => $code,'msg' => $msg ,'data' => $data]);
exit;
?>