<?php
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/openCheck.php');
if (openCheckErr($body,2)) {
  echo json_encode(['code' => 100,'msg' => '非法请求' ]);
  exit;
}
$api = $body -> api;
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/connect.php');
include($_SERVER['DOCUMENT_ROOT'].'/service/api/firend.php');
// 读取站点数据
if($api == "CheckFirend") {
  $ret = CheckFirend($body -> furl);
}
$code = $ret[0] ? 200 : 100;
$msg = $ret[0] ? '读取成功':$ret[1];
$data = $ret[1] ? $ret[1]:null;
mysqli_close($conn);
echo json_encode(['code' => $code,'msg' => $msg, 'data' => $data ]);
exit;
?>