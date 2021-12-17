<?php
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/openCheck.php');
if (openCheckErr($body,3)) {
  echo json_encode(['code' => 100,'msg' => '非法请求' ]);
  exit;
}
include($_SERVER['DOCUMENT_ROOT'].'/service/api/img.php');
$ret = [];
$imgs = $body -> imgs;
$ret = UploadBaseArray($imgs,"firend/");
$code = $ret[0] ? 200 : 100;
$msg = $ret[0] ? '操作成功':$ret[1];
$data = $ret[0] ? $ret[1]: null;
echo json_encode(['code' => $code,'msg' => $msg ,'data' => $data]);
exit;
?>