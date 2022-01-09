<?php
header('Access-Control-Allow-Headers: MTOKEN,Content-Type');
$token = $_SERVER['HTTP_MTOKEN'];//校验登陆合法性早-变_并追加HTTP
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/auth.php');
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
$api = $body -> api;
$ret = [];
include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/doConfig.php');
if($api == "GetConf") {
  $ret = [ true ,
    [
      'cdnUrl' => getConfig("cdnUrl","./comm/system.php","string"),
      'taskTime' => getConfig("taskTime","./comm/system.php","int"),
      'taskLastRun' => getConfig("taskLastRun","./comm/system.php","string"),
      'taskStatus' => getConfig("taskStatus","./comm/system.php","int"),
      'qqUrl' => getConfig("qqUrl","./comm/system.php","string"),
      'githubUrl' => getConfig("githubUrl","./comm/system.php","string"),
      'giteeUrl' => getConfig("giteeUrl","./comm/system.php","string"),
      'emailUrl' => getConfig("emailUrl","./comm/system.php","string")
      
    ]
  ];
}
if($api == "SetConf") {
  $cdnUrl = $body -> cdnUrl;
  $taskTime = $body -> taskTime;
  $qqUrl = $body -> qqUrl;
  $githubUrl = $body -> githubUrl;
  $giteeUrl = $body -> giteeUrl;
  $emailUrl = $body -> emailUrl;
  setConfig("cdnUrl",$cdnUrl,"./comm/system.php","string");
  setConfig("taskTime",$taskTime,"./comm/system.php","int");
  setConfig("qqUrl",$qqUrl,"./comm/system.php","string");
  setConfig("githubUrl",$githubUrl,"./comm/system.php","string");
  setConfig("giteeUrl",$giteeUrl,"./comm/system.php","string");
  setConfig("emailUrl",$emailUrl,"./comm/system.php","string");
  $ret = [ true , null ];
}
$code = $ret[0] ? 200 : 100;
$msg = $ret[0] ? '操作成功':$ret[1];
$data = $ret[0] ? $ret[1]: null;
echo json_encode(['code' => $code,'msg' => $msg ,'data' => $data]);
exit;
?>