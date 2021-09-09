<?php
//header('Access-Control-Allow-Origin: *');
//用户登录 - 单用户暂不设计数据库方式鉴权
$name = $_GET['username'];
$pwd = $_GET['password'];
$code = 200;
$msg = '操作成功';
$data = [];
if($name == null || $pwd == null)//0 false null 空都为null
{
    $code = 100;
    $msg = '非法请求';
}
if($code == 200)
{
  include($_SERVER['DOCUMENT_ROOT'].'/service/comm/doConfig.php');
	if($name == getConfig("adminName") && $pwd == getConfig("adminPwd"))
	{
		//密码正确，生成Token，并存放至本地用于快速校验登陆
		$token = "";
		for ($i = 0; $i < 32; $i++) 
		{ 
		    $token .= chr(mt_rand(97,122)); 
		}
		setConfig("adminToken",$token);
		$data = $token;
	}else{
    $code = 100;
		$msg = '账号或密码不正确';
	}
}
echo json_encode(['code' => $code,'msg' => $msg ,'data' => $data]);
exit;
?>