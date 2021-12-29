<?php
header('Access-Control-Allow-Headers: MTOKEN,Content-Type');
$token = $_SERVER['HTTP_MTOKEN'];//校验登陆合法性早-变_并追加HTTP
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/auth.php');
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
$api = $body -> api;
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/connect.php');
include($_SERVER['DOCUMENT_ROOT'].'/service/api/firend.php');
$ret = [];
if($api == "PageFirends") {
  $query = [ 'size' => $body -> size,'page' => $body -> page ];
  $ret = GetFirendsPageManage($conn,$query);
}
if($api == "ModFirend") {
  $update = [ 'id' => $body -> id,'furl' => $body -> furl,'fname' => $body -> fname,'ficon' => $body -> ficon,'fdesc' => $body -> fdesc,'reson' => $body -> reson,'status' => $body -> status ];
  $ret = ModFirendsManage($conn,$update);
}
if($api == "AdminReplayComms") {
  // 添加回复评论
  $ret = AddComms($conn,$body,2,1);
  if($ret[0]) { // 更新回复评论状态
    $ret = UpdateCommsStatus($conn,$body -> fid,2);
  }
}
$code = $ret[0] ? 200 : 100;
$msg = $ret[0] ? '操作成功':$ret[1];
$data = $ret[0] ? $ret[1]: null;
mysqli_close($conn);
echo json_encode(['code' => $code,'msg' => $msg ,'data' => $data]);
exit;
?>