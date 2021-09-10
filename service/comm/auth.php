<?php
//include(__DIR__.'/connect.php'); 
//校验TOKEN - 通过快速Token鉴权，暂不通过数据库
include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/doConfig.php'); 
$servToken = getConfig("adminToken");
if($servToken != $token) {
	echo json_encode(['code' => 50008,'msg' => 'TOKEN非法' ,'data' => null]);
	exit;
}
?>