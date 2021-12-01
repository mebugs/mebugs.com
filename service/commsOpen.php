<?php
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/openCheck.php');
if (openCheckErr($body,1)) {
  echo json_encode(['code' => 100,'msg' => '非法请求' ]);
  exit;
}
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/connect.php');
mysqli_close($conn);
$code = 200;
$msg = "提交成功！";
// 添加评论
echo json_encode(['code' => $code,'msg' => $msg ]);
exit;
?>