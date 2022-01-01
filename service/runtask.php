<?php
header('Access-Control-Allow-Headers: MTOKEN,Content-Type');
$token = $_SERVER['HTTP_MTOKEN'];//校验登陆合法性早-变_并追加HTTP
include($_SERVER['DOCUMENT_ROOT'].'/service/comm/auth.php');
//界面数据JSON对象接受
$body = json_decode(file_get_contents("php://input"));
$ret = [];
include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/doConfig.php');
setConfig("taskStatus",1,"./comm/system.php","int");
// 返回结果
ob_end_clean();
ob_start();
echo json_encode(['code' => 200,'msg' => 'OK' ,'data' => 'ok']);

// Windows服务器需要加上这行。
echo str_repeat(" ",4096);
$size = ob_get_length();
// 发送请求的相应返回给来源请求
header("Content-Length: $size");
header('Connection: close');
header("HTTP/1.1 200 OK");
header("Content-Type: application/json;charset=utf-8");
ob_end_flush();
if(ob_get_length())
	ob_flush();
flush();
if (function_exists("fastcgi_finish_request")) { // yii或yaf默认不会立即输出，加上此句即可（前提是用的fpm）
    fastcgi_finish_request(); // 响应完成, 立即返回到前端,关闭连接
}
//在关闭连接后，继续运行php脚本
ignore_user_abort(true);
//不设置超时时间（根据实际情况使用）
set_time_limit(0); 
//继续运行的代码
// 启动定询任务
include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/runtask.php');
autoRun();
exit;
?>