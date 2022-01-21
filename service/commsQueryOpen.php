<?php
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/openCheck.php');
if (openCheckErr($body,4)) {
  echo json_encode(['code' => 100,'msg' => '非法请求' ]);
  exit;
}
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/connect.php');
include($_SERVER['DOCUMENT_ROOT'].'/service/api/comms.php');
// 添加评论
$ret = PageComms($conn,$body);
$code = $ret[0] ? 200 : 100;
$msg = $ret[0] ? '获取成功':$ret[1];
$data = $ret[0] ? $ret[1]: null;
mysqli_close($conn);
echo json_encode(['code' => $code,'msg' => $msg ,'data' => $data]);
exit;
?>